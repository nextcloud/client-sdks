# \FilesDirectEditingAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FilesDirectEditingCreate**](FilesDirectEditingAPI.md#FilesDirectEditingCreate) | **Post** /ocs/v2.php/apps/files/api/v1/directEditing/create | Create a file for direct editing
[**FilesDirectEditingInfo**](FilesDirectEditingAPI.md#FilesDirectEditingInfo) | **Get** /ocs/v2.php/apps/files/api/v1/directEditing | Get the direct editing capabilities
[**FilesDirectEditingOpen**](FilesDirectEditingAPI.md#FilesDirectEditingOpen) | **Post** /ocs/v2.php/apps/files/api/v1/directEditing/open | Open a file for direct editing
[**FilesDirectEditingTemplates**](FilesDirectEditingAPI.md#FilesDirectEditingTemplates) | **Get** /ocs/v2.php/apps/files/api/v1/directEditing/templates/{editorId}/{creatorId} | Get the templates for direct editing



## FilesDirectEditingCreate

> DavDirectGetUrl200Response FilesDirectEditingCreate(ctx).Path(path).EditorId(editorId).CreatorId(creatorId).OCSAPIRequest(oCSAPIRequest).TemplateId(templateId).Execute()

Create a file for direct editing

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
    editorId := "editorId_example" // string | ID of the editor
    creatorId := "creatorId_example" // string | ID of the creator
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")
    templateId := "templateId_example" // string | ID of the template (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FilesDirectEditingAPI.FilesDirectEditingCreate(context.Background()).Path(path).EditorId(editorId).CreatorId(creatorId).OCSAPIRequest(oCSAPIRequest).TemplateId(templateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesDirectEditingAPI.FilesDirectEditingCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FilesDirectEditingCreate`: DavDirectGetUrl200Response
    fmt.Fprintf(os.Stdout, "Response from `FilesDirectEditingAPI.FilesDirectEditingCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFilesDirectEditingCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | Path of the file | 
 **editorId** | **string** | ID of the editor | 
 **creatorId** | **string** | ID of the creator | 
 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]
 **templateId** | **string** | ID of the template | 

### Return type

[**DavDirectGetUrl200Response**](DavDirectGetUrl200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FilesDirectEditingInfo

> FilesDirectEditingInfo200Response FilesDirectEditingInfo(ctx).OCSAPIRequest(oCSAPIRequest).Execute()

Get the direct editing capabilities

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
    resp, r, err := apiClient.FilesDirectEditingAPI.FilesDirectEditingInfo(context.Background()).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesDirectEditingAPI.FilesDirectEditingInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FilesDirectEditingInfo`: FilesDirectEditingInfo200Response
    fmt.Fprintf(os.Stdout, "Response from `FilesDirectEditingAPI.FilesDirectEditingInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFilesDirectEditingInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]

### Return type

[**FilesDirectEditingInfo200Response**](FilesDirectEditingInfo200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FilesDirectEditingOpen

> DavDirectGetUrl200Response FilesDirectEditingOpen(ctx).Path(path).OCSAPIRequest(oCSAPIRequest).EditorId(editorId).FileId(fileId).Execute()

Open a file for direct editing

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
    editorId := "editorId_example" // string | ID of the editor (optional)
    fileId := int64(789) // int64 | ID of the file (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FilesDirectEditingAPI.FilesDirectEditingOpen(context.Background()).Path(path).OCSAPIRequest(oCSAPIRequest).EditorId(editorId).FileId(fileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesDirectEditingAPI.FilesDirectEditingOpen``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FilesDirectEditingOpen`: DavDirectGetUrl200Response
    fmt.Fprintf(os.Stdout, "Response from `FilesDirectEditingAPI.FilesDirectEditingOpen`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFilesDirectEditingOpenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | Path of the file | 
 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]
 **editorId** | **string** | ID of the editor | 
 **fileId** | **int64** | ID of the file | 

### Return type

[**DavDirectGetUrl200Response**](DavDirectGetUrl200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FilesDirectEditingTemplates

> FilesDirectEditingTemplates200Response FilesDirectEditingTemplates(ctx, editorId, creatorId).OCSAPIRequest(oCSAPIRequest).Execute()

Get the templates for direct editing

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
    editorId := "editorId_example" // string | ID of the editor
    creatorId := "creatorId_example" // string | ID of the creator
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FilesDirectEditingAPI.FilesDirectEditingTemplates(context.Background(), editorId, creatorId).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesDirectEditingAPI.FilesDirectEditingTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FilesDirectEditingTemplates`: FilesDirectEditingTemplates200Response
    fmt.Fprintf(os.Stdout, "Response from `FilesDirectEditingAPI.FilesDirectEditingTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**editorId** | **string** | ID of the editor | 
**creatorId** | **string** | ID of the creator | 

### Other Parameters

Other parameters are passed through a pointer to a apiFilesDirectEditingTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]

### Return type

[**FilesDirectEditingTemplates200Response**](FilesDirectEditingTemplates200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

