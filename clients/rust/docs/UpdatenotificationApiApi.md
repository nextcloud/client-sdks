# \UpdatenotificationApiApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**updatenotification_api_get_app_list**](UpdatenotificationApiApi.md#updatenotification_api_get_app_list) | **GET** /ocs/v2.php/apps/updatenotification/api/{apiVersion}/applist/{newVersion} | List available updates for apps



## updatenotification_api_get_app_list

> crate::models::UpdatenotificationApiGetAppList200Response updatenotification_api_get_app_list(api_version, new_version, ocs_api_request)
List available updates for apps

This endpoint requires admin access

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**api_version** | **String** |  | [required] |[default to v1]
**new_version** | **String** | Server version to check updates for | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::UpdatenotificationApiGetAppList200Response**](updatenotification_api_get_app_list_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

