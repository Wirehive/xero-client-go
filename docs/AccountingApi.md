# \AccountingApi

All URIs are relative to *https://api.xero.com/api.xro/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccount**](AccountingApi.md#CreateAccount) | **Put** /Accounts | Allows you to create a new chart of accounts
[**CreateAccountAttachmentByFileName**](AccountingApi.md#CreateAccountAttachmentByFileName) | **Put** /Accounts/{AccountID}/Attachments/{FileName} | Allows you to create Attachment on Account
[**CreateBankTransactionAttachmentByFileName**](AccountingApi.md#CreateBankTransactionAttachmentByFileName) | **Put** /BankTransactions/{BankTransactionID}/Attachments/{FileName} | Allows you to createa an Attachment on BankTransaction by Filename
[**CreateBankTransactionHistoryRecord**](AccountingApi.md#CreateBankTransactionHistoryRecord) | **Put** /BankTransactions/{BankTransactionID}/History | Allows you to create history record for a bank transactions
[**CreateBankTransactions**](AccountingApi.md#CreateBankTransactions) | **Put** /BankTransactions | Allows you to create one or more spend or receive money transaction
[**CreateBankTransfer**](AccountingApi.md#CreateBankTransfer) | **Put** /BankTransfers | Allows you to create a bank transfers
[**CreateBankTransferAttachmentByFileName**](AccountingApi.md#CreateBankTransferAttachmentByFileName) | **Put** /BankTransfers/{BankTransferID}/Attachments/{FileName} | 
[**CreateBatchPayment**](AccountingApi.md#CreateBatchPayment) | **Put** /BatchPayments | Create one or many BatchPayments for invoices
[**CreateBatchPaymentHistoryRecord**](AccountingApi.md#CreateBatchPaymentHistoryRecord) | **Put** /BatchPayments/{BatchPaymentID}/History | Allows you to create a history record for a Batch Payment
[**CreateBrandingThemePaymentServices**](AccountingApi.md#CreateBrandingThemePaymentServices) | **Post** /BrandingThemes/{BrandingThemeID}/PaymentServices | Allow for the creation of new custom payment service for specified Branding Theme
[**CreateContactAttachmentByFileName**](AccountingApi.md#CreateContactAttachmentByFileName) | **Put** /Contacts/{ContactID}/Attachments/{FileName} | 
[**CreateContactGroup**](AccountingApi.md#CreateContactGroup) | **Put** /ContactGroups | Allows you to create a contact group
[**CreateContactGroupContacts**](AccountingApi.md#CreateContactGroupContacts) | **Put** /ContactGroups/{ContactGroupID}/Contacts | Allows you to add Contacts to a Contact Group
[**CreateContacts**](AccountingApi.md#CreateContacts) | **Put** /Contacts | Allows you to create a multiple contacts (bulk) in a Xero organisation
[**CreateCreditNoteAllocation**](AccountingApi.md#CreateCreditNoteAllocation) | **Put** /CreditNotes/{CreditNoteID}/Allocations | Allows you to create Allocation on CreditNote
[**CreateCreditNoteAttachmentByFileName**](AccountingApi.md#CreateCreditNoteAttachmentByFileName) | **Put** /CreditNotes/{CreditNoteID}/Attachments/{FileName} | Allows you to create Attachments on CreditNote by file name
[**CreateCreditNoteHistory**](AccountingApi.md#CreateCreditNoteHistory) | **Put** /CreditNotes/{CreditNoteID}/History | Allows you to retrieve a history records of an CreditNote
[**CreateCreditNotes**](AccountingApi.md#CreateCreditNotes) | **Put** /CreditNotes | Allows you to create a credit note
[**CreateCurrency**](AccountingApi.md#CreateCurrency) | **Put** /Currencies | 
[**CreateEmployees**](AccountingApi.md#CreateEmployees) | **Put** /Employees | Allows you to create new employees used in Xero payrun
[**CreateExpenseClaims**](AccountingApi.md#CreateExpenseClaims) | **Put** /ExpenseClaims | Allows you to retrieve expense claims
[**CreateInvoiceAttachmentByFileName**](AccountingApi.md#CreateInvoiceAttachmentByFileName) | **Put** /Invoices/{InvoiceID}/Attachments/{FileName} | Allows you to create an Attachment on invoices or purchase bills by it&#39;s filename
[**CreateInvoiceHistory**](AccountingApi.md#CreateInvoiceHistory) | **Put** /Invoices/{InvoiceID}/History | Allows you to retrieve a history records of an invoice
[**CreateInvoices**](AccountingApi.md#CreateInvoices) | **Put** /Invoices | Allows you to create one or more sales invoices or purchase bills
[**CreateItems**](AccountingApi.md#CreateItems) | **Put** /Items | Allows you to create one or more items
[**CreateLinkedTransaction**](AccountingApi.md#CreateLinkedTransaction) | **Put** /LinkedTransactions | Allows you to create linked transactions (billable expenses)
[**CreateManualJournalAttachmentByFileName**](AccountingApi.md#CreateManualJournalAttachmentByFileName) | **Put** /ManualJournals/{ManualJournalID}/Attachments/{FileName} | Allows you to create a specified Attachment on ManualJournal by file name
[**CreateManualJournals**](AccountingApi.md#CreateManualJournals) | **Put** /ManualJournals | Allows you to create one or more manual journals
[**CreateOverpaymentAllocations**](AccountingApi.md#CreateOverpaymentAllocations) | **Put** /Overpayments/{OverpaymentID}/Allocations | Allows you to create a single allocation for an overpayment
[**CreatePayment**](AccountingApi.md#CreatePayment) | **Post** /Payments | Allows you to create a single payment for invoices or credit notes
[**CreatePaymentService**](AccountingApi.md#CreatePaymentService) | **Put** /PaymentServices | Allows you to create payment services
[**CreatePayments**](AccountingApi.md#CreatePayments) | **Put** /Payments | Allows you to create multiple payments for invoices or credit notes
[**CreatePrepaymentAllocations**](AccountingApi.md#CreatePrepaymentAllocations) | **Put** /Prepayments/{PrepaymentID}/Allocations | Allows you to create an Allocation for prepayments
[**CreatePurchaseOrderHistory**](AccountingApi.md#CreatePurchaseOrderHistory) | **Put** /PurchaseOrders/{PurchaseOrderID}/History | Allows you to create HistoryRecord for purchase orders
[**CreatePurchaseOrders**](AccountingApi.md#CreatePurchaseOrders) | **Put** /PurchaseOrders | Allows you to create one or more purchase orders
[**CreateQuoteAttachmentByFileName**](AccountingApi.md#CreateQuoteAttachmentByFileName) | **Put** /Quotes/{QuoteID}/Attachments/{FileName} | Allows you to create Attachment on Quote
[**CreateQuoteHistory**](AccountingApi.md#CreateQuoteHistory) | **Put** /Quotes/{QuoteID}/History | Allows you to retrieve a history records of an quote
[**CreateQuotes**](AccountingApi.md#CreateQuotes) | **Put** /Quotes | Allows you to create one or more quotes
[**CreateReceipt**](AccountingApi.md#CreateReceipt) | **Put** /Receipts | Allows you to create draft expense claim receipts for any user
[**CreateReceiptAttachmentByFileName**](AccountingApi.md#CreateReceiptAttachmentByFileName) | **Put** /Receipts/{ReceiptID}/Attachments/{FileName} | Allows you to create Attachment on expense claim receipts by file name
[**CreateRepeatingInvoiceAttachmentByFileName**](AccountingApi.md#CreateRepeatingInvoiceAttachmentByFileName) | **Put** /RepeatingInvoices/{RepeatingInvoiceID}/Attachments/{FileName} | Allows you to create attachment on repeating invoices by file name
[**CreateTaxRates**](AccountingApi.md#CreateTaxRates) | **Put** /TaxRates | Allows you to create one or more Tax Rates
[**CreateTrackingCategory**](AccountingApi.md#CreateTrackingCategory) | **Put** /TrackingCategories | Allows you to create tracking categories
[**CreateTrackingOptions**](AccountingApi.md#CreateTrackingOptions) | **Put** /TrackingCategories/{TrackingCategoryID}/Options | Allows you to create options for a specified tracking category
[**DeleteAccount**](AccountingApi.md#DeleteAccount) | **Delete** /Accounts/{AccountID} | Allows you to delete a chart of accounts
[**DeleteContactGroupContact**](AccountingApi.md#DeleteContactGroupContact) | **Delete** /ContactGroups/{ContactGroupID}/Contacts/{ContactID} | Allows you to delete a specific Contact from a Contact Group
[**DeleteContactGroupContacts**](AccountingApi.md#DeleteContactGroupContacts) | **Delete** /ContactGroups/{ContactGroupID}/Contacts | Allows you to delete  all Contacts from a Contact Group
[**DeleteItem**](AccountingApi.md#DeleteItem) | **Delete** /Items/{ItemID} | Allows you to delete a specified item
[**DeleteLinkedTransaction**](AccountingApi.md#DeleteLinkedTransaction) | **Delete** /LinkedTransactions/{LinkedTransactionID} | Allows you to delete a specified linked transactions (billable expenses)
[**DeletePayment**](AccountingApi.md#DeletePayment) | **Post** /Payments/{PaymentID} | Allows you to update a specified payment for invoices and credit notes
[**DeleteTrackingCategory**](AccountingApi.md#DeleteTrackingCategory) | **Delete** /TrackingCategories/{TrackingCategoryID} | Allows you to delete tracking categories
[**DeleteTrackingOptions**](AccountingApi.md#DeleteTrackingOptions) | **Delete** /TrackingCategories/{TrackingCategoryID}/Options/{TrackingOptionID} | Allows you to delete a specified option for a specified tracking category
[**EmailInvoice**](AccountingApi.md#EmailInvoice) | **Post** /Invoices/{InvoiceID}/Email | Allows you to email a copy of invoice to related Contact
[**GetAccount**](AccountingApi.md#GetAccount) | **Get** /Accounts/{AccountID} | Allows you to retrieve a single chart of accounts
[**GetAccountAttachmentByFileName**](AccountingApi.md#GetAccountAttachmentByFileName) | **Get** /Accounts/{AccountID}/Attachments/{FileName} | Allows you to retrieve Attachment on Account by Filename
[**GetAccountAttachmentById**](AccountingApi.md#GetAccountAttachmentById) | **Get** /Accounts/{AccountID}/Attachments/{AttachmentID} | Allows you to retrieve specific Attachment on Account
[**GetAccountAttachments**](AccountingApi.md#GetAccountAttachments) | **Get** /Accounts/{AccountID}/Attachments | Allows you to retrieve Attachments for accounts
[**GetAccounts**](AccountingApi.md#GetAccounts) | **Get** /Accounts | Allows you to retrieve the full chart of accounts
[**GetBankTransaction**](AccountingApi.md#GetBankTransaction) | **Get** /BankTransactions/{BankTransactionID} | Allows you to retrieve a single spend or receive money transaction
[**GetBankTransactionAttachmentByFileName**](AccountingApi.md#GetBankTransactionAttachmentByFileName) | **Get** /BankTransactions/{BankTransactionID}/Attachments/{FileName} | Allows you to retrieve Attachments on BankTransaction by Filename
[**GetBankTransactionAttachmentById**](AccountingApi.md#GetBankTransactionAttachmentById) | **Get** /BankTransactions/{BankTransactionID}/Attachments/{AttachmentID} | Allows you to retrieve Attachments on a specific BankTransaction
[**GetBankTransactionAttachments**](AccountingApi.md#GetBankTransactionAttachments) | **Get** /BankTransactions/{BankTransactionID}/Attachments | Allows you to retrieve any attachments to bank transactions
[**GetBankTransactions**](AccountingApi.md#GetBankTransactions) | **Get** /BankTransactions | Allows you to retrieve any spend or receive money transactions
[**GetBankTransactionsHistory**](AccountingApi.md#GetBankTransactionsHistory) | **Get** /BankTransactions/{BankTransactionID}/History | Allows you to retrieve history from a bank transactions
[**GetBankTransfer**](AccountingApi.md#GetBankTransfer) | **Get** /BankTransfers/{BankTransferID} | Allows you to retrieve any bank transfers
[**GetBankTransferAttachmentByFileName**](AccountingApi.md#GetBankTransferAttachmentByFileName) | **Get** /BankTransfers/{BankTransferID}/Attachments/{FileName} | Allows you to retrieve Attachments on BankTransfer by file name
[**GetBankTransferAttachmentById**](AccountingApi.md#GetBankTransferAttachmentById) | **Get** /BankTransfers/{BankTransferID}/Attachments/{AttachmentID} | Allows you to retrieve Attachments on BankTransfer
[**GetBankTransferAttachments**](AccountingApi.md#GetBankTransferAttachments) | **Get** /BankTransfers/{BankTransferID}/Attachments | Allows you to retrieve Attachments from  bank transfers
[**GetBankTransferHistory**](AccountingApi.md#GetBankTransferHistory) | **Get** /BankTransfers/{BankTransferID}/History | Allows you to retrieve history from a bank transfers
[**GetBankTransfers**](AccountingApi.md#GetBankTransfers) | **Get** /BankTransfers | Allows you to retrieve all bank transfers
[**GetBatchPaymentHistory**](AccountingApi.md#GetBatchPaymentHistory) | **Get** /BatchPayments/{BatchPaymentID}/History | Allows you to retrieve history from a Batch Payment
[**GetBatchPayments**](AccountingApi.md#GetBatchPayments) | **Get** /BatchPayments | Retrieve either one or many BatchPayments for invoices
[**GetBrandingTheme**](AccountingApi.md#GetBrandingTheme) | **Get** /BrandingThemes/{BrandingThemeID} | Allows you to retrieve a specific BrandingThemes
[**GetBrandingThemePaymentServices**](AccountingApi.md#GetBrandingThemePaymentServices) | **Get** /BrandingThemes/{BrandingThemeID}/PaymentServices | Allows you to retrieve the Payment services for a Branding Theme
[**GetBrandingThemes**](AccountingApi.md#GetBrandingThemes) | **Get** /BrandingThemes | Allows you to retrieve all the BrandingThemes
[**GetContact**](AccountingApi.md#GetContact) | **Get** /Contacts/{ContactID} | Allows you to retrieve a single contacts in a Xero organisation
[**GetContactAttachmentByFileName**](AccountingApi.md#GetContactAttachmentByFileName) | **Get** /Contacts/{ContactID}/Attachments/{FileName} | Allows you to retrieve Attachments on Contacts by file name
[**GetContactAttachmentById**](AccountingApi.md#GetContactAttachmentById) | **Get** /Contacts/{ContactID}/Attachments/{AttachmentID} | Allows you to retrieve Attachments on Contacts
[**GetContactAttachments**](AccountingApi.md#GetContactAttachments) | **Get** /Contacts/{ContactID}/Attachments | Allows you to retrieve, add and update contacts in a Xero organisation
[**GetContactByContactNumber**](AccountingApi.md#GetContactByContactNumber) | **Get** /Contacts/{ContactNumber} | Allows you to retrieve a single contact by Contact Number in a Xero organisation
[**GetContactCISSettings**](AccountingApi.md#GetContactCISSettings) | **Get** /Contacts/{ContactID}/CISSettings | Allows you to retrieve CISSettings for a contact in a Xero organisation
[**GetContactGroup**](AccountingApi.md#GetContactGroup) | **Get** /ContactGroups/{ContactGroupID} | Allows you to retrieve a unique Contact Group by ID
[**GetContactGroups**](AccountingApi.md#GetContactGroups) | **Get** /ContactGroups | Allows you to retrieve the ContactID and Name of all the contacts in a contact group
[**GetContactHistory**](AccountingApi.md#GetContactHistory) | **Get** /Contacts/{ContactID}/History | Allows you to retrieve a history records of an Contact
[**GetContacts**](AccountingApi.md#GetContacts) | **Get** /Contacts | Allows you to retrieve all contacts in a Xero organisation
[**GetCreditNote**](AccountingApi.md#GetCreditNote) | **Get** /CreditNotes/{CreditNoteID} | Allows you to retrieve a specific credit note
[**GetCreditNoteAsPdf**](AccountingApi.md#GetCreditNoteAsPdf) | **Get** /CreditNotes/{CreditNoteID}/pdf | Allows you to retrieve Credit Note as PDF files
[**GetCreditNoteAttachmentByFileName**](AccountingApi.md#GetCreditNoteAttachmentByFileName) | **Get** /CreditNotes/{CreditNoteID}/Attachments/{FileName} | Allows you to retrieve Attachments on CreditNote by file name
[**GetCreditNoteAttachmentById**](AccountingApi.md#GetCreditNoteAttachmentById) | **Get** /CreditNotes/{CreditNoteID}/Attachments/{AttachmentID} | Allows you to retrieve Attachments on CreditNote
[**GetCreditNoteAttachments**](AccountingApi.md#GetCreditNoteAttachments) | **Get** /CreditNotes/{CreditNoteID}/Attachments | Allows you to retrieve Attachments for credit notes
[**GetCreditNoteHistory**](AccountingApi.md#GetCreditNoteHistory) | **Get** /CreditNotes/{CreditNoteID}/History | Allows you to retrieve a history records of an CreditNote
[**GetCreditNotes**](AccountingApi.md#GetCreditNotes) | **Get** /CreditNotes | Allows you to retrieve any credit notes
[**GetCurrencies**](AccountingApi.md#GetCurrencies) | **Get** /Currencies | Allows you to retrieve currencies for your organisation
[**GetEmployee**](AccountingApi.md#GetEmployee) | **Get** /Employees/{EmployeeID} | Allows you to retrieve a specific employee used in Xero payrun
[**GetEmployees**](AccountingApi.md#GetEmployees) | **Get** /Employees | Allows you to retrieve employees used in Xero payrun
[**GetExpenseClaim**](AccountingApi.md#GetExpenseClaim) | **Get** /ExpenseClaims/{ExpenseClaimID} | Allows you to retrieve a specified expense claim
[**GetExpenseClaimHistory**](AccountingApi.md#GetExpenseClaimHistory) | **Get** /ExpenseClaims/{ExpenseClaimID}/History | Allows you to retrieve a history records of an ExpenseClaim
[**GetExpenseClaims**](AccountingApi.md#GetExpenseClaims) | **Get** /ExpenseClaims | Allows you to retrieve expense claims
[**GetInvoice**](AccountingApi.md#GetInvoice) | **Get** /Invoices/{InvoiceID} | Allows you to retrieve a specified sales invoice or purchase bill
[**GetInvoiceAsPdf**](AccountingApi.md#GetInvoiceAsPdf) | **Get** /Invoices/{InvoiceID}/pdf | Allows you to retrieve invoices or purchase bills as PDF files
[**GetInvoiceAttachmentByFileName**](AccountingApi.md#GetInvoiceAttachmentByFileName) | **Get** /Invoices/{InvoiceID}/Attachments/{FileName} | Allows you to retrieve Attachment on invoices or purchase bills by it&#39;s filename
[**GetInvoiceAttachmentById**](AccountingApi.md#GetInvoiceAttachmentById) | **Get** /Invoices/{InvoiceID}/Attachments/{AttachmentID} | Allows you to retrieve a specified Attachment on invoices or purchase bills by it&#39;s ID
[**GetInvoiceAttachments**](AccountingApi.md#GetInvoiceAttachments) | **Get** /Invoices/{InvoiceID}/Attachments | Allows you to retrieve Attachments on invoices or purchase bills
[**GetInvoiceHistory**](AccountingApi.md#GetInvoiceHistory) | **Get** /Invoices/{InvoiceID}/History | Allows you to retrieve a history records of an invoice
[**GetInvoiceReminders**](AccountingApi.md#GetInvoiceReminders) | **Get** /InvoiceReminders/Settings | Allows you to retrieve invoice reminder settings
[**GetInvoices**](AccountingApi.md#GetInvoices) | **Get** /Invoices | Allows you to retrieve any sales invoices or purchase bills
[**GetItem**](AccountingApi.md#GetItem) | **Get** /Items/{ItemID} | Allows you to retrieve a specified item
[**GetItemHistory**](AccountingApi.md#GetItemHistory) | **Get** /Items/{ItemID}/History | Allows you to retrieve history for items
[**GetItems**](AccountingApi.md#GetItems) | **Get** /Items | Allows you to retrieve any items
[**GetJournal**](AccountingApi.md#GetJournal) | **Get** /Journals/{JournalID} | Allows you to retrieve a specified journals.
[**GetJournals**](AccountingApi.md#GetJournals) | **Get** /Journals | Allows you to retrieve any journals.
[**GetLinkedTransaction**](AccountingApi.md#GetLinkedTransaction) | **Get** /LinkedTransactions/{LinkedTransactionID} | Allows you to retrieve a specified linked transactions (billable expenses)
[**GetLinkedTransactions**](AccountingApi.md#GetLinkedTransactions) | **Get** /LinkedTransactions | Retrieve linked transactions (billable expenses)
[**GetManualJournal**](AccountingApi.md#GetManualJournal) | **Get** /ManualJournals/{ManualJournalID} | Allows you to retrieve a specified manual journals
[**GetManualJournalAttachmentByFileName**](AccountingApi.md#GetManualJournalAttachmentByFileName) | **Get** /ManualJournals/{ManualJournalID}/Attachments/{FileName} | Allows you to retrieve specified Attachment on ManualJournal by file name
[**GetManualJournalAttachmentById**](AccountingApi.md#GetManualJournalAttachmentById) | **Get** /ManualJournals/{ManualJournalID}/Attachments/{AttachmentID} | Allows you to retrieve specified Attachment on ManualJournals
[**GetManualJournalAttachments**](AccountingApi.md#GetManualJournalAttachments) | **Get** /ManualJournals/{ManualJournalID}/Attachments | Allows you to retrieve Attachment for manual journals
[**GetManualJournals**](AccountingApi.md#GetManualJournals) | **Get** /ManualJournals | Allows you to retrieve any manual journals
[**GetOnlineInvoice**](AccountingApi.md#GetOnlineInvoice) | **Get** /Invoices/{InvoiceID}/OnlineInvoice | Allows you to retrieve a URL to an online invoice
[**GetOrganisationCISSettings**](AccountingApi.md#GetOrganisationCISSettings) | **Get** /Organisation/{OrganisationID}/CISSettings | Allows you To verify if an organisation is using contruction industry scheme, you can retrieve the CIS settings for the organistaion.
[**GetOrganisations**](AccountingApi.md#GetOrganisations) | **Get** /Organisation | Allows you to retrieve Organisation details
[**GetOverpayment**](AccountingApi.md#GetOverpayment) | **Get** /Overpayments/{OverpaymentID} | Allows you to retrieve a specified overpayments
[**GetOverpaymentHistory**](AccountingApi.md#GetOverpaymentHistory) | **Get** /Overpayments/{OverpaymentID}/History | Allows you to retrieve a history records of an Overpayment
[**GetOverpayments**](AccountingApi.md#GetOverpayments) | **Get** /Overpayments | Allows you to retrieve overpayments
[**GetPayment**](AccountingApi.md#GetPayment) | **Get** /Payments/{PaymentID} | Allows you to retrieve a specified payment for invoices and credit notes
[**GetPaymentHistory**](AccountingApi.md#GetPaymentHistory) | **Get** /Payments/{PaymentID}/History | Allows you to retrieve history records of a payment
[**GetPaymentServices**](AccountingApi.md#GetPaymentServices) | **Get** /PaymentServices | Allows you to retrieve payment services
[**GetPayments**](AccountingApi.md#GetPayments) | **Get** /Payments | Allows you to retrieve payments for invoices and credit notes
[**GetPrepayment**](AccountingApi.md#GetPrepayment) | **Get** /Prepayments/{PrepaymentID} | Allows you to retrieve a specified prepayments
[**GetPrepaymentHistory**](AccountingApi.md#GetPrepaymentHistory) | **Get** /Prepayments/{PrepaymentID}/History | Allows you to retrieve a history records of an Prepayment
[**GetPrepayments**](AccountingApi.md#GetPrepayments) | **Get** /Prepayments | Allows you to retrieve prepayments
[**GetPurchaseOrder**](AccountingApi.md#GetPurchaseOrder) | **Get** /PurchaseOrders/{PurchaseOrderID} | Allows you to retrieve a specified purchase orders
[**GetPurchaseOrderAsPdf**](AccountingApi.md#GetPurchaseOrderAsPdf) | **Get** /PurchaseOrders/{PurchaseOrderID}/pdf | Allows you to retrieve purchase orders as PDF files
[**GetPurchaseOrderByNumber**](AccountingApi.md#GetPurchaseOrderByNumber) | **Get** /PurchaseOrders/{PurchaseOrderNumber} | Allows you to retrieve a specified purchase orders
[**GetPurchaseOrderHistory**](AccountingApi.md#GetPurchaseOrderHistory) | **Get** /PurchaseOrders/{PurchaseOrderID}/History | Allows you to retrieve history for PurchaseOrder
[**GetPurchaseOrders**](AccountingApi.md#GetPurchaseOrders) | **Get** /PurchaseOrders | Allows you to retrieve purchase orders
[**GetQuote**](AccountingApi.md#GetQuote) | **Get** /Quotes/{QuoteID} | Allows you to retrieve a specified quote
[**GetQuoteAsPdf**](AccountingApi.md#GetQuoteAsPdf) | **Get** /Quotes/{QuoteID}/pdf | Allows you to retrieve quotes as PDF files
[**GetQuoteAttachmentByFileName**](AccountingApi.md#GetQuoteAttachmentByFileName) | **Get** /Quotes/{QuoteID}/Attachments/{FileName} | Allows you to retrieve Attachment on Quote by Filename
[**GetQuoteAttachmentById**](AccountingApi.md#GetQuoteAttachmentById) | **Get** /Quotes/{QuoteID}/Attachments/{AttachmentID} | Allows you to retrieve specific Attachment on Quote
[**GetQuoteAttachments**](AccountingApi.md#GetQuoteAttachments) | **Get** /Quotes/{QuoteID}/Attachments | Allows you to retrieve Attachments for Quotes
[**GetQuoteHistory**](AccountingApi.md#GetQuoteHistory) | **Get** /Quotes/{QuoteID}/History | Allows you to retrieve a history records of an quote
[**GetQuotes**](AccountingApi.md#GetQuotes) | **Get** /Quotes | Allows you to retrieve any sales quotes
[**GetReceipt**](AccountingApi.md#GetReceipt) | **Get** /Receipts/{ReceiptID} | Allows you to retrieve a specified draft expense claim receipts
[**GetReceiptAttachmentByFileName**](AccountingApi.md#GetReceiptAttachmentByFileName) | **Get** /Receipts/{ReceiptID}/Attachments/{FileName} | Allows you to retrieve Attachments on expense claim receipts by file name
[**GetReceiptAttachmentById**](AccountingApi.md#GetReceiptAttachmentById) | **Get** /Receipts/{ReceiptID}/Attachments/{AttachmentID} | Allows you to retrieve Attachments on expense claim receipts by ID
[**GetReceiptAttachments**](AccountingApi.md#GetReceiptAttachments) | **Get** /Receipts/{ReceiptID}/Attachments | Allows you to retrieve Attachments for expense claim receipts
[**GetReceiptHistory**](AccountingApi.md#GetReceiptHistory) | **Get** /Receipts/{ReceiptID}/History | Allows you to retrieve a history records of an Receipt
[**GetReceipts**](AccountingApi.md#GetReceipts) | **Get** /Receipts | Allows you to retrieve draft expense claim receipts for any user
[**GetRepeatingInvoice**](AccountingApi.md#GetRepeatingInvoice) | **Get** /RepeatingInvoices/{RepeatingInvoiceID} | Allows you to retrieve a specified repeating invoice
[**GetRepeatingInvoiceAttachmentByFileName**](AccountingApi.md#GetRepeatingInvoiceAttachmentByFileName) | **Get** /RepeatingInvoices/{RepeatingInvoiceID}/Attachments/{FileName} | Allows you to retrieve specified attachment on repeating invoices by file name
[**GetRepeatingInvoiceAttachmentById**](AccountingApi.md#GetRepeatingInvoiceAttachmentById) | **Get** /RepeatingInvoices/{RepeatingInvoiceID}/Attachments/{AttachmentID} | Allows you to retrieve a specified Attachments on repeating invoices
[**GetRepeatingInvoiceAttachments**](AccountingApi.md#GetRepeatingInvoiceAttachments) | **Get** /RepeatingInvoices/{RepeatingInvoiceID}/Attachments | Allows you to retrieve Attachments on repeating invoice
[**GetRepeatingInvoiceHistory**](AccountingApi.md#GetRepeatingInvoiceHistory) | **Get** /RepeatingInvoices/{RepeatingInvoiceID}/History | Allows you to retrieve history for a repeating invoice
[**GetRepeatingInvoices**](AccountingApi.md#GetRepeatingInvoices) | **Get** /RepeatingInvoices | Allows you to retrieve any repeating invoices
[**GetReportAgedPayablesByContact**](AccountingApi.md#GetReportAgedPayablesByContact) | **Get** /Reports/AgedPayablesByContact | Allows you to retrieve report for AgedPayablesByContact
[**GetReportAgedReceivablesByContact**](AccountingApi.md#GetReportAgedReceivablesByContact) | **Get** /Reports/AgedReceivablesByContact | Allows you to retrieve report for AgedReceivablesByContact
[**GetReportBASorGST**](AccountingApi.md#GetReportBASorGST) | **Get** /Reports/{ReportID} | Allows you to retrieve report for BAS only valid for AU orgs
[**GetReportBASorGSTList**](AccountingApi.md#GetReportBASorGSTList) | **Get** /Reports | Allows you to retrieve report for BAS only valid for AU orgs
[**GetReportBalanceSheet**](AccountingApi.md#GetReportBalanceSheet) | **Get** /Reports/BalanceSheet | Allows you to retrieve report for BalanceSheet
[**GetReportBankSummary**](AccountingApi.md#GetReportBankSummary) | **Get** /Reports/BankSummary | Allows you to retrieve report for BankSummary
[**GetReportBudgetSummary**](AccountingApi.md#GetReportBudgetSummary) | **Get** /Reports/BudgetSummary | Allows you to retrieve report for Budget Summary
[**GetReportExecutiveSummary**](AccountingApi.md#GetReportExecutiveSummary) | **Get** /Reports/ExecutiveSummary | Allows you to retrieve report for ExecutiveSummary
[**GetReportProfitAndLoss**](AccountingApi.md#GetReportProfitAndLoss) | **Get** /Reports/ProfitAndLoss | Allows you to retrieve report for ProfitAndLoss
[**GetReportTenNinetyNine**](AccountingApi.md#GetReportTenNinetyNine) | **Get** /Reports/TenNinetyNine | Allows you to retrieve report for TenNinetyNine
[**GetReportTrialBalance**](AccountingApi.md#GetReportTrialBalance) | **Get** /Reports/TrialBalance | Allows you to retrieve report for TrialBalance
[**GetTaxRates**](AccountingApi.md#GetTaxRates) | **Get** /TaxRates | Allows you to retrieve Tax Rates
[**GetTrackingCategories**](AccountingApi.md#GetTrackingCategories) | **Get** /TrackingCategories | Allows you to retrieve tracking categories and options
[**GetTrackingCategory**](AccountingApi.md#GetTrackingCategory) | **Get** /TrackingCategories/{TrackingCategoryID} | Allows you to retrieve tracking categories and options for specified category
[**GetUser**](AccountingApi.md#GetUser) | **Get** /Users/{UserID} | Allows you to retrieve a specified user
[**GetUsers**](AccountingApi.md#GetUsers) | **Get** /Users | Allows you to retrieve users
[**UpdateAccount**](AccountingApi.md#UpdateAccount) | **Post** /Accounts/{AccountID} | Allows you to update a chart of accounts
[**UpdateAccountAttachmentByFileName**](AccountingApi.md#UpdateAccountAttachmentByFileName) | **Post** /Accounts/{AccountID}/Attachments/{FileName} | Allows you to update Attachment on Account by Filename
[**UpdateBankTransaction**](AccountingApi.md#UpdateBankTransaction) | **Post** /BankTransactions/{BankTransactionID} | Allows you to update a single spend or receive money transaction
[**UpdateBankTransactionAttachmentByFileName**](AccountingApi.md#UpdateBankTransactionAttachmentByFileName) | **Post** /BankTransactions/{BankTransactionID}/Attachments/{FileName} | Allows you to update an Attachment on BankTransaction by Filename
[**UpdateBankTransferAttachmentByFileName**](AccountingApi.md#UpdateBankTransferAttachmentByFileName) | **Post** /BankTransfers/{BankTransferID}/Attachments/{FileName} | 
[**UpdateContact**](AccountingApi.md#UpdateContact) | **Post** /Contacts/{ContactID} | 
[**UpdateContactAttachmentByFileName**](AccountingApi.md#UpdateContactAttachmentByFileName) | **Post** /Contacts/{ContactID}/Attachments/{FileName} | 
[**UpdateContactGroup**](AccountingApi.md#UpdateContactGroup) | **Post** /ContactGroups/{ContactGroupID} | Allows you to update a Contact Group
[**UpdateCreditNote**](AccountingApi.md#UpdateCreditNote) | **Post** /CreditNotes/{CreditNoteID} | Allows you to update a specific credit note
[**UpdateCreditNoteAttachmentByFileName**](AccountingApi.md#UpdateCreditNoteAttachmentByFileName) | **Post** /CreditNotes/{CreditNoteID}/Attachments/{FileName} | Allows you to update Attachments on CreditNote by file name
[**UpdateExpenseClaim**](AccountingApi.md#UpdateExpenseClaim) | **Post** /ExpenseClaims/{ExpenseClaimID} | Allows you to update specified expense claims
[**UpdateInvoice**](AccountingApi.md#UpdateInvoice) | **Post** /Invoices/{InvoiceID} | Allows you to update a specified sales invoices or purchase bills
[**UpdateInvoiceAttachmentByFileName**](AccountingApi.md#UpdateInvoiceAttachmentByFileName) | **Post** /Invoices/{InvoiceID}/Attachments/{FileName} | Allows you to update Attachment on invoices or purchase bills by it&#39;s filename
[**UpdateItem**](AccountingApi.md#UpdateItem) | **Post** /Items/{ItemID} | Allows you to update a specified item
[**UpdateLinkedTransaction**](AccountingApi.md#UpdateLinkedTransaction) | **Post** /LinkedTransactions/{LinkedTransactionID} | Allows you to update a specified linked transactions (billable expenses)
[**UpdateManualJournal**](AccountingApi.md#UpdateManualJournal) | **Post** /ManualJournals/{ManualJournalID} | Allows you to update a specified manual journal
[**UpdateManualJournalAttachmentByFileName**](AccountingApi.md#UpdateManualJournalAttachmentByFileName) | **Post** /ManualJournals/{ManualJournalID}/Attachments/{FileName} | Allows you to update a specified Attachment on ManualJournal by file name
[**UpdateOrCreateBankTransactions**](AccountingApi.md#UpdateOrCreateBankTransactions) | **Post** /BankTransactions | Allows you to update or create one or more spend or receive money transaction
[**UpdateOrCreateContacts**](AccountingApi.md#UpdateOrCreateContacts) | **Post** /Contacts | Allows you to update OR create one or more contacts in a Xero organisation
[**UpdateOrCreateCreditNotes**](AccountingApi.md#UpdateOrCreateCreditNotes) | **Post** /CreditNotes | Allows you to update OR create one or more credit notes
[**UpdateOrCreateEmployees**](AccountingApi.md#UpdateOrCreateEmployees) | **Post** /Employees | Allows you to create a single new employees used in Xero payrun
[**UpdateOrCreateInvoices**](AccountingApi.md#UpdateOrCreateInvoices) | **Post** /Invoices | Allows you to update OR create one or more sales invoices or purchase bills
[**UpdateOrCreateItems**](AccountingApi.md#UpdateOrCreateItems) | **Post** /Items | Allows you to update or create one or more items
[**UpdateOrCreateManualJournals**](AccountingApi.md#UpdateOrCreateManualJournals) | **Post** /ManualJournals | Allows you to create a single manual journal
[**UpdateOrCreatePurchaseOrders**](AccountingApi.md#UpdateOrCreatePurchaseOrders) | **Post** /PurchaseOrders | Allows you to update or create one or more purchase orders
[**UpdateOrCreateQuotes**](AccountingApi.md#UpdateOrCreateQuotes) | **Post** /Quotes | Allows you to update OR create one or more quotes
[**UpdatePurchaseOrder**](AccountingApi.md#UpdatePurchaseOrder) | **Post** /PurchaseOrders/{PurchaseOrderID} | Allows you to update a specified purchase order
[**UpdateQuote**](AccountingApi.md#UpdateQuote) | **Post** /Quotes/{QuoteID} | Allows you to update a specified quote
[**UpdateQuoteAttachmentByFileName**](AccountingApi.md#UpdateQuoteAttachmentByFileName) | **Post** /Quotes/{QuoteID}/Attachments/{FileName} | Allows you to update Attachment on Quote by Filename
[**UpdateReceipt**](AccountingApi.md#UpdateReceipt) | **Post** /Receipts/{ReceiptID} | Allows you to retrieve a specified draft expense claim receipts
[**UpdateReceiptAttachmentByFileName**](AccountingApi.md#UpdateReceiptAttachmentByFileName) | **Post** /Receipts/{ReceiptID}/Attachments/{FileName} | Allows you to update Attachment on expense claim receipts by file name
[**UpdateRepeatingInvoiceAttachmentByFileName**](AccountingApi.md#UpdateRepeatingInvoiceAttachmentByFileName) | **Post** /RepeatingInvoices/{RepeatingInvoiceID}/Attachments/{FileName} | Allows you to update specified attachment on repeating invoices by file name
[**UpdateTaxRate**](AccountingApi.md#UpdateTaxRate) | **Post** /TaxRates | Allows you to update Tax Rates
[**UpdateTrackingCategory**](AccountingApi.md#UpdateTrackingCategory) | **Post** /TrackingCategories/{TrackingCategoryID} | Allows you to update tracking categories
[**UpdateTrackingOptions**](AccountingApi.md#UpdateTrackingOptions) | **Post** /TrackingCategories/{TrackingCategoryID}/Options/{TrackingOptionID} | Allows you to update options for a specified tracking category



## CreateAccount

> Accounts CreateAccount(ctx, xeroTenantId, account)

Allows you to create a new chart of accounts

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**account** | [**Account**](Account.md)| Account object in body of request | 

### Return type

[**Accounts**](Accounts.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAccountAttachmentByFileName

> Attachments CreateAccountAttachmentByFileName(ctx, xeroTenantId, accountID, fileName, body)

Allows you to create Attachment on Account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**accountID** | [**string**](.md)| Unique identifier for Account object | 
**fileName** | **string**| Name of the attachment | 
**body** | **string**| Byte array of file in body of request | 

### Return type

[**Attachments**](Attachments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBankTransactionAttachmentByFileName

> Attachments CreateBankTransactionAttachmentByFileName(ctx, xeroTenantId, bankTransactionID, fileName, body)

Allows you to createa an Attachment on BankTransaction by Filename

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**bankTransactionID** | [**string**](.md)| Xero generated unique identifier for a bank transaction | 
**fileName** | **string**| The name of the file being attached | 
**body** | **string**| Byte array of file in body of request | 

### Return type

[**Attachments**](Attachments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBankTransactionHistoryRecord

> HistoryRecords CreateBankTransactionHistoryRecord(ctx, xeroTenantId, bankTransactionID, historyRecords)

Allows you to create history record for a bank transactions

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**bankTransactionID** | [**string**](.md)| Xero generated unique identifier for a bank transaction | 
**historyRecords** | [**HistoryRecords**](HistoryRecords.md)| HistoryRecords containing an array of HistoryRecord objects in body of request | 

### Return type

[**HistoryRecords**](HistoryRecords.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBankTransactions

> BankTransactions CreateBankTransactions(ctx, xeroTenantId, bankTransactions, optional)

Allows you to create one or more spend or receive money transaction

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**bankTransactions** | [**BankTransactions**](BankTransactions.md)| BankTransactions with an array of BankTransaction objects in body of request | 
 **optional** | ***CreateBankTransactionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateBankTransactionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **summarizeErrors** | **optional.Bool**| If false return 200 OK and mix of successfully created obejcts and any with validation errors | [default to false]
 **unitdp** | **optional.Int32**| e.g. unitdp&#x3D;4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts | 

### Return type

[**BankTransactions**](BankTransactions.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBankTransfer

> BankTransfers CreateBankTransfer(ctx, xeroTenantId, bankTransfers)

Allows you to create a bank transfers

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**bankTransfers** | [**BankTransfers**](BankTransfers.md)| BankTransfers with array of BankTransfer objects in request body | 

### Return type

[**BankTransfers**](BankTransfers.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBankTransferAttachmentByFileName

> Attachments CreateBankTransferAttachmentByFileName(ctx, xeroTenantId, bankTransferID, fileName, body)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**bankTransferID** | [**string**](.md)| Xero generated unique identifier for a bank transfer | 
**fileName** | **string**| The name of the file being attached to a Bank Transfer | 
**body** | **string**| Byte array of file in body of request | 

### Return type

[**Attachments**](Attachments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBatchPayment

> BatchPayments CreateBatchPayment(ctx, xeroTenantId, batchPayments, optional)

Create one or many BatchPayments for invoices

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**batchPayments** | [**BatchPayments**](BatchPayments.md)| BatchPayments with an array of Payments in body of request | 
 **optional** | ***CreateBatchPaymentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateBatchPaymentOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **summarizeErrors** | **optional.Bool**| If false return 200 OK and mix of successfully created obejcts and any with validation errors | [default to false]

### Return type

[**BatchPayments**](BatchPayments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBatchPaymentHistoryRecord

> HistoryRecords CreateBatchPaymentHistoryRecord(ctx, xeroTenantId, batchPaymentID, historyRecords)

Allows you to create a history record for a Batch Payment

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**batchPaymentID** | [**string**](.md)| Unique identifier for BatchPayment | 
**historyRecords** | [**HistoryRecords**](HistoryRecords.md)| HistoryRecords containing an array of HistoryRecord objects in body of request | 

### Return type

[**HistoryRecords**](HistoryRecords.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBrandingThemePaymentServices

> PaymentServices CreateBrandingThemePaymentServices(ctx, xeroTenantId, brandingThemeID, paymentService)

Allow for the creation of new custom payment service for specified Branding Theme

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**brandingThemeID** | [**string**](.md)| Unique identifier for a Branding Theme | 
**paymentService** | [**PaymentService**](PaymentService.md)| PaymentService object in body of request | 

### Return type

[**PaymentServices**](PaymentServices.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateContactAttachmentByFileName

> Attachments CreateContactAttachmentByFileName(ctx, xeroTenantId, contactID, fileName, body)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**contactID** | [**string**](.md)| Unique identifier for a Contact | 
**fileName** | **string**| Name for the file you are attaching | 
**body** | **string**| Byte array of file in body of request | 

### Return type

[**Attachments**](Attachments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateContactGroup

> ContactGroups CreateContactGroup(ctx, xeroTenantId, contactGroups)

Allows you to create a contact group

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**contactGroups** | [**ContactGroups**](ContactGroups.md)| ContactGroups with an array of names in request body | 

### Return type

[**ContactGroups**](ContactGroups.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateContactGroupContacts

> Contacts CreateContactGroupContacts(ctx, xeroTenantId, contactGroupID, contacts)

Allows you to add Contacts to a Contact Group

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**contactGroupID** | [**string**](.md)| Unique identifier for a Contact Group | 
**contacts** | [**Contacts**](Contacts.md)| Contacts with array of contacts specifiying the ContactID to be added to ContactGroup in body of request | 

### Return type

[**Contacts**](Contacts.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateContacts

> Contacts CreateContacts(ctx, xeroTenantId, contacts, optional)

Allows you to create a multiple contacts (bulk) in a Xero organisation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**contacts** | [**Contacts**](Contacts.md)| Contacts with an array of Contact objects to create in body of request | 
 **optional** | ***CreateContactsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateContactsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **summarizeErrors** | **optional.Bool**| If false return 200 OK and mix of successfully created obejcts and any with validation errors | [default to false]

### Return type

[**Contacts**](Contacts.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCreditNoteAllocation

> Allocations CreateCreditNoteAllocation(ctx, xeroTenantId, creditNoteID, allocations, optional)

Allows you to create Allocation on CreditNote

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**creditNoteID** | [**string**](.md)| Unique identifier for a Credit Note | 
**allocations** | [**Allocations**](Allocations.md)| Allocations with array of Allocation object in body of request. | 
 **optional** | ***CreateCreditNoteAllocationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateCreditNoteAllocationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **summarizeErrors** | **optional.Bool**| If false return 200 OK and mix of successfully created obejcts and any with validation errors | [default to false]

### Return type

[**Allocations**](Allocations.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCreditNoteAttachmentByFileName

> Attachments CreateCreditNoteAttachmentByFileName(ctx, xeroTenantId, creditNoteID, fileName, body, optional)

Allows you to create Attachments on CreditNote by file name

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**creditNoteID** | [**string**](.md)| Unique identifier for a Credit Note | 
**fileName** | **string**| Name of the file you are attaching to Credit Note | 
**body** | **string**| Byte array of file in body of request | 
 **optional** | ***CreateCreditNoteAttachmentByFileNameOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateCreditNoteAttachmentByFileNameOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **includeOnline** | **optional.Bool**| Allows an attachment to be seen by the end customer within their online invoice | [default to false]

### Return type

[**Attachments**](Attachments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCreditNoteHistory

> HistoryRecords CreateCreditNoteHistory(ctx, xeroTenantId, creditNoteID, historyRecords)

Allows you to retrieve a history records of an CreditNote

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**creditNoteID** | [**string**](.md)| Unique identifier for a Credit Note | 
**historyRecords** | [**HistoryRecords**](HistoryRecords.md)| HistoryRecords containing an array of HistoryRecord objects in body of request | 

### Return type

[**HistoryRecords**](HistoryRecords.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCreditNotes

> CreditNotes CreateCreditNotes(ctx, xeroTenantId, creditNotes, optional)

Allows you to create a credit note

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**creditNotes** | [**CreditNotes**](CreditNotes.md)| Credit Notes with array of CreditNote object in body of request | 
 **optional** | ***CreateCreditNotesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateCreditNotesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **summarizeErrors** | **optional.Bool**| If false return 200 OK and mix of successfully created obejcts and any with validation errors | [default to false]
 **unitdp** | **optional.Int32**| e.g. unitdp&#x3D;4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts | 

### Return type

[**CreditNotes**](CreditNotes.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCurrency

> Currencies CreateCurrency(ctx, xeroTenantId, currency)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**currency** | [**Currency**](Currency.md)| Currency obejct in the body of request | 

### Return type

[**Currencies**](Currencies.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEmployees

> Employees CreateEmployees(ctx, xeroTenantId, employees, optional)

Allows you to create new employees used in Xero payrun

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**employees** | [**Employees**](Employees.md)| Employees with array of Employee object in body of request | 
 **optional** | ***CreateEmployeesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateEmployeesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **summarizeErrors** | **optional.Bool**| If false return 200 OK and mix of successfully created obejcts and any with validation errors | [default to false]

### Return type

[**Employees**](Employees.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateExpenseClaims

> ExpenseClaims CreateExpenseClaims(ctx, xeroTenantId, expenseClaims)

Allows you to retrieve expense claims

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**expenseClaims** | [**ExpenseClaims**](ExpenseClaims.md)| ExpenseClaims with array of ExpenseClaim object in body of request | 

### Return type

[**ExpenseClaims**](ExpenseClaims.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInvoiceAttachmentByFileName

> Attachments CreateInvoiceAttachmentByFileName(ctx, xeroTenantId, invoiceID, fileName, body, optional)

Allows you to create an Attachment on invoices or purchase bills by it's filename

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**invoiceID** | [**string**](.md)| Unique identifier for an Invoice | 
**fileName** | **string**| Name of the file you are attaching | 
**body** | **string**| Byte array of file in body of request | 
 **optional** | ***CreateInvoiceAttachmentByFileNameOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateInvoiceAttachmentByFileNameOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **includeOnline** | **optional.Bool**| Allows an attachment to be seen by the end customer within their online invoice | [default to false]

### Return type

[**Attachments**](Attachments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInvoiceHistory

> HistoryRecords CreateInvoiceHistory(ctx, xeroTenantId, invoiceID, historyRecords)

Allows you to retrieve a history records of an invoice

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**invoiceID** | [**string**](.md)| Unique identifier for an Invoice | 
**historyRecords** | [**HistoryRecords**](HistoryRecords.md)| HistoryRecords containing an array of HistoryRecord objects in body of request | 

### Return type

[**HistoryRecords**](HistoryRecords.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInvoices

> Invoices CreateInvoices(ctx, xeroTenantId, invoices, optional)

Allows you to create one or more sales invoices or purchase bills

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**invoices** | [**Invoices**](Invoices.md)| Invoices with an array of invoice objects in body of request | 
 **optional** | ***CreateInvoicesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateInvoicesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **summarizeErrors** | **optional.Bool**| If false return 200 OK and mix of successfully created obejcts and any with validation errors | [default to false]
 **unitdp** | **optional.Int32**| e.g. unitdp&#x3D;4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts | 

### Return type

[**Invoices**](Invoices.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateItems

> Items CreateItems(ctx, xeroTenantId, items, optional)

Allows you to create one or more items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**items** | [**Items**](Items.md)| Items with an array of Item objects in body of request | 
 **optional** | ***CreateItemsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateItemsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **summarizeErrors** | **optional.Bool**| If false return 200 OK and mix of successfully created obejcts and any with validation errors | [default to false]
 **unitdp** | **optional.Int32**| e.g. unitdp&#x3D;4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts | 

### Return type

[**Items**](Items.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLinkedTransaction

> LinkedTransactions CreateLinkedTransaction(ctx, xeroTenantId, linkedTransaction)

Allows you to create linked transactions (billable expenses)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**linkedTransaction** | [**LinkedTransaction**](LinkedTransaction.md)| LinkedTransaction object in body of request | 

### Return type

[**LinkedTransactions**](LinkedTransactions.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateManualJournalAttachmentByFileName

> Attachments CreateManualJournalAttachmentByFileName(ctx, xeroTenantId, manualJournalID, fileName, body)

Allows you to create a specified Attachment on ManualJournal by file name

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**manualJournalID** | [**string**](.md)| Unique identifier for a ManualJournal | 
**fileName** | **string**| The name of the file being attached to a ManualJournal | 
**body** | **string**| Byte array of file in body of request | 

### Return type

[**Attachments**](Attachments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateManualJournals

> ManualJournals CreateManualJournals(ctx, xeroTenantId, manualJournals, optional)

Allows you to create one or more manual journals

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**manualJournals** | [**ManualJournals**](ManualJournals.md)| ManualJournals array with ManualJournal object in body of request | 
 **optional** | ***CreateManualJournalsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateManualJournalsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **summarizeErrors** | **optional.Bool**| If false return 200 OK and mix of successfully created obejcts and any with validation errors | [default to false]

### Return type

[**ManualJournals**](ManualJournals.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOverpaymentAllocations

> Allocations CreateOverpaymentAllocations(ctx, xeroTenantId, overpaymentID, allocations, optional)

Allows you to create a single allocation for an overpayment

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**overpaymentID** | [**string**](.md)| Unique identifier for a Overpayment | 
**allocations** | [**Allocations**](Allocations.md)| Allocations array with Allocation object in body of request | 
 **optional** | ***CreateOverpaymentAllocationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateOverpaymentAllocationsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **summarizeErrors** | **optional.Bool**| If false return 200 OK and mix of successfully created obejcts and any with validation errors | [default to false]

### Return type

[**Allocations**](Allocations.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePayment

> Payments CreatePayment(ctx, xeroTenantId, payment)

Allows you to create a single payment for invoices or credit notes

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**payment** | [**Payment**](Payment.md)| Request body with a single Payment object | 

### Return type

[**Payments**](Payments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePaymentService

> PaymentServices CreatePaymentService(ctx, xeroTenantId, paymentServices)

Allows you to create payment services

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**paymentServices** | [**PaymentServices**](PaymentServices.md)| PaymentServices array with PaymentService object in body of request | 

### Return type

[**PaymentServices**](PaymentServices.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePayments

> Payments CreatePayments(ctx, xeroTenantId, payments, optional)

Allows you to create multiple payments for invoices or credit notes

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**payments** | [**Payments**](Payments.md)| Payments array with Payment object in body of request | 
 **optional** | ***CreatePaymentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreatePaymentsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **summarizeErrors** | **optional.Bool**| If false return 200 OK and mix of successfully created obejcts and any with validation errors | [default to false]

### Return type

[**Payments**](Payments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePrepaymentAllocations

> Allocations CreatePrepaymentAllocations(ctx, xeroTenantId, prepaymentID, allocations, optional)

Allows you to create an Allocation for prepayments

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**prepaymentID** | [**string**](.md)| Unique identifier for Prepayment | 
**allocations** | [**Allocations**](Allocations.md)| Allocations with an array of Allocation object in body of request | 
 **optional** | ***CreatePrepaymentAllocationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreatePrepaymentAllocationsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **summarizeErrors** | **optional.Bool**| If false return 200 OK and mix of successfully created obejcts and any with validation errors | [default to false]

### Return type

[**Allocations**](Allocations.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePurchaseOrderHistory

> HistoryRecords CreatePurchaseOrderHistory(ctx, xeroTenantId, purchaseOrderID, historyRecords)

Allows you to create HistoryRecord for purchase orders

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**purchaseOrderID** | [**string**](.md)| Unique identifier for a PurchaseOrder | 
**historyRecords** | [**HistoryRecords**](HistoryRecords.md)| HistoryRecords containing an array of HistoryRecord objects in body of request | 

### Return type

[**HistoryRecords**](HistoryRecords.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePurchaseOrders

> PurchaseOrders CreatePurchaseOrders(ctx, xeroTenantId, purchaseOrders, optional)

Allows you to create one or more purchase orders

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**purchaseOrders** | [**PurchaseOrders**](PurchaseOrders.md)| PurchaseOrders with an array of PurchaseOrder object in body of request | 
 **optional** | ***CreatePurchaseOrdersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreatePurchaseOrdersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **summarizeErrors** | **optional.Bool**| If false return 200 OK and mix of successfully created obejcts and any with validation errors | [default to false]

### Return type

[**PurchaseOrders**](PurchaseOrders.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateQuoteAttachmentByFileName

> Attachments CreateQuoteAttachmentByFileName(ctx, xeroTenantId, quoteID, fileName, body)

Allows you to create Attachment on Quote

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**quoteID** | [**string**](.md)| Unique identifier for Quote object | 
**fileName** | **string**| Name of the attachment | 
**body** | **string**| Byte array of file in body of request | 

### Return type

[**Attachments**](Attachments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateQuoteHistory

> HistoryRecords CreateQuoteHistory(ctx, xeroTenantId, quoteID, historyRecords)

Allows you to retrieve a history records of an quote

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**quoteID** | [**string**](.md)| Unique identifier for an Quote | 
**historyRecords** | [**HistoryRecords**](HistoryRecords.md)| HistoryRecords containing an array of HistoryRecord objects in body of request | 

### Return type

[**HistoryRecords**](HistoryRecords.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateQuotes

> Quotes CreateQuotes(ctx, xeroTenantId, quotes, optional)

Allows you to create one or more quotes

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**quotes** | [**Quotes**](Quotes.md)| Quotes with an array of Quote object in body of request | 
 **optional** | ***CreateQuotesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateQuotesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **summarizeErrors** | **optional.Bool**| If false return 200 OK and mix of successfully created obejcts and any with validation errors | [default to false]

### Return type

[**Quotes**](Quotes.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateReceipt

> Receipts CreateReceipt(ctx, xeroTenantId, receipts, optional)

Allows you to create draft expense claim receipts for any user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**receipts** | [**Receipts**](Receipts.md)| Receipts with an array of Receipt object in body of request | 
 **optional** | ***CreateReceiptOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateReceiptOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **unitdp** | **optional.Int32**| e.g. unitdp&#x3D;4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts | 

### Return type

[**Receipts**](Receipts.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateReceiptAttachmentByFileName

> Attachments CreateReceiptAttachmentByFileName(ctx, xeroTenantId, receiptID, fileName, body)

Allows you to create Attachment on expense claim receipts by file name

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**receiptID** | [**string**](.md)| Unique identifier for a Receipt | 
**fileName** | **string**| The name of the file being attached to the Receipt | 
**body** | **string**| Byte array of file in body of request | 

### Return type

[**Attachments**](Attachments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRepeatingInvoiceAttachmentByFileName

> Attachments CreateRepeatingInvoiceAttachmentByFileName(ctx, xeroTenantId, repeatingInvoiceID, fileName, body)

Allows you to create attachment on repeating invoices by file name

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**repeatingInvoiceID** | [**string**](.md)| Unique identifier for a Repeating Invoice | 
**fileName** | **string**| The name of the file being attached to a Repeating Invoice | 
**body** | **string**| Byte array of file in body of request | 

### Return type

[**Attachments**](Attachments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTaxRates

> TaxRates CreateTaxRates(ctx, xeroTenantId, taxRates)

Allows you to create one or more Tax Rates

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**taxRates** | [**TaxRates**](TaxRates.md)| TaxRates array with TaxRate object in body of request | 

### Return type

[**TaxRates**](TaxRates.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTrackingCategory

> TrackingCategories CreateTrackingCategory(ctx, xeroTenantId, trackingCategory)

Allows you to create tracking categories

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**trackingCategory** | [**TrackingCategory**](TrackingCategory.md)| TrackingCategory object in body of request | 

### Return type

[**TrackingCategories**](TrackingCategories.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTrackingOptions

> TrackingOptions CreateTrackingOptions(ctx, xeroTenantId, trackingCategoryID, trackingOption)

Allows you to create options for a specified tracking category

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**trackingCategoryID** | [**string**](.md)| Unique identifier for a TrackingCategory | 
**trackingOption** | [**TrackingOption**](TrackingOption.md)| TrackingOption object in body of request | 

### Return type

[**TrackingOptions**](TrackingOptions.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAccount

> Accounts DeleteAccount(ctx, xeroTenantId, accountID)

Allows you to delete a chart of accounts

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**accountID** | [**string**](.md)| Unique identifier for retrieving single object | 

### Return type

[**Accounts**](Accounts.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteContactGroupContact

> DeleteContactGroupContact(ctx, xeroTenantId, contactGroupID, contactID)

Allows you to delete a specific Contact from a Contact Group

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**contactGroupID** | [**string**](.md)| Unique identifier for a Contact Group | 
**contactID** | [**string**](.md)| Unique identifier for a Contact | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteContactGroupContacts

> DeleteContactGroupContacts(ctx, xeroTenantId, contactGroupID)

Allows you to delete  all Contacts from a Contact Group

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**contactGroupID** | [**string**](.md)| Unique identifier for a Contact Group | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteItem

> DeleteItem(ctx, xeroTenantId, itemID)

Allows you to delete a specified item

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**itemID** | [**string**](.md)| Unique identifier for an Item | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLinkedTransaction

> DeleteLinkedTransaction(ctx, xeroTenantId, linkedTransactionID)

Allows you to delete a specified linked transactions (billable expenses)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**linkedTransactionID** | [**string**](.md)| Unique identifier for a LinkedTransaction | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePayment

> Payments DeletePayment(ctx, xeroTenantId, paymentID, paymentDelete)

Allows you to update a specified payment for invoices and credit notes

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**paymentID** | [**string**](.md)| Unique identifier for a Payment | 
**paymentDelete** | [**PaymentDelete**](PaymentDelete.md)|  | 

### Return type

[**Payments**](Payments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTrackingCategory

> TrackingCategories DeleteTrackingCategory(ctx, xeroTenantId, trackingCategoryID)

Allows you to delete tracking categories

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**trackingCategoryID** | [**string**](.md)| Unique identifier for a TrackingCategory | 

### Return type

[**TrackingCategories**](TrackingCategories.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTrackingOptions

> TrackingOptions DeleteTrackingOptions(ctx, xeroTenantId, trackingCategoryID, trackingOptionID)

Allows you to delete a specified option for a specified tracking category

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**trackingCategoryID** | [**string**](.md)| Unique identifier for a TrackingCategory | 
**trackingOptionID** | [**string**](.md)| Unique identifier for a Tracking Option | 

### Return type

[**TrackingOptions**](TrackingOptions.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmailInvoice

> EmailInvoice(ctx, xeroTenantId, invoiceID, requestEmpty)

Allows you to email a copy of invoice to related Contact

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**invoiceID** | [**string**](.md)| Unique identifier for an Invoice | 
**requestEmpty** | [**RequestEmpty**](RequestEmpty.md)|  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccount

> Accounts GetAccount(ctx, xeroTenantId, accountID)

Allows you to retrieve a single chart of accounts

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**accountID** | [**string**](.md)| Unique identifier for retrieving single object | 

### Return type

[**Accounts**](Accounts.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountAttachmentByFileName

> *os.File GetAccountAttachmentByFileName(ctx, xeroTenantId, accountID, fileName, contentType)

Allows you to retrieve Attachment on Account by Filename

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**accountID** | [**string**](.md)| Unique identifier for Account object | 
**fileName** | **string**| Name of the attachment | 
**contentType** | **string**| The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountAttachmentById

> *os.File GetAccountAttachmentById(ctx, xeroTenantId, accountID, attachmentID, contentType)

Allows you to retrieve specific Attachment on Account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**accountID** | [**string**](.md)| Unique identifier for Account object | 
**attachmentID** | [**string**](.md)| Unique identifier for Attachment object | 
**contentType** | **string**| The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountAttachments

> Attachments GetAccountAttachments(ctx, xeroTenantId, accountID)

Allows you to retrieve Attachments for accounts

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**accountID** | [**string**](.md)| Unique identifier for Account object | 

### Return type

[**Attachments**](Attachments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccounts

> Accounts GetAccounts(ctx, xeroTenantId, optional)

Allows you to retrieve the full chart of accounts

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
 **optional** | ***GetAccountsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAccountsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifModifiedSince** | **optional.Time**| Only records created or modified since this timestamp will be returned | 
 **where** | **optional.String**| Filter by an any element | 
 **order** | **optional.String**| Order by an any element | 

### Return type

[**Accounts**](Accounts.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBankTransaction

> BankTransactions GetBankTransaction(ctx, xeroTenantId, bankTransactionID, optional)

Allows you to retrieve a single spend or receive money transaction

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**bankTransactionID** | [**string**](.md)| Xero generated unique identifier for a bank transaction | 
 **optional** | ***GetBankTransactionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetBankTransactionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **unitdp** | **optional.Int32**| e.g. unitdp&#x3D;4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts | 

### Return type

[**BankTransactions**](BankTransactions.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBankTransactionAttachmentByFileName

> *os.File GetBankTransactionAttachmentByFileName(ctx, xeroTenantId, bankTransactionID, fileName, contentType)

Allows you to retrieve Attachments on BankTransaction by Filename

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**bankTransactionID** | [**string**](.md)| Xero generated unique identifier for a bank transaction | 
**fileName** | **string**| The name of the file being attached | 
**contentType** | **string**| The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBankTransactionAttachmentById

> *os.File GetBankTransactionAttachmentById(ctx, xeroTenantId, bankTransactionID, attachmentID, contentType)

Allows you to retrieve Attachments on a specific BankTransaction

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**bankTransactionID** | [**string**](.md)| Xero generated unique identifier for a bank transaction | 
**attachmentID** | [**string**](.md)| Xero generated unique identifier for an attachment | 
**contentType** | **string**| The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBankTransactionAttachments

> Attachments GetBankTransactionAttachments(ctx, xeroTenantId, bankTransactionID)

Allows you to retrieve any attachments to bank transactions

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**bankTransactionID** | [**string**](.md)| Xero generated unique identifier for a bank transaction | 

### Return type

[**Attachments**](Attachments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBankTransactions

> BankTransactions GetBankTransactions(ctx, xeroTenantId, optional)

Allows you to retrieve any spend or receive money transactions

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
 **optional** | ***GetBankTransactionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetBankTransactionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifModifiedSince** | **optional.Time**| Only records created or modified since this timestamp will be returned | 
 **where** | **optional.String**| Filter by an any element | 
 **order** | **optional.String**| Order by an any element | 
 **page** | **optional.Int32**| Up to 100 bank transactions will be returned in a single API call with line items details | 
 **unitdp** | **optional.Int32**| e.g. unitdp&#x3D;4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts | 

### Return type

[**BankTransactions**](BankTransactions.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBankTransactionsHistory

> HistoryRecords GetBankTransactionsHistory(ctx, xeroTenantId, bankTransactionID)

Allows you to retrieve history from a bank transactions

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**bankTransactionID** | [**string**](.md)| Xero generated unique identifier for a bank transaction | 

### Return type

[**HistoryRecords**](HistoryRecords.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBankTransfer

> BankTransfers GetBankTransfer(ctx, xeroTenantId, bankTransferID)

Allows you to retrieve any bank transfers

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**bankTransferID** | [**string**](.md)| Xero generated unique identifier for a bank transfer | 

### Return type

[**BankTransfers**](BankTransfers.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBankTransferAttachmentByFileName

> *os.File GetBankTransferAttachmentByFileName(ctx, xeroTenantId, bankTransferID, fileName, contentType)

Allows you to retrieve Attachments on BankTransfer by file name

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**bankTransferID** | [**string**](.md)| Xero generated unique identifier for a bank transfer | 
**fileName** | **string**| The name of the file being attached to a Bank Transfer | 
**contentType** | **string**| The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBankTransferAttachmentById

> *os.File GetBankTransferAttachmentById(ctx, xeroTenantId, bankTransferID, attachmentID, contentType)

Allows you to retrieve Attachments on BankTransfer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**bankTransferID** | [**string**](.md)| Xero generated unique identifier for a bank transfer | 
**attachmentID** | [**string**](.md)| Xero generated unique identifier for an Attachment to a bank transfer | 
**contentType** | **string**| The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBankTransferAttachments

> Attachments GetBankTransferAttachments(ctx, xeroTenantId, bankTransferID)

Allows you to retrieve Attachments from  bank transfers

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**bankTransferID** | [**string**](.md)| Xero generated unique identifier for a bank transfer | 

### Return type

[**Attachments**](Attachments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBankTransferHistory

> HistoryRecords GetBankTransferHistory(ctx, xeroTenantId, bankTransferID)

Allows you to retrieve history from a bank transfers

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**bankTransferID** | [**string**](.md)| Xero generated unique identifier for a bank transfer | 

### Return type

[**HistoryRecords**](HistoryRecords.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBankTransfers

> BankTransfers GetBankTransfers(ctx, xeroTenantId, optional)

Allows you to retrieve all bank transfers

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
 **optional** | ***GetBankTransfersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetBankTransfersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifModifiedSince** | **optional.Time**| Only records created or modified since this timestamp will be returned | 
 **where** | **optional.String**| Filter by an any element | 
 **order** | **optional.String**| Order by an any element | 

### Return type

[**BankTransfers**](BankTransfers.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBatchPaymentHistory

> HistoryRecords GetBatchPaymentHistory(ctx, xeroTenantId, batchPaymentID)

Allows you to retrieve history from a Batch Payment

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**batchPaymentID** | [**string**](.md)| Unique identifier for BatchPayment | 

### Return type

[**HistoryRecords**](HistoryRecords.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBatchPayments

> BatchPayments GetBatchPayments(ctx, xeroTenantId, optional)

Retrieve either one or many BatchPayments for invoices

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
 **optional** | ***GetBatchPaymentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetBatchPaymentsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifModifiedSince** | **optional.Time**| Only records created or modified since this timestamp will be returned | 
 **where** | **optional.String**| Filter by an any element | 
 **order** | **optional.String**| Order by an any element | 

### Return type

[**BatchPayments**](BatchPayments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBrandingTheme

> BrandingThemes GetBrandingTheme(ctx, xeroTenantId, brandingThemeID)

Allows you to retrieve a specific BrandingThemes

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**brandingThemeID** | [**string**](.md)| Unique identifier for a Branding Theme | 

### Return type

[**BrandingThemes**](BrandingThemes.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBrandingThemePaymentServices

> PaymentServices GetBrandingThemePaymentServices(ctx, xeroTenantId, brandingThemeID)

Allows you to retrieve the Payment services for a Branding Theme

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**brandingThemeID** | [**string**](.md)| Unique identifier for a Branding Theme | 

### Return type

[**PaymentServices**](PaymentServices.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBrandingThemes

> BrandingThemes GetBrandingThemes(ctx, xeroTenantId)

Allows you to retrieve all the BrandingThemes

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 

### Return type

[**BrandingThemes**](BrandingThemes.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContact

> Contacts GetContact(ctx, xeroTenantId, contactID)

Allows you to retrieve a single contacts in a Xero organisation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**contactID** | [**string**](.md)| Unique identifier for a Contact | 

### Return type

[**Contacts**](Contacts.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContactAttachmentByFileName

> *os.File GetContactAttachmentByFileName(ctx, xeroTenantId, contactID, fileName, contentType)

Allows you to retrieve Attachments on Contacts by file name

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**contactID** | [**string**](.md)| Unique identifier for a Contact | 
**fileName** | **string**| Name for the file you are attaching | 
**contentType** | **string**| The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContactAttachmentById

> *os.File GetContactAttachmentById(ctx, xeroTenantId, contactID, attachmentID, contentType)

Allows you to retrieve Attachments on Contacts

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**contactID** | [**string**](.md)| Unique identifier for a Contact | 
**attachmentID** | [**string**](.md)| Unique identifier for a Attachment | 
**contentType** | **string**| The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContactAttachments

> Attachments GetContactAttachments(ctx, xeroTenantId, contactID)

Allows you to retrieve, add and update contacts in a Xero organisation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**contactID** | [**string**](.md)| Unique identifier for a Contact | 

### Return type

[**Attachments**](Attachments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContactByContactNumber

> Contacts GetContactByContactNumber(ctx, xeroTenantId, contactNumber)

Allows you to retrieve a single contact by Contact Number in a Xero organisation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**contactNumber** | **string**| This field is read only on the Xero contact screen, used to identify contacts in external systems (max length &#x3D; 50). | 

### Return type

[**Contacts**](Contacts.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContactCISSettings

> CisSettings GetContactCISSettings(ctx, xeroTenantId, contactID)

Allows you to retrieve CISSettings for a contact in a Xero organisation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**contactID** | [**string**](.md)| Unique identifier for a Contact | 

### Return type

[**CisSettings**](CISSettings.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContactGroup

> ContactGroups GetContactGroup(ctx, xeroTenantId, contactGroupID)

Allows you to retrieve a unique Contact Group by ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**contactGroupID** | [**string**](.md)| Unique identifier for a Contact Group | 

### Return type

[**ContactGroups**](ContactGroups.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContactGroups

> ContactGroups GetContactGroups(ctx, xeroTenantId, optional)

Allows you to retrieve the ContactID and Name of all the contacts in a contact group

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
 **optional** | ***GetContactGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetContactGroupsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **where** | **optional.String**| Filter by an any element | 
 **order** | **optional.String**| Order by an any element | 

### Return type

[**ContactGroups**](ContactGroups.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContactHistory

> HistoryRecords GetContactHistory(ctx, xeroTenantId, contactID)

Allows you to retrieve a history records of an Contact

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**contactID** | [**string**](.md)| Unique identifier for a Contact | 

### Return type

[**HistoryRecords**](HistoryRecords.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContacts

> Contacts GetContacts(ctx, xeroTenantId, optional)

Allows you to retrieve all contacts in a Xero organisation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
 **optional** | ***GetContactsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetContactsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifModifiedSince** | **optional.Time**| Only records created or modified since this timestamp will be returned | 
 **where** | **optional.String**| Filter by an any element | 
 **order** | **optional.String**| Order by an any element | 
 **iDs** | [**optional.Interface of []string**](string.md)| Filter by a comma separated list of ContactIDs. Allows you to retrieve a specific set of contacts in a single call. | 
 **page** | **optional.Int32**| e.g. page&#x3D;1 - Up to 100 contacts will be returned in a single API call. | 
 **includeArchived** | **optional.Bool**| e.g. includeArchived&#x3D;true - Contacts with a status of ARCHIVED will be included in the response | 

### Return type

[**Contacts**](Contacts.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCreditNote

> CreditNotes GetCreditNote(ctx, xeroTenantId, creditNoteID, optional)

Allows you to retrieve a specific credit note

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**creditNoteID** | [**string**](.md)| Unique identifier for a Credit Note | 
 **optional** | ***GetCreditNoteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCreditNoteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **unitdp** | **optional.Int32**| e.g. unitdp&#x3D;4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts | 

### Return type

[**CreditNotes**](CreditNotes.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCreditNoteAsPdf

> *os.File GetCreditNoteAsPdf(ctx, xeroTenantId, creditNoteID)

Allows you to retrieve Credit Note as PDF files

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**creditNoteID** | [**string**](.md)| Unique identifier for a Credit Note | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pdf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCreditNoteAttachmentByFileName

> *os.File GetCreditNoteAttachmentByFileName(ctx, xeroTenantId, creditNoteID, fileName, contentType)

Allows you to retrieve Attachments on CreditNote by file name

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**creditNoteID** | [**string**](.md)| Unique identifier for a Credit Note | 
**fileName** | **string**| Name of the file you are attaching to Credit Note | 
**contentType** | **string**| The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCreditNoteAttachmentById

> *os.File GetCreditNoteAttachmentById(ctx, xeroTenantId, creditNoteID, attachmentID, contentType)

Allows you to retrieve Attachments on CreditNote

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**creditNoteID** | [**string**](.md)| Unique identifier for a Credit Note | 
**attachmentID** | [**string**](.md)| Unique identifier for a Attachment | 
**contentType** | **string**| The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCreditNoteAttachments

> Attachments GetCreditNoteAttachments(ctx, xeroTenantId, creditNoteID)

Allows you to retrieve Attachments for credit notes

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**creditNoteID** | [**string**](.md)| Unique identifier for a Credit Note | 

### Return type

[**Attachments**](Attachments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCreditNoteHistory

> HistoryRecords GetCreditNoteHistory(ctx, xeroTenantId, creditNoteID)

Allows you to retrieve a history records of an CreditNote

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**creditNoteID** | [**string**](.md)| Unique identifier for a Credit Note | 

### Return type

[**HistoryRecords**](HistoryRecords.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCreditNotes

> CreditNotes GetCreditNotes(ctx, xeroTenantId, optional)

Allows you to retrieve any credit notes

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
 **optional** | ***GetCreditNotesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCreditNotesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifModifiedSince** | **optional.Time**| Only records created or modified since this timestamp will be returned | 
 **where** | **optional.String**| Filter by an any element | 
 **order** | **optional.String**| Order by an any element | 
 **page** | **optional.Int32**| e.g. page&#x3D;1 – Up to 100 credit notes will be returned in a single API call with line items shown for each credit note | 
 **unitdp** | **optional.Int32**| e.g. unitdp&#x3D;4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts | 

### Return type

[**CreditNotes**](CreditNotes.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrencies

> Currencies GetCurrencies(ctx, xeroTenantId, optional)

Allows you to retrieve currencies for your organisation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
 **optional** | ***GetCurrenciesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCurrenciesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **where** | **optional.String**| Filter by an any element | 
 **order** | **optional.String**| Order by an any element | 

### Return type

[**Currencies**](Currencies.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmployee

> Employees GetEmployee(ctx, xeroTenantId, employeeID)

Allows you to retrieve a specific employee used in Xero payrun

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**employeeID** | [**string**](.md)| Unique identifier for a Employee | 

### Return type

[**Employees**](Employees.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmployees

> Employees GetEmployees(ctx, xeroTenantId, optional)

Allows you to retrieve employees used in Xero payrun

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
 **optional** | ***GetEmployeesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetEmployeesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifModifiedSince** | **optional.Time**| Only records created or modified since this timestamp will be returned | 
 **where** | **optional.String**| Filter by an any element | 
 **order** | **optional.String**| Order by an any element | 

### Return type

[**Employees**](Employees.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExpenseClaim

> ExpenseClaims GetExpenseClaim(ctx, xeroTenantId, expenseClaimID)

Allows you to retrieve a specified expense claim

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**expenseClaimID** | [**string**](.md)| Unique identifier for a ExpenseClaim | 

### Return type

[**ExpenseClaims**](ExpenseClaims.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExpenseClaimHistory

> HistoryRecords GetExpenseClaimHistory(ctx, xeroTenantId, expenseClaimID)

Allows you to retrieve a history records of an ExpenseClaim

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**expenseClaimID** | [**string**](.md)| Unique identifier for a ExpenseClaim | 

### Return type

[**HistoryRecords**](HistoryRecords.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExpenseClaims

> ExpenseClaims GetExpenseClaims(ctx, xeroTenantId, optional)

Allows you to retrieve expense claims

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
 **optional** | ***GetExpenseClaimsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetExpenseClaimsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifModifiedSince** | **optional.Time**| Only records created or modified since this timestamp will be returned | 
 **where** | **optional.String**| Filter by an any element | 
 **order** | **optional.String**| Order by an any element | 

### Return type

[**ExpenseClaims**](ExpenseClaims.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoice

> Invoices GetInvoice(ctx, xeroTenantId, invoiceID, optional)

Allows you to retrieve a specified sales invoice or purchase bill

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**invoiceID** | [**string**](.md)| Unique identifier for an Invoice | 
 **optional** | ***GetInvoiceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetInvoiceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **unitdp** | **optional.Int32**| e.g. unitdp&#x3D;4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts | 

### Return type

[**Invoices**](Invoices.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoiceAsPdf

> *os.File GetInvoiceAsPdf(ctx, xeroTenantId, invoiceID)

Allows you to retrieve invoices or purchase bills as PDF files

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**invoiceID** | [**string**](.md)| Unique identifier for an Invoice | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pdf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoiceAttachmentByFileName

> *os.File GetInvoiceAttachmentByFileName(ctx, xeroTenantId, invoiceID, fileName, contentType)

Allows you to retrieve Attachment on invoices or purchase bills by it's filename

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**invoiceID** | [**string**](.md)| Unique identifier for an Invoice | 
**fileName** | **string**| Name of the file you are attaching | 
**contentType** | **string**| The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoiceAttachmentById

> *os.File GetInvoiceAttachmentById(ctx, xeroTenantId, invoiceID, attachmentID, contentType)

Allows you to retrieve a specified Attachment on invoices or purchase bills by it's ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**invoiceID** | [**string**](.md)| Unique identifier for an Invoice | 
**attachmentID** | [**string**](.md)| Unique identifier for an Attachment | 
**contentType** | **string**| The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoiceAttachments

> Attachments GetInvoiceAttachments(ctx, xeroTenantId, invoiceID)

Allows you to retrieve Attachments on invoices or purchase bills

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**invoiceID** | [**string**](.md)| Unique identifier for an Invoice | 

### Return type

[**Attachments**](Attachments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoiceHistory

> HistoryRecords GetInvoiceHistory(ctx, xeroTenantId, invoiceID)

Allows you to retrieve a history records of an invoice

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**invoiceID** | [**string**](.md)| Unique identifier for an Invoice | 

### Return type

[**HistoryRecords**](HistoryRecords.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoiceReminders

> InvoiceReminders GetInvoiceReminders(ctx, xeroTenantId)

Allows you to retrieve invoice reminder settings

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 

### Return type

[**InvoiceReminders**](InvoiceReminders.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoices

> Invoices GetInvoices(ctx, xeroTenantId, optional)

Allows you to retrieve any sales invoices or purchase bills

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
 **optional** | ***GetInvoicesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetInvoicesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifModifiedSince** | **optional.Time**| Only records created or modified since this timestamp will be returned | 
 **where** | **optional.String**| Filter by an any element | 
 **order** | **optional.String**| Order by an any element | 
 **iDs** | [**optional.Interface of []string**](string.md)| Filter by a comma-separated list of InvoicesIDs. | 
 **invoiceNumbers** | [**optional.Interface of []string**](string.md)| Filter by a comma-separated list of InvoiceNumbers. | 
 **contactIDs** | [**optional.Interface of []string**](string.md)| Filter by a comma-separated list of ContactIDs. | 
 **statuses** | [**optional.Interface of []string**](string.md)| Filter by a comma-separated list Statuses. For faster response times we recommend using these explicit parameters instead of passing OR conditions into the Where filter. | 
 **page** | **optional.Int32**| e.g. page&#x3D;1 – Up to 100 invoices will be returned in a single API call with line items shown for each invoice | 
 **includeArchived** | **optional.Bool**| e.g. includeArchived&#x3D;true - Contacts with a status of ARCHIVED will be included in the response | 
 **createdByMyApp** | **optional.Bool**| When set to true you&#39;ll only retrieve Invoices created by your app | 
 **unitdp** | **optional.Int32**| e.g. unitdp&#x3D;4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts | 

### Return type

[**Invoices**](Invoices.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItem

> Items GetItem(ctx, xeroTenantId, itemID, optional)

Allows you to retrieve a specified item

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**itemID** | [**string**](.md)| Unique identifier for an Item | 
 **optional** | ***GetItemOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetItemOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **unitdp** | **optional.Int32**| e.g. unitdp&#x3D;4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts | 

### Return type

[**Items**](Items.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemHistory

> HistoryRecords GetItemHistory(ctx, xeroTenantId, itemID)

Allows you to retrieve history for items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**itemID** | [**string**](.md)| Unique identifier for an Item | 

### Return type

[**HistoryRecords**](HistoryRecords.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItems

> Items GetItems(ctx, xeroTenantId, optional)

Allows you to retrieve any items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
 **optional** | ***GetItemsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetItemsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifModifiedSince** | **optional.Time**| Only records created or modified since this timestamp will be returned | 
 **where** | **optional.String**| Filter by an any element | 
 **order** | **optional.String**| Order by an any element | 
 **unitdp** | **optional.Int32**| e.g. unitdp&#x3D;4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts | 

### Return type

[**Items**](Items.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJournal

> Journals GetJournal(ctx, xeroTenantId, journalID)

Allows you to retrieve a specified journals.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**journalID** | [**string**](.md)| Unique identifier for a Journal | 

### Return type

[**Journals**](Journals.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJournals

> Journals GetJournals(ctx, xeroTenantId, optional)

Allows you to retrieve any journals.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
 **optional** | ***GetJournalsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetJournalsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifModifiedSince** | **optional.Time**| Only records created or modified since this timestamp will be returned | 
 **offset** | **optional.Int32**| Offset by a specified journal number. e.g. journals with a JournalNumber greater than the offset will be returned | 
 **paymentsOnly** | **optional.Bool**| Filter to retrieve journals on a cash basis. Journals are returned on an accrual basis by default. | 

### Return type

[**Journals**](Journals.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLinkedTransaction

> LinkedTransactions GetLinkedTransaction(ctx, xeroTenantId, linkedTransactionID)

Allows you to retrieve a specified linked transactions (billable expenses)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**linkedTransactionID** | [**string**](.md)| Unique identifier for a LinkedTransaction | 

### Return type

[**LinkedTransactions**](LinkedTransactions.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLinkedTransactions

> LinkedTransactions GetLinkedTransactions(ctx, xeroTenantId, optional)

Retrieve linked transactions (billable expenses)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
 **optional** | ***GetLinkedTransactionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetLinkedTransactionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Up to 100 linked transactions will be returned in a single API call. Use the page parameter to specify the page to be returned e.g. page&#x3D;1. | 
 **linkedTransactionID** | **optional.String**| The Xero identifier for an Linked Transaction | 
 **sourceTransactionID** | **optional.String**| Filter by the SourceTransactionID. Get the linked transactions created from a particular ACCPAY invoice | 
 **contactID** | **optional.String**| Filter by the ContactID. Get all the linked transactions that have been assigned to a particular customer. | 
 **status** | **optional.String**| Filter by the combination of ContactID and Status. Get  the linked transactions associaed to a  customer and with a status | 
 **targetTransactionID** | **optional.String**| Filter by the TargetTransactionID. Get all the linked transactions allocated to a particular ACCREC invoice | 

### Return type

[**LinkedTransactions**](LinkedTransactions.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManualJournal

> ManualJournals GetManualJournal(ctx, xeroTenantId, manualJournalID)

Allows you to retrieve a specified manual journals

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**manualJournalID** | [**string**](.md)| Unique identifier for a ManualJournal | 

### Return type

[**ManualJournals**](ManualJournals.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManualJournalAttachmentByFileName

> *os.File GetManualJournalAttachmentByFileName(ctx, xeroTenantId, manualJournalID, fileName, contentType)

Allows you to retrieve specified Attachment on ManualJournal by file name

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**manualJournalID** | [**string**](.md)| Unique identifier for a ManualJournal | 
**fileName** | **string**| The name of the file being attached to a ManualJournal | 
**contentType** | **string**| The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManualJournalAttachmentById

> *os.File GetManualJournalAttachmentById(ctx, xeroTenantId, manualJournalID, attachmentID, contentType)

Allows you to retrieve specified Attachment on ManualJournals

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**manualJournalID** | [**string**](.md)| Unique identifier for a ManualJournal | 
**attachmentID** | [**string**](.md)| Unique identifier for a Attachment | 
**contentType** | **string**| The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManualJournalAttachments

> Attachments GetManualJournalAttachments(ctx, xeroTenantId, manualJournalID)

Allows you to retrieve Attachment for manual journals

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**manualJournalID** | [**string**](.md)| Unique identifier for a ManualJournal | 

### Return type

[**Attachments**](Attachments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManualJournals

> ManualJournals GetManualJournals(ctx, xeroTenantId, optional)

Allows you to retrieve any manual journals

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
 **optional** | ***GetManualJournalsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetManualJournalsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifModifiedSince** | **optional.Time**| Only records created or modified since this timestamp will be returned | 
 **where** | **optional.String**| Filter by an any element | 
 **order** | **optional.String**| Order by an any element | 
 **page** | **optional.Int32**| e.g. page&#x3D;1 – Up to 100 manual journals will be returned in a single API call with line items shown for each overpayment | 

### Return type

[**ManualJournals**](ManualJournals.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOnlineInvoice

> OnlineInvoices GetOnlineInvoice(ctx, xeroTenantId, invoiceID)

Allows you to retrieve a URL to an online invoice

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**invoiceID** | [**string**](.md)| Unique identifier for an Invoice | 

### Return type

[**OnlineInvoices**](OnlineInvoices.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganisationCISSettings

> CisOrgSetting GetOrganisationCISSettings(ctx, xeroTenantId, organisationID)

Allows you To verify if an organisation is using contruction industry scheme, you can retrieve the CIS settings for the organistaion.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**organisationID** | [**string**](.md)| The unique Xero identifier for an organisation | 

### Return type

[**CisOrgSetting**](CISOrgSetting.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganisations

> Organisations GetOrganisations(ctx, xeroTenantId)

Allows you to retrieve Organisation details

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 

### Return type

[**Organisations**](Organisations.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOverpayment

> Overpayments GetOverpayment(ctx, xeroTenantId, overpaymentID)

Allows you to retrieve a specified overpayments

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**overpaymentID** | [**string**](.md)| Unique identifier for a Overpayment | 

### Return type

[**Overpayments**](Overpayments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOverpaymentHistory

> HistoryRecords GetOverpaymentHistory(ctx, xeroTenantId, overpaymentID)

Allows you to retrieve a history records of an Overpayment

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**overpaymentID** | [**string**](.md)| Unique identifier for a Overpayment | 

### Return type

[**HistoryRecords**](HistoryRecords.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOverpayments

> Overpayments GetOverpayments(ctx, xeroTenantId, optional)

Allows you to retrieve overpayments

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
 **optional** | ***GetOverpaymentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetOverpaymentsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifModifiedSince** | **optional.Time**| Only records created or modified since this timestamp will be returned | 
 **where** | **optional.String**| Filter by an any element | 
 **order** | **optional.String**| Order by an any element | 
 **page** | **optional.Int32**| e.g. page&#x3D;1 – Up to 100 overpayments will be returned in a single API call with line items shown for each overpayment | 
 **unitdp** | **optional.Int32**| e.g. unitdp&#x3D;4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts | 

### Return type

[**Overpayments**](Overpayments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPayment

> Payments GetPayment(ctx, xeroTenantId, paymentID)

Allows you to retrieve a specified payment for invoices and credit notes

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**paymentID** | [**string**](.md)| Unique identifier for a Payment | 

### Return type

[**Payments**](Payments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPaymentHistory

> HistoryRecords GetPaymentHistory(ctx, xeroTenantId, paymentID)

Allows you to retrieve history records of a payment

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**paymentID** | [**string**](.md)| Unique identifier for a Payment | 

### Return type

[**HistoryRecords**](HistoryRecords.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPaymentServices

> PaymentServices GetPaymentServices(ctx, xeroTenantId)

Allows you to retrieve payment services

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 

### Return type

[**PaymentServices**](PaymentServices.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPayments

> Payments GetPayments(ctx, xeroTenantId, optional)

Allows you to retrieve payments for invoices and credit notes

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
 **optional** | ***GetPaymentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetPaymentsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifModifiedSince** | **optional.Time**| Only records created or modified since this timestamp will be returned | 
 **where** | **optional.String**| Filter by an any element | 
 **order** | **optional.String**| Order by an any element | 
 **page** | **optional.Int32**| Up to 100 payments will be returned in a single API call | 

### Return type

[**Payments**](Payments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrepayment

> Prepayments GetPrepayment(ctx, xeroTenantId, prepaymentID)

Allows you to retrieve a specified prepayments

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**prepaymentID** | [**string**](.md)| Unique identifier for a PrePayment | 

### Return type

[**Prepayments**](Prepayments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrepaymentHistory

> HistoryRecords GetPrepaymentHistory(ctx, xeroTenantId, prepaymentID)

Allows you to retrieve a history records of an Prepayment

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**prepaymentID** | [**string**](.md)| Unique identifier for a PrePayment | 

### Return type

[**HistoryRecords**](HistoryRecords.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrepayments

> Prepayments GetPrepayments(ctx, xeroTenantId, optional)

Allows you to retrieve prepayments

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
 **optional** | ***GetPrepaymentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetPrepaymentsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifModifiedSince** | **optional.Time**| Only records created or modified since this timestamp will be returned | 
 **where** | **optional.String**| Filter by an any element | 
 **order** | **optional.String**| Order by an any element | 
 **page** | **optional.Int32**| e.g. page&#x3D;1 – Up to 100 prepayments will be returned in a single API call with line items shown for each overpayment | 
 **unitdp** | **optional.Int32**| e.g. unitdp&#x3D;4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts | 

### Return type

[**Prepayments**](Prepayments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPurchaseOrder

> PurchaseOrders GetPurchaseOrder(ctx, xeroTenantId, purchaseOrderID)

Allows you to retrieve a specified purchase orders

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**purchaseOrderID** | [**string**](.md)| Unique identifier for a PurchaseOrder | 

### Return type

[**PurchaseOrders**](PurchaseOrders.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPurchaseOrderAsPdf

> *os.File GetPurchaseOrderAsPdf(ctx, xeroTenantId, purchaseOrderID)

Allows you to retrieve purchase orders as PDF files

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**purchaseOrderID** | [**string**](.md)| Unique identifier for an Purchase Order | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pdf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPurchaseOrderByNumber

> PurchaseOrders GetPurchaseOrderByNumber(ctx, xeroTenantId, purchaseOrderNumber)

Allows you to retrieve a specified purchase orders

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**purchaseOrderNumber** | **string**| Unique identifier for a PurchaseOrder | 

### Return type

[**PurchaseOrders**](PurchaseOrders.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPurchaseOrderHistory

> HistoryRecords GetPurchaseOrderHistory(ctx, xeroTenantId, purchaseOrderID)

Allows you to retrieve history for PurchaseOrder

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**purchaseOrderID** | [**string**](.md)| Unique identifier for a PurchaseOrder | 

### Return type

[**HistoryRecords**](HistoryRecords.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPurchaseOrders

> PurchaseOrders GetPurchaseOrders(ctx, xeroTenantId, optional)

Allows you to retrieve purchase orders

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
 **optional** | ***GetPurchaseOrdersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetPurchaseOrdersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifModifiedSince** | **optional.Time**| Only records created or modified since this timestamp will be returned | 
 **status** | **optional.String**| Filter by purchase order status | 
 **dateFrom** | **optional.String**| Filter by purchase order date (e.g. GET https://.../PurchaseOrders?DateFrom&#x3D;2015-12-01&amp;DateTo&#x3D;2015-12-31 | 
 **dateTo** | **optional.String**| Filter by purchase order date (e.g. GET https://.../PurchaseOrders?DateFrom&#x3D;2015-12-01&amp;DateTo&#x3D;2015-12-31 | 
 **order** | **optional.String**| Order by an any element | 
 **page** | **optional.Int32**| To specify a page, append the page parameter to the URL e.g. ?page&#x3D;1. If there are 100 records in the response you will need to check if there is any more data by fetching the next page e.g ?page&#x3D;2 and continuing this process until no more results are returned. | 

### Return type

[**PurchaseOrders**](PurchaseOrders.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQuote

> Quotes GetQuote(ctx, xeroTenantId, quoteID)

Allows you to retrieve a specified quote

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**quoteID** | [**string**](.md)| Unique identifier for an Quote | 

### Return type

[**Quotes**](Quotes.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQuoteAsPdf

> *os.File GetQuoteAsPdf(ctx, xeroTenantId, quoteID)

Allows you to retrieve quotes as PDF files

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**quoteID** | [**string**](.md)| Unique identifier for an Quote | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pdf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQuoteAttachmentByFileName

> *os.File GetQuoteAttachmentByFileName(ctx, xeroTenantId, quoteID, fileName, contentType)

Allows you to retrieve Attachment on Quote by Filename

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**quoteID** | [**string**](.md)| Unique identifier for Quote object | 
**fileName** | **string**| Name of the attachment | 
**contentType** | **string**| The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQuoteAttachmentById

> *os.File GetQuoteAttachmentById(ctx, xeroTenantId, quoteID, attachmentID, contentType)

Allows you to retrieve specific Attachment on Quote

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**quoteID** | [**string**](.md)| Unique identifier for Quote object | 
**attachmentID** | [**string**](.md)| Unique identifier for Attachment object | 
**contentType** | **string**| The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQuoteAttachments

> Attachments GetQuoteAttachments(ctx, xeroTenantId, quoteID)

Allows you to retrieve Attachments for Quotes

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**quoteID** | [**string**](.md)| Unique identifier for Quote object | 

### Return type

[**Attachments**](Attachments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQuoteHistory

> HistoryRecords GetQuoteHistory(ctx, xeroTenantId, quoteID)

Allows you to retrieve a history records of an quote

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**quoteID** | [**string**](.md)| Unique identifier for an Quote | 

### Return type

[**HistoryRecords**](HistoryRecords.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQuotes

> Quotes GetQuotes(ctx, xeroTenantId, optional)

Allows you to retrieve any sales quotes

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
 **optional** | ***GetQuotesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetQuotesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifModifiedSince** | **optional.Time**| Only records created or modified since this timestamp will be returned | 
 **dateFrom** | **optional.String**| Filter for quotes after a particular date | 
 **dateTo** | **optional.String**| Filter for quotes before a particular date | 
 **expiryDateFrom** | **optional.String**| Filter for quotes expiring after a particular date | 
 **expiryDateTo** | **optional.String**| Filter for quotes before a particular date | 
 **contactID** | [**optional.Interface of string**](.md)| Filter for quotes belonging to a particular contact | 
 **status** | **optional.String**| Filter for quotes of a particular Status | 
 **page** | **optional.Int32**| e.g. page&#x3D;1 – Up to 100 Quotes will be returned in a single API call with line items shown for each quote | 
 **order** | **optional.String**| Order by an any element | 

### Return type

[**Quotes**](Quotes.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReceipt

> Receipts GetReceipt(ctx, xeroTenantId, receiptID, optional)

Allows you to retrieve a specified draft expense claim receipts

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**receiptID** | [**string**](.md)| Unique identifier for a Receipt | 
 **optional** | ***GetReceiptOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetReceiptOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **unitdp** | **optional.Int32**| e.g. unitdp&#x3D;4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts | 

### Return type

[**Receipts**](Receipts.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReceiptAttachmentByFileName

> *os.File GetReceiptAttachmentByFileName(ctx, xeroTenantId, receiptID, fileName, contentType)

Allows you to retrieve Attachments on expense claim receipts by file name

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**receiptID** | [**string**](.md)| Unique identifier for a Receipt | 
**fileName** | **string**| The name of the file being attached to the Receipt | 
**contentType** | **string**| The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReceiptAttachmentById

> *os.File GetReceiptAttachmentById(ctx, xeroTenantId, receiptID, attachmentID, contentType)

Allows you to retrieve Attachments on expense claim receipts by ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**receiptID** | [**string**](.md)| Unique identifier for a Receipt | 
**attachmentID** | [**string**](.md)| Unique identifier for a Attachment | 
**contentType** | **string**| The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReceiptAttachments

> Attachments GetReceiptAttachments(ctx, xeroTenantId, receiptID)

Allows you to retrieve Attachments for expense claim receipts

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**receiptID** | [**string**](.md)| Unique identifier for a Receipt | 

### Return type

[**Attachments**](Attachments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReceiptHistory

> HistoryRecords GetReceiptHistory(ctx, xeroTenantId, receiptID)

Allows you to retrieve a history records of an Receipt

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**receiptID** | [**string**](.md)| Unique identifier for a Receipt | 

### Return type

[**HistoryRecords**](HistoryRecords.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReceipts

> Receipts GetReceipts(ctx, xeroTenantId, optional)

Allows you to retrieve draft expense claim receipts for any user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
 **optional** | ***GetReceiptsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetReceiptsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifModifiedSince** | **optional.Time**| Only records created or modified since this timestamp will be returned | 
 **where** | **optional.String**| Filter by an any element | 
 **order** | **optional.String**| Order by an any element | 
 **unitdp** | **optional.Int32**| e.g. unitdp&#x3D;4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts | 

### Return type

[**Receipts**](Receipts.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepeatingInvoice

> RepeatingInvoices GetRepeatingInvoice(ctx, xeroTenantId, repeatingInvoiceID)

Allows you to retrieve a specified repeating invoice

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**repeatingInvoiceID** | [**string**](.md)| Unique identifier for a Repeating Invoice | 

### Return type

[**RepeatingInvoices**](RepeatingInvoices.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepeatingInvoiceAttachmentByFileName

> *os.File GetRepeatingInvoiceAttachmentByFileName(ctx, xeroTenantId, repeatingInvoiceID, fileName, contentType)

Allows you to retrieve specified attachment on repeating invoices by file name

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**repeatingInvoiceID** | [**string**](.md)| Unique identifier for a Repeating Invoice | 
**fileName** | **string**| The name of the file being attached to a Repeating Invoice | 
**contentType** | **string**| The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepeatingInvoiceAttachmentById

> *os.File GetRepeatingInvoiceAttachmentById(ctx, xeroTenantId, repeatingInvoiceID, attachmentID, contentType)

Allows you to retrieve a specified Attachments on repeating invoices

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**repeatingInvoiceID** | [**string**](.md)| Unique identifier for a Repeating Invoice | 
**attachmentID** | [**string**](.md)| Unique identifier for a Attachment | 
**contentType** | **string**| The mime type of the attachment file you are retrieving i.e image/jpg, application/pdf | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepeatingInvoiceAttachments

> Attachments GetRepeatingInvoiceAttachments(ctx, xeroTenantId, repeatingInvoiceID)

Allows you to retrieve Attachments on repeating invoice

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**repeatingInvoiceID** | [**string**](.md)| Unique identifier for a Repeating Invoice | 

### Return type

[**Attachments**](Attachments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepeatingInvoiceHistory

> HistoryRecords GetRepeatingInvoiceHistory(ctx, xeroTenantId, repeatingInvoiceID)

Allows you to retrieve history for a repeating invoice

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**repeatingInvoiceID** | [**string**](.md)| Unique identifier for a Repeating Invoice | 

### Return type

[**HistoryRecords**](HistoryRecords.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepeatingInvoices

> RepeatingInvoices GetRepeatingInvoices(ctx, xeroTenantId, optional)

Allows you to retrieve any repeating invoices

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
 **optional** | ***GetRepeatingInvoicesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetRepeatingInvoicesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **where** | **optional.String**| Filter by an any element | 
 **order** | **optional.String**| Order by an any element | 

### Return type

[**RepeatingInvoices**](RepeatingInvoices.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportAgedPayablesByContact

> ReportWithRows GetReportAgedPayablesByContact(ctx, xeroTenantId, contactId, optional)

Allows you to retrieve report for AgedPayablesByContact

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**contactId** | [**string**](.md)| Unique identifier for a Contact | 
 **optional** | ***GetReportAgedPayablesByContactOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetReportAgedPayablesByContactOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **date** | **optional.String**| The date of the Aged Payables By Contact report | 
 **fromDate** | **optional.String**| The from date of the Aged Payables By Contact report | 
 **toDate** | **optional.String**| The to date of the Aged Payables By Contact report | 

### Return type

[**ReportWithRows**](ReportWithRows.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportAgedReceivablesByContact

> ReportWithRows GetReportAgedReceivablesByContact(ctx, xeroTenantId, contactId, optional)

Allows you to retrieve report for AgedReceivablesByContact

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**contactId** | [**string**](.md)| Unique identifier for a Contact | 
 **optional** | ***GetReportAgedReceivablesByContactOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetReportAgedReceivablesByContactOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **date** | **optional.String**| The date of the Aged Receivables By Contact report | 
 **fromDate** | **optional.String**| The from date of the Aged Receivables By Contact report | 
 **toDate** | **optional.String**| The to date of the Aged Receivables By Contact report | 

### Return type

[**ReportWithRows**](ReportWithRows.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportBASorGST

> ReportWithRows GetReportBASorGST(ctx, xeroTenantId, reportID)

Allows you to retrieve report for BAS only valid for AU orgs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**reportID** | **string**| Unique identifier for a Report | 

### Return type

[**ReportWithRows**](ReportWithRows.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportBASorGSTList

> ReportWithRows GetReportBASorGSTList(ctx, xeroTenantId)

Allows you to retrieve report for BAS only valid for AU orgs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 

### Return type

[**ReportWithRows**](ReportWithRows.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportBalanceSheet

> ReportWithRows GetReportBalanceSheet(ctx, xeroTenantId, optional)

Allows you to retrieve report for BalanceSheet

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
 **optional** | ***GetReportBalanceSheetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetReportBalanceSheetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **date** | **optional.String**| The date of the Balance Sheet report | 
 **periods** | **optional.Int32**| The number of periods for the Balance Sheet report | 
 **timeframe** | **optional.String**| The period size to compare to (MONTH, QUARTER, YEAR) | 
 **trackingOptionID1** | **optional.String**| The tracking option 1 for the Balance Sheet report | 
 **trackingOptionID2** | **optional.String**| The tracking option 2 for the Balance Sheet report | 
 **standardLayout** | **optional.Bool**| The standard layout boolean for the Balance Sheet report | 
 **paymentsOnly** | **optional.Bool**| return a cash basis for the Balance Sheet report | 

### Return type

[**ReportWithRows**](ReportWithRows.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportBankSummary

> ReportWithRows GetReportBankSummary(ctx, xeroTenantId, optional)

Allows you to retrieve report for BankSummary

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
 **optional** | ***GetReportBankSummaryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetReportBankSummaryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fromDate** | **optional.String**| The from date for the Bank Summary report e.g. 2018-03-31 | 
 **toDate** | **optional.String**| The to date for the Bank Summary report e.g. 2018-03-31 | 

### Return type

[**ReportWithRows**](ReportWithRows.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportBudgetSummary

> ReportWithRows GetReportBudgetSummary(ctx, xeroTenantId, optional)

Allows you to retrieve report for Budget Summary

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
 **optional** | ***GetReportBudgetSummaryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetReportBudgetSummaryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **date** | **optional.String**| The date for the Bank Summary report e.g. 2018-03-31 | 
 **period** | **optional.Int32**| The number of periods to compare (integer between 1 and 12) | 
 **timeframe** | **optional.Int32**| The period size to compare to (1&#x3D;month, 3&#x3D;quarter, 12&#x3D;year) | 

### Return type

[**ReportWithRows**](ReportWithRows.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportExecutiveSummary

> ReportWithRows GetReportExecutiveSummary(ctx, xeroTenantId, optional)

Allows you to retrieve report for ExecutiveSummary

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
 **optional** | ***GetReportExecutiveSummaryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetReportExecutiveSummaryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **date** | **optional.String**| The date for the Bank Summary report e.g. 2018-03-31 | 

### Return type

[**ReportWithRows**](ReportWithRows.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportProfitAndLoss

> ReportWithRows GetReportProfitAndLoss(ctx, xeroTenantId, optional)

Allows you to retrieve report for ProfitAndLoss

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
 **optional** | ***GetReportProfitAndLossOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetReportProfitAndLossOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fromDate** | **optional.String**| The from date for the ProfitAndLoss report e.g. 2018-03-31 | 
 **toDate** | **optional.String**| The to date for the ProfitAndLoss report e.g. 2018-03-31 | 
 **periods** | **optional.Int32**| The number of periods to compare (integer between 1 and 12) | 
 **timeframe** | **optional.String**| The period size to compare to (MONTH, QUARTER, YEAR) | 
 **trackingCategoryID** | **optional.String**| The trackingCategory 1 for the ProfitAndLoss report | 
 **trackingCategoryID2** | **optional.String**| The trackingCategory 2 for the ProfitAndLoss report | 
 **trackingOptionID** | **optional.String**| The tracking option 1 for the ProfitAndLoss report | 
 **trackingOptionID2** | **optional.String**| The tracking option 2 for the ProfitAndLoss report | 
 **standardLayout** | **optional.Bool**| Return the standard layout for the ProfitAndLoss report | 
 **paymentsOnly** | **optional.Bool**| Return cash only basis for the ProfitAndLoss report | 

### Return type

[**ReportWithRows**](ReportWithRows.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportTenNinetyNine

> Reports GetReportTenNinetyNine(ctx, xeroTenantId, optional)

Allows you to retrieve report for TenNinetyNine

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
 **optional** | ***GetReportTenNinetyNineOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetReportTenNinetyNineOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reportYear** | **optional.String**| The year of the 1099 report | 

### Return type

[**Reports**](Reports.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportTrialBalance

> ReportWithRows GetReportTrialBalance(ctx, xeroTenantId, optional)

Allows you to retrieve report for TrialBalance

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
 **optional** | ***GetReportTrialBalanceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetReportTrialBalanceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **date** | **optional.String**| The date for the Trial Balance report e.g. 2018-03-31 | 
 **paymentsOnly** | **optional.Bool**| Return cash only basis for the Trial Balance report | 

### Return type

[**ReportWithRows**](ReportWithRows.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaxRates

> TaxRates GetTaxRates(ctx, xeroTenantId, optional)

Allows you to retrieve Tax Rates

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
 **optional** | ***GetTaxRatesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetTaxRatesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **where** | **optional.String**| Filter by an any element | 
 **order** | **optional.String**| Order by an any element | 
 **taxType** | **optional.String**| Filter by tax type | 

### Return type

[**TaxRates**](TaxRates.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrackingCategories

> TrackingCategories GetTrackingCategories(ctx, xeroTenantId, optional)

Allows you to retrieve tracking categories and options

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
 **optional** | ***GetTrackingCategoriesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetTrackingCategoriesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **where** | **optional.String**| Filter by an any element | 
 **order** | **optional.String**| Order by an any element | 
 **includeArchived** | **optional.Bool**| e.g. includeArchived&#x3D;true - Categories and options with a status of ARCHIVED will be included in the response | 

### Return type

[**TrackingCategories**](TrackingCategories.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrackingCategory

> TrackingCategories GetTrackingCategory(ctx, xeroTenantId, trackingCategoryID)

Allows you to retrieve tracking categories and options for specified category

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**trackingCategoryID** | [**string**](.md)| Unique identifier for a TrackingCategory | 

### Return type

[**TrackingCategories**](TrackingCategories.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUser

> Users GetUser(ctx, xeroTenantId, userID)

Allows you to retrieve a specified user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**userID** | [**string**](.md)| Unique identifier for a User | 

### Return type

[**Users**](Users.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsers

> Users GetUsers(ctx, xeroTenantId, optional)

Allows you to retrieve users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
 **optional** | ***GetUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetUsersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifModifiedSince** | **optional.Time**| Only records created or modified since this timestamp will be returned | 
 **where** | **optional.String**| Filter by an any element | 
 **order** | **optional.String**| Order by an any element | 

### Return type

[**Users**](Users.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccount

> Accounts UpdateAccount(ctx, xeroTenantId, accountID, accounts)

Allows you to update a chart of accounts

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**accountID** | [**string**](.md)| Unique identifier for retrieving single object | 
**accounts** | [**Accounts**](Accounts.md)| Request of type Accounts array with one Account | 

### Return type

[**Accounts**](Accounts.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccountAttachmentByFileName

> Attachments UpdateAccountAttachmentByFileName(ctx, xeroTenantId, accountID, fileName, body)

Allows you to update Attachment on Account by Filename

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**accountID** | [**string**](.md)| Unique identifier for Account object | 
**fileName** | **string**| Name of the attachment | 
**body** | **string**| Byte array of file in body of request | 

### Return type

[**Attachments**](Attachments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBankTransaction

> BankTransactions UpdateBankTransaction(ctx, xeroTenantId, bankTransactionID, bankTransactions, optional)

Allows you to update a single spend or receive money transaction

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**bankTransactionID** | [**string**](.md)| Xero generated unique identifier for a bank transaction | 
**bankTransactions** | [**BankTransactions**](BankTransactions.md)|  | 
 **optional** | ***UpdateBankTransactionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateBankTransactionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **unitdp** | **optional.Int32**| e.g. unitdp&#x3D;4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts | 

### Return type

[**BankTransactions**](BankTransactions.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBankTransactionAttachmentByFileName

> Attachments UpdateBankTransactionAttachmentByFileName(ctx, xeroTenantId, bankTransactionID, fileName, body)

Allows you to update an Attachment on BankTransaction by Filename

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**bankTransactionID** | [**string**](.md)| Xero generated unique identifier for a bank transaction | 
**fileName** | **string**| The name of the file being attached | 
**body** | **string**| Byte array of file in body of request | 

### Return type

[**Attachments**](Attachments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBankTransferAttachmentByFileName

> Attachments UpdateBankTransferAttachmentByFileName(ctx, xeroTenantId, bankTransferID, fileName, body)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**bankTransferID** | [**string**](.md)| Xero generated unique identifier for a bank transfer | 
**fileName** | **string**| The name of the file being attached to a Bank Transfer | 
**body** | **string**| Byte array of file in body of request | 

### Return type

[**Attachments**](Attachments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateContact

> Contacts UpdateContact(ctx, xeroTenantId, contactID, contacts)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**contactID** | [**string**](.md)| Unique identifier for a Contact | 
**contacts** | [**Contacts**](Contacts.md)| an array of Contacts containing single Contact object with properties to update | 

### Return type

[**Contacts**](Contacts.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateContactAttachmentByFileName

> Attachments UpdateContactAttachmentByFileName(ctx, xeroTenantId, contactID, fileName, body)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**contactID** | [**string**](.md)| Unique identifier for a Contact | 
**fileName** | **string**| Name for the file you are attaching | 
**body** | **string**| Byte array of file in body of request | 

### Return type

[**Attachments**](Attachments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateContactGroup

> ContactGroups UpdateContactGroup(ctx, xeroTenantId, contactGroupID, contactGroups)

Allows you to update a Contact Group

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**contactGroupID** | [**string**](.md)| Unique identifier for a Contact Group | 
**contactGroups** | [**ContactGroups**](ContactGroups.md)| an array of Contact groups with Name of specific group to update | 

### Return type

[**ContactGroups**](ContactGroups.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCreditNote

> CreditNotes UpdateCreditNote(ctx, xeroTenantId, creditNoteID, creditNotes, optional)

Allows you to update a specific credit note

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**creditNoteID** | [**string**](.md)| Unique identifier for a Credit Note | 
**creditNotes** | [**CreditNotes**](CreditNotes.md)| an array of Credit Notes containing credit note details to update | 
 **optional** | ***UpdateCreditNoteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateCreditNoteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **unitdp** | **optional.Int32**| e.g. unitdp&#x3D;4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts | 

### Return type

[**CreditNotes**](CreditNotes.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCreditNoteAttachmentByFileName

> Attachments UpdateCreditNoteAttachmentByFileName(ctx, xeroTenantId, creditNoteID, fileName, body)

Allows you to update Attachments on CreditNote by file name

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**creditNoteID** | [**string**](.md)| Unique identifier for a Credit Note | 
**fileName** | **string**| Name of the file you are attaching to Credit Note | 
**body** | **string**| Byte array of file in body of request | 

### Return type

[**Attachments**](Attachments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateExpenseClaim

> ExpenseClaims UpdateExpenseClaim(ctx, xeroTenantId, expenseClaimID, expenseClaims)

Allows you to update specified expense claims

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**expenseClaimID** | [**string**](.md)| Unique identifier for a ExpenseClaim | 
**expenseClaims** | [**ExpenseClaims**](ExpenseClaims.md)|  | 

### Return type

[**ExpenseClaims**](ExpenseClaims.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInvoice

> Invoices UpdateInvoice(ctx, xeroTenantId, invoiceID, invoices, optional)

Allows you to update a specified sales invoices or purchase bills

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**invoiceID** | [**string**](.md)| Unique identifier for an Invoice | 
**invoices** | [**Invoices**](Invoices.md)|  | 
 **optional** | ***UpdateInvoiceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateInvoiceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **unitdp** | **optional.Int32**| e.g. unitdp&#x3D;4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts | 

### Return type

[**Invoices**](Invoices.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInvoiceAttachmentByFileName

> Attachments UpdateInvoiceAttachmentByFileName(ctx, xeroTenantId, invoiceID, fileName, body)

Allows you to update Attachment on invoices or purchase bills by it's filename

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**invoiceID** | [**string**](.md)| Unique identifier for an Invoice | 
**fileName** | **string**| Name of the file you are attaching | 
**body** | **string**| Byte array of file in body of request | 

### Return type

[**Attachments**](Attachments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateItem

> Items UpdateItem(ctx, xeroTenantId, itemID, items, optional)

Allows you to update a specified item

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**itemID** | [**string**](.md)| Unique identifier for an Item | 
**items** | [**Items**](Items.md)|  | 
 **optional** | ***UpdateItemOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateItemOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **unitdp** | **optional.Int32**| e.g. unitdp&#x3D;4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts | 

### Return type

[**Items**](Items.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLinkedTransaction

> LinkedTransactions UpdateLinkedTransaction(ctx, xeroTenantId, linkedTransactionID, linkedTransactions)

Allows you to update a specified linked transactions (billable expenses)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**linkedTransactionID** | [**string**](.md)| Unique identifier for a LinkedTransaction | 
**linkedTransactions** | [**LinkedTransactions**](LinkedTransactions.md)|  | 

### Return type

[**LinkedTransactions**](LinkedTransactions.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateManualJournal

> ManualJournals UpdateManualJournal(ctx, xeroTenantId, manualJournalID, manualJournals)

Allows you to update a specified manual journal

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**manualJournalID** | [**string**](.md)| Unique identifier for a ManualJournal | 
**manualJournals** | [**ManualJournals**](ManualJournals.md)|  | 

### Return type

[**ManualJournals**](ManualJournals.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateManualJournalAttachmentByFileName

> Attachments UpdateManualJournalAttachmentByFileName(ctx, xeroTenantId, manualJournalID, fileName, body)

Allows you to update a specified Attachment on ManualJournal by file name

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**manualJournalID** | [**string**](.md)| Unique identifier for a ManualJournal | 
**fileName** | **string**| The name of the file being attached to a ManualJournal | 
**body** | **string**| Byte array of file in body of request | 

### Return type

[**Attachments**](Attachments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrCreateBankTransactions

> BankTransactions UpdateOrCreateBankTransactions(ctx, xeroTenantId, bankTransactions, optional)

Allows you to update or create one or more spend or receive money transaction

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**bankTransactions** | [**BankTransactions**](BankTransactions.md)|  | 
 **optional** | ***UpdateOrCreateBankTransactionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateOrCreateBankTransactionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **summarizeErrors** | **optional.Bool**| If false return 200 OK and mix of successfully created obejcts and any with validation errors | [default to false]
 **unitdp** | **optional.Int32**| e.g. unitdp&#x3D;4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts | 

### Return type

[**BankTransactions**](BankTransactions.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrCreateContacts

> Contacts UpdateOrCreateContacts(ctx, xeroTenantId, contacts, optional)

Allows you to update OR create one or more contacts in a Xero organisation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**contacts** | [**Contacts**](Contacts.md)|  | 
 **optional** | ***UpdateOrCreateContactsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateOrCreateContactsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **summarizeErrors** | **optional.Bool**| If false return 200 OK and mix of successfully created obejcts and any with validation errors | [default to false]

### Return type

[**Contacts**](Contacts.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrCreateCreditNotes

> CreditNotes UpdateOrCreateCreditNotes(ctx, xeroTenantId, creditNotes, optional)

Allows you to update OR create one or more credit notes

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**creditNotes** | [**CreditNotes**](CreditNotes.md)| an array of Credit Notes with a single CreditNote object. | 
 **optional** | ***UpdateOrCreateCreditNotesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateOrCreateCreditNotesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **summarizeErrors** | **optional.Bool**| If false return 200 OK and mix of successfully created obejcts and any with validation errors | [default to false]
 **unitdp** | **optional.Int32**| e.g. unitdp&#x3D;4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts | 

### Return type

[**CreditNotes**](CreditNotes.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrCreateEmployees

> Employees UpdateOrCreateEmployees(ctx, xeroTenantId, employees, optional)

Allows you to create a single new employees used in Xero payrun

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**employees** | [**Employees**](Employees.md)| Employees with array of Employee object in body of request | 
 **optional** | ***UpdateOrCreateEmployeesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateOrCreateEmployeesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **summarizeErrors** | **optional.Bool**| If false return 200 OK and mix of successfully created obejcts and any with validation errors | [default to false]

### Return type

[**Employees**](Employees.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrCreateInvoices

> Invoices UpdateOrCreateInvoices(ctx, xeroTenantId, invoices, optional)

Allows you to update OR create one or more sales invoices or purchase bills

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**invoices** | [**Invoices**](Invoices.md)|  | 
 **optional** | ***UpdateOrCreateInvoicesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateOrCreateInvoicesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **summarizeErrors** | **optional.Bool**| If false return 200 OK and mix of successfully created obejcts and any with validation errors | [default to false]
 **unitdp** | **optional.Int32**| e.g. unitdp&#x3D;4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts | 

### Return type

[**Invoices**](Invoices.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrCreateItems

> Items UpdateOrCreateItems(ctx, xeroTenantId, items, optional)

Allows you to update or create one or more items

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**items** | [**Items**](Items.md)|  | 
 **optional** | ***UpdateOrCreateItemsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateOrCreateItemsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **summarizeErrors** | **optional.Bool**| If false return 200 OK and mix of successfully created obejcts and any with validation errors | [default to false]
 **unitdp** | **optional.Int32**| e.g. unitdp&#x3D;4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts | 

### Return type

[**Items**](Items.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrCreateManualJournals

> ManualJournals UpdateOrCreateManualJournals(ctx, xeroTenantId, manualJournals, optional)

Allows you to create a single manual journal

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**manualJournals** | [**ManualJournals**](ManualJournals.md)| ManualJournals array with ManualJournal object in body of request | 
 **optional** | ***UpdateOrCreateManualJournalsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateOrCreateManualJournalsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **summarizeErrors** | **optional.Bool**| If false return 200 OK and mix of successfully created obejcts and any with validation errors | [default to false]

### Return type

[**ManualJournals**](ManualJournals.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrCreatePurchaseOrders

> PurchaseOrders UpdateOrCreatePurchaseOrders(ctx, xeroTenantId, purchaseOrders, optional)

Allows you to update or create one or more purchase orders

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**purchaseOrders** | [**PurchaseOrders**](PurchaseOrders.md)|  | 
 **optional** | ***UpdateOrCreatePurchaseOrdersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateOrCreatePurchaseOrdersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **summarizeErrors** | **optional.Bool**| If false return 200 OK and mix of successfully created obejcts and any with validation errors | [default to false]

### Return type

[**PurchaseOrders**](PurchaseOrders.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrCreateQuotes

> Quotes UpdateOrCreateQuotes(ctx, xeroTenantId, quotes, optional)

Allows you to update OR create one or more quotes

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**quotes** | [**Quotes**](Quotes.md)|  | 
 **optional** | ***UpdateOrCreateQuotesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateOrCreateQuotesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **summarizeErrors** | **optional.Bool**| If false return 200 OK and mix of successfully created obejcts and any with validation errors | [default to false]

### Return type

[**Quotes**](Quotes.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePurchaseOrder

> PurchaseOrders UpdatePurchaseOrder(ctx, xeroTenantId, purchaseOrderID, purchaseOrders)

Allows you to update a specified purchase order

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**purchaseOrderID** | [**string**](.md)| Unique identifier for a PurchaseOrder | 
**purchaseOrders** | [**PurchaseOrders**](PurchaseOrders.md)|  | 

### Return type

[**PurchaseOrders**](PurchaseOrders.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateQuote

> Quotes UpdateQuote(ctx, xeroTenantId, quoteID, quotes)

Allows you to update a specified quote

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**quoteID** | [**string**](.md)| Unique identifier for an Quote | 
**quotes** | [**Quotes**](Quotes.md)|  | 

### Return type

[**Quotes**](Quotes.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateQuoteAttachmentByFileName

> Attachments UpdateQuoteAttachmentByFileName(ctx, xeroTenantId, quoteID, fileName, body)

Allows you to update Attachment on Quote by Filename

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**quoteID** | [**string**](.md)| Unique identifier for Quote object | 
**fileName** | **string**| Name of the attachment | 
**body** | **string**| Byte array of file in body of request | 

### Return type

[**Attachments**](Attachments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateReceipt

> Receipts UpdateReceipt(ctx, xeroTenantId, receiptID, receipts, optional)

Allows you to retrieve a specified draft expense claim receipts

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**receiptID** | [**string**](.md)| Unique identifier for a Receipt | 
**receipts** | [**Receipts**](Receipts.md)|  | 
 **optional** | ***UpdateReceiptOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateReceiptOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **unitdp** | **optional.Int32**| e.g. unitdp&#x3D;4 – (Unit Decimal Places) You can opt in to use four decimal places for unit amounts | 

### Return type

[**Receipts**](Receipts.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateReceiptAttachmentByFileName

> Attachments UpdateReceiptAttachmentByFileName(ctx, xeroTenantId, receiptID, fileName, body)

Allows you to update Attachment on expense claim receipts by file name

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**receiptID** | [**string**](.md)| Unique identifier for a Receipt | 
**fileName** | **string**| The name of the file being attached to the Receipt | 
**body** | **string**| Byte array of file in body of request | 

### Return type

[**Attachments**](Attachments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRepeatingInvoiceAttachmentByFileName

> Attachments UpdateRepeatingInvoiceAttachmentByFileName(ctx, xeroTenantId, repeatingInvoiceID, fileName, body)

Allows you to update specified attachment on repeating invoices by file name

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**repeatingInvoiceID** | [**string**](.md)| Unique identifier for a Repeating Invoice | 
**fileName** | **string**| The name of the file being attached to a Repeating Invoice | 
**body** | **string**| Byte array of file in body of request | 

### Return type

[**Attachments**](Attachments.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTaxRate

> TaxRates UpdateTaxRate(ctx, xeroTenantId, taxRates)

Allows you to update Tax Rates

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**taxRates** | [**TaxRates**](TaxRates.md)|  | 

### Return type

[**TaxRates**](TaxRates.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTrackingCategory

> TrackingCategories UpdateTrackingCategory(ctx, xeroTenantId, trackingCategoryID, trackingCategory)

Allows you to update tracking categories

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**trackingCategoryID** | [**string**](.md)| Unique identifier for a TrackingCategory | 
**trackingCategory** | [**TrackingCategory**](TrackingCategory.md)|  | 

### Return type

[**TrackingCategories**](TrackingCategories.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTrackingOptions

> TrackingOptions UpdateTrackingOptions(ctx, xeroTenantId, trackingCategoryID, trackingOptionID, trackingOption)

Allows you to update options for a specified tracking category

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xeroTenantId** | **string**| Xero identifier for Tenant | 
**trackingCategoryID** | [**string**](.md)| Unique identifier for a TrackingCategory | 
**trackingOptionID** | [**string**](.md)| Unique identifier for a Tracking Option | 
**trackingOption** | [**TrackingOption**](TrackingOption.md)|  | 

### Return type

[**TrackingOptions**](TrackingOptions.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

