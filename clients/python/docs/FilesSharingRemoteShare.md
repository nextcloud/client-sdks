# FilesSharingRemoteShare


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**accepted** | **bool** |  | 
**file_id** | **int** |  | 
**id** | **int** |  | 
**mimetype** | **str** |  | 
**mountpoint** | **str** |  | 
**mtime** | **int** |  | 
**name** | **str** |  | 
**owner** | **str** |  | 
**parent** | **int** |  | 
**permissions** | **int** |  | 
**remote** | **str** |  | 
**remote_id** | **str** |  | 
**share_token** | **str** |  | 
**share_type** | **int** |  | 
**type** | **str** |  | 
**user** | **str** |  | 

## Example

```python
from nextcloud_client.models.files_sharing_remote_share import FilesSharingRemoteShare

# TODO update the JSON string below
json = "{}"
# create an instance of FilesSharingRemoteShare from a JSON string
files_sharing_remote_share_instance = FilesSharingRemoteShare.from_json(json)
# print the JSON string representation of the object
print FilesSharingRemoteShare.to_json()

# convert the object into a dict
files_sharing_remote_share_dict = files_sharing_remote_share_instance.to_dict()
# create an instance of FilesSharingRemoteShare from a dict
files_sharing_remote_share_form_dict = files_sharing_remote_share.from_dict(files_sharing_remote_share_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


