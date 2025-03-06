package models

import (
	"testing"
)

func TestFunctionAction_Build_SetsName(t *testing.T) {
	action := &FunctionAction{}
	err := action.Build("", "myFunctionName")
	if err != nil {
		t.Errorf("Expected nil errors, got %s", err)
	}

	if action.Name != "myFunctionName" {
		t.Errorf("Expected Name myFunctionName, got %s", action.Name)
	}
}
