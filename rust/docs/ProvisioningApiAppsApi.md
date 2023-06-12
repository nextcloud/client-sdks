# \ProvisioningApiAppsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**provisioning_api_apps_disable**](ProvisioningApiAppsApi.md#provisioning_api_apps_disable) | **DELETE** /ocs/v2.php/cloud/apps/{app} | Disable an app
[**provisioning_api_apps_enable**](ProvisioningApiAppsApi.md#provisioning_api_apps_enable) | **POST** /ocs/v2.php/cloud/apps/{app} | Enable an app
[**provisioning_api_apps_get_app_info**](ProvisioningApiAppsApi.md#provisioning_api_apps_get_app_info) | **GET** /ocs/v2.php/cloud/apps/{app} | Get the app info for an app
[**provisioning_api_apps_get_apps**](ProvisioningApiAppsApi.md#provisioning_api_apps_get_apps) | **GET** /ocs/v2.php/cloud/apps | Get a list of installed apps



## provisioning_api_apps_disable

> crate::models::CoreWhatsNewDismiss200Response provisioning_api_apps_disable(app, ocs_api_request)
Disable an app

This endpoint requires admin access

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**app** | **String** | ID of the app | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::CoreWhatsNewDismiss200Response**](core_whats_new_dismiss_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## provisioning_api_apps_enable

> crate::models::CoreWhatsNewDismiss200Response provisioning_api_apps_enable(app, ocs_api_request)
Enable an app

This endpoint requires admin access

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**app** | **String** | ID of the app | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::CoreWhatsNewDismiss200Response**](core_whats_new_dismiss_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## provisioning_api_apps_get_app_info

> crate::models::ProvisioningApiAppsGetAppInfo200Response provisioning_api_apps_get_app_info(app, ocs_api_request)
Get the app info for an app

This endpoint requires admin access

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**app** | **String** | ID of the app | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::ProvisioningApiAppsGetAppInfo200Response**](provisioning_api_apps_get_app_info_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## provisioning_api_apps_get_apps

> crate::models::ProvisioningApiAppsGetApps200Response provisioning_api_apps_get_apps(ocs_api_request, filter)
Get a list of installed apps

This endpoint requires admin access

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**ocs_api_request** | **String** |  | [required] |[default to true]
**filter** | Option<**String**> | Filter for enabled or disabled apps |  |

### Return type

[**crate::models::ProvisioningApiAppsGetApps200Response**](provisioning_api_apps_get_apps_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

