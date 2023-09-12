# CoreReferenceProvider


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **str** |  | 
**title** | **str** |  | 
**icon_url** | **str** |  | 
**order** | **int** |  | 
**search_providers_ids** | **List[str]** |  | 

## Example

```python
from nextcloud_client.models.core_reference_provider import CoreReferenceProvider

# TODO update the JSON string below
json = "{}"
# create an instance of CoreReferenceProvider from a JSON string
core_reference_provider_instance = CoreReferenceProvider.from_json(json)
# print the JSON string representation of the object
print CoreReferenceProvider.to_json()

# convert the object into a dict
core_reference_provider_dict = core_reference_provider_instance.to_dict()
# create an instance of CoreReferenceProvider from a dict
core_reference_provider_form_dict = core_reference_provider.from_dict(core_reference_provider_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


