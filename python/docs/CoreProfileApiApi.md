# openapi_client.CoreProfileApiApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**core_profile_api_set_visibility**](CoreProfileApiApi.md#core_profile_api_set_visibility) | **PUT** /ocs/v2.php/profile/{targetUserId} | Update the visiblity of a parameter


# **core_profile_api_set_visibility**
> CoreWhatsNewDismiss200Response core_profile_api_set_visibility(param_id, visibility, target_user_id, ocs_api_request)

Update the visiblity of a parameter

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import openapi_client
from openapi_client.models.core_whats_new_dismiss200_response import CoreWhatsNewDismiss200Response
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
    api_instance = openapi_client.CoreProfileApiApi(api_client)
    param_id = 'param_id_example' # str | ID of the parameter
    visibility = 'visibility_example' # str | New visibility
    target_user_id = 'target_user_id_example' # str | ID of the user
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Update the visiblity of a parameter
        api_response = api_instance.core_profile_api_set_visibility(param_id, visibility, target_user_id, ocs_api_request)
        print("The response of CoreProfileApiApi->core_profile_api_set_visibility:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CoreProfileApiApi->core_profile_api_set_visibility: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **param_id** | **str**| ID of the parameter | 
 **visibility** | **str**| New visibility | 
 **target_user_id** | **str**| ID of the user | 
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
**200** | Visibility updated successfully |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

