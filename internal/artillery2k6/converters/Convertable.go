package converters

type Convertable interface {
	Convert() (statements []string, imports []string)
}
