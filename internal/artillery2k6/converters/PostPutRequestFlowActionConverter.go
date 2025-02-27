package converters

import (
	"encoding/json"
	"github.com/cjsaurusrex/arillery2k6/internal/artillery2k6/models"
)

type PostPutRequestFlowActionConverter struct {
	*models.PostPutRequestAction
}

func (r *PostPutRequestFlowActionConverter) Convert() ([]string, []string) {
	var params = make(map[string]any)
	imports, statements := []string{}, []string{}
	if r.Headers != nil {
		params["headers"] = r.Headers
	}

	paramsJson, _ := json.Marshal(params)

	if r.Json != nil {
		jsonString, _ := json.Marshal(r.Json)
		statements = append(statements, "let requestData = "+string(jsonString))
	}
	statements = append(statements, "http."+r.Method+"(\""+r.URL+"\", JSON.stringify(requestData), "+string(paramsJson)+");")

	// Convert Expect
	expectStatements, expectImports := ConvertExpect("req", r.RequestAction)
	statements = append(statements, expectStatements...)
	imports = append(imports, expectImports...)

	return statements, imports
}
