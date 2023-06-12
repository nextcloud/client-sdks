# FilesSharingShareeUser


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**count** | **int** |  | 
**label** | **str** |  | 
**subline** | **str** |  | 
**icon** | **str** |  | 
**share_with_display_name_unique** | **str** |  | 
**status** | [**FilesSharingShareeUserAllOfStatus**](FilesSharingShareeUserAllOfStatus.md) |  | 
**value** | [**FilesSharingShareeValue**](FilesSharingShareeValue.md) |  | 

## Example

```python
from openapi_client.models.files_sharing_sharee_user import FilesSharingShareeUser

# TODO update the JSON string below
json = "{}"
# create an instance of FilesSharingShareeUser from a JSON string
files_sharing_sharee_user_instance = FilesSharingShareeUser.from_json(json)
# print the JSON string representation of the object
print FilesSharingShareeUser.to_json()

# convert the object into a dict
files_sharing_sharee_user_dict = files_sharing_sharee_user_instance.to_dict()
# create an instance of FilesSharingShareeUser from a dict
files_sharing_sharee_user_form_dict = files_sharing_sharee_user.from_dict(files_sharing_sharee_user_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


