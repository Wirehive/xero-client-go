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
// Currency struct for Currency
type Currency struct {
	Code CurrencyCode `json:"Code,omitempty"`
	// Name of Currency
	Description string `json:"Description,omitempty"`
}
