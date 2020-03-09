# gqlgen-kmakeapi - WIP

A GQL interface over the top of [kmake-controller](https://github.com/bythepowerof/kmake-controller)

Allows for operations on runs and queries on CRDs in kmake

Sample client written in python [kmake_gql_client](https://github.com/bythepowerof/kmake_gql_client)

### Build status

| Platform    | CI Status | Coverage | Report Card | Documentation |
|------------|:-------|:------------|:------- |----------- |
linux       |[![Build Status](https://travis-ci.org/bythepowerof/gqlgen-kmakeapi.svg?branch=master)](https://travis-ci.org/bythepowerof/gqlgen-kmakeapi)|[![codecov](https://codecov.io/gh/bythepowerof/gqlgen-kmakeapi/branch/master/graph/badge.svg)](https://codecov.io/gh/bythepowerof/gqlgen-kmakeapi)|[![Go Report Card](https://goreportcard.com/badge/github.com/bythepowerof/gqlgen-kmakeapi)](https://goreportcard.com/report/github.com/bythepowerof/gqlgen-kmakeapi)|[![GoDoc](https://godoc.org/github.com/bythepowerof/gqlgen-kmakeapi?status.svg)](https://godoc.org/github.com/bythepowerof/gqlgen-kmakeapi)|


### TODO

* Add a bootstrap to load new kmake/kmake-run from a docker image
* Dockerise it so it can run in the cluster
