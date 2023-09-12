# \FilesVersionsPreviewAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FilesVersionsPreviewGetPreview**](FilesVersionsPreviewAPI.md#FilesVersionsPreviewGetPreview) | **Get** /index.php/apps/files_versions/preview | Get the preview for a file version



## FilesVersionsPreviewGetPreview

> *os.File FilesVersionsPreviewGetPreview(ctx).File(file).X(x).Y(y).Version(version).Execute()

Get the preview for a file version

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
    x := int64(789) // int64 | Width of the preview (optional) (default to 44)
    y := int64(789) // int64 | Height of the preview (optional) (default to 44)
    version := "version_example" // string | Version of the file to get the preview for (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FilesVersionsPreviewAPI.FilesVersionsPreviewGetPreview(context.Background()).File(file).X(x).Y(y).Version(version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesVersionsPreviewAPI.FilesVersionsPreviewGetPreview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FilesVersionsPreviewGetPreview`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `FilesVersionsPreviewAPI.FilesVersionsPreviewGetPreview`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFilesVersionsPreviewGetPreviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | **string** | Path of the file | [default to &quot;&quot;]
 **x** | **int64** | Width of the preview | [default to 44]
 **y** | **int64** | Height of the preview | [default to 44]
 **version** | **string** | Version of the file to get the preview for | [default to &quot;&quot;]

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

