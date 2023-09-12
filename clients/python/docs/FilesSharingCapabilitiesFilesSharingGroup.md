# FilesSharingCapabilitiesFilesSharingGroup


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**enabled** | **bool** |  | 
**expire_date** | [**FilesSharingCapabilitiesFilesSharingUserExpireDate**](FilesSharingCapabilitiesFilesSharingUserExpireDate.md) |  | [optional] 

## Example

```python
from nextcloud_client.models.files_sharing_capabilities_files_sharing_group import FilesSharingCapabilitiesFilesSharingGroup

# TODO update the JSON string below
json = "{}"
# create an instance of FilesSharingCapabilitiesFilesSharingGroup from a JSON string
files_sharing_capabilities_files_sharing_group_instance = FilesSharingCapabilitiesFilesSharingGroup.from_json(json)
# print the JSON string representation of the object
print FilesSharingCapabilitiesFilesSharingGroup.to_json()

# convert the object into a dict
files_sharing_capabilities_files_sharing_group_dict = files_sharing_capabilities_files_sharing_group_instance.to_dict()
# create an instance of FilesSharingCapabilitiesFilesSharingGroup from a dict
files_sharing_capabilities_files_sharing_group_form_dict = files_sharing_capabilities_files_sharing_group.from_dict(files_sharing_capabilities_files_sharing_group_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


