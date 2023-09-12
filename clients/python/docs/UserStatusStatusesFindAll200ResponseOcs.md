# UserStatusStatusesFindAll200ResponseOcs


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**meta** | [**OCSMeta**](OCSMeta.md) |  | 
**data** | [**List[UserStatusPublic]**](UserStatusPublic.md) |  | 

## Example

```python
from nextcloud_client.models.user_status_statuses_find_all200_response_ocs import UserStatusStatusesFindAll200ResponseOcs

# TODO update the JSON string below
json = "{}"
# create an instance of UserStatusStatusesFindAll200ResponseOcs from a JSON string
user_status_statuses_find_all200_response_ocs_instance = UserStatusStatusesFindAll200ResponseOcs.from_json(json)
# print the JSON string representation of the object
print UserStatusStatusesFindAll200ResponseOcs.to_json()

# convert the object into a dict
user_status_statuses_find_all200_response_ocs_dict = user_status_statuses_find_all200_response_ocs_instance.to_dict()
# create an instance of UserStatusStatusesFindAll200ResponseOcs from a dict
user_status_statuses_find_all200_response_ocs_form_dict = user_status_statuses_find_all200_response_ocs.from_dict(user_status_statuses_find_all200_response_ocs_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


