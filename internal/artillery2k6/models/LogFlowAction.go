package models

type LogAction struct {
	Value string
}

func (l *LogAction) Build(_ string, data any) error {
	l.Value = data.(string)
	return nil
}
