package models

type FlowAction interface {
	Build(string, any) error
}
