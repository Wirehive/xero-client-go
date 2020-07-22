# CreditNote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | See Credit Note Types | [optional] 
**Contact** | [**Contact**](Contact.md) |  | [optional] 
**Date** | **string** | The date the credit note is issued YYYY-MM-DD. If the Date element is not specified then it will default to the current date based on the timezone setting of the organisation | [optional] 
**Status** | **string** | See Credit Note Status Codes | [optional] 
**LineAmountTypes** | [**LineAmountTypes**](LineAmountTypes.md) |  | [optional] 
**LineItems** | [**[]LineItem**](LineItem.md) | See Invoice Line Items | [optional] 
**SubTotal** | **float64** | The subtotal of the credit note excluding taxes | [optional] 
**TotalTax** | **float64** | The total tax on the credit note | [optional] 
**Total** | **float64** | The total of the Credit Note(subtotal + total tax) | [optional] 
**UpdatedDateUTC** | [**time.Time**](time.Time.md) | UTC timestamp of last update to the credit note | [optional] [readonly] 
**CurrencyCode** | [**CurrencyCode**](CurrencyCode.md) |  | [optional] 
**FullyPaidOnDate** | **string** | Date when credit note was fully paid(UTC format) | [optional] 
**CreditNoteID** | **string** | Xero generated unique identifier | [optional] 
**CreditNoteNumber** | **string** | ACCRECCREDIT – Unique alpha numeric code identifying credit note (when missing will auto-generate from your Organisation Invoice Settings) | [optional] 
**Reference** | **string** | ACCRECCREDIT only – additional reference number | [optional] 
**SentToContact** | **bool** | boolean to indicate if a credit note has been sent to a contact via  the Xero app (currently read only) | [optional] [readonly] 
**CurrencyRate** | **float64** | The currency rate for a multicurrency invoice. If no rate is specified, the XE.com day rate is used | [optional] 
**RemainingCredit** | **float64** | The remaining credit balance on the Credit Note | [optional] 
**Allocations** | [**[]Allocation**](Allocation.md) | See Allocations | [optional] 
**AppliedAmount** | **float64** | The amount of applied to an invoice | [optional] 
**Payments** | [**[]Payment**](Payment.md) | See Payments | [optional] 
**BrandingThemeID** | **string** | See BrandingThemes | [optional] 
**StatusAttributeString** | **string** | A string to indicate if a invoice status | [optional] 
**HasAttachments** | **bool** | boolean to indicate if a credit note has an attachment | [optional] [default to false]
**HasErrors** | **bool** | A boolean to indicate if a credit note has an validation errors | [optional] [default to false]
**ValidationErrors** | [**[]ValidationError**](ValidationError.md) | Displays array of validation error messages from the API | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


