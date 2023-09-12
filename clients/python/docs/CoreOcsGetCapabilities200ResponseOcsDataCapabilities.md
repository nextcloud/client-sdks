# CoreOcsGetCapabilities200ResponseOcsDataCapabilities


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**files** | [**FilesVersionsCapabilitiesFiles**](FilesVersionsCapabilitiesFiles.md) |  | 
**dav** | [**DavCapabilitiesDav**](DavCapabilitiesDav.md) |  | 
**files_sharing** | [**SharebymailCapabilitiesFilesSharing**](SharebymailCapabilitiesFilesSharing.md) |  | 
**provisioning_api** | [**ProvisioningApiCapabilitiesProvisioningApi**](ProvisioningApiCapabilitiesProvisioningApi.md) |  | 
**theming** | [**ThemingPublicCapabilitiesTheming**](ThemingPublicCapabilitiesTheming.md) |  | 
**user_status** | [**UserStatusCapabilitiesUserStatus**](UserStatusCapabilitiesUserStatus.md) |  | 
**weather_status** | [**FilesSharingCapabilitiesFilesSharingUserExpireDate**](FilesSharingCapabilitiesFilesSharingUserExpireDate.md) |  | 

## Example

```python
from nextcloud_client.models.core_ocs_get_capabilities200_response_ocs_data_capabilities import CoreOcsGetCapabilities200ResponseOcsDataCapabilities

# TODO update the JSON string below
json = "{}"
# create an instance of CoreOcsGetCapabilities200ResponseOcsDataCapabilities from a JSON string
core_ocs_get_capabilities200_response_ocs_data_capabilities_instance = CoreOcsGetCapabilities200ResponseOcsDataCapabilities.from_json(json)
# print the JSON string representation of the object
print CoreOcsGetCapabilities200ResponseOcsDataCapabilities.to_json()

# convert the object into a dict
core_ocs_get_capabilities200_response_ocs_data_capabilities_dict = core_ocs_get_capabilities200_response_ocs_data_capabilities_instance.to_dict()
# create an instance of CoreOcsGetCapabilities200ResponseOcsDataCapabilities from a dict
core_ocs_get_capabilities200_response_ocs_data_capabilities_form_dict = core_ocs_get_capabilities200_response_ocs_data_capabilities.from_dict(core_ocs_get_capabilities200_response_ocs_data_capabilities_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


