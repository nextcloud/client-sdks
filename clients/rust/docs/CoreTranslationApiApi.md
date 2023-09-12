# \CoreTranslationApiApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**core_translation_api_languages**](CoreTranslationApiApi.md#core_translation_api_languages) | **GET** /ocs/v2.php/translation/languages | Get the list of supported languages
[**core_translation_api_translate**](CoreTranslationApiApi.md#core_translation_api_translate) | **POST** /ocs/v2.php/translation/translate | Translate a text



## core_translation_api_languages

> crate::models::CoreTranslationApiLanguages200Response core_translation_api_languages(ocs_api_request)
Get the list of supported languages

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::CoreTranslationApiLanguages200Response**](core_translation_api_languages_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## core_translation_api_translate

> crate::models::CoreTranslationApiTranslate200Response core_translation_api_translate(text, to_language, ocs_api_request, from_language)
Translate a text

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**text** | **String** | Text to be translated | [required] |
**to_language** | **String** | Language to translate to | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]
**from_language** | Option<**String**> | Language to translate from |  |

### Return type

[**crate::models::CoreTranslationApiTranslate200Response**](core_translation_api_translate_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

