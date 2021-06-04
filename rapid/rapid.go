package rapid

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	BaseURL            string = "https://api.sandbox.ewaypayments.com/"
	AuthHeader         string = "Authorization"
	RequestContentType string = "application/json"
	RequestAccept      string = "application/json"
	TokenType          string = "Basic"
	Connection         string = "keep-alive"
	APIKeyEnv          string = "EWAY_API_KEY"
	APIPasswordEnv     string = "EWAY_API_PASSWORD"
	XEwayAPIVersion    string = "X-EWAY-APIVERSION"
)

type Response struct {
	*http.Response
	content []byte
}

type Client struct {
	BaseURL            *url.URL
	apiKey             string
	apiPassword        string
	payNowButtonApiKey string
	userAgent          string
	client             *http.Client
	config             *Config
	common             service
	Transaction        *TransactionService
	Encryption         *EncryptionService
}

type service struct {
	client *Client
}

func (c *Client) WithAuthenticationValue(k string, p string) error {
	if k == "" {
		return errEmptyApiKey
	}

	if p == "" {
		return errEmptyApiPassword
	}

	c.apiKey = strings.TrimSpace(k)
	c.apiPassword = strings.TrimSpace(p)
	return nil
}

func (c *Client) NewAPIRequest(method string, uri string, body interface{}) (req *http.Request, err error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, errBadBaseURL
	}

	u, err := c.BaseURL.Parse(uri)
	if err != nil {
		return nil, err
	}

	if c.config.testing {
		u.Query().Add("testmode", "true")
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err = http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	var token = base64.StdEncoding.EncodeToString([]byte(strings.Join([]string{c.apiKey, c.apiPassword}, ":")))

	req.Header.Add(AuthHeader, strings.Join([]string{TokenType, token}, " "))
	req.Header.Add(XEwayAPIVersion, "40") // TODO: Move it to the config
	req.Header.Set("Content-Type", RequestContentType)
	req.Header.Set("Accept", RequestAccept)
	req.Header.Set("User-Agent", c.userAgent)
	req.Header.Set("Connection", Connection)
	return
}

func (c *Client) NewApiEncryptedRequest(method string, uri string, body interface{}) (req *http.Request, err error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, errBadBaseURL
	}

	u, err := c.BaseURL.Parse(uri)
	if err != nil {
		return nil, err
	}

	if c.config.testing {
		u.Query().Add("testmode", "true")
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err = http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	var token = base64.StdEncoding.EncodeToString([]byte(strings.Join([]string{c.payNowButtonApiKey, ""}, ":")))

	req.Header.Add(AuthHeader, strings.Join([]string{TokenType, token}, " "))
	req.Header.Add(XEwayAPIVersion, "40")
	req.Header.Set("Content-Type", RequestContentType)
	req.Header.Set("Accept", RequestAccept)
	req.Header.Set("User-Agent", c.userAgent)
	req.Header.Set("Connection", Connection)
	return
}

func (c *Client) Do(req *http.Request) (*Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	response, _ := newResponse(resp)
	err = CheckResponse(resp)
	if err != nil {
		return response, err
	}

	return response, nil
}

func newResponse(r *http.Response) (*Response, error) {
	var res Response
	c, err := ioutil.ReadAll(r.Body)
	if err == nil {
		res.content = c
	}
	err = json.NewDecoder(r.Body).Decode(&res)
	r.Body = ioutil.NopCloser(bytes.NewBuffer(c))
	res.Response = r
	return &res, err
}

func CheckResponse(r *http.Response) error {
	if r.StatusCode >= http.StatusMultipleChoices {
		return newError(r)
	}
	return nil
}
