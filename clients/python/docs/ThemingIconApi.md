# nextcloud_client.ThemingIconApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**theming_icon_get_favicon**](ThemingIconApi.md#theming_icon_get_favicon) | **GET** /index.php/apps/theming/favicon/{app} | Return a 32x32 favicon as png
[**theming_icon_get_themed_icon**](ThemingIconApi.md#theming_icon_get_themed_icon) | **GET** /index.php/apps/theming/img/{app}/{image} | Get a themed icon
[**theming_icon_get_touch_icon**](ThemingIconApi.md#theming_icon_get_touch_icon) | **GET** /index.php/apps/theming/icon/{app} | Return a 512x512 icon for touch devices


# **theming_icon_get_favicon**
> bytearray theming_icon_get_favicon(app)

Return a 32x32 favicon as png

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
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
    api_instance = nextcloud_client.ThemingIconApi(api_client)
    app = 'core' # str | ID of the app (default to 'core')

    try:
        # Return a 32x32 favicon as png
        api_response = api_instance.theming_icon_get_favicon(app)
        print("The response of ThemingIconApi->theming_icon_get_favicon:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ThemingIconApi->theming_icon_get_favicon: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app** | **str**| ID of the app | [default to &#39;core&#39;]

### Return type

**bytearray**

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: image/x-icon

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Favicon returned |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **theming_icon_get_themed_icon**
> bytearray theming_icon_get_themed_icon(app, image)

Get a themed icon

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
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
    api_instance = nextcloud_client.ThemingIconApi(api_client)
    app = 'app_example' # str | ID of the app
    image = 'image_example' # str | image file name (svg required)

    try:
        # Get a themed icon
        api_response = api_instance.theming_icon_get_themed_icon(app, image)
        print("The response of ThemingIconApi->theming_icon_get_themed_icon:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ThemingIconApi->theming_icon_get_themed_icon: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app** | **str**| ID of the app | 
 **image** | **str**| image file name (svg required) | 

### Return type

**bytearray**

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: image/svg+xml

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Themed icon returned |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **theming_icon_get_touch_icon**
> bytearray theming_icon_get_touch_icon(app)

Return a 512x512 icon for touch devices

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
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
    api_instance = nextcloud_client.ThemingIconApi(api_client)
    app = 'core' # str | ID of the app (default to 'core')

    try:
        # Return a 512x512 icon for touch devices
        api_response = api_instance.theming_icon_get_touch_icon(app)
        print("The response of ThemingIconApi->theming_icon_get_touch_icon:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ThemingIconApi->theming_icon_get_touch_icon: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app** | **str**| ID of the app | [default to &#39;core&#39;]

### Return type

**bytearray**

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: image/png, image/x-icon

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Touch icon returned |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

