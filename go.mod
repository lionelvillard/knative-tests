module github.com/knative-tests

go 1.14


require (
	github.com/cloudevents/sdk-go/v2 v2.2.0
	github.com/octago/sflags v0.2.0
	github.com/onsi/gomega v1.10.2
	go.opencensus.io v0.22.4
	k8s.io/api v0.19.2
	k8s.io/apimachinery v0.19.2
	k8s.io/client-go v11.0.1-0.20190805182717-6502b5e7b1b5+incompatible
	knative.dev/pkg v0.0.0-20200923200940-01609c886287
	knative.dev/test-infra v0.0.0-20200923231240-18d2c662bb83
)

replace (
	k8s.io/api => k8s.io/api v0.18.8
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.18.8
	k8s.io/apimachinery => k8s.io/apimachinery v0.18.8
	k8s.io/client-go => k8s.io/client-go v0.18.8
	k8s.io/code-generator => k8s.io/code-generator v0.18.8
)
