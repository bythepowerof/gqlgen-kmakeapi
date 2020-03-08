package gqlgen_kmakeapi

import ()

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.


func (r *Resolver) Kmake() KmakeResolver {
	return &kmakeResolver{r}
}
func (r *Resolver) KmakeNowScheduler() KmakeNowSchedulerResolver {
	return &kmakeNowSchedulerResolver{r}
}
func (r *Resolver) KmakeRun() KmakeRunResolver {
	return &kmakeRunResolver{r}
}
func (r *Resolver) KmakeRunJob() KmakeRunJobResolver {
	return &kmakeRunJobResolver{r}
}
func (r *Resolver) KmakeScheduleRun() KmakeScheduleRunResolver {
	return &kmakeScheduleRunResolver{r}
}
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Namespace() NamespaceResolver {
	return &namespaceResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type kmakeResolver struct{ *Resolver }


type kmakeNowSchedulerResolver struct{ *Resolver }


type kmakeRunResolver struct{ *Resolver }


type kmakeRunJobResolver struct{ *Resolver }


type kmakeScheduleRunResolver struct{ *Resolver }


type mutationResolver struct{ *Resolver }


type namespaceResolver struct{ *Resolver }


type queryResolver struct{ *Resolver }

