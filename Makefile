
server: build
	go run server/server.go

view/resolver.go: view/schema.graphql view/gqlgen.yml
	-mv view/resolver.go view/resolver.go.sav
	cd view; go run github.com/99designs/gqlgen --verbose


fix: view/resolver.go
	cd view; ./fix.pl

diff: fix
	-diff -q --from-file view/resolver.go viw/resolver.go.sav 

build: diff bin
	go fmt ./...
	go vet ./...
	go run server/server.go

bin:
	mkdir $@

test:
	go test ./...

.PHONY: server build test fix diff
