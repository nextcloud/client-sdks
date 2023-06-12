# ProvisioningApiUserDetailsQuota


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**free** | **int** |  | 
**quota** | [**ProvisioningApiUserDetailsQuotaQuota**](ProvisioningApiUserDetailsQuotaQuota.md) |  | 
**relative** | **float** |  | 
**total** | **int** |  | 
**used** | **int** |  | 

## Example

```python
from openapi_client.models.provisioning_api_user_details_quota import ProvisioningApiUserDetailsQuota

# TODO update the JSON string below
json = "{}"
# create an instance of ProvisioningApiUserDetailsQuota from a JSON string
provisioning_api_user_details_quota_instance = ProvisioningApiUserDetailsQuota.from_json(json)
# print the JSON string representation of the object
print ProvisioningApiUserDetailsQuota.to_json()

# convert the object into a dict
provisioning_api_user_details_quota_dict = provisioning_api_user_details_quota_instance.to_dict()
# create an instance of ProvisioningApiUserDetailsQuota from a dict
provisioning_api_user_details_quota_form_dict = provisioning_api_user_details_quota.from_dict(provisioning_api_user_details_quota_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


