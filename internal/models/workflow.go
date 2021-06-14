package models

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
