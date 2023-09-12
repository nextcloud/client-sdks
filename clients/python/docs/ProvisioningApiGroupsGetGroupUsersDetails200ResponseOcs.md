# ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcs


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**meta** | [**OCSMeta**](OCSMeta.md) |  | 
**data** | [**ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsData**](ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsData.md) |  | 

## Example

```python
from nextcloud_client.models.provisioning_api_groups_get_group_users_details200_response_ocs import ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcs

# TODO update the JSON string below
json = "{}"
# create an instance of ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcs from a JSON string
provisioning_api_groups_get_group_users_details200_response_ocs_instance = ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcs.from_json(json)
# print the JSON string representation of the object
print ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcs.to_json()

# convert the object into a dict
provisioning_api_groups_get_group_users_details200_response_ocs_dict = provisioning_api_groups_get_group_users_details200_response_ocs_instance.to_dict()
# create an instance of ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcs from a dict
provisioning_api_groups_get_group_users_details200_response_ocs_form_dict = provisioning_api_groups_get_group_users_details200_response_ocs.from_dict(provisioning_api_groups_get_group_users_details200_response_ocs_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


