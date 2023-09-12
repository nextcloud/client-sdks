# FilesSharingShareInfo


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **int** |  | 
**parent_id** | **int** |  | 
**mtime** | **int** |  | 
**name** | **str** |  | 
**permissions** | **int** |  | 
**mimetype** | **str** |  | 
**size** | [**FilesSharingShareInfoSize**](FilesSharingShareInfoSize.md) |  | 
**type** | **str** |  | 
**etag** | **str** |  | 
**children** | **List[Dict[str, object]]** |  | [optional] 

## Example

```python
from nextcloud_client.models.files_sharing_share_info import FilesSharingShareInfo

# TODO update the JSON string below
json = "{}"
# create an instance of FilesSharingShareInfo from a JSON string
files_sharing_share_info_instance = FilesSharingShareInfo.from_json(json)
# print the JSON string representation of the object
print FilesSharingShareInfo.to_json()

# convert the object into a dict
files_sharing_share_info_dict = files_sharing_share_info_instance.to_dict()
# create an instance of FilesSharingShareInfo from a dict
files_sharing_share_info_form_dict = files_sharing_share_info.from_dict(files_sharing_share_info_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


