# ProvisioningApiUserDetailsBackendCapabilities


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**set_display_name** | **bool** |  | 
**set_password** | **bool** |  | 

## Example

```python
from nextcloud_client.models.provisioning_api_user_details_backend_capabilities import ProvisioningApiUserDetailsBackendCapabilities

# TODO update the JSON string below
json = "{}"
# create an instance of ProvisioningApiUserDetailsBackendCapabilities from a JSON string
provisioning_api_user_details_backend_capabilities_instance = ProvisioningApiUserDetailsBackendCapabilities.from_json(json)
# print the JSON string representation of the object
print ProvisioningApiUserDetailsBackendCapabilities.to_json()

# convert the object into a dict
provisioning_api_user_details_backend_capabilities_dict = provisioning_api_user_details_backend_capabilities_instance.to_dict()
# create an instance of ProvisioningApiUserDetailsBackendCapabilities from a dict
provisioning_api_user_details_backend_capabilities_form_dict = provisioning_api_user_details_backend_capabilities.from_dict(provisioning_api_user_details_backend_capabilities_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


