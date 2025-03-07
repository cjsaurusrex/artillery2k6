package artillery2k6

import (
	"fmt"
	"github.com/cjsaurusrex/artillery2k6/internal/artillery2k6/helpers"
	"github.com/cjsaurusrex/artillery2k6/internal/artillery2k6/models"
	"path/filepath"
	"slices"
)

type K6Script struct {
	Imports    []string
	Variables  map[string]any
	Statements []string
	Stages     []string
	Processor  []Processor
}

func Build(script models.ArtilleryScript) K6Script {
	k6Script := K6Script{}
	k6Script.Imports = append(k6Script.Imports, "import http from \"k6/http\"")

	for _, phase := range script.Config.Phases {
		results, _ := Convert(&phase)
		k6Script.Stages = append(k6Script.Stages, results...)
	}

	k6Script.Variables = buildVariables(script.Config.Variables)

	for _, scenario := range script.Scenarios {
		for _, action := range scenario.Flow.FlowActions {
			statements, imports := Convert(action)
			statements = append(statements, "") // Blank line to separate statement blocks
			k6Script.Statements = append(k6Script.Statements, statements...)

			for _, imp := range imports {
				if !slices.Contains(k6Script.Imports, imp) {
					k6Script.Imports = append(k6Script.Imports, imp)
				}
			}
		}
	}

	if script.Config.Processor != "" {
		k6Script.Processor = BuildProcessor(script.Config.Processor, filepath.Dir(script.Path), nil)
	}

	return k6Script
}

func buildVariables(variables map[string]any) map[string]any {
	vars := map[string]any{}
	for key, value := range variables {
		if str, ok := value.(string); ok {
			vars[helpers.BuildVariableName(key)] = fmt.Sprintf(`"%s"`, str)
		} else {
			vars[helpers.BuildVariableName(key)] = value
		}
	}

	return vars
}
