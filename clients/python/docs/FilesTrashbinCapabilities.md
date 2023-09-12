# FilesTrashbinCapabilities


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**files** | [**FilesTrashbinCapabilitiesFiles**](FilesTrashbinCapabilitiesFiles.md) |  | 

## Example

```python
from nextcloud_client.models.files_trashbin_capabilities import FilesTrashbinCapabilities

# TODO update the JSON string below
json = "{}"
# create an instance of FilesTrashbinCapabilities from a JSON string
files_trashbin_capabilities_instance = FilesTrashbinCapabilities.from_json(json)
# print the JSON string representation of the object
print FilesTrashbinCapabilities.to_json()

# convert the object into a dict
files_trashbin_capabilities_dict = files_trashbin_capabilities_instance.to_dict()
# create an instance of FilesTrashbinCapabilities from a dict
files_trashbin_capabilities_form_dict = files_trashbin_capabilities.from_dict(files_trashbin_capabilities_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


