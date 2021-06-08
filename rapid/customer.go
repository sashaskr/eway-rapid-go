package rapid

import (
	"encoding/json"
	"net/http"
)

type CustomerService service

type Customer struct {
	TokenCustomerID string `json:"TokenCustomerID"`
	Reference       string `json:"Reference,omitempty"`
	Title           string `json:"Title,omitempty"`
	FirstName       string `json:"FirstName,omitempty"`
	LastName        string `json:"LastName,omitempty"`
	CompanyName     string `json:"CompanyName,omitempty"`
	JobDescription  string `json:"JobDescription,omitempty"`
	Country         string `json:"Country,omitempty"`
	Mobile          string `json:"Mobile,omitempty"`
	Url             string `json:"Url,omitempty"`
	Address
	CardDetails CardDetails `json:"CardDetails,omitempty"`
}

type CustomerTokenRequest struct {
	Customer        Customer `json:"Customer"`
	Payment         Payment  `json:"Payment"`
	Method          string   `json:"Method"`
	TransactionType string   `json:"TransactionType"`
}

type Customers struct {
	Customers []Customer `json:"Customers,omitempty"`
}

func (cs *CustomerService) CreateToken(customer *Customer) (c *Customers, err error) {
	var customerRequest = &CustomerTokenRequest{
		Customer: *customer,
		//Payment:         Payment{TotalAmount: 0},
		Method:          "CreateTokenCustomer",
		TransactionType: "Purchase",
	}
	req, err := cs.client.NewAPIRequest(http.MethodPost, "AccessCodes", customerRequest)
	if err != nil {
		panic(err)
	}
	res, err := cs.client.Do(req)
	if err != nil {
		panic(err)
	}
	if err = json.Unmarshal(res.content, &c); err != nil {
		panic(err)
	}

	return
}
