# \FilesSharingPublicPreviewAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FilesSharingPublicPreviewDirectLink**](FilesSharingPublicPreviewAPI.md#FilesSharingPublicPreviewDirectLink) | **Get** /index.php/s/{token}/preview | Get a direct link preview for a shared file
[**FilesSharingPublicPreviewGetPreview**](FilesSharingPublicPreviewAPI.md#FilesSharingPublicPreviewGetPreview) | **Get** /index.php/apps/files_sharing/publicpreview/{token} | Get a preview for a shared file



## FilesSharingPublicPreviewDirectLink

> *os.File FilesSharingPublicPreviewDirectLink(ctx, token).OCSAPIRequest(oCSAPIRequest).Execute()

Get a direct link preview for a shared file

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
    token := "token_example" // string | Token of the share
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FilesSharingPublicPreviewAPI.FilesSharingPublicPreviewDirectLink(context.Background(), token).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesSharingPublicPreviewAPI.FilesSharingPublicPreviewDirectLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FilesSharingPublicPreviewDirectLink`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `FilesSharingPublicPreviewAPI.FilesSharingPublicPreviewDirectLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Token of the share | 

### Other Parameters

Other parameters are passed through a pointer to a apiFilesSharingPublicPreviewDirectLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]

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


## FilesSharingPublicPreviewGetPreview

> *os.File FilesSharingPublicPreviewGetPreview(ctx, token).OCSAPIRequest(oCSAPIRequest).File(file).X(x).Y(y).A(a).Execute()

Get a preview for a shared file

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
    token := "token_example" // string | Token of the share
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")
    file := "file_example" // string | File in the share (optional) (default to "")
    x := int64(789) // int64 | Width of the preview (optional) (default to 32)
    y := int64(789) // int64 | Height of the preview (optional) (default to 32)
    a := int32(56) // int32 | Whether to not crop the preview (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FilesSharingPublicPreviewAPI.FilesSharingPublicPreviewGetPreview(context.Background(), token).OCSAPIRequest(oCSAPIRequest).File(file).X(x).Y(y).A(a).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesSharingPublicPreviewAPI.FilesSharingPublicPreviewGetPreview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FilesSharingPublicPreviewGetPreview`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `FilesSharingPublicPreviewAPI.FilesSharingPublicPreviewGetPreview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Token of the share | 

### Other Parameters

Other parameters are passed through a pointer to a apiFilesSharingPublicPreviewGetPreviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]
 **file** | **string** | File in the share | [default to &quot;&quot;]
 **x** | **int64** | Width of the preview | [default to 32]
 **y** | **int64** | Height of the preview | [default to 32]
 **a** | **int32** | Whether to not crop the preview | [default to 0]

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

