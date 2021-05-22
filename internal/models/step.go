package models

import "github.com/coredevelopment/workflow/internal/constants"

type Step struct {
	Name string             `json:"name" yaml:"name"`
	ID   string             `json:"id" yaml:"id"`
	Type constants.StepType `json:"type" yaml:"type"`
	logicStep
	APIStep
	Async         bool      `json:"async" yaml:"async"`
	Delay         string    `json:"delay" yaml:"delay"`
	Timeout       string    `json:"timeout" yaml:"timeout"`
	NextStep      string    `json:"next-step" yaml:"next-step"`
	Break         bool      `json:"break" yaml:"break"`
	Error         StepError `json:"on-error" yaml:"on-error"`
	PreCondition  string    `json:"pre-condition" yaml:"pre-condition"`
	PostCondition string    `json:"post-condition" yaml:"post-condition"`
}

type logicStep struct {
	Runtime constants.Runtime `json:"runtime" yaml:"runtime"`
	ExePath string            `json:"exe-path" yaml:"exe-path"`
	Handler string            `json:"handler" yaml:"handler"`
}

type APIStep struct {
	Endpoint       string            `json:"endpoint" yaml:"endpoint"`
	Method         string            `json:"method" yaml:"method"`
	Payload        string            `json:"payload" yaml:"payload"`
	IncludeHeaders bool              `json:"include-headers" yaml:"include-headers"`
	CustomHeaders  map[string]string `json:"custom-headers" yaml:"custom-headers"`
}

func (l *logicStep) Execute() {

}
func (a *APIStep) Execute() {

}

type StepError struct {
	Retry bool   `json:"retry" yaml:"retry"`
	Goto  string `json:"goto" yaml:"goto"`
}