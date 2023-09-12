# CoreAppPasswordGetAppPassword200ResponseOcs


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**meta** | [**OCSMeta**](OCSMeta.md) |  | 
**data** | [**CoreAppPasswordGetAppPassword200ResponseOcsData**](CoreAppPasswordGetAppPassword200ResponseOcsData.md) |  | 

## Example

```python
from nextcloud_client.models.core_app_password_get_app_password200_response_ocs import CoreAppPasswordGetAppPassword200ResponseOcs

# TODO update the JSON string below
json = "{}"
# create an instance of CoreAppPasswordGetAppPassword200ResponseOcs from a JSON string
core_app_password_get_app_password200_response_ocs_instance = CoreAppPasswordGetAppPassword200ResponseOcs.from_json(json)
# print the JSON string representation of the object
print CoreAppPasswordGetAppPassword200ResponseOcs.to_json()

# convert the object into a dict
core_app_password_get_app_password200_response_ocs_dict = core_app_password_get_app_password200_response_ocs_instance.to_dict()
# create an instance of CoreAppPasswordGetAppPassword200ResponseOcs from a dict
core_app_password_get_app_password200_response_ocs_form_dict = core_app_password_get_app_password200_response_ocs.from_dict(core_app_password_get_app_password200_response_ocs_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


