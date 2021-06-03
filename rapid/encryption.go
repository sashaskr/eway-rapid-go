package rapid

import (
	"encoding/json"
	"net/http"
)

type EncryptionService service

type Encrypt struct {
	Method string `json:"Method,omitempty"`
	Items  []Item `json:"Items"`
}

type Item struct {
	Name  string `json:"Name"`
	Value string `json:"Value"`
}

type EncryptResponse struct {
	Encrypt
	Errors string `json:"Errors"`
}

func (es *EncryptionService) Encrypt(e *Encrypt) (er *EncryptResponse, err error) {
	e.Method = "eCrypt"
	req, err := es.client.NewAPIRequest(http.MethodPost, "encrypt", e)
	if err != nil {
		return
	}

	res, err := es.client.Do(req)
	if err = json.Unmarshal(res.content, &er); err != nil {
		return
	}

	return
}
