# \ProvisioningApiPreferencesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**provisioning_api_preferences_delete_multiple_preference**](ProvisioningApiPreferencesApi.md#provisioning_api_preferences_delete_multiple_preference) | **DELETE** /ocs/v2.php/apps/provisioning_api/api/v1/config/users/{appId} | Delete multiple preferences for an app
[**provisioning_api_preferences_delete_preference**](ProvisioningApiPreferencesApi.md#provisioning_api_preferences_delete_preference) | **DELETE** /ocs/v2.php/apps/provisioning_api/api/v1/config/users/{appId}/{configKey} | Delete a preference for an app
[**provisioning_api_preferences_set_multiple_preferences**](ProvisioningApiPreferencesApi.md#provisioning_api_preferences_set_multiple_preferences) | **POST** /ocs/v2.php/apps/provisioning_api/api/v1/config/users/{appId} | Update multiple preference values of an app
[**provisioning_api_preferences_set_preference**](ProvisioningApiPreferencesApi.md#provisioning_api_preferences_set_preference) | **POST** /ocs/v2.php/apps/provisioning_api/api/v1/config/users/{appId}/{configKey} | Update a preference value of an app



## provisioning_api_preferences_delete_multiple_preference

> crate::models::CoreWhatsNewDismiss200Response provisioning_api_preferences_delete_multiple_preference(config_keys_left_square_bracket_right_square_bracket, app_id, ocs_api_request)
Delete multiple preferences for an app

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**config_keys_left_square_bracket_right_square_bracket** | [**Vec<String>**](String.md) | Keys to delete | [required] |
**app_id** | **String** | ID of the app | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::CoreWhatsNewDismiss200Response**](core_whats_new_dismiss_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## provisioning_api_preferences_delete_preference

> crate::models::CoreWhatsNewDismiss200Response provisioning_api_preferences_delete_preference(app_id, config_key, ocs_api_request)
Delete a preference for an app

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**app_id** | **String** | ID of the app | [required] |
**config_key** | **String** | Key to delete | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::CoreWhatsNewDismiss200Response**](core_whats_new_dismiss_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## provisioning_api_preferences_set_multiple_preferences

> crate::models::CoreWhatsNewDismiss200Response provisioning_api_preferences_set_multiple_preferences(configs, app_id, ocs_api_request)
Update multiple preference values of an app

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**configs** | **String** | Key-value pairs of the preferences | [required] |
**app_id** | **String** | ID of the app | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::CoreWhatsNewDismiss200Response**](core_whats_new_dismiss_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## provisioning_api_preferences_set_preference

> crate::models::CoreWhatsNewDismiss200Response provisioning_api_preferences_set_preference(config_value, app_id, config_key, ocs_api_request)
Update a preference value of an app

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**config_value** | **String** | New value | [required] |
**app_id** | **String** | ID of the app | [required] |
**config_key** | **String** | Key of the preference | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::CoreWhatsNewDismiss200Response**](core_whats_new_dismiss_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

