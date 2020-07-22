# BankTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | See Bank Transaction Types | 
**Contact** | [**Contact**](Contact.md) |  | 
**LineItems** | [**[]LineItem**](LineItem.md) | See LineItems | 
**BankAccount** | [**Account**](Account.md) |  | 
**IsReconciled** | **bool** | Boolean to show if transaction is reconciled | [optional] 
**Date** | **string** | Date of transaction – YYYY-MM-DD | [optional] 
**Reference** | **string** | Reference for the transaction. Only supported for SPEND and RECEIVE transactions. | [optional] 
**CurrencyCode** | [**CurrencyCode**](CurrencyCode.md) |  | [optional] 
**CurrencyRate** | **float64** | Exchange rate to base currency when money is spent or received. e.g.0.7500 Only used for bank transactions in non base currency. If this isn’t specified for non base currency accounts then either the user-defined rate (preference) or the XE.com day rate will be used. Setting currency is only supported on overpayments. | [optional] 
**Url** | **string** | URL link to a source document – shown as “Go to App Name” | [optional] 
**Status** | **string** | See Bank Transaction Status Codes | [optional] 
**LineAmountTypes** | [**LineAmountTypes**](LineAmountTypes.md) |  | [optional] 
**SubTotal** | **float64** | Total of bank transaction excluding taxes | [optional] 
**TotalTax** | **float64** | Total tax on bank transaction | [optional] 
**Total** | **float64** | Total of bank transaction tax inclusive | [optional] 
**BankTransactionID** | **string** | Xero generated unique identifier for bank transaction | [optional] 
**PrepaymentID** | **string** | Xero generated unique identifier for a Prepayment. This will be returned on BankTransactions with a Type of SPEND-PREPAYMENT or RECEIVE-PREPAYMENT | [optional] [readonly] 
**OverpaymentID** | **string** | Xero generated unique identifier for an Overpayment. This will be returned on BankTransactions with a Type of SPEND-OVERPAYMENT or RECEIVE-OVERPAYMENT | [optional] [readonly] 
**UpdatedDateUTC** | [**time.Time**](time.Time.md) | Last modified date UTC format | [optional] [readonly] 
**HasAttachments** | **bool** | Boolean to indicate if a bank transaction has an attachment | [optional] [readonly] [default to false]
**StatusAttributeString** | **string** | A string to indicate if a invoice status | [optional] 
**ValidationErrors** | [**[]ValidationError**](ValidationError.md) | Displays array of validation error messages from the API | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


