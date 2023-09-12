# \FilesSharingDeletedShareapiApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**files_sharing_deleted_shareapi_list**](FilesSharingDeletedShareapiApi.md#files_sharing_deleted_shareapi_list) | **GET** /ocs/v2.php/apps/files_sharing/api/v1/deletedshares | Get a list of all deleted shares
[**files_sharing_deleted_shareapi_undelete**](FilesSharingDeletedShareapiApi.md#files_sharing_deleted_shareapi_undelete) | **POST** /ocs/v2.php/apps/files_sharing/api/v1/deletedshares/{id} | Undelete a deleted share



## files_sharing_deleted_shareapi_list

> crate::models::FilesSharingDeletedShareapiList200Response files_sharing_deleted_shareapi_list(ocs_api_request)
Get a list of all deleted shares

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::FilesSharingDeletedShareapiList200Response**](files_sharing_deleted_shareapi_list_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## files_sharing_deleted_shareapi_undelete

> crate::models::CoreWhatsNewDismiss200Response files_sharing_deleted_shareapi_undelete(id, ocs_api_request)
Undelete a deleted share

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** | ID of the share | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::CoreWhatsNewDismiss200Response**](core_whats_new_dismiss_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

