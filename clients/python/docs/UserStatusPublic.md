# UserStatusPublic


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**user_id** | **str** |  | 
**message** | **str** |  | 
**icon** | **str** |  | 
**clear_at** | **int** |  | 
**status** | **str** |  | 

## Example

```python
from nextcloud_client.models.user_status_public import UserStatusPublic

# TODO update the JSON string below
json = "{}"
# create an instance of UserStatusPublic from a JSON string
user_status_public_instance = UserStatusPublic.from_json(json)
# print the JSON string representation of the object
print UserStatusPublic.to_json()

# convert the object into a dict
user_status_public_dict = user_status_public_instance.to_dict()
# create an instance of UserStatusPublic from a dict
user_status_public_form_dict = user_status_public.from_dict(user_status_public_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


