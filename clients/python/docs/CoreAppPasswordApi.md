# nextcloud_client.CoreAppPasswordApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**core_app_password_delete_app_password**](CoreAppPasswordApi.md#core_app_password_delete_app_password) | **DELETE** /ocs/v2.php/core/apppassword | Delete app password
[**core_app_password_get_app_password**](CoreAppPasswordApi.md#core_app_password_get_app_password) | **GET** /ocs/v2.php/core/getapppassword | Create app password
[**core_app_password_rotate_app_password**](CoreAppPasswordApi.md#core_app_password_rotate_app_password) | **POST** /ocs/v2.php/core/apppassword/rotate | Rotate app password


# **core_app_password_delete_app_password**
> CoreWhatsNewDismiss200Response core_app_password_delete_app_password(ocs_api_request)

Delete app password

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
    api_instance = nextcloud_client.CoreAppPasswordApi(api_client)
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Delete app password
        api_response = api_instance.core_app_password_delete_app_password(ocs_api_request)
        print("The response of CoreAppPasswordApi->core_app_password_delete_app_password:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CoreAppPasswordApi->core_app_password_delete_app_password: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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
**200** | App password deleted successfully |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **core_app_password_get_app_password**
> CoreAppPasswordGetAppPassword200Response core_app_password_get_app_password(ocs_api_request)

Create app password

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.core_app_password_get_app_password200_response import CoreAppPasswordGetAppPassword200Response
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
    api_instance = nextcloud_client.CoreAppPasswordApi(api_client)
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Create app password
        api_response = api_instance.core_app_password_get_app_password(ocs_api_request)
        print("The response of CoreAppPasswordApi->core_app_password_get_app_password:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CoreAppPasswordApi->core_app_password_get_app_password: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]

### Return type

[**CoreAppPasswordGetAppPassword200Response**](CoreAppPasswordGetAppPassword200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | App password returned |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **core_app_password_rotate_app_password**
> CoreAppPasswordGetAppPassword200Response core_app_password_rotate_app_password(ocs_api_request)

Rotate app password

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.core_app_password_get_app_password200_response import CoreAppPasswordGetAppPassword200Response
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
    api_instance = nextcloud_client.CoreAppPasswordApi(api_client)
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Rotate app password
        api_response = api_instance.core_app_password_rotate_app_password(ocs_api_request)
        print("The response of CoreAppPasswordApi->core_app_password_rotate_app_password:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CoreAppPasswordApi->core_app_password_rotate_app_password: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]

### Return type

[**CoreAppPasswordGetAppPassword200Response**](CoreAppPasswordGetAppPassword200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | App password returned |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

