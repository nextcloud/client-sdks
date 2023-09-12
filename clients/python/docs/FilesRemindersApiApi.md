# nextcloud_client.FilesRemindersApiApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**files_reminders_api_get**](FilesRemindersApiApi.md#files_reminders_api_get) | **GET** /ocs/v2.php/apps/files_reminders/api/v{version}/{fileId} | Get a reminder
[**files_reminders_api_remove**](FilesRemindersApiApi.md#files_reminders_api_remove) | **DELETE** /ocs/v2.php/apps/files_reminders/api/v{version}/{fileId} | Remove a reminder
[**files_reminders_api_set**](FilesRemindersApiApi.md#files_reminders_api_set) | **PUT** /ocs/v2.php/apps/files_reminders/api/v{version}/{fileId} | Set a reminder


# **files_reminders_api_get**
> FilesRemindersApiGet200Response files_reminders_api_get(version, file_id, ocs_api_request)

Get a reminder

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.files_reminders_api_get200_response import FilesRemindersApiGet200Response
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
    api_instance = nextcloud_client.FilesRemindersApiApi(api_client)
    version = 'version_example' # str | 
    file_id = 56 # int | ID of the file
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Get a reminder
        api_response = api_instance.files_reminders_api_get(version, file_id, ocs_api_request)
        print("The response of FilesRemindersApiApi->files_reminders_api_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FilesRemindersApiApi->files_reminders_api_get: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **str**|  | 
 **file_id** | **int**| ID of the file | 
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]

### Return type

[**FilesRemindersApiGet200Response**](FilesRemindersApiGet200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Reminder returned |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **files_reminders_api_remove**
> CoreWhatsNewDismiss200Response files_reminders_api_remove(version, file_id, ocs_api_request)

Remove a reminder

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.core_whats_new_dismiss200_response import CoreWhatsNewDismiss200Response
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
    api_instance = nextcloud_client.FilesRemindersApiApi(api_client)
    version = 'version_example' # str | 
    file_id = 56 # int | ID of the file
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Remove a reminder
        api_response = api_instance.files_reminders_api_remove(version, file_id, ocs_api_request)
        print("The response of FilesRemindersApiApi->files_reminders_api_remove:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FilesRemindersApiApi->files_reminders_api_remove: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **str**|  | 
 **file_id** | **int**| ID of the file | 
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]

### Return type

[**CoreWhatsNewDismiss200Response**](CoreWhatsNewDismiss200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Reminder deleted successfully |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **files_reminders_api_set**
> CoreWhatsNewDismiss200Response files_reminders_api_set(due_date, version, file_id, ocs_api_request)

Set a reminder

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.core_whats_new_dismiss200_response import CoreWhatsNewDismiss200Response
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
    api_instance = nextcloud_client.FilesRemindersApiApi(api_client)
    due_date = 'due_date_example' # str | ISO 8601 formatted date time string
    version = 'version_example' # str | 
    file_id = 56 # int | ID of the file
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Set a reminder
        api_response = api_instance.files_reminders_api_set(due_date, version, file_id, ocs_api_request)
        print("The response of FilesRemindersApiApi->files_reminders_api_set:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FilesRemindersApiApi->files_reminders_api_set: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **due_date** | **str**| ISO 8601 formatted date time string | 
 **version** | **str**|  | 
 **file_id** | **int**| ID of the file | 
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]

### Return type

[**CoreWhatsNewDismiss200Response**](CoreWhatsNewDismiss200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Reminder updated |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

