# openapi_client.CoreAvatarApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**core_avatar_get_avatar**](CoreAvatarApi.md#core_avatar_get_avatar) | **GET** /index.php/avatar/{userId}/{size} | Get the avatar
[**core_avatar_get_avatar_dark**](CoreAvatarApi.md#core_avatar_get_avatar_dark) | **GET** /index.php/avatar/{userId}/{size}/dark | Get the dark avatar


# **core_avatar_get_avatar**
> bytearray core_avatar_get_avatar(user_id, size)

Get the avatar

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
    api_instance = openapi_client.CoreAvatarApi(api_client)
    user_id = 'user_id_example' # str | ID of the user
    size = 56 # int | Size of the avatar

    try:
        # Get the avatar
        api_response = api_instance.core_avatar_get_avatar(user_id, size)
        print("The response of CoreAvatarApi->core_avatar_get_avatar:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CoreAvatarApi->core_avatar_get_avatar: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user_id** | **str**| ID of the user | 
 **size** | **int**| Size of the avatar | 

### Return type

**bytearray**

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Avatar returned |  * Content-Disposition -  <br>  * X-NC-IsCustomAvatar -  <br>  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **core_avatar_get_avatar_dark**
> bytearray core_avatar_get_avatar_dark(user_id, size)

Get the dark avatar

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
    api_instance = openapi_client.CoreAvatarApi(api_client)
    user_id = 'user_id_example' # str | ID of the user
    size = 56 # int | Size of the avatar

    try:
        # Get the dark avatar
        api_response = api_instance.core_avatar_get_avatar_dark(user_id, size)
        print("The response of CoreAvatarApi->core_avatar_get_avatar_dark:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CoreAvatarApi->core_avatar_get_avatar_dark: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user_id** | **str**| ID of the user | 
 **size** | **int**| Size of the avatar | 

### Return type

**bytearray**

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Avatar returned |  * Content-Disposition -  <br>  * X-NC-IsCustomAvatar -  <br>  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

