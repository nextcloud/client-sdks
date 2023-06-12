# ProvisioningApiUsersAddUser200ResponseOcs


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**meta** | [**OCSMeta**](OCSMeta.md) |  | 
**data** | [**ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValueOneOf**](ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValueOneOf.md) |  | 

## Example

```python
from openapi_client.models.provisioning_api_users_add_user200_response_ocs import ProvisioningApiUsersAddUser200ResponseOcs

# TODO update the JSON string below
json = "{}"
# create an instance of ProvisioningApiUsersAddUser200ResponseOcs from a JSON string
provisioning_api_users_add_user200_response_ocs_instance = ProvisioningApiUsersAddUser200ResponseOcs.from_json(json)
# print the JSON string representation of the object
print ProvisioningApiUsersAddUser200ResponseOcs.to_json()

# convert the object into a dict
provisioning_api_users_add_user200_response_ocs_dict = provisioning_api_users_add_user200_response_ocs_instance.to_dict()
# create an instance of ProvisioningApiUsersAddUser200ResponseOcs from a dict
provisioning_api_users_add_user200_response_ocs_form_dict = provisioning_api_users_add_user200_response_ocs.from_dict(provisioning_api_users_add_user200_response_ocs_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


