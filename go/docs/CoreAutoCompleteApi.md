# \CoreAutoCompleteApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CoreAutoCompleteGet**](CoreAutoCompleteApi.md#CoreAutoCompleteGet) | **Get** /ocs/v2.php/core/autocomplete/get | Autocomplete a query



## CoreAutoCompleteGet

> CoreAutoCompleteGet200Response CoreAutoCompleteGet(ctx).Search(search).OCSAPIRequest(oCSAPIRequest).ItemType(itemType).ItemId(itemId).Sorter(sorter).ShareTypes(shareTypes).Limit(limit).Execute()

Autocomplete a query

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
    search := "search_example" // string | Text to search for
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")
    itemType := "itemType_example" // string | Type of the items to search for (optional)
    itemId := "itemId_example" // string | ID of the items to search for (optional)
    sorter := "sorter_example" // string | can be piped, top prio first, e.g.: \"commenters|share-recipients\" (optional)
    shareTypes := "shareTypes_example" // string | Types of shares to search for (optional)
    limit := int64(789) // int64 | Maximum number of results to return (optional) (default to 10)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoreAutoCompleteApi.CoreAutoCompleteGet(context.Background()).Search(search).OCSAPIRequest(oCSAPIRequest).ItemType(itemType).ItemId(itemId).Sorter(sorter).ShareTypes(shareTypes).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreAutoCompleteApi.CoreAutoCompleteGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CoreAutoCompleteGet`: CoreAutoCompleteGet200Response
    fmt.Fprintf(os.Stdout, "Response from `CoreAutoCompleteApi.CoreAutoCompleteGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCoreAutoCompleteGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | Text to search for | 
 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]
 **itemType** | **string** | Type of the items to search for | 
 **itemId** | **string** | ID of the items to search for | 
 **sorter** | **string** | can be piped, top prio first, e.g.: \&quot;commenters|share-recipients\&quot; | 
 **shareTypes** | **string** | Types of shares to search for | 
 **limit** | **int64** | Maximum number of results to return | [default to 10]

### Return type

[**CoreAutoCompleteGet200Response**](CoreAutoCompleteGet200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

