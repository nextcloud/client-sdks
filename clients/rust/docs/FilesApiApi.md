# \FilesApiApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**files_api_get_thumbnail**](FilesApiApi.md#files_api_get_thumbnail) | **GET** /index.php/apps/files/api/v1/thumbnail/{x}/{y}/{file} | Gets a thumbnail of the specified file
[**files_api_service_worker**](FilesApiApi.md#files_api_service_worker) | **GET** /index.php/apps/files/preview-service-worker.js | Get the service-worker Javascript for previews



## files_api_get_thumbnail

> std::path::PathBuf files_api_get_thumbnail(x, y, file)
Gets a thumbnail of the specified file

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**x** | **i64** | Width of the thumbnail | [required] |
**y** | **i64** | Height of the thumbnail | [required] |
**file** | **String** | URL-encoded filename | [required] |

### Return type

[**std::path::PathBuf**](std::path::PathBuf.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## files_api_service_worker

> std::path::PathBuf files_api_service_worker()
Get the service-worker Javascript for previews

### Parameters

This endpoint does not need any parameter.

### Return type

[**std::path::PathBuf**](std::path::PathBuf.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

