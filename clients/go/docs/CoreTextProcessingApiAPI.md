# \CoreTextProcessingApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CoreTextProcessingApiDeleteTask**](CoreTextProcessingApiAPI.md#CoreTextProcessingApiDeleteTask) | **Delete** /ocs/v2.php/textprocessing/task/{id} | This endpoint allows to delete a scheduled task for a user
[**CoreTextProcessingApiGetTask**](CoreTextProcessingApiAPI.md#CoreTextProcessingApiGetTask) | **Get** /ocs/v2.php/textprocessing/task/{id} | This endpoint allows checking the status and results of a task. Tasks are removed 1 week after receiving their last update.
[**CoreTextProcessingApiListTasksByApp**](CoreTextProcessingApiAPI.md#CoreTextProcessingApiListTasksByApp) | **Get** /ocs/v2.php/textprocessing/tasks/app/{appId} | This endpoint returns a list of tasks of a user that are related with a specific appId and optionally with an identifier
[**CoreTextProcessingApiSchedule**](CoreTextProcessingApiAPI.md#CoreTextProcessingApiSchedule) | **Post** /ocs/v2.php/textprocessing/schedule | This endpoint allows scheduling a language model task
[**CoreTextProcessingApiTaskTypes**](CoreTextProcessingApiAPI.md#CoreTextProcessingApiTaskTypes) | **Get** /ocs/v2.php/textprocessing/tasktypes | This endpoint returns all available LanguageModel task types



## CoreTextProcessingApiDeleteTask

> CoreTextProcessingApiSchedule200Response CoreTextProcessingApiDeleteTask(ctx, id).OCSAPIRequest(oCSAPIRequest).Execute()

This endpoint allows to delete a scheduled task for a user

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
    id := int64(789) // int64 | The id of the task
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoreTextProcessingApiAPI.CoreTextProcessingApiDeleteTask(context.Background(), id).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreTextProcessingApiAPI.CoreTextProcessingApiDeleteTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CoreTextProcessingApiDeleteTask`: CoreTextProcessingApiSchedule200Response
    fmt.Fprintf(os.Stdout, "Response from `CoreTextProcessingApiAPI.CoreTextProcessingApiDeleteTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The id of the task | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreTextProcessingApiDeleteTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]

### Return type

[**CoreTextProcessingApiSchedule200Response**](CoreTextProcessingApiSchedule200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreTextProcessingApiGetTask

> CoreTextProcessingApiSchedule200Response CoreTextProcessingApiGetTask(ctx, id).OCSAPIRequest(oCSAPIRequest).Execute()

This endpoint allows checking the status and results of a task. Tasks are removed 1 week after receiving their last update.

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
    id := int64(789) // int64 | The id of the task
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoreTextProcessingApiAPI.CoreTextProcessingApiGetTask(context.Background(), id).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreTextProcessingApiAPI.CoreTextProcessingApiGetTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CoreTextProcessingApiGetTask`: CoreTextProcessingApiSchedule200Response
    fmt.Fprintf(os.Stdout, "Response from `CoreTextProcessingApiAPI.CoreTextProcessingApiGetTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The id of the task | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreTextProcessingApiGetTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]

### Return type

[**CoreTextProcessingApiSchedule200Response**](CoreTextProcessingApiSchedule200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreTextProcessingApiListTasksByApp

> CoreTextProcessingApiListTasksByApp200Response CoreTextProcessingApiListTasksByApp(ctx, appId).OCSAPIRequest(oCSAPIRequest).Identifier(identifier).Execute()

This endpoint returns a list of tasks of a user that are related with a specific appId and optionally with an identifier

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
    appId := "appId_example" // string | ID of the app
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")
    identifier := "identifier_example" // string | An arbitrary identifier for the task (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoreTextProcessingApiAPI.CoreTextProcessingApiListTasksByApp(context.Background(), appId).OCSAPIRequest(oCSAPIRequest).Identifier(identifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreTextProcessingApiAPI.CoreTextProcessingApiListTasksByApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CoreTextProcessingApiListTasksByApp`: CoreTextProcessingApiListTasksByApp200Response
    fmt.Fprintf(os.Stdout, "Response from `CoreTextProcessingApiAPI.CoreTextProcessingApiListTasksByApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | ID of the app | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreTextProcessingApiListTasksByAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]
 **identifier** | **string** | An arbitrary identifier for the task | 

### Return type

[**CoreTextProcessingApiListTasksByApp200Response**](CoreTextProcessingApiListTasksByApp200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreTextProcessingApiSchedule

> CoreTextProcessingApiSchedule200Response CoreTextProcessingApiSchedule(ctx).Input(input).Type_(type_).AppId(appId).OCSAPIRequest(oCSAPIRequest).Identifier(identifier).Execute()

This endpoint allows scheduling a language model task

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
    input := "input_example" // string | Input text
    type_ := "type__example" // string | Type of the task
    appId := "appId_example" // string | ID of the app that will execute the task
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")
    identifier := "identifier_example" // string | An arbitrary identifier for the task (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoreTextProcessingApiAPI.CoreTextProcessingApiSchedule(context.Background()).Input(input).Type_(type_).AppId(appId).OCSAPIRequest(oCSAPIRequest).Identifier(identifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreTextProcessingApiAPI.CoreTextProcessingApiSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CoreTextProcessingApiSchedule`: CoreTextProcessingApiSchedule200Response
    fmt.Fprintf(os.Stdout, "Response from `CoreTextProcessingApiAPI.CoreTextProcessingApiSchedule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCoreTextProcessingApiScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **input** | **string** | Input text | 
 **type_** | **string** | Type of the task | 
 **appId** | **string** | ID of the app that will execute the task | 
 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]
 **identifier** | **string** | An arbitrary identifier for the task | [default to &quot;&quot;]

### Return type

[**CoreTextProcessingApiSchedule200Response**](CoreTextProcessingApiSchedule200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreTextProcessingApiTaskTypes

> CoreTextProcessingApiTaskTypes200Response CoreTextProcessingApiTaskTypes(ctx).OCSAPIRequest(oCSAPIRequest).Execute()

This endpoint returns all available LanguageModel task types

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
    resp, r, err := apiClient.CoreTextProcessingApiAPI.CoreTextProcessingApiTaskTypes(context.Background()).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreTextProcessingApiAPI.CoreTextProcessingApiTaskTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CoreTextProcessingApiTaskTypes`: CoreTextProcessingApiTaskTypes200Response
    fmt.Fprintf(os.Stdout, "Response from `CoreTextProcessingApiAPI.CoreTextProcessingApiTaskTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCoreTextProcessingApiTaskTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]

### Return type

[**CoreTextProcessingApiTaskTypes200Response**](CoreTextProcessingApiTaskTypes200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

