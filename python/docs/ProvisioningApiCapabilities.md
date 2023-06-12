# ProvisioningApiCapabilities


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**provisioning_api** | [**CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesProvisioningApi**](CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesProvisioningApi.md) |  | 

## Example

```python
from openapi_client.models.provisioning_api_capabilities import ProvisioningApiCapabilities

# TODO update the JSON string below
json = "{}"
# create an instance of ProvisioningApiCapabilities from a JSON string
provisioning_api_capabilities_instance = ProvisioningApiCapabilities.from_json(json)
# print the JSON string representation of the object
print ProvisioningApiCapabilities.to_json()

# convert the object into a dict
provisioning_api_capabilities_dict = provisioning_api_capabilities_instance.to_dict()
# create an instance of ProvisioningApiCapabilities from a dict
provisioning_api_capabilities_form_dict = provisioning_api_capabilities.from_dict(provisioning_api_capabilities_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


