# FilesSharingDeletedShare


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **str** |  | 
**share_type** | **int** |  | 
**uid_owner** | **str** |  | 
**displayname_owner** | **str** |  | 
**permissions** | **int** |  | 
**stime** | **int** |  | 
**uid_file_owner** | **str** |  | 
**displayname_file_owner** | **str** |  | 
**path** | **str** |  | 
**item_type** | **str** |  | 
**mimetype** | **str** |  | 
**storage** | **int** |  | 
**item_source** | **int** |  | 
**file_source** | **int** |  | 
**file_parent** | **int** |  | 
**file_target** | **int** |  | 
**expiration** | **str** |  | 
**share_with** | **str** |  | 
**share_with_displayname** | **str** |  | 
**share_with_link** | **str** |  | 

## Example

```python
from nextcloud_client.models.files_sharing_deleted_share import FilesSharingDeletedShare

# TODO update the JSON string below
json = "{}"
# create an instance of FilesSharingDeletedShare from a JSON string
files_sharing_deleted_share_instance = FilesSharingDeletedShare.from_json(json)
# print the JSON string representation of the object
print FilesSharingDeletedShare.to_json()

# convert the object into a dict
files_sharing_deleted_share_dict = files_sharing_deleted_share_instance.to_dict()
# create an instance of FilesSharingDeletedShare from a dict
files_sharing_deleted_share_form_dict = files_sharing_deleted_share.from_dict(files_sharing_deleted_share_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


