# \FilesSharingPublicPreviewApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**files_sharing_public_preview_direct_link**](FilesSharingPublicPreviewApi.md#files_sharing_public_preview_direct_link) | **GET** /index.php/s/{token}/preview | Get a direct link preview for a shared file
[**files_sharing_public_preview_get_preview**](FilesSharingPublicPreviewApi.md#files_sharing_public_preview_get_preview) | **GET** /index.php/apps/files_sharing/publicpreview/{token} | Get a preview for a shared file



## files_sharing_public_preview_direct_link

> std::path::PathBuf files_sharing_public_preview_direct_link(token, ocs_api_request)
Get a direct link preview for a shared file

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**token** | **String** | Token of the share | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**std::path::PathBuf**](std::path::PathBuf.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## files_sharing_public_preview_get_preview

> std::path::PathBuf files_sharing_public_preview_get_preview(token, ocs_api_request, file, x, y, a)
Get a preview for a shared file

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**token** | **String** | Token of the share | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]
**file** | Option<**String**> | File in the share |  |[default to ]
**x** | Option<**i64**> | Width of the preview |  |[default to 32]
**y** | Option<**i64**> | Height of the preview |  |[default to 32]
**a** | Option<**i32**> | Whether to not crop the preview |  |[default to 0]

### Return type

[**std::path::PathBuf**](std::path::PathBuf.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

