# ProvisioningApiAppsGetAppInfo200Response


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ocs** | [**ProvisioningApiAppsGetAppInfo200ResponseOcs**](ProvisioningApiAppsGetAppInfo200ResponseOcs.md) |  | 

## Example

```python
from nextcloud_client.models.provisioning_api_apps_get_app_info200_response import ProvisioningApiAppsGetAppInfo200Response

# TODO update the JSON string below
json = "{}"
# create an instance of ProvisioningApiAppsGetAppInfo200Response from a JSON string
provisioning_api_apps_get_app_info200_response_instance = ProvisioningApiAppsGetAppInfo200Response.from_json(json)
# print the JSON string representation of the object
print ProvisioningApiAppsGetAppInfo200Response.to_json()

# convert the object into a dict
provisioning_api_apps_get_app_info200_response_dict = provisioning_api_apps_get_app_info200_response_instance.to_dict()
# create an instance of ProvisioningApiAppsGetAppInfo200Response from a dict
provisioning_api_apps_get_app_info200_response_form_dict = provisioning_api_apps_get_app_info200_response.from_dict(provisioning_api_apps_get_app_info200_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


