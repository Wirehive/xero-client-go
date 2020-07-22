# LineItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LineItemID** | **string** | LineItem unique ID | [optional] 
**Description** | **string** | Description needs to be at least 1 char long. A line item with just a description (i.e no unit amount or quantity) can be created by specifying just a &lt;Description&gt; element that contains at least 1 character | [optional] 
**Quantity** | **float64** | LineItem Quantity | [optional] 
**UnitAmount** | **float64** | LineItem Unit Amount | [optional] 
**ItemCode** | **string** | See Items | [optional] 
**AccountCode** | **string** | See Accounts | [optional] 
**TaxType** | **string** | The tax type from TaxRates | [optional] 
**TaxAmount** | **float64** | The tax amount is auto calculated as a percentage of the line amount (see below) based on the tax rate. This value can be overriden if the calculated &lt;TaxAmount&gt; is not correct. | [optional] 
**LineAmount** | **float64** | If you wish to omit either of the &lt;Quantity&gt; or &lt;UnitAmount&gt; you can provide a LineAmount and Xero will calculate the missing amount for you. The line amount reflects the discounted price if a DiscountRate has been used . i.e LineAmount &#x3D; Quantity * Unit Amount * ((100 – DiscountRate)/100) | [optional] 
**Tracking** | [**[]LineItemTracking**](LineItemTracking.md) | Optional Tracking Category – see Tracking.  Any LineItem can have a  maximum of 2 &lt;TrackingCategory&gt; elements. | [optional] 
**DiscountRate** | **float64** | Percentage discount being applied to a line item (only supported on  ACCREC invoices – ACC PAY invoices and credit notes in Xero do not support discounts | [optional] 
**DiscountAmount** | **float64** | Discount amount being applied to a line item. Only supported on ACCREC invoices - ACCPAY invoices and credit notes in Xero do not support discounts. | [optional] 
**RepeatingInvoiceID** | **string** | The Xero identifier for a Repeating Invoice | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


