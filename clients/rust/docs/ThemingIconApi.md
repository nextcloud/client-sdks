# \ThemingIconApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**theming_icon_get_favicon**](ThemingIconApi.md#theming_icon_get_favicon) | **GET** /index.php/apps/theming/favicon/{app} | Return a 32x32 favicon as png
[**theming_icon_get_themed_icon**](ThemingIconApi.md#theming_icon_get_themed_icon) | **GET** /index.php/apps/theming/img/{app}/{image} | Get a themed icon
[**theming_icon_get_touch_icon**](ThemingIconApi.md#theming_icon_get_touch_icon) | **GET** /index.php/apps/theming/icon/{app} | Return a 512x512 icon for touch devices



## theming_icon_get_favicon

> std::path::PathBuf theming_icon_get_favicon(app)
Return a 32x32 favicon as png

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**app** | **String** | ID of the app | [required] |[default to core]

### Return type

[**std::path::PathBuf**](std::path::PathBuf.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/x-icon

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## theming_icon_get_themed_icon

> std::path::PathBuf theming_icon_get_themed_icon(app, image)
Get a themed icon

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**app** | **String** | ID of the app | [required] |
**image** | **String** | image file name (svg required) | [required] |

### Return type

[**std::path::PathBuf**](std::path::PathBuf.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/svg+xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## theming_icon_get_touch_icon

> std::path::PathBuf theming_icon_get_touch_icon(app)
Return a 512x512 icon for touch devices

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**app** | **String** | ID of the app | [required] |[default to core]

### Return type

[**std::path::PathBuf**](std::path::PathBuf.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/png, image/x-icon

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

