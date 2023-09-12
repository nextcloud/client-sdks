# \CoreAppPasswordApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**core_app_password_delete_app_password**](CoreAppPasswordApi.md#core_app_password_delete_app_password) | **DELETE** /ocs/v2.php/core/apppassword | Delete app password
[**core_app_password_get_app_password**](CoreAppPasswordApi.md#core_app_password_get_app_password) | **GET** /ocs/v2.php/core/getapppassword | Create app password
[**core_app_password_rotate_app_password**](CoreAppPasswordApi.md#core_app_password_rotate_app_password) | **POST** /ocs/v2.php/core/apppassword/rotate | Rotate app password



## core_app_password_delete_app_password

> crate::models::CoreWhatsNewDismiss200Response core_app_password_delete_app_password(ocs_api_request)
Delete app password

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::CoreWhatsNewDismiss200Response**](core_whats_new_dismiss_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## core_app_password_get_app_password

> crate::models::CoreAppPasswordGetAppPassword200Response core_app_password_get_app_password(ocs_api_request)
Create app password

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::CoreAppPasswordGetAppPassword200Response**](core_app_password_get_app_password_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## core_app_password_rotate_app_password

> crate::models::CoreAppPasswordGetAppPassword200Response core_app_password_rotate_app_password(ocs_api_request)
Rotate app password

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::CoreAppPasswordGetAppPassword200Response**](core_app_password_get_app_password_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

