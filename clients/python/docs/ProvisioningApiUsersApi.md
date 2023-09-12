# nextcloud_client.ProvisioningApiUsersApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**provisioning_api_users_add_sub_admin**](ProvisioningApiUsersApi.md#provisioning_api_users_add_sub_admin) | **POST** /ocs/v2.php/cloud/users/{userId}/subadmins | Make a user a subadmin of a group
[**provisioning_api_users_add_to_group**](ProvisioningApiUsersApi.md#provisioning_api_users_add_to_group) | **POST** /ocs/v2.php/cloud/users/{userId}/groups | Add a user to a group
[**provisioning_api_users_add_user**](ProvisioningApiUsersApi.md#provisioning_api_users_add_user) | **POST** /ocs/v2.php/cloud/users | Create a new user
[**provisioning_api_users_delete_user**](ProvisioningApiUsersApi.md#provisioning_api_users_delete_user) | **DELETE** /ocs/v2.php/cloud/users/{userId} | Delete a user
[**provisioning_api_users_disable_user**](ProvisioningApiUsersApi.md#provisioning_api_users_disable_user) | **PUT** /ocs/v2.php/cloud/users/{userId}/disable | Disable a user
[**provisioning_api_users_edit_user**](ProvisioningApiUsersApi.md#provisioning_api_users_edit_user) | **PUT** /ocs/v2.php/cloud/users/{userId} | Update a value of the user&#39;s details
[**provisioning_api_users_edit_user_multi_value**](ProvisioningApiUsersApi.md#provisioning_api_users_edit_user_multi_value) | **PUT** /ocs/v2.php/cloud/users/{userId}/{collectionName} | Update multiple values of the user&#39;s details
[**provisioning_api_users_enable_user**](ProvisioningApiUsersApi.md#provisioning_api_users_enable_user) | **PUT** /ocs/v2.php/cloud/users/{userId}/enable | Enable a user
[**provisioning_api_users_get_current_user**](ProvisioningApiUsersApi.md#provisioning_api_users_get_current_user) | **GET** /ocs/v2.php/cloud/user | Get the details of the current user
[**provisioning_api_users_get_editable_fields**](ProvisioningApiUsersApi.md#provisioning_api_users_get_editable_fields) | **GET** /ocs/v2.php/cloud/user/fields | Get a list of fields that are editable for the current user
[**provisioning_api_users_get_editable_fields_for_user**](ProvisioningApiUsersApi.md#provisioning_api_users_get_editable_fields_for_user) | **GET** /ocs/v2.php/cloud/user/fields/{userId} | Get a list of fields that are editable for a user
[**provisioning_api_users_get_user**](ProvisioningApiUsersApi.md#provisioning_api_users_get_user) | **GET** /ocs/v2.php/cloud/users/{userId} | Get the details of a user
[**provisioning_api_users_get_user_sub_admin_groups**](ProvisioningApiUsersApi.md#provisioning_api_users_get_user_sub_admin_groups) | **GET** /ocs/v2.php/cloud/users/{userId}/subadmins | Get the groups a user is a subadmin of
[**provisioning_api_users_get_users**](ProvisioningApiUsersApi.md#provisioning_api_users_get_users) | **GET** /ocs/v2.php/cloud/users | Get a list of users
[**provisioning_api_users_get_users_details**](ProvisioningApiUsersApi.md#provisioning_api_users_get_users_details) | **GET** /ocs/v2.php/cloud/users/details | Get a list of users and their details
[**provisioning_api_users_get_users_groups**](ProvisioningApiUsersApi.md#provisioning_api_users_get_users_groups) | **GET** /ocs/v2.php/cloud/users/{userId}/groups | Get a list of groups the user belongs to
[**provisioning_api_users_remove_from_group**](ProvisioningApiUsersApi.md#provisioning_api_users_remove_from_group) | **DELETE** /ocs/v2.php/cloud/users/{userId}/groups | Remove a user from a group
[**provisioning_api_users_remove_sub_admin**](ProvisioningApiUsersApi.md#provisioning_api_users_remove_sub_admin) | **DELETE** /ocs/v2.php/cloud/users/{userId}/subadmins | Remove a user from the subadmins of a group
[**provisioning_api_users_resend_welcome_message**](ProvisioningApiUsersApi.md#provisioning_api_users_resend_welcome_message) | **POST** /ocs/v2.php/cloud/users/{userId}/welcome | Resend the welcome message
[**provisioning_api_users_search_by_phone_numbers**](ProvisioningApiUsersApi.md#provisioning_api_users_search_by_phone_numbers) | **POST** /ocs/v2.php/cloud/users/search/by-phone | Search users by their phone numbers
[**provisioning_api_users_wipe_user_devices**](ProvisioningApiUsersApi.md#provisioning_api_users_wipe_user_devices) | **POST** /ocs/v2.php/cloud/users/{userId}/wipe | Wipe all devices of a user


# **provisioning_api_users_add_sub_admin**
> CoreWhatsNewDismiss200Response provisioning_api_users_add_sub_admin(groupid, user_id, ocs_api_request)

Make a user a subadmin of a group

This endpoint requires admin access

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
    api_instance = nextcloud_client.ProvisioningApiUsersApi(api_client)
    groupid = 'groupid_example' # str | ID of the group
    user_id = 'user_id_example' # str | ID of the user
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Make a user a subadmin of a group
        api_response = api_instance.provisioning_api_users_add_sub_admin(groupid, user_id, ocs_api_request)
        print("The response of ProvisioningApiUsersApi->provisioning_api_users_add_sub_admin:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ProvisioningApiUsersApi->provisioning_api_users_add_sub_admin: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupid** | **str**| ID of the group | 
 **user_id** | **str**| ID of the user | 
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

# **provisioning_api_users_add_to_group**
> CoreWhatsNewDismiss200Response provisioning_api_users_add_to_group(user_id, ocs_api_request, groupid=groupid)

Add a user to a group

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
    api_instance = nextcloud_client.ProvisioningApiUsersApi(api_client)
    user_id = 'user_id_example' # str | ID of the user
    ocs_api_request = 'true' # str |  (default to 'true')
    groupid = '' # str | ID of the group (optional) (default to '')

    try:
        # Add a user to a group
        api_response = api_instance.provisioning_api_users_add_to_group(user_id, ocs_api_request, groupid=groupid)
        print("The response of ProvisioningApiUsersApi->provisioning_api_users_add_to_group:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ProvisioningApiUsersApi->provisioning_api_users_add_to_group: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user_id** | **str**| ID of the user | 
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]
 **groupid** | **str**| ID of the group | [optional] [default to &#39;&#39;]

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

# **provisioning_api_users_add_user**
> ProvisioningApiUsersAddUser200Response provisioning_api_users_add_user(userid, ocs_api_request, password=password, display_name=display_name, email=email, groups=groups, subadmin=subadmin, quota=quota, language=language, manager=manager)

Create a new user

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.provisioning_api_users_add_user200_response import ProvisioningApiUsersAddUser200Response
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
    api_instance = nextcloud_client.ProvisioningApiUsersApi(api_client)
    userid = 'userid_example' # str | ID of the user
    ocs_api_request = 'true' # str |  (default to 'true')
    password = '' # str | Password of the user (optional) (default to '')
    display_name = '' # str | Display name of the user (optional) (default to '')
    email = '' # str | Email of the user (optional) (default to '')
    groups = [] # List[str] | Groups of the user (optional) (default to [])
    subadmin = [] # List[str] | Groups where the user is subadmin (optional) (default to [])
    quota = '' # str | Quota of the user (optional) (default to '')
    language = '' # str | Language of the user (optional) (default to '')
    manager = 'manager_example' # str | Manager of the user (optional)

    try:
        # Create a new user
        api_response = api_instance.provisioning_api_users_add_user(userid, ocs_api_request, password=password, display_name=display_name, email=email, groups=groups, subadmin=subadmin, quota=quota, language=language, manager=manager)
        print("The response of ProvisioningApiUsersApi->provisioning_api_users_add_user:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ProvisioningApiUsersApi->provisioning_api_users_add_user: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userid** | **str**| ID of the user | 
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]
 **password** | **str**| Password of the user | [optional] [default to &#39;&#39;]
 **display_name** | **str**| Display name of the user | [optional] [default to &#39;&#39;]
 **email** | **str**| Email of the user | [optional] [default to &#39;&#39;]
 **groups** | [**List[str]**](str.md)| Groups of the user | [optional] [default to []]
 **subadmin** | [**List[str]**](str.md)| Groups where the user is subadmin | [optional] [default to []]
 **quota** | **str**| Quota of the user | [optional] [default to &#39;&#39;]
 **language** | **str**| Language of the user | [optional] [default to &#39;&#39;]
 **manager** | **str**| Manager of the user | [optional] 

### Return type

[**ProvisioningApiUsersAddUser200Response**](ProvisioningApiUsersAddUser200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | User added successfully |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **provisioning_api_users_delete_user**
> CoreWhatsNewDismiss200Response provisioning_api_users_delete_user(user_id, ocs_api_request)

Delete a user

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
    api_instance = nextcloud_client.ProvisioningApiUsersApi(api_client)
    user_id = 'user_id_example' # str | ID of the user
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Delete a user
        api_response = api_instance.provisioning_api_users_delete_user(user_id, ocs_api_request)
        print("The response of ProvisioningApiUsersApi->provisioning_api_users_delete_user:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ProvisioningApiUsersApi->provisioning_api_users_delete_user: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user_id** | **str**| ID of the user | 
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

# **provisioning_api_users_disable_user**
> CoreWhatsNewDismiss200Response provisioning_api_users_disable_user(user_id, ocs_api_request)

Disable a user

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
    api_instance = nextcloud_client.ProvisioningApiUsersApi(api_client)
    user_id = 'user_id_example' # str | ID of the user
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Disable a user
        api_response = api_instance.provisioning_api_users_disable_user(user_id, ocs_api_request)
        print("The response of ProvisioningApiUsersApi->provisioning_api_users_disable_user:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ProvisioningApiUsersApi->provisioning_api_users_disable_user: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user_id** | **str**| ID of the user | 
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

# **provisioning_api_users_edit_user**
> CoreWhatsNewDismiss200Response provisioning_api_users_edit_user(key, value, user_id, ocs_api_request)

Update a value of the user's details

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
    api_instance = nextcloud_client.ProvisioningApiUsersApi(api_client)
    key = 'key_example' # str | Key that will be updated
    value = 'value_example' # str | New value for the key
    user_id = 'user_id_example' # str | ID of the user
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Update a value of the user's details
        api_response = api_instance.provisioning_api_users_edit_user(key, value, user_id, ocs_api_request)
        print("The response of ProvisioningApiUsersApi->provisioning_api_users_edit_user:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ProvisioningApiUsersApi->provisioning_api_users_edit_user: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **key** | **str**| Key that will be updated | 
 **value** | **str**| New value for the key | 
 **user_id** | **str**| ID of the user | 
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

# **provisioning_api_users_edit_user_multi_value**
> CoreWhatsNewDismiss200Response provisioning_api_users_edit_user_multi_value(key, value, user_id, collection_name, ocs_api_request)

Update multiple values of the user's details

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
    api_instance = nextcloud_client.ProvisioningApiUsersApi(api_client)
    key = 'key_example' # str | Key that will be updated
    value = 'value_example' # str | New value for the key
    user_id = 'user_id_example' # str | ID of the user
    collection_name = 'collection_name_example' # str | Collection to update
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Update multiple values of the user's details
        api_response = api_instance.provisioning_api_users_edit_user_multi_value(key, value, user_id, collection_name, ocs_api_request)
        print("The response of ProvisioningApiUsersApi->provisioning_api_users_edit_user_multi_value:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ProvisioningApiUsersApi->provisioning_api_users_edit_user_multi_value: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **key** | **str**| Key that will be updated | 
 **value** | **str**| New value for the key | 
 **user_id** | **str**| ID of the user | 
 **collection_name** | **str**| Collection to update | 
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

# **provisioning_api_users_enable_user**
> CoreWhatsNewDismiss200Response provisioning_api_users_enable_user(user_id, ocs_api_request)

Enable a user

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
    api_instance = nextcloud_client.ProvisioningApiUsersApi(api_client)
    user_id = 'user_id_example' # str | ID of the user
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Enable a user
        api_response = api_instance.provisioning_api_users_enable_user(user_id, ocs_api_request)
        print("The response of ProvisioningApiUsersApi->provisioning_api_users_enable_user:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ProvisioningApiUsersApi->provisioning_api_users_enable_user: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user_id** | **str**| ID of the user | 
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

# **provisioning_api_users_get_current_user**
> ProvisioningApiUsersGetUser200Response provisioning_api_users_get_current_user(ocs_api_request)

Get the details of the current user

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.provisioning_api_users_get_user200_response import ProvisioningApiUsersGetUser200Response
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
    api_instance = nextcloud_client.ProvisioningApiUsersApi(api_client)
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Get the details of the current user
        api_response = api_instance.provisioning_api_users_get_current_user(ocs_api_request)
        print("The response of ProvisioningApiUsersApi->provisioning_api_users_get_current_user:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ProvisioningApiUsersApi->provisioning_api_users_get_current_user: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]

### Return type

[**ProvisioningApiUsersGetUser200Response**](ProvisioningApiUsersGetUser200Response.md)

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

# **provisioning_api_users_get_editable_fields**
> ProvisioningApiGroupsGetSubAdminsOfGroup200Response provisioning_api_users_get_editable_fields(ocs_api_request)

Get a list of fields that are editable for the current user

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.provisioning_api_groups_get_sub_admins_of_group200_response import ProvisioningApiGroupsGetSubAdminsOfGroup200Response
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
    api_instance = nextcloud_client.ProvisioningApiUsersApi(api_client)
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Get a list of fields that are editable for the current user
        api_response = api_instance.provisioning_api_users_get_editable_fields(ocs_api_request)
        print("The response of ProvisioningApiUsersApi->provisioning_api_users_get_editable_fields:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ProvisioningApiUsersApi->provisioning_api_users_get_editable_fields: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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

# **provisioning_api_users_get_editable_fields_for_user**
> ProvisioningApiGroupsGetSubAdminsOfGroup200Response provisioning_api_users_get_editable_fields_for_user(user_id, ocs_api_request)

Get a list of fields that are editable for a user

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.provisioning_api_groups_get_sub_admins_of_group200_response import ProvisioningApiGroupsGetSubAdminsOfGroup200Response
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
    api_instance = nextcloud_client.ProvisioningApiUsersApi(api_client)
    user_id = 'user_id_example' # str | ID of the user
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Get a list of fields that are editable for a user
        api_response = api_instance.provisioning_api_users_get_editable_fields_for_user(user_id, ocs_api_request)
        print("The response of ProvisioningApiUsersApi->provisioning_api_users_get_editable_fields_for_user:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ProvisioningApiUsersApi->provisioning_api_users_get_editable_fields_for_user: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user_id** | **str**| ID of the user | 
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

# **provisioning_api_users_get_user**
> ProvisioningApiUsersGetUser200Response provisioning_api_users_get_user(user_id, ocs_api_request)

Get the details of a user

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.provisioning_api_users_get_user200_response import ProvisioningApiUsersGetUser200Response
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
    api_instance = nextcloud_client.ProvisioningApiUsersApi(api_client)
    user_id = 'user_id_example' # str | ID of the user
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Get the details of a user
        api_response = api_instance.provisioning_api_users_get_user(user_id, ocs_api_request)
        print("The response of ProvisioningApiUsersApi->provisioning_api_users_get_user:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ProvisioningApiUsersApi->provisioning_api_users_get_user: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user_id** | **str**| ID of the user | 
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]

### Return type

[**ProvisioningApiUsersGetUser200Response**](ProvisioningApiUsersGetUser200Response.md)

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

# **provisioning_api_users_get_user_sub_admin_groups**
> ProvisioningApiGroupsGetSubAdminsOfGroup200Response provisioning_api_users_get_user_sub_admin_groups(user_id, ocs_api_request)

Get the groups a user is a subadmin of

This endpoint requires admin access

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.provisioning_api_groups_get_sub_admins_of_group200_response import ProvisioningApiGroupsGetSubAdminsOfGroup200Response
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
    api_instance = nextcloud_client.ProvisioningApiUsersApi(api_client)
    user_id = 'user_id_example' # str | ID if the user
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Get the groups a user is a subadmin of
        api_response = api_instance.provisioning_api_users_get_user_sub_admin_groups(user_id, ocs_api_request)
        print("The response of ProvisioningApiUsersApi->provisioning_api_users_get_user_sub_admin_groups:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ProvisioningApiUsersApi->provisioning_api_users_get_user_sub_admin_groups: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user_id** | **str**| ID if the user | 
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

# **provisioning_api_users_get_users**
> ProvisioningApiGroupsGetGroupUsers200Response provisioning_api_users_get_users(ocs_api_request, search=search, limit=limit, offset=offset)

Get a list of users

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.provisioning_api_groups_get_group_users200_response import ProvisioningApiGroupsGetGroupUsers200Response
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
    api_instance = nextcloud_client.ProvisioningApiUsersApi(api_client)
    ocs_api_request = 'true' # str |  (default to 'true')
    search = '' # str | Text to search for (optional) (default to '')
    limit = 56 # int | Limit the amount of groups returned (optional)
    offset = 0 # int | Offset for searching for groups (optional) (default to 0)

    try:
        # Get a list of users
        api_response = api_instance.provisioning_api_users_get_users(ocs_api_request, search=search, limit=limit, offset=offset)
        print("The response of ProvisioningApiUsersApi->provisioning_api_users_get_users:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ProvisioningApiUsersApi->provisioning_api_users_get_users: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]
 **search** | **str**| Text to search for | [optional] [default to &#39;&#39;]
 **limit** | **int**| Limit the amount of groups returned | [optional] 
 **offset** | **int**| Offset for searching for groups | [optional] [default to 0]

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

# **provisioning_api_users_get_users_details**
> ProvisioningApiGroupsGetGroupUsersDetails200Response provisioning_api_users_get_users_details(ocs_api_request, search=search, limit=limit, offset=offset)

Get a list of users and their details

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.provisioning_api_groups_get_group_users_details200_response import ProvisioningApiGroupsGetGroupUsersDetails200Response
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
    api_instance = nextcloud_client.ProvisioningApiUsersApi(api_client)
    ocs_api_request = 'true' # str |  (default to 'true')
    search = '' # str | Text to search for (optional) (default to '')
    limit = 56 # int | Limit the amount of groups returned (optional)
    offset = 0 # int | Offset for searching for groups (optional) (default to 0)

    try:
        # Get a list of users and their details
        api_response = api_instance.provisioning_api_users_get_users_details(ocs_api_request, search=search, limit=limit, offset=offset)
        print("The response of ProvisioningApiUsersApi->provisioning_api_users_get_users_details:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ProvisioningApiUsersApi->provisioning_api_users_get_users_details: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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

# **provisioning_api_users_get_users_groups**
> ProvisioningApiGroupsGetGroups200Response provisioning_api_users_get_users_groups(user_id, ocs_api_request)

Get a list of groups the user belongs to

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.provisioning_api_groups_get_groups200_response import ProvisioningApiGroupsGetGroups200Response
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
    api_instance = nextcloud_client.ProvisioningApiUsersApi(api_client)
    user_id = 'user_id_example' # str | ID of the user
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Get a list of groups the user belongs to
        api_response = api_instance.provisioning_api_users_get_users_groups(user_id, ocs_api_request)
        print("The response of ProvisioningApiUsersApi->provisioning_api_users_get_users_groups:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ProvisioningApiUsersApi->provisioning_api_users_get_users_groups: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user_id** | **str**| ID of the user | 
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]

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

# **provisioning_api_users_remove_from_group**
> CoreWhatsNewDismiss200Response provisioning_api_users_remove_from_group(groupid, user_id, ocs_api_request)

Remove a user from a group

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
    api_instance = nextcloud_client.ProvisioningApiUsersApi(api_client)
    groupid = 'groupid_example' # str | ID of the group
    user_id = 'user_id_example' # str | ID of the user
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Remove a user from a group
        api_response = api_instance.provisioning_api_users_remove_from_group(groupid, user_id, ocs_api_request)
        print("The response of ProvisioningApiUsersApi->provisioning_api_users_remove_from_group:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ProvisioningApiUsersApi->provisioning_api_users_remove_from_group: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupid** | **str**| ID of the group | 
 **user_id** | **str**| ID of the user | 
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

# **provisioning_api_users_remove_sub_admin**
> CoreWhatsNewDismiss200Response provisioning_api_users_remove_sub_admin(groupid, user_id, ocs_api_request)

Remove a user from the subadmins of a group

This endpoint requires admin access

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
    api_instance = nextcloud_client.ProvisioningApiUsersApi(api_client)
    groupid = 'groupid_example' # str | ID of the group
    user_id = 'user_id_example' # str | ID of the user
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Remove a user from the subadmins of a group
        api_response = api_instance.provisioning_api_users_remove_sub_admin(groupid, user_id, ocs_api_request)
        print("The response of ProvisioningApiUsersApi->provisioning_api_users_remove_sub_admin:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ProvisioningApiUsersApi->provisioning_api_users_remove_sub_admin: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupid** | **str**| ID of the group | 
 **user_id** | **str**| ID of the user | 
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

# **provisioning_api_users_resend_welcome_message**
> CoreWhatsNewDismiss200Response provisioning_api_users_resend_welcome_message(user_id, ocs_api_request)

Resend the welcome message

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
    api_instance = nextcloud_client.ProvisioningApiUsersApi(api_client)
    user_id = 'user_id_example' # str | ID if the user
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Resend the welcome message
        api_response = api_instance.provisioning_api_users_resend_welcome_message(user_id, ocs_api_request)
        print("The response of ProvisioningApiUsersApi->provisioning_api_users_resend_welcome_message:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ProvisioningApiUsersApi->provisioning_api_users_resend_welcome_message: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user_id** | **str**| ID if the user | 
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

# **provisioning_api_users_search_by_phone_numbers**
> ProvisioningApiUsersSearchByPhoneNumbers200Response provisioning_api_users_search_by_phone_numbers(location, search, ocs_api_request)

Search users by their phone numbers

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.provisioning_api_users_search_by_phone_numbers200_response import ProvisioningApiUsersSearchByPhoneNumbers200Response
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
    api_instance = nextcloud_client.ProvisioningApiUsersApi(api_client)
    location = 'location_example' # str | Location of the phone number (for country code)
    search = 'search_example' # str | Phone numbers to search for
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Search users by their phone numbers
        api_response = api_instance.provisioning_api_users_search_by_phone_numbers(location, search, ocs_api_request)
        print("The response of ProvisioningApiUsersApi->provisioning_api_users_search_by_phone_numbers:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ProvisioningApiUsersApi->provisioning_api_users_search_by_phone_numbers: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **location** | **str**| Location of the phone number (for country code) | 
 **search** | **str**| Phone numbers to search for | 
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]

### Return type

[**ProvisioningApiUsersSearchByPhoneNumbers200Response**](ProvisioningApiUsersSearchByPhoneNumbers200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Users returned |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **provisioning_api_users_wipe_user_devices**
> CoreWhatsNewDismiss200Response provisioning_api_users_wipe_user_devices(user_id, ocs_api_request)

Wipe all devices of a user

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
    api_instance = nextcloud_client.ProvisioningApiUsersApi(api_client)
    user_id = 'user_id_example' # str | ID of the user
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Wipe all devices of a user
        api_response = api_instance.provisioning_api_users_wipe_user_devices(user_id, ocs_api_request)
        print("The response of ProvisioningApiUsersApi->provisioning_api_users_wipe_user_devices:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ProvisioningApiUsersApi->provisioning_api_users_wipe_user_devices: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user_id** | **str**| ID of the user | 
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

