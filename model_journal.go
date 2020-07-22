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
// Journal struct for Journal
type Journal struct {
	// Xero identifier
	JournalID string `json:"JournalID,omitempty"`
	// Date the journal was posted
	JournalDate string `json:"JournalDate,omitempty"`
	// Xero generated journal number
	JournalNumber int32 `json:"JournalNumber,omitempty"`
	// Created date UTC format
	CreatedDateUTC time.Time `json:"CreatedDateUTC,omitempty"`
	// reference field for additional indetifying information
	Reference string `json:"Reference,omitempty"`
	// The identifier for the source transaction (e.g. InvoiceID)
	SourceID string `json:"SourceID,omitempty"`
	// The journal source type. The type of transaction that created the journal
	SourceType string `json:"SourceType,omitempty"`
	// See JournalLines
	JournalLines []JournalLine `json:"JournalLines,omitempty"`
}
