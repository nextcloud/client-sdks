# CoreResource


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**rich_object_type** | **str** |  | 
**rich_object** | **Dict[str, object]** |  | 
**open_graph_object** | [**CoreOpenGraphObject**](CoreOpenGraphObject.md) |  | 
**accessible** | **bool** |  | 

## Example

```python
from nextcloud_client.models.core_resource import CoreResource

# TODO update the JSON string below
json = "{}"
# create an instance of CoreResource from a JSON string
core_resource_instance = CoreResource.from_json(json)
# print the JSON string representation of the object
print CoreResource.to_json()

# convert the object into a dict
core_resource_dict = core_resource_instance.to_dict()
# create an instance of CoreResource from a dict
core_resource_form_dict = core_resource.from_dict(core_resource_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


