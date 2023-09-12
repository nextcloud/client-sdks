# FilesSharingSharee


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**count** | **int** |  | 
**label** | **str** |  | 

## Example

```python
from nextcloud_client.models.files_sharing_sharee import FilesSharingSharee

# TODO update the JSON string below
json = "{}"
# create an instance of FilesSharingSharee from a JSON string
files_sharing_sharee_instance = FilesSharingSharee.from_json(json)
# print the JSON string representation of the object
print FilesSharingSharee.to_json()

# convert the object into a dict
files_sharing_sharee_dict = files_sharing_sharee_instance.to_dict()
# create an instance of FilesSharingSharee from a dict
files_sharing_sharee_form_dict = files_sharing_sharee.from_dict(files_sharing_sharee_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


