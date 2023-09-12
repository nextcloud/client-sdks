# FilesCapabilitiesFiles


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**bigfilechunking** | **bool** |  | 
**blacklisted_files** | **List[object]** |  | 
**direct_editing** | [**FilesCapabilitiesFilesDirectEditing**](FilesCapabilitiesFilesDirectEditing.md) |  | 

## Example

```python
from nextcloud_client.models.files_capabilities_files import FilesCapabilitiesFiles

# TODO update the JSON string below
json = "{}"
# create an instance of FilesCapabilitiesFiles from a JSON string
files_capabilities_files_instance = FilesCapabilitiesFiles.from_json(json)
# print the JSON string representation of the object
print FilesCapabilitiesFiles.to_json()

# convert the object into a dict
files_capabilities_files_dict = files_capabilities_files_instance.to_dict()
# create an instance of FilesCapabilitiesFiles from a dict
files_capabilities_files_form_dict = files_capabilities_files.from_dict(files_capabilities_files_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


