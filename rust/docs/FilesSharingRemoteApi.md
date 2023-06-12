# \FilesSharingRemoteApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**files_sharing_remote_accept_share**](FilesSharingRemoteApi.md#files_sharing_remote_accept_share) | **POST** /ocs/v2.php/apps/files_sharing/api/v1/remote_shares/pending/{id} | Accept a remote share
[**files_sharing_remote_decline_share**](FilesSharingRemoteApi.md#files_sharing_remote_decline_share) | **DELETE** /ocs/v2.php/apps/files_sharing/api/v1/remote_shares/pending/{id} | Decline a remote share
[**files_sharing_remote_get_open_shares**](FilesSharingRemoteApi.md#files_sharing_remote_get_open_shares) | **GET** /ocs/v2.php/apps/files_sharing/api/v1/remote_shares/pending | Get list of pending remote shares
[**files_sharing_remote_get_share**](FilesSharingRemoteApi.md#files_sharing_remote_get_share) | **GET** /ocs/v2.php/apps/files_sharing/api/v1/remote_shares/{id} | Get info of a remote share
[**files_sharing_remote_get_shares**](FilesSharingRemoteApi.md#files_sharing_remote_get_shares) | **GET** /ocs/v2.php/apps/files_sharing/api/v1/remote_shares | Get a list of accepted remote shares
[**files_sharing_remote_unshare**](FilesSharingRemoteApi.md#files_sharing_remote_unshare) | **DELETE** /ocs/v2.php/apps/files_sharing/api/v1/remote_shares/{id} | Unshare a remote share



## files_sharing_remote_accept_share

> crate::models::CoreWhatsNewDismiss200Response files_sharing_remote_accept_share(id, ocs_api_request)
Accept a remote share

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **i64** | ID of the share | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::CoreWhatsNewDismiss200Response**](core_whats_new_dismiss_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## files_sharing_remote_decline_share

> crate::models::CoreWhatsNewDismiss200Response files_sharing_remote_decline_share(id, ocs_api_request)
Decline a remote share

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **i64** | ID of the share | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::CoreWhatsNewDismiss200Response**](core_whats_new_dismiss_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## files_sharing_remote_get_open_shares

> crate::models::FilesSharingRemoteGetShares200Response files_sharing_remote_get_open_shares(ocs_api_request)
Get list of pending remote shares

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::FilesSharingRemoteGetShares200Response**](files_sharing_remote_get_shares_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## files_sharing_remote_get_share

> crate::models::FilesSharingRemoteGetShare200Response files_sharing_remote_get_share(id, ocs_api_request)
Get info of a remote share

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **i64** | ID of the share | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::FilesSharingRemoteGetShare200Response**](files_sharing_remote_get_share_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## files_sharing_remote_get_shares

> crate::models::FilesSharingRemoteGetShares200Response files_sharing_remote_get_shares(ocs_api_request)
Get a list of accepted remote shares

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::FilesSharingRemoteGetShares200Response**](files_sharing_remote_get_shares_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## files_sharing_remote_unshare

> crate::models::CoreWhatsNewDismiss200Response files_sharing_remote_unshare(id, ocs_api_request)
Unshare a remote share

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **i64** | ID of the share | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::CoreWhatsNewDismiss200Response**](core_whats_new_dismiss_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

