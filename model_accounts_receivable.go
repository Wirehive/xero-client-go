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
// AccountsReceivable struct for AccountsReceivable
type AccountsReceivable struct {
	Outstanding float64 `json:"Outstanding,omitempty"`
	Overdue float64 `json:"Overdue,omitempty"`
}
