# nextcloud_client.UserStatusUserStatusApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**user_status_user_status_clear_message**](UserStatusUserStatusApi.md#user_status_user_status_clear_message) | **DELETE** /ocs/v2.php/apps/user_status/api/v1/user_status/message | Clear the message of the current user
[**user_status_user_status_get_status**](UserStatusUserStatusApi.md#user_status_user_status_get_status) | **GET** /ocs/v2.php/apps/user_status/api/v1/user_status | Get the status of the current user
[**user_status_user_status_revert_status**](UserStatusUserStatusApi.md#user_status_user_status_revert_status) | **DELETE** /ocs/v2.php/apps/user_status/api/v1/user_status/revert/{messageId} | Revert the status to the previous status
[**user_status_user_status_set_custom_message**](UserStatusUserStatusApi.md#user_status_user_status_set_custom_message) | **PUT** /ocs/v2.php/apps/user_status/api/v1/user_status/message/custom | Set the message to a custom message for the current user
[**user_status_user_status_set_predefined_message**](UserStatusUserStatusApi.md#user_status_user_status_set_predefined_message) | **PUT** /ocs/v2.php/apps/user_status/api/v1/user_status/message/predefined | Set the message to a predefined message for the current user
[**user_status_user_status_set_status**](UserStatusUserStatusApi.md#user_status_user_status_set_status) | **PUT** /ocs/v2.php/apps/user_status/api/v1/user_status/status | Update the status type of the current user


# **user_status_user_status_clear_message**
> CoreWhatsNewDismiss200Response user_status_user_status_clear_message(ocs_api_request)

Clear the message of the current user

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.core_whats_new_dismiss200_response import CoreWhatsNewDismiss200Response
from nextcloud_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = nextcloud_client.Configuration(
    host = "http://localhost"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure HTTP basic authorization: basic_auth
configuration = nextcloud_client.Configuration(
    username = os.environ["USERNAME"],
    password = os.environ["PASSWORD"]
)

# Configure Bearer authorization: bearer_auth
configuration = nextcloud_client.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with nextcloud_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = nextcloud_client.UserStatusUserStatusApi(api_client)
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Clear the message of the current user
        api_response = api_instance.user_status_user_status_clear_message(ocs_api_request)
        print("The response of UserStatusUserStatusApi->user_status_user_status_clear_message:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling UserStatusUserStatusApi->user_status_user_status_clear_message: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]

### Return type

[**CoreWhatsNewDismiss200Response**](CoreWhatsNewDismiss200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** |  |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **user_status_user_status_get_status**
> UserStatusUserStatusGetStatus200Response user_status_user_status_get_status(ocs_api_request)

Get the status of the current user

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.user_status_user_status_get_status200_response import UserStatusUserStatusGetStatus200Response
from nextcloud_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = nextcloud_client.Configuration(
    host = "http://localhost"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure HTTP basic authorization: basic_auth
configuration = nextcloud_client.Configuration(
    username = os.environ["USERNAME"],
    password = os.environ["PASSWORD"]
)

# Configure Bearer authorization: bearer_auth
configuration = nextcloud_client.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with nextcloud_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = nextcloud_client.UserStatusUserStatusApi(api_client)
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Get the status of the current user
        api_response = api_instance.user_status_user_status_get_status(ocs_api_request)
        print("The response of UserStatusUserStatusApi->user_status_user_status_get_status:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling UserStatusUserStatusApi->user_status_user_status_get_status: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]

### Return type

[**UserStatusUserStatusGetStatus200Response**](UserStatusUserStatusGetStatus200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | The status was found successfully |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **user_status_user_status_revert_status**
> UserStatusUserStatusRevertStatus200Response user_status_user_status_revert_status(message_id, ocs_api_request)

Revert the status to the previous status

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.user_status_user_status_revert_status200_response import UserStatusUserStatusRevertStatus200Response
from nextcloud_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = nextcloud_client.Configuration(
    host = "http://localhost"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure HTTP basic authorization: basic_auth
configuration = nextcloud_client.Configuration(
    username = os.environ["USERNAME"],
    password = os.environ["PASSWORD"]
)

# Configure Bearer authorization: bearer_auth
configuration = nextcloud_client.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with nextcloud_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = nextcloud_client.UserStatusUserStatusApi(api_client)
    message_id = 'message_id_example' # str | ID of the message to delete
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Revert the status to the previous status
        api_response = api_instance.user_status_user_status_revert_status(message_id, ocs_api_request)
        print("The response of UserStatusUserStatusApi->user_status_user_status_revert_status:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling UserStatusUserStatusApi->user_status_user_status_revert_status: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **message_id** | **str**| ID of the message to delete | 
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]

### Return type

[**UserStatusUserStatusRevertStatus200Response**](UserStatusUserStatusRevertStatus200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Status reverted |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **user_status_user_status_set_custom_message**
> UserStatusUserStatusGetStatus200Response user_status_user_status_set_custom_message(ocs_api_request, status_icon=status_icon, message=message, clear_at=clear_at)

Set the message to a custom message for the current user

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.user_status_user_status_get_status200_response import UserStatusUserStatusGetStatus200Response
from nextcloud_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = nextcloud_client.Configuration(
    host = "http://localhost"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure HTTP basic authorization: basic_auth
configuration = nextcloud_client.Configuration(
    username = os.environ["USERNAME"],
    password = os.environ["PASSWORD"]
)

# Configure Bearer authorization: bearer_auth
configuration = nextcloud_client.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with nextcloud_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = nextcloud_client.UserStatusUserStatusApi(api_client)
    ocs_api_request = 'true' # str |  (default to 'true')
    status_icon = 'status_icon_example' # str | Icon of the status (optional)
    message = 'message_example' # str | Message of the status (optional)
    clear_at = 56 # int | When the message should be cleared (optional)

    try:
        # Set the message to a custom message for the current user
        api_response = api_instance.user_status_user_status_set_custom_message(ocs_api_request, status_icon=status_icon, message=message, clear_at=clear_at)
        print("The response of UserStatusUserStatusApi->user_status_user_status_set_custom_message:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling UserStatusUserStatusApi->user_status_user_status_set_custom_message: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]
 **status_icon** | **str**| Icon of the status | [optional] 
 **message** | **str**| Message of the status | [optional] 
 **clear_at** | **int**| When the message should be cleared | [optional] 

### Return type

[**UserStatusUserStatusGetStatus200Response**](UserStatusUserStatusGetStatus200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | The message was updated successfully |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **user_status_user_status_set_predefined_message**
> UserStatusUserStatusGetStatus200Response user_status_user_status_set_predefined_message(message_id, ocs_api_request, clear_at=clear_at)

Set the message to a predefined message for the current user

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.user_status_user_status_get_status200_response import UserStatusUserStatusGetStatus200Response
from nextcloud_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = nextcloud_client.Configuration(
    host = "http://localhost"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure HTTP basic authorization: basic_auth
configuration = nextcloud_client.Configuration(
    username = os.environ["USERNAME"],
    password = os.environ["PASSWORD"]
)

# Configure Bearer authorization: bearer_auth
configuration = nextcloud_client.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with nextcloud_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = nextcloud_client.UserStatusUserStatusApi(api_client)
    message_id = 'message_id_example' # str | ID of the predefined message
    ocs_api_request = 'true' # str |  (default to 'true')
    clear_at = 56 # int | When the message should be cleared (optional)

    try:
        # Set the message to a predefined message for the current user
        api_response = api_instance.user_status_user_status_set_predefined_message(message_id, ocs_api_request, clear_at=clear_at)
        print("The response of UserStatusUserStatusApi->user_status_user_status_set_predefined_message:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling UserStatusUserStatusApi->user_status_user_status_set_predefined_message: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **message_id** | **str**| ID of the predefined message | 
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]
 **clear_at** | **int**| When the message should be cleared | [optional] 

### Return type

[**UserStatusUserStatusGetStatus200Response**](UserStatusUserStatusGetStatus200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | The message was updated successfully |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **user_status_user_status_set_status**
> UserStatusUserStatusGetStatus200Response user_status_user_status_set_status(status_type, ocs_api_request)

Update the status type of the current user

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.user_status_user_status_get_status200_response import UserStatusUserStatusGetStatus200Response
from nextcloud_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = nextcloud_client.Configuration(
    host = "http://localhost"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure HTTP basic authorization: basic_auth
configuration = nextcloud_client.Configuration(
    username = os.environ["USERNAME"],
    password = os.environ["PASSWORD"]
)

# Configure Bearer authorization: bearer_auth
configuration = nextcloud_client.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with nextcloud_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = nextcloud_client.UserStatusUserStatusApi(api_client)
    status_type = 'status_type_example' # str | The new status type
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Update the status type of the current user
        api_response = api_instance.user_status_user_status_set_status(status_type, ocs_api_request)
        print("The response of UserStatusUserStatusApi->user_status_user_status_set_status:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling UserStatusUserStatusApi->user_status_user_status_set_status: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status_type** | **str**| The new status type | 
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]

### Return type

[**UserStatusUserStatusGetStatus200Response**](UserStatusUserStatusGetStatus200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | The status was updated successfully |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

