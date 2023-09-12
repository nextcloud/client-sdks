# DashboardWidgetItems


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**items** | [**List[DashboardWidgetItem]**](DashboardWidgetItem.md) |  | 
**empty_content_message** | **str** |  | 
**half_empty_content_message** | **str** |  | 

## Example

```python
from nextcloud_client.models.dashboard_widget_items import DashboardWidgetItems

# TODO update the JSON string below
json = "{}"
# create an instance of DashboardWidgetItems from a JSON string
dashboard_widget_items_instance = DashboardWidgetItems.from_json(json)
# print the JSON string representation of the object
print DashboardWidgetItems.to_json()

# convert the object into a dict
dashboard_widget_items_dict = dashboard_widget_items_instance.to_dict()
# create an instance of DashboardWidgetItems from a dict
dashboard_widget_items_form_dict = dashboard_widget_items.from_dict(dashboard_widget_items_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


