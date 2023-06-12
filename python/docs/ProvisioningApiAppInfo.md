# ProvisioningApiAppInfo


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**active** | **bool** |  | 
**activity** | **object** |  | 
**author** | **object** |  | 
**background_jobs** | **object** |  | 
**bugs** | **object** |  | 
**category** | **object** |  | 
**collaboration** | **object** |  | 
**commands** | **object** |  | 
**default_enable** | **object** |  | 
**dependencies** | **object** |  | 
**description** | **str** |  | 
**discussion** | **object** |  | 
**documentation** | **object** |  | 
**groups** | **object** |  | 
**id** | **str** |  | 
**info** | **object** |  | 
**internal** | **bool** |  | 
**level** | **int** |  | 
**licence** | **object** |  | 
**name** | **str** |  | 
**namespace** | **object** |  | 
**navigations** | **object** |  | 
**preview** | **object** |  | 
**preview_as_icon** | **bool** |  | 
**public** | **object** |  | 
**remote** | **object** |  | 
**removable** | **bool** |  | 
**repair_steps** | **object** |  | 
**repository** | **object** |  | 
**sabre** | **object** |  | 
**screenshot** | **object** |  | 
**settings** | **object** |  | 
**summary** | **str** |  | 
**trash** | **object** |  | 
**two_factor_providers** | **object** |  | 
**types** | **object** |  | 
**version** | **str** |  | 
**versions** | **object** |  | 
**website** | **object** |  | 

## Example

```python
from openapi_client.models.provisioning_api_app_info import ProvisioningApiAppInfo

# TODO update the JSON string below
json = "{}"
# create an instance of ProvisioningApiAppInfo from a JSON string
provisioning_api_app_info_instance = ProvisioningApiAppInfo.from_json(json)
# print the JSON string representation of the object
print ProvisioningApiAppInfo.to_json()

# convert the object into a dict
provisioning_api_app_info_dict = provisioning_api_app_info_instance.to_dict()
# create an instance of ProvisioningApiAppInfo from a dict
provisioning_api_app_info_form_dict = provisioning_api_app_info.from_dict(provisioning_api_app_info_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


