package converters

import (
	"encoding/json"
	"github.com/cjsaurusrex/artillery2k6/internal/artillery2k6/helpers"
	"github.com/cjsaurusrex/artillery2k6/internal/artillery2k6/models"
	"strconv"
)

type PhaseConverter struct {
	Base     *models.Phase `json:"-"`
	Duration string        `json:"duration"`
	Target   string        `json:"target"`
}

func (p *PhaseConverter) Convert(_ *helpers.BuilderConfig) ([]string, []string) {
	p.Duration = "60"
	if p.Base.Duration != nil {
		p.Duration = strconv.Itoa(*p.Base.Duration)
	}
	p.Target = strconv.Itoa(*p.Base.ArrivalRate)

	res, _ := json.Marshal(p)
	return []string{string(res)}, []string{}

	// Todo: Implement RampTo and ArrivalCount
}
