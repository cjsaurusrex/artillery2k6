package models

import (
	"errors"
	"fmt"
	"slices"
	"strings"
)

type Capture struct {
	Type  string
	Value string
	As    string
}

type RequestAction struct {
	Count    int
	Name     string
	Method   string
	URL      string
	Headers  map[string]any
	Expect   map[string]any
	Captures []Capture
}

func NewRequestAction(requestCount int) *RequestAction {
	return &RequestAction{
		Count:   requestCount,
		Headers: make(map[string]any),
		Expect:  make(map[string]any),
	}
}

func (r *RequestAction) Build(key string, data any) error {
	if values, ok := data.(map[any]any); ok {
		r.Method = key
		r.URL = values["url"].(string)

		if name, ok := values["name"].(string); ok {
			r.Name = name
		} else {
			r.Name = fmt.Sprintf("Req %d", r.Count)
		}

		if headers, ok := values["headers"].(map[any]any); ok {
			r.Headers = make(map[string]any)
			for key, value := range headers {
				r.Headers[key.(string)] = value
			}
		}

		if expectList, ok := values["expect"].([]any); ok {
			r.Expect = make(map[string]any)
			for _, expect := range expectList {
				for key, value := range expect.(map[any]any) {
					r.Expect[key.(string)] = value
				}
			}
		}

		if values["capture"] != nil {
			caps := []Capture{}
			for _, capture := range values["capture"].([]any) {
				caps = append(caps, parseCapture(capture.(map[any]any)))
			}
			r.Captures = caps
		}

		return nil
	} else {
		return errors.New(`invalid request format`)
	}
}

func parseCapture(capture map[any]any) Capture {
	validTypes := []string{"json"}
	c := Capture{}
	for k, v := range capture {
		key := fmt.Sprintf("%v", k)
		if strings.ToLower(key) == "as" {
			c.As = fmt.Sprintf("%v", v)
			continue
		}
		if slices.Contains(validTypes, key) {
			c.Type = key
			c.Value = fmt.Sprintf("%v", v)
		}
	}
	return c
}
