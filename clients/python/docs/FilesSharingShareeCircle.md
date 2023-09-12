# FilesSharingShareeCircle


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**count** | **int** |  | 
**label** | **str** |  | 
**share_with_description** | **str** |  | 
**value** | [**FilesSharingShareeCircleAllOfValue**](FilesSharingShareeCircleAllOfValue.md) |  | 

## Example

```python
from nextcloud_client.models.files_sharing_sharee_circle import FilesSharingShareeCircle

# TODO update the JSON string below
json = "{}"
# create an instance of FilesSharingShareeCircle from a JSON string
files_sharing_sharee_circle_instance = FilesSharingShareeCircle.from_json(json)
# print the JSON string representation of the object
print FilesSharingShareeCircle.to_json()

# convert the object into a dict
files_sharing_sharee_circle_dict = files_sharing_sharee_circle_instance.to_dict()
# create an instance of FilesSharingShareeCircle from a dict
files_sharing_sharee_circle_form_dict = files_sharing_sharee_circle.from_dict(files_sharing_sharee_circle_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


