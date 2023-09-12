# FilesVersionsCapabilities


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**files** | [**FilesVersionsCapabilitiesFiles**](FilesVersionsCapabilitiesFiles.md) |  | 

## Example

```python
from nextcloud_client.models.files_versions_capabilities import FilesVersionsCapabilities

# TODO update the JSON string below
json = "{}"
# create an instance of FilesVersionsCapabilities from a JSON string
files_versions_capabilities_instance = FilesVersionsCapabilities.from_json(json)
# print the JSON string representation of the object
print FilesVersionsCapabilities.to_json()

# convert the object into a dict
files_versions_capabilities_dict = files_versions_capabilities_instance.to_dict()
# create an instance of FilesVersionsCapabilities from a dict
files_versions_capabilities_form_dict = files_versions_capabilities.from_dict(files_versions_capabilities_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


