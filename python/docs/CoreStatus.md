# CoreStatus


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**installed** | **bool** |  | 
**maintenance** | **bool** |  | 
**needs_db_upgrade** | **bool** |  | 
**version** | **str** |  | 
**versionstring** | **str** |  | 
**edition** | **str** |  | 
**productname** | **str** |  | 
**extended_support** | **bool** |  | 

## Example

```python
from openapi_client.models.core_status import CoreStatus

# TODO update the JSON string below
json = "{}"
# create an instance of CoreStatus from a JSON string
core_status_instance = CoreStatus.from_json(json)
# print the JSON string representation of the object
print CoreStatus.to_json()

# convert the object into a dict
core_status_dict = core_status_instance.to_dict()
# create an instance of CoreStatus from a dict
core_status_form_dict = core_status.from_dict(core_status_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


