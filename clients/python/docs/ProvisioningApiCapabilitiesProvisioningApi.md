# ProvisioningApiCapabilitiesProvisioningApi


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**version** | **str** |  | 
**account_property_scopes_version** | **int** |  | 
**account_property_scopes_federated_enabled** | **bool** |  | 
**account_property_scopes_published_enabled** | **bool** |  | 

## Example

```python
from nextcloud_client.models.provisioning_api_capabilities_provisioning_api import ProvisioningApiCapabilitiesProvisioningApi

# TODO update the JSON string below
json = "{}"
# create an instance of ProvisioningApiCapabilitiesProvisioningApi from a JSON string
provisioning_api_capabilities_provisioning_api_instance = ProvisioningApiCapabilitiesProvisioningApi.from_json(json)
# print the JSON string representation of the object
print ProvisioningApiCapabilitiesProvisioningApi.to_json()

# convert the object into a dict
provisioning_api_capabilities_provisioning_api_dict = provisioning_api_capabilities_provisioning_api_instance.to_dict()
# create an instance of ProvisioningApiCapabilitiesProvisioningApi from a dict
provisioning_api_capabilities_provisioning_api_form_dict = provisioning_api_capabilities_provisioning_api.from_dict(provisioning_api_capabilities_provisioning_api_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


