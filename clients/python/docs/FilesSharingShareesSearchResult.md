# FilesSharingShareesSearchResult


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**exact** | [**FilesSharingShareesSearchResultExact**](FilesSharingShareesSearchResultExact.md) |  | 
**circles** | [**List[FilesSharingShareeCircle]**](FilesSharingShareeCircle.md) |  | 
**emails** | [**List[FilesSharingShareeEmail]**](FilesSharingShareeEmail.md) |  | 
**groups** | [**List[FilesSharingSharee]**](FilesSharingSharee.md) |  | 
**lookup** | [**List[FilesSharingShareeLookup]**](FilesSharingShareeLookup.md) |  | 
**remote_groups** | [**List[FilesSharingShareeRemoteGroup]**](FilesSharingShareeRemoteGroup.md) |  | 
**remotes** | [**List[FilesSharingShareeRemote]**](FilesSharingShareeRemote.md) |  | 
**rooms** | [**List[FilesSharingSharee]**](FilesSharingSharee.md) |  | 
**users** | [**List[FilesSharingShareeUser]**](FilesSharingShareeUser.md) |  | 
**lookup_enabled** | **bool** |  | 

## Example

```python
from nextcloud_client.models.files_sharing_sharees_search_result import FilesSharingShareesSearchResult

# TODO update the JSON string below
json = "{}"
# create an instance of FilesSharingShareesSearchResult from a JSON string
files_sharing_sharees_search_result_instance = FilesSharingShareesSearchResult.from_json(json)
# print the JSON string representation of the object
print FilesSharingShareesSearchResult.to_json()

# convert the object into a dict
files_sharing_sharees_search_result_dict = files_sharing_sharees_search_result_instance.to_dict()
# create an instance of FilesSharingShareesSearchResult from a dict
files_sharing_sharees_search_result_form_dict = files_sharing_sharees_search_result.from_dict(files_sharing_sharees_search_result_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


