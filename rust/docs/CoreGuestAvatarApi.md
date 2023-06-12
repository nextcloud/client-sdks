# \CoreGuestAvatarApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**core_guest_avatar_get_avatar**](CoreGuestAvatarApi.md#core_guest_avatar_get_avatar) | **GET** /index.php/avatar/guest/{guestName}/{size} | Returns a guest avatar image response
[**core_guest_avatar_get_avatar_dark**](CoreGuestAvatarApi.md#core_guest_avatar_get_avatar_dark) | **GET** /index.php/avatar/guest/{guestName}/{size}/dark | Returns a dark guest avatar image response



## core_guest_avatar_get_avatar

> std::path::PathBuf core_guest_avatar_get_avatar(guest_name, size, dark_theme)
Returns a guest avatar image response

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**guest_name** | **String** | The guest name, e.g. \"Albert\" | [required] |
**size** | **String** | The desired avatar size, e.g. 64 for 64x64px | [required] |
**dark_theme** | Option<**i32**> | Return dark avatar |  |[default to 0]

### Return type

[**std::path::PathBuf**](std::path::PathBuf.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## core_guest_avatar_get_avatar_dark

> std::path::PathBuf core_guest_avatar_get_avatar_dark(guest_name, size)
Returns a dark guest avatar image response

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**guest_name** | **String** | The guest name, e.g. \"Albert\" | [required] |
**size** | **String** | The desired avatar size, e.g. 64 for 64x64px | [required] |

### Return type

[**std::path::PathBuf**](std::path::PathBuf.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

