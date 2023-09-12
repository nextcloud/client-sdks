# \CoreWhatsNewAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CoreWhatsNewDismiss**](CoreWhatsNewAPI.md#CoreWhatsNewDismiss) | **Post** /ocs/v2.php/core/whatsnew | Dismiss the changes
[**CoreWhatsNewGet**](CoreWhatsNewAPI.md#CoreWhatsNewGet) | **Get** /ocs/v2.php/core/whatsnew | Get the changes



## CoreWhatsNewDismiss

> CoreWhatsNewDismiss200Response CoreWhatsNewDismiss(ctx).Version(version).OCSAPIRequest(oCSAPIRequest).Execute()

Dismiss the changes

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
    version := "version_example" // string | Version to dismiss the changes for
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoreWhatsNewAPI.CoreWhatsNewDismiss(context.Background()).Version(version).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreWhatsNewAPI.CoreWhatsNewDismiss``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CoreWhatsNewDismiss`: CoreWhatsNewDismiss200Response
    fmt.Fprintf(os.Stdout, "Response from `CoreWhatsNewAPI.CoreWhatsNewDismiss`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCoreWhatsNewDismissRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **string** | Version to dismiss the changes for | 
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


## CoreWhatsNewGet

> CoreWhatsNewGet200Response CoreWhatsNewGet(ctx).OCSAPIRequest(oCSAPIRequest).Execute()

Get the changes

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
    resp, r, err := apiClient.CoreWhatsNewAPI.CoreWhatsNewGet(context.Background()).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreWhatsNewAPI.CoreWhatsNewGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CoreWhatsNewGet`: CoreWhatsNewGet200Response
    fmt.Fprintf(os.Stdout, "Response from `CoreWhatsNewAPI.CoreWhatsNewGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCoreWhatsNewGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]

### Return type

[**CoreWhatsNewGet200Response**](CoreWhatsNewGet200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

