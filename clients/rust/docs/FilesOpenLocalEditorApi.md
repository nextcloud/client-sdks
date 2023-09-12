# \FilesOpenLocalEditorApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**files_open_local_editor_create**](FilesOpenLocalEditorApi.md#files_open_local_editor_create) | **POST** /ocs/v2.php/apps/files/api/v1/openlocaleditor | Create a local editor
[**files_open_local_editor_validate**](FilesOpenLocalEditorApi.md#files_open_local_editor_validate) | **POST** /ocs/v2.php/apps/files/api/v1/openlocaleditor/{token} | Validate a local editor



## files_open_local_editor_create

> crate::models::FilesOpenLocalEditorCreate200Response files_open_local_editor_create(path, ocs_api_request)
Create a local editor

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**path** | **String** | Path of the file | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::FilesOpenLocalEditorCreate200Response**](files_open_local_editor_create_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## files_open_local_editor_validate

> crate::models::FilesOpenLocalEditorValidate200Response files_open_local_editor_validate(path, token, ocs_api_request)
Validate a local editor

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**path** | **String** | Path of the file | [required] |
**token** | **String** | Token of the local editor | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::FilesOpenLocalEditorValidate200Response**](files_open_local_editor_validate_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

