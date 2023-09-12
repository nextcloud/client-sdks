# UserStatusClearAt


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**type** | **str** |  | 
**time** | [**UserStatusClearAtTime**](UserStatusClearAtTime.md) |  | 

## Example

```python
from nextcloud_client.models.user_status_clear_at import UserStatusClearAt

# TODO update the JSON string below
json = "{}"
# create an instance of UserStatusClearAt from a JSON string
user_status_clear_at_instance = UserStatusClearAt.from_json(json)
# print the JSON string representation of the object
print UserStatusClearAt.to_json()

# convert the object into a dict
user_status_clear_at_dict = user_status_clear_at_instance.to_dict()
# create an instance of UserStatusClearAt from a dict
user_status_clear_at_form_dict = user_status_clear_at.from_dict(user_status_clear_at_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


