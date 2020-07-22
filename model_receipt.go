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
// Receipt struct for Receipt
type Receipt struct {
	// Date of receipt – YYYY-MM-DD
	Date string `json:"Date,omitempty"`
	Contact Contact `json:"Contact,omitempty"`
	LineItems []LineItem `json:"LineItems,omitempty"`
	User User `json:"User,omitempty"`
	// Additional reference number
	Reference string `json:"Reference,omitempty"`
	LineAmountTypes LineAmountTypes `json:"LineAmountTypes,omitempty"`
	// Total of receipt excluding taxes
	SubTotal float64 `json:"SubTotal,omitempty"`
	// Total tax on receipt
	TotalTax float64 `json:"TotalTax,omitempty"`
	// Total of receipt tax inclusive (i.e. SubTotal + TotalTax)
	Total float64 `json:"Total,omitempty"`
	// Xero generated unique identifier for receipt
	ReceiptID string `json:"ReceiptID,omitempty"`
	// Current status of receipt – see status types
	Status string `json:"Status,omitempty"`
	// Xero generated sequence number for receipt in current claim for a given user
	ReceiptNumber string `json:"ReceiptNumber,omitempty"`
	// Last modified date UTC format
	UpdatedDateUTC time.Time `json:"UpdatedDateUTC,omitempty"`
	// boolean to indicate if a receipt has an attachment
	HasAttachments bool `json:"HasAttachments,omitempty"`
	// URL link to a source document – shown as “Go to [appName]” in the Xero app
	Url string `json:"Url,omitempty"`
	// Displays array of validation error messages from the API
	ValidationErrors []ValidationError `json:"ValidationErrors,omitempty"`
	// Displays array of warning messages from the API
	Warnings []ValidationError `json:"Warnings,omitempty"`
	// Displays array of attachments from the API
	Attachments []Attachment `json:"Attachments,omitempty"`
}