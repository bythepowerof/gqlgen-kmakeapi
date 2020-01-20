package gqlgen_kmakeapi

//go:generate go run github.com/99designs/gqlgen

import (
	"github.com/bythepowerof/gqlgen-kmakeapi/controller"
)

type Resolver struct {
	KmakeController controller.KmakeController
}
