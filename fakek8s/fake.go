package fakek8s

import (
	// "github.com/99designs/gqlgen/client"
	// "github.com/99designs/gqlgen/graphql/handler"
	// "github.com/bythepowerof/gqlgen-kmakeapi/controller"
	// "github.com/bythepowerof/gqlgen-kmakeapi/view"
	bythepowerofv1 "github.com/bythepowerof/kmake-controller/api/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	k8sclient "sigs.k8s.io/controller-runtime/pkg/client"
	k8sfakeclient "sigs.k8s.io/controller-runtime/pkg/client/fake"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"

)

func FakeK8sClient() (k8sclient.Client, error) {
	scheme := runtime.NewScheme()
	_ = clientgoscheme.AddToScheme(scheme)
	_ = bythepowerofv1.AddToScheme(scheme)


	kmake := &bythepowerofv1.Kmake{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "Kmake",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-kmake",
			Namespace: "ns1",
		},
	}

	return k8sfakeclient.NewFakeClientWithScheme(scheme,
		kmake,
		// kmakerun,
		// kmakeschedulerun,
		// kmakesnowscheduler,
		// namespace,
	), nil
}

// func FakeHTTPServer(c k8sclient.Client) {
// 	client.New(handler.NewDefaultServer(gqlgen_kmakeapi.NewExecutableSchema(
// 		gqlgen_kmakeapi.Config{
// 			Resolvers: &gqlgen_kmakeapi.Resolver{
// 				KmakeController: &controller.KubernetesController{
// 					Client: c,
// 				},
// 			},
// 		},
// 	)))
// }
