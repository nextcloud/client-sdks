# nextcloud_client.FilesDirectEditingApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**files_direct_editing_create**](FilesDirectEditingApi.md#files_direct_editing_create) | **POST** /ocs/v2.php/apps/files/api/v1/directEditing/create | Create a file for direct editing
[**files_direct_editing_info**](FilesDirectEditingApi.md#files_direct_editing_info) | **GET** /ocs/v2.php/apps/files/api/v1/directEditing | Get the direct editing capabilities
[**files_direct_editing_open**](FilesDirectEditingApi.md#files_direct_editing_open) | **POST** /ocs/v2.php/apps/files/api/v1/directEditing/open | Open a file for direct editing
[**files_direct_editing_templates**](FilesDirectEditingApi.md#files_direct_editing_templates) | **GET** /ocs/v2.php/apps/files/api/v1/directEditing/templates/{editorId}/{creatorId} | Get the templates for direct editing


# **files_direct_editing_create**
> DavDirectGetUrl200Response files_direct_editing_create(path, editor_id, creator_id, ocs_api_request, template_id=template_id)

Create a file for direct editing

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
    api_instance = nextcloud_client.FilesDirectEditingApi(api_client)
    path = 'path_example' # str | Path of the file
    editor_id = 'editor_id_example' # str | ID of the editor
    creator_id = 'creator_id_example' # str | ID of the creator
    ocs_api_request = 'true' # str |  (default to 'true')
    template_id = 'template_id_example' # str | ID of the template (optional)

    try:
        # Create a file for direct editing
        api_response = api_instance.files_direct_editing_create(path, editor_id, creator_id, ocs_api_request, template_id=template_id)
        print("The response of FilesDirectEditingApi->files_direct_editing_create:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FilesDirectEditingApi->files_direct_editing_create: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **str**| Path of the file | 
 **editor_id** | **str**| ID of the editor | 
 **creator_id** | **str**| ID of the creator | 
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]
 **template_id** | **str**| ID of the template | [optional] 

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
**200** | URL for direct editing returned |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **files_direct_editing_info**
> FilesDirectEditingInfo200Response files_direct_editing_info(ocs_api_request)

Get the direct editing capabilities

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.files_direct_editing_info200_response import FilesDirectEditingInfo200Response
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
    api_instance = nextcloud_client.FilesDirectEditingApi(api_client)
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Get the direct editing capabilities
        api_response = api_instance.files_direct_editing_info(ocs_api_request)
        print("The response of FilesDirectEditingApi->files_direct_editing_info:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FilesDirectEditingApi->files_direct_editing_info: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]

### Return type

[**FilesDirectEditingInfo200Response**](FilesDirectEditingInfo200Response.md)

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

# **files_direct_editing_open**
> DavDirectGetUrl200Response files_direct_editing_open(path, ocs_api_request, editor_id=editor_id, file_id=file_id)

Open a file for direct editing

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
    api_instance = nextcloud_client.FilesDirectEditingApi(api_client)
    path = 'path_example' # str | Path of the file
    ocs_api_request = 'true' # str |  (default to 'true')
    editor_id = 'editor_id_example' # str | ID of the editor (optional)
    file_id = 56 # int | ID of the file (optional)

    try:
        # Open a file for direct editing
        api_response = api_instance.files_direct_editing_open(path, ocs_api_request, editor_id=editor_id, file_id=file_id)
        print("The response of FilesDirectEditingApi->files_direct_editing_open:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FilesDirectEditingApi->files_direct_editing_open: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **str**| Path of the file | 
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]
 **editor_id** | **str**| ID of the editor | [optional] 
 **file_id** | **int**| ID of the file | [optional] 

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
**200** | URL for direct editing returned |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **files_direct_editing_templates**
> FilesDirectEditingTemplates200Response files_direct_editing_templates(editor_id, creator_id, ocs_api_request)

Get the templates for direct editing

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.files_direct_editing_templates200_response import FilesDirectEditingTemplates200Response
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
    api_instance = nextcloud_client.FilesDirectEditingApi(api_client)
    editor_id = 'editor_id_example' # str | ID of the editor
    creator_id = 'creator_id_example' # str | ID of the creator
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Get the templates for direct editing
        api_response = api_instance.files_direct_editing_templates(editor_id, creator_id, ocs_api_request)
        print("The response of FilesDirectEditingApi->files_direct_editing_templates:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FilesDirectEditingApi->files_direct_editing_templates: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **editor_id** | **str**| ID of the editor | 
 **creator_id** | **str**| ID of the creator | 
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]

### Return type

[**FilesDirectEditingTemplates200Response**](FilesDirectEditingTemplates200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Templates returned |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

