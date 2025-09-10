package namedmodels

// +k8s:openapi-model-package=io.k8s.kube-openapi.test.integration.testdata.namedmodels
type Struct struct { // The generator should see the +k8s:openapi-model-package and assume that a OpenAPIModelName is (or will be) generated.
	Field      ContainedStruct
	OtherField int
}

type ContainedStruct struct{} // The generator should use the go package name since there is no OpenAPIModelName declared.

type AtomicStruct struct { // The generator should respect a manually declared OpenAPIModelName.
	Field int
}

// OpenAPIModelName returns the OpenAPI model name for this type.
func (in AtomicStruct) OpenAPIModelName() string {
	return "io.k8s.kube-openapi.test.integration.testdata.namedmodels.AtomicStruct"
}
