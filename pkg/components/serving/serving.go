/*
Copyright 2019 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package serving

import (
	"fmt"

	"github.com/knative-tests/pkg/config"
	"github.com/knative-tests/pkg/framework"
)

const (
	artifactURITemplate = "https://github.com/knative/serving/releases/download/v%s/%s"
)

var (
	Component = &servingComponent{}
)

type servingComponent struct {
}

var _ framework.Component = (*servingComponent)(nil)

// Scope returns the component scope
func (s *servingComponent) Scope() framework.ComponentScope {
	return framework.ComponentScopeCluster
}

func (s *servingComponent) Required(rc framework.ResourceContext, rcfg config.Config) {
	cfg := config.GetConfig(rcfg, "components/serving").(ServingConfig)

	rc.Apply(fmt.Sprintf(artifactURITemplate, cfg.Version, "serving-crds.yaml"))
	rc.Apply(fmt.Sprintf(artifactURITemplate, cfg.Version, "serving-core.yaml"))
}
