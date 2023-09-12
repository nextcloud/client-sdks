# ProvisioningApiGroupsGetGroupsDetails200ResponseOcs


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**meta** | [**OCSMeta**](OCSMeta.md) |  | 
**data** | [**ProvisioningApiGroupsGetGroupsDetails200ResponseOcsData**](ProvisioningApiGroupsGetGroupsDetails200ResponseOcsData.md) |  | 

## Example

```python
from nextcloud_client.models.provisioning_api_groups_get_groups_details200_response_ocs import ProvisioningApiGroupsGetGroupsDetails200ResponseOcs

# TODO update the JSON string below
json = "{}"
# create an instance of ProvisioningApiGroupsGetGroupsDetails200ResponseOcs from a JSON string
provisioning_api_groups_get_groups_details200_response_ocs_instance = ProvisioningApiGroupsGetGroupsDetails200ResponseOcs.from_json(json)
# print the JSON string representation of the object
print ProvisioningApiGroupsGetGroupsDetails200ResponseOcs.to_json()

# convert the object into a dict
provisioning_api_groups_get_groups_details200_response_ocs_dict = provisioning_api_groups_get_groups_details200_response_ocs_instance.to_dict()
# create an instance of ProvisioningApiGroupsGetGroupsDetails200ResponseOcs from a dict
provisioning_api_groups_get_groups_details200_response_ocs_form_dict = provisioning_api_groups_get_groups_details200_response_ocs.from_dict(provisioning_api_groups_get_groups_details200_response_ocs_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


