# FilesSharingShareeRemoteGroup


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**count** | **int** |  | 
**label** | **str** |  | 
**guid** | **str** |  | 
**name** | **str** |  | 
**value** | [**FilesSharingShareeRemoteAllOfValue**](FilesSharingShareeRemoteAllOfValue.md) |  | 

## Example

```python
from nextcloud_client.models.files_sharing_sharee_remote_group import FilesSharingShareeRemoteGroup

# TODO update the JSON string below
json = "{}"
# create an instance of FilesSharingShareeRemoteGroup from a JSON string
files_sharing_sharee_remote_group_instance = FilesSharingShareeRemoteGroup.from_json(json)
# print the JSON string representation of the object
print FilesSharingShareeRemoteGroup.to_json()

# convert the object into a dict
files_sharing_sharee_remote_group_dict = files_sharing_sharee_remote_group_instance.to_dict()
# create an instance of FilesSharingShareeRemoteGroup from a dict
files_sharing_sharee_remote_group_form_dict = files_sharing_sharee_remote_group.from_dict(files_sharing_sharee_remote_group_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


