# nextcloud_client.CoreUnifiedSearchApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**core_unified_search_get_providers**](CoreUnifiedSearchApi.md#core_unified_search_get_providers) | **GET** /ocs/v2.php/search/providers | Get the providers for unified search
[**core_unified_search_search**](CoreUnifiedSearchApi.md#core_unified_search_search) | **GET** /ocs/v2.php/search/providers/{providerId}/search | Search


# **core_unified_search_get_providers**
> CoreUnifiedSearchGetProviders200Response core_unified_search_get_providers(ocs_api_request, var_from=var_from)

Get the providers for unified search

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.core_unified_search_get_providers200_response import CoreUnifiedSearchGetProviders200Response
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
    api_instance = nextcloud_client.CoreUnifiedSearchApi(api_client)
    ocs_api_request = 'true' # str |  (default to 'true')
    var_from = '' # str | the url the user is currently at (optional) (default to '')

    try:
        # Get the providers for unified search
        api_response = api_instance.core_unified_search_get_providers(ocs_api_request, var_from=var_from)
        print("The response of CoreUnifiedSearchApi->core_unified_search_get_providers:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CoreUnifiedSearchApi->core_unified_search_get_providers: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]
 **var_from** | **str**| the url the user is currently at | [optional] [default to &#39;&#39;]

### Return type

[**CoreUnifiedSearchGetProviders200Response**](CoreUnifiedSearchGetProviders200Response.md)

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

# **core_unified_search_search**
> CoreUnifiedSearchSearch200Response core_unified_search_search(provider_id, ocs_api_request, term=term, sort_order=sort_order, limit=limit, cursor=cursor, var_from=var_from)

Search

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.core_unified_search_search200_response import CoreUnifiedSearchSearch200Response
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
    api_instance = nextcloud_client.CoreUnifiedSearchApi(api_client)
    provider_id = 'provider_id_example' # str | ID of the provider
    ocs_api_request = 'true' # str |  (default to 'true')
    term = '' # str | Term to search (optional) (default to '')
    sort_order = 56 # int | Order of entries (optional)
    limit = 56 # int | Maximum amount of entries (optional)
    cursor = 'cursor_example' # str | Offset for searching (optional)
    var_from = '' # str | The current user URL (optional) (default to '')

    try:
        # Search
        api_response = api_instance.core_unified_search_search(provider_id, ocs_api_request, term=term, sort_order=sort_order, limit=limit, cursor=cursor, var_from=var_from)
        print("The response of CoreUnifiedSearchApi->core_unified_search_search:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CoreUnifiedSearchApi->core_unified_search_search: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **provider_id** | **str**| ID of the provider | 
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]
 **term** | **str**| Term to search | [optional] [default to &#39;&#39;]
 **sort_order** | **int**| Order of entries | [optional] 
 **limit** | **int**| Maximum amount of entries | [optional] 
 **cursor** | **str**| Offset for searching | [optional] 
 **var_from** | **str**| The current user URL | [optional] [default to &#39;&#39;]

### Return type

[**CoreUnifiedSearchSearch200Response**](CoreUnifiedSearchSearch200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Search entries returned |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

