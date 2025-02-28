package converters

import (
	"encoding/json"
	"fmt"
	"github.com/cjsaurusrex/arillery2k6/internal/artillery2k6/models"
)

type PostPutRequestFlowActionConverter struct {
	*models.PostPutRequestAction
}

func (r *PostPutRequestFlowActionConverter) Convert() ([]string, []string) {
	reqName := convertReqName(r.Name)
	var params = make(map[string]any)
	imports, statements := []string{}, []string{}
	if r.Headers != nil {
		params["headers"] = r.Headers
	}

	paramsJson, _ := json.Marshal(params)

	if r.Json != nil {
		rdn := fmt.Sprintf("%sData", reqName)
		jsonString, _ := json.Marshal(r.Json)
		statements = append(statements, fmt.Sprintf("let %s = %s;", rdn, string(jsonString)))
		statements = append(statements, fmt.Sprintf("let %s = http.%s(\"%s\", JSON.stringify(%s), %s);", reqName, r.Method, r.URL, rdn, string(paramsJson)))
	} else {
		statements = append(statements, fmt.Sprintf("let %s = http.%s(\"%s\", %s);", reqName, r.Method, r.URL, string(paramsJson)))
	}

	// Convert Expect
	expectStatements, expectImports := ConvertExpect("req", r.RequestAction)
	statements = append(statements, expectStatements...)
	imports = append(imports, expectImports...)

	// Convert Captures
	captureStatements, captureImports := ConvertCapture(r.RequestAction)
	statements = append(statements, captureStatements...)
	imports = append(imports, captureImports...)

	return statements, imports
}
