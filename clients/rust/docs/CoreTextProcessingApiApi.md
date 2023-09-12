# \CoreTextProcessingApiApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**core_text_processing_api_delete_task**](CoreTextProcessingApiApi.md#core_text_processing_api_delete_task) | **DELETE** /ocs/v2.php/textprocessing/task/{id} | This endpoint allows to delete a scheduled task for a user
[**core_text_processing_api_get_task**](CoreTextProcessingApiApi.md#core_text_processing_api_get_task) | **GET** /ocs/v2.php/textprocessing/task/{id} | This endpoint allows checking the status and results of a task. Tasks are removed 1 week after receiving their last update.
[**core_text_processing_api_list_tasks_by_app**](CoreTextProcessingApiApi.md#core_text_processing_api_list_tasks_by_app) | **GET** /ocs/v2.php/textprocessing/tasks/app/{appId} | This endpoint returns a list of tasks of a user that are related with a specific appId and optionally with an identifier
[**core_text_processing_api_schedule**](CoreTextProcessingApiApi.md#core_text_processing_api_schedule) | **POST** /ocs/v2.php/textprocessing/schedule | This endpoint allows scheduling a language model task
[**core_text_processing_api_task_types**](CoreTextProcessingApiApi.md#core_text_processing_api_task_types) | **GET** /ocs/v2.php/textprocessing/tasktypes | This endpoint returns all available LanguageModel task types



## core_text_processing_api_delete_task

> crate::models::CoreTextProcessingApiSchedule200Response core_text_processing_api_delete_task(id, ocs_api_request)
This endpoint allows to delete a scheduled task for a user

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **i64** | The id of the task | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::CoreTextProcessingApiSchedule200Response**](core_text_processing_api_schedule_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## core_text_processing_api_get_task

> crate::models::CoreTextProcessingApiSchedule200Response core_text_processing_api_get_task(id, ocs_api_request)
This endpoint allows checking the status and results of a task. Tasks are removed 1 week after receiving their last update.

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **i64** | The id of the task | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::CoreTextProcessingApiSchedule200Response**](core_text_processing_api_schedule_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## core_text_processing_api_list_tasks_by_app

> crate::models::CoreTextProcessingApiListTasksByApp200Response core_text_processing_api_list_tasks_by_app(app_id, ocs_api_request, identifier)
This endpoint returns a list of tasks of a user that are related with a specific appId and optionally with an identifier

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**app_id** | **String** | ID of the app | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]
**identifier** | Option<**String**> | An arbitrary identifier for the task |  |

### Return type

[**crate::models::CoreTextProcessingApiListTasksByApp200Response**](core_text_processing_api_list_tasks_by_app_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## core_text_processing_api_schedule

> crate::models::CoreTextProcessingApiSchedule200Response core_text_processing_api_schedule(input, r#type, app_id, ocs_api_request, identifier)
This endpoint allows scheduling a language model task

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**input** | **String** | Input text | [required] |
**r#type** | **String** | Type of the task | [required] |
**app_id** | **String** | ID of the app that will execute the task | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]
**identifier** | Option<**String**> | An arbitrary identifier for the task |  |[default to ]

### Return type

[**crate::models::CoreTextProcessingApiSchedule200Response**](core_text_processing_api_schedule_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## core_text_processing_api_task_types

> crate::models::CoreTextProcessingApiTaskTypes200Response core_text_processing_api_task_types(ocs_api_request)
This endpoint returns all available LanguageModel task types

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::CoreTextProcessingApiTaskTypes200Response**](core_text_processing_api_task_types_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

