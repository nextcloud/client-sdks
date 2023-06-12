# \FilesSharingRemoteApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FilesSharingRemoteAcceptShare**](FilesSharingRemoteApi.md#FilesSharingRemoteAcceptShare) | **Post** /ocs/v2.php/apps/files_sharing/api/v1/remote_shares/pending/{id} | Accept a remote share
[**FilesSharingRemoteDeclineShare**](FilesSharingRemoteApi.md#FilesSharingRemoteDeclineShare) | **Delete** /ocs/v2.php/apps/files_sharing/api/v1/remote_shares/pending/{id} | Decline a remote share
[**FilesSharingRemoteGetOpenShares**](FilesSharingRemoteApi.md#FilesSharingRemoteGetOpenShares) | **Get** /ocs/v2.php/apps/files_sharing/api/v1/remote_shares/pending | Get list of pending remote shares
[**FilesSharingRemoteGetShare**](FilesSharingRemoteApi.md#FilesSharingRemoteGetShare) | **Get** /ocs/v2.php/apps/files_sharing/api/v1/remote_shares/{id} | Get info of a remote share
[**FilesSharingRemoteGetShares**](FilesSharingRemoteApi.md#FilesSharingRemoteGetShares) | **Get** /ocs/v2.php/apps/files_sharing/api/v1/remote_shares | Get a list of accepted remote shares
[**FilesSharingRemoteUnshare**](FilesSharingRemoteApi.md#FilesSharingRemoteUnshare) | **Delete** /ocs/v2.php/apps/files_sharing/api/v1/remote_shares/{id} | Unshare a remote share



## FilesSharingRemoteAcceptShare

> CoreWhatsNewDismiss200Response FilesSharingRemoteAcceptShare(ctx, id).OCSAPIRequest(oCSAPIRequest).Execute()

Accept a remote share

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
    id := int64(789) // int64 | ID of the share
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FilesSharingRemoteApi.FilesSharingRemoteAcceptShare(context.Background(), id).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesSharingRemoteApi.FilesSharingRemoteAcceptShare``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FilesSharingRemoteAcceptShare`: CoreWhatsNewDismiss200Response
    fmt.Fprintf(os.Stdout, "Response from `FilesSharingRemoteApi.FilesSharingRemoteAcceptShare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | ID of the share | 

### Other Parameters

Other parameters are passed through a pointer to a apiFilesSharingRemoteAcceptShareRequest struct via the builder pattern


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


## FilesSharingRemoteDeclineShare

> CoreWhatsNewDismiss200Response FilesSharingRemoteDeclineShare(ctx, id).OCSAPIRequest(oCSAPIRequest).Execute()

Decline a remote share

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
    id := int64(789) // int64 | ID of the share
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FilesSharingRemoteApi.FilesSharingRemoteDeclineShare(context.Background(), id).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesSharingRemoteApi.FilesSharingRemoteDeclineShare``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FilesSharingRemoteDeclineShare`: CoreWhatsNewDismiss200Response
    fmt.Fprintf(os.Stdout, "Response from `FilesSharingRemoteApi.FilesSharingRemoteDeclineShare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | ID of the share | 

### Other Parameters

Other parameters are passed through a pointer to a apiFilesSharingRemoteDeclineShareRequest struct via the builder pattern


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


## FilesSharingRemoteGetOpenShares

> FilesSharingRemoteGetShares200Response FilesSharingRemoteGetOpenShares(ctx).OCSAPIRequest(oCSAPIRequest).Execute()

Get list of pending remote shares

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
    resp, r, err := apiClient.FilesSharingRemoteApi.FilesSharingRemoteGetOpenShares(context.Background()).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesSharingRemoteApi.FilesSharingRemoteGetOpenShares``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FilesSharingRemoteGetOpenShares`: FilesSharingRemoteGetShares200Response
    fmt.Fprintf(os.Stdout, "Response from `FilesSharingRemoteApi.FilesSharingRemoteGetOpenShares`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFilesSharingRemoteGetOpenSharesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]

### Return type

[**FilesSharingRemoteGetShares200Response**](FilesSharingRemoteGetShares200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FilesSharingRemoteGetShare

> FilesSharingRemoteGetShare200Response FilesSharingRemoteGetShare(ctx, id).OCSAPIRequest(oCSAPIRequest).Execute()

Get info of a remote share

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
    id := int64(789) // int64 | ID of the share
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FilesSharingRemoteApi.FilesSharingRemoteGetShare(context.Background(), id).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesSharingRemoteApi.FilesSharingRemoteGetShare``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FilesSharingRemoteGetShare`: FilesSharingRemoteGetShare200Response
    fmt.Fprintf(os.Stdout, "Response from `FilesSharingRemoteApi.FilesSharingRemoteGetShare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | ID of the share | 

### Other Parameters

Other parameters are passed through a pointer to a apiFilesSharingRemoteGetShareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]

### Return type

[**FilesSharingRemoteGetShare200Response**](FilesSharingRemoteGetShare200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FilesSharingRemoteGetShares

> FilesSharingRemoteGetShares200Response FilesSharingRemoteGetShares(ctx).OCSAPIRequest(oCSAPIRequest).Execute()

Get a list of accepted remote shares

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
    resp, r, err := apiClient.FilesSharingRemoteApi.FilesSharingRemoteGetShares(context.Background()).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesSharingRemoteApi.FilesSharingRemoteGetShares``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FilesSharingRemoteGetShares`: FilesSharingRemoteGetShares200Response
    fmt.Fprintf(os.Stdout, "Response from `FilesSharingRemoteApi.FilesSharingRemoteGetShares`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFilesSharingRemoteGetSharesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]

### Return type

[**FilesSharingRemoteGetShares200Response**](FilesSharingRemoteGetShares200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FilesSharingRemoteUnshare

> CoreWhatsNewDismiss200Response FilesSharingRemoteUnshare(ctx, id).OCSAPIRequest(oCSAPIRequest).Execute()

Unshare a remote share

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
    id := int64(789) // int64 | ID of the share
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FilesSharingRemoteApi.FilesSharingRemoteUnshare(context.Background(), id).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesSharingRemoteApi.FilesSharingRemoteUnshare``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FilesSharingRemoteUnshare`: CoreWhatsNewDismiss200Response
    fmt.Fprintf(os.Stdout, "Response from `FilesSharingRemoteApi.FilesSharingRemoteUnshare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | ID of the share | 

### Other Parameters

Other parameters are passed through a pointer to a apiFilesSharingRemoteUnshareRequest struct via the builder pattern


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

