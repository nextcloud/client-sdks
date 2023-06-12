# openapi_client.CoreGuestAvatarApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**core_guest_avatar_get_avatar**](CoreGuestAvatarApi.md#core_guest_avatar_get_avatar) | **GET** /index.php/avatar/guest/{guestName}/{size} | Returns a guest avatar image response
[**core_guest_avatar_get_avatar_dark**](CoreGuestAvatarApi.md#core_guest_avatar_get_avatar_dark) | **GET** /index.php/avatar/guest/{guestName}/{size}/dark | Returns a dark guest avatar image response


# **core_guest_avatar_get_avatar**
> bytearray core_guest_avatar_get_avatar(guest_name, size, dark_theme=dark_theme)

Returns a guest avatar image response

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
    api_instance = openapi_client.CoreGuestAvatarApi(api_client)
    guest_name = 'guest_name_example' # str | The guest name, e.g. \"Albert\"
    size = 'size_example' # str | The desired avatar size, e.g. 64 for 64x64px
    dark_theme = 0 # int | Return dark avatar (optional) (default to 0)

    try:
        # Returns a guest avatar image response
        api_response = api_instance.core_guest_avatar_get_avatar(guest_name, size, dark_theme=dark_theme)
        print("The response of CoreGuestAvatarApi->core_guest_avatar_get_avatar:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CoreGuestAvatarApi->core_guest_avatar_get_avatar: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **guest_name** | **str**| The guest name, e.g. \&quot;Albert\&quot; | 
 **size** | **str**| The desired avatar size, e.g. 64 for 64x64px | 
 **dark_theme** | **int**| Return dark avatar | [optional] [default to 0]

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
**200** | Custom avatar returned |  * Content-Disposition -  <br>  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **core_guest_avatar_get_avatar_dark**
> bytearray core_guest_avatar_get_avatar_dark(guest_name, size)

Returns a dark guest avatar image response

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
    api_instance = openapi_client.CoreGuestAvatarApi(api_client)
    guest_name = 'guest_name_example' # str | The guest name, e.g. \"Albert\"
    size = 'size_example' # str | The desired avatar size, e.g. 64 for 64x64px

    try:
        # Returns a dark guest avatar image response
        api_response = api_instance.core_guest_avatar_get_avatar_dark(guest_name, size)
        print("The response of CoreGuestAvatarApi->core_guest_avatar_get_avatar_dark:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CoreGuestAvatarApi->core_guest_avatar_get_avatar_dark: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **guest_name** | **str**| The guest name, e.g. \&quot;Albert\&quot; | 
 **size** | **str**| The desired avatar size, e.g. 64 for 64x64px | 

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
**200** | Custom avatar returned |  * Content-Disposition -  <br>  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

