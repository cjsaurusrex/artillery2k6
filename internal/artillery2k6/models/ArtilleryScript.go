package models

type ArtilleryScript struct {
	Path      string     `yaml:"-"`
	Config    Config     `yaml:"config"`
	Scenarios []Scenario `yaml:"scenarios"`
}

type Config struct {
	Target    string  `yaml:"target"`
	Processor string  `yaml:"processor"`
	Phases    []Phase `yaml:"phases"`
	//Environments []Environment `yaml:"environments"`
	Variables map[string]any `yaml:"variables"`
}

type Phase struct {
	Name         *string `yaml:"name,omitempty"`
	Duration     *int    `yaml:"duration,omitempty"`
	ArrivalRate  *int    `yaml:"arrivalRate,omitempty"`
	RampTo       *int    `yaml:"rampTo,omitempty"`
	ArrivalCount *int    `yaml:"arrivalCount,omitempty"`
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

	for _, rawAction := range rawActions {
		for key, value := range rawAction {
			var action FlowAction
			switch key {
			case "log":
				action = &LogAction{}
			case "get":
				action = &RequestAction{}
			case "post", "put":
				action = &PostPutRequestAction{RequestAction: &RequestAction{}}
			}
			if err := action.Build(key, value); err != nil {
				return err
			}
			f.FlowActions = append(f.FlowActions, action)
		}
	}
	return nil
}
