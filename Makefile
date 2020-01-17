
server: resolver.go server/%.go
	go run server/server.go

%.go server/%.go:
	go fmt ./...
	go vet ./...
	go build ./...

test:
	go test ./...

resolver.go: schema.graphql
	mv resolver.go resolver.go.sav
	go run github.com/99designs/gqlgen --verbose
	diff -q --from-file resolver.go resolver.go.sav 
	
.PHONY: server build test generate
