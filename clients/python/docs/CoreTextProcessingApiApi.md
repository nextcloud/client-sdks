# nextcloud_client.CoreTextProcessingApiApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**core_text_processing_api_delete_task**](CoreTextProcessingApiApi.md#core_text_processing_api_delete_task) | **DELETE** /ocs/v2.php/textprocessing/task/{id} | This endpoint allows to delete a scheduled task for a user
[**core_text_processing_api_get_task**](CoreTextProcessingApiApi.md#core_text_processing_api_get_task) | **GET** /ocs/v2.php/textprocessing/task/{id} | This endpoint allows checking the status and results of a task. Tasks are removed 1 week after receiving their last update.
[**core_text_processing_api_list_tasks_by_app**](CoreTextProcessingApiApi.md#core_text_processing_api_list_tasks_by_app) | **GET** /ocs/v2.php/textprocessing/tasks/app/{appId} | This endpoint returns a list of tasks of a user that are related with a specific appId and optionally with an identifier
[**core_text_processing_api_schedule**](CoreTextProcessingApiApi.md#core_text_processing_api_schedule) | **POST** /ocs/v2.php/textprocessing/schedule | This endpoint allows scheduling a language model task
[**core_text_processing_api_task_types**](CoreTextProcessingApiApi.md#core_text_processing_api_task_types) | **GET** /ocs/v2.php/textprocessing/tasktypes | This endpoint returns all available LanguageModel task types


# **core_text_processing_api_delete_task**
> CoreTextProcessingApiSchedule200Response core_text_processing_api_delete_task(id, ocs_api_request)

This endpoint allows to delete a scheduled task for a user

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.core_text_processing_api_schedule200_response import CoreTextProcessingApiSchedule200Response
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
    api_instance = nextcloud_client.CoreTextProcessingApiApi(api_client)
    id = 56 # int | The id of the task
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # This endpoint allows to delete a scheduled task for a user
        api_response = api_instance.core_text_processing_api_delete_task(id, ocs_api_request)
        print("The response of CoreTextProcessingApiApi->core_text_processing_api_delete_task:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CoreTextProcessingApiApi->core_text_processing_api_delete_task: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| The id of the task | 
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]

### Return type

[**CoreTextProcessingApiSchedule200Response**](CoreTextProcessingApiSchedule200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Task returned |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **core_text_processing_api_get_task**
> CoreTextProcessingApiSchedule200Response core_text_processing_api_get_task(id, ocs_api_request)

This endpoint allows checking the status and results of a task. Tasks are removed 1 week after receiving their last update.

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.core_text_processing_api_schedule200_response import CoreTextProcessingApiSchedule200Response
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
    api_instance = nextcloud_client.CoreTextProcessingApiApi(api_client)
    id = 56 # int | The id of the task
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # This endpoint allows checking the status and results of a task. Tasks are removed 1 week after receiving their last update.
        api_response = api_instance.core_text_processing_api_get_task(id, ocs_api_request)
        print("The response of CoreTextProcessingApiApi->core_text_processing_api_get_task:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CoreTextProcessingApiApi->core_text_processing_api_get_task: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| The id of the task | 
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]

### Return type

[**CoreTextProcessingApiSchedule200Response**](CoreTextProcessingApiSchedule200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Task returned |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **core_text_processing_api_list_tasks_by_app**
> CoreTextProcessingApiListTasksByApp200Response core_text_processing_api_list_tasks_by_app(app_id, ocs_api_request, identifier=identifier)

This endpoint returns a list of tasks of a user that are related with a specific appId and optionally with an identifier

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.core_text_processing_api_list_tasks_by_app200_response import CoreTextProcessingApiListTasksByApp200Response
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
    api_instance = nextcloud_client.CoreTextProcessingApiApi(api_client)
    app_id = 'app_id_example' # str | ID of the app
    ocs_api_request = 'true' # str |  (default to 'true')
    identifier = 'identifier_example' # str | An arbitrary identifier for the task (optional)

    try:
        # This endpoint returns a list of tasks of a user that are related with a specific appId and optionally with an identifier
        api_response = api_instance.core_text_processing_api_list_tasks_by_app(app_id, ocs_api_request, identifier=identifier)
        print("The response of CoreTextProcessingApiApi->core_text_processing_api_list_tasks_by_app:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CoreTextProcessingApiApi->core_text_processing_api_list_tasks_by_app: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app_id** | **str**| ID of the app | 
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]
 **identifier** | **str**| An arbitrary identifier for the task | [optional] 

### Return type

[**CoreTextProcessingApiListTasksByApp200Response**](CoreTextProcessingApiListTasksByApp200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Task list returned |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **core_text_processing_api_schedule**
> CoreTextProcessingApiSchedule200Response core_text_processing_api_schedule(input, type, app_id, ocs_api_request, identifier=identifier)

This endpoint allows scheduling a language model task

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.core_text_processing_api_schedule200_response import CoreTextProcessingApiSchedule200Response
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
    api_instance = nextcloud_client.CoreTextProcessingApiApi(api_client)
    input = 'input_example' # str | Input text
    type = 'type_example' # str | Type of the task
    app_id = 'app_id_example' # str | ID of the app that will execute the task
    ocs_api_request = 'true' # str |  (default to 'true')
    identifier = '' # str | An arbitrary identifier for the task (optional) (default to '')

    try:
        # This endpoint allows scheduling a language model task
        api_response = api_instance.core_text_processing_api_schedule(input, type, app_id, ocs_api_request, identifier=identifier)
        print("The response of CoreTextProcessingApiApi->core_text_processing_api_schedule:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CoreTextProcessingApiApi->core_text_processing_api_schedule: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **input** | **str**| Input text | 
 **type** | **str**| Type of the task | 
 **app_id** | **str**| ID of the app that will execute the task | 
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]
 **identifier** | **str**| An arbitrary identifier for the task | [optional] [default to &#39;&#39;]

### Return type

[**CoreTextProcessingApiSchedule200Response**](CoreTextProcessingApiSchedule200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Task scheduled successfully |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **core_text_processing_api_task_types**
> CoreTextProcessingApiTaskTypes200Response core_text_processing_api_task_types(ocs_api_request)

This endpoint returns all available LanguageModel task types

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.core_text_processing_api_task_types200_response import CoreTextProcessingApiTaskTypes200Response
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
    api_instance = nextcloud_client.CoreTextProcessingApiApi(api_client)
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # This endpoint returns all available LanguageModel task types
        api_response = api_instance.core_text_processing_api_task_types(ocs_api_request)
        print("The response of CoreTextProcessingApiApi->core_text_processing_api_task_types:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CoreTextProcessingApiApi->core_text_processing_api_task_types: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]

### Return type

[**CoreTextProcessingApiTaskTypes200Response**](CoreTextProcessingApiTaskTypes200Response.md)

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

