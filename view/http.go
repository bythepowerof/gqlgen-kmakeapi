package gqlgen_kmakeapi

import (
	"context"
	"fmt"
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

func RealHTTPServer(c client.Client, namespace string, port string, trace bool) {
	cl := cors.Default()

	kc := controller.NewKubernetesController(c, namespace)

	kc.AddListener()
	kc.GetListener().KmakeChanges(namespace)

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

	if trace {
		srv.AroundFields(func(ctx context.Context, next graphql.Resolver) (res interface{}, err error) {
			rc := graphql.GetFieldContext(ctx)
			fmt.Println("Entered", rc.Object, rc.Field.Name)
			res, err = next(ctx)
			fmt.Println("Left", rc.Object, rc.Field.Name, "=>", res, err)
			return res, err
		})
	}

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", handlers.LoggingHandler(os.Stdout, cl.Handler(srv)))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func FakeHTTPServer(c k8sclient.Client) *gclient.Client {
	return gclient.New(handler.NewDefaultServer(NewExecutableSchema(
		Config{
			Resolvers: &Resolver{
				controller.NewKubernetesController(c, "all"),
				// controller.NewKubernetesController(c, nil, "all"),
			},
		},
	)))
}
