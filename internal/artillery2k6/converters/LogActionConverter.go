package converters

import "github.com/cjsaurusrex/artillery2k6/internal/artillery2k6/models"

type LogActionConverter struct {
	*models.LogAction
}

func (l *LogActionConverter) Convert() (statements []string, imports []string) {
	logLine := "console.log(\"" + l.Value + "\")"
	return []string{logLine}, []string{}
}
