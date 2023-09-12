# \UserStatusHeartbeatAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserStatusHeartbeatHeartbeat**](UserStatusHeartbeatAPI.md#UserStatusHeartbeatHeartbeat) | **Put** /ocs/v2.php/apps/user_status/api/v1/heartbeat | Keep the status alive



## UserStatusHeartbeatHeartbeat

> UserStatusUserStatusGetStatus200Response UserStatusHeartbeatHeartbeat(ctx).Status(status).OCSAPIRequest(oCSAPIRequest).Execute()

Keep the status alive

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
    status := "status_example" // string | Only online, away
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserStatusHeartbeatAPI.UserStatusHeartbeatHeartbeat(context.Background()).Status(status).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserStatusHeartbeatAPI.UserStatusHeartbeatHeartbeat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserStatusHeartbeatHeartbeat`: UserStatusUserStatusGetStatus200Response
    fmt.Fprintf(os.Stdout, "Response from `UserStatusHeartbeatAPI.UserStatusHeartbeatHeartbeat`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserStatusHeartbeatHeartbeatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | Only online, away | 
 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]

### Return type

[**UserStatusUserStatusGetStatus200Response**](UserStatusUserStatusGetStatus200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

