# Golang client for eWay payment provider.

Example of usage
```go
config := rapid.NewConfig(true, "API_KEY_HERE", "API_PASSWORD_HERE")
// In case you are going to use DirectConnection you should also add:
config.SetPublicApiKey("HERE_IS_YOUR_PUBLIC_KEY")
client, err := rapid.NewClient(nil, config)
if err != nil {
    panic(err)
}
```
There are three services:
1. Transaction
2. Encryption
3. Customer(WIP)

Example of creating transaction:
```go
    directConnection := &rapid.Transaction{
		Method:          "ProcessPayment",
		TransactionType: "Purchase",
		Payment: rapid.Payment{
			TotalAmount: 1000000,
			CurrencyCode: "NZD",
			InvoiceReference: "uni-1",
		},
		Customer: rapid.Customer{
			CardDetails: rapid.CardDetails{
				Name:        "John Doe",
				Number:      "4444333322221111",
				ExpiryMonth: "12",
				ExpiryYear:  "22",
				CVN:         "123",
			},
		},
	}
	
    transaction, err := client.Transaction.DirectConnection(directConnection, client.Encryption)
    if err != nil {
        panic(err)
    }
    fmt.Println(transaction)
```