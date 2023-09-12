# UserStatusUserStatusRevertStatus200ResponseOcsData


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**user_id** | **str** |  | 
**message** | **str** |  | 
**icon** | **str** |  | 
**clear_at** | **int** |  | 
**status** | **str** |  | 
**message_id** | **str** |  | 
**message_is_predefined** | **bool** |  | 
**status_is_user_defined** | **bool** |  | 

## Example

```python
from nextcloud_client.models.user_status_user_status_revert_status200_response_ocs_data import UserStatusUserStatusRevertStatus200ResponseOcsData

# TODO update the JSON string below
json = "{}"
# create an instance of UserStatusUserStatusRevertStatus200ResponseOcsData from a JSON string
user_status_user_status_revert_status200_response_ocs_data_instance = UserStatusUserStatusRevertStatus200ResponseOcsData.from_json(json)
# print the JSON string representation of the object
print UserStatusUserStatusRevertStatus200ResponseOcsData.to_json()

# convert the object into a dict
user_status_user_status_revert_status200_response_ocs_data_dict = user_status_user_status_revert_status200_response_ocs_data_instance.to_dict()
# create an instance of UserStatusUserStatusRevertStatus200ResponseOcsData from a dict
user_status_user_status_revert_status200_response_ocs_data_form_dict = user_status_user_status_revert_status200_response_ocs_data.from_dict(user_status_user_status_revert_status200_response_ocs_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


