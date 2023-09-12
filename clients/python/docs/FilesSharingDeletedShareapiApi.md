# nextcloud_client.FilesSharingDeletedShareapiApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**files_sharing_deleted_shareapi_list**](FilesSharingDeletedShareapiApi.md#files_sharing_deleted_shareapi_list) | **GET** /ocs/v2.php/apps/files_sharing/api/v1/deletedshares | Get a list of all deleted shares
[**files_sharing_deleted_shareapi_undelete**](FilesSharingDeletedShareapiApi.md#files_sharing_deleted_shareapi_undelete) | **POST** /ocs/v2.php/apps/files_sharing/api/v1/deletedshares/{id} | Undelete a deleted share


# **files_sharing_deleted_shareapi_list**
> FilesSharingDeletedShareapiList200Response files_sharing_deleted_shareapi_list(ocs_api_request)

Get a list of all deleted shares

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.files_sharing_deleted_shareapi_list200_response import FilesSharingDeletedShareapiList200Response
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
    api_instance = nextcloud_client.FilesSharingDeletedShareapiApi(api_client)
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Get a list of all deleted shares
        api_response = api_instance.files_sharing_deleted_shareapi_list(ocs_api_request)
        print("The response of FilesSharingDeletedShareapiApi->files_sharing_deleted_shareapi_list:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FilesSharingDeletedShareapiApi->files_sharing_deleted_shareapi_list: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]

### Return type

[**FilesSharingDeletedShareapiList200Response**](FilesSharingDeletedShareapiList200Response.md)

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

# **files_sharing_deleted_shareapi_undelete**
> CoreWhatsNewDismiss200Response files_sharing_deleted_shareapi_undelete(id, ocs_api_request)

Undelete a deleted share

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
    api_instance = nextcloud_client.FilesSharingDeletedShareapiApi(api_client)
    id = 'id_example' # str | ID of the share
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Undelete a deleted share
        api_response = api_instance.files_sharing_deleted_shareapi_undelete(id, ocs_api_request)
        print("The response of FilesSharingDeletedShareapiApi->files_sharing_deleted_shareapi_undelete:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FilesSharingDeletedShareapiApi->files_sharing_deleted_shareapi_undelete: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **str**| ID of the share | 
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
**200** | Share undeleted successfully |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

