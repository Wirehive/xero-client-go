# JournalLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JournalLineID** | **string** | Xero identifier for Journal | [optional] 
**AccountID** | **string** | See Accounts | [optional] 
**AccountCode** | **string** | See Accounts | [optional] 
**AccountType** | [**AccountType**](AccountType.md) |  | [optional] 
**AccountName** | **string** | See AccountCodes | [optional] 
**Description** | **string** | The description from the source transaction line item. Only returned if populated. | [optional] 
**NetAmount** | **float64** | Net amount of journal line. This will be a positive value for a debit and negative for a credit | [optional] 
**GrossAmount** | **float64** | Gross amount of journal line (NetAmount + TaxAmount). | [optional] 
**TaxAmount** | **float64** | Total tax on a journal line | [optional] [readonly] 
**TaxType** | **string** | The tax type from TaxRates | [optional] 
**TaxName** | **string** | see TaxRates | [optional] 
**TrackingCategories** | [**[]TrackingCategory**](TrackingCategory.md) | Optional Tracking Category – see Tracking. Any JournalLine can have a maximum of 2 &lt;TrackingCategory&gt; elements. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


