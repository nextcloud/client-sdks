# openapi_client.CorePreviewApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**core_preview_get_preview**](CorePreviewApi.md#core_preview_get_preview) | **GET** /index.php/core/preview.png | Get a preview by file ID
[**core_preview_get_preview_by_file_id**](CorePreviewApi.md#core_preview_get_preview_by_file_id) | **GET** /index.php/core/preview | Get a preview by file ID


# **core_preview_get_preview**
> bytearray core_preview_get_preview(file=file, x=x, y=y, a=a, force_icon=force_icon, mode=mode)

Get a preview by file ID

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
    api_instance = openapi_client.CorePreviewApi(api_client)
    file = '' # str | Path of the file (optional) (default to '')
    x = 32 # int | Width of the preview (optional) (default to 32)
    y = 32 # int | Height of the preview (optional) (default to 32)
    a = 0 # int | Not crop the preview (optional) (default to 0)
    force_icon = 1 # int | Force returning an icon (optional) (default to 1)
    mode = 'fill' # str | How to crop the image (optional) (default to 'fill')

    try:
        # Get a preview by file ID
        api_response = api_instance.core_preview_get_preview(file=file, x=x, y=y, a=a, force_icon=force_icon, mode=mode)
        print("The response of CorePreviewApi->core_preview_get_preview:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CorePreviewApi->core_preview_get_preview: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | **str**| Path of the file | [optional] [default to &#39;&#39;]
 **x** | **int**| Width of the preview | [optional] [default to 32]
 **y** | **int**| Height of the preview | [optional] [default to 32]
 **a** | **int**| Not crop the preview | [optional] [default to 0]
 **force_icon** | **int**| Force returning an icon | [optional] [default to 1]
 **mode** | **str**| How to crop the image | [optional] [default to &#39;fill&#39;]

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
**200** | Preview returned |  * Content-Disposition -  <br>  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **core_preview_get_preview_by_file_id**
> bytearray core_preview_get_preview_by_file_id(file_id=file_id, x=x, y=y, a=a, force_icon=force_icon, mode=mode)

Get a preview by file ID

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
    api_instance = openapi_client.CorePreviewApi(api_client)
    file_id = 1 # int | ID of the file (optional) (default to 1)
    x = 32 # int | Width of the preview (optional) (default to 32)
    y = 32 # int | Height of the preview (optional) (default to 32)
    a = 0 # int | Not crop the preview (optional) (default to 0)
    force_icon = 1 # int | Force returning an icon (optional) (default to 1)
    mode = 'fill' # str | How to crop the image (optional) (default to 'fill')

    try:
        # Get a preview by file ID
        api_response = api_instance.core_preview_get_preview_by_file_id(file_id=file_id, x=x, y=y, a=a, force_icon=force_icon, mode=mode)
        print("The response of CorePreviewApi->core_preview_get_preview_by_file_id:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CorePreviewApi->core_preview_get_preview_by_file_id: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file_id** | **int**| ID of the file | [optional] [default to 1]
 **x** | **int**| Width of the preview | [optional] [default to 32]
 **y** | **int**| Height of the preview | [optional] [default to 32]
 **a** | **int**| Not crop the preview | [optional] [default to 0]
 **force_icon** | **int**| Force returning an icon | [optional] [default to 1]
 **mode** | **str**| How to crop the image | [optional] [default to &#39;fill&#39;]

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
**200** | Preview returned |  * Content-Disposition -  <br>  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

