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
// Invoice struct for Invoice
type Invoice struct {
	// See Invoice Types
	Type string `json:"Type,omitempty"`
	Contact Contact `json:"Contact,omitempty"`
	// See LineItems
	LineItems []LineItem `json:"LineItems,omitempty"`
	// Date invoice was issued – YYYY-MM-DD. If the Date element is not specified it will default to the current date based on the timezone setting of the organisation
	Date string `json:"Date,omitempty"`
	// Date invoice is due – YYYY-MM-DD
	DueDate string `json:"DueDate,omitempty"`
	LineAmountTypes LineAmountTypes `json:"LineAmountTypes,omitempty"`
	// ACCREC – Unique alpha numeric code identifying invoice (when missing will auto-generate from your Organisation Invoice Settings) (max length = 255)
	InvoiceNumber string `json:"InvoiceNumber,omitempty"`
	// ACCREC only – additional reference number (max length = 255)
	Reference string `json:"Reference,omitempty"`
	// See BrandingThemes
	BrandingThemeID string `json:"BrandingThemeID,omitempty"`
	// URL link to a source document – shown as “Go to [appName]” in the Xero app
	Url string `json:"Url,omitempty"`
	CurrencyCode CurrencyCode `json:"CurrencyCode,omitempty"`
	// The currency rate for a multicurrency invoice. If no rate is specified, the XE.com day rate is used. (max length = [18].[6])
	CurrencyRate float64 `json:"CurrencyRate,omitempty"`
	// See Invoice Status Codes
	Status string `json:"Status,omitempty"`
	// Boolean to set whether the invoice in the Xero app should be marked as “sent”. This can be set only on invoices that have been approved
	SentToContact bool `json:"SentToContact,omitempty"`
	// Shown on sales invoices (Accounts Receivable) when this has been set
	ExpectedPaymentDate string `json:"ExpectedPaymentDate,omitempty"`
	// Shown on bills (Accounts Payable) when this has been set
	PlannedPaymentDate string `json:"PlannedPaymentDate,omitempty"`
	// CIS deduction for UK contractors
	CISDeduction float64 `json:"CISDeduction,omitempty"`
	// Total of invoice excluding taxes
	SubTotal float64 `json:"SubTotal,omitempty"`
	// Total tax on invoice
	TotalTax float64 `json:"TotalTax,omitempty"`
	// Total of Invoice tax inclusive (i.e. SubTotal + TotalTax). This will be ignored if it doesn’t equal the sum of the LineAmounts
	Total float64 `json:"Total,omitempty"`
	// Total of discounts applied on the invoice line items
	TotalDiscount float64 `json:"TotalDiscount,omitempty"`
	// Xero generated unique identifier for invoice
	InvoiceID string `json:"InvoiceID,omitempty"`
	// boolean to indicate if an invoice has an attachment
	HasAttachments bool `json:"HasAttachments,omitempty"`
	// boolean to indicate if an invoice has a discount
	IsDiscounted bool `json:"IsDiscounted,omitempty"`
	// See Payments
	Payments []Payment `json:"Payments,omitempty"`
	// See Prepayments
	Prepayments []Prepayment `json:"Prepayments,omitempty"`
	// See Overpayments
	Overpayments []Overpayment `json:"Overpayments,omitempty"`
	// Amount remaining to be paid on invoice
	AmountDue float64 `json:"AmountDue,omitempty"`
	// Sum of payments received for invoice
	AmountPaid float64 `json:"AmountPaid,omitempty"`
	// The date the invoice was fully paid. Only returned on fully paid invoices
	FullyPaidOnDate string `json:"FullyPaidOnDate,omitempty"`
	// Sum of all credit notes, over-payments and pre-payments applied to invoice
	AmountCredited float64 `json:"AmountCredited,omitempty"`
	// Last modified date UTC format
	UpdatedDateUTC time.Time `json:"UpdatedDateUTC,omitempty"`
	// Details of credit notes that have been applied to an invoice
	CreditNotes []CreditNote `json:"CreditNotes,omitempty"`
	// Displays array of attachments from the API
	Attachments []Attachment `json:"Attachments,omitempty"`
	// A boolean to indicate if a invoice has an validation errors
	HasErrors bool `json:"HasErrors,omitempty"`
	// A string to indicate if a invoice status
	StatusAttributeString string `json:"StatusAttributeString,omitempty"`
	// Displays array of validation error messages from the API
	ValidationErrors []ValidationError `json:"ValidationErrors,omitempty"`
	// Displays array of warning messages from the API
	Warnings []ValidationError `json:"Warnings,omitempty"`
}
