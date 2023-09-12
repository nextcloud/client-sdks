# nextcloud_client.CoreClientFlowLoginV2Api

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**core_client_flow_login_v2_init**](CoreClientFlowLoginV2Api.md#core_client_flow_login_v2_init) | **POST** /index.php/login/v2 | Init a login flow
[**core_client_flow_login_v2_poll**](CoreClientFlowLoginV2Api.md#core_client_flow_login_v2_poll) | **POST** /index.php/login/v2/poll | Poll the login flow credentials


# **core_client_flow_login_v2_init**
> CoreLoginFlowV2 core_client_flow_login_v2_init()

Init a login flow

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.core_login_flow_v2 import CoreLoginFlowV2
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
    api_instance = nextcloud_client.CoreClientFlowLoginV2Api(api_client)

    try:
        # Init a login flow
        api_response = api_instance.core_client_flow_login_v2_init()
        print("The response of CoreClientFlowLoginV2Api->core_client_flow_login_v2_init:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CoreClientFlowLoginV2Api->core_client_flow_login_v2_init: %s\n" % e)
```



### Parameters
This endpoint does not need any parameter.

### Return type

[**CoreLoginFlowV2**](CoreLoginFlowV2.md)

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

# **core_client_flow_login_v2_poll**
> CoreLoginFlowV2Credentials core_client_flow_login_v2_poll(token)

Poll the login flow credentials

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.core_login_flow_v2_credentials import CoreLoginFlowV2Credentials
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
    api_instance = nextcloud_client.CoreClientFlowLoginV2Api(api_client)
    token = 'token_example' # str | Token of the flow

    try:
        # Poll the login flow credentials
        api_response = api_instance.core_client_flow_login_v2_poll(token)
        print("The response of CoreClientFlowLoginV2Api->core_client_flow_login_v2_poll:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CoreClientFlowLoginV2Api->core_client_flow_login_v2_poll: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **str**| Token of the flow | 

### Return type

[**CoreLoginFlowV2Credentials**](CoreLoginFlowV2Credentials.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Login flow credentials returned |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

