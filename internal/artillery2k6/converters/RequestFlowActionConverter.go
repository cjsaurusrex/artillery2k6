package converters

import (
	"encoding/json"
	"github.com/cjsaurusrex/arillery2k6/internal/artillery2k6/models"
	"strconv"
)

type RequestFlowActionConverter struct {
	*models.RequestAction
}

func (r *RequestFlowActionConverter) Convert() ([]string, []string) {
	var params = make(map[string]any)
	imports, statements := []string{}, []string{}
	if r.Headers != nil {
		params["headers"] = r.Headers
	}

	json, _ := json.Marshal(params)
	statements = append(statements, "http."+r.Method+"(\""+r.URL+"\", "+string(json)+");")

	if r.Expect != nil {
		expectStatements, expectImports := ConvertExpect("req", r.RequestAction)
		statements = append(statements, expectStatements...)
		imports = append(imports, expectImports...)
	}
	
	return statements, imports
}

func ConvertExpect(requestName string, r *models.RequestAction) ([]string, []string) {
	imports, statements := []string{}, []string{}

	if r.Expect != nil && len(r.Expect) > 0 {
		for key, value := range r.Expect {
			switch key {
			case "statusCode":
				statusCode := strconv.Itoa(value.(int))
				imports = append(imports, "import { check } from 'k6'")
				statements = append(statements, "check("+requestName+", { 'status is "+statusCode+"': (r) => r.status === "+statusCode+" })")
			}
		}
	}

	return statements, imports
}
