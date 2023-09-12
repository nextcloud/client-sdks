# \CorePreviewAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CorePreviewGetPreview**](CorePreviewAPI.md#CorePreviewGetPreview) | **Get** /index.php/core/preview.png | Get a preview by file path
[**CorePreviewGetPreviewByFileId**](CorePreviewAPI.md#CorePreviewGetPreviewByFileId) | **Get** /index.php/core/preview | Get a preview by file ID



## CorePreviewGetPreview

> *os.File CorePreviewGetPreview(ctx).File(file).X(x).Y(y).A(a).ForceIcon(forceIcon).Mode(mode).MimeFallback(mimeFallback).Execute()

Get a preview by file path

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
    file := "file_example" // string | Path of the file (optional) (default to "")
    x := int64(789) // int64 | Width of the preview (optional) (default to 32)
    y := int64(789) // int64 | Height of the preview (optional) (default to 32)
    a := int32(56) // int32 | Whether to not crop the preview (optional) (default to 0)
    forceIcon := int32(56) // int32 | Force returning an icon (optional) (default to 1)
    mode := "mode_example" // string | How to crop the image (optional) (default to "fill")
    mimeFallback := int32(56) // int32 | Whether to fallback to the mime icon if no preview is available (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorePreviewAPI.CorePreviewGetPreview(context.Background()).File(file).X(x).Y(y).A(a).ForceIcon(forceIcon).Mode(mode).MimeFallback(mimeFallback).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorePreviewAPI.CorePreviewGetPreview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CorePreviewGetPreview`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `CorePreviewAPI.CorePreviewGetPreview`: %v\n", resp)
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
 **a** | **int32** | Whether to not crop the preview | [default to 0]
 **forceIcon** | **int32** | Force returning an icon | [default to 1]
 **mode** | **string** | How to crop the image | [default to &quot;fill&quot;]
 **mimeFallback** | **int32** | Whether to fallback to the mime icon if no preview is available | [default to 0]

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

> *os.File CorePreviewGetPreviewByFileId(ctx).FileId(fileId).X(x).Y(y).A(a).ForceIcon(forceIcon).Mode(mode).MimeFallback(mimeFallback).Execute()

Get a preview by file ID

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
    fileId := int64(789) // int64 | ID of the file (optional) (default to -1)
    x := int64(789) // int64 | Width of the preview (optional) (default to 32)
    y := int64(789) // int64 | Height of the preview (optional) (default to 32)
    a := int32(56) // int32 | Whether to not crop the preview (optional) (default to 0)
    forceIcon := int32(56) // int32 | Force returning an icon (optional) (default to 1)
    mode := "mode_example" // string | How to crop the image (optional) (default to "fill")
    mimeFallback := int32(56) // int32 | Whether to fallback to the mime icon if no preview is available (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CorePreviewAPI.CorePreviewGetPreviewByFileId(context.Background()).FileId(fileId).X(x).Y(y).A(a).ForceIcon(forceIcon).Mode(mode).MimeFallback(mimeFallback).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorePreviewAPI.CorePreviewGetPreviewByFileId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CorePreviewGetPreviewByFileId`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `CorePreviewAPI.CorePreviewGetPreviewByFileId`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCorePreviewGetPreviewByFileIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fileId** | **int64** | ID of the file | [default to -1]
 **x** | **int64** | Width of the preview | [default to 32]
 **y** | **int64** | Height of the preview | [default to 32]
 **a** | **int32** | Whether to not crop the preview | [default to 0]
 **forceIcon** | **int32** | Force returning an icon | [default to 1]
 **mode** | **string** | How to crop the image | [default to &quot;fill&quot;]
 **mimeFallback** | **int32** | Whether to fallback to the mime icon if no preview is available | [default to 0]

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

