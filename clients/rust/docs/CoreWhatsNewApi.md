# \CoreWhatsNewApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**core_whats_new_dismiss**](CoreWhatsNewApi.md#core_whats_new_dismiss) | **POST** /ocs/v2.php/core/whatsnew | Dismiss the changes
[**core_whats_new_get**](CoreWhatsNewApi.md#core_whats_new_get) | **GET** /ocs/v2.php/core/whatsnew | Get the changes



## core_whats_new_dismiss

> crate::models::CoreWhatsNewDismiss200Response core_whats_new_dismiss(version, ocs_api_request)
Dismiss the changes

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**version** | **String** | Version to dismiss the changes for | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::CoreWhatsNewDismiss200Response**](core_whats_new_dismiss_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## core_whats_new_get

> crate::models::CoreWhatsNewGet200Response core_whats_new_get(ocs_api_request)
Get the changes

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::CoreWhatsNewGet200Response**](core_whats_new_get_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

