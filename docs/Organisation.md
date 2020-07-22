# Organisation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganisationID** | **string** | Unique Xero identifier | [optional] 
**APIKey** | **string** | Display a unique key used for Xero-to-Xero transactions | [optional] 
**Name** | **string** | Display name of organisation shown in Xero | [optional] 
**LegalName** | **string** | Organisation name shown on Reports | [optional] 
**PaysTax** | **bool** | Boolean to describe if organisation is registered with a local tax authority i.e. true, false | [optional] 
**Version** | **string** | See Version Types | [optional] 
**OrganisationType** | **string** | Organisation Type | [optional] 
**BaseCurrency** | [**CurrencyCode**](CurrencyCode.md) |  | [optional] 
**CountryCode** | [**CountryCode**](CountryCode.md) |  | [optional] 
**IsDemoCompany** | **bool** | Boolean to describe if organisation is a demo company. | [optional] 
**OrganisationStatus** | **string** | Will be set to ACTIVE if you can connect to organisation via the Xero API | [optional] 
**RegistrationNumber** | **string** | Shows for New Zealand, Australian and UK organisations | [optional] 
**EmployerIdentificationNumber** | **string** | Shown if set. US Only. | [optional] 
**TaxNumber** | **string** | Shown if set. Displays in the Xero UI as Tax File Number (AU), GST Number (NZ), VAT Number (UK) and Tax ID Number (US &amp; Global). | [optional] 
**FinancialYearEndDay** | **int32** | Calendar day e.g. 0-31 | [optional] 
**FinancialYearEndMonth** | **int32** | Calendar Month e.g. 1-12 | [optional] 
**SalesTaxBasis** | **string** | The accounting basis used for tax returns. See Sales Tax Basis | [optional] 
**SalesTaxPeriod** | **string** | The frequency with which tax returns are processed. See Sales Tax Period | [optional] 
**DefaultSalesTax** | **string** | The default for LineAmountTypes on sales transactions | [optional] 
**DefaultPurchasesTax** | **string** | The default for LineAmountTypes on purchase transactions | [optional] 
**PeriodLockDate** | **string** | Shown if set. See lock dates | [optional] 
**EndOfYearLockDate** | **string** | Shown if set. See lock dates | [optional] 
**CreatedDateUTC** | [**time.Time**](time.Time.md) | Timestamp when the organisation was created in Xero | [optional] [readonly] 
**Timezone** | [**TimeZone**](TimeZone.md) |  | [optional] 
**OrganisationEntityType** | **string** | Organisation Type | [optional] 
**ShortCode** | **string** | A unique identifier for the organisation. Potential uses. | [optional] 
**Class** | **string** | Organisation Classes describe which plan the Xero organisation is on (e.g. DEMO, TRIAL, PREMIUM) | [optional] 
**Edition** | **string** | BUSINESS or PARTNER. Partner edition organisations are sold exclusively through accounting partners and have restricted functionality (e.g. no access to invoicing) | [optional] 
**LineOfBusiness** | **string** | Description of business type as defined in Organisation settings | [optional] 
**Addresses** | [**[]Address**](Address.md) | Address details for organisation – see Addresses | [optional] 
**Phones** | [**[]Phone**](Phone.md) | Phones details for organisation – see Phones | [optional] 
**ExternalLinks** | [**[]ExternalLink**](ExternalLink.md) | Organisation profile links for popular services such as Facebook,Twitter, GooglePlus and LinkedIn. You can also add link to your website here. Shown if Organisation settings  is updated in Xero. See ExternalLinks below | [optional] 
**PaymentTerms** | [**PaymentTerm**](PaymentTerm.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


