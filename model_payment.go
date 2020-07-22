/*
 * Accounting API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.2.8
 * Contact: api@xero.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
import (
	"time"
)
// Payment struct for Payment
type Payment struct {
	Invoice Invoice `json:"Invoice,omitempty"`
	CreditNote CreditNote `json:"CreditNote,omitempty"`
	Prepayment Prepayment `json:"Prepayment,omitempty"`
	Overpayment Overpayment `json:"Overpayment,omitempty"`
	// Number of invoice or credit note you are applying payment to e.g.INV-4003
	InvoiceNumber string `json:"InvoiceNumber,omitempty"`
	// Number of invoice or credit note you are applying payment to e.g. INV-4003
	CreditNoteNumber string `json:"CreditNoteNumber,omitempty"`
	Account Account `json:"Account,omitempty"`
	// Code of account you are using to make the payment e.g. 001 (note- not all accounts have a code value)
	Code string `json:"Code,omitempty"`
	// Date the payment is being made (YYYY-MM-DD) e.g. 2009-09-06
	Date string `json:"Date,omitempty"`
	// Exchange rate when payment is received. Only used for non base currency invoices and credit notes e.g. 0.7500
	CurrencyRate float64 `json:"CurrencyRate,omitempty"`
	// The amount of the payment. Must be less than or equal to the outstanding amount owing on the invoice e.g. 200.00
	Amount float64 `json:"Amount,omitempty"`
	// An optional description for the payment e.g. Direct Debit
	Reference string `json:"Reference,omitempty"`
	// An optional parameter for the payment. A boolean indicating whether you would like the payment to be created as reconciled when using PUT, or whether a payment has been reconciled when using GET
	IsReconciled bool `json:"IsReconciled,omitempty"`
	// The status of the payment.
	Status string `json:"Status,omitempty"`
	// See Payment Types.
	PaymentType string `json:"PaymentType,omitempty"`
	// UTC timestamp of last update to the payment
	UpdatedDateUTC time.Time `json:"UpdatedDateUTC,omitempty"`
	// The Xero identifier for an Payment e.g. 297c2dc5-cc47-4afd-8ec8-74990b8761e9
	PaymentID string `json:"PaymentID,omitempty"`
	// The suppliers bank account number the payment is being made to
	BankAccountNumber string `json:"BankAccountNumber,omitempty"`
	// The suppliers bank account number the payment is being made to
	Particulars string `json:"Particulars,omitempty"`
	// The information to appear on the supplier's bank account
	Details string `json:"Details,omitempty"`
	// A boolean to indicate if a contact has an validation errors
	HasAccount bool `json:"HasAccount,omitempty"`
	// A boolean to indicate if a contact has an validation errors
	HasValidationErrors bool `json:"HasValidationErrors,omitempty"`
	// A string to indicate if a invoice status
	StatusAttributeString string `json:"StatusAttributeString,omitempty"`
	// Displays array of validation error messages from the API
	ValidationErrors []ValidationError `json:"ValidationErrors,omitempty"`
}
