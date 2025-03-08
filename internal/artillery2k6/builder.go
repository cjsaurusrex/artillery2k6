package artillery2k6

import (
	"encoding/json"
	"fmt"
	"github.com/cjsaurusrex/artillery2k6/internal/artillery2k6/helpers"
	"github.com/cjsaurusrex/artillery2k6/internal/artillery2k6/models"
	"path/filepath"
	"slices"
)

func BuildScript(config *helpers.BuilderConfig, script models.ArtilleryScript) K6Script {
	k6Script := K6Script{}
	k6Script.InitLifecycle.Imports = append(k6Script.InitLifecycle.Imports, "import http from \"k6/http\"")

	for _, phase := range script.Config.Phases {
		results, _ := Convert(config, &phase)
		k6Script.InitLifecycle.Stages = append(k6Script.InitLifecycle.Stages, results...)
	}

	if config.EnvironmentsInUse {

		for _, env := range script.Config.Environments {
			if env.Variables != nil {
				updateEnvironmentVariableNames(&env.Variables)
			}
		}

		json, _ := json.MarshalIndent(script.Config.Environments, "", "  ")
		k6Script.InitLifecycle.Statements = append(k6Script.InitLifecycle.Statements, fmt.Sprintf("let environments = %s", string(json)))
	}

	k6Script.InitLifecycle.Statements = append(k6Script.InitLifecycle.Statements, buildVariables(config, script.Config.Variables)...)

	for _, scenario := range script.Scenarios {
		for _, action := range scenario.Flow.FlowActions {
			statements, imports := Convert(config, action)
			statements = append(statements, "") // Blank line to separate statement blocks
			k6Script.VULifecycle.Statements = append(k6Script.VULifecycle.Statements, statements...)

			for _, imp := range imports {
				if !slices.Contains(k6Script.InitLifecycle.Imports, imp) {
					k6Script.InitLifecycle.Imports = append(k6Script.InitLifecycle.Imports, imp)
				}
			}
		}
	}

	if script.Config.Processor != "" {
		k6Script.Processor = BuildProcessor(script.Config.Processor, filepath.Dir(script.Path), nil)
	}

	return k6Script
}

func BuildContext(script K6Script, config helpers.BuilderConfig) K6ScriptContext {
	return K6ScriptContext{
		Script: script,
		Config: config,
	}
}

func updateEnvironmentVariableNames(vars *map[string]any) {
	newVars := map[string]any{}
	for key, value := range *vars {
		newVars[helpers.BuildVariableName(key)] = value
	}
	*vars = newVars
}

func buildVariables(config *helpers.BuilderConfig, variables map[string]any) []string {
	vars := []string{}
	for key, value := range variables {
		v := helpers.BuildVariableDefinitionPrefix(config, key)

		if str, ok := value.(string); ok {
			v = fmt.Sprintf(`%s = "%s"`, v, str)
		} else {
			v = fmt.Sprintf("%s = %v", v, value)
		}

		vars = append(vars, v)
	}

	return vars
}
