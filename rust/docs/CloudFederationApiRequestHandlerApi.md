# \CloudFederationApiRequestHandlerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**cloud_federation_api_request_handler_add_share**](CloudFederationApiRequestHandlerApi.md#cloud_federation_api_request_handler_add_share) | **POST** /index.php/ocm/shares | Add share
[**cloud_federation_api_request_handler_receive_notification**](CloudFederationApiRequestHandlerApi.md#cloud_federation_api_request_handler_receive_notification) | **POST** /index.php/ocm/notifications | Send a notification about an existing share



## cloud_federation_api_request_handler_add_share

> crate::models::CloudFederationApiAddShare cloud_federation_api_request_handler_add_share(share_with, name, provider_id, owner, protocol, share_type, resource_type, description, owner_display_name, shared_by, shared_by_display_name)
Add share

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**share_with** | **String** | The user who the share will be shared with | [required] |
**name** | **String** | The resource name (e.g. document.odt) | [required] |
**provider_id** | **String** | Resource UID on the provider side | [required] |
**owner** | **String** | Provider specific UID of the user who owns the resource | [required] |
**protocol** | **String** | e,.g. ['name' => 'webdav', 'options' => ['username' => 'john', 'permissions' => 31]] | [required] |
**share_type** | **String** | 'group' or 'user' share | [required] |
**resource_type** | **String** | 'file', 'calendar',... | [required] |
**description** | Option<**String**> | Share description |  |
**owner_display_name** | Option<**String**> | Display name of the user who shared the item |  |
**shared_by** | Option<**String**> | Provider specific UID of the user who shared the resource |  |
**shared_by_display_name** | Option<**String**> | Display name of the user who shared the resource |  |

### Return type

[**crate::models::CloudFederationApiAddShare**](CloudFederationApiAddShare.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## cloud_federation_api_request_handler_receive_notification

> Vec<serde_json::Value> cloud_federation_api_request_handler_receive_notification(notification_type, resource_type, provider_id, notification)
Send a notification about an existing share

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**notification_type** | **String** | Notification type, e.g. SHARE_ACCEPTED | [required] |
**resource_type** | **String** | calendar, file, contact,... | [required] |
**provider_id** | Option<**String**> | ID of the share |  |
**notification** | Option<**String**> | The actual payload of the notification |  |

### Return type

[**Vec<serde_json::Value>**](serde_json::Value.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

