# \FilesTrashbinPreviewAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FilesTrashbinPreviewGetPreview**](FilesTrashbinPreviewAPI.md#FilesTrashbinPreviewGetPreview) | **Get** /index.php/apps/files_trashbin/preview | Get the preview for a file



## FilesTrashbinPreviewGetPreview

> *os.File FilesTrashbinPreviewGetPreview(ctx).FileId(fileId).X(x).Y(y).A(a).Execute()

Get the preview for a file

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FilesTrashbinPreviewAPI.FilesTrashbinPreviewGetPreview(context.Background()).FileId(fileId).X(x).Y(y).A(a).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesTrashbinPreviewAPI.FilesTrashbinPreviewGetPreview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FilesTrashbinPreviewGetPreview`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `FilesTrashbinPreviewAPI.FilesTrashbinPreviewGetPreview`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFilesTrashbinPreviewGetPreviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fileId** | **int64** | ID of the file | [default to -1]
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

