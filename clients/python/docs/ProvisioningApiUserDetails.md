# ProvisioningApiUserDetails


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**additional_mail** | **List[str]** |  | 
**additional_mail_scope** | **List[str]** |  | [optional] 
**address** | **str** |  | 
**address_scope** | **str** |  | [optional] 
**avatar_scope** | **str** |  | [optional] 
**backend** | **str** |  | 
**backend_capabilities** | [**ProvisioningApiUserDetailsBackendCapabilities**](ProvisioningApiUserDetailsBackendCapabilities.md) |  | 
**biography** | **str** |  | 
**biography_scope** | **str** |  | [optional] 
**display_name** | **str** |  | 
**displayname** | **str** |  | 
**displayname_scope** | **str** |  | [optional] 
**email** | **str** |  | 
**email_scope** | **str** |  | [optional] 
**enabled** | **bool** |  | [optional] 
**fediverse** | **str** |  | 
**fediverse_scope** | **str** |  | [optional] 
**groups** | **List[str]** |  | 
**headline** | **str** |  | 
**headline_scope** | **str** |  | [optional] 
**id** | **str** |  | 
**language** | **str** |  | 
**last_login** | **int** |  | 
**locale** | **str** |  | 
**manager** | **str** |  | 
**notify_email** | **str** |  | 
**organisation** | **str** |  | 
**organisation_scope** | **str** |  | [optional] 
**phone** | **str** |  | 
**phone_scope** | **str** |  | [optional] 
**profile_enabled** | **str** |  | 
**profile_enabled_scope** | **str** |  | [optional] 
**quota** | [**ProvisioningApiUserDetailsQuota**](ProvisioningApiUserDetailsQuota.md) |  | 
**role** | **str** |  | 
**role_scope** | **str** |  | [optional] 
**storage_location** | **str** |  | [optional] 
**subadmin** | **List[str]** |  | 
**twitter** | **str** |  | 
**twitter_scope** | **str** |  | [optional] 
**website** | **str** |  | 
**website_scope** | **str** |  | [optional] 

## Example

```python
from nextcloud_client.models.provisioning_api_user_details import ProvisioningApiUserDetails

# TODO update the JSON string below
json = "{}"
# create an instance of ProvisioningApiUserDetails from a JSON string
provisioning_api_user_details_instance = ProvisioningApiUserDetails.from_json(json)
# print the JSON string representation of the object
print ProvisioningApiUserDetails.to_json()

# convert the object into a dict
provisioning_api_user_details_dict = provisioning_api_user_details_instance.to_dict()
# create an instance of ProvisioningApiUserDetails from a dict
provisioning_api_user_details_form_dict = provisioning_api_user_details.from_dict(provisioning_api_user_details_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


