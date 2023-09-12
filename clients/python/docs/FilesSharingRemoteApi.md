# nextcloud_client.FilesSharingRemoteApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**files_sharing_remote_accept_share**](FilesSharingRemoteApi.md#files_sharing_remote_accept_share) | **POST** /ocs/v2.php/apps/files_sharing/api/v1/remote_shares/pending/{id} | Accept a remote share
[**files_sharing_remote_decline_share**](FilesSharingRemoteApi.md#files_sharing_remote_decline_share) | **DELETE** /ocs/v2.php/apps/files_sharing/api/v1/remote_shares/pending/{id} | Decline a remote share
[**files_sharing_remote_get_open_shares**](FilesSharingRemoteApi.md#files_sharing_remote_get_open_shares) | **GET** /ocs/v2.php/apps/files_sharing/api/v1/remote_shares/pending | Get list of pending remote shares
[**files_sharing_remote_get_share**](FilesSharingRemoteApi.md#files_sharing_remote_get_share) | **GET** /ocs/v2.php/apps/files_sharing/api/v1/remote_shares/{id} | Get info of a remote share
[**files_sharing_remote_get_shares**](FilesSharingRemoteApi.md#files_sharing_remote_get_shares) | **GET** /ocs/v2.php/apps/files_sharing/api/v1/remote_shares | Get a list of accepted remote shares
[**files_sharing_remote_unshare**](FilesSharingRemoteApi.md#files_sharing_remote_unshare) | **DELETE** /ocs/v2.php/apps/files_sharing/api/v1/remote_shares/{id} | Unshare a remote share


# **files_sharing_remote_accept_share**
> CoreWhatsNewDismiss200Response files_sharing_remote_accept_share(id, ocs_api_request)

Accept a remote share

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
    api_instance = nextcloud_client.FilesSharingRemoteApi(api_client)
    id = 56 # int | ID of the share
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Accept a remote share
        api_response = api_instance.files_sharing_remote_accept_share(id, ocs_api_request)
        print("The response of FilesSharingRemoteApi->files_sharing_remote_accept_share:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FilesSharingRemoteApi->files_sharing_remote_accept_share: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| ID of the share | 
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
**200** | Share accepted successfully |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **files_sharing_remote_decline_share**
> CoreWhatsNewDismiss200Response files_sharing_remote_decline_share(id, ocs_api_request)

Decline a remote share

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
    api_instance = nextcloud_client.FilesSharingRemoteApi(api_client)
    id = 56 # int | ID of the share
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Decline a remote share
        api_response = api_instance.files_sharing_remote_decline_share(id, ocs_api_request)
        print("The response of FilesSharingRemoteApi->files_sharing_remote_decline_share:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FilesSharingRemoteApi->files_sharing_remote_decline_share: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| ID of the share | 
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
**200** | Share declined successfully |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **files_sharing_remote_get_open_shares**
> FilesSharingRemoteGetShares200Response files_sharing_remote_get_open_shares(ocs_api_request)

Get list of pending remote shares

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.files_sharing_remote_get_shares200_response import FilesSharingRemoteGetShares200Response
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
    api_instance = nextcloud_client.FilesSharingRemoteApi(api_client)
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Get list of pending remote shares
        api_response = api_instance.files_sharing_remote_get_open_shares(ocs_api_request)
        print("The response of FilesSharingRemoteApi->files_sharing_remote_get_open_shares:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FilesSharingRemoteApi->files_sharing_remote_get_open_shares: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]

### Return type

[**FilesSharingRemoteGetShares200Response**](FilesSharingRemoteGetShares200Response.md)

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

# **files_sharing_remote_get_share**
> FilesSharingRemoteGetShare200Response files_sharing_remote_get_share(id, ocs_api_request)

Get info of a remote share

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.files_sharing_remote_get_share200_response import FilesSharingRemoteGetShare200Response
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
    api_instance = nextcloud_client.FilesSharingRemoteApi(api_client)
    id = 56 # int | ID of the share
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Get info of a remote share
        api_response = api_instance.files_sharing_remote_get_share(id, ocs_api_request)
        print("The response of FilesSharingRemoteApi->files_sharing_remote_get_share:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FilesSharingRemoteApi->files_sharing_remote_get_share: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| ID of the share | 
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]

### Return type

[**FilesSharingRemoteGetShare200Response**](FilesSharingRemoteGetShare200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Share returned |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **files_sharing_remote_get_shares**
> FilesSharingRemoteGetShares200Response files_sharing_remote_get_shares(ocs_api_request)

Get a list of accepted remote shares

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.files_sharing_remote_get_shares200_response import FilesSharingRemoteGetShares200Response
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
    api_instance = nextcloud_client.FilesSharingRemoteApi(api_client)
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Get a list of accepted remote shares
        api_response = api_instance.files_sharing_remote_get_shares(ocs_api_request)
        print("The response of FilesSharingRemoteApi->files_sharing_remote_get_shares:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FilesSharingRemoteApi->files_sharing_remote_get_shares: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]

### Return type

[**FilesSharingRemoteGetShares200Response**](FilesSharingRemoteGetShares200Response.md)

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

# **files_sharing_remote_unshare**
> CoreWhatsNewDismiss200Response files_sharing_remote_unshare(id, ocs_api_request)

Unshare a remote share

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
    api_instance = nextcloud_client.FilesSharingRemoteApi(api_client)
    id = 56 # int | ID of the share
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Unshare a remote share
        api_response = api_instance.files_sharing_remote_unshare(id, ocs_api_request)
        print("The response of FilesSharingRemoteApi->files_sharing_remote_unshare:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FilesSharingRemoteApi->files_sharing_remote_unshare: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**| ID of the share | 
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
**200** | Share unshared successfully |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

