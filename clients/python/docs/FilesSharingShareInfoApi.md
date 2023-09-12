# nextcloud_client.FilesSharingShareInfoApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**files_sharing_share_info_info**](FilesSharingShareInfoApi.md#files_sharing_share_info_info) | **POST** /index.php/apps/files_sharing/shareinfo | Get the info about a share


# **files_sharing_share_info_info**
> FilesSharingShareInfo files_sharing_share_info_info(t, password=password, dir=dir, depth=depth)

Get the info about a share

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.files_sharing_share_info import FilesSharingShareInfo
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
    api_instance = nextcloud_client.FilesSharingShareInfoApi(api_client)
    t = 't_example' # str | Token of the share
    password = 'password_example' # str | Password of the share (optional)
    dir = 'dir_example' # str | Subdirectory to get info about (optional)
    depth = -1 # int | Maximum depth to get info about (optional) (default to -1)

    try:
        # Get the info about a share
        api_response = api_instance.files_sharing_share_info_info(t, password=password, dir=dir, depth=depth)
        print("The response of FilesSharingShareInfoApi->files_sharing_share_info_info:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FilesSharingShareInfoApi->files_sharing_share_info_info: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **t** | **str**| Token of the share | 
 **password** | **str**| Password of the share | [optional] 
 **dir** | **str**| Subdirectory to get info about | [optional] 
 **depth** | **int**| Maximum depth to get info about | [optional] [default to -1]

### Return type

[**FilesSharingShareInfo**](FilesSharingShareInfo.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Share info returned |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

