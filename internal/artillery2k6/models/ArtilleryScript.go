package models

type ArtilleryScript struct {
	Path      string     `yaml:"-"`
	Config    Config     `yaml:"config"`
	Scenarios []Scenario `yaml:"scenarios"`
}

type Config struct {
	Target       string                 `yaml:"target"`
	Processor    string                 `yaml:"processor"`
	Phases       []Phase                `yaml:"phases"`
	Environments map[string]Environment `yaml:"environments"`
	Variables    map[string]any         `yaml:"variables"`
	Payloads     Payload                `yaml:"payload"`
}

type Phase struct {
	Name         *string `yaml:"name,omitempty"`
	Duration     *int    `yaml:"duration,omitempty"`
	ArrivalRate  *int    `yaml:"arrivalRate,omitempty"`
	RampTo       *int    `yaml:"rampTo,omitempty"`
	ArrivalCount *int    `yaml:"arrivalCount,omitempty"`
}

type Environment struct {
	Target    string         `yaml:"target,omitempty" json:"-"`
	Phases    []Phase        `yaml:"phases,omitempty" json:"-"`
	Variables map[string]any `yaml:"variables,omitempty" json:"variables,omitempty"`
	Payload   Payload        `yaml:"payload,omitempty" json:"payloads,omitempty"`
}

type Scenario struct {
	Name string `yaml:"name"`
	Flow Flow   `yaml:"flow"`
}

type Flow struct {
	FlowActions []FlowAction
}

func (f *Flow) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var rawActions []map[string]any
	if err := unmarshal(&rawActions); err != nil {
		return err
	}

	reqActionCount := 0

	for _, rawAction := range rawActions {
		for key, value := range rawAction {
			var action FlowAction
			switch key {
			case "log":
				action = &LogAction{}
			case "think":
				action = &ThinkAction{}
			case "function":
				action = &FunctionAction{}
			case "get":
				action = NewRequestAction(reqActionCount)
				reqActionCount++
			case "post", "put":
				action = &PostPutRequestAction{RequestAction: NewRequestAction(reqActionCount)}
				reqActionCount++
			}
			if err := action.Build(key, value); err != nil {
				return err
			}
			f.FlowActions = append(f.FlowActions, action)
		}
	}
	return nil
}
