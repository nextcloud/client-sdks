# UserStatusCapabilitiesUserStatus


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**enabled** | **bool** |  | 
**restore** | **bool** |  | 
**supports_emoji** | **bool** |  | 

## Example

```python
from nextcloud_client.models.user_status_capabilities_user_status import UserStatusCapabilitiesUserStatus

# TODO update the JSON string below
json = "{}"
# create an instance of UserStatusCapabilitiesUserStatus from a JSON string
user_status_capabilities_user_status_instance = UserStatusCapabilitiesUserStatus.from_json(json)
# print the JSON string representation of the object
print UserStatusCapabilitiesUserStatus.to_json()

# convert the object into a dict
user_status_capabilities_user_status_dict = user_status_capabilities_user_status_instance.to_dict()
# create an instance of UserStatusCapabilitiesUserStatus from a dict
user_status_capabilities_user_status_form_dict = user_status_capabilities_user_status.from_dict(user_status_capabilities_user_status_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


