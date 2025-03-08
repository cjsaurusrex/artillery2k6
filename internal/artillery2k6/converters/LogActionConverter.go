package converters

import (
	"github.com/cjsaurusrex/artillery2k6/internal/artillery2k6/helpers"
	"github.com/cjsaurusrex/artillery2k6/internal/artillery2k6/models"
)

type LogActionConverter struct {
	*models.LogAction
}

func (l *LogActionConverter) Convert(config *helpers.BuilderConfig) (statements []string, imports []string) {
	logLine := helpers.InterpolateArtilleryVariables(config, "console.log(\""+l.Value+"\")")
	return []string{logLine}, []string{}
}
