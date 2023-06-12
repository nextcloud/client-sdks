# UserStatusCapabilities


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**user_status** | [**CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesUserStatus**](CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesUserStatus.md) |  | 

## Example

```python
from openapi_client.models.user_status_capabilities import UserStatusCapabilities

# TODO update the JSON string below
json = "{}"
# create an instance of UserStatusCapabilities from a JSON string
user_status_capabilities_instance = UserStatusCapabilities.from_json(json)
# print the JSON string representation of the object
print UserStatusCapabilities.to_json()

# convert the object into a dict
user_status_capabilities_dict = user_status_capabilities_instance.to_dict()
# create an instance of UserStatusCapabilities from a dict
user_status_capabilities_form_dict = user_status_capabilities.from_dict(user_status_capabilities_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


