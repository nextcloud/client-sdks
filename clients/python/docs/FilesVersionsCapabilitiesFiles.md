# FilesVersionsCapabilitiesFiles


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**versioning** | **bool** |  | 
**version_labeling** | **bool** |  | 
**version_deletion** | **bool** |  | 

## Example

```python
from nextcloud_client.models.files_versions_capabilities_files import FilesVersionsCapabilitiesFiles

# TODO update the JSON string below
json = "{}"
# create an instance of FilesVersionsCapabilitiesFiles from a JSON string
files_versions_capabilities_files_instance = FilesVersionsCapabilitiesFiles.from_json(json)
# print the JSON string representation of the object
print FilesVersionsCapabilitiesFiles.to_json()

# convert the object into a dict
files_versions_capabilities_files_dict = files_versions_capabilities_files_instance.to_dict()
# create an instance of FilesVersionsCapabilitiesFiles from a dict
files_versions_capabilities_files_form_dict = files_versions_capabilities_files.from_dict(files_versions_capabilities_files_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


