# \ThemingThemingApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**theming_theming_get_image**](ThemingThemingApi.md#theming_theming_get_image) | **GET** /index.php/apps/theming/image/{key} | Get an image
[**theming_theming_get_manifest**](ThemingThemingApi.md#theming_theming_get_manifest) | **GET** /index.php/apps/theming/manifest/{app} | Get the manifest for an app
[**theming_theming_get_theme_stylesheet**](ThemingThemingApi.md#theming_theming_get_theme_stylesheet) | **GET** /index.php/apps/theming/theme/{themeId}.css | Get the CSS stylesheet for a theme



## theming_theming_get_image

> std::path::PathBuf theming_theming_get_image(key, use_svg)
Get an image

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**key** | **String** | Key of the image | [required] |
**use_svg** | Option<**i32**> | Return image as SVG |  |[default to 1]

### Return type

[**std::path::PathBuf**](std::path::PathBuf.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## theming_theming_get_manifest

> crate::models::ThemingThemingGetManifest200Response theming_theming_get_manifest(app)
Get the manifest for an app

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**app** | **String** | ID of the app | [required] |

### Return type

[**crate::models::ThemingThemingGetManifest200Response**](theming_theming_get_manifest_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## theming_theming_get_theme_stylesheet

> std::path::PathBuf theming_theming_get_theme_stylesheet(theme_id, plain, with_custom_css)
Get the CSS stylesheet for a theme

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**theme_id** | **String** | ID of the theme | [required] |
**plain** | Option<**i32**> | Let the browser decide the CSS priority |  |[default to 0]
**with_custom_css** | Option<**i32**> | Include custom CSS |  |[default to 0]

### Return type

[**std::path::PathBuf**](std::path::PathBuf.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/css

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

