package rapid

type Config struct {
	testing            bool
	apiKey             string
	apiPassword        string
	payNowButtonApiKey string
}

func NewConfig(t bool, apiKey string, apiPassword string) *Config {
	return &Config{
		testing:            t,
		apiKey:             apiKey,
		apiPassword:        apiPassword,
		payNowButtonApiKey: "",
	}
}

func (c *Config) SetPublicApiKey(key string) *Config {
	c.payNowButtonApiKey = key
	return c
}
