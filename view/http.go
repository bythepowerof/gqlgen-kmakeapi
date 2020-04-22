package gqlgen_kmakeapi

import (
	"log"
	"net/http"
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
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

func RealHTTPServer(c client.Client, m manager.Manager, namespace string, port string) {
	cl := cors.Default()

	kc := controller.NewKubernetesController(c, namespace)

	kc.AddListener(m)
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

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", cl.Handler(srv))

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
