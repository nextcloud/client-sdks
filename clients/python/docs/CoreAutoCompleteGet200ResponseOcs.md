# CoreAutoCompleteGet200ResponseOcs


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**meta** | [**OCSMeta**](OCSMeta.md) |  | 
**data** | [**List[CoreAutocompleteResult]**](CoreAutocompleteResult.md) |  | 

## Example

```python
from nextcloud_client.models.core_auto_complete_get200_response_ocs import CoreAutoCompleteGet200ResponseOcs

# TODO update the JSON string below
json = "{}"
# create an instance of CoreAutoCompleteGet200ResponseOcs from a JSON string
core_auto_complete_get200_response_ocs_instance = CoreAutoCompleteGet200ResponseOcs.from_json(json)
# print the JSON string representation of the object
print CoreAutoCompleteGet200ResponseOcs.to_json()

# convert the object into a dict
core_auto_complete_get200_response_ocs_dict = core_auto_complete_get200_response_ocs_instance.to_dict()
# create an instance of CoreAutoCompleteGet200ResponseOcs from a dict
core_auto_complete_get200_response_ocs_form_dict = core_auto_complete_get200_response_ocs.from_dict(core_auto_complete_get200_response_ocs_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


