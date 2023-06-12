# \FilesSharingShareesapiApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FilesSharingShareesapiFindRecommended**](FilesSharingShareesapiApi.md#FilesSharingShareesapiFindRecommended) | **Get** /ocs/v2.php/apps/files_sharing/api/v1/sharees_recommended | Find recommended sharees
[**FilesSharingShareesapiSearch**](FilesSharingShareesapiApi.md#FilesSharingShareesapiSearch) | **Get** /ocs/v2.php/apps/files_sharing/api/v1/sharees | Search for sharees



## FilesSharingShareesapiFindRecommended

> FilesSharingShareesapiFindRecommended200Response FilesSharingShareesapiFindRecommended(ctx).ItemType(itemType).OCSAPIRequest(oCSAPIRequest).ShareType(shareType).Execute()

Find recommended sharees

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/nextcloud/api-sdk"
)

func main() {
    itemType := "itemType_example" // string | Limit to specific item types
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")
    shareType := "shareType_example" // string | Limit to specific share types (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FilesSharingShareesapiApi.FilesSharingShareesapiFindRecommended(context.Background()).ItemType(itemType).OCSAPIRequest(oCSAPIRequest).ShareType(shareType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesSharingShareesapiApi.FilesSharingShareesapiFindRecommended``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FilesSharingShareesapiFindRecommended`: FilesSharingShareesapiFindRecommended200Response
    fmt.Fprintf(os.Stdout, "Response from `FilesSharingShareesapiApi.FilesSharingShareesapiFindRecommended`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFilesSharingShareesapiFindRecommendedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **itemType** | **string** | Limit to specific item types | 
 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]
 **shareType** | **string** | Limit to specific share types | 

### Return type

[**FilesSharingShareesapiFindRecommended200Response**](FilesSharingShareesapiFindRecommended200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FilesSharingShareesapiSearch

> FilesSharingShareesapiSearch200Response FilesSharingShareesapiSearch(ctx).OCSAPIRequest(oCSAPIRequest).Search(search).ItemType(itemType).Page(page).PerPage(perPage).ShareType(shareType).Lookup(lookup).Execute()

Search for sharees

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/nextcloud/api-sdk"
)

func main() {
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")
    search := "search_example" // string | Text to search for (optional) (default to "")
    itemType := "itemType_example" // string | Limit to specific item types (optional)
    page := int64(789) // int64 | Page offset for searching (optional) (default to 1)
    perPage := int64(789) // int64 | Limit amount of search results per page (optional) (default to 200)
    shareType := "shareType_example" // string | Limit to specific share types (optional)
    lookup := int32(56) // int32 | If a global lookup should be performed too (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FilesSharingShareesapiApi.FilesSharingShareesapiSearch(context.Background()).OCSAPIRequest(oCSAPIRequest).Search(search).ItemType(itemType).Page(page).PerPage(perPage).ShareType(shareType).Lookup(lookup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesSharingShareesapiApi.FilesSharingShareesapiSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FilesSharingShareesapiSearch`: FilesSharingShareesapiSearch200Response
    fmt.Fprintf(os.Stdout, "Response from `FilesSharingShareesapiApi.FilesSharingShareesapiSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFilesSharingShareesapiSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]
 **search** | **string** | Text to search for | [default to &quot;&quot;]
 **itemType** | **string** | Limit to specific item types | 
 **page** | **int64** | Page offset for searching | [default to 1]
 **perPage** | **int64** | Limit amount of search results per page | [default to 200]
 **shareType** | **string** | Limit to specific share types | 
 **lookup** | **int32** | If a global lookup should be performed too | [default to 0]

### Return type

[**FilesSharingShareesapiSearch200Response**](FilesSharingShareesapiSearch200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

