package gqlgen_kmakeapi

import (
	"net/http"
	"sync"
	"time"

	gclient "github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	// "github.com/99designs/gqlgen/handler"
	"github.com/bythepowerof/gqlgen-kmakeapi/controller"
	"sigs.k8s.io/controller-runtime/pkg/client"
	k8sclient "sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/bythepowerof/kmake-controller/gql"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
)

func RealHTTPServer(c client.Client) {
	cl := cors.Default()

	kc := &controller.KubernetesController{
		Client:  c,
		Mutex:   sync.Mutex{},
		Changes: map[int]chan gql.KmakeObject{},
	}

	kc.KmakeChanges("xxx")

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

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", cl.Handler(srv))

	http.ListenAndServe(":8080", nil)
}

func FakeHTTPServer(c k8sclient.Client) *gclient.Client {
	return gclient.New(handler.NewDefaultServer(NewExecutableSchema(
		Config{
			Resolvers: &Resolver{
				KmakeController: &controller.KubernetesController{
					Client:  c,
					Mutex:   sync.Mutex{},
					Changes: map[int]chan gql.KmakeObject{},
				},
			},
		},
	)))
}
