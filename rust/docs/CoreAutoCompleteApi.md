# \CoreAutoCompleteApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**core_auto_complete_get**](CoreAutoCompleteApi.md#core_auto_complete_get) | **GET** /ocs/v2.php/core/autocomplete/get | Autocomplete a query



## core_auto_complete_get

> crate::models::CoreAutoCompleteGet200Response core_auto_complete_get(search, ocs_api_request, item_type, item_id, sorter, share_types, limit)
Autocomplete a query

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**search** | **String** | Text to search for | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]
**item_type** | Option<**String**> | Type of the items to search for |  |
**item_id** | Option<**String**> | ID of the items to search for |  |
**sorter** | Option<**String**> | can be piped, top prio first, e.g.: \"commenters|share-recipients\" |  |
**share_types** | Option<**String**> | Types of shares to search for |  |
**limit** | Option<**i64**> | Maximum number of results to return |  |[default to 10]

### Return type

[**crate::models::CoreAutoCompleteGet200Response**](core_auto_complete_get_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

