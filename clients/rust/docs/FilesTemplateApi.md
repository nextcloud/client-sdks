# \FilesTemplateApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**files_template_create**](FilesTemplateApi.md#files_template_create) | **POST** /ocs/v2.php/apps/files/api/v1/templates/create | Create a template
[**files_template_list**](FilesTemplateApi.md#files_template_list) | **GET** /ocs/v2.php/apps/files/api/v1/templates | List the available templates
[**files_template_path**](FilesTemplateApi.md#files_template_path) | **POST** /ocs/v2.php/apps/files/api/v1/templates/path | Initialize the template directory



## files_template_create

> crate::models::FilesTemplateCreate200Response files_template_create(file_path, ocs_api_request, template_path, template_type)
Create a template

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**file_path** | **String** | Path of the file | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]
**template_path** | Option<**String**> | Name of the template |  |[default to ]
**template_type** | Option<**String**> | Type of the template |  |[default to user]

### Return type

[**crate::models::FilesTemplateCreate200Response**](files_template_create_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## files_template_list

> crate::models::FilesTemplateList200Response files_template_list(ocs_api_request)
List the available templates

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::FilesTemplateList200Response**](files_template_list_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## files_template_path

> crate::models::FilesTemplatePath200Response files_template_path(ocs_api_request, template_path, copy_system_templates)
Initialize the template directory

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**ocs_api_request** | **String** |  | [required] |[default to true]
**template_path** | Option<**String**> | Path of the template directory |  |[default to ]
**copy_system_templates** | Option<**i32**> | Whether to copy the system templates to the template directory |  |[default to 0]

### Return type

[**crate::models::FilesTemplatePath200Response**](files_template_path_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

