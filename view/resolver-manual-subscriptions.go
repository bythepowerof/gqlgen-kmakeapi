package gqlgen_kmakeapi

//go:generate go run github.com/99designs/gqlgen
import (
	"context"
	"github.com/bythepowerof/gqlgen-kmakeapi/controller"
	// "github.com/bythepowerof/kmake-controller/api/v1"
	"github.com/bythepowerof/kmake-controller/gql"
	// "strconv"
)

func (r *subscriptionResolver) Changed(ctx context.Context, input *controller.SubNamespace) (<-chan gql.KmakeObject, error) {
	return r.KmakeController.GetListener().AddChangeClient(ctx, input.Namespace)
	// return nil, nil
}
