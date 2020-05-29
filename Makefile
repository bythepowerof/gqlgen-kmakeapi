# Image URL to use all building/pushing image targets
IMG ?= bythepowerof/gqlgen-kmakeapi:latest
NAMESPACE ?= all

# PKGS := github.com/bythepowerof/gqlgen-kmakeapi,github.com/bythepowerof/gqlgen-kmakeapi/controller,github.com/bythepowerof/gqlgen-kmakeapi/k8s,github.com/bythepowerof/gqlgen-kmakeapi/view

# Build manager binary
api: bin fmt vet
	go build -o bin/api 

server: build
	go run main.go -namespace ${NAMESPACE}

view/resolver.go: view/schema.graphql view/gqlgen.yml
	-mv view/resolver.go view/resolver.go.sav
	cd view; go run github.com/99designs/gqlgen --verbose


fix: view/resolver.go
	cd view; ./fix.pl

diff: fix
	-diff -q --from-file view/resolver.go viw/resolver.go.sav 

build: diff bin tidy fmt vet

tidy:
	go mod tidy

fmt:
	go fmt ./...

vet:
	go vet ./...

bin:
	mkdir $@

test:
	go test -count 1  -coverprofile cover.out ./...
	# go test -count 1 -coverpkg $(PKGS) -coverprofile cover.out ./...

cover: test
	go tool cover -html=cover.out

# Build the docker image
docker-build: test
	docker build . -t ${IMG}

# Push the docker image
docker-push:
	docker push ${IMG}	

# Deploy controller in the configured Kubernetes cluster in ~/.kube/config
deploy:
	cd config/api && kustomize edit set image api=${IMG}
	kustomize build config/default | kubectl apply -f -

clean:
	-rm -fr dist bin cover.out coverage.txt cp.out

.PHONY: server build test fix diff fmt vet tidy cover docker-push docker-build
