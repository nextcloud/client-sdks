# FilesExternalStorageConfig


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**applicable_groups** | **List[str]** |  | [optional] 
**applicable_users** | **List[str]** |  | [optional] 
**auth_mechanism** | **str** |  | 
**backend** | **str** |  | 
**backend_options** | **Dict[str, object]** |  | 
**id** | **int** |  | [optional] 
**mount_options** | **Dict[str, object]** |  | [optional] 
**mount_point** | **str** |  | 
**priority** | **int** |  | [optional] 
**status** | **int** |  | [optional] 
**status_message** | **str** |  | [optional] 
**type** | **str** |  | 
**user_provided** | **bool** |  | 

## Example

```python
from nextcloud_client.models.files_external_storage_config import FilesExternalStorageConfig

# TODO update the JSON string below
json = "{}"
# create an instance of FilesExternalStorageConfig from a JSON string
files_external_storage_config_instance = FilesExternalStorageConfig.from_json(json)
# print the JSON string representation of the object
print FilesExternalStorageConfig.to_json()

# convert the object into a dict
files_external_storage_config_dict = files_external_storage_config_instance.to_dict()
# create an instance of FilesExternalStorageConfig from a dict
files_external_storage_config_form_dict = files_external_storage_config.from_dict(files_external_storage_config_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


