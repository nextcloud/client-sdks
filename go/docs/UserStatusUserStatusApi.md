# \UserStatusUserStatusApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserStatusUserStatusClearMessage**](UserStatusUserStatusApi.md#UserStatusUserStatusClearMessage) | **Delete** /ocs/v2.php/apps/user_status/api/v1/user_status/message | Clear the message of the current user
[**UserStatusUserStatusGetStatus**](UserStatusUserStatusApi.md#UserStatusUserStatusGetStatus) | **Get** /ocs/v2.php/apps/user_status/api/v1/user_status | Get the status of the current user
[**UserStatusUserStatusRevertStatus**](UserStatusUserStatusApi.md#UserStatusUserStatusRevertStatus) | **Delete** /ocs/v2.php/apps/user_status/api/v1/user_status/revert/{messageId} | Revert the status to the previous status
[**UserStatusUserStatusSetCustomMessage**](UserStatusUserStatusApi.md#UserStatusUserStatusSetCustomMessage) | **Put** /ocs/v2.php/apps/user_status/api/v1/user_status/message/custom | Set the message to a custom message for the current user
[**UserStatusUserStatusSetPredefinedMessage**](UserStatusUserStatusApi.md#UserStatusUserStatusSetPredefinedMessage) | **Put** /ocs/v2.php/apps/user_status/api/v1/user_status/message/predefined | Set the message to a predefined message for the current user
[**UserStatusUserStatusSetStatus**](UserStatusUserStatusApi.md#UserStatusUserStatusSetStatus) | **Put** /ocs/v2.php/apps/user_status/api/v1/user_status/status | Update the status type of the current user



## UserStatusUserStatusClearMessage

> CoreWhatsNewDismiss200Response UserStatusUserStatusClearMessage(ctx).OCSAPIRequest(oCSAPIRequest).Execute()

Clear the message of the current user

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
    resp, r, err := apiClient.UserStatusUserStatusApi.UserStatusUserStatusClearMessage(context.Background()).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserStatusUserStatusApi.UserStatusUserStatusClearMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserStatusUserStatusClearMessage`: CoreWhatsNewDismiss200Response
    fmt.Fprintf(os.Stdout, "Response from `UserStatusUserStatusApi.UserStatusUserStatusClearMessage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserStatusUserStatusClearMessageRequest struct via the builder pattern


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


## UserStatusUserStatusGetStatus

> UserStatusUserStatusGetStatus200Response UserStatusUserStatusGetStatus(ctx).OCSAPIRequest(oCSAPIRequest).Execute()

Get the status of the current user

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
    resp, r, err := apiClient.UserStatusUserStatusApi.UserStatusUserStatusGetStatus(context.Background()).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserStatusUserStatusApi.UserStatusUserStatusGetStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserStatusUserStatusGetStatus`: UserStatusUserStatusGetStatus200Response
    fmt.Fprintf(os.Stdout, "Response from `UserStatusUserStatusApi.UserStatusUserStatusGetStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserStatusUserStatusGetStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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


## UserStatusUserStatusRevertStatus

> UserStatusUserStatusGetStatus200Response UserStatusUserStatusRevertStatus(ctx, messageId).OCSAPIRequest(oCSAPIRequest).Execute()

Revert the status to the previous status

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
    messageId := "messageId_example" // string | ID of the message to delete
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserStatusUserStatusApi.UserStatusUserStatusRevertStatus(context.Background(), messageId).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserStatusUserStatusApi.UserStatusUserStatusRevertStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserStatusUserStatusRevertStatus`: UserStatusUserStatusGetStatus200Response
    fmt.Fprintf(os.Stdout, "Response from `UserStatusUserStatusApi.UserStatusUserStatusRevertStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** | ID of the message to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserStatusUserStatusRevertStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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


## UserStatusUserStatusSetCustomMessage

> UserStatusUserStatusGetStatus200Response UserStatusUserStatusSetCustomMessage(ctx).OCSAPIRequest(oCSAPIRequest).StatusIcon(statusIcon).Message(message).ClearAt(clearAt).Execute()

Set the message to a custom message for the current user

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
    statusIcon := "statusIcon_example" // string | Icon of the status (optional)
    message := "message_example" // string | Message of the status (optional)
    clearAt := int64(789) // int64 | When the message should be cleared (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserStatusUserStatusApi.UserStatusUserStatusSetCustomMessage(context.Background()).OCSAPIRequest(oCSAPIRequest).StatusIcon(statusIcon).Message(message).ClearAt(clearAt).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserStatusUserStatusApi.UserStatusUserStatusSetCustomMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserStatusUserStatusSetCustomMessage`: UserStatusUserStatusGetStatus200Response
    fmt.Fprintf(os.Stdout, "Response from `UserStatusUserStatusApi.UserStatusUserStatusSetCustomMessage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserStatusUserStatusSetCustomMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]
 **statusIcon** | **string** | Icon of the status | 
 **message** | **string** | Message of the status | 
 **clearAt** | **int64** | When the message should be cleared | 

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


## UserStatusUserStatusSetPredefinedMessage

> UserStatusUserStatusGetStatus200Response UserStatusUserStatusSetPredefinedMessage(ctx).MessageId(messageId).OCSAPIRequest(oCSAPIRequest).ClearAt(clearAt).Execute()

Set the message to a predefined message for the current user

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
    messageId := "messageId_example" // string | ID of the predefined message
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")
    clearAt := int64(789) // int64 | When the message should be cleared (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserStatusUserStatusApi.UserStatusUserStatusSetPredefinedMessage(context.Background()).MessageId(messageId).OCSAPIRequest(oCSAPIRequest).ClearAt(clearAt).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserStatusUserStatusApi.UserStatusUserStatusSetPredefinedMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserStatusUserStatusSetPredefinedMessage`: UserStatusUserStatusGetStatus200Response
    fmt.Fprintf(os.Stdout, "Response from `UserStatusUserStatusApi.UserStatusUserStatusSetPredefinedMessage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserStatusUserStatusSetPredefinedMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **messageId** | **string** | ID of the predefined message | 
 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]
 **clearAt** | **int64** | When the message should be cleared | 

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


## UserStatusUserStatusSetStatus

> UserStatusUserStatusGetStatus200Response UserStatusUserStatusSetStatus(ctx).StatusType(statusType).OCSAPIRequest(oCSAPIRequest).Execute()

Update the status type of the current user

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
    statusType := "statusType_example" // string | The new status type
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserStatusUserStatusApi.UserStatusUserStatusSetStatus(context.Background()).StatusType(statusType).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserStatusUserStatusApi.UserStatusUserStatusSetStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserStatusUserStatusSetStatus`: UserStatusUserStatusGetStatus200Response
    fmt.Fprintf(os.Stdout, "Response from `UserStatusUserStatusApi.UserStatusUserStatusSetStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserStatusUserStatusSetStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **statusType** | **string** | The new status type | 
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

