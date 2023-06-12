# \UserStatusUserStatusApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**user_status_user_status_clear_message**](UserStatusUserStatusApi.md#user_status_user_status_clear_message) | **DELETE** /ocs/v2.php/apps/user_status/api/v1/user_status/message | Clear the message of the current user
[**user_status_user_status_get_status**](UserStatusUserStatusApi.md#user_status_user_status_get_status) | **GET** /ocs/v2.php/apps/user_status/api/v1/user_status | Get the status of the current user
[**user_status_user_status_revert_status**](UserStatusUserStatusApi.md#user_status_user_status_revert_status) | **DELETE** /ocs/v2.php/apps/user_status/api/v1/user_status/revert/{messageId} | Revert the status to the previous status
[**user_status_user_status_set_custom_message**](UserStatusUserStatusApi.md#user_status_user_status_set_custom_message) | **PUT** /ocs/v2.php/apps/user_status/api/v1/user_status/message/custom | Set the message to a custom message for the current user
[**user_status_user_status_set_predefined_message**](UserStatusUserStatusApi.md#user_status_user_status_set_predefined_message) | **PUT** /ocs/v2.php/apps/user_status/api/v1/user_status/message/predefined | Set the message to a predefined message for the current user
[**user_status_user_status_set_status**](UserStatusUserStatusApi.md#user_status_user_status_set_status) | **PUT** /ocs/v2.php/apps/user_status/api/v1/user_status/status | Update the status type of the current user



## user_status_user_status_clear_message

> crate::models::CoreWhatsNewDismiss200Response user_status_user_status_clear_message(ocs_api_request)
Clear the message of the current user

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::CoreWhatsNewDismiss200Response**](core_whats_new_dismiss_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## user_status_user_status_get_status

> crate::models::UserStatusUserStatusGetStatus200Response user_status_user_status_get_status(ocs_api_request)
Get the status of the current user

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::UserStatusUserStatusGetStatus200Response**](user_status_user_status_get_status_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## user_status_user_status_revert_status

> crate::models::UserStatusUserStatusGetStatus200Response user_status_user_status_revert_status(message_id, ocs_api_request)
Revert the status to the previous status

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**message_id** | **String** | ID of the message to delete | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::UserStatusUserStatusGetStatus200Response**](user_status_user_status_get_status_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## user_status_user_status_set_custom_message

> crate::models::UserStatusUserStatusGetStatus200Response user_status_user_status_set_custom_message(ocs_api_request, status_icon, message, clear_at)
Set the message to a custom message for the current user

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**ocs_api_request** | **String** |  | [required] |[default to true]
**status_icon** | Option<**String**> | Icon of the status |  |
**message** | Option<**String**> | Message of the status |  |
**clear_at** | Option<**i64**> | When the message should be cleared |  |

### Return type

[**crate::models::UserStatusUserStatusGetStatus200Response**](user_status_user_status_get_status_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## user_status_user_status_set_predefined_message

> crate::models::UserStatusUserStatusGetStatus200Response user_status_user_status_set_predefined_message(message_id, ocs_api_request, clear_at)
Set the message to a predefined message for the current user

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**message_id** | **String** | ID of the predefined message | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]
**clear_at** | Option<**i64**> | When the message should be cleared |  |

### Return type

[**crate::models::UserStatusUserStatusGetStatus200Response**](user_status_user_status_get_status_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## user_status_user_status_set_status

> crate::models::UserStatusUserStatusGetStatus200Response user_status_user_status_set_status(status_type, ocs_api_request)
Update the status type of the current user

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**status_type** | **String** | The new status type | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::UserStatusUserStatusGetStatus200Response**](user_status_user_status_get_status_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

