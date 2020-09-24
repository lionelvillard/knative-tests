package eventing_github

import (
	"testing"

	"github.com/knative-tests/pkg/framework"
)

func TestGitHubSource(t *testing.T) {
	framework.
		NewTest(t).
		Run("GitHubSource", func(tc framework.TestContext) {
			tc.CreateFromURIOrFail("testdata", false)
		})
}
