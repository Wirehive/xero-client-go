# ExpenseClaim

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpenseClaimID** | **string** | Xero generated unique identifier for an expense claim | [optional] 
**Status** | **string** | Current status of an expense claim – see status types | [optional] 
**Payments** | [**[]Payment**](Payment.md) | See Payments | [optional] 
**User** | [**User**](User.md) |  | [optional] 
**Receipts** | [**[]Receipt**](Receipt.md) |  | [optional] 
**UpdatedDateUTC** | [**time.Time**](time.Time.md) | Last modified date UTC format | [optional] [readonly] 
**Total** | **float64** | The total of an expense claim being paid | [optional] [readonly] 
**AmountDue** | **float64** | The amount due to be paid for an expense claim | [optional] [readonly] 
**AmountPaid** | **float64** | The amount still to pay for an expense claim | [optional] [readonly] 
**PaymentDueDate** | **string** | The date when the expense claim is due to be paid YYYY-MM-DD | [optional] [readonly] 
**ReportingDate** | **string** | The date the expense claim will be reported in Xero YYYY-MM-DD | [optional] [readonly] 
**ReceiptID** | **string** | The Xero identifier for the Receipt e.g.  e59a2c7f-1306-4078-a0f3-73537afcbba9 | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


