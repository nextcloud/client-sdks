# openapi_client.UserStatusPredefinedStatusApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**user_status_predefined_status_find_all**](UserStatusPredefinedStatusApi.md#user_status_predefined_status_find_all) | **GET** /ocs/v2.php/apps/user_status/api/v1/predefined_statuses | Get all predefined messages


# **user_status_predefined_status_find_all**
> UserStatusPredefinedStatusFindAll200Response user_status_predefined_status_find_all(ocs_api_request)

Get all predefined messages

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import openapi_client
from openapi_client.models.user_status_predefined_status_find_all200_response import UserStatusPredefinedStatusFindAll200Response
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
    api_instance = openapi_client.UserStatusPredefinedStatusApi(api_client)
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Get all predefined messages
        api_response = api_instance.user_status_predefined_status_find_all(ocs_api_request)
        print("The response of UserStatusPredefinedStatusApi->user_status_predefined_status_find_all:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling UserStatusPredefinedStatusApi->user_status_predefined_status_find_all: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]

### Return type

[**UserStatusPredefinedStatusFindAll200Response**](UserStatusPredefinedStatusFindAll200Response.md)

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

