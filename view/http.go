
package gqlgen_kmakeapi


import (
	"net/http"

	"github.com/99designs/gqlgen/handler"
	"github.com/bythepowerof/gqlgen-kmakeapi/controller"
	k8sclient "sigs.k8s.io/controller-runtime/pkg/client"
	gclient "github.com/99designs/gqlgen/client"
	"sigs.k8s.io/controller-runtime/pkg/client"
	myhandler "github.com/99designs/gqlgen/graphql/handler"
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

func FakeHTTPServer(c k8sclient.Client) {
	gclient.New(myhandler.NewDefaultServer(NewExecutableSchema(
		Config{
			Resolvers: &Resolver{
				KmakeController: &controller.KubernetesController{
					Client: c,
				},
			},
		},
	)))
}