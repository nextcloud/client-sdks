# CoreUnifiedSearchResult


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** |  | 
**is_paginated** | **bool** |  | 
**entries** | [**List[CoreUnifiedSearchResultEntry]**](CoreUnifiedSearchResultEntry.md) |  | 
**cursor** | [**CoreUnifiedSearchResultCursor**](CoreUnifiedSearchResultCursor.md) |  | 

## Example

```python
from nextcloud_client.models.core_unified_search_result import CoreUnifiedSearchResult

# TODO update the JSON string below
json = "{}"
# create an instance of CoreUnifiedSearchResult from a JSON string
core_unified_search_result_instance = CoreUnifiedSearchResult.from_json(json)
# print the JSON string representation of the object
print CoreUnifiedSearchResult.to_json()

# convert the object into a dict
core_unified_search_result_dict = core_unified_search_result_instance.to_dict()
# create an instance of CoreUnifiedSearchResult from a dict
core_unified_search_result_form_dict = core_unified_search_result.from_dict(core_unified_search_result_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


