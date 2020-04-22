package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/bythepowerof/gqlgen-kmakeapi/k8s"
	"github.com/bythepowerof/gqlgen-kmakeapi/view"
	bythepowerofv1 "github.com/bythepowerof/kmake-controller/api/v1"
	"github.com/namsral/flag"
	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

const defaultPort = "8080"

func processArgs(fakeK8sClient *bool, fakeHTTPServer *bool, port *string, namespace *string) {
	flag.BoolVar(fakeK8sClient, "fake-k8s", false, "Use fake k8s client")
	flag.BoolVar(fakeHTTPServer, "fake-http", false, "Use fake k8s server")
	flag.StringVar(port, "port", defaultPort, "Port to listen to")
	flag.StringVar(namespace, "namespace", "all",
		"Namespace to watch - use 'all' for all namespaces")
	flag.Parse()
}

func main() {
	var fakeK8sClient bool
	var fakeHTTPServer bool
	var port string
	var namespace string

	processArgs(&fakeK8sClient, &fakeHTTPServer, &port, &namespace)

	scheme := runtime.NewScheme()
	_ = clientgoscheme.AddToScheme(scheme)
	_ = bythepowerofv1.AddToScheme(scheme)

	var c client.Client
	var err error

	if fakeK8sClient {
		fo := &k8s.FakeObjects{}
		c, err = fo.FakeK8sClient()
	} else {
		c, err = k8s.RealK8sClient(config.GetConfigOrDie(), client.Options{Scheme: scheme})
	}

	if err != nil {
		fmt.Println("failed to create client")
		os.Exit(1)
	}

	mo := manager.Options{Scheme: scheme, MetricsBindAddress: "0"}

	if strings.ToLower(namespace) != "all" {
		mo.Namespace = namespace
	}

	m, err := k8s.RealK8sManager(config.GetConfigOrDie(), mo)
	if err != nil {
		fmt.Println("failed to create manager")
		os.Exit(1)
	}

	if fakeHTTPServer {
		gqlgen_kmakeapi.FakeHTTPServer(c)
	} else {
		gqlgen_kmakeapi.RealHTTPServer(c, m, namespace, port)
	}

	// log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	// log.Fatal(http.ListenAndServe(":"+port, nil))
}
