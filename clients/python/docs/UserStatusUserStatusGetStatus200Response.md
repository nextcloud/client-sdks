# UserStatusUserStatusGetStatus200Response


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ocs** | [**UserStatusUserStatusGetStatus200ResponseOcs**](UserStatusUserStatusGetStatus200ResponseOcs.md) |  | 

## Example

```python
from nextcloud_client.models.user_status_user_status_get_status200_response import UserStatusUserStatusGetStatus200Response

# TODO update the JSON string below
json = "{}"
# create an instance of UserStatusUserStatusGetStatus200Response from a JSON string
user_status_user_status_get_status200_response_instance = UserStatusUserStatusGetStatus200Response.from_json(json)
# print the JSON string representation of the object
print UserStatusUserStatusGetStatus200Response.to_json()

# convert the object into a dict
user_status_user_status_get_status200_response_dict = user_status_user_status_get_status200_response_instance.to_dict()
# create an instance of UserStatusUserStatusGetStatus200Response from a dict
user_status_user_status_get_status200_response_form_dict = user_status_user_status_get_status200_response.from_dict(user_status_user_status_get_status200_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


