package artillery2k6

import (
	"github.com/cjsaurusrex/arillery2k6/internal/artillery2k6/models"
	"path/filepath"
	"slices"
)

type K6Script struct {
	Imports    []string
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
		//processorPath := filepath.Join(filepath.Dir(script.Path), script.Config.Processor)
		k6Script.Processor = BuildProcessor(script.Config.Processor, filepath.Dir(script.Path), nil)
	}

	return k6Script
}
