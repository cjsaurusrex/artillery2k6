package converters

import (
	"github.com/cjsaurusrex/artillery2k6/internal/artillery2k6/models"
	"testing"
)

func TestLogActionConverter_Convert_StatementIsCorrectFormat(t *testing.T) {
	lac := &LogActionConverter{&models.LogAction{Value: "my test log"}}
	statement, _ := lac.Convert(nil)
	if statement[0] != "console.log(\"my test log\")" {
		t.Errorf("Expected 'console.log(\"my test log\")', got %s", statement[0])
	}
}

func TestLogActionConverter_Convert_NoImportsRequired(t *testing.T) {
	lac := &LogActionConverter{&models.LogAction{Value: "my test log"}}
	_, imports := lac.Convert(nil)
	if len(imports) != 0 {
		t.Errorf("Expected no imports, got %d", len(imports))
	}
}
