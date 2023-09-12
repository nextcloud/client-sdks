# FilesSharingShareeLookup


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**count** | **int** |  | 
**label** | **str** |  | 
**extra** | [**FilesSharingShareeLookupAllOfExtra**](FilesSharingShareeLookupAllOfExtra.md) |  | 
**value** | [**FilesSharingShareeLookupAllOfValue**](FilesSharingShareeLookupAllOfValue.md) |  | 

## Example

```python
from nextcloud_client.models.files_sharing_sharee_lookup import FilesSharingShareeLookup

# TODO update the JSON string below
json = "{}"
# create an instance of FilesSharingShareeLookup from a JSON string
files_sharing_sharee_lookup_instance = FilesSharingShareeLookup.from_json(json)
# print the JSON string representation of the object
print FilesSharingShareeLookup.to_json()

# convert the object into a dict
files_sharing_sharee_lookup_dict = files_sharing_sharee_lookup_instance.to_dict()
# create an instance of FilesSharingShareeLookup from a dict
files_sharing_sharee_lookup_form_dict = files_sharing_sharee_lookup.from_dict(files_sharing_sharee_lookup_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


