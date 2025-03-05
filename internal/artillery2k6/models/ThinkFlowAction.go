package models

import (
	"errors"
	"strconv"
	"time"
)

type ThinkAction struct {
	Duration float64
}

func (t *ThinkAction) Build(_ string, data any) error {
	if i, ok := data.(int); ok {
		t.Duration = float64(i)
		return nil
	}

	if f, ok := data.(float64); ok {
		t.Duration = f
		return nil
	}

	if s, ok := data.(string); ok {
		t.Duration = handleStringValue(s)
		return nil
	}

	return errors.New("invalid data type")
}

func handleStringValue(data string) float64 {
	if f, err := strconv.ParseFloat(data, 64); err == nil {
		return f
	}

	if i, err := strconv.Atoi(data); err == nil {
		return float64(i)
	}

	if d, err := time.ParseDuration(data); err == nil {
		return d.Seconds()
	}

	return 0
}
