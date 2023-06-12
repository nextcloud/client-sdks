# openapi_client.ProvisioningApiGroupsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**provisioning_api_groups_add_group**](ProvisioningApiGroupsApi.md#provisioning_api_groups_add_group) | **POST** /ocs/v2.php/cloud/groups | Create a new group
[**provisioning_api_groups_delete_group**](ProvisioningApiGroupsApi.md#provisioning_api_groups_delete_group) | **DELETE** /ocs/v2.php/cloud/groups/{groupId} | Delete a group
[**provisioning_api_groups_get_group**](ProvisioningApiGroupsApi.md#provisioning_api_groups_get_group) | **GET** /ocs/v2.php/cloud/groups/{groupId} | Get a list of users in the specified group
[**provisioning_api_groups_get_group_users**](ProvisioningApiGroupsApi.md#provisioning_api_groups_get_group_users) | **GET** /ocs/v2.php/cloud/groups/{groupId}/users | Get a list of users in the specified group
[**provisioning_api_groups_get_group_users_details**](ProvisioningApiGroupsApi.md#provisioning_api_groups_get_group_users_details) | **GET** /ocs/v2.php/cloud/groups/{groupId}/users/details | Get a list of users details in the specified group
[**provisioning_api_groups_get_groups**](ProvisioningApiGroupsApi.md#provisioning_api_groups_get_groups) | **GET** /ocs/v2.php/cloud/groups | Get a list of groups
[**provisioning_api_groups_get_groups_details**](ProvisioningApiGroupsApi.md#provisioning_api_groups_get_groups_details) | **GET** /ocs/v2.php/cloud/groups/details | Get a list of groups details
[**provisioning_api_groups_get_sub_admins_of_group**](ProvisioningApiGroupsApi.md#provisioning_api_groups_get_sub_admins_of_group) | **GET** /ocs/v2.php/cloud/groups/{groupId}/subadmins | Get the list of user IDs that are a subadmin of the group
[**provisioning_api_groups_update_group**](ProvisioningApiGroupsApi.md#provisioning_api_groups_update_group) | **PUT** /ocs/v2.php/cloud/groups/{groupId} | Update a group


# **provisioning_api_groups_add_group**
> CoreWhatsNewDismiss200Response provisioning_api_groups_add_group(groupid, ocs_api_request, displayname=displayname)

Create a new group

This endpoint requires admin access

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import openapi_client
from openapi_client.models.core_whats_new_dismiss200_response import CoreWhatsNewDismiss200Response
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "http://localhost"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure HTTP basic authorization: basic_auth
configuration = openapi_client.Configuration(
    username = os.environ["USERNAME"],
    password = os.environ["PASSWORD"]
)

# Configure Bearer authorization: bearer_auth
configuration = openapi_client.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.ProvisioningApiGroupsApi(api_client)
    groupid = 'groupid_example' # str | ID of the group
    ocs_api_request = 'true' # str |  (default to 'true')
    displayname = '' # str | Display name of the group (optional) (default to '')

    try:
        # Create a new group
        api_response = api_instance.provisioning_api_groups_add_group(groupid, ocs_api_request, displayname=displayname)
        print("The response of ProvisioningApiGroupsApi->provisioning_api_groups_add_group:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ProvisioningApiGroupsApi->provisioning_api_groups_add_group: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupid** | **str**| ID of the group | 
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]
 **displayname** | **str**| Display name of the group | [optional] [default to &#39;&#39;]

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
**200** |  |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **provisioning_api_groups_delete_group**
> CoreWhatsNewDismiss200Response provisioning_api_groups_delete_group(group_id, ocs_api_request)

Delete a group

This endpoint requires admin access

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import openapi_client
from openapi_client.models.core_whats_new_dismiss200_response import CoreWhatsNewDismiss200Response
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "http://localhost"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure HTTP basic authorization: basic_auth
configuration = openapi_client.Configuration(
    username = os.environ["USERNAME"],
    password = os.environ["PASSWORD"]
)

# Configure Bearer authorization: bearer_auth
configuration = openapi_client.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.ProvisioningApiGroupsApi(api_client)
    group_id = 'group_id_example' # str | ID of the group
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Delete a group
        api_response = api_instance.provisioning_api_groups_delete_group(group_id, ocs_api_request)
        print("The response of ProvisioningApiGroupsApi->provisioning_api_groups_delete_group:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ProvisioningApiGroupsApi->provisioning_api_groups_delete_group: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group_id** | **str**| ID of the group | 
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
**200** |  |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **provisioning_api_groups_get_group**
> ProvisioningApiGroupsGetGroupUsers200Response provisioning_api_groups_get_group(group_id, ocs_api_request)

Get a list of users in the specified group

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import openapi_client
from openapi_client.models.provisioning_api_groups_get_group_users200_response import ProvisioningApiGroupsGetGroupUsers200Response
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "http://localhost"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure HTTP basic authorization: basic_auth
configuration = openapi_client.Configuration(
    username = os.environ["USERNAME"],
    password = os.environ["PASSWORD"]
)

# Configure Bearer authorization: bearer_auth
configuration = openapi_client.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.ProvisioningApiGroupsApi(api_client)
    group_id = 'group_id_example' # str | ID of the group
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Get a list of users in the specified group
        api_response = api_instance.provisioning_api_groups_get_group(group_id, ocs_api_request)
        print("The response of ProvisioningApiGroupsApi->provisioning_api_groups_get_group:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ProvisioningApiGroupsApi->provisioning_api_groups_get_group: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group_id** | **str**| ID of the group | 
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]

### Return type

[**ProvisioningApiGroupsGetGroupUsers200Response**](ProvisioningApiGroupsGetGroupUsers200Response.md)

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

# **provisioning_api_groups_get_group_users**
> ProvisioningApiGroupsGetGroupUsers200Response provisioning_api_groups_get_group_users(group_id, ocs_api_request)

Get a list of users in the specified group

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import openapi_client
from openapi_client.models.provisioning_api_groups_get_group_users200_response import ProvisioningApiGroupsGetGroupUsers200Response
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "http://localhost"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure HTTP basic authorization: basic_auth
configuration = openapi_client.Configuration(
    username = os.environ["USERNAME"],
    password = os.environ["PASSWORD"]
)

# Configure Bearer authorization: bearer_auth
configuration = openapi_client.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.ProvisioningApiGroupsApi(api_client)
    group_id = 'group_id_example' # str | ID of the group
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Get a list of users in the specified group
        api_response = api_instance.provisioning_api_groups_get_group_users(group_id, ocs_api_request)
        print("The response of ProvisioningApiGroupsApi->provisioning_api_groups_get_group_users:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ProvisioningApiGroupsApi->provisioning_api_groups_get_group_users: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group_id** | **str**| ID of the group | 
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]

### Return type

[**ProvisioningApiGroupsGetGroupUsers200Response**](ProvisioningApiGroupsGetGroupUsers200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | User IDs returned |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **provisioning_api_groups_get_group_users_details**
> ProvisioningApiGroupsGetGroupUsersDetails200Response provisioning_api_groups_get_group_users_details(group_id, ocs_api_request, search=search, limit=limit, offset=offset)

Get a list of users details in the specified group

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import openapi_client
from openapi_client.models.provisioning_api_groups_get_group_users_details200_response import ProvisioningApiGroupsGetGroupUsersDetails200Response
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "http://localhost"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure HTTP basic authorization: basic_auth
configuration = openapi_client.Configuration(
    username = os.environ["USERNAME"],
    password = os.environ["PASSWORD"]
)

# Configure Bearer authorization: bearer_auth
configuration = openapi_client.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.ProvisioningApiGroupsApi(api_client)
    group_id = 'group_id_example' # str | ID of the group
    ocs_api_request = 'true' # str |  (default to 'true')
    search = '' # str | Text to search for (optional) (default to '')
    limit = 56 # int | Limit the amount of groups returned (optional)
    offset = 0 # int | Offset for searching for groups (optional) (default to 0)

    try:
        # Get a list of users details in the specified group
        api_response = api_instance.provisioning_api_groups_get_group_users_details(group_id, ocs_api_request, search=search, limit=limit, offset=offset)
        print("The response of ProvisioningApiGroupsApi->provisioning_api_groups_get_group_users_details:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ProvisioningApiGroupsApi->provisioning_api_groups_get_group_users_details: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group_id** | **str**| ID of the group | 
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]
 **search** | **str**| Text to search for | [optional] [default to &#39;&#39;]
 **limit** | **int**| Limit the amount of groups returned | [optional] 
 **offset** | **int**| Offset for searching for groups | [optional] [default to 0]

### Return type

[**ProvisioningApiGroupsGetGroupUsersDetails200Response**](ProvisioningApiGroupsGetGroupUsersDetails200Response.md)

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

# **provisioning_api_groups_get_groups**
> ProvisioningApiGroupsGetGroups200Response provisioning_api_groups_get_groups(ocs_api_request, search=search, limit=limit, offset=offset)

Get a list of groups

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import openapi_client
from openapi_client.models.provisioning_api_groups_get_groups200_response import ProvisioningApiGroupsGetGroups200Response
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "http://localhost"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure HTTP basic authorization: basic_auth
configuration = openapi_client.Configuration(
    username = os.environ["USERNAME"],
    password = os.environ["PASSWORD"]
)

# Configure Bearer authorization: bearer_auth
configuration = openapi_client.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.ProvisioningApiGroupsApi(api_client)
    ocs_api_request = 'true' # str |  (default to 'true')
    search = '' # str | Text to search for (optional) (default to '')
    limit = 56 # int | Limit the amount of groups returned (optional)
    offset = 0 # int | Offset for searching for groups (optional) (default to 0)

    try:
        # Get a list of groups
        api_response = api_instance.provisioning_api_groups_get_groups(ocs_api_request, search=search, limit=limit, offset=offset)
        print("The response of ProvisioningApiGroupsApi->provisioning_api_groups_get_groups:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ProvisioningApiGroupsApi->provisioning_api_groups_get_groups: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]
 **search** | **str**| Text to search for | [optional] [default to &#39;&#39;]
 **limit** | **int**| Limit the amount of groups returned | [optional] 
 **offset** | **int**| Offset for searching for groups | [optional] [default to 0]

### Return type

[**ProvisioningApiGroupsGetGroups200Response**](ProvisioningApiGroupsGetGroups200Response.md)

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

# **provisioning_api_groups_get_groups_details**
> ProvisioningApiGroupsGetGroupsDetails200Response provisioning_api_groups_get_groups_details(ocs_api_request, search=search, limit=limit, offset=offset)

Get a list of groups details

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import openapi_client
from openapi_client.models.provisioning_api_groups_get_groups_details200_response import ProvisioningApiGroupsGetGroupsDetails200Response
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "http://localhost"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure HTTP basic authorization: basic_auth
configuration = openapi_client.Configuration(
    username = os.environ["USERNAME"],
    password = os.environ["PASSWORD"]
)

# Configure Bearer authorization: bearer_auth
configuration = openapi_client.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.ProvisioningApiGroupsApi(api_client)
    ocs_api_request = 'true' # str |  (default to 'true')
    search = '' # str | Text to search for (optional) (default to '')
    limit = 56 # int | Limit the amount of groups returned (optional)
    offset = 0 # int | Offset for searching for groups (optional) (default to 0)

    try:
        # Get a list of groups details
        api_response = api_instance.provisioning_api_groups_get_groups_details(ocs_api_request, search=search, limit=limit, offset=offset)
        print("The response of ProvisioningApiGroupsApi->provisioning_api_groups_get_groups_details:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ProvisioningApiGroupsApi->provisioning_api_groups_get_groups_details: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]
 **search** | **str**| Text to search for | [optional] [default to &#39;&#39;]
 **limit** | **int**| Limit the amount of groups returned | [optional] 
 **offset** | **int**| Offset for searching for groups | [optional] [default to 0]

### Return type

[**ProvisioningApiGroupsGetGroupsDetails200Response**](ProvisioningApiGroupsGetGroupsDetails200Response.md)

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

# **provisioning_api_groups_get_sub_admins_of_group**
> ProvisioningApiGroupsGetSubAdminsOfGroup200Response provisioning_api_groups_get_sub_admins_of_group(group_id, ocs_api_request)

Get the list of user IDs that are a subadmin of the group

This endpoint requires admin access

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import openapi_client
from openapi_client.models.provisioning_api_groups_get_sub_admins_of_group200_response import ProvisioningApiGroupsGetSubAdminsOfGroup200Response
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "http://localhost"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure HTTP basic authorization: basic_auth
configuration = openapi_client.Configuration(
    username = os.environ["USERNAME"],
    password = os.environ["PASSWORD"]
)

# Configure Bearer authorization: bearer_auth
configuration = openapi_client.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.ProvisioningApiGroupsApi(api_client)
    group_id = 'group_id_example' # str | ID of the group
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Get the list of user IDs that are a subadmin of the group
        api_response = api_instance.provisioning_api_groups_get_sub_admins_of_group(group_id, ocs_api_request)
        print("The response of ProvisioningApiGroupsApi->provisioning_api_groups_get_sub_admins_of_group:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ProvisioningApiGroupsApi->provisioning_api_groups_get_sub_admins_of_group: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group_id** | **str**| ID of the group | 
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]

### Return type

[**ProvisioningApiGroupsGetSubAdminsOfGroup200Response**](ProvisioningApiGroupsGetSubAdminsOfGroup200Response.md)

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

# **provisioning_api_groups_update_group**
> CoreWhatsNewDismiss200Response provisioning_api_groups_update_group(key, value, group_id, ocs_api_request)

Update a group

This endpoint requires admin access

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import openapi_client
from openapi_client.models.core_whats_new_dismiss200_response import CoreWhatsNewDismiss200Response
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "http://localhost"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure HTTP basic authorization: basic_auth
configuration = openapi_client.Configuration(
    username = os.environ["USERNAME"],
    password = os.environ["PASSWORD"]
)

# Configure Bearer authorization: bearer_auth
configuration = openapi_client.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.ProvisioningApiGroupsApi(api_client)
    key = 'key_example' # str | Key to update, only 'displayname'
    value = 'value_example' # str | New value for the key
    group_id = 'group_id_example' # str | ID of the group
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Update a group
        api_response = api_instance.provisioning_api_groups_update_group(key, value, group_id, ocs_api_request)
        print("The response of ProvisioningApiGroupsApi->provisioning_api_groups_update_group:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ProvisioningApiGroupsApi->provisioning_api_groups_update_group: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **key** | **str**| Key to update, only &#39;displayname&#39; | 
 **value** | **str**| New value for the key | 
 **group_id** | **str**| ID of the group | 
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
**200** |  |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

