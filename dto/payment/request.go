// Package payment request/response objects
package payment

// Request Request struct read from http request body
type Request struct {
	// ties each request to a merchant
	MerchantId int64 `json:"-"`
}
