package main

import (
	"log"
	"net/http"
	"os"
	"fmt"

	"github.com/99designs/gqlgen/handler"
	gqlgen_todos "github.com/bythepowerof/gqlgen-kmakeapi"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	bythepowerofv1 "github.com/bythepowerof/kmake-controller/api/v1"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/apimachinery/pkg/runtime"

)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	scheme   := runtime.NewScheme()
	_ = clientgoscheme.AddToScheme(scheme)
	_ = bythepowerofv1.AddToScheme(scheme)


	c, err := client.New(config.GetConfigOrDie(), client.Options{Scheme: scheme})
	if err != nil {
		fmt.Println("failed to create client")
		os.Exit(1)
	}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(gqlgen_todos.NewExecutableSchema(gqlgen_todos.Config{Resolvers: &gqlgen_todos.Resolver{Client: c}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
