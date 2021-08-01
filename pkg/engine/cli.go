package engine

type Config struct {
	Port   int
	Prefix string
}

var engConfig *Config

// SetConfig : running config for workflow
func SetConfig(port int, prefix string) {
	engConfig = &Config{Port: port, Prefix: prefix}
}

// GetConfig : get existing config
func GetEngConfig() *Config {
	return engConfig
}
