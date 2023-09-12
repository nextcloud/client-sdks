# \FilesOpenLocalEditorAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FilesOpenLocalEditorCreate**](FilesOpenLocalEditorAPI.md#FilesOpenLocalEditorCreate) | **Post** /ocs/v2.php/apps/files/api/v1/openlocaleditor | Create a local editor
[**FilesOpenLocalEditorValidate**](FilesOpenLocalEditorAPI.md#FilesOpenLocalEditorValidate) | **Post** /ocs/v2.php/apps/files/api/v1/openlocaleditor/{token} | Validate a local editor



## FilesOpenLocalEditorCreate

> FilesOpenLocalEditorCreate200Response FilesOpenLocalEditorCreate(ctx).Path(path).OCSAPIRequest(oCSAPIRequest).Execute()

Create a local editor

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
    path := "path_example" // string | Path of the file
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FilesOpenLocalEditorAPI.FilesOpenLocalEditorCreate(context.Background()).Path(path).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesOpenLocalEditorAPI.FilesOpenLocalEditorCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FilesOpenLocalEditorCreate`: FilesOpenLocalEditorCreate200Response
    fmt.Fprintf(os.Stdout, "Response from `FilesOpenLocalEditorAPI.FilesOpenLocalEditorCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFilesOpenLocalEditorCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | Path of the file | 
 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]

### Return type

[**FilesOpenLocalEditorCreate200Response**](FilesOpenLocalEditorCreate200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FilesOpenLocalEditorValidate

> FilesOpenLocalEditorValidate200Response FilesOpenLocalEditorValidate(ctx, token).Path(path).OCSAPIRequest(oCSAPIRequest).Execute()

Validate a local editor

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
    path := "path_example" // string | Path of the file
    token := "token_example" // string | Token of the local editor
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FilesOpenLocalEditorAPI.FilesOpenLocalEditorValidate(context.Background(), token).Path(path).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesOpenLocalEditorAPI.FilesOpenLocalEditorValidate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FilesOpenLocalEditorValidate`: FilesOpenLocalEditorValidate200Response
    fmt.Fprintf(os.Stdout, "Response from `FilesOpenLocalEditorAPI.FilesOpenLocalEditorValidate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | Token of the local editor | 

### Other Parameters

Other parameters are passed through a pointer to a apiFilesOpenLocalEditorValidateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | Path of the file | 

 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]

### Return type

[**FilesOpenLocalEditorValidate200Response**](FilesOpenLocalEditorValidate200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

