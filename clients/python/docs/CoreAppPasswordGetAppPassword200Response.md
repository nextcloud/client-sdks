# CoreAppPasswordGetAppPassword200Response


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ocs** | [**CoreAppPasswordGetAppPassword200ResponseOcs**](CoreAppPasswordGetAppPassword200ResponseOcs.md) |  | 

## Example

```python
from nextcloud_client.models.core_app_password_get_app_password200_response import CoreAppPasswordGetAppPassword200Response

# TODO update the JSON string below
json = "{}"
# create an instance of CoreAppPasswordGetAppPassword200Response from a JSON string
core_app_password_get_app_password200_response_instance = CoreAppPasswordGetAppPassword200Response.from_json(json)
# print the JSON string representation of the object
print CoreAppPasswordGetAppPassword200Response.to_json()

# convert the object into a dict
core_app_password_get_app_password200_response_dict = core_app_password_get_app_password200_response_instance.to_dict()
# create an instance of CoreAppPasswordGetAppPassword200Response from a dict
core_app_password_get_app_password200_response_form_dict = core_app_password_get_app_password200_response.from_dict(core_app_password_get_app_password200_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

