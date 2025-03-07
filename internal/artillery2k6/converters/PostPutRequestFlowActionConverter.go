package converters

import (
	"encoding/json"
	"fmt"
	"github.com/cjsaurusrex/artillery2k6/internal/artillery2k6/helpers"
	"github.com/cjsaurusrex/artillery2k6/internal/artillery2k6/models"
)

type PostPutRequestFlowActionConverter struct {
	*models.PostPutRequestAction
}

func (r *PostPutRequestFlowActionConverter) Convert(config *helpers.BuilderConfig) ([]string, []string) {
	reqName := convertReqName(r.Name)
	var params = make(map[string]any)
	imports, statements := []string{}, []string{}

	if r.BeforeRequest != nil {
		for _, f := range r.BeforeRequest {
			statements = append(statements, fmt.Sprintf("%s()", f))
		}
	}

	if r.Headers != nil && len(r.Headers) > 0 {
		params["headers"] = r.Headers
	}

	paramsJson, _ := json.Marshal(params)

	if r.Json != nil {
		rdn := fmt.Sprintf("%sData", reqName)
		jsonString, _ := json.Marshal(r.Json)
		statements = append(statements, helpers.InterpolateArtilleryVariables(config, fmt.Sprintf("let %s = %s", rdn, string(jsonString))))
		statements = append(statements, helpers.InterpolateArtilleryVariables(config, fmt.Sprintf("let %s = http.%s(\"%s\", JSON.stringify(%s), %s)", reqName, r.Method, r.URL, rdn, string(paramsJson))))
	} else {
		statements = append(statements, helpers.InterpolateArtilleryVariables(config, fmt.Sprintf("let %s = http.%s(\"%s\", %s)", reqName, r.Method, r.URL, string(paramsJson))))
	}

	// Convert Expect & Capture
	convertExpect(r.RequestAction, &statements, &imports)
	convertCapture(r.RequestAction, &statements, &imports)

	if r.AfterRequest != nil {
		for _, f := range r.AfterRequest {
			statements = append(statements, fmt.Sprintf("%s()", f))
		}
	}

	return statements, imports
}
