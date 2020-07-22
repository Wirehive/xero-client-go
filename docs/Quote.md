# Quote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuoteID** | **string** | QuoteID GUID is automatically generated and is returned after create or GET. | [optional] 
**QuoteNumber** | **string** | Unique alpha numeric code identifying a quote (Max Length &#x3D; 255) | [optional] 
**Reference** | **string** | Additional reference number | [optional] 
**Terms** | **string** | Terms of the quote | [optional] 
**Contact** | [**Contact**](Contact.md) |  | [optional] 
**LineItems** | [**[]LineItem**](LineItem.md) | See LineItems | [optional] 
**Date** | **string** | Date quote was issued – YYYY-MM-DD. If the Date element is not specified it will default to the current date based on the timezone setting of the organisation | [optional] 
**DateString** | **string** | Date the quote was issued (YYYY-MM-DD) | [optional] 
**ExpiryDate** | **string** | Date the quote expires – YYYY-MM-DD. | [optional] 
**ExpiryDateString** | **string** | Date the quote expires – YYYY-MM-DD. | [optional] 
**Status** | [**QuoteStatusCodes**](QuoteStatusCodes.md) |  | [optional] 
**CurrencyCode** | [**CurrencyCode**](CurrencyCode.md) |  | [optional] 
**CurrencyRate** | **float64** | The currency rate for a multicurrency quote | [optional] 
**SubTotal** | **float64** | Total of quote excluding taxes. | [optional] [readonly] 
**TotalTax** | **float64** | Total tax on quote | [optional] [readonly] 
**Total** | **float64** | Total of Quote tax inclusive (i.e. SubTotal + TotalTax). This will be ignored if it doesn’t equal the sum of the LineAmounts | [optional] [readonly] 
**TotalDiscount** | **float64** | Total of discounts applied on the quote line items | [optional] [readonly] 
**Title** | **string** | Title text for the quote | [optional] 
**Summary** | **string** | Summary text for the quote | [optional] 
**BrandingThemeID** | **string** | See BrandingThemes | [optional] 
**UpdatedDateUTC** | [**time.Time**](time.Time.md) | Last modified date UTC format | [optional] [readonly] 
**LineAmountTypes** | [**QuoteLineAmountTypes**](QuoteLineAmountTypes.md) |  | [optional] 
**StatusAttributeString** | **string** | A string to indicate if a invoice status | [optional] 
**ValidationErrors** | [**[]ValidationError**](ValidationError.md) | Displays array of validation error messages from the API | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


