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
// TrackingCategory struct for TrackingCategory
type TrackingCategory struct {
	// The Xero identifier for a tracking category e.g. 297c2dc5-cc47-4afd-8ec8-74990b8761e9
	TrackingCategoryID string `json:"TrackingCategoryID,omitempty"`
	// The Xero identifier for a tracking option e.g. dc54c220-0140-495a-b925-3246adc0075f
	TrackingOptionID string `json:"TrackingOptionID,omitempty"`
	// The name of the tracking category e.g. Department, Region (max length = 100)
	Name string `json:"Name,omitempty"`
	// The option name of the tracking option e.g. East, West (max length = 100)
	Option string `json:"Option,omitempty"`
	// The status of a tracking category
	Status string `json:"Status,omitempty"`
	// See Tracking Options
	Options []TrackingOption `json:"Options,omitempty"`
}
