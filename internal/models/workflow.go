package models

type Workflow struct {
	Name       string       `json:"name" yaml:"name"`
	ID         string       `json:"id" yaml:"id"`
	Version    string       `json:"version" yaml:"version"`
	PrimaryKey []PrimaryKey `json:"primary-key" yaml:"primary-key"`
	Authorizer Authorizer   `json:"authorizer" yaml:"authorizer"`
	CORS       interface{}  `json:"cors" yaml:"cors"`
	Steps      []Step       `json:"steps" yaml:"steps"`
}

type PrimaryKey struct {
	PKey  string `json:"key" yaml:"key"`
	Input string `json:"input" yaml:"input"`
}

type Authorizer struct {
	Type   string `json:"type" yaml:"type"`
	AKey   string `json:"key" yaml:"key"`
	Input  string `json:"input" yaml:"input"`
	Secret string `json:"secret" yaml:"secret"`
}

type CORS struct {
	AllowOrigin  string `json:"allow-origin" yaml:"allow-origin" default:"*"`
	AllowMethods string `json:"allow-methods" yaml:"allow-methods" default:"*"`
	AllowHeaders string `json:"allow-headers" yaml:"allow-headers" default:"*"`
	MaxAge       int64  `json:"maxage" yaml:"maxage"`
}
