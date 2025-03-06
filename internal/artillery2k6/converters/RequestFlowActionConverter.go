package converters

import (
	"encoding/json"
	"fmt"
	"github.com/cjsaurusrex/artillery2k6/internal/artillery2k6/models"
	"strconv"
	"strings"
	"unicode"
)

type RequestFlowActionConverter struct {
	*models.RequestAction
}

func (r *RequestFlowActionConverter) Convert() ([]string, []string) {
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

	json, _ := json.Marshal(params)
	statements = append(statements, fmt.Sprintf("let %s = http.%s(\"%s\", %s)", convertReqName(r.Name), r.Method, r.URL, string(json)))

	convertExpect(r.RequestAction, &statements, &imports)
	convertCapture(r.RequestAction, &statements, &imports)

	if r.AfterRequest != nil {
		for _, f := range r.AfterRequest {
			statements = append(statements, fmt.Sprintf("%s()", f))
		}
	}

	return statements, imports
}

func convertExpect(r *models.RequestAction, statements *[]string, imports *[]string) {
	if r.Expect != nil && len(r.Expect) > 0 {
		for key, value := range r.Expect {
			switch key {
			case "statusCode":
				statusCode := strconv.Itoa(value.(int))
				*imports = append(*imports, "import { check } from 'k6'")
				*statements = append(*statements, fmt.Sprintf("check(%s, { 'status is %s': (r) => r.status === %s })", convertReqName(r.Name), statusCode, statusCode))
			}
		}
	}
}

func convertCapture(r *models.RequestAction, statements *[]string, imports *[]string) {
	if r.Captures != nil && len(r.Captures) > 0 {
		for _, capture := range r.Captures {
			switch capture.Type {
			case "json":
				*statements = append(*statements, fmt.Sprintf(`let %s = %s.json("%s")`, capture.As, convertReqName(r.Name), capture.Value))
			}
		}
	}
}

func convertReqName(name string) string {
	n := []rune(strings.ReplaceAll(name, " ", ""))
	n[0] = unicode.ToLower(n[0])

	return string(n)
}
