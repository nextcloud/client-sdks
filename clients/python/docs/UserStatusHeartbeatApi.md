# nextcloud_client.UserStatusHeartbeatApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**user_status_heartbeat_heartbeat**](UserStatusHeartbeatApi.md#user_status_heartbeat_heartbeat) | **PUT** /ocs/v2.php/apps/user_status/api/v1/heartbeat | Keep the status alive


# **user_status_heartbeat_heartbeat**
> UserStatusUserStatusGetStatus200Response user_status_heartbeat_heartbeat(status, ocs_api_request)

Keep the status alive

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
    api_instance = nextcloud_client.UserStatusHeartbeatApi(api_client)
    status = 'status_example' # str | Only online, away
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Keep the status alive
        api_response = api_instance.user_status_heartbeat_heartbeat(status, ocs_api_request)
        print("The response of UserStatusHeartbeatApi->user_status_heartbeat_heartbeat:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling UserStatusHeartbeatApi->user_status_heartbeat_heartbeat: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **str**| Only online, away | 
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
**200** | Status successfully updated |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

