package gqlgen_kmakeapi

//go:generate go run github.com/99designs/gqlgen

import (
	"github.com/bythepowerof/gqlgen-kmakeapi/controller"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Resolver struct {
	Client          client.Client
	KmakeController controller.KmakeController
}
