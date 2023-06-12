# \FilesSharingShareapiApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**files_sharing_shareapi_accept_share**](FilesSharingShareapiApi.md#files_sharing_shareapi_accept_share) | **POST** /ocs/v2.php/apps/files_sharing/api/v1/shares/pending/{id} | Accept a share
[**files_sharing_shareapi_create_share**](FilesSharingShareapiApi.md#files_sharing_shareapi_create_share) | **POST** /ocs/v2.php/apps/files_sharing/api/v1/shares | Create a share
[**files_sharing_shareapi_delete_share**](FilesSharingShareapiApi.md#files_sharing_shareapi_delete_share) | **DELETE** /ocs/v2.php/apps/files_sharing/api/v1/shares/{id} | Delete a share
[**files_sharing_shareapi_get_inherited_shares**](FilesSharingShareapiApi.md#files_sharing_shareapi_get_inherited_shares) | **GET** /ocs/v2.php/apps/files_sharing/api/v1/shares/inherited | Get all shares relative to a file, including parent folders shares rights
[**files_sharing_shareapi_get_share**](FilesSharingShareapiApi.md#files_sharing_shareapi_get_share) | **GET** /ocs/v2.php/apps/files_sharing/api/v1/shares/{id} | Get a specific share by id
[**files_sharing_shareapi_get_shares**](FilesSharingShareapiApi.md#files_sharing_shareapi_get_shares) | **GET** /ocs/v2.php/apps/files_sharing/api/v1/shares | Get shares of the current user
[**files_sharing_shareapi_pending_shares**](FilesSharingShareapiApi.md#files_sharing_shareapi_pending_shares) | **GET** /ocs/v2.php/apps/files_sharing/api/v1/shares/pending | Get all shares that are still pending
[**files_sharing_shareapi_update_share**](FilesSharingShareapiApi.md#files_sharing_shareapi_update_share) | **PUT** /ocs/v2.php/apps/files_sharing/api/v1/shares/{id} | Update a share



## files_sharing_shareapi_accept_share

> crate::models::CoreWhatsNewDismiss200Response files_sharing_shareapi_accept_share(id, ocs_api_request)
Accept a share

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


## files_sharing_shareapi_create_share

> crate::models::FilesSharingShareapiCreateShare200Response files_sharing_shareapi_create_share(ocs_api_request, path, permissions, share_type, share_with, public_upload, password, send_password_by_talk, expire_date, note, label, attributes)
Create a share

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**ocs_api_request** | **String** |  | [required] |[default to true]
**path** | Option<**String**> | Path of the share |  |
**permissions** | Option<**i64**> | Permissions for the share |  |
**share_type** | Option<**i64**> | Type of the share |  |[default to 1]
**share_with** | Option<**String**> | The entity this should be shared with |  |
**public_upload** | Option<**String**> | If public uploading is allowed |  |[default to false]
**password** | Option<**String**> | Password for the share |  |[default to ]
**send_password_by_talk** | Option<**String**> | Send the password for the share over Talk |  |
**expire_date** | Option<**String**> | Expiry date of the share |  |[default to ]
**note** | Option<**String**> | Note for the share |  |[default to ]
**label** | Option<**String**> | Label for the share (only used in link and email) |  |[default to ]
**attributes** | Option<**String**> | Additional attributes for the share |  |

### Return type

[**crate::models::FilesSharingShareapiCreateShare200Response**](files_sharing_shareapi_create_share_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## files_sharing_shareapi_delete_share

> crate::models::CoreWhatsNewDismiss200Response files_sharing_shareapi_delete_share(id, ocs_api_request)
Delete a share

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


## files_sharing_shareapi_get_inherited_shares

> crate::models::FilesSharingShareapiGetShares200Response files_sharing_shareapi_get_inherited_shares(path, ocs_api_request)
Get all shares relative to a file, including parent folders shares rights

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**path** | **String** | Path all shares will be relative to | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::FilesSharingShareapiGetShares200Response**](files_sharing_shareapi_get_shares_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## files_sharing_shareapi_get_share

> crate::models::FilesSharingShareapiCreateShare200Response files_sharing_shareapi_get_share(id, ocs_api_request, include_tags)
Get a specific share by id

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** | ID of the share | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]
**include_tags** | Option<**i32**> | Include tags in the share |  |[default to 0]

### Return type

[**crate::models::FilesSharingShareapiCreateShare200Response**](files_sharing_shareapi_create_share_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## files_sharing_shareapi_get_shares

> crate::models::FilesSharingShareapiGetShares200Response files_sharing_shareapi_get_shares(ocs_api_request, shared_with_me, reshares, subfiles, path, include_tags)
Get shares of the current user

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**ocs_api_request** | **String** |  | [required] |[default to true]
**shared_with_me** | Option<**String**> | Only get shares with the current user |  |[default to false]
**reshares** | Option<**String**> | Only get shares by the current user and reshares |  |[default to false]
**subfiles** | Option<**String**> | Only get all shares in a folder |  |[default to false]
**path** | Option<**String**> | Get shares for a specific path |  |[default to ]
**include_tags** | Option<**String**> | Include tags in the share |  |[default to false]

### Return type

[**crate::models::FilesSharingShareapiGetShares200Response**](files_sharing_shareapi_get_shares_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## files_sharing_shareapi_pending_shares

> crate::models::FilesSharingShareapiGetShares200Response files_sharing_shareapi_pending_shares(ocs_api_request)
Get all shares that are still pending

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::FilesSharingShareapiGetShares200Response**](files_sharing_shareapi_get_shares_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## files_sharing_shareapi_update_share

> crate::models::FilesSharingShareapiCreateShare200Response files_sharing_shareapi_update_share(id, ocs_api_request, permissions, password, send_password_by_talk, public_upload, expire_date, note, label, hide_download, attributes)
Update a share

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **String** | ID of the share | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]
**permissions** | Option<**i64**> | New permissions |  |
**password** | Option<**String**> | New password |  |
**send_password_by_talk** | Option<**String**> | New condition if the password should be send over Talk |  |
**public_upload** | Option<**String**> | New condition if public uploading is allowed |  |
**expire_date** | Option<**String**> | New expiry date |  |
**note** | Option<**String**> | New note |  |
**label** | Option<**String**> | New label |  |
**hide_download** | Option<**String**> | New condition if the download should be hidden |  |
**attributes** | Option<**String**> | New additional attributes |  |

### Return type

[**crate::models::FilesSharingShareapiCreateShare200Response**](files_sharing_shareapi_create_share_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

