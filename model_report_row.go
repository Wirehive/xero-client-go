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
// ReportRow struct for ReportRow
type ReportRow struct {
	RowType RowType `json:"RowType,omitempty"`
	Title string `json:"Title,omitempty"`
	Cells []ReportCell `json:"Cells,omitempty"`
}