package converters

import (
	"fmt"
	"github.com/cjsaurusrex/artillery2k6/internal/artillery2k6/helpers"
	"github.com/cjsaurusrex/artillery2k6/internal/artillery2k6/models"
)

type ThinkActionConverter struct {
	*models.ThinkAction
}

func (t *ThinkActionConverter) Convert(_ *helpers.BuilderConfig) (statements []string, imports []string) {
	sleep := fmt.Sprintf("sleep(%.2f)", t.Duration)
	return []string{sleep}, []string{"import { sleep } from 'k6'"}
}
