# \FilesSharingDeletedShareapiApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FilesSharingDeletedShareapiList**](FilesSharingDeletedShareapiApi.md#FilesSharingDeletedShareapiList) | **Get** /ocs/v2.php/apps/files_sharing/api/v1/deletedshares | Get a list of all deleted shares
[**FilesSharingDeletedShareapiUndelete**](FilesSharingDeletedShareapiApi.md#FilesSharingDeletedShareapiUndelete) | **Post** /ocs/v2.php/apps/files_sharing/api/v1/deletedshares/{id} | Undelete a deleted share



## FilesSharingDeletedShareapiList

> FilesSharingDeletedShareapiList200Response FilesSharingDeletedShareapiList(ctx).OCSAPIRequest(oCSAPIRequest).Execute()

Get a list of all deleted shares

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
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FilesSharingDeletedShareapiApi.FilesSharingDeletedShareapiList(context.Background()).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesSharingDeletedShareapiApi.FilesSharingDeletedShareapiList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FilesSharingDeletedShareapiList`: FilesSharingDeletedShareapiList200Response
    fmt.Fprintf(os.Stdout, "Response from `FilesSharingDeletedShareapiApi.FilesSharingDeletedShareapiList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFilesSharingDeletedShareapiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]

### Return type

[**FilesSharingDeletedShareapiList200Response**](FilesSharingDeletedShareapiList200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FilesSharingDeletedShareapiUndelete

> CoreWhatsNewDismiss200Response FilesSharingDeletedShareapiUndelete(ctx, id).OCSAPIRequest(oCSAPIRequest).Execute()

Undelete a deleted share

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
    id := "id_example" // string | ID of the share
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FilesSharingDeletedShareapiApi.FilesSharingDeletedShareapiUndelete(context.Background(), id).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesSharingDeletedShareapiApi.FilesSharingDeletedShareapiUndelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FilesSharingDeletedShareapiUndelete`: CoreWhatsNewDismiss200Response
    fmt.Fprintf(os.Stdout, "Response from `FilesSharingDeletedShareapiApi.FilesSharingDeletedShareapiUndelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the share | 

### Other Parameters

Other parameters are passed through a pointer to a apiFilesSharingDeletedShareapiUndeleteRequest struct via the builder pattern


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

