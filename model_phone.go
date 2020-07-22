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
// Phone struct for Phone
type Phone struct {
	PhoneType string `json:"PhoneType,omitempty"`
	// max length = 50
	PhoneNumber string `json:"PhoneNumber,omitempty"`
	// max length = 10
	PhoneAreaCode string `json:"PhoneAreaCode,omitempty"`
	// max length = 20
	PhoneCountryCode string `json:"PhoneCountryCode,omitempty"`
}