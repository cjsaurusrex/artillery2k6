package models

import (
	"errors"
)

type RequestAction struct {
	Name    string
	Method  string
	URL     string
	Headers map[string]any
	Expect  map[string]any
}

func (r *RequestAction) Build(key string, data any) error {
	if values, ok := data.(map[any]any); ok {
		r.Method = key
		r.URL = values["url"].(string)

		if name, ok := values["name"].(string); ok {
			r.Name = name
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

		return nil
	} else {
		return errors.New(`invalid request format`)
	}
}
