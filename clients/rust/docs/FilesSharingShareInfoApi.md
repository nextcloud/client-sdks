# \FilesSharingShareInfoApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**files_sharing_share_info_info**](FilesSharingShareInfoApi.md#files_sharing_share_info_info) | **POST** /index.php/apps/files_sharing/shareinfo | Get the info about a share



## files_sharing_share_info_info

> crate::models::FilesSharingShareInfo files_sharing_share_info_info(t, password, dir, depth)
Get the info about a share

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**t** | **String** | Token of the share | [required] |
**password** | Option<**String**> | Password of the share |  |
**dir** | Option<**String**> | Subdirectory to get info about |  |
**depth** | Option<**i64**> | Maximum depth to get info about |  |[default to -1]

### Return type

[**crate::models::FilesSharingShareInfo**](FilesSharingShareInfo.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

