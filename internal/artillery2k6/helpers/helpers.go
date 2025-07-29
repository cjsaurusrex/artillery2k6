package helpers

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

func BuildVariableName(input string) string {
	n := []rune(strings.ReplaceAll(input, " ", ""))
	n[0] = unicode.ToLower(n[0])

	return string(n)
}

func BuildVariableReference(config *BuilderConfig, input string) string {
	if config != nil && config.EnvironmentsInUse {
		return fmt.Sprintf(`%s("%s")`, config.VariableFromEnvironmentFuncName(), BuildVariableName(input))
	}
	if config != nil && config.RootVariableFormat == GlobalThis {
		return fmt.Sprintf(`globalThis["%s"]`, BuildVariableName(input))
	}

	return BuildVariableName(input)
}

func BuildVariableDefinitionPrefix(config *BuilderConfig, input string) string {
	if config != nil && config.RootVariableFormat == GlobalThis {
		return fmt.Sprintf(`globalThis["%s"]`, BuildVariableName(input))
	}

	return fmt.Sprintf("let %s", BuildVariableName(input))
}

func InterpolateArtilleryVariables(config *BuilderConfig, input string) string {
	if !strings.Contains(input, "{{") {
		return input
	}

	result := input

	regex := regexp.MustCompile(`{{\s*(\w+)\s*}}`)
	matches := regex.FindAllStringSubmatch(input, -1)

	if len(matches) == 0 {
		return input
	}

	for _, match := range matches {
		result = strings.ReplaceAll(result, match[0], fmt.Sprintf("${ %s }", BuildVariableReference(config, match[1])))
	}

	dq := regexp.MustCompile(`\"([^"]*\$\{[^}]+\}[^"]*)\"`) // Find double quotes with variables
	dqMatches := dq.FindAllStringSubmatch(result, -1)

	for _, match := range dqMatches {
		result = strings.ReplaceAll(result, match[0], fmt.Sprintf("`%s`", match[1]))
	}

	return result
}

func AnyMatch(slice []any, condition func(any) bool) bool {
	for _, v := range slice {
		if condition(v) {
			return true
		}
	}
	return false
}
