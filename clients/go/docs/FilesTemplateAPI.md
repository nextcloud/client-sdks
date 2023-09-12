# \FilesTemplateAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FilesTemplateCreate**](FilesTemplateAPI.md#FilesTemplateCreate) | **Post** /ocs/v2.php/apps/files/api/v1/templates/create | Create a template
[**FilesTemplateList**](FilesTemplateAPI.md#FilesTemplateList) | **Get** /ocs/v2.php/apps/files/api/v1/templates | List the available templates
[**FilesTemplatePath**](FilesTemplateAPI.md#FilesTemplatePath) | **Post** /ocs/v2.php/apps/files/api/v1/templates/path | Initialize the template directory



## FilesTemplateCreate

> FilesTemplateCreate200Response FilesTemplateCreate(ctx).FilePath(filePath).OCSAPIRequest(oCSAPIRequest).TemplatePath(templatePath).TemplateType(templateType).Execute()

Create a template

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
    filePath := "filePath_example" // string | Path of the file
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")
    templatePath := "templatePath_example" // string | Name of the template (optional) (default to "")
    templateType := "templateType_example" // string | Type of the template (optional) (default to "user")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FilesTemplateAPI.FilesTemplateCreate(context.Background()).FilePath(filePath).OCSAPIRequest(oCSAPIRequest).TemplatePath(templatePath).TemplateType(templateType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesTemplateAPI.FilesTemplateCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FilesTemplateCreate`: FilesTemplateCreate200Response
    fmt.Fprintf(os.Stdout, "Response from `FilesTemplateAPI.FilesTemplateCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFilesTemplateCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filePath** | **string** | Path of the file | 
 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]
 **templatePath** | **string** | Name of the template | [default to &quot;&quot;]
 **templateType** | **string** | Type of the template | [default to &quot;user&quot;]

### Return type

[**FilesTemplateCreate200Response**](FilesTemplateCreate200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FilesTemplateList

> FilesTemplateList200Response FilesTemplateList(ctx).OCSAPIRequest(oCSAPIRequest).Execute()

List the available templates

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
    resp, r, err := apiClient.FilesTemplateAPI.FilesTemplateList(context.Background()).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesTemplateAPI.FilesTemplateList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FilesTemplateList`: FilesTemplateList200Response
    fmt.Fprintf(os.Stdout, "Response from `FilesTemplateAPI.FilesTemplateList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFilesTemplateListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]

### Return type

[**FilesTemplateList200Response**](FilesTemplateList200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FilesTemplatePath

> FilesTemplatePath200Response FilesTemplatePath(ctx).OCSAPIRequest(oCSAPIRequest).TemplatePath(templatePath).CopySystemTemplates(copySystemTemplates).Execute()

Initialize the template directory

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
    templatePath := "templatePath_example" // string | Path of the template directory (optional) (default to "")
    copySystemTemplates := int32(56) // int32 | Whether to copy the system templates to the template directory (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FilesTemplateAPI.FilesTemplatePath(context.Background()).OCSAPIRequest(oCSAPIRequest).TemplatePath(templatePath).CopySystemTemplates(copySystemTemplates).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesTemplateAPI.FilesTemplatePath``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FilesTemplatePath`: FilesTemplatePath200Response
    fmt.Fprintf(os.Stdout, "Response from `FilesTemplateAPI.FilesTemplatePath`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFilesTemplatePathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]
 **templatePath** | **string** | Path of the template directory | [default to &quot;&quot;]
 **copySystemTemplates** | **int32** | Whether to copy the system templates to the template directory | [default to 0]

### Return type

[**FilesTemplatePath200Response**](FilesTemplatePath200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

