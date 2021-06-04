package rapid

import (
	"encoding/json"
	"net/http"
)

type TransactionService service

type Transaction struct {
	RedirectUrl     string          `json:"RedirectUrl,omitempty"`
	CustomerIP      string          `json:"CustomerIP,omitempty"`
	Method          string          `json:"Method"`
	TransactionType string          `json:"TransactionType"`
	DeviceID        string          `json:"DeviceID,omitempty"`
	PartnerID       string          `json:"PartnerID,omitempty"`
	CheckoutPayment bool            `json:"CheckoutPayment,omitempty"`
	CheckoutUrl     string          `json:"CheckoutUrl,omitempty"`
	Capture         bool            `json:"Capture,omitempty"`
	SaveCustomer    bool            `json:"SaveCustomer,omitempty"`
	Payment         Payment         `json:"Payment,omitempty"`
	Customer        Customer        `json:"Customer,omitempty"`
	ShippingAddress ShippingAddress `json:"ShippingAddress,omitempty"`
}

type Payment struct {
	TotalAmount        int    `json:"TotalAmount"`
	InvoiceNumber      string `json:"InvoiceNumber,omitempty"`
	InvoiceDescription string `json:"InvoiceDescription,omitempty"`
	InvoiceReference   string `json:"InvoiceReference,omitempty"`
	CurrencyCode       string `json:"CurrencyCode,omitempty"`
}

type Customer struct {
	TokenCustomerID string `json:"TokenCustomerID"`
	Reference       string `json:"Reference,omitempty"`
	Title           string `json:"Title"`
	FirstName       string `json:"FirstName"`
	LastName        string `json:"LastName"`
	CompanyName     string `json:"CompanyName,omitempty"`
	JobDescription  string `json:"JobDescription,omitempty"`
	Address
	Mobile      string      `json:"Mobile,omitempty"`
	Url         string      `json:"Url,omitempty"`
	CardDetails CardDetails `json:"CardDetails,omitempty"`
}

type Address struct {
	Street1    string `json:"Street1,omitempty"`
	Street2    string `json:"Street2,omitempty"`
	City       string `json:"City,omitempty"`
	State      string `json:"State,omitempty"`
	PostalCode string `json:"PostalCode,omitempty"`
	Country    string `json:"Country,omitempty"`
	Email      string `json:"Email,omitempty"`
	Phone      string `json:"Phone,omitempty"`
	Fax        string `json:"Fax,omitempty"`
}

type ShippingAddress struct {
	ShippingMethod string `json:"ShippingMethod,omitempty"`
	Address
}

type CardDetails struct {
	Name        string `json:"Name"`
	Number      string `json:"Number"`
	ExpiryMonth int8   `json:"ExpiryMonth"`
	ExpiryYear  int8   `json:"ExpiryYear"`
	StartMonth  int8   `json:"StartMonth,omitempty"`
	StartYear   int8   `json:"StartYear,omitempty"`
	IssueNumber int8   `json:"IssueNumber,omitempty"`
	CVN         string `json:"CVN"`
}

type ResponseTransaction struct {
	AccessCode          string  `json:"AccessCode"`
	FormActionURL       string  `json:"FormActionURL"`
	CompleteCheckoutURL string  `json:"CompleteCheckoutURL"`
	Payment             Payment `json:"Payment"`
}

type ResponseDirectConnection struct {
	AuthorisationCode string  `json:"AuthorisationCode"`
	ResponseCode      string  `json:"ResponseCode"`
	ResponseMessage   string  `json:"ResponseMessage"`
	TransactionID     int32   `json:"TransactionID"`
	TransactionStatus bool    `json:"TransactionStatus"`
	TransactionType   string  `json:"TransactionType"`
	TotalAmount       int     `json:"TotalAmount"`
	BeagleScore       string  `json:"BeagleScore"`
	Errors            string  `json:"Errors"`
	Payment           Payment `json:"Payment"`
}

type Context struct{}

func (ts *TransactionService) AccessCodes(t *Transaction) (rt *ResponseTransaction, err error) {
	req, err := ts.client.NewAPIRequest(http.MethodPost, "AccessCodes", t)
	if err != nil {
		return
	}

	res, err := ts.client.Do(req)
	if err != nil {
		panic(err)
	}
	if err = json.Unmarshal(res.content, &rt); err != nil {
		return
	}

	return
}

func (ts *TransactionService) DirectConnection(t *Transaction, encryptionService *EncryptionService) (rt *ResponseTransaction, err error) {
	encryptionService.EncryptCardDetails(t)
	req, err := ts.client.NewAPIRequest(http.MethodPost, "Transaction", t)

	if err != nil {
		return
	}

	res, err := ts.client.Do(req)
	if err != nil {
		panic(err)
	}
	if err = json.Unmarshal(res.content, &rt); err != nil {
		return
	}

	return
}
