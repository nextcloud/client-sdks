# openapi_client.CoreCollaborationResourcesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**core_collaboration_resources_add_resource**](CoreCollaborationResourcesApi.md#core_collaboration_resources_add_resource) | **POST** /ocs/v2.php/collaboration/resources/collections/{collectionId} | Add a resource to a collection
[**core_collaboration_resources_create_collection_on_resource**](CoreCollaborationResourcesApi.md#core_collaboration_resources_create_collection_on_resource) | **POST** /ocs/v2.php/collaboration/resources/{baseResourceType}/{baseResourceId} | Create a collection for a resource
[**core_collaboration_resources_get_collections_by_resource**](CoreCollaborationResourcesApi.md#core_collaboration_resources_get_collections_by_resource) | **GET** /ocs/v2.php/collaboration/resources/{resourceType}/{resourceId} | Get collections by resource
[**core_collaboration_resources_list_collection**](CoreCollaborationResourcesApi.md#core_collaboration_resources_list_collection) | **GET** /ocs/v2.php/collaboration/resources/collections/{collectionId} | Get a collection
[**core_collaboration_resources_remove_resource**](CoreCollaborationResourcesApi.md#core_collaboration_resources_remove_resource) | **DELETE** /ocs/v2.php/collaboration/resources/collections/{collectionId} | Remove a resource from a collection
[**core_collaboration_resources_rename_collection**](CoreCollaborationResourcesApi.md#core_collaboration_resources_rename_collection) | **PUT** /ocs/v2.php/collaboration/resources/collections/{collectionId} | Rename a collection
[**core_collaboration_resources_search_collections**](CoreCollaborationResourcesApi.md#core_collaboration_resources_search_collections) | **GET** /ocs/v2.php/collaboration/resources/collections/search/{filter} | Search for collections


# **core_collaboration_resources_add_resource**
> CoreCollaborationResourcesListCollection200Response core_collaboration_resources_add_resource(resource_type, resource_id, collection_id, ocs_api_request)

Add a resource to a collection

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import openapi_client
from openapi_client.models.core_collaboration_resources_list_collection200_response import CoreCollaborationResourcesListCollection200Response
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
    api_instance = openapi_client.CoreCollaborationResourcesApi(api_client)
    resource_type = 'resource_type_example' # str | Name of the resource
    resource_id = 'resource_id_example' # str | ID of the resource
    collection_id = 56 # int | ID of the collection
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Add a resource to a collection
        api_response = api_instance.core_collaboration_resources_add_resource(resource_type, resource_id, collection_id, ocs_api_request)
        print("The response of CoreCollaborationResourcesApi->core_collaboration_resources_add_resource:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CoreCollaborationResourcesApi->core_collaboration_resources_add_resource: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resource_type** | **str**| Name of the resource | 
 **resource_id** | **str**| ID of the resource | 
 **collection_id** | **int**| ID of the collection | 
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]

### Return type

[**CoreCollaborationResourcesListCollection200Response**](CoreCollaborationResourcesListCollection200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Collection returned |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **core_collaboration_resources_create_collection_on_resource**
> CoreCollaborationResourcesListCollection200Response core_collaboration_resources_create_collection_on_resource(name, base_resource_type, base_resource_id, ocs_api_request)

Create a collection for a resource

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import openapi_client
from openapi_client.models.core_collaboration_resources_list_collection200_response import CoreCollaborationResourcesListCollection200Response
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
    api_instance = openapi_client.CoreCollaborationResourcesApi(api_client)
    name = 'name_example' # str | Name of the collection
    base_resource_type = 'base_resource_type_example' # str | Type of the base resource
    base_resource_id = 'base_resource_id_example' # str | ID of the base resource
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Create a collection for a resource
        api_response = api_instance.core_collaboration_resources_create_collection_on_resource(name, base_resource_type, base_resource_id, ocs_api_request)
        print("The response of CoreCollaborationResourcesApi->core_collaboration_resources_create_collection_on_resource:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CoreCollaborationResourcesApi->core_collaboration_resources_create_collection_on_resource: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **str**| Name of the collection | 
 **base_resource_type** | **str**| Type of the base resource | 
 **base_resource_id** | **str**| ID of the base resource | 
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]

### Return type

[**CoreCollaborationResourcesListCollection200Response**](CoreCollaborationResourcesListCollection200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Collection returned |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **core_collaboration_resources_get_collections_by_resource**
> CoreCollaborationResourcesSearchCollections200Response core_collaboration_resources_get_collections_by_resource(resource_type, resource_id, ocs_api_request)

Get collections by resource

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import openapi_client
from openapi_client.models.core_collaboration_resources_search_collections200_response import CoreCollaborationResourcesSearchCollections200Response
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
    api_instance = openapi_client.CoreCollaborationResourcesApi(api_client)
    resource_type = 'resource_type_example' # str | Type of the resource
    resource_id = 'resource_id_example' # str | ID of the resource
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Get collections by resource
        api_response = api_instance.core_collaboration_resources_get_collections_by_resource(resource_type, resource_id, ocs_api_request)
        print("The response of CoreCollaborationResourcesApi->core_collaboration_resources_get_collections_by_resource:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CoreCollaborationResourcesApi->core_collaboration_resources_get_collections_by_resource: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resource_type** | **str**| Type of the resource | 
 **resource_id** | **str**| ID of the resource | 
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]

### Return type

[**CoreCollaborationResourcesSearchCollections200Response**](CoreCollaborationResourcesSearchCollections200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Collections returned |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **core_collaboration_resources_list_collection**
> CoreCollaborationResourcesListCollection200Response core_collaboration_resources_list_collection(collection_id, ocs_api_request)

Get a collection

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import openapi_client
from openapi_client.models.core_collaboration_resources_list_collection200_response import CoreCollaborationResourcesListCollection200Response
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
    api_instance = openapi_client.CoreCollaborationResourcesApi(api_client)
    collection_id = 56 # int | ID of the collection
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Get a collection
        api_response = api_instance.core_collaboration_resources_list_collection(collection_id, ocs_api_request)
        print("The response of CoreCollaborationResourcesApi->core_collaboration_resources_list_collection:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CoreCollaborationResourcesApi->core_collaboration_resources_list_collection: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **collection_id** | **int**| ID of the collection | 
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]

### Return type

[**CoreCollaborationResourcesListCollection200Response**](CoreCollaborationResourcesListCollection200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Collection returned |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **core_collaboration_resources_remove_resource**
> CoreCollaborationResourcesListCollection200Response core_collaboration_resources_remove_resource(resource_type, resource_id, collection_id, ocs_api_request)

Remove a resource from a collection

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import openapi_client
from openapi_client.models.core_collaboration_resources_list_collection200_response import CoreCollaborationResourcesListCollection200Response
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
    api_instance = openapi_client.CoreCollaborationResourcesApi(api_client)
    resource_type = 'resource_type_example' # str | Name of the resource
    resource_id = 'resource_id_example' # str | ID of the resource
    collection_id = 56 # int | ID of the collection
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Remove a resource from a collection
        api_response = api_instance.core_collaboration_resources_remove_resource(resource_type, resource_id, collection_id, ocs_api_request)
        print("The response of CoreCollaborationResourcesApi->core_collaboration_resources_remove_resource:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CoreCollaborationResourcesApi->core_collaboration_resources_remove_resource: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resource_type** | **str**| Name of the resource | 
 **resource_id** | **str**| ID of the resource | 
 **collection_id** | **int**| ID of the collection | 
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]

### Return type

[**CoreCollaborationResourcesListCollection200Response**](CoreCollaborationResourcesListCollection200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Collection returned |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **core_collaboration_resources_rename_collection**
> CoreCollaborationResourcesListCollection200Response core_collaboration_resources_rename_collection(collection_name, collection_id, ocs_api_request)

Rename a collection

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import openapi_client
from openapi_client.models.core_collaboration_resources_list_collection200_response import CoreCollaborationResourcesListCollection200Response
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
    api_instance = openapi_client.CoreCollaborationResourcesApi(api_client)
    collection_name = 'collection_name_example' # str | New name
    collection_id = 56 # int | ID of the collection
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Rename a collection
        api_response = api_instance.core_collaboration_resources_rename_collection(collection_name, collection_id, ocs_api_request)
        print("The response of CoreCollaborationResourcesApi->core_collaboration_resources_rename_collection:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CoreCollaborationResourcesApi->core_collaboration_resources_rename_collection: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **collection_name** | **str**| New name | 
 **collection_id** | **int**| ID of the collection | 
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]

### Return type

[**CoreCollaborationResourcesListCollection200Response**](CoreCollaborationResourcesListCollection200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Collection returned |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **core_collaboration_resources_search_collections**
> CoreCollaborationResourcesSearchCollections200Response core_collaboration_resources_search_collections(filter, ocs_api_request)

Search for collections

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import openapi_client
from openapi_client.models.core_collaboration_resources_search_collections200_response import CoreCollaborationResourcesSearchCollections200Response
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
    api_instance = openapi_client.CoreCollaborationResourcesApi(api_client)
    filter = 'filter_example' # str | Filter collections
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Search for collections
        api_response = api_instance.core_collaboration_resources_search_collections(filter, ocs_api_request)
        print("The response of CoreCollaborationResourcesApi->core_collaboration_resources_search_collections:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CoreCollaborationResourcesApi->core_collaboration_resources_search_collections: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **str**| Filter collections | 
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]

### Return type

[**CoreCollaborationResourcesSearchCollections200Response**](CoreCollaborationResourcesSearchCollections200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Collections returned |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

