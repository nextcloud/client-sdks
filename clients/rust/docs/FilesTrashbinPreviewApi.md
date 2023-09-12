# \FilesTrashbinPreviewApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**files_trashbin_preview_get_preview**](FilesTrashbinPreviewApi.md#files_trashbin_preview_get_preview) | **GET** /index.php/apps/files_trashbin/preview | Get the preview for a file



## files_trashbin_preview_get_preview

> std::path::PathBuf files_trashbin_preview_get_preview(file_id, x, y, a)
Get the preview for a file

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**file_id** | Option<**i64**> | ID of the file |  |[default to -1]
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

