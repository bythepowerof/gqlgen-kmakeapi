package gqlgen_todos

import ()

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

func (r *Resolver) Kmake() KmakeResolver {
	return &kmakeResolver{r}
}
func (r *Resolver) KmakeRun() KmakeRunResolver {
	return &kmakeRunResolver{r}
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
func (r *Resolver) Todo() TodoResolver {
	return &todoResolver{r}
}

type kmakeResolver struct{ *Resolver }

type kmakeRunResolver struct{ *Resolver }

type mutationResolver struct{ *Resolver }

type namespaceResolver struct{ *Resolver }

type queryResolver struct{ *Resolver }

type todoResolver struct{ *Resolver }
