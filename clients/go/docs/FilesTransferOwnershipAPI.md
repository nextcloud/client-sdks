# \FilesTransferOwnershipAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FilesTransferOwnershipAccept**](FilesTransferOwnershipAPI.md#FilesTransferOwnershipAccept) | **Post** /ocs/v2.php/apps/files/api/v1/transferownership/{id} | Accept an ownership transfer
[**FilesTransferOwnershipReject**](FilesTransferOwnershipAPI.md#FilesTransferOwnershipReject) | **Delete** /ocs/v2.php/apps/files/api/v1/transferownership/{id} | Reject an ownership transfer
[**FilesTransferOwnershipTransfer**](FilesTransferOwnershipAPI.md#FilesTransferOwnershipTransfer) | **Post** /ocs/v2.php/apps/files/api/v1/transferownership | Transfer the ownership to another user



## FilesTransferOwnershipAccept

> CoreWhatsNewDismiss200Response FilesTransferOwnershipAccept(ctx, id).OCSAPIRequest(oCSAPIRequest).Execute()

Accept an ownership transfer

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
    id := int64(789) // int64 | ID of the ownership transfer
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FilesTransferOwnershipAPI.FilesTransferOwnershipAccept(context.Background(), id).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesTransferOwnershipAPI.FilesTransferOwnershipAccept``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FilesTransferOwnershipAccept`: CoreWhatsNewDismiss200Response
    fmt.Fprintf(os.Stdout, "Response from `FilesTransferOwnershipAPI.FilesTransferOwnershipAccept`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | ID of the ownership transfer | 

### Other Parameters

Other parameters are passed through a pointer to a apiFilesTransferOwnershipAcceptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]

### Return type

[**CoreWhatsNewDismiss200Response**](CoreWhatsNewDismiss200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FilesTransferOwnershipReject

> CoreWhatsNewDismiss200Response FilesTransferOwnershipReject(ctx, id).OCSAPIRequest(oCSAPIRequest).Execute()

Reject an ownership transfer

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
    id := int64(789) // int64 | ID of the ownership transfer
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FilesTransferOwnershipAPI.FilesTransferOwnershipReject(context.Background(), id).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesTransferOwnershipAPI.FilesTransferOwnershipReject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FilesTransferOwnershipReject`: CoreWhatsNewDismiss200Response
    fmt.Fprintf(os.Stdout, "Response from `FilesTransferOwnershipAPI.FilesTransferOwnershipReject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | ID of the ownership transfer | 

### Other Parameters

Other parameters are passed through a pointer to a apiFilesTransferOwnershipRejectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]

### Return type

[**CoreWhatsNewDismiss200Response**](CoreWhatsNewDismiss200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FilesTransferOwnershipTransfer

> CoreWhatsNewDismiss200Response FilesTransferOwnershipTransfer(ctx).Recipient(recipient).Path(path).OCSAPIRequest(oCSAPIRequest).Execute()

Transfer the ownership to another user

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
    recipient := "recipient_example" // string | Username of the recipient
    path := "path_example" // string | Path of the file
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FilesTransferOwnershipAPI.FilesTransferOwnershipTransfer(context.Background()).Recipient(recipient).Path(path).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesTransferOwnershipAPI.FilesTransferOwnershipTransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FilesTransferOwnershipTransfer`: CoreWhatsNewDismiss200Response
    fmt.Fprintf(os.Stdout, "Response from `FilesTransferOwnershipAPI.FilesTransferOwnershipTransfer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFilesTransferOwnershipTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **recipient** | **string** | Username of the recipient | 
 **path** | **string** | Path of the file | 
 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]

### Return type

[**CoreWhatsNewDismiss200Response**](CoreWhatsNewDismiss200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

