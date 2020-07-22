# BatchPayment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | [**Account**](Account.md) |  | [optional] 
**Reference** | **string** | (NZ Only) Optional references for the batch payment transaction. It will also show with the batch payment transaction in the bank reconciliation Find &amp; Match screen. Depending on your individual bank, the detail may also show on the bank statement you import into Xero. | [optional] 
**Particulars** | **string** | (NZ Only) Optional references for the batch payment transaction. It will also show with the batch payment transaction in the bank reconciliation Find &amp; Match screen. Depending on your individual bank, the detail may also show on the bank statement you import into Xero. | [optional] 
**Code** | **string** | (NZ Only) Optional references for the batch payment transaction. It will also show with the batch payment transaction in the bank reconciliation Find &amp; Match screen. Depending on your individual bank, the detail may also show on the bank statement you import into Xero. | [optional] 
**Details** | **string** | (Non-NZ Only) These details are sent to the org’s bank as a reference for the batch payment transaction. They will also show with the batch payment transaction in the bank reconciliation Find &amp; Match screen. Depending on your individual bank, the detail may also show on the bank statement imported into Xero. Maximum field length &#x3D; 18 | [optional] 
**Narrative** | **string** | (UK Only) Only shows on the statement line in Xero. Max length &#x3D;18 | [optional] 
**BatchPaymentID** | **string** | The Xero generated unique identifier for the bank transaction (read-only) | [optional] [readonly] 
**DateString** | **string** | Date the payment is being made (YYYY-MM-DD) e.g. 2009-09-06 | [optional] 
**Date** | **string** | Date the payment is being made (YYYY-MM-DD) e.g. 2009-09-06 | [optional] 
**Amount** | **float64** | The amount of the payment. Must be less than or equal to the outstanding amount owing on the invoice e.g. 200.00 | [optional] 
**Payments** | [**[]Payment**](Payment.md) |  | [optional] 
**Type** | **string** | PAYBATCH for bill payments or RECBATCH for sales invoice payments (read-only) | [optional] [readonly] 
**Status** | **string** | AUTHORISED or DELETED (read-only). New batch payments will have a status of AUTHORISED. It is not possible to delete batch payments via the API. | [optional] [readonly] 
**TotalAmount** | **string** | The total of the payments that make up the batch (read-only) | [optional] [readonly] 
**UpdatedDateUTC** | [**time.Time**](time.Time.md) | UTC timestamp of last update to the payment | [optional] [readonly] 
**IsReconciled** | **string** | Booelan that tells you if the batch payment has been reconciled (read-only) | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


