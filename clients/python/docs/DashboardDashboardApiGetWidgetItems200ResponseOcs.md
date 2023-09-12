# DashboardDashboardApiGetWidgetItems200ResponseOcs


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**meta** | [**OCSMeta**](OCSMeta.md) |  | 
**data** | **Dict[str, List[DashboardWidgetItem]]** |  | 

## Example

```python
from nextcloud_client.models.dashboard_dashboard_api_get_widget_items200_response_ocs import DashboardDashboardApiGetWidgetItems200ResponseOcs

# TODO update the JSON string below
json = "{}"
# create an instance of DashboardDashboardApiGetWidgetItems200ResponseOcs from a JSON string
dashboard_dashboard_api_get_widget_items200_response_ocs_instance = DashboardDashboardApiGetWidgetItems200ResponseOcs.from_json(json)
# print the JSON string representation of the object
print DashboardDashboardApiGetWidgetItems200ResponseOcs.to_json()

# convert the object into a dict
dashboard_dashboard_api_get_widget_items200_response_ocs_dict = dashboard_dashboard_api_get_widget_items200_response_ocs_instance.to_dict()
# create an instance of DashboardDashboardApiGetWidgetItems200ResponseOcs from a dict
dashboard_dashboard_api_get_widget_items200_response_ocs_form_dict = dashboard_dashboard_api_get_widget_items200_response_ocs.from_dict(dashboard_dashboard_api_get_widget_items200_response_ocs_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

