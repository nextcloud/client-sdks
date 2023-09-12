# CoreUnifiedSearchProvider


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **str** |  | 
**name** | **str** |  | 
**order** | **int** |  | 

## Example

```python
from nextcloud_client.models.core_unified_search_provider import CoreUnifiedSearchProvider

# TODO update the JSON string below
json = "{}"
# create an instance of CoreUnifiedSearchProvider from a JSON string
core_unified_search_provider_instance = CoreUnifiedSearchProvider.from_json(json)
# print the JSON string representation of the object
print CoreUnifiedSearchProvider.to_json()

# convert the object into a dict
core_unified_search_provider_dict = core_unified_search_provider_instance.to_dict()
# create an instance of CoreUnifiedSearchProvider from a dict
core_unified_search_provider_form_dict = core_unified_search_provider.from_dict(core_unified_search_provider_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


