package factory_test

import (
	"dataStructuresAndAlgorithms/factory"
	"testing"
)

func TestNewConfigParse(t *testing.T) {
	_, err := factory.NewConfigParse("json01")
	if err != nil {
		t.Logf("this is wrong param")
	} else {
		t.Errorf("logic error")
	}

	_, err = factory.NewConfigParse("json")
	if err != nil {
		t.Errorf("failed to parse: %s", err)
	}
	t.Log("create json parse succeeded")
	_, err = factory.NewConfigParse("json")
	if err != nil {
		t.Errorf("failed to parse: %s", err)
	}
	t.Log("create json parse succeeded")
}
