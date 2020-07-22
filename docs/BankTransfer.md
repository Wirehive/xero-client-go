# BankTransfer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromBankAccount** | [**Account**](Account.md) |  | 
**ToBankAccount** | [**Account**](Account.md) |  | 
**Amount** | **float64** | amount of the transaction | 
**Date** | **string** | The date of the Transfer YYYY-MM-DD | [optional] 
**BankTransferID** | **string** | The identifier of the Bank Transfer | [optional] [readonly] 
**CurrencyRate** | **float64** | The currency rate | [optional] [readonly] 
**FromBankTransactionID** | **string** | The Bank Transaction ID for the source account | [optional] [readonly] 
**ToBankTransactionID** | **string** | The Bank Transaction ID for the destination account | [optional] [readonly] 
**HasAttachments** | **bool** | Boolean to indicate if a Bank Transfer has an attachment | [optional] [readonly] [default to false]
**CreatedDateUTC** | [**time.Time**](time.Time.md) | UTC timestamp of creation date of bank transfer | [optional] [readonly] 
**ValidationErrors** | [**[]ValidationError**](ValidationError.md) | Displays array of validation error messages from the API | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


