# \UpdatenotificationApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdatenotificationApiGetAppList**](UpdatenotificationApiAPI.md#UpdatenotificationApiGetAppList) | **Get** /ocs/v2.php/apps/updatenotification/api/{apiVersion}/applist/{newVersion} | List available updates for apps



## UpdatenotificationApiGetAppList

> UpdatenotificationApiGetAppList200Response UpdatenotificationApiGetAppList(ctx, apiVersion, newVersion).OCSAPIRequest(oCSAPIRequest).Execute()

List available updates for apps



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
    apiVersion := "apiVersion_example" // string |  (default to "v1")
    newVersion := "newVersion_example" // string | Server version to check updates for
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UpdatenotificationApiAPI.UpdatenotificationApiGetAppList(context.Background(), apiVersion, newVersion).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UpdatenotificationApiAPI.UpdatenotificationApiGetAppList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatenotificationApiGetAppList`: UpdatenotificationApiGetAppList200Response
    fmt.Fprintf(os.Stdout, "Response from `UpdatenotificationApiAPI.UpdatenotificationApiGetAppList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** |  | [default to &quot;v1&quot;]
**newVersion** | **string** | Server version to check updates for | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatenotificationApiGetAppListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]

### Return type

[**UpdatenotificationApiGetAppList200Response**](UpdatenotificationApiGetAppList200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

