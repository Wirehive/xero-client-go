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
// RepeatingInvoice struct for RepeatingInvoice
type RepeatingInvoice struct {
	// See Invoice Types
	Type string `json:"Type,omitempty"`
	Contact Contact `json:"Contact,omitempty"`
	Schedule Schedule `json:"Schedule,omitempty"`
	// See LineItems
	LineItems []LineItem `json:"LineItems,omitempty"`
	LineAmountTypes LineAmountTypes `json:"LineAmountTypes,omitempty"`
	// ACCREC only – additional reference number
	Reference string `json:"Reference,omitempty"`
	// See BrandingThemes
	BrandingThemeID string `json:"BrandingThemeID,omitempty"`
	CurrencyCode CurrencyCode `json:"CurrencyCode,omitempty"`
	// One of the following - DRAFT or AUTHORISED – See Invoice Status Codes
	Status string `json:"Status,omitempty"`
	// Total of invoice excluding taxes
	SubTotal float64 `json:"SubTotal,omitempty"`
	// Total tax on invoice
	TotalTax float64 `json:"TotalTax,omitempty"`
	// Total of Invoice tax inclusive (i.e. SubTotal + TotalTax)
	Total float64 `json:"Total,omitempty"`
	// Xero generated unique identifier for repeating invoice template
	RepeatingInvoiceID string `json:"RepeatingInvoiceID,omitempty"`
	// Xero generated unique identifier for repeating invoice template
	ID string `json:"ID,omitempty"`
	// boolean to indicate if an invoice has an attachment
	HasAttachments bool `json:"HasAttachments,omitempty"`
	// Displays array of attachments from the API
	Attachments []Attachment `json:"Attachments,omitempty"`
}
