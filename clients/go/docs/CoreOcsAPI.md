# \CoreOcsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CoreOcsGetCapabilities**](CoreOcsAPI.md#CoreOcsGetCapabilities) | **Get** /ocs/v2.php/cloud/capabilities | Get the capabilities



## CoreOcsGetCapabilities

> CoreOcsGetCapabilities200Response CoreOcsGetCapabilities(ctx).OCSAPIRequest(oCSAPIRequest).Execute()

Get the capabilities

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
    resp, r, err := apiClient.CoreOcsAPI.CoreOcsGetCapabilities(context.Background()).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreOcsAPI.CoreOcsGetCapabilities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CoreOcsGetCapabilities`: CoreOcsGetCapabilities200Response
    fmt.Fprintf(os.Stdout, "Response from `CoreOcsAPI.CoreOcsGetCapabilities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCoreOcsGetCapabilitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]

### Return type

[**CoreOcsGetCapabilities200Response**](CoreOcsGetCapabilities200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

