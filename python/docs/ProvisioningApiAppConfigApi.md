# openapi_client.ProvisioningApiAppConfigApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**provisioning_api_app_config_delete_key**](ProvisioningApiAppConfigApi.md#provisioning_api_app_config_delete_key) | **DELETE** /ocs/v2.php/apps/provisioning_api/api/v1/config/apps/{app}/{key} | Delete a config key of an app
[**provisioning_api_app_config_get_apps**](ProvisioningApiAppConfigApi.md#provisioning_api_app_config_get_apps) | **GET** /ocs/v2.php/apps/provisioning_api/api/v1/config/apps | Get a list of apps
[**provisioning_api_app_config_get_keys**](ProvisioningApiAppConfigApi.md#provisioning_api_app_config_get_keys) | **GET** /ocs/v2.php/apps/provisioning_api/api/v1/config/apps/{app} | Get the config keys of an app
[**provisioning_api_app_config_get_value**](ProvisioningApiAppConfigApi.md#provisioning_api_app_config_get_value) | **GET** /ocs/v2.php/apps/provisioning_api/api/v1/config/apps/{app}/{key} | Get a the config value of an app
[**provisioning_api_app_config_set_value**](ProvisioningApiAppConfigApi.md#provisioning_api_app_config_set_value) | **POST** /ocs/v2.php/apps/provisioning_api/api/v1/config/apps/{app}/{key} | Update the config value of an app


# **provisioning_api_app_config_delete_key**
> CoreWhatsNewDismiss200Response provisioning_api_app_config_delete_key(app, key, ocs_api_request)

Delete a config key of an app

This endpoint requires admin access

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
    api_instance = openapi_client.ProvisioningApiAppConfigApi(api_client)
    app = 'app_example' # str | ID of the app
    key = 'key_example' # str | Key to delete
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Delete a config key of an app
        api_response = api_instance.provisioning_api_app_config_delete_key(app, key, ocs_api_request)
        print("The response of ProvisioningApiAppConfigApi->provisioning_api_app_config_delete_key:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ProvisioningApiAppConfigApi->provisioning_api_app_config_delete_key: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app** | **str**| ID of the app | 
 **key** | **str**| Key to delete | 
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
**200** | Key deleted successfully |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **provisioning_api_app_config_get_apps**
> ProvisioningApiAppConfigGetApps200Response provisioning_api_app_config_get_apps(ocs_api_request)

Get a list of apps

This endpoint requires admin access

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import openapi_client
from openapi_client.models.provisioning_api_app_config_get_apps200_response import ProvisioningApiAppConfigGetApps200Response
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
    api_instance = openapi_client.ProvisioningApiAppConfigApi(api_client)
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Get a list of apps
        api_response = api_instance.provisioning_api_app_config_get_apps(ocs_api_request)
        print("The response of ProvisioningApiAppConfigApi->provisioning_api_app_config_get_apps:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ProvisioningApiAppConfigApi->provisioning_api_app_config_get_apps: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]

### Return type

[**ProvisioningApiAppConfigGetApps200Response**](ProvisioningApiAppConfigGetApps200Response.md)

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

# **provisioning_api_app_config_get_keys**
> ProvisioningApiAppConfigGetApps200Response provisioning_api_app_config_get_keys(app, ocs_api_request)

Get the config keys of an app

This endpoint requires admin access

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import openapi_client
from openapi_client.models.provisioning_api_app_config_get_apps200_response import ProvisioningApiAppConfigGetApps200Response
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
    api_instance = openapi_client.ProvisioningApiAppConfigApi(api_client)
    app = 'app_example' # str | ID of the app
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Get the config keys of an app
        api_response = api_instance.provisioning_api_app_config_get_keys(app, ocs_api_request)
        print("The response of ProvisioningApiAppConfigApi->provisioning_api_app_config_get_keys:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ProvisioningApiAppConfigApi->provisioning_api_app_config_get_keys: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app** | **str**| ID of the app | 
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]

### Return type

[**ProvisioningApiAppConfigGetApps200Response**](ProvisioningApiAppConfigGetApps200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Keys returned |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **provisioning_api_app_config_get_value**
> ProvisioningApiAppConfigGetValue200Response provisioning_api_app_config_get_value(app, key, ocs_api_request, default_value=default_value)

Get a the config value of an app

This endpoint requires admin access

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import openapi_client
from openapi_client.models.provisioning_api_app_config_get_value200_response import ProvisioningApiAppConfigGetValue200Response
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
    api_instance = openapi_client.ProvisioningApiAppConfigApi(api_client)
    app = 'app_example' # str | ID if the app
    key = 'key_example' # str | Key
    ocs_api_request = 'true' # str |  (default to 'true')
    default_value = '' # str | Default returned value if the value is empty (optional) (default to '')

    try:
        # Get a the config value of an app
        api_response = api_instance.provisioning_api_app_config_get_value(app, key, ocs_api_request, default_value=default_value)
        print("The response of ProvisioningApiAppConfigApi->provisioning_api_app_config_get_value:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ProvisioningApiAppConfigApi->provisioning_api_app_config_get_value: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app** | **str**| ID if the app | 
 **key** | **str**| Key | 
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]
 **default_value** | **str**| Default returned value if the value is empty | [optional] [default to &#39;&#39;]

### Return type

[**ProvisioningApiAppConfigGetValue200Response**](ProvisioningApiAppConfigGetValue200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Value returned |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **provisioning_api_app_config_set_value**
> CoreWhatsNewDismiss200Response provisioning_api_app_config_set_value(value, app, key, ocs_api_request)

Update the config value of an app

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
    api_instance = openapi_client.ProvisioningApiAppConfigApi(api_client)
    value = 'value_example' # str | New value for the key
    app = 'app_example' # str | ID of the app
    key = 'key_example' # str | Key to update
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Update the config value of an app
        api_response = api_instance.provisioning_api_app_config_set_value(value, app, key, ocs_api_request)
        print("The response of ProvisioningApiAppConfigApi->provisioning_api_app_config_set_value:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ProvisioningApiAppConfigApi->provisioning_api_app_config_set_value: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **value** | **str**| New value for the key | 
 **app** | **str**| ID of the app | 
 **key** | **str**| Key to update | 
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
**200** | Value updated successfully |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

