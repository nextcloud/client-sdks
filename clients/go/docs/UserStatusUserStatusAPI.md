# \UserStatusUserStatusAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserStatusUserStatusClearMessage**](UserStatusUserStatusAPI.md#UserStatusUserStatusClearMessage) | **Delete** /ocs/v2.php/apps/user_status/api/v1/user_status/message | Clear the message of the current user
[**UserStatusUserStatusGetStatus**](UserStatusUserStatusAPI.md#UserStatusUserStatusGetStatus) | **Get** /ocs/v2.php/apps/user_status/api/v1/user_status | Get the status of the current user
[**UserStatusUserStatusRevertStatus**](UserStatusUserStatusAPI.md#UserStatusUserStatusRevertStatus) | **Delete** /ocs/v2.php/apps/user_status/api/v1/user_status/revert/{messageId} | Revert the status to the previous status
[**UserStatusUserStatusSetCustomMessage**](UserStatusUserStatusAPI.md#UserStatusUserStatusSetCustomMessage) | **Put** /ocs/v2.php/apps/user_status/api/v1/user_status/message/custom | Set the message to a custom message for the current user
[**UserStatusUserStatusSetPredefinedMessage**](UserStatusUserStatusAPI.md#UserStatusUserStatusSetPredefinedMessage) | **Put** /ocs/v2.php/apps/user_status/api/v1/user_status/message/predefined | Set the message to a predefined message for the current user
[**UserStatusUserStatusSetStatus**](UserStatusUserStatusAPI.md#UserStatusUserStatusSetStatus) | **Put** /ocs/v2.php/apps/user_status/api/v1/user_status/status | Update the status type of the current user



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
    openapiclient "github.com/nextcloud/client-sdks"
)

func main() {
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserStatusUserStatusAPI.UserStatusUserStatusClearMessage(context.Background()).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserStatusUserStatusAPI.UserStatusUserStatusClearMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserStatusUserStatusClearMessage`: CoreWhatsNewDismiss200Response
    fmt.Fprintf(os.Stdout, "Response from `UserStatusUserStatusAPI.UserStatusUserStatusClearMessage`: %v\n", resp)
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
    openapiclient "github.com/nextcloud/client-sdks"
)

func main() {
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserStatusUserStatusAPI.UserStatusUserStatusGetStatus(context.Background()).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserStatusUserStatusAPI.UserStatusUserStatusGetStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserStatusUserStatusGetStatus`: UserStatusUserStatusGetStatus200Response
    fmt.Fprintf(os.Stdout, "Response from `UserStatusUserStatusAPI.UserStatusUserStatusGetStatus`: %v\n", resp)
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

> UserStatusUserStatusRevertStatus200Response UserStatusUserStatusRevertStatus(ctx, messageId).OCSAPIRequest(oCSAPIRequest).Execute()

Revert the status to the previous status

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
    messageId := "messageId_example" // string | ID of the message to delete
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserStatusUserStatusAPI.UserStatusUserStatusRevertStatus(context.Background(), messageId).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserStatusUserStatusAPI.UserStatusUserStatusRevertStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserStatusUserStatusRevertStatus`: UserStatusUserStatusRevertStatus200Response
    fmt.Fprintf(os.Stdout, "Response from `UserStatusUserStatusAPI.UserStatusUserStatusRevertStatus`: %v\n", resp)
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

[**UserStatusUserStatusRevertStatus200Response**](UserStatusUserStatusRevertStatus200Response.md)

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
    openapiclient "github.com/nextcloud/client-sdks"
)

func main() {
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")
    statusIcon := "statusIcon_example" // string | Icon of the status (optional)
    message := "message_example" // string | Message of the status (optional)
    clearAt := int64(789) // int64 | When the message should be cleared (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserStatusUserStatusAPI.UserStatusUserStatusSetCustomMessage(context.Background()).OCSAPIRequest(oCSAPIRequest).StatusIcon(statusIcon).Message(message).ClearAt(clearAt).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserStatusUserStatusAPI.UserStatusUserStatusSetCustomMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserStatusUserStatusSetCustomMessage`: UserStatusUserStatusGetStatus200Response
    fmt.Fprintf(os.Stdout, "Response from `UserStatusUserStatusAPI.UserStatusUserStatusSetCustomMessage`: %v\n", resp)
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
    openapiclient "github.com/nextcloud/client-sdks"
)

func main() {
    messageId := "messageId_example" // string | ID of the predefined message
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")
    clearAt := int64(789) // int64 | When the message should be cleared (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserStatusUserStatusAPI.UserStatusUserStatusSetPredefinedMessage(context.Background()).MessageId(messageId).OCSAPIRequest(oCSAPIRequest).ClearAt(clearAt).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserStatusUserStatusAPI.UserStatusUserStatusSetPredefinedMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserStatusUserStatusSetPredefinedMessage`: UserStatusUserStatusGetStatus200Response
    fmt.Fprintf(os.Stdout, "Response from `UserStatusUserStatusAPI.UserStatusUserStatusSetPredefinedMessage`: %v\n", resp)
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
    openapiclient "github.com/nextcloud/client-sdks"
)

func main() {
    statusType := "statusType_example" // string | The new status type
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserStatusUserStatusAPI.UserStatusUserStatusSetStatus(context.Background()).StatusType(statusType).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserStatusUserStatusAPI.UserStatusUserStatusSetStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserStatusUserStatusSetStatus`: UserStatusUserStatusGetStatus200Response
    fmt.Fprintf(os.Stdout, "Response from `UserStatusUserStatusAPI.UserStatusUserStatusSetStatus`: %v\n", resp)
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

