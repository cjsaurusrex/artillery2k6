package helpers

type BuilderConfig struct {
	// Determines if 'Environments' are being used in the Artillery script
	EnvironmentsInUse bool

	// Name of the function that retrieves a variable from the environment
	GetVariableFromEnvironmentFuncName string

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
		RootVariableFormat:                 Local,
		GetVariableFromEnvironmentFuncName: "getVariableFromEnvironmentOrDefault",
		LoadCsvFunctionName:                "loadCsvFile",
		CsvOpenAlias:                       "csv_open",
		TargetVariableName:                 "target",
	}
}
