# nextcloud_client.FilesTransferOwnershipApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**files_transfer_ownership_accept**](FilesTransferOwnershipApi.md#files_transfer_ownership_accept) | **POST** /ocs/v2.php/apps/files/api/v1/transferownership/{id} | Accept an ownership transfer
[**files_transfer_ownership_reject**](FilesTransferOwnershipApi.md#files_transfer_ownership_reject) | **DELETE** /ocs/v2.php/apps/files/api/v1/transferownership/{id} | Reject an ownership transfer
[**files_transfer_ownership_transfer**](FilesTransferOwnershipApi.md#files_transfer_ownership_transfer) | **POST** /ocs/v2.php/apps/files/api/v1/transferownership | Transfer the ownership to another user


# **files_transfer_ownership_accept**
> CoreWhatsNewDismiss200Response files_transfer_ownership_accept(id, ocs_api_request)

Accept an ownership transfer

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
    api_instance = nextcloud_client.FilesTransferOwnershipApi(api_client)
    id = 56 # int | ID of the ownership transfer
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Accept an ownership transfer
        api_response = api_instance.files_transfer_ownership_accept(id, ocs_api_request)
        print("The response of FilesTransferOwnershipApi->files_transfer_ownership_accept:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FilesTransferOwnershipApi->files_transfer_ownership_accept: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| ID of the ownership transfer | 
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
**200** | Ownership transfer accepted successfully |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **files_transfer_ownership_reject**
> CoreWhatsNewDismiss200Response files_transfer_ownership_reject(id, ocs_api_request)

Reject an ownership transfer

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
    api_instance = nextcloud_client.FilesTransferOwnershipApi(api_client)
    id = 56 # int | ID of the ownership transfer
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Reject an ownership transfer
        api_response = api_instance.files_transfer_ownership_reject(id, ocs_api_request)
        print("The response of FilesTransferOwnershipApi->files_transfer_ownership_reject:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FilesTransferOwnershipApi->files_transfer_ownership_reject: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| ID of the ownership transfer | 
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
**200** | Ownership transfer rejected successfully |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **files_transfer_ownership_transfer**
> CoreWhatsNewDismiss200Response files_transfer_ownership_transfer(recipient, path, ocs_api_request)

Transfer the ownership to another user

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
    api_instance = nextcloud_client.FilesTransferOwnershipApi(api_client)
    recipient = 'recipient_example' # str | Username of the recipient
    path = 'path_example' # str | Path of the file
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Transfer the ownership to another user
        api_response = api_instance.files_transfer_ownership_transfer(recipient, path, ocs_api_request)
        print("The response of FilesTransferOwnershipApi->files_transfer_ownership_transfer:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FilesTransferOwnershipApi->files_transfer_ownership_transfer: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **recipient** | **str**| Username of the recipient | 
 **path** | **str**| Path of the file | 
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
**200** | Ownership transferred successfully |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

