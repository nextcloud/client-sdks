# nextcloud_client.FilesTemplateApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**files_template_create**](FilesTemplateApi.md#files_template_create) | **POST** /ocs/v2.php/apps/files/api/v1/templates/create | Create a template
[**files_template_list**](FilesTemplateApi.md#files_template_list) | **GET** /ocs/v2.php/apps/files/api/v1/templates | List the available templates
[**files_template_path**](FilesTemplateApi.md#files_template_path) | **POST** /ocs/v2.php/apps/files/api/v1/templates/path | Initialize the template directory


# **files_template_create**
> FilesTemplateCreate200Response files_template_create(file_path, ocs_api_request, template_path=template_path, template_type=template_type)

Create a template

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.files_template_create200_response import FilesTemplateCreate200Response
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
    api_instance = nextcloud_client.FilesTemplateApi(api_client)
    file_path = 'file_path_example' # str | Path of the file
    ocs_api_request = 'true' # str |  (default to 'true')
    template_path = '' # str | Name of the template (optional) (default to '')
    template_type = 'user' # str | Type of the template (optional) (default to 'user')

    try:
        # Create a template
        api_response = api_instance.files_template_create(file_path, ocs_api_request, template_path=template_path, template_type=template_type)
        print("The response of FilesTemplateApi->files_template_create:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FilesTemplateApi->files_template_create: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file_path** | **str**| Path of the file | 
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]
 **template_path** | **str**| Name of the template | [optional] [default to &#39;&#39;]
 **template_type** | **str**| Type of the template | [optional] [default to &#39;user&#39;]

### Return type

[**FilesTemplateCreate200Response**](FilesTemplateCreate200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Template created successfully |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **files_template_list**
> FilesTemplateList200Response files_template_list(ocs_api_request)

List the available templates

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.files_template_list200_response import FilesTemplateList200Response
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
    api_instance = nextcloud_client.FilesTemplateApi(api_client)
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # List the available templates
        api_response = api_instance.files_template_list(ocs_api_request)
        print("The response of FilesTemplateApi->files_template_list:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FilesTemplateApi->files_template_list: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]

### Return type

[**FilesTemplateList200Response**](FilesTemplateList200Response.md)

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

# **files_template_path**
> FilesTemplatePath200Response files_template_path(ocs_api_request, template_path=template_path, copy_system_templates=copy_system_templates)

Initialize the template directory

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.files_template_path200_response import FilesTemplatePath200Response
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
    api_instance = nextcloud_client.FilesTemplateApi(api_client)
    ocs_api_request = 'true' # str |  (default to 'true')
    template_path = '' # str | Path of the template directory (optional) (default to '')
    copy_system_templates = 0 # int | Whether to copy the system templates to the template directory (optional) (default to 0)

    try:
        # Initialize the template directory
        api_response = api_instance.files_template_path(ocs_api_request, template_path=template_path, copy_system_templates=copy_system_templates)
        print("The response of FilesTemplateApi->files_template_path:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FilesTemplateApi->files_template_path: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]
 **template_path** | **str**| Path of the template directory | [optional] [default to &#39;&#39;]
 **copy_system_templates** | **int**| Whether to copy the system templates to the template directory | [optional] [default to 0]

### Return type

[**FilesTemplatePath200Response**](FilesTemplatePath200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Template directory initialized successfully |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

