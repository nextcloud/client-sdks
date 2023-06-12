# \CoreAvatarApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**core_avatar_get_avatar**](CoreAvatarApi.md#core_avatar_get_avatar) | **GET** /index.php/avatar/{userId}/{size} | Get the avatar
[**core_avatar_get_avatar_dark**](CoreAvatarApi.md#core_avatar_get_avatar_dark) | **GET** /index.php/avatar/{userId}/{size}/dark | Get the dark avatar



## core_avatar_get_avatar

> std::path::PathBuf core_avatar_get_avatar(user_id, size)
Get the avatar

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**user_id** | **String** | ID of the user | [required] |
**size** | **i64** | Size of the avatar | [required] |

### Return type

[**std::path::PathBuf**](std::path::PathBuf.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## core_avatar_get_avatar_dark

> std::path::PathBuf core_avatar_get_avatar_dark(user_id, size)
Get the dark avatar

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**user_id** | **String** | ID of the user | [required] |
**size** | **i64** | Size of the avatar | [required] |

### Return type

[**std::path::PathBuf**](std::path::PathBuf.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

