// Copyright 2016-2019, KhulnaSoft, Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ciutil

import (
	"os"
)

// genericCICI represents a generic CI/CD system.
// It provides a way for users to workaround the fact that the CLI
// may not know about their CI system. This is sort of an
// escape hatch to give the CLI a hint about the CI environment.
type genericCICI struct {
	baseCI
}

// DetectVars detects the env vars for a Generic CI system.
func (g genericCICI) DetectVars() Vars {
	v := Vars{}
	v.Name = SystemName(os.Getenv("KHULNASOFT_CI_SYSTEM"))
	v.BranchName = os.Getenv("KHULNASOFT_CI_BRANCH_NAME")
	v.BuildID = os.Getenv("KHULNASOFT_CI_BUILD_ID")
	v.BuildNumber = os.Getenv("KHULNASOFT_CI_BUILD_NUMBER")
	v.BuildType = os.Getenv("KHULNASOFT_CI_BUILD_TYPE")
	v.BuildURL = os.Getenv("KHULNASOFT_CI_BUILD_URL")
	v.CommitMessage = os.Getenv("KHULNASOFT_COMMIT_MESSAGE")
	v.PRNumber = os.Getenv("KHULNASOFT_PR_NUMBER")
	v.SHA = os.Getenv("KHULNASOFT_CI_PULL_REQUEST_SHA")

	return v
}
