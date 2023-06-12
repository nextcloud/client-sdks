# openapi_client.CoreHoverCardApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**core_hover_card_get_user**](CoreHoverCardApi.md#core_hover_card_get_user) | **GET** /ocs/v2.php/hovercard/v1/{userId} | Get the user details for a hovercard


# **core_hover_card_get_user**
> CoreHoverCardGetUser200Response core_hover_card_get_user(user_id, ocs_api_request)

Get the user details for a hovercard

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import openapi_client
from openapi_client.models.core_hover_card_get_user200_response import CoreHoverCardGetUser200Response
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
    api_instance = openapi_client.CoreHoverCardApi(api_client)
    user_id = 'user_id_example' # str | ID of the user
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Get the user details for a hovercard
        api_response = api_instance.core_hover_card_get_user(user_id, ocs_api_request)
        print("The response of CoreHoverCardApi->core_hover_card_get_user:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CoreHoverCardApi->core_hover_card_get_user: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user_id** | **str**| ID of the user | 
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]

### Return type

[**CoreHoverCardGetUser200Response**](CoreHoverCardGetUser200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | User details returned |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

