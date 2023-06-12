# openapi_client.UserStatusStatusesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**user_status_statuses_find**](UserStatusStatusesApi.md#user_status_statuses_find) | **GET** /ocs/v2.php/apps/user_status/api/v1/statuses/{userId} | Find the status of a user
[**user_status_statuses_find_all**](UserStatusStatusesApi.md#user_status_statuses_find_all) | **GET** /ocs/v2.php/apps/user_status/api/v1/statuses | Find statuses of users


# **user_status_statuses_find**
> UserStatusStatusesFind200Response user_status_statuses_find(user_id, ocs_api_request)

Find the status of a user

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import openapi_client
from openapi_client.models.user_status_statuses_find200_response import UserStatusStatusesFind200Response
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
    api_instance = openapi_client.UserStatusStatusesApi(api_client)
    user_id = 'user_id_example' # str | ID of the user
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Find the status of a user
        api_response = api_instance.user_status_statuses_find(user_id, ocs_api_request)
        print("The response of UserStatusStatusesApi->user_status_statuses_find:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling UserStatusStatusesApi->user_status_statuses_find: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user_id** | **str**| ID of the user | 
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]

### Return type

[**UserStatusStatusesFind200Response**](UserStatusStatusesFind200Response.md)

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

# **user_status_statuses_find_all**
> UserStatusStatusesFindAll200Response user_status_statuses_find_all(ocs_api_request, limit=limit, offset=offset)

Find statuses of users

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import openapi_client
from openapi_client.models.user_status_statuses_find_all200_response import UserStatusStatusesFindAll200Response
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
    api_instance = openapi_client.UserStatusStatusesApi(api_client)
    ocs_api_request = 'true' # str |  (default to 'true')
    limit = 56 # int | Maximum number of statuses to find (optional)
    offset = 56 # int | Offset for finding statuses (optional)

    try:
        # Find statuses of users
        api_response = api_instance.user_status_statuses_find_all(ocs_api_request, limit=limit, offset=offset)
        print("The response of UserStatusStatusesApi->user_status_statuses_find_all:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling UserStatusStatusesApi->user_status_statuses_find_all: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]
 **limit** | **int**| Maximum number of statuses to find | [optional] 
 **offset** | **int**| Offset for finding statuses | [optional] 

### Return type

[**UserStatusStatusesFindAll200Response**](UserStatusStatusesFindAll200Response.md)

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

