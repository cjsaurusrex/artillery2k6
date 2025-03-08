package helpers

type BuilderConfig struct {
	EnvironmentsInUse                  bool
	GetVariableFromEnvironmentFuncName string
	RootVariableFormat                 VariableFormat
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
	}
}
