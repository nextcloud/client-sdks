# \FilesSharingShareesapiApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**files_sharing_shareesapi_find_recommended**](FilesSharingShareesapiApi.md#files_sharing_shareesapi_find_recommended) | **GET** /ocs/v2.php/apps/files_sharing/api/v1/sharees_recommended | Find recommended sharees
[**files_sharing_shareesapi_search**](FilesSharingShareesapiApi.md#files_sharing_shareesapi_search) | **GET** /ocs/v2.php/apps/files_sharing/api/v1/sharees | Search for sharees



## files_sharing_shareesapi_find_recommended

> crate::models::FilesSharingShareesapiFindRecommended200Response files_sharing_shareesapi_find_recommended(item_type, ocs_api_request, share_type)
Find recommended sharees

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**item_type** | **String** | Limit to specific item types | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]
**share_type** | Option<**String**> | Limit to specific share types |  |

### Return type

[**crate::models::FilesSharingShareesapiFindRecommended200Response**](files_sharing_shareesapi_find_recommended_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## files_sharing_shareesapi_search

> crate::models::FilesSharingShareesapiSearch200Response files_sharing_shareesapi_search(ocs_api_request, search, item_type, page, per_page, share_type, lookup)
Search for sharees

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**ocs_api_request** | **String** |  | [required] |[default to true]
**search** | Option<**String**> | Text to search for |  |[default to ]
**item_type** | Option<**String**> | Limit to specific item types |  |
**page** | Option<**i64**> | Page offset for searching |  |[default to 1]
**per_page** | Option<**i64**> | Limit amount of search results per page |  |[default to 200]
**share_type** | Option<**String**> | Limit to specific share types |  |
**lookup** | Option<**i32**> | If a global lookup should be performed too |  |[default to 0]

### Return type

[**crate::models::FilesSharingShareesapiSearch200Response**](files_sharing_shareesapi_search_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

