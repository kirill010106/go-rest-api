package apiserver

type Config struct {
	BindAddr    string `toml:"bind_addr"`
	LogLevel    string `toml:"log_level"`
	DatabaseURL string `toml:"database_url"`
	SessionKey  string `toml:"session_key"` // Must be 32 bytes long.
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}
