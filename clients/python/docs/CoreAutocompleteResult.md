# CoreAutocompleteResult


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **str** |  | 
**label** | **str** |  | 
**icon** | **str** |  | 
**source** | **str** |  | 
**status** | **str** |  | 
**subline** | **str** |  | 
**share_with_display_name_unique** | **str** |  | 

## Example

```python
from nextcloud_client.models.core_autocomplete_result import CoreAutocompleteResult

# TODO update the JSON string below
json = "{}"
# create an instance of CoreAutocompleteResult from a JSON string
core_autocomplete_result_instance = CoreAutocompleteResult.from_json(json)
# print the JSON string representation of the object
print CoreAutocompleteResult.to_json()

# convert the object into a dict
core_autocomplete_result_dict = core_autocomplete_result_instance.to_dict()
# create an instance of CoreAutocompleteResult from a dict
core_autocomplete_result_form_dict = core_autocomplete_result.from_dict(core_autocomplete_result_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


