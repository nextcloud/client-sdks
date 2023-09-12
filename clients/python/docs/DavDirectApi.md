# nextcloud_client.DavDirectApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**dav_direct_get_url**](DavDirectApi.md#dav_direct_get_url) | **POST** /ocs/v2.php/apps/dav/api/v1/direct | Get a direct link to a file


# **dav_direct_get_url**
> DavDirectGetUrl200Response dav_direct_get_url(file_id, ocs_api_request, expiration_time=expiration_time)

Get a direct link to a file

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.dav_direct_get_url200_response import DavDirectGetUrl200Response
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
    api_instance = nextcloud_client.DavDirectApi(api_client)
    file_id = 56 # int | ID of the file
    ocs_api_request = 'true' # str |  (default to 'true')
    expiration_time = 56 # int | Duration until the link expires (optional)

    try:
        # Get a direct link to a file
        api_response = api_instance.dav_direct_get_url(file_id, ocs_api_request, expiration_time=expiration_time)
        print("The response of DavDirectApi->dav_direct_get_url:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DavDirectApi->dav_direct_get_url: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file_id** | **int**| ID of the file | 
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]
 **expiration_time** | **int**| Duration until the link expires | [optional] 

### Return type

[**DavDirectGetUrl200Response**](DavDirectGetUrl200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Direct link returned |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

