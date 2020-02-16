package gqlgen_kmakeapi

import (
	"net/http"

	gclient "github.com/99designs/gqlgen/client"
	myhandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/handler"
	"github.com/bythepowerof/gqlgen-kmakeapi/controller"
	"sigs.k8s.io/controller-runtime/pkg/client"
	k8sclient "sigs.k8s.io/controller-runtime/pkg/client"
)

func RealHTTPServer(c client.Client) {
	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(
		NewExecutableSchema(
			Config{
				Resolvers: &Resolver{
					KmakeController: &controller.KubernetesController{Client: c},
				},
			},
		),
	),
	)
}

func FakeHTTPServer(c k8sclient.Client) *gclient.Client {
	return gclient.New(myhandler.NewDefaultServer(NewExecutableSchema(
		Config{
			Resolvers: &Resolver{
				KmakeController: &controller.KubernetesController{
					Client: c,
				},
			},
		},
	)))
}
