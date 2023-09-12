# \UserStatusPredefinedStatusAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserStatusPredefinedStatusFindAll**](UserStatusPredefinedStatusAPI.md#UserStatusPredefinedStatusFindAll) | **Get** /ocs/v2.php/apps/user_status/api/v1/predefined_statuses | Get all predefined messages



## UserStatusPredefinedStatusFindAll

> UserStatusPredefinedStatusFindAll200Response UserStatusPredefinedStatusFindAll(ctx).OCSAPIRequest(oCSAPIRequest).Execute()

Get all predefined messages

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
    resp, r, err := apiClient.UserStatusPredefinedStatusAPI.UserStatusPredefinedStatusFindAll(context.Background()).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserStatusPredefinedStatusAPI.UserStatusPredefinedStatusFindAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserStatusPredefinedStatusFindAll`: UserStatusPredefinedStatusFindAll200Response
    fmt.Fprintf(os.Stdout, "Response from `UserStatusPredefinedStatusAPI.UserStatusPredefinedStatusFindAll`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserStatusPredefinedStatusFindAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]

### Return type

[**UserStatusPredefinedStatusFindAll200Response**](UserStatusPredefinedStatusFindAll200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

