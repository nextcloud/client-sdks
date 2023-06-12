# \ProvisioningApiAppConfigApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**provisioning_api_app_config_delete_key**](ProvisioningApiAppConfigApi.md#provisioning_api_app_config_delete_key) | **DELETE** /ocs/v2.php/apps/provisioning_api/api/v1/config/apps/{app}/{key} | Delete a config key of an app
[**provisioning_api_app_config_get_apps**](ProvisioningApiAppConfigApi.md#provisioning_api_app_config_get_apps) | **GET** /ocs/v2.php/apps/provisioning_api/api/v1/config/apps | Get a list of apps
[**provisioning_api_app_config_get_keys**](ProvisioningApiAppConfigApi.md#provisioning_api_app_config_get_keys) | **GET** /ocs/v2.php/apps/provisioning_api/api/v1/config/apps/{app} | Get the config keys of an app
[**provisioning_api_app_config_get_value**](ProvisioningApiAppConfigApi.md#provisioning_api_app_config_get_value) | **GET** /ocs/v2.php/apps/provisioning_api/api/v1/config/apps/{app}/{key} | Get a the config value of an app
[**provisioning_api_app_config_set_value**](ProvisioningApiAppConfigApi.md#provisioning_api_app_config_set_value) | **POST** /ocs/v2.php/apps/provisioning_api/api/v1/config/apps/{app}/{key} | Update the config value of an app



## provisioning_api_app_config_delete_key

> crate::models::CoreWhatsNewDismiss200Response provisioning_api_app_config_delete_key(app, key, ocs_api_request)
Delete a config key of an app

This endpoint requires admin access

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**app** | **String** | ID of the app | [required] |
**key** | **String** | Key to delete | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::CoreWhatsNewDismiss200Response**](core_whats_new_dismiss_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## provisioning_api_app_config_get_apps

> crate::models::ProvisioningApiAppConfigGetApps200Response provisioning_api_app_config_get_apps(ocs_api_request)
Get a list of apps

This endpoint requires admin access

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::ProvisioningApiAppConfigGetApps200Response**](provisioning_api_app_config_get_apps_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## provisioning_api_app_config_get_keys

> crate::models::ProvisioningApiAppConfigGetApps200Response provisioning_api_app_config_get_keys(app, ocs_api_request)
Get the config keys of an app

This endpoint requires admin access

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**app** | **String** | ID of the app | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::ProvisioningApiAppConfigGetApps200Response**](provisioning_api_app_config_get_apps_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## provisioning_api_app_config_get_value

> crate::models::ProvisioningApiAppConfigGetValue200Response provisioning_api_app_config_get_value(app, key, ocs_api_request, default_value)
Get a the config value of an app

This endpoint requires admin access

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**app** | **String** | ID if the app | [required] |
**key** | **String** | Key | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]
**default_value** | Option<**String**> | Default returned value if the value is empty |  |[default to ]

### Return type

[**crate::models::ProvisioningApiAppConfigGetValue200Response**](provisioning_api_app_config_get_value_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## provisioning_api_app_config_set_value

> crate::models::CoreWhatsNewDismiss200Response provisioning_api_app_config_set_value(value, app, key, ocs_api_request)
Update the config value of an app

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**value** | **String** | New value for the key | [required] |
**app** | **String** | ID of the app | [required] |
**key** | **String** | Key to update | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::CoreWhatsNewDismiss200Response**](core_whats_new_dismiss_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

