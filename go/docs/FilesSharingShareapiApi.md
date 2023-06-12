# \FilesSharingShareapiApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FilesSharingShareapiAcceptShare**](FilesSharingShareapiApi.md#FilesSharingShareapiAcceptShare) | **Post** /ocs/v2.php/apps/files_sharing/api/v1/shares/pending/{id} | Accept a share
[**FilesSharingShareapiCreateShare**](FilesSharingShareapiApi.md#FilesSharingShareapiCreateShare) | **Post** /ocs/v2.php/apps/files_sharing/api/v1/shares | Create a share
[**FilesSharingShareapiDeleteShare**](FilesSharingShareapiApi.md#FilesSharingShareapiDeleteShare) | **Delete** /ocs/v2.php/apps/files_sharing/api/v1/shares/{id} | Delete a share
[**FilesSharingShareapiGetInheritedShares**](FilesSharingShareapiApi.md#FilesSharingShareapiGetInheritedShares) | **Get** /ocs/v2.php/apps/files_sharing/api/v1/shares/inherited | Get all shares relative to a file, including parent folders shares rights
[**FilesSharingShareapiGetShare**](FilesSharingShareapiApi.md#FilesSharingShareapiGetShare) | **Get** /ocs/v2.php/apps/files_sharing/api/v1/shares/{id} | Get a specific share by id
[**FilesSharingShareapiGetShares**](FilesSharingShareapiApi.md#FilesSharingShareapiGetShares) | **Get** /ocs/v2.php/apps/files_sharing/api/v1/shares | Get shares of the current user
[**FilesSharingShareapiPendingShares**](FilesSharingShareapiApi.md#FilesSharingShareapiPendingShares) | **Get** /ocs/v2.php/apps/files_sharing/api/v1/shares/pending | Get all shares that are still pending
[**FilesSharingShareapiUpdateShare**](FilesSharingShareapiApi.md#FilesSharingShareapiUpdateShare) | **Put** /ocs/v2.php/apps/files_sharing/api/v1/shares/{id} | Update a share



## FilesSharingShareapiAcceptShare

> CoreWhatsNewDismiss200Response FilesSharingShareapiAcceptShare(ctx, id).OCSAPIRequest(oCSAPIRequest).Execute()

Accept a share

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
    resp, r, err := apiClient.FilesSharingShareapiApi.FilesSharingShareapiAcceptShare(context.Background(), id).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesSharingShareapiApi.FilesSharingShareapiAcceptShare``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FilesSharingShareapiAcceptShare`: CoreWhatsNewDismiss200Response
    fmt.Fprintf(os.Stdout, "Response from `FilesSharingShareapiApi.FilesSharingShareapiAcceptShare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the share | 

### Other Parameters

Other parameters are passed through a pointer to a apiFilesSharingShareapiAcceptShareRequest struct via the builder pattern


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


## FilesSharingShareapiCreateShare

> FilesSharingShareapiCreateShare200Response FilesSharingShareapiCreateShare(ctx).OCSAPIRequest(oCSAPIRequest).Path(path).Permissions(permissions).ShareType(shareType).ShareWith(shareWith).PublicUpload(publicUpload).Password(password).SendPasswordByTalk(sendPasswordByTalk).ExpireDate(expireDate).Note(note).Label(label).Attributes(attributes).Execute()

Create a share

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
    path := "path_example" // string | Path of the share (optional)
    permissions := int64(789) // int64 | Permissions for the share (optional)
    shareType := int64(789) // int64 | Type of the share (optional) (default to 1)
    shareWith := "shareWith_example" // string | The entity this should be shared with (optional)
    publicUpload := "publicUpload_example" // string | If public uploading is allowed (optional) (default to "false")
    password := "password_example" // string | Password for the share (optional) (default to "")
    sendPasswordByTalk := "sendPasswordByTalk_example" // string | Send the password for the share over Talk (optional)
    expireDate := "expireDate_example" // string | Expiry date of the share (optional) (default to "")
    note := "note_example" // string | Note for the share (optional) (default to "")
    label := "label_example" // string | Label for the share (only used in link and email) (optional) (default to "")
    attributes := "attributes_example" // string | Additional attributes for the share (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FilesSharingShareapiApi.FilesSharingShareapiCreateShare(context.Background()).OCSAPIRequest(oCSAPIRequest).Path(path).Permissions(permissions).ShareType(shareType).ShareWith(shareWith).PublicUpload(publicUpload).Password(password).SendPasswordByTalk(sendPasswordByTalk).ExpireDate(expireDate).Note(note).Label(label).Attributes(attributes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesSharingShareapiApi.FilesSharingShareapiCreateShare``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FilesSharingShareapiCreateShare`: FilesSharingShareapiCreateShare200Response
    fmt.Fprintf(os.Stdout, "Response from `FilesSharingShareapiApi.FilesSharingShareapiCreateShare`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFilesSharingShareapiCreateShareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]
 **path** | **string** | Path of the share | 
 **permissions** | **int64** | Permissions for the share | 
 **shareType** | **int64** | Type of the share | [default to 1]
 **shareWith** | **string** | The entity this should be shared with | 
 **publicUpload** | **string** | If public uploading is allowed | [default to &quot;false&quot;]
 **password** | **string** | Password for the share | [default to &quot;&quot;]
 **sendPasswordByTalk** | **string** | Send the password for the share over Talk | 
 **expireDate** | **string** | Expiry date of the share | [default to &quot;&quot;]
 **note** | **string** | Note for the share | [default to &quot;&quot;]
 **label** | **string** | Label for the share (only used in link and email) | [default to &quot;&quot;]
 **attributes** | **string** | Additional attributes for the share | 

### Return type

[**FilesSharingShareapiCreateShare200Response**](FilesSharingShareapiCreateShare200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FilesSharingShareapiDeleteShare

> CoreWhatsNewDismiss200Response FilesSharingShareapiDeleteShare(ctx, id).OCSAPIRequest(oCSAPIRequest).Execute()

Delete a share

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
    resp, r, err := apiClient.FilesSharingShareapiApi.FilesSharingShareapiDeleteShare(context.Background(), id).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesSharingShareapiApi.FilesSharingShareapiDeleteShare``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FilesSharingShareapiDeleteShare`: CoreWhatsNewDismiss200Response
    fmt.Fprintf(os.Stdout, "Response from `FilesSharingShareapiApi.FilesSharingShareapiDeleteShare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the share | 

### Other Parameters

Other parameters are passed through a pointer to a apiFilesSharingShareapiDeleteShareRequest struct via the builder pattern


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


## FilesSharingShareapiGetInheritedShares

> FilesSharingShareapiGetShares200Response FilesSharingShareapiGetInheritedShares(ctx).Path(path).OCSAPIRequest(oCSAPIRequest).Execute()

Get all shares relative to a file, including parent folders shares rights

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
    path := "path_example" // string | Path all shares will be relative to
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FilesSharingShareapiApi.FilesSharingShareapiGetInheritedShares(context.Background()).Path(path).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesSharingShareapiApi.FilesSharingShareapiGetInheritedShares``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FilesSharingShareapiGetInheritedShares`: FilesSharingShareapiGetShares200Response
    fmt.Fprintf(os.Stdout, "Response from `FilesSharingShareapiApi.FilesSharingShareapiGetInheritedShares`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFilesSharingShareapiGetInheritedSharesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | Path all shares will be relative to | 
 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]

### Return type

[**FilesSharingShareapiGetShares200Response**](FilesSharingShareapiGetShares200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FilesSharingShareapiGetShare

> FilesSharingShareapiCreateShare200Response FilesSharingShareapiGetShare(ctx, id).OCSAPIRequest(oCSAPIRequest).IncludeTags(includeTags).Execute()

Get a specific share by id

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
    includeTags := int32(56) // int32 | Include tags in the share (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FilesSharingShareapiApi.FilesSharingShareapiGetShare(context.Background(), id).OCSAPIRequest(oCSAPIRequest).IncludeTags(includeTags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesSharingShareapiApi.FilesSharingShareapiGetShare``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FilesSharingShareapiGetShare`: FilesSharingShareapiCreateShare200Response
    fmt.Fprintf(os.Stdout, "Response from `FilesSharingShareapiApi.FilesSharingShareapiGetShare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the share | 

### Other Parameters

Other parameters are passed through a pointer to a apiFilesSharingShareapiGetShareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]
 **includeTags** | **int32** | Include tags in the share | [default to 0]

### Return type

[**FilesSharingShareapiCreateShare200Response**](FilesSharingShareapiCreateShare200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FilesSharingShareapiGetShares

> FilesSharingShareapiGetShares200Response FilesSharingShareapiGetShares(ctx).OCSAPIRequest(oCSAPIRequest).SharedWithMe(sharedWithMe).Reshares(reshares).Subfiles(subfiles).Path(path).IncludeTags(includeTags).Execute()

Get shares of the current user

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
    sharedWithMe := "sharedWithMe_example" // string | Only get shares with the current user (optional) (default to "false")
    reshares := "reshares_example" // string | Only get shares by the current user and reshares (optional) (default to "false")
    subfiles := "subfiles_example" // string | Only get all shares in a folder (optional) (default to "false")
    path := "path_example" // string | Get shares for a specific path (optional) (default to "")
    includeTags := "includeTags_example" // string | Include tags in the share (optional) (default to "false")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FilesSharingShareapiApi.FilesSharingShareapiGetShares(context.Background()).OCSAPIRequest(oCSAPIRequest).SharedWithMe(sharedWithMe).Reshares(reshares).Subfiles(subfiles).Path(path).IncludeTags(includeTags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesSharingShareapiApi.FilesSharingShareapiGetShares``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FilesSharingShareapiGetShares`: FilesSharingShareapiGetShares200Response
    fmt.Fprintf(os.Stdout, "Response from `FilesSharingShareapiApi.FilesSharingShareapiGetShares`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFilesSharingShareapiGetSharesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]
 **sharedWithMe** | **string** | Only get shares with the current user | [default to &quot;false&quot;]
 **reshares** | **string** | Only get shares by the current user and reshares | [default to &quot;false&quot;]
 **subfiles** | **string** | Only get all shares in a folder | [default to &quot;false&quot;]
 **path** | **string** | Get shares for a specific path | [default to &quot;&quot;]
 **includeTags** | **string** | Include tags in the share | [default to &quot;false&quot;]

### Return type

[**FilesSharingShareapiGetShares200Response**](FilesSharingShareapiGetShares200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FilesSharingShareapiPendingShares

> FilesSharingShareapiGetShares200Response FilesSharingShareapiPendingShares(ctx).OCSAPIRequest(oCSAPIRequest).Execute()

Get all shares that are still pending

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
    resp, r, err := apiClient.FilesSharingShareapiApi.FilesSharingShareapiPendingShares(context.Background()).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesSharingShareapiApi.FilesSharingShareapiPendingShares``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FilesSharingShareapiPendingShares`: FilesSharingShareapiGetShares200Response
    fmt.Fprintf(os.Stdout, "Response from `FilesSharingShareapiApi.FilesSharingShareapiPendingShares`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFilesSharingShareapiPendingSharesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]

### Return type

[**FilesSharingShareapiGetShares200Response**](FilesSharingShareapiGetShares200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FilesSharingShareapiUpdateShare

> FilesSharingShareapiCreateShare200Response FilesSharingShareapiUpdateShare(ctx, id).OCSAPIRequest(oCSAPIRequest).Permissions(permissions).Password(password).SendPasswordByTalk(sendPasswordByTalk).PublicUpload(publicUpload).ExpireDate(expireDate).Note(note).Label(label).HideDownload(hideDownload).Attributes(attributes).Execute()

Update a share

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
    permissions := int64(789) // int64 | New permissions (optional)
    password := "password_example" // string | New password (optional)
    sendPasswordByTalk := "sendPasswordByTalk_example" // string | New condition if the password should be send over Talk (optional)
    publicUpload := "publicUpload_example" // string | New condition if public uploading is allowed (optional)
    expireDate := "expireDate_example" // string | New expiry date (optional)
    note := "note_example" // string | New note (optional)
    label := "label_example" // string | New label (optional)
    hideDownload := "hideDownload_example" // string | New condition if the download should be hidden (optional)
    attributes := "attributes_example" // string | New additional attributes (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FilesSharingShareapiApi.FilesSharingShareapiUpdateShare(context.Background(), id).OCSAPIRequest(oCSAPIRequest).Permissions(permissions).Password(password).SendPasswordByTalk(sendPasswordByTalk).PublicUpload(publicUpload).ExpireDate(expireDate).Note(note).Label(label).HideDownload(hideDownload).Attributes(attributes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesSharingShareapiApi.FilesSharingShareapiUpdateShare``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FilesSharingShareapiUpdateShare`: FilesSharingShareapiCreateShare200Response
    fmt.Fprintf(os.Stdout, "Response from `FilesSharingShareapiApi.FilesSharingShareapiUpdateShare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the share | 

### Other Parameters

Other parameters are passed through a pointer to a apiFilesSharingShareapiUpdateShareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]
 **permissions** | **int64** | New permissions | 
 **password** | **string** | New password | 
 **sendPasswordByTalk** | **string** | New condition if the password should be send over Talk | 
 **publicUpload** | **string** | New condition if public uploading is allowed | 
 **expireDate** | **string** | New expiry date | 
 **note** | **string** | New note | 
 **label** | **string** | New label | 
 **hideDownload** | **string** | New condition if the download should be hidden | 
 **attributes** | **string** | New additional attributes | 

### Return type

[**FilesSharingShareapiCreateShare200Response**](FilesSharingShareapiCreateShare200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

