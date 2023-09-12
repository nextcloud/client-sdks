# nextcloud_client.FilesApiApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**files_api_get_thumbnail**](FilesApiApi.md#files_api_get_thumbnail) | **GET** /index.php/apps/files/api/v1/thumbnail/{x}/{y}/{file} | Gets a thumbnail of the specified file
[**files_api_service_worker**](FilesApiApi.md#files_api_service_worker) | **GET** /index.php/apps/files/preview-service-worker.js | Get the service-worker Javascript for previews


# **files_api_get_thumbnail**
> bytearray files_api_get_thumbnail(x, y, file)

Gets a thumbnail of the specified file

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
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
    api_instance = nextcloud_client.FilesApiApi(api_client)
    x = 56 # int | Width of the thumbnail
    y = 56 # int | Height of the thumbnail
    file = 'file_example' # str | URL-encoded filename

    try:
        # Gets a thumbnail of the specified file
        api_response = api_instance.files_api_get_thumbnail(x, y, file)
        print("The response of FilesApiApi->files_api_get_thumbnail:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FilesApiApi->files_api_get_thumbnail: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **x** | **int**| Width of the thumbnail | 
 **y** | **int**| Height of the thumbnail | 
 **file** | **str**| URL-encoded filename | 

### Return type

**bytearray**

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Thumbnail returned |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **files_api_service_worker**
> bytearray files_api_service_worker()

Get the service-worker Javascript for previews

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
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
    api_instance = nextcloud_client.FilesApiApi(api_client)

    try:
        # Get the service-worker Javascript for previews
        api_response = api_instance.files_api_service_worker()
        print("The response of FilesApiApi->files_api_service_worker:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FilesApiApi->files_api_service_worker: %s\n" % e)
```



### Parameters
This endpoint does not need any parameter.

### Return type

**bytearray**

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/javascript

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** |  |  * Service-Worker-Allowed -  <br>  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

