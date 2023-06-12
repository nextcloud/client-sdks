# \CoreCollaborationResourcesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CoreCollaborationResourcesAddResource**](CoreCollaborationResourcesApi.md#CoreCollaborationResourcesAddResource) | **Post** /ocs/v2.php/collaboration/resources/collections/{collectionId} | Add a resource to a collection
[**CoreCollaborationResourcesCreateCollectionOnResource**](CoreCollaborationResourcesApi.md#CoreCollaborationResourcesCreateCollectionOnResource) | **Post** /ocs/v2.php/collaboration/resources/{baseResourceType}/{baseResourceId} | Create a collection for a resource
[**CoreCollaborationResourcesGetCollectionsByResource**](CoreCollaborationResourcesApi.md#CoreCollaborationResourcesGetCollectionsByResource) | **Get** /ocs/v2.php/collaboration/resources/{resourceType}/{resourceId} | Get collections by resource
[**CoreCollaborationResourcesListCollection**](CoreCollaborationResourcesApi.md#CoreCollaborationResourcesListCollection) | **Get** /ocs/v2.php/collaboration/resources/collections/{collectionId} | Get a collection
[**CoreCollaborationResourcesRemoveResource**](CoreCollaborationResourcesApi.md#CoreCollaborationResourcesRemoveResource) | **Delete** /ocs/v2.php/collaboration/resources/collections/{collectionId} | Remove a resource from a collection
[**CoreCollaborationResourcesRenameCollection**](CoreCollaborationResourcesApi.md#CoreCollaborationResourcesRenameCollection) | **Put** /ocs/v2.php/collaboration/resources/collections/{collectionId} | Rename a collection
[**CoreCollaborationResourcesSearchCollections**](CoreCollaborationResourcesApi.md#CoreCollaborationResourcesSearchCollections) | **Get** /ocs/v2.php/collaboration/resources/collections/search/{filter} | Search for collections



## CoreCollaborationResourcesAddResource

> CoreCollaborationResourcesListCollection200Response CoreCollaborationResourcesAddResource(ctx, collectionId).ResourceType(resourceType).ResourceId(resourceId).OCSAPIRequest(oCSAPIRequest).Execute()

Add a resource to a collection

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
    resourceType := "resourceType_example" // string | Name of the resource
    resourceId := "resourceId_example" // string | ID of the resource
    collectionId := int64(789) // int64 | ID of the collection
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoreCollaborationResourcesApi.CoreCollaborationResourcesAddResource(context.Background(), collectionId).ResourceType(resourceType).ResourceId(resourceId).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreCollaborationResourcesApi.CoreCollaborationResourcesAddResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CoreCollaborationResourcesAddResource`: CoreCollaborationResourcesListCollection200Response
    fmt.Fprintf(os.Stdout, "Response from `CoreCollaborationResourcesApi.CoreCollaborationResourcesAddResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | ID of the collection | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreCollaborationResourcesAddResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resourceType** | **string** | Name of the resource | 
 **resourceId** | **string** | ID of the resource | 

 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]

### Return type

[**CoreCollaborationResourcesListCollection200Response**](CoreCollaborationResourcesListCollection200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreCollaborationResourcesCreateCollectionOnResource

> CoreCollaborationResourcesListCollection200Response CoreCollaborationResourcesCreateCollectionOnResource(ctx, baseResourceType, baseResourceId).Name(name).OCSAPIRequest(oCSAPIRequest).Execute()

Create a collection for a resource

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
    name := "name_example" // string | Name of the collection
    baseResourceType := "baseResourceType_example" // string | Type of the base resource
    baseResourceId := "baseResourceId_example" // string | ID of the base resource
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoreCollaborationResourcesApi.CoreCollaborationResourcesCreateCollectionOnResource(context.Background(), baseResourceType, baseResourceId).Name(name).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreCollaborationResourcesApi.CoreCollaborationResourcesCreateCollectionOnResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CoreCollaborationResourcesCreateCollectionOnResource`: CoreCollaborationResourcesListCollection200Response
    fmt.Fprintf(os.Stdout, "Response from `CoreCollaborationResourcesApi.CoreCollaborationResourcesCreateCollectionOnResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**baseResourceType** | **string** | Type of the base resource | 
**baseResourceId** | **string** | ID of the base resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreCollaborationResourcesCreateCollectionOnResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Name of the collection | 


 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]

### Return type

[**CoreCollaborationResourcesListCollection200Response**](CoreCollaborationResourcesListCollection200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreCollaborationResourcesGetCollectionsByResource

> CoreCollaborationResourcesSearchCollections200Response CoreCollaborationResourcesGetCollectionsByResource(ctx, resourceType, resourceId).OCSAPIRequest(oCSAPIRequest).Execute()

Get collections by resource

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
    resourceType := "resourceType_example" // string | Type of the resource
    resourceId := "resourceId_example" // string | ID of the resource
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoreCollaborationResourcesApi.CoreCollaborationResourcesGetCollectionsByResource(context.Background(), resourceType, resourceId).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreCollaborationResourcesApi.CoreCollaborationResourcesGetCollectionsByResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CoreCollaborationResourcesGetCollectionsByResource`: CoreCollaborationResourcesSearchCollections200Response
    fmt.Fprintf(os.Stdout, "Response from `CoreCollaborationResourcesApi.CoreCollaborationResourcesGetCollectionsByResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceType** | **string** | Type of the resource | 
**resourceId** | **string** | ID of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreCollaborationResourcesGetCollectionsByResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]

### Return type

[**CoreCollaborationResourcesSearchCollections200Response**](CoreCollaborationResourcesSearchCollections200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreCollaborationResourcesListCollection

> CoreCollaborationResourcesListCollection200Response CoreCollaborationResourcesListCollection(ctx, collectionId).OCSAPIRequest(oCSAPIRequest).Execute()

Get a collection

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
    collectionId := int64(789) // int64 | ID of the collection
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoreCollaborationResourcesApi.CoreCollaborationResourcesListCollection(context.Background(), collectionId).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreCollaborationResourcesApi.CoreCollaborationResourcesListCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CoreCollaborationResourcesListCollection`: CoreCollaborationResourcesListCollection200Response
    fmt.Fprintf(os.Stdout, "Response from `CoreCollaborationResourcesApi.CoreCollaborationResourcesListCollection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | ID of the collection | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreCollaborationResourcesListCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]

### Return type

[**CoreCollaborationResourcesListCollection200Response**](CoreCollaborationResourcesListCollection200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreCollaborationResourcesRemoveResource

> CoreCollaborationResourcesListCollection200Response CoreCollaborationResourcesRemoveResource(ctx, collectionId).ResourceType(resourceType).ResourceId(resourceId).OCSAPIRequest(oCSAPIRequest).Execute()

Remove a resource from a collection

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
    resourceType := "resourceType_example" // string | Name of the resource
    resourceId := "resourceId_example" // string | ID of the resource
    collectionId := int64(789) // int64 | ID of the collection
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoreCollaborationResourcesApi.CoreCollaborationResourcesRemoveResource(context.Background(), collectionId).ResourceType(resourceType).ResourceId(resourceId).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreCollaborationResourcesApi.CoreCollaborationResourcesRemoveResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CoreCollaborationResourcesRemoveResource`: CoreCollaborationResourcesListCollection200Response
    fmt.Fprintf(os.Stdout, "Response from `CoreCollaborationResourcesApi.CoreCollaborationResourcesRemoveResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | ID of the collection | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreCollaborationResourcesRemoveResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resourceType** | **string** | Name of the resource | 
 **resourceId** | **string** | ID of the resource | 

 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]

### Return type

[**CoreCollaborationResourcesListCollection200Response**](CoreCollaborationResourcesListCollection200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreCollaborationResourcesRenameCollection

> CoreCollaborationResourcesListCollection200Response CoreCollaborationResourcesRenameCollection(ctx, collectionId).CollectionName(collectionName).OCSAPIRequest(oCSAPIRequest).Execute()

Rename a collection

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
    collectionName := "collectionName_example" // string | New name
    collectionId := int64(789) // int64 | ID of the collection
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoreCollaborationResourcesApi.CoreCollaborationResourcesRenameCollection(context.Background(), collectionId).CollectionName(collectionName).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreCollaborationResourcesApi.CoreCollaborationResourcesRenameCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CoreCollaborationResourcesRenameCollection`: CoreCollaborationResourcesListCollection200Response
    fmt.Fprintf(os.Stdout, "Response from `CoreCollaborationResourcesApi.CoreCollaborationResourcesRenameCollection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | ID of the collection | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreCollaborationResourcesRenameCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **collectionName** | **string** | New name | 

 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]

### Return type

[**CoreCollaborationResourcesListCollection200Response**](CoreCollaborationResourcesListCollection200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreCollaborationResourcesSearchCollections

> CoreCollaborationResourcesSearchCollections200Response CoreCollaborationResourcesSearchCollections(ctx, filter).OCSAPIRequest(oCSAPIRequest).Execute()

Search for collections

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
    filter := "filter_example" // string | Filter collections
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoreCollaborationResourcesApi.CoreCollaborationResourcesSearchCollections(context.Background(), filter).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreCollaborationResourcesApi.CoreCollaborationResourcesSearchCollections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CoreCollaborationResourcesSearchCollections`: CoreCollaborationResourcesSearchCollections200Response
    fmt.Fprintf(os.Stdout, "Response from `CoreCollaborationResourcesApi.CoreCollaborationResourcesSearchCollections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filter** | **string** | Filter collections | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreCollaborationResourcesSearchCollectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]

### Return type

[**CoreCollaborationResourcesSearchCollections200Response**](CoreCollaborationResourcesSearchCollections200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

