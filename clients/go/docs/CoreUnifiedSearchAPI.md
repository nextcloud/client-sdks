# \CoreUnifiedSearchAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CoreUnifiedSearchGetProviders**](CoreUnifiedSearchAPI.md#CoreUnifiedSearchGetProviders) | **Get** /ocs/v2.php/search/providers | Get the providers for unified search
[**CoreUnifiedSearchSearch**](CoreUnifiedSearchAPI.md#CoreUnifiedSearchSearch) | **Get** /ocs/v2.php/search/providers/{providerId}/search | Search



## CoreUnifiedSearchGetProviders

> CoreUnifiedSearchGetProviders200Response CoreUnifiedSearchGetProviders(ctx).OCSAPIRequest(oCSAPIRequest).From(from).Execute()

Get the providers for unified search

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
    from := "from_example" // string | the url the user is currently at (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoreUnifiedSearchAPI.CoreUnifiedSearchGetProviders(context.Background()).OCSAPIRequest(oCSAPIRequest).From(from).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreUnifiedSearchAPI.CoreUnifiedSearchGetProviders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CoreUnifiedSearchGetProviders`: CoreUnifiedSearchGetProviders200Response
    fmt.Fprintf(os.Stdout, "Response from `CoreUnifiedSearchAPI.CoreUnifiedSearchGetProviders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCoreUnifiedSearchGetProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]
 **from** | **string** | the url the user is currently at | [default to &quot;&quot;]

### Return type

[**CoreUnifiedSearchGetProviders200Response**](CoreUnifiedSearchGetProviders200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreUnifiedSearchSearch

> CoreUnifiedSearchSearch200Response CoreUnifiedSearchSearch(ctx, providerId).OCSAPIRequest(oCSAPIRequest).Term(term).SortOrder(sortOrder).Limit(limit).Cursor(cursor).From(from).Execute()

Search

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
    term := "term_example" // string | Term to search (optional) (default to "")
    sortOrder := int64(789) // int64 | Order of entries (optional)
    limit := int64(789) // int64 | Maximum amount of entries (optional)
    cursor := "cursor_example" // string | Offset for searching (optional)
    from := "from_example" // string | The current user URL (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoreUnifiedSearchAPI.CoreUnifiedSearchSearch(context.Background(), providerId).OCSAPIRequest(oCSAPIRequest).Term(term).SortOrder(sortOrder).Limit(limit).Cursor(cursor).From(from).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreUnifiedSearchAPI.CoreUnifiedSearchSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CoreUnifiedSearchSearch`: CoreUnifiedSearchSearch200Response
    fmt.Fprintf(os.Stdout, "Response from `CoreUnifiedSearchAPI.CoreUnifiedSearchSearch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**providerId** | **string** | ID of the provider | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreUnifiedSearchSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]
 **term** | **string** | Term to search | [default to &quot;&quot;]
 **sortOrder** | **int64** | Order of entries | 
 **limit** | **int64** | Maximum amount of entries | 
 **cursor** | **string** | Offset for searching | 
 **from** | **string** | The current user URL | [default to &quot;&quot;]

### Return type

[**CoreUnifiedSearchSearch200Response**](CoreUnifiedSearchSearch200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

