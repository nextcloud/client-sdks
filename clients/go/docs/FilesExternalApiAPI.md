# \FilesExternalApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FilesExternalApiGetUserMounts**](FilesExternalApiAPI.md#FilesExternalApiGetUserMounts) | **Get** /ocs/v2.php/apps/files_external/api/v1/mounts | Get the mount points visible for this user



## FilesExternalApiGetUserMounts

> FilesExternalApiGetUserMounts200Response FilesExternalApiGetUserMounts(ctx).OCSAPIRequest(oCSAPIRequest).Execute()

Get the mount points visible for this user

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
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FilesExternalApiAPI.FilesExternalApiGetUserMounts(context.Background()).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesExternalApiAPI.FilesExternalApiGetUserMounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FilesExternalApiGetUserMounts`: FilesExternalApiGetUserMounts200Response
    fmt.Fprintf(os.Stdout, "Response from `FilesExternalApiAPI.FilesExternalApiGetUserMounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFilesExternalApiGetUserMountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]

### Return type

[**FilesExternalApiGetUserMounts200Response**](FilesExternalApiGetUserMounts200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

