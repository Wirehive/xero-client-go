# RepeatingInvoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | See Invoice Types | [optional] 
**Contact** | [**Contact**](Contact.md) |  | [optional] 
**Schedule** | [**Schedule**](Schedule.md) |  | [optional] 
**LineItems** | [**[]LineItem**](LineItem.md) | See LineItems | [optional] 
**LineAmountTypes** | [**LineAmountTypes**](LineAmountTypes.md) |  | [optional] 
**Reference** | **string** | ACCREC only – additional reference number | [optional] 
**BrandingThemeID** | **string** | See BrandingThemes | [optional] 
**CurrencyCode** | [**CurrencyCode**](CurrencyCode.md) |  | [optional] 
**Status** | **string** | One of the following - DRAFT or AUTHORISED – See Invoice Status Codes | [optional] 
**SubTotal** | **float64** | Total of invoice excluding taxes | [optional] 
**TotalTax** | **float64** | Total tax on invoice | [optional] 
**Total** | **float64** | Total of Invoice tax inclusive (i.e. SubTotal + TotalTax) | [optional] 
**RepeatingInvoiceID** | **string** | Xero generated unique identifier for repeating invoice template | [optional] 
**ID** | **string** | Xero generated unique identifier for repeating invoice template | [optional] 
**HasAttachments** | **bool** | boolean to indicate if an invoice has an attachment | [optional] [readonly] [default to false]
**Attachments** | [**[]Attachment**](Attachment.md) | Displays array of attachments from the API | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


