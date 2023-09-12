# nextcloud_client.DashboardDashboardApiApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**dashboard_dashboard_api_get_widget_items**](DashboardDashboardApiApi.md#dashboard_dashboard_api_get_widget_items) | **GET** /ocs/v2.php/apps/dashboard/api/v1/widget-items | Get the items for the widgets
[**dashboard_dashboard_api_get_widget_items_v2**](DashboardDashboardApiApi.md#dashboard_dashboard_api_get_widget_items_v2) | **GET** /ocs/v2.php/apps/dashboard/api/v2/widget-items | Get the items for the widgets
[**dashboard_dashboard_api_get_widgets**](DashboardDashboardApiApi.md#dashboard_dashboard_api_get_widgets) | **GET** /ocs/v2.php/apps/dashboard/api/v1/widgets | Get the widgets


# **dashboard_dashboard_api_get_widget_items**
> DashboardDashboardApiGetWidgetItems200Response dashboard_dashboard_api_get_widget_items(ocs_api_request, since_ids=since_ids, limit=limit, widgets=widgets)

Get the items for the widgets

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.dashboard_dashboard_api_get_widget_items200_response import DashboardDashboardApiGetWidgetItems200Response
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
    api_instance = nextcloud_client.DashboardDashboardApiApi(api_client)
    ocs_api_request = 'true' # str |  (default to 'true')
    since_ids = 'since_ids_example' # str | Array indexed by widget Ids, contains date/id from which we want the new items (optional)
    limit = 7 # int | Limit number of result items per widget (optional) (default to 7)
    widgets = [] # List[str] | Limit results to specific widgets (optional) (default to [])

    try:
        # Get the items for the widgets
        api_response = api_instance.dashboard_dashboard_api_get_widget_items(ocs_api_request, since_ids=since_ids, limit=limit, widgets=widgets)
        print("The response of DashboardDashboardApiApi->dashboard_dashboard_api_get_widget_items:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DashboardDashboardApiApi->dashboard_dashboard_api_get_widget_items: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]
 **since_ids** | **str**| Array indexed by widget Ids, contains date/id from which we want the new items | [optional] 
 **limit** | **int**| Limit number of result items per widget | [optional] [default to 7]
 **widgets** | [**List[str]**](str.md)| Limit results to specific widgets | [optional] [default to []]

### Return type

[**DashboardDashboardApiGetWidgetItems200Response**](DashboardDashboardApiGetWidgetItems200Response.md)

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

# **dashboard_dashboard_api_get_widget_items_v2**
> DashboardDashboardApiGetWidgetItemsV2200Response dashboard_dashboard_api_get_widget_items_v2(ocs_api_request, since_ids=since_ids, limit=limit, widgets=widgets)

Get the items for the widgets

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.dashboard_dashboard_api_get_widget_items_v2200_response import DashboardDashboardApiGetWidgetItemsV2200Response
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
    api_instance = nextcloud_client.DashboardDashboardApiApi(api_client)
    ocs_api_request = 'true' # str |  (default to 'true')
    since_ids = 'since_ids_example' # str | Array indexed by widget Ids, contains date/id from which we want the new items (optional)
    limit = 7 # int | Limit number of result items per widget (optional) (default to 7)
    widgets = [] # List[str] | Limit results to specific widgets (optional) (default to [])

    try:
        # Get the items for the widgets
        api_response = api_instance.dashboard_dashboard_api_get_widget_items_v2(ocs_api_request, since_ids=since_ids, limit=limit, widgets=widgets)
        print("The response of DashboardDashboardApiApi->dashboard_dashboard_api_get_widget_items_v2:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DashboardDashboardApiApi->dashboard_dashboard_api_get_widget_items_v2: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]
 **since_ids** | **str**| Array indexed by widget Ids, contains date/id from which we want the new items | [optional] 
 **limit** | **int**| Limit number of result items per widget | [optional] [default to 7]
 **widgets** | [**List[str]**](str.md)| Limit results to specific widgets | [optional] [default to []]

### Return type

[**DashboardDashboardApiGetWidgetItemsV2200Response**](DashboardDashboardApiGetWidgetItemsV2200Response.md)

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

# **dashboard_dashboard_api_get_widgets**
> DashboardDashboardApiGetWidgets200Response dashboard_dashboard_api_get_widgets(ocs_api_request)

Get the widgets

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.dashboard_dashboard_api_get_widgets200_response import DashboardDashboardApiGetWidgets200Response
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
    api_instance = nextcloud_client.DashboardDashboardApiApi(api_client)
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Get the widgets
        api_response = api_instance.dashboard_dashboard_api_get_widgets(ocs_api_request)
        print("The response of DashboardDashboardApiApi->dashboard_dashboard_api_get_widgets:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DashboardDashboardApiApi->dashboard_dashboard_api_get_widgets: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]

### Return type

[**DashboardDashboardApiGetWidgets200Response**](DashboardDashboardApiGetWidgets200Response.md)

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
