# FilesExternalApiGetUserMounts200ResponseOcs


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**meta** | [**OCSMeta**](OCSMeta.md) |  | 
**data** | [**List[FilesExternalMount]**](FilesExternalMount.md) |  | 

## Example

```python
from nextcloud_client.models.files_external_api_get_user_mounts200_response_ocs import FilesExternalApiGetUserMounts200ResponseOcs

# TODO update the JSON string below
json = "{}"
# create an instance of FilesExternalApiGetUserMounts200ResponseOcs from a JSON string
files_external_api_get_user_mounts200_response_ocs_instance = FilesExternalApiGetUserMounts200ResponseOcs.from_json(json)
# print the JSON string representation of the object
print FilesExternalApiGetUserMounts200ResponseOcs.to_json()

# convert the object into a dict
files_external_api_get_user_mounts200_response_ocs_dict = files_external_api_get_user_mounts200_response_ocs_instance.to_dict()
# create an instance of FilesExternalApiGetUserMounts200ResponseOcs from a dict
files_external_api_get_user_mounts200_response_ocs_form_dict = files_external_api_get_user_mounts200_response_ocs.from_dict(files_external_api_get_user_mounts200_response_ocs_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


