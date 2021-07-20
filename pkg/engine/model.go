package engine

import (
	"github.com/corepackage/workflow/internal/constants"
)

// Workflow - Structure of workflow configuration
type Workflow struct {
	Name       string                 `json:"name" yaml:"name"`
	ID         string                 `json:"id" yaml:"id"`
	Version    string                 `json:"version" yaml:"version"`
	PrimaryKey []PrimaryKey           `json:"primary-key" yaml:"primary-key"`
	Authorizer *Authorizer            `json:"authorizer" yaml:"authorizer"`
	CORS       map[string]interface{} `json:"cors" yaml:"cors"`
	Steps      []Step                 `json:"steps" yaml:"steps"`
}

// PrimaryKey - Unique key to identify a single instance
type PrimaryKey struct {
	PKey  string `json:"key" yaml:"key"`
	Input string `json:"input" yaml:"input"`
}

// Authorizer - To specify authorization method for the respective workflow
type Authorizer struct {
	Type    string `json:"type" yaml:"type"`
	AKey    string `json:"key" yaml:"key"`
	Input   string `json:"input" yaml:"input"`
	ExePath string `json:"exe-path" yaml:"exe-path"`
	Handler string `json:"handler" yaml:"handler"`
}

// CORS - properties for the access control to workflow
type CORS struct {
	AllowOrigin  string `json:"allow-origin" yaml:"allow-origin" default:"*"`
	AllowMethods string `json:"allow-methods" yaml:"allow-methods" default:"*"`
	AllowHeaders string `json:"allow-headers" yaml:"allow-headers" default:"*"`
	MaxAge       int64  `json:"maxage" yaml:"maxage"`
}

// Step - It defines a single step
type Step struct {
	LogicStep     `yaml:",inline"`
	APIStep       `yaml:",inline"`
	Name          string             `json:"name" yaml:"name"`
	ID            string             `json:"id" yaml:"id"`
	Type          constants.StepType `json:"type" yaml:"type"`
	Async         bool               `json:"async" yaml:"async"`
	Authorize     bool               `json:"authorize" yaml:"authorize"`
	Delay         string             `json:"delay" yaml:"delay"`
	Timeout       string             `json:"timeout" yaml:"timeout"`
	NextStep      string             `json:"next-step" yaml:"next-step"`
	Users         []string           `json:"users" yaml:"users"`
	Break         bool               `json:"break" yaml:"break"`
	Error         *StepError         `json:"on-error" yaml:"on-error"`
	PreCondition  interface{}        `json:"pre-condition" yaml:"pre-condition"`
	PostCondition interface{}        `json:"post-condition" yaml:"post-condition"`
}

// StepError - properties defined for step error
type StepError struct {
	Retry bool   `json:"retry" yaml:"retry"`
	Goto  string `json:"goto" yaml:"goto"`
}
