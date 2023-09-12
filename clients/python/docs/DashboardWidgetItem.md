# DashboardWidgetItem


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**subtitle** | **str** |  | 
**title** | **str** |  | 
**link** | **str** |  | 
**icon_url** | **str** |  | 
**overlay_icon_url** | **str** |  | 
**since_id** | **str** |  | 

## Example

```python
from nextcloud_client.models.dashboard_widget_item import DashboardWidgetItem

# TODO update the JSON string below
json = "{}"
# create an instance of DashboardWidgetItem from a JSON string
dashboard_widget_item_instance = DashboardWidgetItem.from_json(json)
# print the JSON string representation of the object
print DashboardWidgetItem.to_json()

# convert the object into a dict
dashboard_widget_item_dict = dashboard_widget_item_instance.to_dict()
# create an instance of DashboardWidgetItem from a dict
dashboard_widget_item_form_dict = dashboard_widget_item.from_dict(dashboard_widget_item_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


