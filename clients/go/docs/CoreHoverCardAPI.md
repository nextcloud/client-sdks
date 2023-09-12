# \CoreHoverCardAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CoreHoverCardGetUser**](CoreHoverCardAPI.md#CoreHoverCardGetUser) | **Get** /ocs/v2.php/hovercard/v1/{userId} | Get the user details for a hovercard



## CoreHoverCardGetUser

> CoreHoverCardGetUser200Response CoreHoverCardGetUser(ctx, userId).OCSAPIRequest(oCSAPIRequest).Execute()

Get the user details for a hovercard

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
    userId := "userId_example" // string | ID of the user
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoreHoverCardAPI.CoreHoverCardGetUser(context.Background(), userId).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreHoverCardAPI.CoreHoverCardGetUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CoreHoverCardGetUser`: CoreHoverCardGetUser200Response
    fmt.Fprintf(os.Stdout, "Response from `CoreHoverCardAPI.CoreHoverCardGetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of the user | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreHoverCardGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]

### Return type

[**CoreHoverCardGetUser200Response**](CoreHoverCardGetUser200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

