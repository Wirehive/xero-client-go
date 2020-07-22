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
// Address struct for Address
type Address struct {
	// define the type of address
	AddressType string `json:"AddressType,omitempty"`
	// max length = 500
	AddressLine1 string `json:"AddressLine1,omitempty"`
	// max length = 500
	AddressLine2 string `json:"AddressLine2,omitempty"`
	// max length = 500
	AddressLine3 string `json:"AddressLine3,omitempty"`
	// max length = 500
	AddressLine4 string `json:"AddressLine4,omitempty"`
	// max length = 255
	City string `json:"City,omitempty"`
	// max length = 255
	Region string `json:"Region,omitempty"`
	// max length = 50
	PostalCode string `json:"PostalCode,omitempty"`
	// max length = 50, [A-Z], [a-z] only
	Country string `json:"Country,omitempty"`
	// max length = 255
	AttentionTo string `json:"AttentionTo,omitempty"`
}