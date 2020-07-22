# Invoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | See Invoice Types | [optional] 
**Contact** | [**Contact**](Contact.md) |  | [optional] 
**LineItems** | [**[]LineItem**](LineItem.md) | See LineItems | [optional] 
**Date** | **string** | Date invoice was issued – YYYY-MM-DD. If the Date element is not specified it will default to the current date based on the timezone setting of the organisation | [optional] 
**DueDate** | **string** | Date invoice is due – YYYY-MM-DD | [optional] 
**LineAmountTypes** | [**LineAmountTypes**](LineAmountTypes.md) |  | [optional] 
**InvoiceNumber** | **string** | ACCREC – Unique alpha numeric code identifying invoice (when missing will auto-generate from your Organisation Invoice Settings) (max length &#x3D; 255) | [optional] 
**Reference** | **string** | ACCREC only – additional reference number (max length &#x3D; 255) | [optional] 
**BrandingThemeID** | **string** | See BrandingThemes | [optional] 
**Url** | **string** | URL link to a source document – shown as “Go to [appName]” in the Xero app | [optional] 
**CurrencyCode** | [**CurrencyCode**](CurrencyCode.md) |  | [optional] 
**CurrencyRate** | **float64** | The currency rate for a multicurrency invoice. If no rate is specified, the XE.com day rate is used. (max length &#x3D; [18].[6]) | [optional] 
**Status** | **string** | See Invoice Status Codes | [optional] 
**SentToContact** | **bool** | Boolean to set whether the invoice in the Xero app should be marked as “sent”. This can be set only on invoices that have been approved | [optional] 
**ExpectedPaymentDate** | **string** | Shown on sales invoices (Accounts Receivable) when this has been set | [optional] 
**PlannedPaymentDate** | **string** | Shown on bills (Accounts Payable) when this has been set | [optional] 
**CISDeduction** | **float64** | CIS deduction for UK contractors | [optional] [readonly] 
**SubTotal** | **float64** | Total of invoice excluding taxes | [optional] [readonly] 
**TotalTax** | **float64** | Total tax on invoice | [optional] [readonly] 
**Total** | **float64** | Total of Invoice tax inclusive (i.e. SubTotal + TotalTax). This will be ignored if it doesn’t equal the sum of the LineAmounts | [optional] [readonly] 
**TotalDiscount** | **float64** | Total of discounts applied on the invoice line items | [optional] [readonly] 
**InvoiceID** | **string** | Xero generated unique identifier for invoice | [optional] 
**HasAttachments** | **bool** | boolean to indicate if an invoice has an attachment | [optional] [readonly] [default to false]
**IsDiscounted** | **bool** | boolean to indicate if an invoice has a discount | [optional] [readonly] 
**Payments** | [**[]Payment**](Payment.md) | See Payments | [optional] [readonly] 
**Prepayments** | [**[]Prepayment**](Prepayment.md) | See Prepayments | [optional] [readonly] 
**Overpayments** | [**[]Overpayment**](Overpayment.md) | See Overpayments | [optional] [readonly] 
**AmountDue** | **float64** | Amount remaining to be paid on invoice | [optional] [readonly] 
**AmountPaid** | **float64** | Sum of payments received for invoice | [optional] [readonly] 
**FullyPaidOnDate** | **string** | The date the invoice was fully paid. Only returned on fully paid invoices | [optional] [readonly] 
**AmountCredited** | **float64** | Sum of all credit notes, over-payments and pre-payments applied to invoice | [optional] [readonly] 
**UpdatedDateUTC** | [**time.Time**](time.Time.md) | Last modified date UTC format | [optional] [readonly] 
**CreditNotes** | [**[]CreditNote**](CreditNote.md) | Details of credit notes that have been applied to an invoice | [optional] [readonly] 
**Attachments** | [**[]Attachment**](Attachment.md) | Displays array of attachments from the API | [optional] 
**HasErrors** | **bool** | A boolean to indicate if a invoice has an validation errors | [optional] [default to false]
**StatusAttributeString** | **string** | A string to indicate if a invoice status | [optional] 
**ValidationErrors** | [**[]ValidationError**](ValidationError.md) | Displays array of validation error messages from the API | [optional] 
**Warnings** | [**[]ValidationError**](ValidationError.md) | Displays array of warning messages from the API | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


