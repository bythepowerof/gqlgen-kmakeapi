
server: build
	go run server/server.go

resolver.go: schema.graphql
	mv resolver.go resolver.go.sav
	go run github.com/99designs/gqlgen --verbose
	./fix.pl
	diff -q --from-file resolver.go resolver.go.sav 

build: resolver.go bin
	go fmt ./...
	go vet ./...
	go run server/server.go

bin:
	mkdir $@

test:
	go test ./...

.PHONY: server build test 
