package models

import (
	"testing"
)

func TestFlowAction_Build_FloatIsUsedAsDuration(t *testing.T) {
	action := &ThinkAction{}
	err := action.Build("", 1.0)
	if err != nil {
		t.Errorf("Expected nil errors, got %s", err)
	}

	if action.Duration != 1.0 {
		t.Errorf("Expected Duration 1.0, got %f", action.Duration)
	}
}

func TestFlowAction_Build_IntIsUsedAsDuration(t *testing.T) {
	action := &ThinkAction{}
	err := action.Build("", 1)
	if err != nil {
		t.Errorf("Expected nil errors, got %s", err)
	}

	if action.Duration != 1.0 {
		t.Errorf("Expected Duration 1.0, got %f", action.Duration)
	}
}

func TestFlowAction_Build_StringFloatIsUsedAsDuration(t *testing.T) {
	action := &ThinkAction{}
	err := action.Build("", "1.0")
	if err != nil {
		t.Errorf("Expected nil errors, got %s", err)
	}

	if action.Duration != 1.0 {
		t.Errorf("Expected Duration 1.0, got %f", action.Duration)
	}

}

func TestFlowAction_Build_StringIsUsedAsDuration(t *testing.T) {
	type test struct {
		name     string
		duration string
	}

	tests := []test{
		{`"1"`, "1"},
		{`"1.0"`, "1.0"},
		{`"1s"`, "1s"},
		{`"1m""`, "1m"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			action := &ThinkAction{}
			err := action.Build("", "1")
			if err != nil {
				t.Errorf("Expected nil errors, got %s", err)
			}

			if action.Duration != 1.0 {
				t.Errorf("Expected Duration 1.0, got %f", action.Duration)
			}
		})
	}
}
