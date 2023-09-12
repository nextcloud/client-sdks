# \DavDirectApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**dav_direct_get_url**](DavDirectApi.md#dav_direct_get_url) | **POST** /ocs/v2.php/apps/dav/api/v1/direct | Get a direct link to a file



## dav_direct_get_url

> crate::models::DavDirectGetUrl200Response dav_direct_get_url(file_id, ocs_api_request, expiration_time)
Get a direct link to a file

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**file_id** | **i64** | ID of the file | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]
**expiration_time** | Option<**i64**> | Duration until the link expires |  |

### Return type

[**crate::models::DavDirectGetUrl200Response**](dav_direct_get_url_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

