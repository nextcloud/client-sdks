# FilesSharingCapabilitiesFilesSharingPublic


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**enabled** | **bool** |  | 
**password** | [**FilesSharingCapabilitiesFilesSharingPublicPassword**](FilesSharingCapabilitiesFilesSharingPublicPassword.md) |  | [optional] 
**multiple_links** | **bool** |  | [optional] 
**expire_date** | [**FilesSharingCapabilitiesFilesSharingPublicExpireDate**](FilesSharingCapabilitiesFilesSharingPublicExpireDate.md) |  | [optional] 
**expire_date_internal** | [**FilesSharingCapabilitiesFilesSharingPublicExpireDate**](FilesSharingCapabilitiesFilesSharingPublicExpireDate.md) |  | [optional] 
**expire_date_remote** | [**FilesSharingCapabilitiesFilesSharingPublicExpireDate**](FilesSharingCapabilitiesFilesSharingPublicExpireDate.md) |  | [optional] 
**send_mail** | **bool** |  | [optional] 
**upload** | **bool** |  | [optional] 
**upload_files_drop** | **bool** |  | [optional] 

## Example

```python
from nextcloud_client.models.files_sharing_capabilities_files_sharing_public import FilesSharingCapabilitiesFilesSharingPublic

# TODO update the JSON string below
json = "{}"
# create an instance of FilesSharingCapabilitiesFilesSharingPublic from a JSON string
files_sharing_capabilities_files_sharing_public_instance = FilesSharingCapabilitiesFilesSharingPublic.from_json(json)
# print the JSON string representation of the object
print FilesSharingCapabilitiesFilesSharingPublic.to_json()

# convert the object into a dict
files_sharing_capabilities_files_sharing_public_dict = files_sharing_capabilities_files_sharing_public_instance.to_dict()
# create an instance of FilesSharingCapabilitiesFilesSharingPublic from a dict
files_sharing_capabilities_files_sharing_public_form_dict = files_sharing_capabilities_files_sharing_public.from_dict(files_sharing_capabilities_files_sharing_public_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


