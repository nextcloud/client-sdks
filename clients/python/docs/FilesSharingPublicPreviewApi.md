# nextcloud_client.FilesSharingPublicPreviewApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**files_sharing_public_preview_direct_link**](FilesSharingPublicPreviewApi.md#files_sharing_public_preview_direct_link) | **GET** /index.php/s/{token}/preview | Get a direct link preview for a shared file
[**files_sharing_public_preview_get_preview**](FilesSharingPublicPreviewApi.md#files_sharing_public_preview_get_preview) | **GET** /index.php/apps/files_sharing/publicpreview/{token} | Get a preview for a shared file


# **files_sharing_public_preview_direct_link**
> bytearray files_sharing_public_preview_direct_link(token, ocs_api_request)

Get a direct link preview for a shared file

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
    api_instance = nextcloud_client.FilesSharingPublicPreviewApi(api_client)
    token = 'token_example' # str | Token of the share
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Get a direct link preview for a shared file
        api_response = api_instance.files_sharing_public_preview_direct_link(token, ocs_api_request)
        print("The response of FilesSharingPublicPreviewApi->files_sharing_public_preview_direct_link:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FilesSharingPublicPreviewApi->files_sharing_public_preview_direct_link: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **str**| Token of the share | 
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]

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
**200** | Preview returned |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **files_sharing_public_preview_get_preview**
> bytearray files_sharing_public_preview_get_preview(token, ocs_api_request, file=file, x=x, y=y, a=a)

Get a preview for a shared file

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
    api_instance = nextcloud_client.FilesSharingPublicPreviewApi(api_client)
    token = 'token_example' # str | Token of the share
    ocs_api_request = 'true' # str |  (default to 'true')
    file = '' # str | File in the share (optional) (default to '')
    x = 32 # int | Width of the preview (optional) (default to 32)
    y = 32 # int | Height of the preview (optional) (default to 32)
    a = 0 # int | Whether to not crop the preview (optional) (default to 0)

    try:
        # Get a preview for a shared file
        api_response = api_instance.files_sharing_public_preview_get_preview(token, ocs_api_request, file=file, x=x, y=y, a=a)
        print("The response of FilesSharingPublicPreviewApi->files_sharing_public_preview_get_preview:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FilesSharingPublicPreviewApi->files_sharing_public_preview_get_preview: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **str**| Token of the share | 
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]
 **file** | **str**| File in the share | [optional] [default to &#39;&#39;]
 **x** | **int**| Width of the preview | [optional] [default to 32]
 **y** | **int**| Height of the preview | [optional] [default to 32]
 **a** | **int**| Whether to not crop the preview | [optional] [default to 0]

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
**200** | Preview returned |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

