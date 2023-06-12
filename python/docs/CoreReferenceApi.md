# openapi_client.CoreReferenceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**core_reference_preview**](CoreReferenceApi.md#core_reference_preview) | **GET** /index.php/core/references/preview/{referenceId} | Get a preview for a reference


# **core_reference_preview**
> bytearray core_reference_preview(reference_id)

Get a preview for a reference

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
    api_instance = openapi_client.CoreReferenceApi(api_client)
    reference_id = 'reference_id_example' # str | the reference cache key

    try:
        # Get a preview for a reference
        api_response = api_instance.core_reference_preview(reference_id)
        print("The response of CoreReferenceApi->core_reference_preview:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CoreReferenceApi->core_reference_preview: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reference_id** | **str**| the reference cache key | 

### Return type

**bytearray**

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: image/*

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Preview returned |  * Content-Disposition -  <br>  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

