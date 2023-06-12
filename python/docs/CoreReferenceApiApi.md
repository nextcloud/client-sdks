# openapi_client.CoreReferenceApiApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**core_reference_api_extract**](CoreReferenceApiApi.md#core_reference_api_extract) | **POST** /ocs/v2.php/references/extract | Extract references from a text
[**core_reference_api_get_providers_info**](CoreReferenceApiApi.md#core_reference_api_get_providers_info) | **GET** /ocs/v2.php/references/providers | Get the providers
[**core_reference_api_resolve**](CoreReferenceApiApi.md#core_reference_api_resolve) | **POST** /ocs/v2.php/references/resolve | Resolve multiple references
[**core_reference_api_resolve_one**](CoreReferenceApiApi.md#core_reference_api_resolve_one) | **GET** /ocs/v2.php/references/resolve | Resolve a reference
[**core_reference_api_touch_provider**](CoreReferenceApiApi.md#core_reference_api_touch_provider) | **PUT** /ocs/v2.php/references/provider/{providerId} | Touch a provider


# **core_reference_api_extract**
> CoreReferenceApiResolveOne200Response core_reference_api_extract(text, ocs_api_request, resolve=resolve, limit=limit)

Extract references from a text

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import openapi_client
from openapi_client.models.core_reference_api_resolve_one200_response import CoreReferenceApiResolveOne200Response
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
    api_instance = openapi_client.CoreReferenceApiApi(api_client)
    text = 'text_example' # str | Text to extract from
    ocs_api_request = 'true' # str |  (default to 'true')
    resolve = 0 # int | Resolve the references (optional) (default to 0)
    limit = 1 # int | Maximum amount of references to extract (optional) (default to 1)

    try:
        # Extract references from a text
        api_response = api_instance.core_reference_api_extract(text, ocs_api_request, resolve=resolve, limit=limit)
        print("The response of CoreReferenceApiApi->core_reference_api_extract:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CoreReferenceApiApi->core_reference_api_extract: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **text** | **str**| Text to extract from | 
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]
 **resolve** | **int**| Resolve the references | [optional] [default to 0]
 **limit** | **int**| Maximum amount of references to extract | [optional] [default to 1]

### Return type

[**CoreReferenceApiResolveOne200Response**](CoreReferenceApiResolveOne200Response.md)

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

# **core_reference_api_get_providers_info**
> CoreReferenceApiGetProvidersInfo200Response core_reference_api_get_providers_info(ocs_api_request)

Get the providers

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import openapi_client
from openapi_client.models.core_reference_api_get_providers_info200_response import CoreReferenceApiGetProvidersInfo200Response
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
    api_instance = openapi_client.CoreReferenceApiApi(api_client)
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Get the providers
        api_response = api_instance.core_reference_api_get_providers_info(ocs_api_request)
        print("The response of CoreReferenceApiApi->core_reference_api_get_providers_info:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CoreReferenceApiApi->core_reference_api_get_providers_info: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]

### Return type

[**CoreReferenceApiGetProvidersInfo200Response**](CoreReferenceApiGetProvidersInfo200Response.md)

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

# **core_reference_api_resolve**
> CoreReferenceApiResolveOne200Response core_reference_api_resolve(references, ocs_api_request, limit=limit)

Resolve multiple references

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import openapi_client
from openapi_client.models.core_reference_api_resolve_one200_response import CoreReferenceApiResolveOne200Response
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
    api_instance = openapi_client.CoreReferenceApiApi(api_client)
    references = 'references_example' # str | References to resolve
    ocs_api_request = 'true' # str |  (default to 'true')
    limit = 1 # int | Maximum amount of references to resolve (optional) (default to 1)

    try:
        # Resolve multiple references
        api_response = api_instance.core_reference_api_resolve(references, ocs_api_request, limit=limit)
        print("The response of CoreReferenceApiApi->core_reference_api_resolve:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CoreReferenceApiApi->core_reference_api_resolve: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **references** | **str**| References to resolve | 
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]
 **limit** | **int**| Maximum amount of references to resolve | [optional] [default to 1]

### Return type

[**CoreReferenceApiResolveOne200Response**](CoreReferenceApiResolveOne200Response.md)

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

# **core_reference_api_resolve_one**
> CoreReferenceApiResolveOne200Response core_reference_api_resolve_one(reference, ocs_api_request)

Resolve a reference

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import openapi_client
from openapi_client.models.core_reference_api_resolve_one200_response import CoreReferenceApiResolveOne200Response
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
    api_instance = openapi_client.CoreReferenceApiApi(api_client)
    reference = 'reference_example' # str | Reference to resolve
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Resolve a reference
        api_response = api_instance.core_reference_api_resolve_one(reference, ocs_api_request)
        print("The response of CoreReferenceApiApi->core_reference_api_resolve_one:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CoreReferenceApiApi->core_reference_api_resolve_one: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reference** | **str**| Reference to resolve | 
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]

### Return type

[**CoreReferenceApiResolveOne200Response**](CoreReferenceApiResolveOne200Response.md)

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

# **core_reference_api_touch_provider**
> CoreReferenceApiTouchProvider200Response core_reference_api_touch_provider(provider_id, ocs_api_request, timestamp=timestamp)

Touch a provider

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import openapi_client
from openapi_client.models.core_reference_api_touch_provider200_response import CoreReferenceApiTouchProvider200Response
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
    api_instance = openapi_client.CoreReferenceApiApi(api_client)
    provider_id = 'provider_id_example' # str | ID of the provider
    ocs_api_request = 'true' # str |  (default to 'true')
    timestamp = 56 # int | Timestamp of the last usage (optional)

    try:
        # Touch a provider
        api_response = api_instance.core_reference_api_touch_provider(provider_id, ocs_api_request, timestamp=timestamp)
        print("The response of CoreReferenceApiApi->core_reference_api_touch_provider:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CoreReferenceApiApi->core_reference_api_touch_provider: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **provider_id** | **str**| ID of the provider | 
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]
 **timestamp** | **int**| Timestamp of the last usage | [optional] 

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

