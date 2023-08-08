package adapter_test

import (
	"dataStructuresAndAlgorithms/adapter"
	"testing"
)

func TestAdapter(t *testing.T) {
	adaptee := &adapter.ConcreteAdaptee{}
	adapter := &adapter.Adapter{
		Adapt: adaptee,
	}
	adapter.Adapt.SpecificRequest()
}
