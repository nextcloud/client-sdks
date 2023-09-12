# nextcloud_client.FilesOpenLocalEditorApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**files_open_local_editor_create**](FilesOpenLocalEditorApi.md#files_open_local_editor_create) | **POST** /ocs/v2.php/apps/files/api/v1/openlocaleditor | Create a local editor
[**files_open_local_editor_validate**](FilesOpenLocalEditorApi.md#files_open_local_editor_validate) | **POST** /ocs/v2.php/apps/files/api/v1/openlocaleditor/{token} | Validate a local editor


# **files_open_local_editor_create**
> FilesOpenLocalEditorCreate200Response files_open_local_editor_create(path, ocs_api_request)

Create a local editor

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.files_open_local_editor_create200_response import FilesOpenLocalEditorCreate200Response
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
    api_instance = nextcloud_client.FilesOpenLocalEditorApi(api_client)
    path = 'path_example' # str | Path of the file
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Create a local editor
        api_response = api_instance.files_open_local_editor_create(path, ocs_api_request)
        print("The response of FilesOpenLocalEditorApi->files_open_local_editor_create:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FilesOpenLocalEditorApi->files_open_local_editor_create: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **str**| Path of the file | 
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]

### Return type

[**FilesOpenLocalEditorCreate200Response**](FilesOpenLocalEditorCreate200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Local editor returned |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **files_open_local_editor_validate**
> FilesOpenLocalEditorValidate200Response files_open_local_editor_validate(path, token, ocs_api_request)

Validate a local editor

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.files_open_local_editor_validate200_response import FilesOpenLocalEditorValidate200Response
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
    api_instance = nextcloud_client.FilesOpenLocalEditorApi(api_client)
    path = 'path_example' # str | Path of the file
    token = 'token_example' # str | Token of the local editor
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Validate a local editor
        api_response = api_instance.files_open_local_editor_validate(path, token, ocs_api_request)
        print("The response of FilesOpenLocalEditorApi->files_open_local_editor_validate:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FilesOpenLocalEditorApi->files_open_local_editor_validate: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **str**| Path of the file | 
 **token** | **str**| Token of the local editor | 
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]

### Return type

[**FilesOpenLocalEditorValidate200Response**](FilesOpenLocalEditorValidate200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Local editor validated successfully |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

