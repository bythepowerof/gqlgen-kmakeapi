#!/usr/bin/env bash

set -e
echo "" > coverage.txt

for d in $(go list ./... | grep -v vendor); do
    go test -race -coverprofile=profile.out \
        -covermode=atomic \
        -coverpkg github.com/bythepowerof/gqlgen-kmakeapi/controller,github.com/bythepowerof/gqlgen-kmakeapi/k8s,github.com/bythepowerof/gqlgen-kmakeapi/view \
        "$d"
    if [ -f profile.out ]; then
        cat profile.out >> coverage.txt
        rm profile.out
    fi
done