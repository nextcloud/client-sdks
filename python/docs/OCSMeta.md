# OCSMeta


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**status** | **str** |  | 
**statuscode** | **int** |  | 
**message** | **str** |  | [optional] 
**totalitems** | **str** |  | [optional] 
**itemsperpage** | **str** |  | [optional] 

## Example

```python
from openapi_client.models.ocs_meta import OCSMeta

# TODO update the JSON string below
json = "{}"
# create an instance of OCSMeta from a JSON string
ocs_meta_instance = OCSMeta.from_json(json)
# print the JSON string representation of the object
print OCSMeta.to_json()

# convert the object into a dict
ocs_meta_dict = ocs_meta_instance.to_dict()
# create an instance of OCSMeta from a dict
ocs_meta_form_dict = ocs_meta.from_dict(ocs_meta_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


