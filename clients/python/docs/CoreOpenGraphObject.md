# CoreOpenGraphObject


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **str** |  | 
**name** | **str** |  | 
**description** | **str** |  | 
**thumb** | **str** |  | 
**link** | **str** |  | 

## Example

```python
from nextcloud_client.models.core_open_graph_object import CoreOpenGraphObject

# TODO update the JSON string below
json = "{}"
# create an instance of CoreOpenGraphObject from a JSON string
core_open_graph_object_instance = CoreOpenGraphObject.from_json(json)
# print the JSON string representation of the object
print CoreOpenGraphObject.to_json()

# convert the object into a dict
core_open_graph_object_dict = core_open_graph_object_instance.to_dict()
# create an instance of CoreOpenGraphObject from a dict
core_open_graph_object_form_dict = core_open_graph_object.from_dict(core_open_graph_object_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


