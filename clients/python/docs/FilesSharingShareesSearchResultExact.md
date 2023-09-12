# FilesSharingShareesSearchResultExact


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**circles** | [**List[FilesSharingShareeCircle]**](FilesSharingShareeCircle.md) |  | 
**emails** | [**List[FilesSharingShareeEmail]**](FilesSharingShareeEmail.md) |  | 
**groups** | [**List[FilesSharingSharee]**](FilesSharingSharee.md) |  | 
**remote_groups** | [**List[FilesSharingShareeRemoteGroup]**](FilesSharingShareeRemoteGroup.md) |  | 
**remotes** | [**List[FilesSharingShareeRemote]**](FilesSharingShareeRemote.md) |  | 
**rooms** | [**List[FilesSharingSharee]**](FilesSharingSharee.md) |  | 
**users** | [**List[FilesSharingShareeUser]**](FilesSharingShareeUser.md) |  | 

## Example

```python
from nextcloud_client.models.files_sharing_sharees_search_result_exact import FilesSharingShareesSearchResultExact

# TODO update the JSON string below
json = "{}"
# create an instance of FilesSharingShareesSearchResultExact from a JSON string
files_sharing_sharees_search_result_exact_instance = FilesSharingShareesSearchResultExact.from_json(json)
# print the JSON string representation of the object
print FilesSharingShareesSearchResultExact.to_json()

# convert the object into a dict
files_sharing_sharees_search_result_exact_dict = files_sharing_sharees_search_result_exact_instance.to_dict()
# create an instance of FilesSharingShareesSearchResultExact from a dict
files_sharing_sharees_search_result_exact_form_dict = files_sharing_sharees_search_result_exact.from_dict(files_sharing_sharees_search_result_exact_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


