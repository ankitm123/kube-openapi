package namedmodels

type Struct struct {
	Field      ContainedStruct
	OtherField int
}

type ContainedStruct struct{}

type AtomicStruct struct {
	Field int
}
