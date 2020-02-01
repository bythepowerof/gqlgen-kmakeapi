package gqlgen_kmakeapi

//go:generate go run github.com/99designs/gqlgen
import (
	"context"
	"github.com/bythepowerof/gqlgen-kmakeapi/controller"
	"github.com/bythepowerof/kmake-controller/api/v1"
	// "github.com/bythepowerof/kmake-controller/gql"
	"strconv"
)

func (r *mutationResolver) Reset(ctx context.Context, input controller.NewReset) (*v1.KmakeScheduleRun, error) {
	opt := make(map[string]string)
	opt["full"] = strconv.FormatBool(input.Full)
	op := controller.RunTypeReset

	return r.KmakeController.CreateScheduleRun(ctx, input.Namespace, nil, nil, &input.Kmakescheduler, &op, opt)
	// func CreateScheduleRun(ctx context.Context, namespace string, kmake *string, kmakerun *string, kmakescheduler *string, runtype *RunType, opts *map[string]string) (*v1.KmakeScheduleRun, error) {

}
