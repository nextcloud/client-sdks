# nextcloud_client.FilesTrashbinPreviewApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**files_trashbin_preview_get_preview**](FilesTrashbinPreviewApi.md#files_trashbin_preview_get_preview) | **GET** /index.php/apps/files_trashbin/preview | Get the preview for a file


# **files_trashbin_preview_get_preview**
> bytearray files_trashbin_preview_get_preview(file_id=file_id, x=x, y=y, a=a)

Get the preview for a file

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
    api_instance = nextcloud_client.FilesTrashbinPreviewApi(api_client)
    file_id = -1 # int | ID of the file (optional) (default to -1)
    x = 32 # int | Width of the preview (optional) (default to 32)
    y = 32 # int | Height of the preview (optional) (default to 32)
    a = 0 # int | Whether to not crop the preview (optional) (default to 0)

    try:
        # Get the preview for a file
        api_response = api_instance.files_trashbin_preview_get_preview(file_id=file_id, x=x, y=y, a=a)
        print("The response of FilesTrashbinPreviewApi->files_trashbin_preview_get_preview:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FilesTrashbinPreviewApi->files_trashbin_preview_get_preview: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file_id** | **int**| ID of the file | [optional] [default to -1]
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

