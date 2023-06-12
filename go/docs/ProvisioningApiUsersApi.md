# \ProvisioningApiUsersApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProvisioningApiUsersAddSubAdmin**](ProvisioningApiUsersApi.md#ProvisioningApiUsersAddSubAdmin) | **Post** /ocs/v2.php/cloud/users/{userId}/subadmins | Make a user a subadmin of a group
[**ProvisioningApiUsersAddToGroup**](ProvisioningApiUsersApi.md#ProvisioningApiUsersAddToGroup) | **Post** /ocs/v2.php/cloud/users/{userId}/groups | Add a user to a group
[**ProvisioningApiUsersAddUser**](ProvisioningApiUsersApi.md#ProvisioningApiUsersAddUser) | **Post** /ocs/v2.php/cloud/users | Create a new user
[**ProvisioningApiUsersDeleteUser**](ProvisioningApiUsersApi.md#ProvisioningApiUsersDeleteUser) | **Delete** /ocs/v2.php/cloud/users/{userId} | Delete a user
[**ProvisioningApiUsersDisableUser**](ProvisioningApiUsersApi.md#ProvisioningApiUsersDisableUser) | **Put** /ocs/v2.php/cloud/users/{userId}/disable | Disable a user
[**ProvisioningApiUsersEditUser**](ProvisioningApiUsersApi.md#ProvisioningApiUsersEditUser) | **Put** /ocs/v2.php/cloud/users/{userId} | Update a value of the user&#39;s details
[**ProvisioningApiUsersEditUserMultiValue**](ProvisioningApiUsersApi.md#ProvisioningApiUsersEditUserMultiValue) | **Put** /ocs/v2.php/cloud/users/{userId}/{collectionName} | Update multiple values of the user&#39;s details
[**ProvisioningApiUsersEnableUser**](ProvisioningApiUsersApi.md#ProvisioningApiUsersEnableUser) | **Put** /ocs/v2.php/cloud/users/{userId}/enable | Enable a user
[**ProvisioningApiUsersGetCurrentUser**](ProvisioningApiUsersApi.md#ProvisioningApiUsersGetCurrentUser) | **Get** /ocs/v2.php/cloud/user | Get the details of the current user
[**ProvisioningApiUsersGetEditableFields**](ProvisioningApiUsersApi.md#ProvisioningApiUsersGetEditableFields) | **Get** /ocs/v2.php/cloud/user/fields | Get a list of fields that are editable for the current user
[**ProvisioningApiUsersGetEditableFieldsForUser**](ProvisioningApiUsersApi.md#ProvisioningApiUsersGetEditableFieldsForUser) | **Get** /ocs/v2.php/cloud/user/fields/{userId} | Get a list of fields that are editable for a user
[**ProvisioningApiUsersGetUser**](ProvisioningApiUsersApi.md#ProvisioningApiUsersGetUser) | **Get** /ocs/v2.php/cloud/users/{userId} | Get the details of a user
[**ProvisioningApiUsersGetUserSubAdminGroups**](ProvisioningApiUsersApi.md#ProvisioningApiUsersGetUserSubAdminGroups) | **Get** /ocs/v2.php/cloud/users/{userId}/subadmins | Get the groups a user is a subadmin of
[**ProvisioningApiUsersGetUsers**](ProvisioningApiUsersApi.md#ProvisioningApiUsersGetUsers) | **Get** /ocs/v2.php/cloud/users | Get a list of users
[**ProvisioningApiUsersGetUsersDetails**](ProvisioningApiUsersApi.md#ProvisioningApiUsersGetUsersDetails) | **Get** /ocs/v2.php/cloud/users/details | Get a list of users and their details
[**ProvisioningApiUsersGetUsersGroups**](ProvisioningApiUsersApi.md#ProvisioningApiUsersGetUsersGroups) | **Get** /ocs/v2.php/cloud/users/{userId}/groups | Get a list of groups the user belongs to
[**ProvisioningApiUsersRemoveFromGroup**](ProvisioningApiUsersApi.md#ProvisioningApiUsersRemoveFromGroup) | **Delete** /ocs/v2.php/cloud/users/{userId}/groups | Remove a user from a group
[**ProvisioningApiUsersRemoveSubAdmin**](ProvisioningApiUsersApi.md#ProvisioningApiUsersRemoveSubAdmin) | **Delete** /ocs/v2.php/cloud/users/{userId}/subadmins | Remove a user from the subadmins of a group
[**ProvisioningApiUsersResendWelcomeMessage**](ProvisioningApiUsersApi.md#ProvisioningApiUsersResendWelcomeMessage) | **Post** /ocs/v2.php/cloud/users/{userId}/welcome | Resend the welcome message
[**ProvisioningApiUsersSearchByPhoneNumbers**](ProvisioningApiUsersApi.md#ProvisioningApiUsersSearchByPhoneNumbers) | **Post** /ocs/v2.php/cloud/users/search/by-phone | Search users by their phone numbers
[**ProvisioningApiUsersWipeUserDevices**](ProvisioningApiUsersApi.md#ProvisioningApiUsersWipeUserDevices) | **Post** /ocs/v2.php/cloud/users/{userId}/wipe | Wipe all devices of a user



## ProvisioningApiUsersAddSubAdmin

> CoreWhatsNewDismiss200Response ProvisioningApiUsersAddSubAdmin(ctx, userId).Groupid(groupid).OCSAPIRequest(oCSAPIRequest).Execute()

Make a user a subadmin of a group



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
    userId := "userId_example" // string | ID of the user
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvisioningApiUsersApi.ProvisioningApiUsersAddSubAdmin(context.Background(), userId).Groupid(groupid).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningApiUsersApi.ProvisioningApiUsersAddSubAdmin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisioningApiUsersAddSubAdmin`: CoreWhatsNewDismiss200Response
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningApiUsersApi.ProvisioningApiUsersAddSubAdmin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of the user | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvisioningApiUsersAddSubAdminRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupid** | **string** | ID of the group | 

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


## ProvisioningApiUsersAddToGroup

> CoreWhatsNewDismiss200Response ProvisioningApiUsersAddToGroup(ctx, userId).OCSAPIRequest(oCSAPIRequest).Groupid(groupid).Execute()

Add a user to a group

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
    groupid := "groupid_example" // string | ID of the group (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvisioningApiUsersApi.ProvisioningApiUsersAddToGroup(context.Background(), userId).OCSAPIRequest(oCSAPIRequest).Groupid(groupid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningApiUsersApi.ProvisioningApiUsersAddToGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisioningApiUsersAddToGroup`: CoreWhatsNewDismiss200Response
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningApiUsersApi.ProvisioningApiUsersAddToGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of the user | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvisioningApiUsersAddToGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]
 **groupid** | **string** | ID of the group | [default to &quot;&quot;]

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


## ProvisioningApiUsersAddUser

> ProvisioningApiUsersAddUser200Response ProvisioningApiUsersAddUser(ctx).Userid(userid).OCSAPIRequest(oCSAPIRequest).Password(password).DisplayName(displayName).Email(email).Groups(groups).Subadmin(subadmin).Quota(quota).Language(language).Manager(manager).Execute()

Create a new user

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
    userid := "userid_example" // string | ID of the user
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")
    password := "password_example" // string | Password of the user (optional) (default to "")
    displayName := "displayName_example" // string | Display name of the user (optional) (default to "")
    email := "email_example" // string | Email of the user (optional) (default to "")
    groups := "groups_example" // string | Groups of the user (optional)
    subadmin := "subadmin_example" // string | Groups where the user is subadmin (optional)
    quota := "quota_example" // string | Quota of the user (optional) (default to "")
    language := "language_example" // string | Language of the user (optional) (default to "")
    manager := "manager_example" // string | Manager of the user (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvisioningApiUsersApi.ProvisioningApiUsersAddUser(context.Background()).Userid(userid).OCSAPIRequest(oCSAPIRequest).Password(password).DisplayName(displayName).Email(email).Groups(groups).Subadmin(subadmin).Quota(quota).Language(language).Manager(manager).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningApiUsersApi.ProvisioningApiUsersAddUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisioningApiUsersAddUser`: ProvisioningApiUsersAddUser200Response
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningApiUsersApi.ProvisioningApiUsersAddUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProvisioningApiUsersAddUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userid** | **string** | ID of the user | 
 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]
 **password** | **string** | Password of the user | [default to &quot;&quot;]
 **displayName** | **string** | Display name of the user | [default to &quot;&quot;]
 **email** | **string** | Email of the user | [default to &quot;&quot;]
 **groups** | **string** | Groups of the user | 
 **subadmin** | **string** | Groups where the user is subadmin | 
 **quota** | **string** | Quota of the user | [default to &quot;&quot;]
 **language** | **string** | Language of the user | [default to &quot;&quot;]
 **manager** | **string** | Manager of the user | 

### Return type

[**ProvisioningApiUsersAddUser200Response**](ProvisioningApiUsersAddUser200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvisioningApiUsersDeleteUser

> CoreWhatsNewDismiss200Response ProvisioningApiUsersDeleteUser(ctx, userId).OCSAPIRequest(oCSAPIRequest).Execute()

Delete a user

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
    resp, r, err := apiClient.ProvisioningApiUsersApi.ProvisioningApiUsersDeleteUser(context.Background(), userId).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningApiUsersApi.ProvisioningApiUsersDeleteUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisioningApiUsersDeleteUser`: CoreWhatsNewDismiss200Response
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningApiUsersApi.ProvisioningApiUsersDeleteUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of the user | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvisioningApiUsersDeleteUserRequest struct via the builder pattern


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


## ProvisioningApiUsersDisableUser

> CoreWhatsNewDismiss200Response ProvisioningApiUsersDisableUser(ctx, userId).OCSAPIRequest(oCSAPIRequest).Execute()

Disable a user

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
    resp, r, err := apiClient.ProvisioningApiUsersApi.ProvisioningApiUsersDisableUser(context.Background(), userId).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningApiUsersApi.ProvisioningApiUsersDisableUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisioningApiUsersDisableUser`: CoreWhatsNewDismiss200Response
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningApiUsersApi.ProvisioningApiUsersDisableUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of the user | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvisioningApiUsersDisableUserRequest struct via the builder pattern


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


## ProvisioningApiUsersEditUser

> CoreWhatsNewDismiss200Response ProvisioningApiUsersEditUser(ctx, userId).Key(key).Value(value).OCSAPIRequest(oCSAPIRequest).Execute()

Update a value of the user's details

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
    key := "key_example" // string | Key that will be updated
    value := "value_example" // string | New value for the key
    userId := "userId_example" // string | ID of the user
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvisioningApiUsersApi.ProvisioningApiUsersEditUser(context.Background(), userId).Key(key).Value(value).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningApiUsersApi.ProvisioningApiUsersEditUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisioningApiUsersEditUser`: CoreWhatsNewDismiss200Response
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningApiUsersApi.ProvisioningApiUsersEditUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of the user | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvisioningApiUsersEditUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **key** | **string** | Key that will be updated | 
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


## ProvisioningApiUsersEditUserMultiValue

> CoreWhatsNewDismiss200Response ProvisioningApiUsersEditUserMultiValue(ctx, userId, collectionName).Key(key).Value(value).OCSAPIRequest(oCSAPIRequest).Execute()

Update multiple values of the user's details

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
    key := "key_example" // string | Key that will be updated
    value := "value_example" // string | New value for the key
    userId := "userId_example" // string | ID of the user
    collectionName := "collectionName_example" // string | Collection to update
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvisioningApiUsersApi.ProvisioningApiUsersEditUserMultiValue(context.Background(), userId, collectionName).Key(key).Value(value).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningApiUsersApi.ProvisioningApiUsersEditUserMultiValue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisioningApiUsersEditUserMultiValue`: CoreWhatsNewDismiss200Response
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningApiUsersApi.ProvisioningApiUsersEditUserMultiValue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of the user | 
**collectionName** | **string** | Collection to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvisioningApiUsersEditUserMultiValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **key** | **string** | Key that will be updated | 
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


## ProvisioningApiUsersEnableUser

> CoreWhatsNewDismiss200Response ProvisioningApiUsersEnableUser(ctx, userId).OCSAPIRequest(oCSAPIRequest).Execute()

Enable a user

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
    resp, r, err := apiClient.ProvisioningApiUsersApi.ProvisioningApiUsersEnableUser(context.Background(), userId).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningApiUsersApi.ProvisioningApiUsersEnableUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisioningApiUsersEnableUser`: CoreWhatsNewDismiss200Response
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningApiUsersApi.ProvisioningApiUsersEnableUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of the user | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvisioningApiUsersEnableUserRequest struct via the builder pattern


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


## ProvisioningApiUsersGetCurrentUser

> ProvisioningApiUsersGetUser200Response ProvisioningApiUsersGetCurrentUser(ctx).OCSAPIRequest(oCSAPIRequest).Execute()

Get the details of the current user

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
    resp, r, err := apiClient.ProvisioningApiUsersApi.ProvisioningApiUsersGetCurrentUser(context.Background()).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningApiUsersApi.ProvisioningApiUsersGetCurrentUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisioningApiUsersGetCurrentUser`: ProvisioningApiUsersGetUser200Response
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningApiUsersApi.ProvisioningApiUsersGetCurrentUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProvisioningApiUsersGetCurrentUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]

### Return type

[**ProvisioningApiUsersGetUser200Response**](ProvisioningApiUsersGetUser200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvisioningApiUsersGetEditableFields

> ProvisioningApiGroupsGetSubAdminsOfGroup200Response ProvisioningApiUsersGetEditableFields(ctx).OCSAPIRequest(oCSAPIRequest).Execute()

Get a list of fields that are editable for the current user

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
    resp, r, err := apiClient.ProvisioningApiUsersApi.ProvisioningApiUsersGetEditableFields(context.Background()).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningApiUsersApi.ProvisioningApiUsersGetEditableFields``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisioningApiUsersGetEditableFields`: ProvisioningApiGroupsGetSubAdminsOfGroup200Response
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningApiUsersApi.ProvisioningApiUsersGetEditableFields`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProvisioningApiUsersGetEditableFieldsRequest struct via the builder pattern


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


## ProvisioningApiUsersGetEditableFieldsForUser

> ProvisioningApiGroupsGetSubAdminsOfGroup200Response ProvisioningApiUsersGetEditableFieldsForUser(ctx, userId).OCSAPIRequest(oCSAPIRequest).Execute()

Get a list of fields that are editable for a user

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
    resp, r, err := apiClient.ProvisioningApiUsersApi.ProvisioningApiUsersGetEditableFieldsForUser(context.Background(), userId).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningApiUsersApi.ProvisioningApiUsersGetEditableFieldsForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisioningApiUsersGetEditableFieldsForUser`: ProvisioningApiGroupsGetSubAdminsOfGroup200Response
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningApiUsersApi.ProvisioningApiUsersGetEditableFieldsForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of the user | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvisioningApiUsersGetEditableFieldsForUserRequest struct via the builder pattern


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


## ProvisioningApiUsersGetUser

> ProvisioningApiUsersGetUser200Response ProvisioningApiUsersGetUser(ctx, userId).OCSAPIRequest(oCSAPIRequest).Execute()

Get the details of a user

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
    resp, r, err := apiClient.ProvisioningApiUsersApi.ProvisioningApiUsersGetUser(context.Background(), userId).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningApiUsersApi.ProvisioningApiUsersGetUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisioningApiUsersGetUser`: ProvisioningApiUsersGetUser200Response
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningApiUsersApi.ProvisioningApiUsersGetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of the user | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvisioningApiUsersGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]

### Return type

[**ProvisioningApiUsersGetUser200Response**](ProvisioningApiUsersGetUser200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvisioningApiUsersGetUserSubAdminGroups

> ProvisioningApiGroupsGetSubAdminsOfGroup200Response ProvisioningApiUsersGetUserSubAdminGroups(ctx, userId).OCSAPIRequest(oCSAPIRequest).Execute()

Get the groups a user is a subadmin of



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
    userId := "userId_example" // string | ID if the user
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvisioningApiUsersApi.ProvisioningApiUsersGetUserSubAdminGroups(context.Background(), userId).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningApiUsersApi.ProvisioningApiUsersGetUserSubAdminGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisioningApiUsersGetUserSubAdminGroups`: ProvisioningApiGroupsGetSubAdminsOfGroup200Response
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningApiUsersApi.ProvisioningApiUsersGetUserSubAdminGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID if the user | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvisioningApiUsersGetUserSubAdminGroupsRequest struct via the builder pattern


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


## ProvisioningApiUsersGetUsers

> ProvisioningApiGroupsGetGroupUsers200Response ProvisioningApiUsersGetUsers(ctx).OCSAPIRequest(oCSAPIRequest).Search(search).Limit(limit).Offset(offset).Execute()

Get a list of users

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
    resp, r, err := apiClient.ProvisioningApiUsersApi.ProvisioningApiUsersGetUsers(context.Background()).OCSAPIRequest(oCSAPIRequest).Search(search).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningApiUsersApi.ProvisioningApiUsersGetUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisioningApiUsersGetUsers`: ProvisioningApiGroupsGetGroupUsers200Response
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningApiUsersApi.ProvisioningApiUsersGetUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProvisioningApiUsersGetUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]
 **search** | **string** | Text to search for | [default to &quot;&quot;]
 **limit** | **int64** | Limit the amount of groups returned | 
 **offset** | **int64** | Offset for searching for groups | [default to 0]

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


## ProvisioningApiUsersGetUsersDetails

> ProvisioningApiGroupsGetGroupUsersDetails200Response ProvisioningApiUsersGetUsersDetails(ctx).OCSAPIRequest(oCSAPIRequest).Search(search).Limit(limit).Offset(offset).Execute()

Get a list of users and their details

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
    resp, r, err := apiClient.ProvisioningApiUsersApi.ProvisioningApiUsersGetUsersDetails(context.Background()).OCSAPIRequest(oCSAPIRequest).Search(search).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningApiUsersApi.ProvisioningApiUsersGetUsersDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisioningApiUsersGetUsersDetails`: ProvisioningApiGroupsGetGroupUsersDetails200Response
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningApiUsersApi.ProvisioningApiUsersGetUsersDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProvisioningApiUsersGetUsersDetailsRequest struct via the builder pattern


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


## ProvisioningApiUsersGetUsersGroups

> ProvisioningApiGroupsGetGroups200Response ProvisioningApiUsersGetUsersGroups(ctx, userId).OCSAPIRequest(oCSAPIRequest).Execute()

Get a list of groups the user belongs to

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
    resp, r, err := apiClient.ProvisioningApiUsersApi.ProvisioningApiUsersGetUsersGroups(context.Background(), userId).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningApiUsersApi.ProvisioningApiUsersGetUsersGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisioningApiUsersGetUsersGroups`: ProvisioningApiGroupsGetGroups200Response
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningApiUsersApi.ProvisioningApiUsersGetUsersGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of the user | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvisioningApiUsersGetUsersGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]

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


## ProvisioningApiUsersRemoveFromGroup

> CoreWhatsNewDismiss200Response ProvisioningApiUsersRemoveFromGroup(ctx, userId).Groupid(groupid).OCSAPIRequest(oCSAPIRequest).Execute()

Remove a user from a group

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
    userId := "userId_example" // string | ID of the user
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvisioningApiUsersApi.ProvisioningApiUsersRemoveFromGroup(context.Background(), userId).Groupid(groupid).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningApiUsersApi.ProvisioningApiUsersRemoveFromGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisioningApiUsersRemoveFromGroup`: CoreWhatsNewDismiss200Response
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningApiUsersApi.ProvisioningApiUsersRemoveFromGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of the user | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvisioningApiUsersRemoveFromGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupid** | **string** | ID of the group | 

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


## ProvisioningApiUsersRemoveSubAdmin

> CoreWhatsNewDismiss200Response ProvisioningApiUsersRemoveSubAdmin(ctx, userId).Groupid(groupid).OCSAPIRequest(oCSAPIRequest).Execute()

Remove a user from the subadmins of a group



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
    userId := "userId_example" // string | ID of the user
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvisioningApiUsersApi.ProvisioningApiUsersRemoveSubAdmin(context.Background(), userId).Groupid(groupid).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningApiUsersApi.ProvisioningApiUsersRemoveSubAdmin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisioningApiUsersRemoveSubAdmin`: CoreWhatsNewDismiss200Response
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningApiUsersApi.ProvisioningApiUsersRemoveSubAdmin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of the user | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvisioningApiUsersRemoveSubAdminRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupid** | **string** | ID of the group | 

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


## ProvisioningApiUsersResendWelcomeMessage

> CoreWhatsNewDismiss200Response ProvisioningApiUsersResendWelcomeMessage(ctx, userId).OCSAPIRequest(oCSAPIRequest).Execute()

Resend the welcome message

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
    userId := "userId_example" // string | ID if the user
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvisioningApiUsersApi.ProvisioningApiUsersResendWelcomeMessage(context.Background(), userId).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningApiUsersApi.ProvisioningApiUsersResendWelcomeMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisioningApiUsersResendWelcomeMessage`: CoreWhatsNewDismiss200Response
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningApiUsersApi.ProvisioningApiUsersResendWelcomeMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID if the user | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvisioningApiUsersResendWelcomeMessageRequest struct via the builder pattern


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


## ProvisioningApiUsersSearchByPhoneNumbers

> ProvisioningApiUsersSearchByPhoneNumbers200Response ProvisioningApiUsersSearchByPhoneNumbers(ctx).Location(location).Search(search).OCSAPIRequest(oCSAPIRequest).Execute()

Search users by their phone numbers

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
    location := "location_example" // string | Location of the phone number (for country code)
    search := "search_example" // string | Phone numbers to search for
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvisioningApiUsersApi.ProvisioningApiUsersSearchByPhoneNumbers(context.Background()).Location(location).Search(search).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningApiUsersApi.ProvisioningApiUsersSearchByPhoneNumbers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisioningApiUsersSearchByPhoneNumbers`: ProvisioningApiUsersSearchByPhoneNumbers200Response
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningApiUsersApi.ProvisioningApiUsersSearchByPhoneNumbers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProvisioningApiUsersSearchByPhoneNumbersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **location** | **string** | Location of the phone number (for country code) | 
 **search** | **string** | Phone numbers to search for | 
 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]

### Return type

[**ProvisioningApiUsersSearchByPhoneNumbers200Response**](ProvisioningApiUsersSearchByPhoneNumbers200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvisioningApiUsersWipeUserDevices

> CoreWhatsNewDismiss200Response ProvisioningApiUsersWipeUserDevices(ctx, userId).OCSAPIRequest(oCSAPIRequest).Execute()

Wipe all devices of a user

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
    resp, r, err := apiClient.ProvisioningApiUsersApi.ProvisioningApiUsersWipeUserDevices(context.Background(), userId).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningApiUsersApi.ProvisioningApiUsersWipeUserDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisioningApiUsersWipeUserDevices`: CoreWhatsNewDismiss200Response
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningApiUsersApi.ProvisioningApiUsersWipeUserDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of the user | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvisioningApiUsersWipeUserDevicesRequest struct via the builder pattern


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

