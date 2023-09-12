# \DashboardDashboardApiApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**dashboard_dashboard_api_get_widget_items**](DashboardDashboardApiApi.md#dashboard_dashboard_api_get_widget_items) | **GET** /ocs/v2.php/apps/dashboard/api/v1/widget-items | Get the items for the widgets
[**dashboard_dashboard_api_get_widget_items_v2**](DashboardDashboardApiApi.md#dashboard_dashboard_api_get_widget_items_v2) | **GET** /ocs/v2.php/apps/dashboard/api/v2/widget-items | Get the items for the widgets
[**dashboard_dashboard_api_get_widgets**](DashboardDashboardApiApi.md#dashboard_dashboard_api_get_widgets) | **GET** /ocs/v2.php/apps/dashboard/api/v1/widgets | Get the widgets



## dashboard_dashboard_api_get_widget_items

> crate::models::DashboardDashboardApiGetWidgetItems200Response dashboard_dashboard_api_get_widget_items(ocs_api_request, since_ids, limit, widgets_left_square_bracket_right_square_bracket)
Get the items for the widgets

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**ocs_api_request** | **String** |  | [required] |[default to true]
**since_ids** | Option<**String**> | Array indexed by widget Ids, contains date/id from which we want the new items |  |
**limit** | Option<**i64**> | Limit number of result items per widget |  |[default to 7]
**widgets_left_square_bracket_right_square_bracket** | Option<[**Vec<String>**](String.md)> | Limit results to specific widgets |  |[default to []]

### Return type

[**crate::models::DashboardDashboardApiGetWidgetItems200Response**](dashboard_dashboard_api_get_widget_items_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## dashboard_dashboard_api_get_widget_items_v2

> crate::models::DashboardDashboardApiGetWidgetItemsV2200Response dashboard_dashboard_api_get_widget_items_v2(ocs_api_request, since_ids, limit, widgets_left_square_bracket_right_square_bracket)
Get the items for the widgets

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**ocs_api_request** | **String** |  | [required] |[default to true]
**since_ids** | Option<**String**> | Array indexed by widget Ids, contains date/id from which we want the new items |  |
**limit** | Option<**i64**> | Limit number of result items per widget |  |[default to 7]
**widgets_left_square_bracket_right_square_bracket** | Option<[**Vec<String>**](String.md)> | Limit results to specific widgets |  |[default to []]

### Return type

[**crate::models::DashboardDashboardApiGetWidgetItemsV2200Response**](dashboard_dashboard_api_get_widget_items_v2_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## dashboard_dashboard_api_get_widgets

> crate::models::DashboardDashboardApiGetWidgets200Response dashboard_dashboard_api_get_widgets(ocs_api_request)
Get the widgets

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::DashboardDashboardApiGetWidgets200Response**](dashboard_dashboard_api_get_widgets_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

