# FilesSharingShareeRemote


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**count** | **int** |  | 
**label** | **str** |  | 
**uuid** | **str** |  | 
**name** | **str** |  | 
**type** | **str** |  | 
**value** | [**FilesSharingShareeRemoteAllOfValue**](FilesSharingShareeRemoteAllOfValue.md) |  | 

## Example

```python
from nextcloud_client.models.files_sharing_sharee_remote import FilesSharingShareeRemote

# TODO update the JSON string below
json = "{}"
# create an instance of FilesSharingShareeRemote from a JSON string
files_sharing_sharee_remote_instance = FilesSharingShareeRemote.from_json(json)
# print the JSON string representation of the object
print FilesSharingShareeRemote.to_json()

# convert the object into a dict
files_sharing_sharee_remote_dict = files_sharing_sharee_remote_instance.to_dict()
# create an instance of FilesSharingShareeRemote from a dict
files_sharing_sharee_remote_form_dict = files_sharing_sharee_remote.from_dict(files_sharing_sharee_remote_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


