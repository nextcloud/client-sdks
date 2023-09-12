# \FilesSharingShareInfoAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FilesSharingShareInfoInfo**](FilesSharingShareInfoAPI.md#FilesSharingShareInfoInfo) | **Post** /index.php/apps/files_sharing/shareinfo | Get the info about a share



## FilesSharingShareInfoInfo

> FilesSharingShareInfo FilesSharingShareInfoInfo(ctx).T(t).Password(password).Dir(dir).Depth(depth).Execute()

Get the info about a share

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
    t := "t_example" // string | Token of the share
    password := "password_example" // string | Password of the share (optional)
    dir := "dir_example" // string | Subdirectory to get info about (optional)
    depth := int64(789) // int64 | Maximum depth to get info about (optional) (default to -1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FilesSharingShareInfoAPI.FilesSharingShareInfoInfo(context.Background()).T(t).Password(password).Dir(dir).Depth(depth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesSharingShareInfoAPI.FilesSharingShareInfoInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FilesSharingShareInfoInfo`: FilesSharingShareInfo
    fmt.Fprintf(os.Stdout, "Response from `FilesSharingShareInfoAPI.FilesSharingShareInfoInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFilesSharingShareInfoInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **t** | **string** | Token of the share | 
 **password** | **string** | Password of the share | 
 **dir** | **string** | Subdirectory to get info about | 
 **depth** | **int64** | Maximum depth to get info about | [default to -1]

### Return type

[**FilesSharingShareInfo**](FilesSharingShareInfo.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

