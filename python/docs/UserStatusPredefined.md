# UserStatusPredefined


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **str** |  | 
**icon** | **str** |  | 
**message** | **str** |  | 
**clear_at** | [**UserStatusClearAt**](UserStatusClearAt.md) |  | 
**visible** | **bool** |  | 

## Example

```python
from openapi_client.models.user_status_predefined import UserStatusPredefined

# TODO update the JSON string below
json = "{}"
# create an instance of UserStatusPredefined from a JSON string
user_status_predefined_instance = UserStatusPredefined.from_json(json)
# print the JSON string representation of the object
print UserStatusPredefined.to_json()

# convert the object into a dict
user_status_predefined_dict = user_status_predefined_instance.to_dict()
# create an instance of UserStatusPredefined from a dict
user_status_predefined_form_dict = user_status_predefined.from_dict(user_status_predefined_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


