# FilesSharingShareeEmailAllOf


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**uuid** | **str** |  | 
**name** | **str** |  | 
**type** | **str** |  | 
**share_with_display_name_unique** | **str** |  | 
**value** | [**FilesSharingShareeValue**](FilesSharingShareeValue.md) |  | 

## Example

```python
from openapi_client.models.files_sharing_sharee_email_all_of import FilesSharingShareeEmailAllOf

# TODO update the JSON string below
json = "{}"
# create an instance of FilesSharingShareeEmailAllOf from a JSON string
files_sharing_sharee_email_all_of_instance = FilesSharingShareeEmailAllOf.from_json(json)
# print the JSON string representation of the object
print FilesSharingShareeEmailAllOf.to_json()

# convert the object into a dict
files_sharing_sharee_email_all_of_dict = files_sharing_sharee_email_all_of_instance.to_dict()
# create an instance of FilesSharingShareeEmailAllOf from a dict
files_sharing_sharee_email_all_of_form_dict = files_sharing_sharee_email_all_of.from_dict(files_sharing_sharee_email_all_of_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


