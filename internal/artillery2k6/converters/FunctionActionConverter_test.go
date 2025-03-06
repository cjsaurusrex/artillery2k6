package converters

import (
	"github.com/cjsaurusrex/artillery2k6/internal/artillery2k6/models"
	"testing"
)

func TestFunctionActionConverter_Convert_StatementIsCorrectFormat(t *testing.T) {
	fac := &FunctionActionConverter{&models.FunctionAction{Name: "myFunction"}}
	statement, _ := fac.Convert()
	if statement[0] != "myFunction()" {
		t.Errorf("Expected 'myFunction()', got %s", statement[0])
	}
}

func TestFunctionActionConverter_Convert_NoImportsRequired(t *testing.T) {
	fac := &FunctionActionConverter{&models.FunctionAction{Name: "myFunction"}}
	_, imports := fac.Convert()
	if len(imports) != 0 {
		t.Errorf("Expected no imports, got %d", len(imports))
	}
}
