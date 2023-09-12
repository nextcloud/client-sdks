# nextcloud_client.UpdatenotificationApiApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**updatenotification_api_get_app_list**](UpdatenotificationApiApi.md#updatenotification_api_get_app_list) | **GET** /ocs/v2.php/apps/updatenotification/api/{apiVersion}/applist/{newVersion} | List available updates for apps


# **updatenotification_api_get_app_list**
> UpdatenotificationApiGetAppList200Response updatenotification_api_get_app_list(api_version, new_version, ocs_api_request)

List available updates for apps

This endpoint requires admin access

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.updatenotification_api_get_app_list200_response import UpdatenotificationApiGetAppList200Response
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
    api_instance = nextcloud_client.UpdatenotificationApiApi(api_client)
    api_version = 'v1' # str |  (default to 'v1')
    new_version = 'new_version_example' # str | Server version to check updates for
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # List available updates for apps
        api_response = api_instance.updatenotification_api_get_app_list(api_version, new_version, ocs_api_request)
        print("The response of UpdatenotificationApiApi->updatenotification_api_get_app_list:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling UpdatenotificationApiApi->updatenotification_api_get_app_list: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **api_version** | **str**|  | [default to &#39;v1&#39;]
 **new_version** | **str**| Server version to check updates for | 
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]

### Return type

[**UpdatenotificationApiGetAppList200Response**](UpdatenotificationApiGetAppList200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Apps returned |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

