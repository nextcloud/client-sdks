# FilesExternalMount


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** |  | 
**path** | **str** |  | 
**type** | **str** |  | 
**backend** | **str** |  | 
**scope** | **str** |  | 
**permissions** | **int** |  | 
**id** | **int** |  | 
**var_class** | **str** |  | 
**config** | [**FilesExternalStorageConfig**](FilesExternalStorageConfig.md) |  | 

## Example

```python
from nextcloud_client.models.files_external_mount import FilesExternalMount

# TODO update the JSON string below
json = "{}"
# create an instance of FilesExternalMount from a JSON string
files_external_mount_instance = FilesExternalMount.from_json(json)
# print the JSON string representation of the object
print FilesExternalMount.to_json()

# convert the object into a dict
files_external_mount_dict = files_external_mount_instance.to_dict()
# create an instance of FilesExternalMount from a dict
files_external_mount_form_dict = files_external_mount.from_dict(files_external_mount_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


