package models

import (
	"errors"
	"fmt"
	"github.com/cjsaurusrex/artillery2k6/internal/artillery2k6/helpers"
	"slices"
	"strings"
)

type Capture struct {
	Type  string
	Value string
	As    string
}

type RequestAction struct {
	Count         int
	Name          string
	Method        string
	URL           string
	Headers       map[string]any
	Expect        map[string]any
	Captures      []Capture
	BeforeRequest []string
	AfterRequest  []string
}

func NewRequestAction(requestCount int) *RequestAction {
	return &RequestAction{
		Count:   requestCount,
		Headers: make(map[string]any),
		Expect:  make(map[string]any),
	}
}

var builderFlow = []func(r *RequestAction, data map[any]any){
	buildHeaders,
	buildExpectations,
	buildCaptures,
	buildProcessorActions,
}

func (r *RequestAction) Build(key string, data any) error {
	if values, ok := data.(map[any]any); ok {
		r.Method = key
		r.URL = values["url"].(string)

		// Set default name
		if name, ok := values["name"].(string); ok {
			r.Name = name
		} else {
			r.Name = fmt.Sprintf("Req %d", r.Count)
		}

		for _, f := range builderFlow {
			f(r, values)
		}

		return nil
	} else {
		return errors.New(`invalid request format`)
	}
}

func buildHeaders(r *RequestAction, data map[any]any) {
	if headers, ok := data["headers"].(map[any]any); ok {
		r.Headers = make(map[string]any)
		for key, value := range headers {
			r.Headers[key.(string)] = value
		}
	}
}

func buildExpectations(r *RequestAction, data map[any]any) {
	if expectList, ok := data["expect"].([]any); ok {
		r.Expect = make(map[string]any)
		for _, expect := range expectList {
			for key, value := range expect.(map[any]any) {
				r.Expect[key.(string)] = value
			}
		}
	}
}

func buildCaptures(r *RequestAction, data map[any]any) {
	if data["capture"] != nil {
		caps := []Capture{}
		for _, capture := range data["capture"].([]any) {
			caps = append(caps, parseCapture(capture.(map[any]any)))
		}
		r.Captures = caps
	}
}

func buildProcessorActions(r *RequestAction, data map[any]any) {
	if data["beforeRequest"] != nil {
		r.BeforeRequest = buildProcessorAction(r, data["beforeRequest"])
	}

	if data["afterRequest"] != nil {
		r.AfterRequest = buildProcessorAction(r, data["afterRequest"])
	}
}

func buildProcessorAction(r *RequestAction, action any) []string {
	if list, ok := action.([]string); ok {
		return list
	} else if single, ok := action.(string); ok {
		return []string{single}
	} else {
		return []string{}
	}
}

func parseCapture(capture map[any]any) Capture {
	validTypes := []string{"json"}
	c := Capture{}
	for k, v := range capture {
		key := fmt.Sprintf("%v", k)
		if strings.ToLower(key) == "as" {
			c.As = helpers.BuildVariableName(fmt.Sprintf("%v", v)) // fmt.Sprintf("%v", v)
			continue
		}
		if slices.Contains(validTypes, key) {
			c.Type = key
			c.Value = fmt.Sprintf("%v", v)
		}
	}
	return c
}
