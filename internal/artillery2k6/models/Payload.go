package models

import (
	"encoding/json"
	"fmt"
	"strings"
	"unicode"
)

type Payload struct {
	Payloads []PayloadConfig
}

type PayloadConfig struct {
	Name       string   `yaml:"name" json:"name"`
	Path       string   `yaml:"path" json:"path"`
	Order      string   `yaml:"order" json:"-"`     // todo: Implement this
	Delimiter  string   `yaml:"delimiter" json:"-"` // todo: Implement this
	Fields     []string `yaml:"fields" json:"-"`
	SkipHeader bool     `yaml:"skipHeader" json:"-"` // todo: Implement this
	LoadAll    bool     `yaml:"loadAll" json:"-"`    // todo: Implement this
}

func (p *PayloadConfig) GenerateName() {
	if p.Name != "" {
		return
	}

	name := ""
	for _, field := range p.Fields {
		n := []rune(strings.ReplaceAll(field, " ", ""))
		n[0] = unicode.ToUpper(n[0])
		name += string(n)
	}

	n := []rune(name)
	n[0] = unicode.ToLower(n[0])
	p.Name = string(n)
}

// UnmarshalYAML Artillery allows `payload` to be either a single object or an array of objects
func (p *Payload) UnmarshalYAML(unmarshal func(interface{}) error) error {
	// Check for individual payload
	var singlePayload PayloadConfig
	if err := unmarshal(&singlePayload); err == nil {
		singlePayload.GenerateName()
		p.Payloads = append(p.Payloads, singlePayload)
		return nil
	}

	var multiplePayloads []PayloadConfig
	if err := unmarshal(&multiplePayloads); err == nil {
		for i := 0; i < len(multiplePayloads); i++ {
			multiplePayloads[i].GenerateName()
		}
		p.Payloads = multiplePayloads
		return nil
	}

	return fmt.Errorf("unable to unmarshal payload")
}

// MarshalJSON Within our k6 script JSON, we simply want the names->paths of the payloads
func (p Payload) MarshalJSON() ([]byte, error) {
	payloads := make(map[string]string)
	for _, config := range p.Payloads {
		payloads[config.Name] = config.Path
	}

	return json.Marshal(payloads)
}
