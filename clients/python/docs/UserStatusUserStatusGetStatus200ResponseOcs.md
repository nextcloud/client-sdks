# UserStatusUserStatusGetStatus200ResponseOcs


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**meta** | [**OCSMeta**](OCSMeta.md) |  | 
**data** | [**UserStatusPrivate**](UserStatusPrivate.md) |  | 

## Example

```python
from nextcloud_client.models.user_status_user_status_get_status200_response_ocs import UserStatusUserStatusGetStatus200ResponseOcs

# TODO update the JSON string below
json = "{}"
# create an instance of UserStatusUserStatusGetStatus200ResponseOcs from a JSON string
user_status_user_status_get_status200_response_ocs_instance = UserStatusUserStatusGetStatus200ResponseOcs.from_json(json)
# print the JSON string representation of the object
print UserStatusUserStatusGetStatus200ResponseOcs.to_json()

# convert the object into a dict
user_status_user_status_get_status200_response_ocs_dict = user_status_user_status_get_status200_response_ocs_instance.to_dict()
# create an instance of UserStatusUserStatusGetStatus200ResponseOcs from a dict
user_status_user_status_get_status200_response_ocs_form_dict = user_status_user_status_get_status200_response_ocs.from_dict(user_status_user_status_get_status200_response_ocs_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


