# FilesCapabilities


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**files** | [**FilesCapabilitiesFiles**](FilesCapabilitiesFiles.md) |  | 

## Example

```python
from nextcloud_client.models.files_capabilities import FilesCapabilities

# TODO update the JSON string below
json = "{}"
# create an instance of FilesCapabilities from a JSON string
files_capabilities_instance = FilesCapabilities.from_json(json)
# print the JSON string representation of the object
print FilesCapabilities.to_json()

# convert the object into a dict
files_capabilities_dict = files_capabilities_instance.to_dict()
# create an instance of FilesCapabilities from a dict
files_capabilities_form_dict = files_capabilities.from_dict(files_capabilities_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


