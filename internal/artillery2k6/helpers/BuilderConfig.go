package helpers

import (
	"fmt"
	"strings"
)

type BuilderConfig struct {
	// Determines if 'Environments' are being used in the Artillery script
	EnvironmentsInUse bool

	// Name of the function that retrieves a variable from the environment
	GetVariableFromEnvironmentFuncName string

	// Name of 'Environments' within the outputted k6 script. This should not be pluralized.
	EnvironmentName string

	// Format of root scoped variables
	// GlobalThis: Variables are defined under `globalThis` keys
	// Local: Variables are defined as local variables
	RootVariableFormat VariableFormat

	// Determines if 'Payloads' are being used in the Artillery script
	PayloadsInUse bool

	// Name of the function that loads a CSV file
	LoadCsvFunctionName string

	// Alias of the 'csv.open' function
	CsvOpenAlias string

	// Name of the Artillery 'Target' in the k6 script
	TargetVariableName string
}

type VariableFormat int

const (
	GlobalThis VariableFormat = iota
	Local
)

func NewBuilderConfig() *BuilderConfig {
	return &BuilderConfig{
		RootVariableFormat:  Local,
		EnvironmentName:     "environment",
		LoadCsvFunctionName: "loadCsvFile",
		CsvOpenAlias:        "csv_open",
		TargetVariableName:  "base_url",
	}
}

func (bc BuilderConfig) VariableFromEnvironmentFuncName() string {
	if bc.GetVariableFromEnvironmentFuncName != "" {
		return bc.GetVariableFromEnvironmentFuncName
	}

	envName := strings.ToUpper(bc.EnvironmentName[:1]) + strings.ToLower(bc.EnvironmentName[1:])
	return fmt.Sprintf("getVariableFrom%sOrDefault", envName)
}

func (bc BuilderConfig) PluralEnvironmentsName() string {
	return fmt.Sprintf("%ss", bc.EnvironmentName)
}

func (bc *BuilderConfig) SetEnvironmentName(envName string) {
	if len(envName) > 0 {
		bc.EnvironmentName = strings.ToUpper(envName[:1]) + strings.ToLower(envName[1:])
	}
}
