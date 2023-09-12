# CoreReference


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**rich_object_type** | **str** |  | 
**rich_object** | **Dict[str, object]** |  | 
**open_graph_object** | [**CoreOpenGraphObject**](CoreOpenGraphObject.md) |  | 
**accessible** | **bool** |  | 

## Example

```python
from nextcloud_client.models.core_reference import CoreReference

# TODO update the JSON string below
json = "{}"
# create an instance of CoreReference from a JSON string
core_reference_instance = CoreReference.from_json(json)
# print the JSON string representation of the object
print CoreReference.to_json()

# convert the object into a dict
core_reference_dict = core_reference_instance.to_dict()
# create an instance of CoreReference from a dict
core_reference_form_dict = core_reference.from_dict(core_reference_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


