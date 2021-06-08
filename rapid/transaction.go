package rapid

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type TransactionService service

type TransactionsResponseDirectConnection struct {
	Transactions []ResponseDirectConnection
}

type Transactions struct {
	Transactions []Transaction
}

type Transaction struct {
	TransactionID   string          `json:"TransactionID,omitempty"`
	RedirectUrl     string          `json:"RedirectUrl,omitempty"`
	CustomerIP      string          `json:"CustomerIP,omitempty"`
	Method          string          `json:"Method"`
	TransactionType string          `json:"TransactionType,omitempty"`
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
	ExpiryMonth string `json:"ExpiryMonth"`
	ExpiryYear  string `json:"ExpiryYear"`
	StartMonth  string `json:"StartMonth,omitempty"`
	StartYear   string `json:"StartYear,omitempty"`
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
	AuthorisationCode     string             `json:"AuthorisationCode,omitempty"`
	ResponseCode          string             `json:"ResponseCode,omitempty"`
	ResponseMessage       string             `json:"ResponseMessage,omitempty"`
	InvoiceNumber         string             `json:"InvoiceNumber,omitempty"`
	InvoiceReference      string             `json:"InvoiceReference,omitempty"`
	TransactionID         int32              `json:"TransactionID,omitempty"`
	TransactionStatus     bool               `json:"TransactionStatus,omitempty"`
	TransactionType       int                `json:"TransactionType,omitempty"`
	TotalAmount           int                `json:"TotalAmount,omitempty"`
	TokenCustomerID       string             `json:"TokenCustomerID,omitempty"`
	BeagleScore           string             `json:"BeagleScore,omitempty"`
	TransactionDateTime   time.Time          `json:"TransactionDateTime,omitempty"`
	FraudAction           string             `json:"FraudAction,omitempty"`
	TransactionCaptured   bool               `json:"TransactionCaptured,omitempty"`
	CurrencyCode          string             `json:"CurrencyCode,omitempty"`
	Source                int                `json:"Source,omitempty"`
	Errors                string             `json:"Errors,omitempty"`
	MaxRefund             int                `json:"MaxRefund,omitempty"`
	OriginalTransactionId int                `json:"OriginalTransactionId,omitempty"`
	Payment               Payment            `json:"Payment,omitempty"`
	Customer              Customer           `json:"Customer,omitempty"`
	Options               []Option           `json:"Options,omitempty"`
	Verification          Verification       `json:"Verification,omitempty"`
	BeagleVerification    BeagleVerification `json:"BeagleVerification,omitempty"`
}

type Verification struct {
	CVN     int `json:"CVN,omitempty"`
	Address int `json:"Address,omitempty"`
	Email   int `json:"Email,omitempty"`
	Mobile  int `json:"Mobile,omitempty"`
	Phone   int `json:"Phone,omitempty"`
}

type BeagleVerification struct {
	Email int `json:"Email,omitempty"`
	Phone int `json:"Phone,omitempty"`
}

type Option struct {
	Value string `json:"Value,omitempty"`
}

type Refund struct {
	TotalAmount int `json:"TotalAmount"`
}

type Context struct{}

func (ts *TransactionService) AccessCodes(t *Transaction) (rt *ResponseTransaction, err error) {
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

func (ts *TransactionService) GetTransaction(identifier string, method string) (tr *TransactionsResponseDirectConnection, err error) {
	u := fmt.Sprintf("%s/%s", method, identifier)
	req, err := ts.client.NewAPIRequest(http.MethodGet, u, nil)
	if err != nil {
		panic(err)
	}

	res, err := ts.client.Do(req)
	if err != nil {
		panic(err)
	}
	if err = json.Unmarshal(res.content, &tr); err != nil {
		panic(err)
	}

	return
}

func (ts *TransactionService) GetTransactionByTransactionID(transactionID string) (tr *TransactionsResponseDirectConnection, err error) {
	u := fmt.Sprintf("Transaction/%s", transactionID)
	req, err := ts.client.NewAPIRequest(http.MethodGet, u, nil)
	if err != nil {
		panic(err)
	}

	res, err := ts.client.Do(req)
	if err != nil {
		panic(err)
	}
	if err = json.Unmarshal(res.content, &tr); err != nil {
		panic(err)
	}

	return
}

func (ts *TransactionService) GetTransactionsByInvoiceNumber(invoiceNumber string) (tr *TransactionsResponseDirectConnection, err error) {
	code, err := ts.GetTransaction(invoiceNumber, "Transaction/InvoiceNumber")
	if err != nil {
		return nil, err
	}
	return code, nil
}

func (ts *TransactionService) GetTransactionsByInvoiceReference(invoiceReference string) (tr *TransactionsResponseDirectConnection, err error) {
	code, err := ts.GetTransaction(invoiceReference, "Transaction/InvoiceRef")
	if err != nil {
		return nil, err
	}
	return code, nil
}

func (ts *TransactionService) Refund(transactionID string, r *Refund) (rt *ResponseTransaction, err error) {
	u := fmt.Sprintf("Transaction/%s/Refund", transactionID)
	req, err := ts.client.NewAPIRequest(http.MethodPost, u, r)
	if err != nil {
		panic(err)
	}

	res, err := ts.client.Do(req)
	if err != nil {
		panic(err)
	}
	if err = json.Unmarshal(res.content, &rt); err != nil {
		panic(err)
	}

	return
}
