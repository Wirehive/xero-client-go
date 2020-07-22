# ReportWithRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportID** | **string** | Report id | [optional] 
**ReportName** | **string** | Name of the report | [optional] 
**ReportTitle** | **string** | Title of the report | [optional] 
**ReportType** | **string** | The type of report (BalanceSheet,ProfitLoss, etc) | [optional] 
**ReportTitles** | **[]string** | Report titles array (3 to 4 strings with the report name, orgnisation name and time frame of report) | [optional] 
**ReportDate** | **string** | Date of report | [optional] 
**Rows** | [**[]ReportRows**](ReportRows.md) |  | [optional] 
**UpdatedDateUTC** | [**time.Time**](time.Time.md) | Updated Date | [optional] [readonly] 
**Fields** | [**[]ReportFields**](ReportFields.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


