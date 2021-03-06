/*
 * Copyright 2020 The Knative Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package framework

// BaseConfig defines the configuration parameters
// controlling the framework behavior.
type BaseConfig struct {
	KubeConfig      string
	ServerURL       string
	Setup			bool
	ImageRepository string
	Requirements    Requirements
}

type Requirements struct {
	Must   bool `desc:"run test mark as Must. Default is true"`
	Should bool `desc:"run test mark as Should. Default is true"`
	May    bool `desc:"run test mark as May. Default is true"`
}
