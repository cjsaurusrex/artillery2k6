package artillery2k6

import (
	"github.com/cjsaurusrex/artillery2k6/internal/artillery2k6/converters"
	"github.com/cjsaurusrex/artillery2k6/internal/artillery2k6/helpers"
	"github.com/cjsaurusrex/artillery2k6/internal/artillery2k6/models"
)

func Convert(config *helpers.BuilderConfig, action any) (results []string, imports []string) {
	var convertable converters.Convertable
	switch a := action.(type) {
	case *models.LogAction:
		convertable = &converters.LogActionConverter{LogAction: a}
	case *models.ThinkAction:
		convertable = &converters.ThinkActionConverter{ThinkAction: a}
	case *models.FunctionAction:
		convertable = &converters.FunctionActionConverter{FunctionAction: a}
	case *models.PostPutRequestAction:
		convertable = &converters.PostPutRequestFlowActionConverter{PostPutRequestAction: a}
	case *models.RequestAction:
		convertable = &converters.RequestFlowActionConverter{RequestAction: a}
	case *models.Phase:
		convertable = &converters.PhaseConverter{Base: a}
	}
	return convertable.Convert(config)
}
