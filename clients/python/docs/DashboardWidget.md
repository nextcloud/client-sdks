# DashboardWidget


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **str** |  | 
**title** | **str** |  | 
**order** | **int** |  | 
**icon_class** | **str** |  | 
**icon_url** | **str** |  | 
**widget_url** | **str** |  | 
**item_icons_round** | **bool** |  | 
**item_api_versions** | **List[int]** |  | 
**reload_interval** | **int** |  | 
**buttons** | [**List[DashboardWidgetButtonsInner]**](DashboardWidgetButtonsInner.md) |  | [optional] 

## Example

```python
from nextcloud_client.models.dashboard_widget import DashboardWidget

# TODO update the JSON string below
json = "{}"
# create an instance of DashboardWidget from a JSON string
dashboard_widget_instance = DashboardWidget.from_json(json)
# print the JSON string representation of the object
print DashboardWidget.to_json()

# convert the object into a dict
dashboard_widget_dict = dashboard_widget_instance.to_dict()
# create an instance of DashboardWidget from a dict
dashboard_widget_form_dict = dashboard_widget.from_dict(dashboard_widget_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


