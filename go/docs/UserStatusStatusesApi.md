# \UserStatusStatusesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserStatusStatusesFind**](UserStatusStatusesApi.md#UserStatusStatusesFind) | **Get** /ocs/v2.php/apps/user_status/api/v1/statuses/{userId} | Find the status of a user
[**UserStatusStatusesFindAll**](UserStatusStatusesApi.md#UserStatusStatusesFindAll) | **Get** /ocs/v2.php/apps/user_status/api/v1/statuses | Find statuses of users



## UserStatusStatusesFind

> UserStatusStatusesFind200Response UserStatusStatusesFind(ctx, userId).OCSAPIRequest(oCSAPIRequest).Execute()

Find the status of a user

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
    userId := "userId_example" // string | ID of the user
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserStatusStatusesApi.UserStatusStatusesFind(context.Background(), userId).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserStatusStatusesApi.UserStatusStatusesFind``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserStatusStatusesFind`: UserStatusStatusesFind200Response
    fmt.Fprintf(os.Stdout, "Response from `UserStatusStatusesApi.UserStatusStatusesFind`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of the user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserStatusStatusesFindRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]

### Return type

[**UserStatusStatusesFind200Response**](UserStatusStatusesFind200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserStatusStatusesFindAll

> UserStatusStatusesFindAll200Response UserStatusStatusesFindAll(ctx).OCSAPIRequest(oCSAPIRequest).Limit(limit).Offset(offset).Execute()

Find statuses of users

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
    limit := int64(789) // int64 | Maximum number of statuses to find (optional)
    offset := int64(789) // int64 | Offset for finding statuses (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserStatusStatusesApi.UserStatusStatusesFindAll(context.Background()).OCSAPIRequest(oCSAPIRequest).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserStatusStatusesApi.UserStatusStatusesFindAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserStatusStatusesFindAll`: UserStatusStatusesFindAll200Response
    fmt.Fprintf(os.Stdout, "Response from `UserStatusStatusesApi.UserStatusStatusesFindAll`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserStatusStatusesFindAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]
 **limit** | **int64** | Maximum number of statuses to find | 
 **offset** | **int64** | Offset for finding statuses | 

### Return type

[**UserStatusStatusesFindAll200Response**](UserStatusStatusesFindAll200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

