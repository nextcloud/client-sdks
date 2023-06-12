# openapi_client.CloudFederationApiRequestHandlerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**cloud_federation_api_request_handler_add_share**](CloudFederationApiRequestHandlerApi.md#cloud_federation_api_request_handler_add_share) | **POST** /index.php/ocm/shares | Add share
[**cloud_federation_api_request_handler_receive_notification**](CloudFederationApiRequestHandlerApi.md#cloud_federation_api_request_handler_receive_notification) | **POST** /index.php/ocm/notifications | Send a notification about an existing share


# **cloud_federation_api_request_handler_add_share**
> CloudFederationApiAddShare cloud_federation_api_request_handler_add_share(share_with, name, provider_id, owner, protocol, share_type, resource_type, description=description, owner_display_name=owner_display_name, shared_by=shared_by, shared_by_display_name=shared_by_display_name)

Add share

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import openapi_client
from openapi_client.models.cloud_federation_api_add_share import CloudFederationApiAddShare
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "http://localhost"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure HTTP basic authorization: basic_auth
configuration = openapi_client.Configuration(
    username = os.environ["USERNAME"],
    password = os.environ["PASSWORD"]
)

# Configure Bearer authorization: bearer_auth
configuration = openapi_client.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.CloudFederationApiRequestHandlerApi(api_client)
    share_with = 'share_with_example' # str | The user who the share will be shared with
    name = 'name_example' # str | The resource name (e.g. document.odt)
    provider_id = 'provider_id_example' # str | Resource UID on the provider side
    owner = 'owner_example' # str | Provider specific UID of the user who owns the resource
    protocol = 'protocol_example' # str | e,.g. ['name' => 'webdav', 'options' => ['username' => 'john', 'permissions' => 31]]
    share_type = 'share_type_example' # str | 'group' or 'user' share
    resource_type = 'resource_type_example' # str | 'file', 'calendar',...
    description = 'description_example' # str | Share description (optional)
    owner_display_name = 'owner_display_name_example' # str | Display name of the user who shared the item (optional)
    shared_by = 'shared_by_example' # str | Provider specific UID of the user who shared the resource (optional)
    shared_by_display_name = 'shared_by_display_name_example' # str | Display name of the user who shared the resource (optional)

    try:
        # Add share
        api_response = api_instance.cloud_federation_api_request_handler_add_share(share_with, name, provider_id, owner, protocol, share_type, resource_type, description=description, owner_display_name=owner_display_name, shared_by=shared_by, shared_by_display_name=shared_by_display_name)
        print("The response of CloudFederationApiRequestHandlerApi->cloud_federation_api_request_handler_add_share:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CloudFederationApiRequestHandlerApi->cloud_federation_api_request_handler_add_share: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **share_with** | **str**| The user who the share will be shared with | 
 **name** | **str**| The resource name (e.g. document.odt) | 
 **provider_id** | **str**| Resource UID on the provider side | 
 **owner** | **str**| Provider specific UID of the user who owns the resource | 
 **protocol** | **str**| e,.g. [&#39;name&#39; &#x3D;&gt; &#39;webdav&#39;, &#39;options&#39; &#x3D;&gt; [&#39;username&#39; &#x3D;&gt; &#39;john&#39;, &#39;permissions&#39; &#x3D;&gt; 31]] | 
 **share_type** | **str**| &#39;group&#39; or &#39;user&#39; share | 
 **resource_type** | **str**| &#39;file&#39;, &#39;calendar&#39;,... | 
 **description** | **str**| Share description | [optional] 
 **owner_display_name** | **str**| Display name of the user who shared the item | [optional] 
 **shared_by** | **str**| Provider specific UID of the user who shared the resource | [optional] 
 **shared_by_display_name** | **str**| Display name of the user who shared the resource | [optional] 

### Return type

[**CloudFederationApiAddShare**](CloudFederationApiAddShare.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**201** | The notification was successfully received. The display name of the recipient might be returned in the body |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **cloud_federation_api_request_handler_receive_notification**
> List[object] cloud_federation_api_request_handler_receive_notification(notification_type, resource_type, provider_id=provider_id, notification=notification)

Send a notification about an existing share

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import openapi_client
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "http://localhost"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure HTTP basic authorization: basic_auth
configuration = openapi_client.Configuration(
    username = os.environ["USERNAME"],
    password = os.environ["PASSWORD"]
)

# Configure Bearer authorization: bearer_auth
configuration = openapi_client.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.CloudFederationApiRequestHandlerApi(api_client)
    notification_type = 'notification_type_example' # str | Notification type, e.g. SHARE_ACCEPTED
    resource_type = 'resource_type_example' # str | calendar, file, contact,...
    provider_id = 'provider_id_example' # str | ID of the share (optional)
    notification = 'notification_example' # str | The actual payload of the notification (optional)

    try:
        # Send a notification about an existing share
        api_response = api_instance.cloud_federation_api_request_handler_receive_notification(notification_type, resource_type, provider_id=provider_id, notification=notification)
        print("The response of CloudFederationApiRequestHandlerApi->cloud_federation_api_request_handler_receive_notification:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CloudFederationApiRequestHandlerApi->cloud_federation_api_request_handler_receive_notification: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **notification_type** | **str**| Notification type, e.g. SHARE_ACCEPTED | 
 **resource_type** | **str**| calendar, file, contact,... | 
 **provider_id** | **str**| ID of the share | [optional] 
 **notification** | **str**| The actual payload of the notification | [optional] 

### Return type

**List[object]**

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**201** | The notification was successfully received |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

