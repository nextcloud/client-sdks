# \FilesRemindersApiApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**files_reminders_api_get**](FilesRemindersApiApi.md#files_reminders_api_get) | **GET** /ocs/v2.php/apps/files_reminders/api/v{version}/{fileId} | Get a reminder
[**files_reminders_api_remove**](FilesRemindersApiApi.md#files_reminders_api_remove) | **DELETE** /ocs/v2.php/apps/files_reminders/api/v{version}/{fileId} | Remove a reminder
[**files_reminders_api_set**](FilesRemindersApiApi.md#files_reminders_api_set) | **PUT** /ocs/v2.php/apps/files_reminders/api/v{version}/{fileId} | Set a reminder



## files_reminders_api_get

> crate::models::FilesRemindersApiGet200Response files_reminders_api_get(version, file_id, ocs_api_request)
Get a reminder

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**version** | **String** |  | [required] |
**file_id** | **i64** | ID of the file | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::FilesRemindersApiGet200Response**](files_reminders_api_get_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## files_reminders_api_remove

> crate::models::CoreWhatsNewDismiss200Response files_reminders_api_remove(version, file_id, ocs_api_request)
Remove a reminder

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**version** | **String** |  | [required] |
**file_id** | **i64** | ID of the file | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::CoreWhatsNewDismiss200Response**](core_whats_new_dismiss_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## files_reminders_api_set

> crate::models::CoreWhatsNewDismiss200Response files_reminders_api_set(due_date, version, file_id, ocs_api_request)
Set a reminder

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**due_date** | **String** | ISO 8601 formatted date time string | [required] |
**version** | **String** |  | [required] |
**file_id** | **i64** | ID of the file | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::CoreWhatsNewDismiss200Response**](core_whats_new_dismiss_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

