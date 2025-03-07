package helpers

import (
	"testing"
)

func TestBuildVariableName_lowercasesStart(t *testing.T) {
	type test struct {
		input    string
		expected string
	}

	tests := []test{
		{"HelloWorld", "helloWorld"},
		{"helloWorld", "helloWorld"},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			actual := BuildVariableName(tc.input)
			if actual != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, actual)
			}
		})
	}
}

func TestInterpolateArtilleryVariables(t *testing.T) {
	type test struct {
		input    string
		expected string
	}

	tests := []test{
		{"{{hello}}", "${ hello }"},
		{"{{ hello }}", "${ hello }"},
		{`"www.{{domain}}.com"`, "`www.${ domain }.com`"},
		{`http.get("www.{{domain}}.com", { headers: {"User-Agent": "{{ userAgent }}" }})`, "http.get(`www.${ domain }.com`, { headers: {\"User-Agent\": `${ userAgent }` }})"},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			actual := InterpolateArtilleryVariables(tc.input)
			if actual != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, actual)
			}
		})
	}
}
