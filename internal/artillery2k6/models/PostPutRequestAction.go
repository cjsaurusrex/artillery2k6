package models

type PostPutRequestAction struct {
	*RequestAction
	Json map[string]any
}

func (r *PostPutRequestAction) Build(key string, data any) error {
	if err := r.RequestAction.Build(key, data); err != nil {
		return err
	}

	if values, ok := data.(map[any]any); ok {
		if values["json"] != nil {
			r.Headers["Content-Type"] = "application/json"
			x := values["json"].(map[any]any)
			r.Json = make(map[string]any)
			for k, v := range x {
				key := k.(string)
				r.Json[key] = v
			}
		}
	}
	return nil
}
