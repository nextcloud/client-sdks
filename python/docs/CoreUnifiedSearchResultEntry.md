# CoreUnifiedSearchResultEntry


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**thumbnail_url** | **str** |  | 
**title** | **str** |  | 
**subline** | **str** |  | 
**resource_url** | **str** |  | 
**icon** | **str** |  | 
**rounded** | **bool** |  | 
**attributes** | **List[str]** |  | 

## Example

```python
from openapi_client.models.core_unified_search_result_entry import CoreUnifiedSearchResultEntry

# TODO update the JSON string below
json = "{}"
# create an instance of CoreUnifiedSearchResultEntry from a JSON string
core_unified_search_result_entry_instance = CoreUnifiedSearchResultEntry.from_json(json)
# print the JSON string representation of the object
print CoreUnifiedSearchResultEntry.to_json()

# convert the object into a dict
core_unified_search_result_entry_dict = core_unified_search_result_entry_instance.to_dict()
# create an instance of CoreUnifiedSearchResultEntry from a dict
core_unified_search_result_entry_form_dict = core_unified_search_result_entry.from_dict(core_unified_search_result_entry_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


