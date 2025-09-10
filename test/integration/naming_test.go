/*
Copyright 2024 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package integration

import (
	"fmt"
	"testing"

	"k8s.io/kube-openapi/pkg/util"
	"k8s.io/kube-openapi/pkg/validation/spec"
	generated "k8s.io/kube-openapi/test/integration/pkg/generated/namedmodels"
	"k8s.io/kube-openapi/test/integration/testdata/namedmodels"
)

func TestCanonicalTypeNames(t *testing.T) {
	defs := generated.GetOpenAPIDefinitions(func(path string) spec.Ref {
		return spec.MustCreateRef(path)
	})

	// Define the types we want to check.
	typeTestCases := []any{namedmodels.Struct{}, namedmodels.ContainedStruct{}, namedmodels.AtomicStruct{}}

	for _, v := range typeTestCases {
		t.Run(fmt.Sprintf("%T", v), func(t *testing.T) {
			// Get the canonical name using the generator's logic.
			canonicalName := util.GetCanonicalTypeName(v)

			// Check if the canonical name exists as a key in the generated map.
			if _, ok := defs[canonicalName]; !ok {
				t.Errorf("canonical type name %q for type %q not found in GetOpenAPIDefinitions", canonicalName, v)
			}
		})
	}

	// Additionally, verify that the number of generated definitions matches our expectation.
	if len(defs) != len(typeTestCases) {
		t.Errorf("expected %d definitions, but got %d", len(typeTestCases), len(defs))
	}
}
