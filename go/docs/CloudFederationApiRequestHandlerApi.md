# \CloudFederationApiRequestHandlerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudFederationApiRequestHandlerAddShare**](CloudFederationApiRequestHandlerApi.md#CloudFederationApiRequestHandlerAddShare) | **Post** /index.php/ocm/shares | Add share
[**CloudFederationApiRequestHandlerReceiveNotification**](CloudFederationApiRequestHandlerApi.md#CloudFederationApiRequestHandlerReceiveNotification) | **Post** /index.php/ocm/notifications | Send a notification about an existing share



## CloudFederationApiRequestHandlerAddShare

> CloudFederationApiAddShare CloudFederationApiRequestHandlerAddShare(ctx).ShareWith(shareWith).Name(name).ProviderId(providerId).Owner(owner).Protocol(protocol).ShareType(shareType).ResourceType(resourceType).Description(description).OwnerDisplayName(ownerDisplayName).SharedBy(sharedBy).SharedByDisplayName(sharedByDisplayName).Execute()

Add share

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
    shareWith := "shareWith_example" // string | The user who the share will be shared with
    name := "name_example" // string | The resource name (e.g. document.odt)
    providerId := "providerId_example" // string | Resource UID on the provider side
    owner := "owner_example" // string | Provider specific UID of the user who owns the resource
    protocol := "protocol_example" // string | e,.g. ['name' => 'webdav', 'options' => ['username' => 'john', 'permissions' => 31]]
    shareType := "shareType_example" // string | 'group' or 'user' share
    resourceType := "resourceType_example" // string | 'file', 'calendar',...
    description := "description_example" // string | Share description (optional)
    ownerDisplayName := "ownerDisplayName_example" // string | Display name of the user who shared the item (optional)
    sharedBy := "sharedBy_example" // string | Provider specific UID of the user who shared the resource (optional)
    sharedByDisplayName := "sharedByDisplayName_example" // string | Display name of the user who shared the resource (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudFederationApiRequestHandlerApi.CloudFederationApiRequestHandlerAddShare(context.Background()).ShareWith(shareWith).Name(name).ProviderId(providerId).Owner(owner).Protocol(protocol).ShareType(shareType).ResourceType(resourceType).Description(description).OwnerDisplayName(ownerDisplayName).SharedBy(sharedBy).SharedByDisplayName(sharedByDisplayName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudFederationApiRequestHandlerApi.CloudFederationApiRequestHandlerAddShare``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloudFederationApiRequestHandlerAddShare`: CloudFederationApiAddShare
    fmt.Fprintf(os.Stdout, "Response from `CloudFederationApiRequestHandlerApi.CloudFederationApiRequestHandlerAddShare`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudFederationApiRequestHandlerAddShareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **shareWith** | **string** | The user who the share will be shared with | 
 **name** | **string** | The resource name (e.g. document.odt) | 
 **providerId** | **string** | Resource UID on the provider side | 
 **owner** | **string** | Provider specific UID of the user who owns the resource | 
 **protocol** | **string** | e,.g. [&#39;name&#39; &#x3D;&gt; &#39;webdav&#39;, &#39;options&#39; &#x3D;&gt; [&#39;username&#39; &#x3D;&gt; &#39;john&#39;, &#39;permissions&#39; &#x3D;&gt; 31]] | 
 **shareType** | **string** | &#39;group&#39; or &#39;user&#39; share | 
 **resourceType** | **string** | &#39;file&#39;, &#39;calendar&#39;,... | 
 **description** | **string** | Share description | 
 **ownerDisplayName** | **string** | Display name of the user who shared the item | 
 **sharedBy** | **string** | Provider specific UID of the user who shared the resource | 
 **sharedByDisplayName** | **string** | Display name of the user who shared the resource | 

### Return type

[**CloudFederationApiAddShare**](CloudFederationApiAddShare.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudFederationApiRequestHandlerReceiveNotification

> []map[string]interface{} CloudFederationApiRequestHandlerReceiveNotification(ctx).NotificationType(notificationType).ResourceType(resourceType).ProviderId(providerId).Notification(notification).Execute()

Send a notification about an existing share

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
    notificationType := "notificationType_example" // string | Notification type, e.g. SHARE_ACCEPTED
    resourceType := "resourceType_example" // string | calendar, file, contact,...
    providerId := "providerId_example" // string | ID of the share (optional)
    notification := "notification_example" // string | The actual payload of the notification (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudFederationApiRequestHandlerApi.CloudFederationApiRequestHandlerReceiveNotification(context.Background()).NotificationType(notificationType).ResourceType(resourceType).ProviderId(providerId).Notification(notification).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudFederationApiRequestHandlerApi.CloudFederationApiRequestHandlerReceiveNotification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloudFederationApiRequestHandlerReceiveNotification`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CloudFederationApiRequestHandlerApi.CloudFederationApiRequestHandlerReceiveNotification`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudFederationApiRequestHandlerReceiveNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **notificationType** | **string** | Notification type, e.g. SHARE_ACCEPTED | 
 **resourceType** | **string** | calendar, file, contact,... | 
 **providerId** | **string** | ID of the share | 
 **notification** | **string** | The actual payload of the notification | 

### Return type

**[]map[string]interface{}**

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

