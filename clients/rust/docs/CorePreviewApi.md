# \CorePreviewApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**core_preview_get_preview**](CorePreviewApi.md#core_preview_get_preview) | **GET** /index.php/core/preview.png | Get a preview by file path
[**core_preview_get_preview_by_file_id**](CorePreviewApi.md#core_preview_get_preview_by_file_id) | **GET** /index.php/core/preview | Get a preview by file ID



## core_preview_get_preview

> std::path::PathBuf core_preview_get_preview(file, x, y, a, force_icon, mode, mime_fallback)
Get a preview by file path

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**file** | Option<**String**> | Path of the file |  |[default to ]
**x** | Option<**i64**> | Width of the preview |  |[default to 32]
**y** | Option<**i64**> | Height of the preview |  |[default to 32]
**a** | Option<**i32**> | Whether to not crop the preview |  |[default to 0]
**force_icon** | Option<**i32**> | Force returning an icon |  |[default to 1]
**mode** | Option<**String**> | How to crop the image |  |[default to fill]
**mime_fallback** | Option<**i32**> | Whether to fallback to the mime icon if no preview is available |  |[default to 0]

### Return type

[**std::path::PathBuf**](std::path::PathBuf.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## core_preview_get_preview_by_file_id

> std::path::PathBuf core_preview_get_preview_by_file_id(file_id, x, y, a, force_icon, mode, mime_fallback)
Get a preview by file ID

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**file_id** | Option<**i64**> | ID of the file |  |[default to -1]
**x** | Option<**i64**> | Width of the preview |  |[default to 32]
**y** | Option<**i64**> | Height of the preview |  |[default to 32]
**a** | Option<**i32**> | Whether to not crop the preview |  |[default to 0]
**force_icon** | Option<**i32**> | Force returning an icon |  |[default to 1]
**mode** | Option<**String**> | How to crop the image |  |[default to fill]
**mime_fallback** | Option<**i32**> | Whether to fallback to the mime icon if no preview is available |  |[default to 0]

### Return type

[**std::path::PathBuf**](std::path::PathBuf.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

