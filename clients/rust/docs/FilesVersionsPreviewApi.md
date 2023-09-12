# \FilesVersionsPreviewApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**files_versions_preview_get_preview**](FilesVersionsPreviewApi.md#files_versions_preview_get_preview) | **GET** /index.php/apps/files_versions/preview | Get the preview for a file version



## files_versions_preview_get_preview

> std::path::PathBuf files_versions_preview_get_preview(file, x, y, version)
Get the preview for a file version

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**file** | Option<**String**> | Path of the file |  |[default to ]
**x** | Option<**i64**> | Width of the preview |  |[default to 44]
**y** | Option<**i64**> | Height of the preview |  |[default to 44]
**version** | Option<**String**> | Version of the file to get the preview for |  |[default to ]

### Return type

[**std::path::PathBuf**](std::path::PathBuf.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

