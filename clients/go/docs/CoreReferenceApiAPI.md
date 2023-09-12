# \CoreReferenceApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CoreReferenceApiExtract**](CoreReferenceApiAPI.md#CoreReferenceApiExtract) | **Post** /ocs/v2.php/references/extract | Extract references from a text
[**CoreReferenceApiGetProvidersInfo**](CoreReferenceApiAPI.md#CoreReferenceApiGetProvidersInfo) | **Get** /ocs/v2.php/references/providers | Get the providers
[**CoreReferenceApiResolve**](CoreReferenceApiAPI.md#CoreReferenceApiResolve) | **Post** /ocs/v2.php/references/resolve | Resolve multiple references
[**CoreReferenceApiResolveOne**](CoreReferenceApiAPI.md#CoreReferenceApiResolveOne) | **Get** /ocs/v2.php/references/resolve | Resolve a reference
[**CoreReferenceApiTouchProvider**](CoreReferenceApiAPI.md#CoreReferenceApiTouchProvider) | **Put** /ocs/v2.php/references/provider/{providerId} | Touch a provider



## CoreReferenceApiExtract

> CoreReferenceApiResolveOne200Response CoreReferenceApiExtract(ctx).Text(text).OCSAPIRequest(oCSAPIRequest).Resolve(resolve).Limit(limit).Execute()

Extract references from a text

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
    text := "text_example" // string | Text to extract from
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")
    resolve := int32(56) // int32 | Resolve the references (optional) (default to 0)
    limit := int64(789) // int64 | Maximum amount of references to extract (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoreReferenceApiAPI.CoreReferenceApiExtract(context.Background()).Text(text).OCSAPIRequest(oCSAPIRequest).Resolve(resolve).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreReferenceApiAPI.CoreReferenceApiExtract``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CoreReferenceApiExtract`: CoreReferenceApiResolveOne200Response
    fmt.Fprintf(os.Stdout, "Response from `CoreReferenceApiAPI.CoreReferenceApiExtract`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCoreReferenceApiExtractRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **text** | **string** | Text to extract from | 
 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]
 **resolve** | **int32** | Resolve the references | [default to 0]
 **limit** | **int64** | Maximum amount of references to extract | [default to 1]

### Return type

[**CoreReferenceApiResolveOne200Response**](CoreReferenceApiResolveOne200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreReferenceApiGetProvidersInfo

> CoreReferenceApiGetProvidersInfo200Response CoreReferenceApiGetProvidersInfo(ctx).OCSAPIRequest(oCSAPIRequest).Execute()

Get the providers

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
    resp, r, err := apiClient.CoreReferenceApiAPI.CoreReferenceApiGetProvidersInfo(context.Background()).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreReferenceApiAPI.CoreReferenceApiGetProvidersInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CoreReferenceApiGetProvidersInfo`: CoreReferenceApiGetProvidersInfo200Response
    fmt.Fprintf(os.Stdout, "Response from `CoreReferenceApiAPI.CoreReferenceApiGetProvidersInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCoreReferenceApiGetProvidersInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]

### Return type

[**CoreReferenceApiGetProvidersInfo200Response**](CoreReferenceApiGetProvidersInfo200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreReferenceApiResolve

> CoreReferenceApiResolveOne200Response CoreReferenceApiResolve(ctx).References(references).OCSAPIRequest(oCSAPIRequest).Limit(limit).Execute()

Resolve multiple references

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
    references := []string{"Inner_example"} // []string | References to resolve
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")
    limit := int64(789) // int64 | Maximum amount of references to resolve (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoreReferenceApiAPI.CoreReferenceApiResolve(context.Background()).References(references).OCSAPIRequest(oCSAPIRequest).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreReferenceApiAPI.CoreReferenceApiResolve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CoreReferenceApiResolve`: CoreReferenceApiResolveOne200Response
    fmt.Fprintf(os.Stdout, "Response from `CoreReferenceApiAPI.CoreReferenceApiResolve`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCoreReferenceApiResolveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **references** | **[]string** | References to resolve | 
 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]
 **limit** | **int64** | Maximum amount of references to resolve | [default to 1]

### Return type

[**CoreReferenceApiResolveOne200Response**](CoreReferenceApiResolveOne200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreReferenceApiResolveOne

> CoreReferenceApiResolveOne200Response CoreReferenceApiResolveOne(ctx).Reference(reference).OCSAPIRequest(oCSAPIRequest).Execute()

Resolve a reference

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
    reference := "reference_example" // string | Reference to resolve
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoreReferenceApiAPI.CoreReferenceApiResolveOne(context.Background()).Reference(reference).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreReferenceApiAPI.CoreReferenceApiResolveOne``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CoreReferenceApiResolveOne`: CoreReferenceApiResolveOne200Response
    fmt.Fprintf(os.Stdout, "Response from `CoreReferenceApiAPI.CoreReferenceApiResolveOne`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCoreReferenceApiResolveOneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reference** | **string** | Reference to resolve | 
 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]

### Return type

[**CoreReferenceApiResolveOne200Response**](CoreReferenceApiResolveOne200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreReferenceApiTouchProvider

> CoreReferenceApiTouchProvider200Response CoreReferenceApiTouchProvider(ctx, providerId).OCSAPIRequest(oCSAPIRequest).Timestamp(timestamp).Execute()

Touch a provider

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
    providerId := "providerId_example" // string | ID of the provider
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")
    timestamp := int64(789) // int64 | Timestamp of the last usage (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoreReferenceApiAPI.CoreReferenceApiTouchProvider(context.Background(), providerId).OCSAPIRequest(oCSAPIRequest).Timestamp(timestamp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreReferenceApiAPI.CoreReferenceApiTouchProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CoreReferenceApiTouchProvider`: CoreReferenceApiTouchProvider200Response
    fmt.Fprintf(os.Stdout, "Response from `CoreReferenceApiAPI.CoreReferenceApiTouchProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** | ID of the provider | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreReferenceApiTouchProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]
 **timestamp** | **int64** | Timestamp of the last usage | 

### Return type

[**CoreReferenceApiTouchProvider200Response**](CoreReferenceApiTouchProvider200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

