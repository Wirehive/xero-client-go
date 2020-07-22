# Prepayment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | See Prepayment Types | [optional] 
**Contact** | [**Contact**](Contact.md) |  | [optional] 
**Date** | **string** | The date the prepayment is created YYYY-MM-DD | [optional] 
**Status** | **string** | See Prepayment Status Codes | [optional] 
**LineAmountTypes** | [**LineAmountTypes**](LineAmountTypes.md) |  | [optional] 
**LineItems** | [**[]LineItem**](LineItem.md) | See Prepayment Line Items | [optional] 
**SubTotal** | **float64** | The subtotal of the prepayment excluding taxes | [optional] 
**TotalTax** | **float64** | The total tax on the prepayment | [optional] 
**Total** | **float64** | The total of the prepayment(subtotal + total tax) | [optional] 
**Reference** | **string** | Returns Invoice number field. Reference field isn&#39;t available. | [optional] [readonly] 
**UpdatedDateUTC** | [**time.Time**](time.Time.md) | UTC timestamp of last update to the prepayment | [optional] [readonly] 
**CurrencyCode** | [**CurrencyCode**](CurrencyCode.md) |  | [optional] 
**PrepaymentID** | **string** | Xero generated unique identifier | [optional] 
**CurrencyRate** | **float64** | The currency rate for a multicurrency prepayment. If no rate is specified, the XE.com day rate is used | [optional] 
**RemainingCredit** | **float64** | The remaining credit balance on the prepayment | [optional] 
**Allocations** | [**[]Allocation**](Allocation.md) | See Allocations | [optional] 
**AppliedAmount** | **float64** | The amount of applied to an invoice | [optional] 
**HasAttachments** | **bool** | boolean to indicate if a prepayment has an attachment | [optional] [readonly] [default to false]
**Attachments** | [**[]Attachment**](Attachment.md) | See Attachments | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


