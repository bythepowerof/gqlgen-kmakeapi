package gqlgen_kmakeapi

import (
	"context"
	"encoding/json"
	"fmt"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"os"
	"time"

	gclient "github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"

	// "github.com/99designs/gqlgen/handler"
	"github.com/bythepowerof/gqlgen-kmakeapi/controller"
	"sigs.k8s.io/controller-runtime/pkg/client"
	k8sclient "sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/JeremyMarshall/gqlgen-jwt/rbac/dummy"
	"github.com/JeremyMarshall/gqlgen-jwt/rbac/types"
	"github.com/gorilla/handlers"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

type ServerOpts struct {
	Namespace string
	Port      string
	Trace     bool
	Jwt       bool
}

func AuthMiddleware(next http.Handler, secret string) http.Handler {
	if len(secret) == 0 {
		log.Fatal("HTTP server unable to start, expected an APP_KEY for JWT auth")
	}
	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
		Debug:         true,
		// Set this to false if you always want a bearer token present
		CredentialsOptional: true,
		UserProperty:        JwtTokenField,
		ErrorHandler: func(w http.ResponseWriter, r *http.Request, err string) {
			data := gqlerror.Error{
				Message: fmt.Sprintf("JWT Auth %s", err),
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(data)
		},
	})
	return jwtMiddleware.Handler(next)
}

type User struct {
	User  string
	Roles []string
}

func GetCurrentUser(ctx context.Context) *User {
	if rawToken := ctx.Value(JwtTokenField); rawToken != nil {
		token := rawToken.(*jwt.Token)

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			u := &User{
				User:  claims["user"].(string),
				Roles: make([]string, 0),
			}
			for _, r := range claims["roles"].([]interface{}) {
				u.Roles = append(u.Roles, fmt.Sprint(r))
			}
			return u
		}
	}
	return &User{}
}

type RbacDomainMiddlewareFunc func(ctx context.Context, obj interface{}, next graphql.Resolver, rbac Rbac) (res interface{}, err error)

func RbacDomainMiddleware(rbacChecker types.Rbac) RbacDomainMiddlewareFunc {
	return func(ctx context.Context, obj interface{}, next graphql.Resolver, rbac Rbac) (res interface{}, err error) {

		if args, ok := obj.(map[string]interface{}); ok {
			if domain, ok := args["namespace"].(string); ok {
				user := GetCurrentUser(ctx)
				if rbacChecker.CheckDomain(user.Roles, &domain, rbac.String()) {
					return next(ctx)
				}
			}
		}
		return nil, fmt.Errorf("Access denied")
	}
}

func RealHTTPServer(c client.Client, opts ServerOpts) {

	cl := cors.Default()

	kc := controller.NewKubernetesController(c, opts.Namespace)

	kc.AddListener()
	kc.GetListener().KmakeChanges(opts.Namespace)

	rbac := &dummy.Dummy{}

	srv := handler.New(NewExecutableSchema(Config{
		Resolvers: &Resolver{
			KmakeController: kc,
			JwtSecret:       JwtSecret,
		},
		Directives: DirectiveRoot{
			HasRbacDomain: RbacDomainMiddleware(rbac),
		},
	}))

	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})
	srv.Use(extension.Introspection{})

	if opts.Trace {
		srv.AroundFields(func(ctx context.Context, next graphql.Resolver) (res interface{}, err error) {
			rc := graphql.GetFieldContext(ctx)
			fmt.Println("Entered", rc.Object, rc.Field.Name)
			res, err = next(ctx)
			fmt.Println("Left", rc.Object, rc.Field.Name, "=>", res, err)
			return res, err
		})
	}

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	if opts.Jwt {
		http.Handle("/query", AuthMiddleware(handlers.LoggingHandler(os.Stdout, cl.Handler(srv)), JwtSecret))
	} else {
		http.Handle("/query", handlers.LoggingHandler(os.Stdout, cl.Handler(srv)))
	}

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", opts.Port)
	log.Fatal(http.ListenAndServe(":"+opts.Port, nil))
}

func FakeHTTPServer(c k8sclient.Client) *gclient.Client {
	return gclient.New(handler.NewDefaultServer(NewExecutableSchema(
		Config{
			Resolvers: &Resolver{
				KmakeController: controller.NewKubernetesController(c, "all"),
				JwtSecret:       JwtSecret,
			},
		},
	)))
}
