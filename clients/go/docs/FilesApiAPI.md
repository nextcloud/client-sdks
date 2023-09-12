# \FilesApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FilesApiGetThumbnail**](FilesApiAPI.md#FilesApiGetThumbnail) | **Get** /index.php/apps/files/api/v1/thumbnail/{x}/{y}/{file} | Gets a thumbnail of the specified file
[**FilesApiServiceWorker**](FilesApiAPI.md#FilesApiServiceWorker) | **Get** /index.php/apps/files/preview-service-worker.js | Get the service-worker Javascript for previews



## FilesApiGetThumbnail

> *os.File FilesApiGetThumbnail(ctx, x, y, file).Execute()

Gets a thumbnail of the specified file

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
    x := int64(789) // int64 | Width of the thumbnail
    y := int64(789) // int64 | Height of the thumbnail
    file := "file_example" // string | URL-encoded filename

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FilesApiAPI.FilesApiGetThumbnail(context.Background(), x, y, file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesApiAPI.FilesApiGetThumbnail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FilesApiGetThumbnail`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `FilesApiAPI.FilesApiGetThumbnail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**x** | **int64** | Width of the thumbnail | 
**y** | **int64** | Height of the thumbnail | 
**file** | **string** | URL-encoded filename | 

### Other Parameters

Other parameters are passed through a pointer to a apiFilesApiGetThumbnailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[***os.File**](*os.File.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FilesApiServiceWorker

> *os.File FilesApiServiceWorker(ctx).Execute()

Get the service-worker Javascript for previews

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FilesApiAPI.FilesApiServiceWorker(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesApiAPI.FilesApiServiceWorker``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FilesApiServiceWorker`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `FilesApiAPI.FilesApiServiceWorker`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFilesApiServiceWorkerRequest struct via the builder pattern


### Return type

[***os.File**](*os.File.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

