# FilesSharingCapabilitiesFilesSharing


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**api_enabled** | **bool** |  | 
**public** | [**FilesSharingCapabilitiesFilesSharingPublic**](FilesSharingCapabilitiesFilesSharingPublic.md) |  | 
**user** | [**FilesSharingCapabilitiesFilesSharingUser**](FilesSharingCapabilitiesFilesSharingUser.md) |  | 
**resharing** | **bool** |  | 
**group_sharing** | **bool** |  | [optional] 
**group** | [**FilesSharingCapabilitiesFilesSharingGroup**](FilesSharingCapabilitiesFilesSharingGroup.md) |  | [optional] 
**default_permissions** | **int** |  | [optional] 
**federation** | [**FilesSharingCapabilitiesFilesSharingFederation**](FilesSharingCapabilitiesFilesSharingFederation.md) |  | 
**sharee** | [**FilesSharingCapabilitiesFilesSharingSharee**](FilesSharingCapabilitiesFilesSharingSharee.md) |  | 

## Example

```python
from nextcloud_client.models.files_sharing_capabilities_files_sharing import FilesSharingCapabilitiesFilesSharing

# TODO update the JSON string below
json = "{}"
# create an instance of FilesSharingCapabilitiesFilesSharing from a JSON string
files_sharing_capabilities_files_sharing_instance = FilesSharingCapabilitiesFilesSharing.from_json(json)
# print the JSON string representation of the object
print FilesSharingCapabilitiesFilesSharing.to_json()

# convert the object into a dict
files_sharing_capabilities_files_sharing_dict = files_sharing_capabilities_files_sharing_instance.to_dict()
# create an instance of FilesSharingCapabilitiesFilesSharing from a dict
files_sharing_capabilities_files_sharing_form_dict = files_sharing_capabilities_files_sharing.from_dict(files_sharing_capabilities_files_sharing_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


