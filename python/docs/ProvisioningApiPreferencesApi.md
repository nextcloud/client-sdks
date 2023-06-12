# openapi_client.ProvisioningApiPreferencesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**provisioning_api_preferences_delete_multiple_preference**](ProvisioningApiPreferencesApi.md#provisioning_api_preferences_delete_multiple_preference) | **DELETE** /ocs/v2.php/apps/provisioning_api/api/v1/config/users/{appId} | Delete multiple preferences for an app
[**provisioning_api_preferences_delete_preference**](ProvisioningApiPreferencesApi.md#provisioning_api_preferences_delete_preference) | **DELETE** /ocs/v2.php/apps/provisioning_api/api/v1/config/users/{appId}/{configKey} | Delete a preference for an app
[**provisioning_api_preferences_set_multiple_preferences**](ProvisioningApiPreferencesApi.md#provisioning_api_preferences_set_multiple_preferences) | **POST** /ocs/v2.php/apps/provisioning_api/api/v1/config/users/{appId} | Update multiple preference values of an app
[**provisioning_api_preferences_set_preference**](ProvisioningApiPreferencesApi.md#provisioning_api_preferences_set_preference) | **POST** /ocs/v2.php/apps/provisioning_api/api/v1/config/users/{appId}/{configKey} | Update a preference value of an app


# **provisioning_api_preferences_delete_multiple_preference**
> CoreWhatsNewDismiss200Response provisioning_api_preferences_delete_multiple_preference(config_keys, app_id, ocs_api_request)

Delete multiple preferences for an app

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
    api_instance = openapi_client.ProvisioningApiPreferencesApi(api_client)
    config_keys = 'config_keys_example' # str | Keys to delete
    app_id = 'app_id_example' # str | ID of the app
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Delete multiple preferences for an app
        api_response = api_instance.provisioning_api_preferences_delete_multiple_preference(config_keys, app_id, ocs_api_request)
        print("The response of ProvisioningApiPreferencesApi->provisioning_api_preferences_delete_multiple_preference:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ProvisioningApiPreferencesApi->provisioning_api_preferences_delete_multiple_preference: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **config_keys** | **str**| Keys to delete | 
 **app_id** | **str**| ID of the app | 
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
**200** | Preferences deleted successfully |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **provisioning_api_preferences_delete_preference**
> CoreWhatsNewDismiss200Response provisioning_api_preferences_delete_preference(app_id, config_key, ocs_api_request)

Delete a preference for an app

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
    api_instance = openapi_client.ProvisioningApiPreferencesApi(api_client)
    app_id = 'app_id_example' # str | ID of the app
    config_key = 'config_key_example' # str | Key to delete
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Delete a preference for an app
        api_response = api_instance.provisioning_api_preferences_delete_preference(app_id, config_key, ocs_api_request)
        print("The response of ProvisioningApiPreferencesApi->provisioning_api_preferences_delete_preference:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ProvisioningApiPreferencesApi->provisioning_api_preferences_delete_preference: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app_id** | **str**| ID of the app | 
 **config_key** | **str**| Key to delete | 
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
**200** | Preference deleted successfully |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **provisioning_api_preferences_set_multiple_preferences**
> CoreWhatsNewDismiss200Response provisioning_api_preferences_set_multiple_preferences(configs, app_id, ocs_api_request)

Update multiple preference values of an app

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
    api_instance = openapi_client.ProvisioningApiPreferencesApi(api_client)
    configs = 'configs_example' # str | Key-value pairs of the preferences
    app_id = 'app_id_example' # str | ID of the app
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Update multiple preference values of an app
        api_response = api_instance.provisioning_api_preferences_set_multiple_preferences(configs, app_id, ocs_api_request)
        print("The response of ProvisioningApiPreferencesApi->provisioning_api_preferences_set_multiple_preferences:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ProvisioningApiPreferencesApi->provisioning_api_preferences_set_multiple_preferences: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configs** | **str**| Key-value pairs of the preferences | 
 **app_id** | **str**| ID of the app | 
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
**200** | Preferences updated successfully |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **provisioning_api_preferences_set_preference**
> CoreWhatsNewDismiss200Response provisioning_api_preferences_set_preference(config_value, app_id, config_key, ocs_api_request)

Update a preference value of an app

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
    api_instance = openapi_client.ProvisioningApiPreferencesApi(api_client)
    config_value = 'config_value_example' # str | New value
    app_id = 'app_id_example' # str | ID of the app
    config_key = 'config_key_example' # str | Key of the preference
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Update a preference value of an app
        api_response = api_instance.provisioning_api_preferences_set_preference(config_value, app_id, config_key, ocs_api_request)
        print("The response of ProvisioningApiPreferencesApi->provisioning_api_preferences_set_preference:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ProvisioningApiPreferencesApi->provisioning_api_preferences_set_preference: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **config_value** | **str**| New value | 
 **app_id** | **str**| ID of the app | 
 **config_key** | **str**| Key of the preference | 
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
**200** | Preference updated successfully |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

