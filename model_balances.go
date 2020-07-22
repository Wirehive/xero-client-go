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
// Balances The raw AccountsReceivable(sales invoices) and AccountsPayable(bills) outstanding and overdue amounts, not converted to base currency (read only)
type Balances struct {
	AccountsReceivable AccountsReceivable `json:"AccountsReceivable,omitempty"`
	AccountsPayable AccountsPayable `json:"AccountsPayable,omitempty"`
}
