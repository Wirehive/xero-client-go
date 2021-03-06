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
// ContactPerson struct for ContactPerson
type ContactPerson struct {
	// First name of person
	FirstName string `json:"FirstName,omitempty"`
	// Last name of person
	LastName string `json:"LastName,omitempty"`
	// Email address of person
	EmailAddress string `json:"EmailAddress,omitempty"`
	// boolean to indicate whether contact should be included on emails with invoices etc.
	IncludeInEmails bool `json:"IncludeInEmails,omitempty"`
}
