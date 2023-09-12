# \FilesExternalApiApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**files_external_api_get_user_mounts**](FilesExternalApiApi.md#files_external_api_get_user_mounts) | **GET** /ocs/v2.php/apps/files_external/api/v1/mounts | Get the mount points visible for this user



## files_external_api_get_user_mounts

> crate::models::FilesExternalApiGetUserMounts200Response files_external_api_get_user_mounts(ocs_api_request)
Get the mount points visible for this user

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::FilesExternalApiGetUserMounts200Response**](files_external_api_get_user_mounts_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

