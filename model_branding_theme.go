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
// BrandingTheme struct for BrandingTheme
type BrandingTheme struct {
	// Xero identifier
	BrandingThemeID string `json:"BrandingThemeID,omitempty"`
	// Name of branding theme
	Name string `json:"Name,omitempty"`
	// The location of the image file used as the logo on this branding theme
	LogoUrl string `json:"LogoUrl,omitempty"`
	// Always INVOICE
	Type string `json:"Type,omitempty"`
	// Integer – ranked order of branding theme. The default branding theme has a value of 0
	SortOrder int32 `json:"SortOrder,omitempty"`
	// UTC timestamp of creation date of branding theme
	CreatedDateUTC time.Time `json:"CreatedDateUTC,omitempty"`
}