# openapi_client.ThemingUserThemeApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**theming_user_theme_delete_background**](ThemingUserThemeApi.md#theming_user_theme_delete_background) | **DELETE** /index.php/apps/theming/background/custom | Delete the background
[**theming_user_theme_disable_theme**](ThemingUserThemeApi.md#theming_user_theme_disable_theme) | **DELETE** /ocs/v2.php/apps/theming/api/v1/theme/{themeId} | Disable theme
[**theming_user_theme_enable_theme**](ThemingUserThemeApi.md#theming_user_theme_enable_theme) | **PUT** /ocs/v2.php/apps/theming/api/v1/theme/{themeId}/enable | Enable theme
[**theming_user_theme_get_background**](ThemingUserThemeApi.md#theming_user_theme_get_background) | **GET** /index.php/apps/theming/background | Get the background image
[**theming_user_theme_set_background**](ThemingUserThemeApi.md#theming_user_theme_set_background) | **POST** /index.php/apps/theming/background/{type} | Set the background


# **theming_user_theme_delete_background**
> ThemingBackground theming_user_theme_delete_background(ocs_api_request)

Delete the background

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import openapi_client
from openapi_client.models.theming_background import ThemingBackground
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
    api_instance = openapi_client.ThemingUserThemeApi(api_client)
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Delete the background
        api_response = api_instance.theming_user_theme_delete_background(ocs_api_request)
        print("The response of ThemingUserThemeApi->theming_user_theme_delete_background:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ThemingUserThemeApi->theming_user_theme_delete_background: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]

### Return type

[**ThemingBackground**](ThemingBackground.md)

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

# **theming_user_theme_disable_theme**
> CoreWhatsNewDismiss200Response theming_user_theme_disable_theme(theme_id, ocs_api_request)

Disable theme

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
    api_instance = openapi_client.ThemingUserThemeApi(api_client)
    theme_id = 'theme_id_example' # str | the theme ID
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Disable theme
        api_response = api_instance.theming_user_theme_disable_theme(theme_id, ocs_api_request)
        print("The response of ThemingUserThemeApi->theming_user_theme_disable_theme:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ThemingUserThemeApi->theming_user_theme_disable_theme: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **theme_id** | **str**| the theme ID | 
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
**200** | Theme disabled successfully |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **theming_user_theme_enable_theme**
> CoreWhatsNewDismiss200Response theming_user_theme_enable_theme(theme_id, ocs_api_request)

Enable theme

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
    api_instance = openapi_client.ThemingUserThemeApi(api_client)
    theme_id = 'theme_id_example' # str | the theme ID
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Enable theme
        api_response = api_instance.theming_user_theme_enable_theme(theme_id, ocs_api_request)
        print("The response of ThemingUserThemeApi->theming_user_theme_enable_theme:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ThemingUserThemeApi->theming_user_theme_enable_theme: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **theme_id** | **str**| the theme ID | 
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
**200** | Theme enabled successfully |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **theming_user_theme_get_background**
> bytearray theming_user_theme_get_background(ocs_api_request)

Get the background image

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
    api_instance = openapi_client.ThemingUserThemeApi(api_client)
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Get the background image
        api_response = api_instance.theming_user_theme_get_background(ocs_api_request)
        print("The response of ThemingUserThemeApi->theming_user_theme_get_background:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ThemingUserThemeApi->theming_user_theme_get_background: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]

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
**200** | Background image returned |  * Content-Disposition -  <br>  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **theming_user_theme_set_background**
> ThemingBackground theming_user_theme_set_background(type, ocs_api_request, value=value, color=color)

Set the background

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import openapi_client
from openapi_client.models.theming_background import ThemingBackground
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
    api_instance = openapi_client.ThemingUserThemeApi(api_client)
    type = 'type_example' # str | Type of background
    ocs_api_request = 'true' # str |  (default to 'true')
    value = '' # str | Path of the background image (optional) (default to '')
    color = 'color_example' # str | Color for the background (optional)

    try:
        # Set the background
        api_response = api_instance.theming_user_theme_set_background(type, ocs_api_request, value=value, color=color)
        print("The response of ThemingUserThemeApi->theming_user_theme_set_background:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ThemingUserThemeApi->theming_user_theme_set_background: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type** | **str**| Type of background | 
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]
 **value** | **str**| Path of the background image | [optional] [default to &#39;&#39;]
 **color** | **str**| Color for the background | [optional] 

### Return type

[**ThemingBackground**](ThemingBackground.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Background set successfully |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

