# CoreNavigationEntry


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **str** |  | 
**order** | [**CoreNavigationEntryOrder**](CoreNavigationEntryOrder.md) |  | 
**href** | **str** |  | 
**icon** | **str** |  | 
**type** | **str** |  | 
**name** | **str** |  | 
**active** | **bool** |  | 
**classes** | **str** |  | 
**unread** | **int** |  | 

## Example

```python
from nextcloud_client.models.core_navigation_entry import CoreNavigationEntry

# TODO update the JSON string below
json = "{}"
# create an instance of CoreNavigationEntry from a JSON string
core_navigation_entry_instance = CoreNavigationEntry.from_json(json)
# print the JSON string representation of the object
print CoreNavigationEntry.to_json()

# convert the object into a dict
core_navigation_entry_dict = core_navigation_entry_instance.to_dict()
# create an instance of CoreNavigationEntry from a dict
core_navigation_entry_form_dict = core_navigation_entry.from_dict(core_navigation_entry_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


