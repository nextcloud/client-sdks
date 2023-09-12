# \FilesDirectEditingApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**files_direct_editing_create**](FilesDirectEditingApi.md#files_direct_editing_create) | **POST** /ocs/v2.php/apps/files/api/v1/directEditing/create | Create a file for direct editing
[**files_direct_editing_info**](FilesDirectEditingApi.md#files_direct_editing_info) | **GET** /ocs/v2.php/apps/files/api/v1/directEditing | Get the direct editing capabilities
[**files_direct_editing_open**](FilesDirectEditingApi.md#files_direct_editing_open) | **POST** /ocs/v2.php/apps/files/api/v1/directEditing/open | Open a file for direct editing
[**files_direct_editing_templates**](FilesDirectEditingApi.md#files_direct_editing_templates) | **GET** /ocs/v2.php/apps/files/api/v1/directEditing/templates/{editorId}/{creatorId} | Get the templates for direct editing



## files_direct_editing_create

> crate::models::DavDirectGetUrl200Response files_direct_editing_create(path, editor_id, creator_id, ocs_api_request, template_id)
Create a file for direct editing

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**path** | **String** | Path of the file | [required] |
**editor_id** | **String** | ID of the editor | [required] |
**creator_id** | **String** | ID of the creator | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]
**template_id** | Option<**String**> | ID of the template |  |

### Return type

[**crate::models::DavDirectGetUrl200Response**](dav_direct_get_url_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## files_direct_editing_info

> crate::models::FilesDirectEditingInfo200Response files_direct_editing_info(ocs_api_request)
Get the direct editing capabilities

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::FilesDirectEditingInfo200Response**](files_direct_editing_info_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## files_direct_editing_open

> crate::models::DavDirectGetUrl200Response files_direct_editing_open(path, ocs_api_request, editor_id, file_id)
Open a file for direct editing

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**path** | **String** | Path of the file | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]
**editor_id** | Option<**String**> | ID of the editor |  |
**file_id** | Option<**i64**> | ID of the file |  |

### Return type

[**crate::models::DavDirectGetUrl200Response**](dav_direct_get_url_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## files_direct_editing_templates

> crate::models::FilesDirectEditingTemplates200Response files_direct_editing_templates(editor_id, creator_id, ocs_api_request)
Get the templates for direct editing

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**editor_id** | **String** | ID of the editor | [required] |
**creator_id** | **String** | ID of the creator | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::FilesDirectEditingTemplates200Response**](files_direct_editing_templates_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

