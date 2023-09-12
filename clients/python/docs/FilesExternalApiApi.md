# nextcloud_client.FilesExternalApiApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**files_external_api_get_user_mounts**](FilesExternalApiApi.md#files_external_api_get_user_mounts) | **GET** /ocs/v2.php/apps/files_external/api/v1/mounts | Get the mount points visible for this user


# **files_external_api_get_user_mounts**
> FilesExternalApiGetUserMounts200Response files_external_api_get_user_mounts(ocs_api_request)

Get the mount points visible for this user

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.files_external_api_get_user_mounts200_response import FilesExternalApiGetUserMounts200Response
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
    api_instance = nextcloud_client.FilesExternalApiApi(api_client)
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Get the mount points visible for this user
        api_response = api_instance.files_external_api_get_user_mounts(ocs_api_request)
        print("The response of FilesExternalApiApi->files_external_api_get_user_mounts:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FilesExternalApiApi->files_external_api_get_user_mounts: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]

### Return type

[**FilesExternalApiGetUserMounts200Response**](FilesExternalApiGetUserMounts200Response.md)

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
