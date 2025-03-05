package models

type FunctionAction struct {
	Name string
}

func (f *FunctionAction) Build(_ string, data any) error {
	f.Name = data.(string)
	return nil
}
