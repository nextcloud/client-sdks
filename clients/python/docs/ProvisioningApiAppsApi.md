# nextcloud_client.ProvisioningApiAppsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**provisioning_api_apps_disable**](ProvisioningApiAppsApi.md#provisioning_api_apps_disable) | **DELETE** /ocs/v2.php/cloud/apps/{app} | Disable an app
[**provisioning_api_apps_enable**](ProvisioningApiAppsApi.md#provisioning_api_apps_enable) | **POST** /ocs/v2.php/cloud/apps/{app} | Enable an app
[**provisioning_api_apps_get_app_info**](ProvisioningApiAppsApi.md#provisioning_api_apps_get_app_info) | **GET** /ocs/v2.php/cloud/apps/{app} | Get the app info for an app
[**provisioning_api_apps_get_apps**](ProvisioningApiAppsApi.md#provisioning_api_apps_get_apps) | **GET** /ocs/v2.php/cloud/apps | Get a list of installed apps


# **provisioning_api_apps_disable**
> CoreWhatsNewDismiss200Response provisioning_api_apps_disable(app, ocs_api_request)

Disable an app

This endpoint requires admin access

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.core_whats_new_dismiss200_response import CoreWhatsNewDismiss200Response
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
    api_instance = nextcloud_client.ProvisioningApiAppsApi(api_client)
    app = 'app_example' # str | ID of the app
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Disable an app
        api_response = api_instance.provisioning_api_apps_disable(app, ocs_api_request)
        print("The response of ProvisioningApiAppsApi->provisioning_api_apps_disable:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ProvisioningApiAppsApi->provisioning_api_apps_disable: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app** | **str**| ID of the app | 
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
**200** |  |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **provisioning_api_apps_enable**
> CoreWhatsNewDismiss200Response provisioning_api_apps_enable(app, ocs_api_request)

Enable an app

This endpoint requires admin access

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.core_whats_new_dismiss200_response import CoreWhatsNewDismiss200Response
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
    api_instance = nextcloud_client.ProvisioningApiAppsApi(api_client)
    app = 'app_example' # str | ID of the app
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Enable an app
        api_response = api_instance.provisioning_api_apps_enable(app, ocs_api_request)
        print("The response of ProvisioningApiAppsApi->provisioning_api_apps_enable:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ProvisioningApiAppsApi->provisioning_api_apps_enable: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app** | **str**| ID of the app | 
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
**200** |  |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **provisioning_api_apps_get_app_info**
> ProvisioningApiAppsGetAppInfo200Response provisioning_api_apps_get_app_info(app, ocs_api_request)

Get the app info for an app

This endpoint requires admin access

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.provisioning_api_apps_get_app_info200_response import ProvisioningApiAppsGetAppInfo200Response
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
    api_instance = nextcloud_client.ProvisioningApiAppsApi(api_client)
    app = 'app_example' # str | ID of the app
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Get the app info for an app
        api_response = api_instance.provisioning_api_apps_get_app_info(app, ocs_api_request)
        print("The response of ProvisioningApiAppsApi->provisioning_api_apps_get_app_info:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ProvisioningApiAppsApi->provisioning_api_apps_get_app_info: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app** | **str**| ID of the app | 
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]

### Return type

[**ProvisioningApiAppsGetAppInfo200Response**](ProvisioningApiAppsGetAppInfo200Response.md)

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

# **provisioning_api_apps_get_apps**
> ProvisioningApiAppsGetApps200Response provisioning_api_apps_get_apps(ocs_api_request, filter=filter)

Get a list of installed apps

This endpoint requires admin access

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.provisioning_api_apps_get_apps200_response import ProvisioningApiAppsGetApps200Response
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
    api_instance = nextcloud_client.ProvisioningApiAppsApi(api_client)
    ocs_api_request = 'true' # str |  (default to 'true')
    filter = 'filter_example' # str | Filter for enabled or disabled apps (optional)

    try:
        # Get a list of installed apps
        api_response = api_instance.provisioning_api_apps_get_apps(ocs_api_request, filter=filter)
        print("The response of ProvisioningApiAppsApi->provisioning_api_apps_get_apps:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ProvisioningApiAppsApi->provisioning_api_apps_get_apps: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]
 **filter** | **str**| Filter for enabled or disabled apps | [optional] 

### Return type

[**ProvisioningApiAppsGetApps200Response**](ProvisioningApiAppsGetApps200Response.md)

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

