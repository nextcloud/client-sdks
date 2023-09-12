# FilesSharingDeletedShareapiList200ResponseOcs


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**meta** | [**OCSMeta**](OCSMeta.md) |  | 
**data** | [**List[FilesSharingDeletedShare]**](FilesSharingDeletedShare.md) |  | 

## Example

```python
from nextcloud_client.models.files_sharing_deleted_shareapi_list200_response_ocs import FilesSharingDeletedShareapiList200ResponseOcs

# TODO update the JSON string below
json = "{}"
# create an instance of FilesSharingDeletedShareapiList200ResponseOcs from a JSON string
files_sharing_deleted_shareapi_list200_response_ocs_instance = FilesSharingDeletedShareapiList200ResponseOcs.from_json(json)
# print the JSON string representation of the object
print FilesSharingDeletedShareapiList200ResponseOcs.to_json()

# convert the object into a dict
files_sharing_deleted_shareapi_list200_response_ocs_dict = files_sharing_deleted_shareapi_list200_response_ocs_instance.to_dict()
# create an instance of FilesSharingDeletedShareapiList200ResponseOcs from a dict
files_sharing_deleted_shareapi_list200_response_ocs_form_dict = files_sharing_deleted_shareapi_list200_response_ocs.from_dict(files_sharing_deleted_shareapi_list200_response_ocs_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


