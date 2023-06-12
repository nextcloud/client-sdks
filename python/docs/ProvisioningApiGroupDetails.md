# ProvisioningApiGroupDetails


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **str** |  | 
**displayname** | **str** |  | 
**usercount** | [**ProvisioningApiGroupDetailsUsercount**](ProvisioningApiGroupDetailsUsercount.md) |  | 
**disabled** | [**ProvisioningApiGroupDetailsUsercount**](ProvisioningApiGroupDetailsUsercount.md) |  | 
**can_add** | **bool** |  | 
**can_remove** | **bool** |  | 

## Example

```python
from openapi_client.models.provisioning_api_group_details import ProvisioningApiGroupDetails

# TODO update the JSON string below
json = "{}"
# create an instance of ProvisioningApiGroupDetails from a JSON string
provisioning_api_group_details_instance = ProvisioningApiGroupDetails.from_json(json)
# print the JSON string representation of the object
print ProvisioningApiGroupDetails.to_json()

# convert the object into a dict
provisioning_api_group_details_dict = provisioning_api_group_details_instance.to_dict()
# create an instance of ProvisioningApiGroupDetails from a dict
provisioning_api_group_details_form_dict = provisioning_api_group_details.from_dict(provisioning_api_group_details_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


