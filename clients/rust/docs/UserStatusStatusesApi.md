# \UserStatusStatusesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**user_status_statuses_find**](UserStatusStatusesApi.md#user_status_statuses_find) | **GET** /ocs/v2.php/apps/user_status/api/v1/statuses/{userId} | Find the status of a user
[**user_status_statuses_find_all**](UserStatusStatusesApi.md#user_status_statuses_find_all) | **GET** /ocs/v2.php/apps/user_status/api/v1/statuses | Find statuses of users



## user_status_statuses_find

> crate::models::UserStatusStatusesFind200Response user_status_statuses_find(user_id, ocs_api_request)
Find the status of a user

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**user_id** | **String** | ID of the user | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::UserStatusStatusesFind200Response**](user_status_statuses_find_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## user_status_statuses_find_all

> crate::models::UserStatusStatusesFindAll200Response user_status_statuses_find_all(ocs_api_request, limit, offset)
Find statuses of users

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**ocs_api_request** | **String** |  | [required] |[default to true]
**limit** | Option<**i64**> | Maximum number of statuses to find |  |
**offset** | Option<**i64**> | Offset for finding statuses |  |

### Return type

[**crate::models::UserStatusStatusesFindAll200Response**](user_status_statuses_find_all_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

