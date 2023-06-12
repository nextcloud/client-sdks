# FilesSharingShareeEmail


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**count** | **int** |  | 
**label** | **str** |  | 
**uuid** | **str** |  | 
**name** | **str** |  | 
**type** | **str** |  | 
**share_with_display_name_unique** | **str** |  | 
**value** | [**FilesSharingShareeValue**](FilesSharingShareeValue.md) |  | 

## Example

```python
from openapi_client.models.files_sharing_sharee_email import FilesSharingShareeEmail

# TODO update the JSON string below
json = "{}"
# create an instance of FilesSharingShareeEmail from a JSON string
files_sharing_sharee_email_instance = FilesSharingShareeEmail.from_json(json)
# print the JSON string representation of the object
print FilesSharingShareeEmail.to_json()

# convert the object into a dict
files_sharing_sharee_email_dict = files_sharing_sharee_email_instance.to_dict()
# create an instance of FilesSharingShareeEmail from a dict
files_sharing_sharee_email_form_dict = files_sharing_sharee_email.from_dict(files_sharing_sharee_email_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


