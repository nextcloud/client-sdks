# nextcloud_client.CoreNavigationApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**core_navigation_get_apps_navigation**](CoreNavigationApi.md#core_navigation_get_apps_navigation) | **GET** /ocs/v2.php/core/navigation/apps | Get the apps navigation
[**core_navigation_get_settings_navigation**](CoreNavigationApi.md#core_navigation_get_settings_navigation) | **GET** /ocs/v2.php/core/navigation/settings | Get the settings navigation


# **core_navigation_get_apps_navigation**
> CoreNavigationGetAppsNavigation200Response core_navigation_get_apps_navigation(ocs_api_request, absolute=absolute)

Get the apps navigation

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.core_navigation_get_apps_navigation200_response import CoreNavigationGetAppsNavigation200Response
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
    api_instance = nextcloud_client.CoreNavigationApi(api_client)
    ocs_api_request = 'true' # str |  (default to 'true')
    absolute = 0 # int | Rewrite URLs to absolute ones (optional) (default to 0)

    try:
        # Get the apps navigation
        api_response = api_instance.core_navigation_get_apps_navigation(ocs_api_request, absolute=absolute)
        print("The response of CoreNavigationApi->core_navigation_get_apps_navigation:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CoreNavigationApi->core_navigation_get_apps_navigation: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]
 **absolute** | **int**| Rewrite URLs to absolute ones | [optional] [default to 0]

### Return type

[**CoreNavigationGetAppsNavigation200Response**](CoreNavigationGetAppsNavigation200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Apps navigation returned |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **core_navigation_get_settings_navigation**
> CoreNavigationGetAppsNavigation200Response core_navigation_get_settings_navigation(ocs_api_request, absolute=absolute)

Get the settings navigation

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.core_navigation_get_apps_navigation200_response import CoreNavigationGetAppsNavigation200Response
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
    api_instance = nextcloud_client.CoreNavigationApi(api_client)
    ocs_api_request = 'true' # str |  (default to 'true')
    absolute = 0 # int | Rewrite URLs to absolute ones (optional) (default to 0)

    try:
        # Get the settings navigation
        api_response = api_instance.core_navigation_get_settings_navigation(ocs_api_request, absolute=absolute)
        print("The response of CoreNavigationApi->core_navigation_get_settings_navigation:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CoreNavigationApi->core_navigation_get_settings_navigation: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]
 **absolute** | **int**| Rewrite URLs to absolute ones | [optional] [default to 0]

### Return type

[**CoreNavigationGetAppsNavigation200Response**](CoreNavigationGetAppsNavigation200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Apps navigation returned |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

