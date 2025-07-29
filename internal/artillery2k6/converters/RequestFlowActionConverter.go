package converters

import (
	"encoding/json"
	"fmt"
	"github.com/cjsaurusrex/artillery2k6/internal/artillery2k6/helpers"
	"github.com/cjsaurusrex/artillery2k6/internal/artillery2k6/models"
	"strconv"
	"strings"
	"unicode"
)

type RequestFlowActionConverter struct {
	*models.RequestAction
}

func (r *RequestFlowActionConverter) formatURL(config *helpers.BuilderConfig, url string) string {
	if strings.HasPrefix(strings.ToLower(url), "http") {
		return fmt.Sprintf("`%s`", url)
	}
	targetReference := helpers.BuildVariableReference(config, "target")
	return fmt.Sprintf("`${%s}%s`", targetReference, strings.TrimPrefix(url, "/"))
}

func (r *RequestFlowActionConverter) Convert(config *helpers.BuilderConfig) ([]string, []string) {
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

	j, _ := json.Marshal(params)

	url := r.URL
	if strings.HasPrefix(strings.ToLower(url), "http") {
		url = fmt.Sprintf("\"%s\"", url)
	} else {
		targetReference := helpers.BuildVariableReference(config, config.TargetVariableName)
		url = fmt.Sprintf("%s + \"%s\"", targetReference, url)
	}

	var statement string
	if params != nil && len(params) > 0 {
		statement = fmt.Sprintf("let %s = http.%s(%s, %s)", convertReqName(r.Name), r.Method, url, string(j))
	} else {
		statement = fmt.Sprintf("let %s = http.%s(%s)", convertReqName(r.Name), r.Method, url)
	}

	statement = helpers.InterpolateArtilleryVariables(config, statement)

	statements = append(statements, statement)

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

func formatURL(config *helpers.BuilderConfig, url string) string {
	if strings.HasPrefix(strings.ToLower(url), "http") {
		return fmt.Sprintf("\"%s\"", url)
	} else {
		targetReference := helpers.BuildVariableReference(config, config.TargetVariableName)
		return fmt.Sprintf("%s + \"%s\"", targetReference, url)
	}
}
