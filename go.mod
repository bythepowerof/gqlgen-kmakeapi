module github.com/bythepowerof/gqlgen-kmakeapi

go 1.13

require (
	github.com/99designs/gqlgen v0.11.3
	github.com/JeremyMarshall/gqlgen-jwt v0.1.1
	github.com/auth0/go-jwt-middleware v0.0.0-20200507191422-d30d7b9ece63
	github.com/bythepowerof/kmake-controller v0.1.12
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/golang/groupcache v0.0.0-20190702054246-869f871628b6 // indirect
	github.com/gorilla/handlers v1.4.2
	github.com/gorilla/websocket v1.4.2
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/mitchellh/mapstructure v1.2.2 // indirect
	github.com/namsral/flag v1.7.4-pre
	github.com/onsi/ginkgo v1.13.0
	github.com/onsi/gomega v1.10.1
	github.com/pkg/errors v0.9.1 // indirect
	github.com/prometheus/client_golang v1.1.0 // indirect
	github.com/rs/cors v1.7.0
	github.com/vektah/gqlparser/v2 v2.0.1
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d // indirect
	golang.org/x/time v0.0.0-20200416051211-89c76fbcd5d1 // indirect
	k8s.io/api v0.18.2
	k8s.io/apimachinery v0.0.0-20191020214737-6c8691705fc5
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
	k8s.io/utils v0.0.0-20200109141947-94aeca20bf09 // indirect
	sigs.k8s.io/controller-runtime v0.2.2
)

replace (
	k8s.io/api => k8s.io/api v0.0.0-20190409021203-6e4e0e4f393b
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190404173353-6a84e37a896d
	k8s.io/client-go => k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
)
