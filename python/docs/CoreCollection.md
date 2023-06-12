# CoreCollection


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **str** |  | 
**name** | **str** |  | 
**resources** | **List[Dict[str, object]]** |  | 

## Example

```python
from openapi_client.models.core_collection import CoreCollection

# TODO update the JSON string below
json = "{}"
# create an instance of CoreCollection from a JSON string
core_collection_instance = CoreCollection.from_json(json)
# print the JSON string representation of the object
print CoreCollection.to_json()

# convert the object into a dict
core_collection_dict = core_collection_instance.to_dict()
# create an instance of CoreCollection from a dict
core_collection_form_dict = core_collection.from_dict(core_collection_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


