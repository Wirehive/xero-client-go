# ManualJournalLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LineAmount** | **float64** | total for line. Debits are positive, credits are negative value | [optional] 
**AccountCode** | **string** | See Accounts | [optional] 
**Description** | **string** | Description for journal line | [optional] 
**TaxType** | **string** | The tax type from TaxRates | [optional] 
**Tracking** | [**[]TrackingCategory**](TrackingCategory.md) | Optional Tracking Category – see Tracking. Any JournalLine can have a maximum of 2 &lt;TrackingCategory&gt; elements. | [optional] 
**TaxAmount** | **float64** | The calculated tax amount based on the TaxType and LineAmount | [optional] [readonly] 
**IsBlank** | **bool** | is the line blank | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


