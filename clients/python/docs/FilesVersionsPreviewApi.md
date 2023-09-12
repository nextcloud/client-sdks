# nextcloud_client.FilesVersionsPreviewApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**files_versions_preview_get_preview**](FilesVersionsPreviewApi.md#files_versions_preview_get_preview) | **GET** /index.php/apps/files_versions/preview | Get the preview for a file version


# **files_versions_preview_get_preview**
> bytearray files_versions_preview_get_preview(file=file, x=x, y=y, version=version)

Get the preview for a file version

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
    api_instance = nextcloud_client.FilesVersionsPreviewApi(api_client)
    file = '' # str | Path of the file (optional) (default to '')
    x = 44 # int | Width of the preview (optional) (default to 44)
    y = 44 # int | Height of the preview (optional) (default to 44)
    version = '' # str | Version of the file to get the preview for (optional) (default to '')

    try:
        # Get the preview for a file version
        api_response = api_instance.files_versions_preview_get_preview(file=file, x=x, y=y, version=version)
        print("The response of FilesVersionsPreviewApi->files_versions_preview_get_preview:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FilesVersionsPreviewApi->files_versions_preview_get_preview: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | **str**| Path of the file | [optional] [default to &#39;&#39;]
 **x** | **int**| Width of the preview | [optional] [default to 44]
 **y** | **int**| Height of the preview | [optional] [default to 44]
 **version** | **str**| Version of the file to get the preview for | [optional] [default to &#39;&#39;]

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

