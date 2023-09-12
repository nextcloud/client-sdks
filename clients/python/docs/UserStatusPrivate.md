# UserStatusPrivate


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
from nextcloud_client.models.user_status_private import UserStatusPrivate

# TODO update the JSON string below
json = "{}"
# create an instance of UserStatusPrivate from a JSON string
user_status_private_instance = UserStatusPrivate.from_json(json)
# print the JSON string representation of the object
print UserStatusPrivate.to_json()

# convert the object into a dict
user_status_private_dict = user_status_private_instance.to_dict()
# create an instance of UserStatusPrivate from a dict
user_status_private_form_dict = user_status_private.from_dict(user_status_private_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


