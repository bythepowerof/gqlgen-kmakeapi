package gqlgen_todos

// //go:generate go run github.com/99designs/gqlgen

import (
	context "context"
)

func (r *queryResolver) Namespaces(ctx context.Context) ([]*Namespace, error) {
	panic("not implemented")
}
