# CoreUnifiedSearchGetProviders200ResponseOcs


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**meta** | [**OCSMeta**](OCSMeta.md) |  | 
**data** | [**List[CoreUnifiedSearchProvider]**](CoreUnifiedSearchProvider.md) |  | 

## Example

```python
from openapi_client.models.core_unified_search_get_providers200_response_ocs import CoreUnifiedSearchGetProviders200ResponseOcs

# TODO update the JSON string below
json = "{}"
# create an instance of CoreUnifiedSearchGetProviders200ResponseOcs from a JSON string
core_unified_search_get_providers200_response_ocs_instance = CoreUnifiedSearchGetProviders200ResponseOcs.from_json(json)
# print the JSON string representation of the object
print CoreUnifiedSearchGetProviders200ResponseOcs.to_json()

# convert the object into a dict
core_unified_search_get_providers200_response_ocs_dict = core_unified_search_get_providers200_response_ocs_instance.to_dict()
# create an instance of CoreUnifiedSearchGetProviders200ResponseOcs from a dict
core_unified_search_get_providers200_response_ocs_form_dict = core_unified_search_get_providers200_response_ocs.from_dict(core_unified_search_get_providers200_response_ocs_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


