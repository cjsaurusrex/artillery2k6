package models

import (
	"testing"
)

func TestLogAction_Build_SetsValue(t *testing.T) {
	action := &LogAction{}
	err := action.Build("", "myValue")
	if err != nil {
		t.Errorf("Expected nil errors, got %s", err)
	}

	if action.Value != "myValue" {
		t.Errorf("Expected Value myValue, got %s", action.Value)
	}
}
