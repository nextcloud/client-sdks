# openapi_client.CoreAutoCompleteApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**core_auto_complete_get**](CoreAutoCompleteApi.md#core_auto_complete_get) | **GET** /ocs/v2.php/core/autocomplete/get | Autocomplete a query


# **core_auto_complete_get**
> CoreAutoCompleteGet200Response core_auto_complete_get(search, ocs_api_request, item_type=item_type, item_id=item_id, sorter=sorter, share_types=share_types, limit=limit)

Autocomplete a query

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import openapi_client
from openapi_client.models.core_auto_complete_get200_response import CoreAutoCompleteGet200Response
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
    api_instance = openapi_client.CoreAutoCompleteApi(api_client)
    search = 'search_example' # str | Text to search for
    ocs_api_request = 'true' # str |  (default to 'true')
    item_type = 'item_type_example' # str | Type of the items to search for (optional)
    item_id = 'item_id_example' # str | ID of the items to search for (optional)
    sorter = 'sorter_example' # str | can be piped, top prio first, e.g.: \"commenters|share-recipients\" (optional)
    share_types = 'share_types_example' # str | Types of shares to search for (optional)
    limit = 10 # int | Maximum number of results to return (optional) (default to 10)

    try:
        # Autocomplete a query
        api_response = api_instance.core_auto_complete_get(search, ocs_api_request, item_type=item_type, item_id=item_id, sorter=sorter, share_types=share_types, limit=limit)
        print("The response of CoreAutoCompleteApi->core_auto_complete_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CoreAutoCompleteApi->core_auto_complete_get: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **str**| Text to search for | 
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]
 **item_type** | **str**| Type of the items to search for | [optional] 
 **item_id** | **str**| ID of the items to search for | [optional] 
 **sorter** | **str**| can be piped, top prio first, e.g.: \&quot;commenters|share-recipients\&quot; | [optional] 
 **share_types** | **str**| Types of shares to search for | [optional] 
 **limit** | **int**| Maximum number of results to return | [optional] [default to 10]

### Return type

[**CoreAutoCompleteGet200Response**](CoreAutoCompleteGet200Response.md)

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

