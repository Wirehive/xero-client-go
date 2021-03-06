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
// Employee struct for Employee
type Employee struct {
	// The Xero identifier for an employee e.g. 297c2dc5-cc47-4afd-8ec8-74990b8761e9
	EmployeeID string `json:"EmployeeID,omitempty"`
	// Current status of an employee – see contact status types
	Status string `json:"Status,omitempty"`
	// First name of an employee (max length = 255)
	FirstName string `json:"FirstName,omitempty"`
	// Last name of an employee (max length = 255)
	LastName string `json:"LastName,omitempty"`
	ExternalLink ExternalLink `json:"ExternalLink,omitempty"`
	UpdatedDateUTC time.Time `json:"UpdatedDateUTC,omitempty"`
	// A string to indicate if a invoice status
	StatusAttributeString string `json:"StatusAttributeString,omitempty"`
	// Displays array of validation error messages from the API
	ValidationErrors []ValidationError `json:"ValidationErrors,omitempty"`
}
