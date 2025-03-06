package converters

import (
	"github.com/cjsaurusrex/arillery2k6/internal/artillery2k6/models"
	"testing"
)

func createTestActionAndConverter(duration float64) *ThinkActionConverter {
	thinkAction := &models.ThinkAction{Duration: duration}
	return &ThinkActionConverter{ThinkAction: thinkAction}
}

func TestThinkActionConverter_Convert_StatementIsCorrectFormat(t *testing.T) {
	type test struct {
		name     string
		duration float64
		expected string
	}

	tests := []test{
		{"5.00", 5.00, "sleep(5.00)"},
		{"5.119", 5.119, "sleep(5.12)"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			thinkActionConverter := createTestActionAndConverter(tc.duration)
			statement, _ := thinkActionConverter.Convert()
			if statement[0] != tc.expected {
				t.Errorf("Expected '%s', got %s", tc.expected, statement[0])
			}
		})
	}
}

func TestThinkActionConverter_Convert_ImportsAreIncluded(t *testing.T) {
	thinkActionConverter := createTestActionAndConverter(5.00)

	_, imports := thinkActionConverter.Convert()
	if imports[0] != "import { sleep } from 'k6'" {
		t.Errorf("Expected 'import { sleep } from 'k6'', got %s", imports[0])
	}
}
