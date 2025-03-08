package converters

import "github.com/cjsaurusrex/artillery2k6/internal/artillery2k6/helpers"

type Convertable interface {
	Convert(config *helpers.BuilderConfig) (statements []string, imports []string)
}
