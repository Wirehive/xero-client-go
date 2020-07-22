# Journal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JournalID** | **string** | Xero identifier | [optional] 
**JournalDate** | **string** | Date the journal was posted | [optional] 
**JournalNumber** | **int32** | Xero generated journal number | [optional] 
**CreatedDateUTC** | [**time.Time**](time.Time.md) | Created date UTC format | [optional] [readonly] 
**Reference** | **string** | reference field for additional indetifying information | [optional] 
**SourceID** | **string** | The identifier for the source transaction (e.g. InvoiceID) | [optional] 
**SourceType** | **string** | The journal source type. The type of transaction that created the journal | [optional] 
**JournalLines** | [**[]JournalLine**](JournalLine.md) | See JournalLines | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


