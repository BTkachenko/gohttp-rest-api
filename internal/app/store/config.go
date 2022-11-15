package store

type Config struct {
	DatabaseURL string `toml:"database_url"`
}


//New Config ...
func NewConfig() *Config {
	return &Config{}
}