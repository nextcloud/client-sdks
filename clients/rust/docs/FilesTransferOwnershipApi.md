# \FilesTransferOwnershipApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**files_transfer_ownership_accept**](FilesTransferOwnershipApi.md#files_transfer_ownership_accept) | **POST** /ocs/v2.php/apps/files/api/v1/transferownership/{id} | Accept an ownership transfer
[**files_transfer_ownership_reject**](FilesTransferOwnershipApi.md#files_transfer_ownership_reject) | **DELETE** /ocs/v2.php/apps/files/api/v1/transferownership/{id} | Reject an ownership transfer
[**files_transfer_ownership_transfer**](FilesTransferOwnershipApi.md#files_transfer_ownership_transfer) | **POST** /ocs/v2.php/apps/files/api/v1/transferownership | Transfer the ownership to another user



## files_transfer_ownership_accept

> crate::models::CoreWhatsNewDismiss200Response files_transfer_ownership_accept(id, ocs_api_request)
Accept an ownership transfer

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **i64** | ID of the ownership transfer | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::CoreWhatsNewDismiss200Response**](core_whats_new_dismiss_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## files_transfer_ownership_reject

> crate::models::CoreWhatsNewDismiss200Response files_transfer_ownership_reject(id, ocs_api_request)
Reject an ownership transfer

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **i64** | ID of the ownership transfer | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::CoreWhatsNewDismiss200Response**](core_whats_new_dismiss_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## files_transfer_ownership_transfer

> crate::models::CoreWhatsNewDismiss200Response files_transfer_ownership_transfer(recipient, path, ocs_api_request)
Transfer the ownership to another user

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**recipient** | **String** | Username of the recipient | [required] |
**path** | **String** | Path of the file | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::CoreWhatsNewDismiss200Response**](core_whats_new_dismiss_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

