# nextcloud_client.FilesSharingShareapiApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**files_sharing_shareapi_accept_share**](FilesSharingShareapiApi.md#files_sharing_shareapi_accept_share) | **POST** /ocs/v2.php/apps/files_sharing/api/v1/shares/pending/{id} | Accept a share
[**files_sharing_shareapi_create_share**](FilesSharingShareapiApi.md#files_sharing_shareapi_create_share) | **POST** /ocs/v2.php/apps/files_sharing/api/v1/shares | Create a share
[**files_sharing_shareapi_delete_share**](FilesSharingShareapiApi.md#files_sharing_shareapi_delete_share) | **DELETE** /ocs/v2.php/apps/files_sharing/api/v1/shares/{id} | Delete a share
[**files_sharing_shareapi_get_inherited_shares**](FilesSharingShareapiApi.md#files_sharing_shareapi_get_inherited_shares) | **GET** /ocs/v2.php/apps/files_sharing/api/v1/shares/inherited | Get all shares relative to a file, including parent folders shares rights
[**files_sharing_shareapi_get_share**](FilesSharingShareapiApi.md#files_sharing_shareapi_get_share) | **GET** /ocs/v2.php/apps/files_sharing/api/v1/shares/{id} | Get a specific share by id
[**files_sharing_shareapi_get_shares**](FilesSharingShareapiApi.md#files_sharing_shareapi_get_shares) | **GET** /ocs/v2.php/apps/files_sharing/api/v1/shares | Get shares of the current user
[**files_sharing_shareapi_pending_shares**](FilesSharingShareapiApi.md#files_sharing_shareapi_pending_shares) | **GET** /ocs/v2.php/apps/files_sharing/api/v1/shares/pending | Get all shares that are still pending
[**files_sharing_shareapi_update_share**](FilesSharingShareapiApi.md#files_sharing_shareapi_update_share) | **PUT** /ocs/v2.php/apps/files_sharing/api/v1/shares/{id} | Update a share


# **files_sharing_shareapi_accept_share**
> CoreWhatsNewDismiss200Response files_sharing_shareapi_accept_share(id, ocs_api_request)

Accept a share

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
    api_instance = nextcloud_client.FilesSharingShareapiApi(api_client)
    id = 'id_example' # str | ID of the share
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Accept a share
        api_response = api_instance.files_sharing_shareapi_accept_share(id, ocs_api_request)
        print("The response of FilesSharingShareapiApi->files_sharing_shareapi_accept_share:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FilesSharingShareapiApi->files_sharing_shareapi_accept_share: %s\n" % e)
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
**200** | Share accepted successfully |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **files_sharing_shareapi_create_share**
> FilesSharingShareapiCreateShare200Response files_sharing_shareapi_create_share(ocs_api_request, path=path, permissions=permissions, share_type=share_type, share_with=share_with, public_upload=public_upload, password=password, send_password_by_talk=send_password_by_talk, expire_date=expire_date, note=note, label=label, attributes=attributes)

Create a share

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.files_sharing_shareapi_create_share200_response import FilesSharingShareapiCreateShare200Response
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
    api_instance = nextcloud_client.FilesSharingShareapiApi(api_client)
    ocs_api_request = 'true' # str |  (default to 'true')
    path = 'path_example' # str | Path of the share (optional)
    permissions = 56 # int | Permissions for the share (optional)
    share_type = -1 # int | Type of the share (optional) (default to -1)
    share_with = 'share_with_example' # str | The entity this should be shared with (optional)
    public_upload = 'false' # str | If public uploading is allowed (optional) (default to 'false')
    password = '' # str | Password for the share (optional) (default to '')
    send_password_by_talk = 'send_password_by_talk_example' # str | Send the password for the share over Talk (optional)
    expire_date = '' # str | Expiry date of the share (optional) (default to '')
    note = '' # str | Note for the share (optional) (default to '')
    label = '' # str | Label for the share (only used in link and email) (optional) (default to '')
    attributes = 'attributes_example' # str | Additional attributes for the share (optional)

    try:
        # Create a share
        api_response = api_instance.files_sharing_shareapi_create_share(ocs_api_request, path=path, permissions=permissions, share_type=share_type, share_with=share_with, public_upload=public_upload, password=password, send_password_by_talk=send_password_by_talk, expire_date=expire_date, note=note, label=label, attributes=attributes)
        print("The response of FilesSharingShareapiApi->files_sharing_shareapi_create_share:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FilesSharingShareapiApi->files_sharing_shareapi_create_share: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]
 **path** | **str**| Path of the share | [optional] 
 **permissions** | **int**| Permissions for the share | [optional] 
 **share_type** | **int**| Type of the share | [optional] [default to -1]
 **share_with** | **str**| The entity this should be shared with | [optional] 
 **public_upload** | **str**| If public uploading is allowed | [optional] [default to &#39;false&#39;]
 **password** | **str**| Password for the share | [optional] [default to &#39;&#39;]
 **send_password_by_talk** | **str**| Send the password for the share over Talk | [optional] 
 **expire_date** | **str**| Expiry date of the share | [optional] [default to &#39;&#39;]
 **note** | **str**| Note for the share | [optional] [default to &#39;&#39;]
 **label** | **str**| Label for the share (only used in link and email) | [optional] [default to &#39;&#39;]
 **attributes** | **str**| Additional attributes for the share | [optional] 

### Return type

[**FilesSharingShareapiCreateShare200Response**](FilesSharingShareapiCreateShare200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Share created |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **files_sharing_shareapi_delete_share**
> CoreWhatsNewDismiss200Response files_sharing_shareapi_delete_share(id, ocs_api_request)

Delete a share

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
    api_instance = nextcloud_client.FilesSharingShareapiApi(api_client)
    id = 'id_example' # str | ID of the share
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Delete a share
        api_response = api_instance.files_sharing_shareapi_delete_share(id, ocs_api_request)
        print("The response of FilesSharingShareapiApi->files_sharing_shareapi_delete_share:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FilesSharingShareapiApi->files_sharing_shareapi_delete_share: %s\n" % e)
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
**200** | Share deleted successfully |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **files_sharing_shareapi_get_inherited_shares**
> FilesSharingShareapiGetShares200Response files_sharing_shareapi_get_inherited_shares(path, ocs_api_request)

Get all shares relative to a file, including parent folders shares rights

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.files_sharing_shareapi_get_shares200_response import FilesSharingShareapiGetShares200Response
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
    api_instance = nextcloud_client.FilesSharingShareapiApi(api_client)
    path = 'path_example' # str | Path all shares will be relative to
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Get all shares relative to a file, including parent folders shares rights
        api_response = api_instance.files_sharing_shareapi_get_inherited_shares(path, ocs_api_request)
        print("The response of FilesSharingShareapiApi->files_sharing_shareapi_get_inherited_shares:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FilesSharingShareapiApi->files_sharing_shareapi_get_inherited_shares: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **str**| Path all shares will be relative to | 
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]

### Return type

[**FilesSharingShareapiGetShares200Response**](FilesSharingShareapiGetShares200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Shares returned |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **files_sharing_shareapi_get_share**
> FilesSharingShareapiCreateShare200Response files_sharing_shareapi_get_share(id, ocs_api_request, include_tags=include_tags)

Get a specific share by id

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.files_sharing_shareapi_create_share200_response import FilesSharingShareapiCreateShare200Response
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
    api_instance = nextcloud_client.FilesSharingShareapiApi(api_client)
    id = 'id_example' # str | ID of the share
    ocs_api_request = 'true' # str |  (default to 'true')
    include_tags = 0 # int | Include tags in the share (optional) (default to 0)

    try:
        # Get a specific share by id
        api_response = api_instance.files_sharing_shareapi_get_share(id, ocs_api_request, include_tags=include_tags)
        print("The response of FilesSharingShareapiApi->files_sharing_shareapi_get_share:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FilesSharingShareapiApi->files_sharing_shareapi_get_share: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **str**| ID of the share | 
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]
 **include_tags** | **int**| Include tags in the share | [optional] [default to 0]

### Return type

[**FilesSharingShareapiCreateShare200Response**](FilesSharingShareapiCreateShare200Response.md)

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

# **files_sharing_shareapi_get_shares**
> FilesSharingShareapiGetShares200Response files_sharing_shareapi_get_shares(ocs_api_request, shared_with_me=shared_with_me, reshares=reshares, subfiles=subfiles, path=path, include_tags=include_tags)

Get shares of the current user

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.files_sharing_shareapi_get_shares200_response import FilesSharingShareapiGetShares200Response
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
    api_instance = nextcloud_client.FilesSharingShareapiApi(api_client)
    ocs_api_request = 'true' # str |  (default to 'true')
    shared_with_me = 'false' # str | Only get shares with the current user (optional) (default to 'false')
    reshares = 'false' # str | Only get shares by the current user and reshares (optional) (default to 'false')
    subfiles = 'false' # str | Only get all shares in a folder (optional) (default to 'false')
    path = '' # str | Get shares for a specific path (optional) (default to '')
    include_tags = 'false' # str | Include tags in the share (optional) (default to 'false')

    try:
        # Get shares of the current user
        api_response = api_instance.files_sharing_shareapi_get_shares(ocs_api_request, shared_with_me=shared_with_me, reshares=reshares, subfiles=subfiles, path=path, include_tags=include_tags)
        print("The response of FilesSharingShareapiApi->files_sharing_shareapi_get_shares:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FilesSharingShareapiApi->files_sharing_shareapi_get_shares: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]
 **shared_with_me** | **str**| Only get shares with the current user | [optional] [default to &#39;false&#39;]
 **reshares** | **str**| Only get shares by the current user and reshares | [optional] [default to &#39;false&#39;]
 **subfiles** | **str**| Only get all shares in a folder | [optional] [default to &#39;false&#39;]
 **path** | **str**| Get shares for a specific path | [optional] [default to &#39;&#39;]
 **include_tags** | **str**| Include tags in the share | [optional] [default to &#39;false&#39;]

### Return type

[**FilesSharingShareapiGetShares200Response**](FilesSharingShareapiGetShares200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Shares returned |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **files_sharing_shareapi_pending_shares**
> FilesSharingShareapiGetShares200Response files_sharing_shareapi_pending_shares(ocs_api_request)

Get all shares that are still pending

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.files_sharing_shareapi_get_shares200_response import FilesSharingShareapiGetShares200Response
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
    api_instance = nextcloud_client.FilesSharingShareapiApi(api_client)
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Get all shares that are still pending
        api_response = api_instance.files_sharing_shareapi_pending_shares(ocs_api_request)
        print("The response of FilesSharingShareapiApi->files_sharing_shareapi_pending_shares:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FilesSharingShareapiApi->files_sharing_shareapi_pending_shares: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]

### Return type

[**FilesSharingShareapiGetShares200Response**](FilesSharingShareapiGetShares200Response.md)

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

# **files_sharing_shareapi_update_share**
> FilesSharingShareapiCreateShare200Response files_sharing_shareapi_update_share(id, ocs_api_request, permissions=permissions, password=password, send_password_by_talk=send_password_by_talk, public_upload=public_upload, expire_date=expire_date, note=note, label=label, hide_download=hide_download, attributes=attributes)

Update a share

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.files_sharing_shareapi_create_share200_response import FilesSharingShareapiCreateShare200Response
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
    api_instance = nextcloud_client.FilesSharingShareapiApi(api_client)
    id = 'id_example' # str | ID of the share
    ocs_api_request = 'true' # str |  (default to 'true')
    permissions = 56 # int | New permissions (optional)
    password = 'password_example' # str | New password (optional)
    send_password_by_talk = 'send_password_by_talk_example' # str | New condition if the password should be send over Talk (optional)
    public_upload = 'public_upload_example' # str | New condition if public uploading is allowed (optional)
    expire_date = 'expire_date_example' # str | New expiry date (optional)
    note = 'note_example' # str | New note (optional)
    label = 'label_example' # str | New label (optional)
    hide_download = 'hide_download_example' # str | New condition if the download should be hidden (optional)
    attributes = 'attributes_example' # str | New additional attributes (optional)

    try:
        # Update a share
        api_response = api_instance.files_sharing_shareapi_update_share(id, ocs_api_request, permissions=permissions, password=password, send_password_by_talk=send_password_by_talk, public_upload=public_upload, expire_date=expire_date, note=note, label=label, hide_download=hide_download, attributes=attributes)
        print("The response of FilesSharingShareapiApi->files_sharing_shareapi_update_share:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FilesSharingShareapiApi->files_sharing_shareapi_update_share: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **str**| ID of the share | 
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]
 **permissions** | **int**| New permissions | [optional] 
 **password** | **str**| New password | [optional] 
 **send_password_by_talk** | **str**| New condition if the password should be send over Talk | [optional] 
 **public_upload** | **str**| New condition if public uploading is allowed | [optional] 
 **expire_date** | **str**| New expiry date | [optional] 
 **note** | **str**| New note | [optional] 
 **label** | **str**| New label | [optional] 
 **hide_download** | **str**| New condition if the download should be hidden | [optional] 
 **attributes** | **str**| New additional attributes | [optional] 

### Return type

[**FilesSharingShareapiCreateShare200Response**](FilesSharingShareapiCreateShare200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Share updated successfully |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

