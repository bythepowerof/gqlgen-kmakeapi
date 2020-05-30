package gqlgen_kmakeapi

import (
	"context"
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
	"github.com/gorilla/handlers"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
)

const (
	APP_KEY = "My Secret"
)

type ServerOpts struct {
	Namespace string
	Port      string
	Trace     bool
	Jwt       bool
}

func AuthMiddleware(next http.Handler) http.Handler {
	if len(APP_KEY) == 0 {
		log.Fatal("HTTP server unable to start, expected an APP_KEY for JWT auth")
	}
	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(APP_KEY), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
		Debug:         true,
	})
	return jwtMiddleware.Handler(next)
}

func RealHTTPServer(c client.Client, opts ServerOpts) {

	cl := cors.Default()

	kc := controller.NewKubernetesController(c, opts.Namespace)

	kc.AddListener()
	kc.GetListener().KmakeChanges(opts.Namespace)

	srv := handler.New(NewExecutableSchema(Config{
		Resolvers: &Resolver{
			KmakeController: kc,
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
		http.Handle("/query", AuthMiddleware(handlers.LoggingHandler(os.Stdout, cl.Handler(srv))))
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
				controller.NewKubernetesController(c, "all"),
			},
		},
	)))
}
