# \CorePreviewApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CorePreviewGetPreview**](CorePreviewApi.md#CorePreviewGetPreview) | **Get** /index.php/core/preview.png | Get a preview by file ID
[**CorePreviewGetPreviewByFileId**](CorePreviewApi.md#CorePreviewGetPreviewByFileId) | **Get** /index.php/core/preview | Get a preview by file ID



## CorePreviewGetPreview

> *os.File CorePreviewGetPreview(ctx).File(file).X(x).Y(y).A(a).ForceIcon(forceIcon).Mode(mode).Execute()

Get a preview by file ID

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
    file := "file_example" // string | Path of the file (optional) (default to "")
    x := int64(789) // int64 | Width of the preview (optional) (default to 32)
    y := int64(789) // int64 | Height of the preview (optional) (default to 32)
    a := int32(56) // int32 | Not crop the preview (optional) (default to 0)
    forceIcon := int32(56) // int32 | Force returning an icon (optional) (default to 1)
    mode := "mode_example" // string | How to crop the image (optional) (default to "fill")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorePreviewApi.CorePreviewGetPreview(context.Background()).File(file).X(x).Y(y).A(a).ForceIcon(forceIcon).Mode(mode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorePreviewApi.CorePreviewGetPreview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CorePreviewGetPreview`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `CorePreviewApi.CorePreviewGetPreview`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCorePreviewGetPreviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | **string** | Path of the file | [default to &quot;&quot;]
 **x** | **int64** | Width of the preview | [default to 32]
 **y** | **int64** | Height of the preview | [default to 32]
 **a** | **int32** | Not crop the preview | [default to 0]
 **forceIcon** | **int32** | Force returning an icon | [default to 1]
 **mode** | **string** | How to crop the image | [default to &quot;fill&quot;]

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


## CorePreviewGetPreviewByFileId

> *os.File CorePreviewGetPreviewByFileId(ctx).FileId(fileId).X(x).Y(y).A(a).ForceIcon(forceIcon).Mode(mode).Execute()

Get a preview by file ID

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
    fileId := int64(789) // int64 | ID of the file (optional) (default to 1)
    x := int64(789) // int64 | Width of the preview (optional) (default to 32)
    y := int64(789) // int64 | Height of the preview (optional) (default to 32)
    a := int32(56) // int32 | Not crop the preview (optional) (default to 0)
    forceIcon := int32(56) // int32 | Force returning an icon (optional) (default to 1)
    mode := "mode_example" // string | How to crop the image (optional) (default to "fill")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorePreviewApi.CorePreviewGetPreviewByFileId(context.Background()).FileId(fileId).X(x).Y(y).A(a).ForceIcon(forceIcon).Mode(mode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorePreviewApi.CorePreviewGetPreviewByFileId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CorePreviewGetPreviewByFileId`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `CorePreviewApi.CorePreviewGetPreviewByFileId`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCorePreviewGetPreviewByFileIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fileId** | **int64** | ID of the file | [default to 1]
 **x** | **int64** | Width of the preview | [default to 32]
 **y** | **int64** | Height of the preview | [default to 32]
 **a** | **int32** | Not crop the preview | [default to 0]
 **forceIcon** | **int32** | Force returning an icon | [default to 1]
 **mode** | **string** | How to crop the image | [default to &quot;fill&quot;]

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

