package eventing_github

import (
	"testing"

	"github.com/knative-tests/pkg"
	"github.com/knative-tests/pkg/components/eventing/sources/github"
	"github.com/knative-tests/pkg/framework"
)

func TestMain(m *testing.M) {
	framework.
		NewSuite(m).
		Configure(&pkg.AllConfig{}).
		Require(github.Component).
		Run()
}
