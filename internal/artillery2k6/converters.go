package artillery2k6

import (
	"github.com/cjsaurusrex/arillery2k6/internal/artillery2k6/converters"
	"github.com/cjsaurusrex/arillery2k6/internal/artillery2k6/models"
)

type Convertable interface {
	Convert() (results []string, imports []string)
}

func Convert(action any) (results []string, imports []string) {
	var convertable Convertable
	switch a := action.(type) {
	case *models.LogAction:
		convertable = &converters.LogActionConverter{LogAction: a}
	case *models.ThinkAction:
		convertable = &converters.ThinkActionConverter{ThinkAction: a}
	case *models.PostPutRequestAction:
		convertable = &converters.PostPutRequestFlowActionConverter{PostPutRequestAction: a}
	case *models.RequestAction:
		convertable = &converters.RequestFlowActionConverter{RequestAction: a}
	case *models.Phase:
		convertable = &converters.PhaseConverter{Base: a}
	}
	return convertable.Convert()
}
