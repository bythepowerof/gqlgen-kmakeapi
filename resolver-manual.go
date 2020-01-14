package gqlgen_todos

//go:generate go run github.com/99designs/gqlgen

import (
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Resolver struct {
	todos []*Todo
	Client client.Client
}
