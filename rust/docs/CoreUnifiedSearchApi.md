# \CoreUnifiedSearchApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**core_unified_search_get_providers**](CoreUnifiedSearchApi.md#core_unified_search_get_providers) | **GET** /ocs/v2.php/search/providers | Get the providers for unified search
[**core_unified_search_search**](CoreUnifiedSearchApi.md#core_unified_search_search) | **GET** /ocs/v2.php/search/providers/{providerId}/search | Search



## core_unified_search_get_providers

> crate::models::CoreUnifiedSearchGetProviders200Response core_unified_search_get_providers(ocs_api_request, from)
Get the providers for unified search

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**ocs_api_request** | **String** |  | [required] |[default to true]
**from** | Option<**String**> | the url the user is currently at |  |[default to ]

### Return type

[**crate::models::CoreUnifiedSearchGetProviders200Response**](core_unified_search_get_providers_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## core_unified_search_search

> crate::models::CoreUnifiedSearchSearch200Response core_unified_search_search(provider_id, ocs_api_request, term, sort_order, limit, cursor, from)
Search

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**provider_id** | **String** | ID of the provider | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]
**term** | Option<**String**> | Term to search |  |[default to ]
**sort_order** | Option<**i64**> | Order of entries |  |
**limit** | Option<**i64**> | Maximum amount of entries |  |
**cursor** | Option<**String**> | Offset for searching |  |
**from** | Option<**String**> | The current user URL |  |[default to ]

### Return type

[**crate::models::CoreUnifiedSearchSearch200Response**](core_unified_search_search_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

