# Overpayment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | See Overpayment Types | [optional] 
**Contact** | [**Contact**](Contact.md) |  | [optional] 
**Date** | **string** | The date the overpayment is created YYYY-MM-DD | [optional] 
**Status** | **string** | See Overpayment Status Codes | [optional] 
**LineAmountTypes** | [**LineAmountTypes**](LineAmountTypes.md) |  | [optional] 
**LineItems** | [**[]LineItem**](LineItem.md) | See Overpayment Line Items | [optional] 
**SubTotal** | **float64** | The subtotal of the overpayment excluding taxes | [optional] 
**TotalTax** | **float64** | The total tax on the overpayment | [optional] 
**Total** | **float64** | The total of the overpayment (subtotal + total tax) | [optional] 
**UpdatedDateUTC** | [**time.Time**](time.Time.md) | UTC timestamp of last update to the overpayment | [optional] [readonly] 
**CurrencyCode** | [**CurrencyCode**](CurrencyCode.md) |  | [optional] 
**OverpaymentID** | **string** | Xero generated unique identifier | [optional] 
**CurrencyRate** | **float64** | The currency rate for a multicurrency overpayment. If no rate is specified, the XE.com day rate is used | [optional] 
**RemainingCredit** | **float64** | The remaining credit balance on the overpayment | [optional] 
**Allocations** | [**[]Allocation**](Allocation.md) | See Allocations | [optional] 
**AppliedAmount** | **float64** | The amount of applied to an invoice | [optional] 
**Payments** | [**[]Payment**](Payment.md) | See Payments | [optional] 
**HasAttachments** | **bool** | boolean to indicate if a overpayment has an attachment | [optional] [readonly] [default to false]
**Attachments** | [**[]Attachment**](Attachment.md) | See Attachments | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


