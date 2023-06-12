# \ThemingUserThemeApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**theming_user_theme_delete_background**](ThemingUserThemeApi.md#theming_user_theme_delete_background) | **DELETE** /index.php/apps/theming/background/custom | Delete the background
[**theming_user_theme_disable_theme**](ThemingUserThemeApi.md#theming_user_theme_disable_theme) | **DELETE** /ocs/v2.php/apps/theming/api/v1/theme/{themeId} | Disable theme
[**theming_user_theme_enable_theme**](ThemingUserThemeApi.md#theming_user_theme_enable_theme) | **PUT** /ocs/v2.php/apps/theming/api/v1/theme/{themeId}/enable | Enable theme
[**theming_user_theme_get_background**](ThemingUserThemeApi.md#theming_user_theme_get_background) | **GET** /index.php/apps/theming/background | Get the background image
[**theming_user_theme_set_background**](ThemingUserThemeApi.md#theming_user_theme_set_background) | **POST** /index.php/apps/theming/background/{type} | Set the background



## theming_user_theme_delete_background

> crate::models::ThemingBackground theming_user_theme_delete_background(ocs_api_request)
Delete the background

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::ThemingBackground**](ThemingBackground.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## theming_user_theme_disable_theme

> crate::models::CoreWhatsNewDismiss200Response theming_user_theme_disable_theme(theme_id, ocs_api_request)
Disable theme

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**theme_id** | **String** | the theme ID | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::CoreWhatsNewDismiss200Response**](core_whats_new_dismiss_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## theming_user_theme_enable_theme

> crate::models::CoreWhatsNewDismiss200Response theming_user_theme_enable_theme(theme_id, ocs_api_request)
Enable theme

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**theme_id** | **String** | the theme ID | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::CoreWhatsNewDismiss200Response**](core_whats_new_dismiss_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## theming_user_theme_get_background

> std::path::PathBuf theming_user_theme_get_background(ocs_api_request)
Get the background image

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**std::path::PathBuf**](std::path::PathBuf.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## theming_user_theme_set_background

> crate::models::ThemingBackground theming_user_theme_set_background(r#type, ocs_api_request, value, color)
Set the background

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**r#type** | **String** | Type of background | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]
**value** | Option<**String**> | Path of the background image |  |[default to ]
**color** | Option<**String**> | Color for the background |  |

### Return type

[**crate::models::ThemingBackground**](ThemingBackground.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

