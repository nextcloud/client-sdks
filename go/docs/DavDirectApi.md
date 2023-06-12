# \DavDirectApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DavDirectGetUrl**](DavDirectApi.md#DavDirectGetUrl) | **Post** /ocs/v2.php/apps/dav/api/v1/direct | Get a direct link to a file



## DavDirectGetUrl

> DavDirectGetUrl200Response DavDirectGetUrl(ctx).FileId(fileId).OCSAPIRequest(oCSAPIRequest).ExpirationTime(expirationTime).Execute()

Get a direct link to a file

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
    fileId := int64(789) // int64 | ID of the file
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")
    expirationTime := int64(789) // int64 | Duration until the link expires (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DavDirectApi.DavDirectGetUrl(context.Background()).FileId(fileId).OCSAPIRequest(oCSAPIRequest).ExpirationTime(expirationTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DavDirectApi.DavDirectGetUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DavDirectGetUrl`: DavDirectGetUrl200Response
    fmt.Fprintf(os.Stdout, "Response from `DavDirectApi.DavDirectGetUrl`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDavDirectGetUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fileId** | **int64** | ID of the file | 
 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]
 **expirationTime** | **int64** | Duration until the link expires | 

### Return type

[**DavDirectGetUrl200Response**](DavDirectGetUrl200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

