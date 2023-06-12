# \CoreReferenceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CoreReferencePreview**](CoreReferenceApi.md#CoreReferencePreview) | **Get** /index.php/core/references/preview/{referenceId} | Get a preview for a reference



## CoreReferencePreview

> *os.File CoreReferencePreview(ctx, referenceId).Execute()

Get a preview for a reference

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
    referenceId := "referenceId_example" // string | the reference cache key

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoreReferenceApi.CoreReferencePreview(context.Background(), referenceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreReferenceApi.CoreReferencePreview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CoreReferencePreview`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `CoreReferenceApi.CoreReferencePreview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**referenceId** | **string** | the reference cache key | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreReferencePreviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

