package converters

import "github.com/cjsaurusrex/artillery2k6/internal/artillery2k6/models"

type FunctionActionConverter struct {
	*models.FunctionAction
}

func (f *FunctionActionConverter) Convert() ([]string, []string) {
	return []string{f.Name + "()"}, []string{}
}
