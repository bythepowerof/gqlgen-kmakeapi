module github.com/bythepowerof/gqlgen-kmakeapi

go 1.13

require (
	github.com/99designs/gqlgen v0.11.3
	github.com/bythepowerof/kmake-controller v0.1.11-0.20200425122911-8a2cda2b0dd1
	github.com/gorilla/websocket v1.4.1
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/mitchellh/mapstructure v1.1.2 // indirect
	github.com/namsral/flag v1.7.4-pre
	github.com/onsi/ginkgo v1.12.0
	github.com/onsi/gomega v1.9.0
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rs/cors v1.6.0
	github.com/vektah/gqlparser/v2 v2.0.1
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d // indirect
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0 // indirect
	golang.org/x/xerrors v0.0.0-20191204190536-9bdfabe68543 // indirect
	gopkg.in/yaml.v2 v2.2.8 // indirect
	k8s.io/api v0.17.0
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
