# FilesCapabilitiesFilesDirectEditing


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**url** | **str** |  | 
**etag** | **str** |  | 
**supports_file_id** | **bool** |  | 

## Example

```python
from nextcloud_client.models.files_capabilities_files_direct_editing import FilesCapabilitiesFilesDirectEditing

# TODO update the JSON string below
json = "{}"
# create an instance of FilesCapabilitiesFilesDirectEditing from a JSON string
files_capabilities_files_direct_editing_instance = FilesCapabilitiesFilesDirectEditing.from_json(json)
# print the JSON string representation of the object
print FilesCapabilitiesFilesDirectEditing.to_json()

# convert the object into a dict
files_capabilities_files_direct_editing_dict = files_capabilities_files_direct_editing_instance.to_dict()
# create an instance of FilesCapabilitiesFilesDirectEditing from a dict
files_capabilities_files_direct_editing_form_dict = files_capabilities_files_direct_editing.from_dict(files_capabilities_files_direct_editing_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


