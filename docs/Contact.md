# Contact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactID** | **string** | Xero identifier | [optional] 
**ContactNumber** | **string** | This can be updated via the API only i.e. This field is read only on the Xero contact screen, used to identify contacts in external systems (max length &#x3D; 50). If the Contact Number is used, this is displayed as Contact Code in the Contacts UI in Xero. | [optional] 
**AccountNumber** | **string** | A user defined account number. This can be updated via the API and the Xero UI (max length &#x3D; 50) | [optional] 
**ContactStatus** | **string** | Current status of a contact – see contact status types | [optional] 
**Name** | **string** | Full name of contact/organisation (max length &#x3D; 255) | [optional] 
**FirstName** | **string** | First name of contact person (max length &#x3D; 255) | [optional] 
**LastName** | **string** | Last name of contact person (max length &#x3D; 255) | [optional] 
**EmailAddress** | **string** | Email address of contact person (umlauts not supported) (max length  &#x3D; 255) | [optional] 
**SkypeUserName** | **string** | Skype user name of contact | [optional] 
**ContactPersons** | [**[]ContactPerson**](ContactPerson.md) | See contact persons | [optional] 
**BankAccountDetails** | **string** | Bank account number of contact | [optional] 
**TaxNumber** | **string** | Tax number of contact – this is also known as the ABN (Australia), GST Number (New Zealand), VAT Number (UK) or Tax ID Number (US and global) in the Xero UI depending on which regionalized version of Xero you are using (max length &#x3D; 50) | [optional] 
**AccountsReceivableTaxType** | **string** | The tax type from TaxRates | [optional] 
**AccountsPayableTaxType** | **string** | The tax type from TaxRates | [optional] 
**Addresses** | [**[]Address**](Address.md) | Store certain address types for a contact – see address types | [optional] 
**Phones** | [**[]Phone**](Phone.md) | Store certain phone types for a contact – see phone types | [optional] 
**IsSupplier** | **bool** | true or false – Boolean that describes if a contact that has any AP  invoices entered against them. Cannot be set via PUT or POST – it is automatically set when an accounts payable invoice is generated against this contact. | [optional] 
**IsCustomer** | **bool** | true or false – Boolean that describes if a contact has any AR invoices entered against them. Cannot be set via PUT or POST – it is automatically set when an accounts receivable invoice is generated against this contact. | [optional] 
**DefaultCurrency** | [**CurrencyCode**](CurrencyCode.md) |  | [optional] 
**XeroNetworkKey** | **string** | Store XeroNetworkKey for contacts. | [optional] 
**SalesDefaultAccountCode** | **string** | The default sales account code for contacts | [optional] 
**PurchasesDefaultAccountCode** | **string** | The default purchases account code for contacts | [optional] 
**SalesTrackingCategories** | [**[]SalesTrackingCategory**](SalesTrackingCategory.md) | The default sales tracking categories for contacts | [optional] 
**PurchasesTrackingCategories** | [**[]SalesTrackingCategory**](SalesTrackingCategory.md) | The default purchases tracking categories for contacts | [optional] 
**TrackingCategoryName** | **string** | The name of the Tracking Category assigned to the contact under SalesTrackingCategories and PurchasesTrackingCategories | [optional] 
**TrackingCategoryOption** | **string** | The name of the Tracking Option assigned to the contact under SalesTrackingCategories and PurchasesTrackingCategories | [optional] 
**PaymentTerms** | [**PaymentTerm**](PaymentTerm.md) |  | [optional] 
**UpdatedDateUTC** | [**time.Time**](time.Time.md) | UTC timestamp of last update to contact | [optional] [readonly] 
**ContactGroups** | [**[]ContactGroup**](ContactGroup.md) | Displays which contact groups a contact is included in | [optional] 
**Website** | **string** | Website address for contact (read only) | [optional] [readonly] 
**BrandingTheme** | [**BrandingTheme**](BrandingTheme.md) |  | [optional] 
**BatchPayments** | [**BatchPaymentDetails**](BatchPaymentDetails.md) |  | [optional] 
**Discount** | **float64** | The default discount rate for the contact (read only) | [optional] [readonly] 
**Balances** | [**Balances**](Balances.md) |  | [optional] 
**Attachments** | [**[]Attachment**](Attachment.md) | Displays array of attachments from the API | [optional] 
**HasAttachments** | **bool** | A boolean to indicate if a contact has an attachment | [optional] [default to false]
**ValidationErrors** | [**[]ValidationError**](ValidationError.md) | Displays validation errors returned from the API | [optional] 
**HasValidationErrors** | **bool** | A boolean to indicate if a contact has an validation errors | [optional] [default to false]
**StatusAttributeString** | **string** | Status of object | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


