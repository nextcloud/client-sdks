# \CoreTranslationApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CoreTranslationApiLanguages**](CoreTranslationApiAPI.md#CoreTranslationApiLanguages) | **Get** /ocs/v2.php/translation/languages | Get the list of supported languages
[**CoreTranslationApiTranslate**](CoreTranslationApiAPI.md#CoreTranslationApiTranslate) | **Post** /ocs/v2.php/translation/translate | Translate a text



## CoreTranslationApiLanguages

> CoreTranslationApiLanguages200Response CoreTranslationApiLanguages(ctx).OCSAPIRequest(oCSAPIRequest).Execute()

Get the list of supported languages

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/nextcloud/client-sdks"
)

func main() {
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoreTranslationApiAPI.CoreTranslationApiLanguages(context.Background()).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreTranslationApiAPI.CoreTranslationApiLanguages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CoreTranslationApiLanguages`: CoreTranslationApiLanguages200Response
    fmt.Fprintf(os.Stdout, "Response from `CoreTranslationApiAPI.CoreTranslationApiLanguages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCoreTranslationApiLanguagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]

### Return type

[**CoreTranslationApiLanguages200Response**](CoreTranslationApiLanguages200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreTranslationApiTranslate

> CoreTranslationApiTranslate200Response CoreTranslationApiTranslate(ctx).Text(text).ToLanguage(toLanguage).OCSAPIRequest(oCSAPIRequest).FromLanguage(fromLanguage).Execute()

Translate a text

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/nextcloud/client-sdks"
)

func main() {
    text := "text_example" // string | Text to be translated
    toLanguage := "toLanguage_example" // string | Language to translate to
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")
    fromLanguage := "fromLanguage_example" // string | Language to translate from (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoreTranslationApiAPI.CoreTranslationApiTranslate(context.Background()).Text(text).ToLanguage(toLanguage).OCSAPIRequest(oCSAPIRequest).FromLanguage(fromLanguage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreTranslationApiAPI.CoreTranslationApiTranslate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CoreTranslationApiTranslate`: CoreTranslationApiTranslate200Response
    fmt.Fprintf(os.Stdout, "Response from `CoreTranslationApiAPI.CoreTranslationApiTranslate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCoreTranslationApiTranslateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **text** | **string** | Text to be translated | 
 **toLanguage** | **string** | Language to translate to | 
 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]
 **fromLanguage** | **string** | Language to translate from | 

### Return type

[**CoreTranslationApiTranslate200Response**](CoreTranslationApiTranslate200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

