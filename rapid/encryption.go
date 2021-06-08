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
	Items  []Item `json:"Items"`
	Errors string `json:"Errors"`
}

func (es *EncryptionService) Encrypt(e *Encrypt) (er *EncryptResponse, err error) {
	e.Method = "eCrypt"
	req, err := es.client.NewApiEncryptedRequest(http.MethodPost, "encrypt", e)
	if err != nil {
		return
	}

	res, err := es.client.Do(req)
	if err != nil {
		panic(err)
	}
	if err = json.Unmarshal(res.content, &er); err != nil {
		return
	}

	return
}

func (es *EncryptionService) EncryptCardDetails(t *Transaction) {
	encrypt := &Encrypt{
		Items: []Item{
			{
				Name:  "CVN",
				Value: t.Customer.CardDetails.CVN,
			},
			{
				Name:  "card",
				Value: t.Customer.CardDetails.Number,
			},
		},
	}
	items, err := es.Encrypt(encrypt)
	if err != nil {
		return
	}
	for _, item := range items.Items {
		if item.Name == "CVN" {
			t.Customer.CardDetails.CVN = item.Value
		}
		if item.Name == "card" {
			t.Customer.CardDetails.Number = item.Value
		}
	}

}
