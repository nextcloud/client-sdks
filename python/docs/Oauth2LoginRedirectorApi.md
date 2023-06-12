# openapi_client.Oauth2LoginRedirectorApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**oauth2_login_redirector_authorize**](Oauth2LoginRedirectorApi.md#oauth2_login_redirector_authorize) | **GET** /index.php/apps/oauth2/authorize | Authorize the user


# **oauth2_login_redirector_authorize**
> str oauth2_login_redirector_authorize(client_id, state, response_type)

Authorize the user

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
    api_instance = openapi_client.Oauth2LoginRedirectorApi(api_client)
    client_id = 'client_id_example' # str | Client ID
    state = 'state_example' # str | State of the flow
    response_type = 'response_type_example' # str | Response type for the flow

    try:
        # Authorize the user
        api_response = api_instance.oauth2_login_redirector_authorize(client_id, state, response_type)
        print("The response of Oauth2LoginRedirectorApi->oauth2_login_redirector_authorize:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling Oauth2LoginRedirectorApi->oauth2_login_redirector_authorize: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **client_id** | **str**| Client ID | 
 **state** | **str**| State of the flow | 
 **response_type** | **str**| Response type for the flow | 

### Return type

**str**

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/html

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Client not found |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

