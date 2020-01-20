package gqlgen_kmakeapi

import ()

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

func (r *Resolver) Kmake() KmakeResolver {
	return &kmakeResolver{r}
}
func (r *Resolver) KmakeRun() KmakeRunResolver {
	return &kmakeRunResolver{r}
}
func (r *Resolver) KmakeRunDummy() KmakeRunDummyResolver {
	return &kmakeRunDummyResolver{r}
}
func (r *Resolver) KmakeRunJob() KmakeRunJobResolver {
	return &kmakeRunJobResolver{r}
}
func (r *Resolver) Namespace() NamespaceResolver {
	return &namespaceResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type kmakeResolver struct{ *Resolver }

type kmakeRunResolver struct{ *Resolver }

type kmakeRunDummyResolver struct{ *Resolver }

type kmakeRunJobResolver struct{ *Resolver }

type namespaceResolver struct{ *Resolver }

type queryResolver struct{ *Resolver }
