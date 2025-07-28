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

		j, _ := json.MarshalIndent(script.Config.Environments, "", "  ")
		k6Script.InitLifecycle.Statements = append(k6Script.InitLifecycle.Statements, fmt.Sprintf("let environments = %s", string(j)))
	}

	k6Script.InitLifecycle.Statements = append(k6Script.InitLifecycle.Statements, buildVariables(config, script.Config.Variables)...)

	if config.PayloadsInUse {
		statements, imports := buildPayloadDeclarations(config, script)
		k6Script.InitLifecycle.Statements = append(k6Script.InitLifecycle.Statements, statements...)
		k6Script.InitLifecycle.Imports = append(k6Script.InitLifecycle.Imports, imports...)

		k6Script.VULifecycle.Statements = append(k6Script.VULifecycle.Statements, buildPayloadVariableStatements(config, script)...)
	}

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

	if script.Config.Target != "" {
		targetStatement := helpers.BuildVariableDefinitionPrefix(config, config.TargetVariableName) + fmt.Sprintf(" = \"%s\"", script.Config.Target)
		k6Script.InitLifecycle.Statements = append(k6Script.InitLifecycle.Statements, targetStatement)
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

func buildPayloadDeclarations(config *helpers.BuilderConfig, script models.ArtilleryScript) (statements []string, imports []string) {
	imports = []string{fmt.Sprintf("import { open as %s } from 'k6/experimental/fs'", config.CsvOpenAlias),
		"import csv from 'k6/experimental/csv'"}
	statements = []string{}

	rootPayloads, environmentPayloads := []string{}, []string{}
	for _, v := range script.Config.Payloads.Payloads {
		rootPayloads = append(rootPayloads, v.Name)
	}

	for _, e := range script.Config.Environments {
		for _, v := range e.Payload.Payloads {
			environmentPayloads = append(environmentPayloads, v.Name)
		}
	}

	completed := []string{}
	for _, payload := range append(rootPayloads, environmentPayloads...) {
		if slices.Contains(completed, payload) {
			continue
		}

		if slices.Contains(rootPayloads, payload) && slices.Contains(environmentPayloads, payload) {
			statements = append(statements, fmt.Sprintf(`let %sData = %s(environments[__ENV.ENVIRONMENT]?.payloads?.%s || "%s")`, payload, config.LoadCsvFunctionName, payload, payload))
		} else if slices.Contains(rootPayloads, payload) {
			statements = append(statements, fmt.Sprintf(`let %sData = %s("%s")`, payload, config.LoadCsvFunctionName, payload))
		} else if slices.Contains(environmentPayloads, payload) {
			statements = append(statements, fmt.Sprintf(`let %sData = %s(environments[__ENV.ENVIRONMENT]?.payloads?.%s)`, payload, config.LoadCsvFunctionName, payload))
		}
		completed = append(completed, payload)
	}

	return statements, imports
}

func buildPayloadVariableStatements(config *helpers.BuilderConfig, script models.ArtilleryScript) []string {
	statements := []string{}
	allPayloads := []models.PayloadConfig{}

	for _, env := range script.Config.Environments {
		allPayloads = append(allPayloads, env.Payload.Payloads...)
	}
	allPayloads = append(allPayloads, script.Config.Payloads.Payloads...)
	completed := []string{}

	for _, payload := range allPayloads {
		if slices.Contains(completed, payload.Name) {
			continue
		}

		for _, field := range payload.Fields {
			v := helpers.BuildVariableDefinitionPrefix(config, field)
			statements = append(statements, fmt.Sprintf("%s = %sData[Math.floor(Math.random()*%sData.length)].%s", v, payload.Name, payload.Name, field))
		}

		completed = append(completed, payload.Name)
	}

	return statements
}
