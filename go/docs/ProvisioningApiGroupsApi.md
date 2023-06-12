# \ProvisioningApiGroupsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProvisioningApiGroupsAddGroup**](ProvisioningApiGroupsApi.md#ProvisioningApiGroupsAddGroup) | **Post** /ocs/v2.php/cloud/groups | Create a new group
[**ProvisioningApiGroupsDeleteGroup**](ProvisioningApiGroupsApi.md#ProvisioningApiGroupsDeleteGroup) | **Delete** /ocs/v2.php/cloud/groups/{groupId} | Delete a group
[**ProvisioningApiGroupsGetGroup**](ProvisioningApiGroupsApi.md#ProvisioningApiGroupsGetGroup) | **Get** /ocs/v2.php/cloud/groups/{groupId} | Get a list of users in the specified group
[**ProvisioningApiGroupsGetGroupUsers**](ProvisioningApiGroupsApi.md#ProvisioningApiGroupsGetGroupUsers) | **Get** /ocs/v2.php/cloud/groups/{groupId}/users | Get a list of users in the specified group
[**ProvisioningApiGroupsGetGroupUsersDetails**](ProvisioningApiGroupsApi.md#ProvisioningApiGroupsGetGroupUsersDetails) | **Get** /ocs/v2.php/cloud/groups/{groupId}/users/details | Get a list of users details in the specified group
[**ProvisioningApiGroupsGetGroups**](ProvisioningApiGroupsApi.md#ProvisioningApiGroupsGetGroups) | **Get** /ocs/v2.php/cloud/groups | Get a list of groups
[**ProvisioningApiGroupsGetGroupsDetails**](ProvisioningApiGroupsApi.md#ProvisioningApiGroupsGetGroupsDetails) | **Get** /ocs/v2.php/cloud/groups/details | Get a list of groups details
[**ProvisioningApiGroupsGetSubAdminsOfGroup**](ProvisioningApiGroupsApi.md#ProvisioningApiGroupsGetSubAdminsOfGroup) | **Get** /ocs/v2.php/cloud/groups/{groupId}/subadmins | Get the list of user IDs that are a subadmin of the group
[**ProvisioningApiGroupsUpdateGroup**](ProvisioningApiGroupsApi.md#ProvisioningApiGroupsUpdateGroup) | **Put** /ocs/v2.php/cloud/groups/{groupId} | Update a group



## ProvisioningApiGroupsAddGroup

> CoreWhatsNewDismiss200Response ProvisioningApiGroupsAddGroup(ctx).Groupid(groupid).OCSAPIRequest(oCSAPIRequest).Displayname(displayname).Execute()

Create a new group



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
    groupid := "groupid_example" // string | ID of the group
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")
    displayname := "displayname_example" // string | Display name of the group (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvisioningApiGroupsApi.ProvisioningApiGroupsAddGroup(context.Background()).Groupid(groupid).OCSAPIRequest(oCSAPIRequest).Displayname(displayname).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningApiGroupsApi.ProvisioningApiGroupsAddGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisioningApiGroupsAddGroup`: CoreWhatsNewDismiss200Response
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningApiGroupsApi.ProvisioningApiGroupsAddGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProvisioningApiGroupsAddGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupid** | **string** | ID of the group | 
 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]
 **displayname** | **string** | Display name of the group | [default to &quot;&quot;]

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


## ProvisioningApiGroupsDeleteGroup

> CoreWhatsNewDismiss200Response ProvisioningApiGroupsDeleteGroup(ctx, groupId).OCSAPIRequest(oCSAPIRequest).Execute()

Delete a group



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
    groupId := "groupId_example" // string | ID of the group
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvisioningApiGroupsApi.ProvisioningApiGroupsDeleteGroup(context.Background(), groupId).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningApiGroupsApi.ProvisioningApiGroupsDeleteGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisioningApiGroupsDeleteGroup`: CoreWhatsNewDismiss200Response
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningApiGroupsApi.ProvisioningApiGroupsDeleteGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | ID of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvisioningApiGroupsDeleteGroupRequest struct via the builder pattern


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


## ProvisioningApiGroupsGetGroup

> ProvisioningApiGroupsGetGroupUsers200Response ProvisioningApiGroupsGetGroup(ctx, groupId).OCSAPIRequest(oCSAPIRequest).Execute()

Get a list of users in the specified group

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
    groupId := "groupId_example" // string | ID of the group
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvisioningApiGroupsApi.ProvisioningApiGroupsGetGroup(context.Background(), groupId).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningApiGroupsApi.ProvisioningApiGroupsGetGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisioningApiGroupsGetGroup`: ProvisioningApiGroupsGetGroupUsers200Response
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningApiGroupsApi.ProvisioningApiGroupsGetGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | ID of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvisioningApiGroupsGetGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]

### Return type

[**ProvisioningApiGroupsGetGroupUsers200Response**](ProvisioningApiGroupsGetGroupUsers200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvisioningApiGroupsGetGroupUsers

> ProvisioningApiGroupsGetGroupUsers200Response ProvisioningApiGroupsGetGroupUsers(ctx, groupId).OCSAPIRequest(oCSAPIRequest).Execute()

Get a list of users in the specified group

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
    groupId := "groupId_example" // string | ID of the group
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvisioningApiGroupsApi.ProvisioningApiGroupsGetGroupUsers(context.Background(), groupId).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningApiGroupsApi.ProvisioningApiGroupsGetGroupUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisioningApiGroupsGetGroupUsers`: ProvisioningApiGroupsGetGroupUsers200Response
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningApiGroupsApi.ProvisioningApiGroupsGetGroupUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | ID of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvisioningApiGroupsGetGroupUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]

### Return type

[**ProvisioningApiGroupsGetGroupUsers200Response**](ProvisioningApiGroupsGetGroupUsers200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvisioningApiGroupsGetGroupUsersDetails

> ProvisioningApiGroupsGetGroupUsersDetails200Response ProvisioningApiGroupsGetGroupUsersDetails(ctx, groupId).OCSAPIRequest(oCSAPIRequest).Search(search).Limit(limit).Offset(offset).Execute()

Get a list of users details in the specified group

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
    groupId := "groupId_example" // string | ID of the group
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")
    search := "search_example" // string | Text to search for (optional) (default to "")
    limit := int64(789) // int64 | Limit the amount of groups returned (optional)
    offset := int64(789) // int64 | Offset for searching for groups (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvisioningApiGroupsApi.ProvisioningApiGroupsGetGroupUsersDetails(context.Background(), groupId).OCSAPIRequest(oCSAPIRequest).Search(search).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningApiGroupsApi.ProvisioningApiGroupsGetGroupUsersDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisioningApiGroupsGetGroupUsersDetails`: ProvisioningApiGroupsGetGroupUsersDetails200Response
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningApiGroupsApi.ProvisioningApiGroupsGetGroupUsersDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | ID of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvisioningApiGroupsGetGroupUsersDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]
 **search** | **string** | Text to search for | [default to &quot;&quot;]
 **limit** | **int64** | Limit the amount of groups returned | 
 **offset** | **int64** | Offset for searching for groups | [default to 0]

### Return type

[**ProvisioningApiGroupsGetGroupUsersDetails200Response**](ProvisioningApiGroupsGetGroupUsersDetails200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvisioningApiGroupsGetGroups

> ProvisioningApiGroupsGetGroups200Response ProvisioningApiGroupsGetGroups(ctx).OCSAPIRequest(oCSAPIRequest).Search(search).Limit(limit).Offset(offset).Execute()

Get a list of groups

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
    search := "search_example" // string | Text to search for (optional) (default to "")
    limit := int64(789) // int64 | Limit the amount of groups returned (optional)
    offset := int64(789) // int64 | Offset for searching for groups (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvisioningApiGroupsApi.ProvisioningApiGroupsGetGroups(context.Background()).OCSAPIRequest(oCSAPIRequest).Search(search).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningApiGroupsApi.ProvisioningApiGroupsGetGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisioningApiGroupsGetGroups`: ProvisioningApiGroupsGetGroups200Response
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningApiGroupsApi.ProvisioningApiGroupsGetGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProvisioningApiGroupsGetGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]
 **search** | **string** | Text to search for | [default to &quot;&quot;]
 **limit** | **int64** | Limit the amount of groups returned | 
 **offset** | **int64** | Offset for searching for groups | [default to 0]

### Return type

[**ProvisioningApiGroupsGetGroups200Response**](ProvisioningApiGroupsGetGroups200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvisioningApiGroupsGetGroupsDetails

> ProvisioningApiGroupsGetGroupsDetails200Response ProvisioningApiGroupsGetGroupsDetails(ctx).OCSAPIRequest(oCSAPIRequest).Search(search).Limit(limit).Offset(offset).Execute()

Get a list of groups details

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
    search := "search_example" // string | Text to search for (optional) (default to "")
    limit := int64(789) // int64 | Limit the amount of groups returned (optional)
    offset := int64(789) // int64 | Offset for searching for groups (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvisioningApiGroupsApi.ProvisioningApiGroupsGetGroupsDetails(context.Background()).OCSAPIRequest(oCSAPIRequest).Search(search).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningApiGroupsApi.ProvisioningApiGroupsGetGroupsDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisioningApiGroupsGetGroupsDetails`: ProvisioningApiGroupsGetGroupsDetails200Response
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningApiGroupsApi.ProvisioningApiGroupsGetGroupsDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProvisioningApiGroupsGetGroupsDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]
 **search** | **string** | Text to search for | [default to &quot;&quot;]
 **limit** | **int64** | Limit the amount of groups returned | 
 **offset** | **int64** | Offset for searching for groups | [default to 0]

### Return type

[**ProvisioningApiGroupsGetGroupsDetails200Response**](ProvisioningApiGroupsGetGroupsDetails200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvisioningApiGroupsGetSubAdminsOfGroup

> ProvisioningApiGroupsGetSubAdminsOfGroup200Response ProvisioningApiGroupsGetSubAdminsOfGroup(ctx, groupId).OCSAPIRequest(oCSAPIRequest).Execute()

Get the list of user IDs that are a subadmin of the group



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
    groupId := "groupId_example" // string | ID of the group
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvisioningApiGroupsApi.ProvisioningApiGroupsGetSubAdminsOfGroup(context.Background(), groupId).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningApiGroupsApi.ProvisioningApiGroupsGetSubAdminsOfGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisioningApiGroupsGetSubAdminsOfGroup`: ProvisioningApiGroupsGetSubAdminsOfGroup200Response
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningApiGroupsApi.ProvisioningApiGroupsGetSubAdminsOfGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | ID of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvisioningApiGroupsGetSubAdminsOfGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]

### Return type

[**ProvisioningApiGroupsGetSubAdminsOfGroup200Response**](ProvisioningApiGroupsGetSubAdminsOfGroup200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvisioningApiGroupsUpdateGroup

> CoreWhatsNewDismiss200Response ProvisioningApiGroupsUpdateGroup(ctx, groupId).Key(key).Value(value).OCSAPIRequest(oCSAPIRequest).Execute()

Update a group



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
    key := "key_example" // string | Key to update, only 'displayname'
    value := "value_example" // string | New value for the key
    groupId := "groupId_example" // string | ID of the group
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvisioningApiGroupsApi.ProvisioningApiGroupsUpdateGroup(context.Background(), groupId).Key(key).Value(value).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningApiGroupsApi.ProvisioningApiGroupsUpdateGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisioningApiGroupsUpdateGroup`: CoreWhatsNewDismiss200Response
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningApiGroupsApi.ProvisioningApiGroupsUpdateGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | ID of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvisioningApiGroupsUpdateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **key** | **string** | Key to update, only &#39;displayname&#39; | 
 **value** | **string** | New value for the key | 

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

