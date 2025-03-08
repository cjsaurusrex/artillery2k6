package artillery2k6

import "github.com/cjsaurusrex/artillery2k6/internal/artillery2k6/helpers"

type K6Script struct {
	InitLifecycle InitLifecycle
	VULifecycle   VULifecycle
	Processor     []Processor
}

type InitLifecycle struct {
	Imports    []string
	Stages     []string
	Statements []string
}

type VULifecycle struct {
	Statements []string
}

type K6ScriptContext struct {
	Script K6Script
	Config helpers.BuilderConfig
}
