# FilesSharingShare


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**attributes** | **str** |  | 
**can_delete** | **bool** |  | 
**can_edit** | **bool** |  | 
**displayname_file_owner** | **str** |  | 
**displayname_owner** | **str** |  | 
**expiration** | **str** |  | 
**file_parent** | **int** |  | 
**file_source** | **int** |  | 
**file_target** | **str** |  | 
**has_preview** | **bool** |  | 
**hide_download** | **int** |  | 
**id** | **str** |  | 
**item_mtime** | **int** |  | 
**item_permissions** | **int** |  | [optional] 
**item_size** | [**FilesSharingShareItemSize**](FilesSharingShareItemSize.md) |  | 
**item_source** | **int** |  | 
**item_type** | **str** |  | 
**label** | **str** |  | 
**mail_send** | **int** |  | 
**mimetype** | **str** |  | 
**note** | **str** |  | 
**parent** | **object** |  | 
**password** | **str** |  | [optional] 
**password_expiration_time** | **str** |  | [optional] 
**path** | **str** |  | 
**permissions** | **int** |  | 
**send_password_by_talk** | **bool** |  | [optional] 
**share_type** | **int** |  | 
**share_with** | **str** |  | [optional] 
**share_with_avatar** | **str** |  | [optional] 
**share_with_displayname** | **str** |  | [optional] 
**share_with_displayname_unique** | **str** |  | [optional] 
**share_with_link** | **str** |  | [optional] 
**status** | [**FilesSharingShareStatus**](FilesSharingShareStatus.md) |  | [optional] 
**stime** | **int** |  | 
**storage** | **int** |  | 
**storage_id** | **str** |  | 
**token** | **str** |  | 
**uid_file_owner** | **str** |  | 
**uid_owner** | **str** |  | 
**url** | **str** |  | [optional] 

## Example

```python
from nextcloud_client.models.files_sharing_share import FilesSharingShare

# TODO update the JSON string below
json = "{}"
# create an instance of FilesSharingShare from a JSON string
files_sharing_share_instance = FilesSharingShare.from_json(json)
# print the JSON string representation of the object
print FilesSharingShare.to_json()

# convert the object into a dict
files_sharing_share_dict = files_sharing_share_instance.to_dict()
# create an instance of FilesSharingShare from a dict
files_sharing_share_form_dict = files_sharing_share.from_dict(files_sharing_share_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


