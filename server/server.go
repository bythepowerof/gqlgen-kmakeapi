package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/bythepowerof/gqlgen-kmakeapi/controller"
	"github.com/bythepowerof/gqlgen-kmakeapi/fake"
	"github.com/bythepowerof/gqlgen-kmakeapi/view"

	bythepowerofv1 "github.com/bythepowerof/kmake-controller/api/v1"
	"github.com/namsral/flag"
	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

const defaultPort = "8080"

func RealK8sClient(config *rest.Config, options client.Options) (client.Client, error) {
	return client.New(config, options)
}

func RealHTTPServer(c client.Client) {
	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(
		gqlgen_kmakeapi.NewExecutableSchema(
			gqlgen_kmakeapi.Config{
				Resolvers: &gqlgen_kmakeapi.Resolver{
					KmakeController: &controller.KubernetesController{Client: c},
				},
			},
		),
	),
	)
}

func main() {
	var fakeK8sClient bool
	var fakeHTTPServer bool
	var port string

	flag.BoolVar(&fakeK8sClient, "fake-k8s", false, "Use fake k8s client")
	flag.BoolVar(&fakeHTTPServer, "fake-http", false, "Use fake k8s server")
	flag.StringVar(&port, "port", defaultPort, "Port to listen to")

	flag.Parse()

	scheme := runtime.NewScheme()
	_ = clientgoscheme.AddToScheme(scheme)
	_ = bythepowerofv1.AddToScheme(scheme)

	var c client.Client
	var err error

	if fakeK8sClient {
		c, err = fake.FakeK8sClient(scheme)
	} else {
		c, err = RealK8sClient(config.GetConfigOrDie(), client.Options{Scheme: scheme})
	}

	if err != nil {
		fmt.Println("failed to create client")
		os.Exit(1)
	}

	if fakeHTTPServer {

	} else {
		RealHTTPServer(c)
	}

	// http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	// http.Handle("/query", handler.GraphQL(
	// 	gqlgen_kmakeapi.NewExecutableSchema(
	// 		gqlgen_kmakeapi.Config{
	// 			Resolvers: &gqlgen_kmakeapi.Resolver{
	// 				KmakeController: &controller.KubernetesController{Client: c},
	// 			},
	// 		},
	// 	),
	// ),
	// )

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
