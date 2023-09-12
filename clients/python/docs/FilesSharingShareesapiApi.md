# nextcloud_client.FilesSharingShareesapiApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**files_sharing_shareesapi_find_recommended**](FilesSharingShareesapiApi.md#files_sharing_shareesapi_find_recommended) | **GET** /ocs/v2.php/apps/files_sharing/api/v1/sharees_recommended | Find recommended sharees
[**files_sharing_shareesapi_search**](FilesSharingShareesapiApi.md#files_sharing_shareesapi_search) | **GET** /ocs/v2.php/apps/files_sharing/api/v1/sharees | Search for sharees


# **files_sharing_shareesapi_find_recommended**
> FilesSharingShareesapiFindRecommended200Response files_sharing_shareesapi_find_recommended(item_type, ocs_api_request, share_type=share_type)

Find recommended sharees

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.files_sharing_shareesapi_find_recommended200_response import FilesSharingShareesapiFindRecommended200Response
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
    api_instance = nextcloud_client.FilesSharingShareesapiApi(api_client)
    item_type = 'item_type_example' # str | Limit to specific item types
    ocs_api_request = 'true' # str |  (default to 'true')
    share_type = 'share_type_example' # str | Limit to specific share types (optional)

    try:
        # Find recommended sharees
        api_response = api_instance.files_sharing_shareesapi_find_recommended(item_type, ocs_api_request, share_type=share_type)
        print("The response of FilesSharingShareesapiApi->files_sharing_shareesapi_find_recommended:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FilesSharingShareesapiApi->files_sharing_shareesapi_find_recommended: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **item_type** | **str**| Limit to specific item types | 
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]
 **share_type** | **str**| Limit to specific share types | [optional] 

### Return type

[**FilesSharingShareesapiFindRecommended200Response**](FilesSharingShareesapiFindRecommended200Response.md)

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

# **files_sharing_shareesapi_search**
> FilesSharingShareesapiSearch200Response files_sharing_shareesapi_search(ocs_api_request, search=search, item_type=item_type, page=page, per_page=per_page, share_type=share_type, lookup=lookup)

Search for sharees

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.files_sharing_shareesapi_search200_response import FilesSharingShareesapiSearch200Response
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
    api_instance = nextcloud_client.FilesSharingShareesapiApi(api_client)
    ocs_api_request = 'true' # str |  (default to 'true')
    search = '' # str | Text to search for (optional) (default to '')
    item_type = 'item_type_example' # str | Limit to specific item types (optional)
    page = 1 # int | Page offset for searching (optional) (default to 1)
    per_page = 200 # int | Limit amount of search results per page (optional) (default to 200)
    share_type = 'share_type_example' # str | Limit to specific share types (optional)
    lookup = 0 # int | If a global lookup should be performed too (optional) (default to 0)

    try:
        # Search for sharees
        api_response = api_instance.files_sharing_shareesapi_search(ocs_api_request, search=search, item_type=item_type, page=page, per_page=per_page, share_type=share_type, lookup=lookup)
        print("The response of FilesSharingShareesapiApi->files_sharing_shareesapi_search:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FilesSharingShareesapiApi->files_sharing_shareesapi_search: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]
 **search** | **str**| Text to search for | [optional] [default to &#39;&#39;]
 **item_type** | **str**| Limit to specific item types | [optional] 
 **page** | **int**| Page offset for searching | [optional] [default to 1]
 **per_page** | **int**| Limit amount of search results per page | [optional] [default to 200]
 **share_type** | **str**| Limit to specific share types | [optional] 
 **lookup** | **int**| If a global lookup should be performed too | [optional] [default to 0]

### Return type

[**FilesSharingShareesapiSearch200Response**](FilesSharingShareesapiSearch200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Sharees search result returned |  * Link -  <br>  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

