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

	return r.KmakeController.CreateScheduleRun(input.Namespace, nil, nil, &input.Kmakescheduler, &op, opt)
}

func (r *mutationResolver) Stop(ctx context.Context, input controller.RunLevelIn) (*v1.KmakeScheduleRun, error) {
	opt := make(map[string]string)
	op := controller.RunTypeStop

	return r.KmakeController.CreateScheduleRun(input.Namespace, nil, &input.Kmakerun, &input.Kmakescheduler, &op, opt)
}

func (r *mutationResolver) Restart(ctx context.Context, input controller.RunLevelIn) (*v1.KmakeScheduleRun, error) {
	opt := make(map[string]string)
	op := controller.RunTypeRestart

	return r.KmakeController.CreateScheduleRun(input.Namespace, nil, &input.Kmakerun, &input.Kmakescheduler, &op, opt)
}
