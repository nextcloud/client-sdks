# CoreHoverCardGetUser200ResponseOcsData


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**user_id** | **str** |  | 
**display_name** | **str** |  | 
**actions** | [**List[CoreContactsAction]**](CoreContactsAction.md) |  | 

## Example

```python
from nextcloud_client.models.core_hover_card_get_user200_response_ocs_data import CoreHoverCardGetUser200ResponseOcsData

# TODO update the JSON string below
json = "{}"
# create an instance of CoreHoverCardGetUser200ResponseOcsData from a JSON string
core_hover_card_get_user200_response_ocs_data_instance = CoreHoverCardGetUser200ResponseOcsData.from_json(json)
# print the JSON string representation of the object
print CoreHoverCardGetUser200ResponseOcsData.to_json()

# convert the object into a dict
core_hover_card_get_user200_response_ocs_data_dict = core_hover_card_get_user200_response_ocs_data_instance.to_dict()
# create an instance of CoreHoverCardGetUser200ResponseOcsData from a dict
core_hover_card_get_user200_response_ocs_data_form_dict = core_hover_card_get_user200_response_ocs_data.from_dict(core_hover_card_get_user200_response_ocs_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


