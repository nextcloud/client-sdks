# FilesSharingCapabilitiesFilesSharingFederation


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**outgoing** | **bool** |  | 
**incoming** | **bool** |  | 
**expire_date** | [**FilesSharingCapabilitiesFilesSharingUserExpireDate**](FilesSharingCapabilitiesFilesSharingUserExpireDate.md) |  | 
**expire_date_supported** | [**FilesSharingCapabilitiesFilesSharingUserExpireDate**](FilesSharingCapabilitiesFilesSharingUserExpireDate.md) |  | 

## Example

```python
from nextcloud_client.models.files_sharing_capabilities_files_sharing_federation import FilesSharingCapabilitiesFilesSharingFederation

# TODO update the JSON string below
json = "{}"
# create an instance of FilesSharingCapabilitiesFilesSharingFederation from a JSON string
files_sharing_capabilities_files_sharing_federation_instance = FilesSharingCapabilitiesFilesSharingFederation.from_json(json)
# print the JSON string representation of the object
print FilesSharingCapabilitiesFilesSharingFederation.to_json()

# convert the object into a dict
files_sharing_capabilities_files_sharing_federation_dict = files_sharing_capabilities_files_sharing_federation_instance.to_dict()
# create an instance of FilesSharingCapabilitiesFilesSharingFederation from a dict
files_sharing_capabilities_files_sharing_federation_form_dict = files_sharing_capabilities_files_sharing_federation.from_dict(files_sharing_capabilities_files_sharing_federation_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


