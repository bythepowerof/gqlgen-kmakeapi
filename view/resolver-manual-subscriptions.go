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
	return r.KmakeController.AddChangeClient(ctx, input.Namespace)

	// messages := make(chan []gql.KmakeObject, 1)
	// r.mutex.Lock()
	// r.index++
	// r.changes[r.index] = messages
	// r.mutex.Unlock()

	// // Delete channel when done
	// go func() {
	// 	<-ctx.Done()
	// 	r.mutex.Lock()
	// 	delete(r.changes, r.index)
	// 	r.mutex.Unlock()
	// }()
	// return nil, nil
}
