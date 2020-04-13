package gqlgen_kmakeapi

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import ()

// Kmake returns KmakeResolver implementation.
func (r *Resolver) Kmake() KmakeResolver { return &kmakeResolver{r} }

// KmakeNowScheduler returns KmakeNowSchedulerResolver implementation.
func (r *Resolver) KmakeNowScheduler() KmakeNowSchedulerResolver { return &kmakeNowSchedulerResolver{r} }

// KmakeRun returns KmakeRunResolver implementation.
func (r *Resolver) KmakeRun() KmakeRunResolver { return &kmakeRunResolver{r} }

// KmakeRunJob returns KmakeRunJobResolver implementation.
func (r *Resolver) KmakeRunJob() KmakeRunJobResolver { return &kmakeRunJobResolver{r} }

// KmakeScheduleRun returns KmakeScheduleRunResolver implementation.
func (r *Resolver) KmakeScheduleRun() KmakeScheduleRunResolver { return &kmakeScheduleRunResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Namespace returns NamespaceResolver implementation.
func (r *Resolver) Namespace() NamespaceResolver { return &namespaceResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type kmakeResolver struct{ *Resolver }
type kmakeNowSchedulerResolver struct{ *Resolver }
type kmakeRunResolver struct{ *Resolver }
type kmakeRunJobResolver struct{ *Resolver }
type kmakeScheduleRunResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type namespaceResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
