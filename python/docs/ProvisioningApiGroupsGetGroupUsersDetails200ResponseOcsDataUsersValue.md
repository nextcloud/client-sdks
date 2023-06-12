# ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**additional_mail** | **List[str]** |  | 
**additional_mail_scope** | **List[str]** |  | 
**address** | **str** |  | 
**address_scope** | **str** |  | 
**avatar_scope** | **str** |  | 
**backend** | **str** |  | 
**backend_capabilities** | [**ProvisioningApiUserDetailsBackendCapabilities**](ProvisioningApiUserDetailsBackendCapabilities.md) |  | 
**biography** | **str** |  | 
**biography_scope** | **str** |  | 
**displayname** | **str** |  | 
**display_name** | **str** |  | 
**displayname_scope** | **str** |  | 
**email** | **str** |  | 
**email_scope** | **str** |  | 
**enabled** | **bool** |  | 
**fediverse** | **str** |  | 
**fediverse_scope** | **str** |  | 
**groups** | **List[str]** |  | 
**headline** | **str** |  | 
**headline_scope** | **str** |  | 
**id** | **str** |  | 
**language** | **str** |  | 
**last_login** | **int** |  | 
**locale** | **str** |  | 
**notify_email** | **str** |  | 
**organisation** | **str** |  | 
**organisation_scope** | **str** |  | 
**phone** | **str** |  | 
**phone_scope** | **str** |  | 
**profile_enabled** | **str** |  | 
**profile_enabled_scope** | **str** |  | 
**quota** | [**ProvisioningApiUserDetailsQuota**](ProvisioningApiUserDetailsQuota.md) |  | 
**role** | **str** |  | 
**role_scope** | **str** |  | 
**storage_location** | **str** |  | 
**subadmin** | **List[str]** |  | 
**twitter** | **str** |  | 
**twitter_scope** | **str** |  | 
**website** | **str** |  | 
**website_scope** | **str** |  | 

## Example

```python
from openapi_client.models.provisioning_api_groups_get_group_users_details200_response_ocs_data_users_value import ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue

# TODO update the JSON string below
json = "{}"
# create an instance of ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue from a JSON string
provisioning_api_groups_get_group_users_details200_response_ocs_data_users_value_instance = ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue.from_json(json)
# print the JSON string representation of the object
print ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue.to_json()

# convert the object into a dict
provisioning_api_groups_get_group_users_details200_response_ocs_data_users_value_dict = provisioning_api_groups_get_group_users_details200_response_ocs_data_users_value_instance.to_dict()
# create an instance of ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue from a dict
provisioning_api_groups_get_group_users_details200_response_ocs_data_users_value_form_dict = provisioning_api_groups_get_group_users_details200_response_ocs_data_users_value.from_dict(provisioning_api_groups_get_group_users_details200_response_ocs_data_users_value_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


