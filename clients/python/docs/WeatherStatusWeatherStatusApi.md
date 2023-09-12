# nextcloud_client.WeatherStatusWeatherStatusApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**weather_status_weather_status_get_favorites**](WeatherStatusWeatherStatusApi.md#weather_status_weather_status_get_favorites) | **GET** /ocs/v2.php/apps/weather_status/api/v1/favorites | Get favorites list
[**weather_status_weather_status_get_forecast**](WeatherStatusWeatherStatusApi.md#weather_status_weather_status_get_forecast) | **GET** /ocs/v2.php/apps/weather_status/api/v1/forecast | Get forecast for current location
[**weather_status_weather_status_get_location**](WeatherStatusWeatherStatusApi.md#weather_status_weather_status_get_location) | **GET** /ocs/v2.php/apps/weather_status/api/v1/location | Get stored user location
[**weather_status_weather_status_set_favorites**](WeatherStatusWeatherStatusApi.md#weather_status_weather_status_set_favorites) | **PUT** /ocs/v2.php/apps/weather_status/api/v1/favorites | Set favorites list
[**weather_status_weather_status_set_location**](WeatherStatusWeatherStatusApi.md#weather_status_weather_status_set_location) | **PUT** /ocs/v2.php/apps/weather_status/api/v1/location | Set address and resolve it to get coordinates or directly set coordinates and get address with reverse geocoding
[**weather_status_weather_status_set_mode**](WeatherStatusWeatherStatusApi.md#weather_status_weather_status_set_mode) | **PUT** /ocs/v2.php/apps/weather_status/api/v1/mode | Change the weather status mode. There are currently 2 modes: - ask the browser - use the user defined address
[**weather_status_weather_status_use_personal_address**](WeatherStatusWeatherStatusApi.md#weather_status_weather_status_use_personal_address) | **PUT** /ocs/v2.php/apps/weather_status/api/v1/use-personal | Try to use the address set in user personal settings as weather location


# **weather_status_weather_status_get_favorites**
> ProvisioningApiGroupsGetSubAdminsOfGroup200Response weather_status_weather_status_get_favorites(ocs_api_request)

Get favorites list

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.provisioning_api_groups_get_sub_admins_of_group200_response import ProvisioningApiGroupsGetSubAdminsOfGroup200Response
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
    api_instance = nextcloud_client.WeatherStatusWeatherStatusApi(api_client)
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Get favorites list
        api_response = api_instance.weather_status_weather_status_get_favorites(ocs_api_request)
        print("The response of WeatherStatusWeatherStatusApi->weather_status_weather_status_get_favorites:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling WeatherStatusWeatherStatusApi->weather_status_weather_status_get_favorites: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]

### Return type

[**ProvisioningApiGroupsGetSubAdminsOfGroup200Response**](ProvisioningApiGroupsGetSubAdminsOfGroup200Response.md)

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

# **weather_status_weather_status_get_forecast**
> WeatherStatusWeatherStatusGetForecast200Response weather_status_weather_status_get_forecast(ocs_api_request)

Get forecast for current location

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.weather_status_weather_status_get_forecast200_response import WeatherStatusWeatherStatusGetForecast200Response
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
    api_instance = nextcloud_client.WeatherStatusWeatherStatusApi(api_client)
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Get forecast for current location
        api_response = api_instance.weather_status_weather_status_get_forecast(ocs_api_request)
        print("The response of WeatherStatusWeatherStatusApi->weather_status_weather_status_get_forecast:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling WeatherStatusWeatherStatusApi->weather_status_weather_status_get_forecast: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]

### Return type

[**WeatherStatusWeatherStatusGetForecast200Response**](WeatherStatusWeatherStatusGetForecast200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Forecast returned |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **weather_status_weather_status_get_location**
> WeatherStatusWeatherStatusGetLocation200Response weather_status_weather_status_get_location(ocs_api_request)

Get stored user location

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.weather_status_weather_status_get_location200_response import WeatherStatusWeatherStatusGetLocation200Response
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
    api_instance = nextcloud_client.WeatherStatusWeatherStatusApi(api_client)
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Get stored user location
        api_response = api_instance.weather_status_weather_status_get_location(ocs_api_request)
        print("The response of WeatherStatusWeatherStatusApi->weather_status_weather_status_get_location:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling WeatherStatusWeatherStatusApi->weather_status_weather_status_get_location: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]

### Return type

[**WeatherStatusWeatherStatusGetLocation200Response**](WeatherStatusWeatherStatusGetLocation200Response.md)

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

# **weather_status_weather_status_set_favorites**
> CoreReferenceApiTouchProvider200Response weather_status_weather_status_set_favorites(favorites, ocs_api_request)

Set favorites list

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.core_reference_api_touch_provider200_response import CoreReferenceApiTouchProvider200Response
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
    api_instance = nextcloud_client.WeatherStatusWeatherStatusApi(api_client)
    favorites = ['favorites_example'] # List[str] | Favorite addresses
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Set favorites list
        api_response = api_instance.weather_status_weather_status_set_favorites(favorites, ocs_api_request)
        print("The response of WeatherStatusWeatherStatusApi->weather_status_weather_status_set_favorites:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling WeatherStatusWeatherStatusApi->weather_status_weather_status_set_favorites: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **favorites** | [**List[str]**](str.md)| Favorite addresses | 
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]

### Return type

[**CoreReferenceApiTouchProvider200Response**](CoreReferenceApiTouchProvider200Response.md)

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

# **weather_status_weather_status_set_location**
> WeatherStatusWeatherStatusUsePersonalAddress200Response weather_status_weather_status_set_location(ocs_api_request, address=address, lat=lat, lon=lon)

Set address and resolve it to get coordinates or directly set coordinates and get address with reverse geocoding

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.weather_status_weather_status_use_personal_address200_response import WeatherStatusWeatherStatusUsePersonalAddress200Response
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
    api_instance = nextcloud_client.WeatherStatusWeatherStatusApi(api_client)
    ocs_api_request = 'true' # str |  (default to 'true')
    address = 'address_example' # str | Any approximative or exact address (optional)
    lat = 3.4 # float | Latitude in decimal degree format (optional)
    lon = 3.4 # float | Longitude in decimal degree format (optional)

    try:
        # Set address and resolve it to get coordinates or directly set coordinates and get address with reverse geocoding
        api_response = api_instance.weather_status_weather_status_set_location(ocs_api_request, address=address, lat=lat, lon=lon)
        print("The response of WeatherStatusWeatherStatusApi->weather_status_weather_status_set_location:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling WeatherStatusWeatherStatusApi->weather_status_weather_status_set_location: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]
 **address** | **str**| Any approximative or exact address | [optional] 
 **lat** | **float**| Latitude in decimal degree format | [optional] 
 **lon** | **float**| Longitude in decimal degree format | [optional] 

### Return type

[**WeatherStatusWeatherStatusUsePersonalAddress200Response**](WeatherStatusWeatherStatusUsePersonalAddress200Response.md)

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

# **weather_status_weather_status_set_mode**
> CoreReferenceApiTouchProvider200Response weather_status_weather_status_set_mode(mode, ocs_api_request)

Change the weather status mode. There are currently 2 modes: - ask the browser - use the user defined address

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.core_reference_api_touch_provider200_response import CoreReferenceApiTouchProvider200Response
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
    api_instance = nextcloud_client.WeatherStatusWeatherStatusApi(api_client)
    mode = 56 # int | New mode
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Change the weather status mode. There are currently 2 modes: - ask the browser - use the user defined address
        api_response = api_instance.weather_status_weather_status_set_mode(mode, ocs_api_request)
        print("The response of WeatherStatusWeatherStatusApi->weather_status_weather_status_set_mode:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling WeatherStatusWeatherStatusApi->weather_status_weather_status_set_mode: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mode** | **int**| New mode | 
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]

### Return type

[**CoreReferenceApiTouchProvider200Response**](CoreReferenceApiTouchProvider200Response.md)

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

# **weather_status_weather_status_use_personal_address**
> WeatherStatusWeatherStatusUsePersonalAddress200Response weather_status_weather_status_use_personal_address(ocs_api_request)

Try to use the address set in user personal settings as weather location

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.weather_status_weather_status_use_personal_address200_response import WeatherStatusWeatherStatusUsePersonalAddress200Response
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
    api_instance = nextcloud_client.WeatherStatusWeatherStatusApi(api_client)
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Try to use the address set in user personal settings as weather location
        api_response = api_instance.weather_status_weather_status_use_personal_address(ocs_api_request)
        print("The response of WeatherStatusWeatherStatusApi->weather_status_weather_status_use_personal_address:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling WeatherStatusWeatherStatusApi->weather_status_weather_status_use_personal_address: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]

### Return type

[**WeatherStatusWeatherStatusUsePersonalAddress200Response**](WeatherStatusWeatherStatusUsePersonalAddress200Response.md)

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

