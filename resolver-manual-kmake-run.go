package gqlgen_todos

import (
	"context"

	"github.com/bythepowerof/kmake-controller/api/v1"
	// v11 "k8s.io/api/core/v1"
	// 	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (r *kmakeRunResolver) Runstatus(ctx context.Context, obj *v1.KmakeRun) (*v1.KmakeRunStatus, error) {
	return &obj.Status, nil
}
