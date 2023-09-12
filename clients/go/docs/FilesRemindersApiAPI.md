# \FilesRemindersApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FilesRemindersApiGet**](FilesRemindersApiAPI.md#FilesRemindersApiGet) | **Get** /ocs/v2.php/apps/files_reminders/api/v{version}/{fileId} | Get a reminder
[**FilesRemindersApiRemove**](FilesRemindersApiAPI.md#FilesRemindersApiRemove) | **Delete** /ocs/v2.php/apps/files_reminders/api/v{version}/{fileId} | Remove a reminder
[**FilesRemindersApiSet**](FilesRemindersApiAPI.md#FilesRemindersApiSet) | **Put** /ocs/v2.php/apps/files_reminders/api/v{version}/{fileId} | Set a reminder



## FilesRemindersApiGet

> FilesRemindersApiGet200Response FilesRemindersApiGet(ctx, version, fileId).OCSAPIRequest(oCSAPIRequest).Execute()

Get a reminder

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
    version := "version_example" // string | 
    fileId := int64(789) // int64 | ID of the file
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FilesRemindersApiAPI.FilesRemindersApiGet(context.Background(), version, fileId).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesRemindersApiAPI.FilesRemindersApiGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FilesRemindersApiGet`: FilesRemindersApiGet200Response
    fmt.Fprintf(os.Stdout, "Response from `FilesRemindersApiAPI.FilesRemindersApiGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**version** | **string** |  | 
**fileId** | **int64** | ID of the file | 

### Other Parameters

Other parameters are passed through a pointer to a apiFilesRemindersApiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]

### Return type

[**FilesRemindersApiGet200Response**](FilesRemindersApiGet200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FilesRemindersApiRemove

> CoreWhatsNewDismiss200Response FilesRemindersApiRemove(ctx, version, fileId).OCSAPIRequest(oCSAPIRequest).Execute()

Remove a reminder

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
    version := "version_example" // string | 
    fileId := int64(789) // int64 | ID of the file
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FilesRemindersApiAPI.FilesRemindersApiRemove(context.Background(), version, fileId).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesRemindersApiAPI.FilesRemindersApiRemove``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FilesRemindersApiRemove`: CoreWhatsNewDismiss200Response
    fmt.Fprintf(os.Stdout, "Response from `FilesRemindersApiAPI.FilesRemindersApiRemove`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**version** | **string** |  | 
**fileId** | **int64** | ID of the file | 

### Other Parameters

Other parameters are passed through a pointer to a apiFilesRemindersApiRemoveRequest struct via the builder pattern


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


## FilesRemindersApiSet

> CoreWhatsNewDismiss200Response FilesRemindersApiSet(ctx, version, fileId).DueDate(dueDate).OCSAPIRequest(oCSAPIRequest).Execute()

Set a reminder

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
    dueDate := "dueDate_example" // string | ISO 8601 formatted date time string
    version := "version_example" // string | 
    fileId := int64(789) // int64 | ID of the file
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FilesRemindersApiAPI.FilesRemindersApiSet(context.Background(), version, fileId).DueDate(dueDate).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesRemindersApiAPI.FilesRemindersApiSet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FilesRemindersApiSet`: CoreWhatsNewDismiss200Response
    fmt.Fprintf(os.Stdout, "Response from `FilesRemindersApiAPI.FilesRemindersApiSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**version** | **string** |  | 
**fileId** | **int64** | ID of the file | 

### Other Parameters

Other parameters are passed through a pointer to a apiFilesRemindersApiSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dueDate** | **string** | ISO 8601 formatted date time string | 


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

