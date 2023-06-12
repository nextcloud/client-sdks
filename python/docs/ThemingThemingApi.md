# openapi_client.ThemingThemingApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**theming_theming_get_image**](ThemingThemingApi.md#theming_theming_get_image) | **GET** /index.php/apps/theming/image/{key} | Get an image
[**theming_theming_get_manifest**](ThemingThemingApi.md#theming_theming_get_manifest) | **GET** /index.php/apps/theming/manifest/{app} | Get the manifest for an app
[**theming_theming_get_theme_stylesheet**](ThemingThemingApi.md#theming_theming_get_theme_stylesheet) | **GET** /index.php/apps/theming/theme/{themeId}.css | Get the CSS stylesheet for a theme


# **theming_theming_get_image**
> bytearray theming_theming_get_image(key, use_svg=use_svg)

Get an image

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
    api_instance = openapi_client.ThemingThemingApi(api_client)
    key = 'key_example' # str | Key of the image
    use_svg = 1 # int | Return image as SVG (optional) (default to 1)

    try:
        # Get an image
        api_response = api_instance.theming_theming_get_image(key, use_svg=use_svg)
        print("The response of ThemingThemingApi->theming_theming_get_image:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ThemingThemingApi->theming_theming_get_image: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **key** | **str**| Key of the image | 
 **use_svg** | **int**| Return image as SVG | [optional] [default to 1]

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
**200** | Image returned |  * Content-Disposition -  <br>  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **theming_theming_get_manifest**
> ThemingThemingGetManifest200Response theming_theming_get_manifest(app)

Get the manifest for an app

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import openapi_client
from openapi_client.models.theming_theming_get_manifest200_response import ThemingThemingGetManifest200Response
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
    api_instance = openapi_client.ThemingThemingApi(api_client)
    app = 'app_example' # str | ID of the app

    try:
        # Get the manifest for an app
        api_response = api_instance.theming_theming_get_manifest(app)
        print("The response of ThemingThemingApi->theming_theming_get_manifest:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ThemingThemingApi->theming_theming_get_manifest: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app** | **str**| ID of the app | 

### Return type

[**ThemingThemingGetManifest200Response**](ThemingThemingGetManifest200Response.md)

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

# **theming_theming_get_theme_stylesheet**
> bytearray theming_theming_get_theme_stylesheet(theme_id, plain=plain, with_custom_css=with_custom_css)

Get the CSS stylesheet for a theme

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
    api_instance = openapi_client.ThemingThemingApi(api_client)
    theme_id = 'theme_id_example' # str | ID of the theme
    plain = 0 # int | Let the browser decide the CSS priority (optional) (default to 0)
    with_custom_css = 0 # int | Include custom CSS (optional) (default to 0)

    try:
        # Get the CSS stylesheet for a theme
        api_response = api_instance.theming_theming_get_theme_stylesheet(theme_id, plain=plain, with_custom_css=with_custom_css)
        print("The response of ThemingThemingApi->theming_theming_get_theme_stylesheet:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ThemingThemingApi->theming_theming_get_theme_stylesheet: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **theme_id** | **str**| ID of the theme | 
 **plain** | **int**| Let the browser decide the CSS priority | [optional] [default to 0]
 **with_custom_css** | **int**| Include custom CSS | [optional] [default to 0]

### Return type

**bytearray**

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/css

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Stylesheet returned |  * Content-Disposition -  <br>  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

