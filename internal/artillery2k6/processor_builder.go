package artillery2k6

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Processor struct {
	Name    string
	Content string
}

func BuildProcessor(processor string, basePath string, processors []Processor) []Processor {
	processorPath := filepath.Join(basePath, processor)
	processorBytes, err := os.ReadFile(processorPath)
	processorName := filepath.Base(processorPath)
	if err != nil {
		return []Processor{}
	}

	processorContent := removeExports(string(processorBytes))
	contentAsLines := strings.Split(processorContent, "\n")
	newContent := []string{}
	additionalProcessors := []string{}

	requireRegex := regexp.MustCompile(`require\(['"](.*)['"]\)`)
	for _, line := range contentAsLines {
		matches := requireRegex.FindStringSubmatch(line)
		if len(matches) > 1 {
			additionalProcessors = append(additionalProcessors, matches[1])
			continue
		}

		newContent = append(newContent, updateVariables(line))
	}

	processorContent = strings.Join(newContent, "\n")

	if processors == nil {
		processors = []Processor{Processor{Name: processorName, Content: processorContent}}
	} else {
		processors = append(processors, Processor{Name: processorName, Content: processorContent})
	}

	if len(additionalProcessors) > 0 {
		for _, p := range additionalProcessors {
			if !containsProcessor(processors, p) {
				processors = BuildProcessor(p, basePath, processors)
			}
		}
	}

	return processors
}

func removeExports(content string) string {
	regex := regexp.MustCompile(`module\.exports ?= ?{[\s\S]*}`)
	return regex.ReplaceAllString(content, "")
}

func updateVariables(content string) string {
	if strings.Contains(content, "context.vars.") {
		regex := regexp.MustCompile(`context\.vars\.(\w+)`)
		matches := regex.FindStringSubmatch(content)
		if len(matches) > 1 {
			return regex.ReplaceAllString(content, "globalThis['"+matches[1]+"']")
		}
	}
	if strings.Contains(content, "context.vars[") {
		regex := regexp.MustCompile(`context\.vars\[['"](\w+)['"]\]`)
		matches := regex.FindStringSubmatch(content)
		if len(matches) > 1 {
			return regex.ReplaceAllString(content, "globalThis['"+matches[1]+"']")
		}
	}

	return content
}

func containsProcessor(processors []Processor, path string) bool {
	for _, processor := range processors {
		if processor.Name == path {
			return true
		}
	}
	return false
}
