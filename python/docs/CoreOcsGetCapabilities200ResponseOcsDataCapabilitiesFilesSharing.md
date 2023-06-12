# CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**api_enabled** | **bool** |  | 
**public** | [**CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingPublic**](CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingPublic.md) |  | 
**user** | [**CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUser**](CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUser.md) |  | 
**resharing** | **bool** |  | 
**group_sharing** | **bool** |  | [optional] 
**group** | [**CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingGroup**](CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingGroup.md) |  | [optional] 
**default_permissions** | **int** |  | [optional] 
**federation** | [**CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingFederation**](CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingFederation.md) |  | 
**sharee** | [**CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingSharee**](CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingSharee.md) |  | 

## Example

```python
from openapi_client.models.core_ocs_get_capabilities200_response_ocs_data_capabilities_files_sharing import CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing

# TODO update the JSON string below
json = "{}"
# create an instance of CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing from a JSON string
core_ocs_get_capabilities200_response_ocs_data_capabilities_files_sharing_instance = CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing.from_json(json)
# print the JSON string representation of the object
print CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing.to_json()

# convert the object into a dict
core_ocs_get_capabilities200_response_ocs_data_capabilities_files_sharing_dict = core_ocs_get_capabilities200_response_ocs_data_capabilities_files_sharing_instance.to_dict()
# create an instance of CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing from a dict
core_ocs_get_capabilities200_response_ocs_data_capabilities_files_sharing_form_dict = core_ocs_get_capabilities200_response_ocs_data_capabilities_files_sharing.from_dict(core_ocs_get_capabilities200_response_ocs_data_capabilities_files_sharing_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


