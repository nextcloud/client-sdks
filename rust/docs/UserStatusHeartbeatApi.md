# \UserStatusHeartbeatApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**user_status_heartbeat_heartbeat**](UserStatusHeartbeatApi.md#user_status_heartbeat_heartbeat) | **PUT** /ocs/v2.php/apps/user_status/api/v1/heartbeat | Keep the current status alive



## user_status_heartbeat_heartbeat

> crate::models::UserStatusUserStatusGetStatus200Response user_status_heartbeat_heartbeat(status, ocs_api_request)
Keep the current status alive

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**status** | **String** | Only online, away | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::UserStatusUserStatusGetStatus200Response**](user_status_user_status_get_status_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

