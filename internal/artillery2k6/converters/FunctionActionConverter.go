package converters

import (
	"github.com/cjsaurusrex/artillery2k6/internal/artillery2k6/helpers"
	"github.com/cjsaurusrex/artillery2k6/internal/artillery2k6/models"
)

type FunctionActionConverter struct {
	*models.FunctionAction
}

func (f *FunctionActionConverter) Convert(_ *helpers.BuilderConfig) ([]string, []string) {
	return []string{f.Name + "()"}, []string{}
}
