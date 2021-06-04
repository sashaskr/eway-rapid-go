package rapid

import (
	"net/http"
	"net/url"
	"runtime"
	"strings"
)

func NewClient(baseClient *http.Client, c *Config) (rapid *Client, err error) {
	if baseClient == nil {
		baseClient = http.DefaultClient
	}

	u, _ := url.Parse(BaseURL)

	rapid = &Client{
		BaseURL: u,
		client:  baseClient,
		config:  c,
	}

	rapid.common.client = rapid

	// services for resources
	rapid.Transaction = (*TransactionService)(&rapid.common)
	rapid.Encryption = (*EncryptionService)(&rapid.common)
	// services end

	rapid.apiKey = c.apiKey
	rapid.apiPassword = c.apiPassword
	rapid.payNowButtonApiKey = c.payNowButtonApiKey
	rapid.userAgent = strings.Join([]string{
		runtime.GOOS,
		runtime.GOARCH,
		runtime.Version(),
	}, ";")
	return
}
