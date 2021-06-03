package rapid

type Config struct {
	testing     bool
	apiKey      string
	apiPassword string
}

func NewConfig(t bool, apiKey string, apiPassword string) *Config {
	return &Config{
		testing:     t,
		apiKey:      apiKey,
		apiPassword: apiPassword,
	}
}
