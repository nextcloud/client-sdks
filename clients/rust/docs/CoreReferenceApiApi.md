# \CoreReferenceApiApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**core_reference_api_extract**](CoreReferenceApiApi.md#core_reference_api_extract) | **POST** /ocs/v2.php/references/extract | Extract references from a text
[**core_reference_api_get_providers_info**](CoreReferenceApiApi.md#core_reference_api_get_providers_info) | **GET** /ocs/v2.php/references/providers | Get the providers
[**core_reference_api_resolve**](CoreReferenceApiApi.md#core_reference_api_resolve) | **POST** /ocs/v2.php/references/resolve | Resolve multiple references
[**core_reference_api_resolve_one**](CoreReferenceApiApi.md#core_reference_api_resolve_one) | **GET** /ocs/v2.php/references/resolve | Resolve a reference
[**core_reference_api_touch_provider**](CoreReferenceApiApi.md#core_reference_api_touch_provider) | **PUT** /ocs/v2.php/references/provider/{providerId} | Touch a provider



## core_reference_api_extract

> crate::models::CoreReferenceApiResolveOne200Response core_reference_api_extract(text, ocs_api_request, resolve, limit)
Extract references from a text

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**text** | **String** | Text to extract from | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]
**resolve** | Option<**i32**> | Resolve the references |  |[default to 0]
**limit** | Option<**i64**> | Maximum amount of references to extract |  |[default to 1]

### Return type

[**crate::models::CoreReferenceApiResolveOne200Response**](core_reference_api_resolve_one_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## core_reference_api_get_providers_info

> crate::models::CoreReferenceApiGetProvidersInfo200Response core_reference_api_get_providers_info(ocs_api_request)
Get the providers

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::CoreReferenceApiGetProvidersInfo200Response**](core_reference_api_get_providers_info_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## core_reference_api_resolve

> crate::models::CoreReferenceApiResolveOne200Response core_reference_api_resolve(references_left_square_bracket_right_square_bracket, ocs_api_request, limit)
Resolve multiple references

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**references_left_square_bracket_right_square_bracket** | [**Vec<String>**](String.md) | References to resolve | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]
**limit** | Option<**i64**> | Maximum amount of references to resolve |  |[default to 1]

### Return type

[**crate::models::CoreReferenceApiResolveOne200Response**](core_reference_api_resolve_one_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## core_reference_api_resolve_one

> crate::models::CoreReferenceApiResolveOne200Response core_reference_api_resolve_one(reference, ocs_api_request)
Resolve a reference

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**reference** | **String** | Reference to resolve | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::CoreReferenceApiResolveOne200Response**](core_reference_api_resolve_one_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## core_reference_api_touch_provider

> crate::models::CoreReferenceApiTouchProvider200Response core_reference_api_touch_provider(provider_id, ocs_api_request, timestamp)
Touch a provider

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**provider_id** | **String** | ID of the provider | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]
**timestamp** | Option<**i64**> | Timestamp of the last usage |  |

### Return type

[**crate::models::CoreReferenceApiTouchProvider200Response**](core_reference_api_touch_provider_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

