# FilesSharingShareesRecommendedResult


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**exact** | [**FilesSharingShareesRecommendedResultExact**](FilesSharingShareesRecommendedResultExact.md) |  | 
**emails** | [**List[FilesSharingShareeEmail]**](FilesSharingShareeEmail.md) |  | 
**groups** | [**List[FilesSharingSharee]**](FilesSharingSharee.md) |  | 
**remote_groups** | [**List[FilesSharingShareeRemoteGroup]**](FilesSharingShareeRemoteGroup.md) |  | 
**remotes** | [**List[FilesSharingShareeRemote]**](FilesSharingShareeRemote.md) |  | 
**users** | [**List[FilesSharingShareeUser]**](FilesSharingShareeUser.md) |  | 

## Example

```python
from nextcloud_client.models.files_sharing_sharees_recommended_result import FilesSharingShareesRecommendedResult

# TODO update the JSON string below
json = "{}"
# create an instance of FilesSharingShareesRecommendedResult from a JSON string
files_sharing_sharees_recommended_result_instance = FilesSharingShareesRecommendedResult.from_json(json)
# print the JSON string representation of the object
print FilesSharingShareesRecommendedResult.to_json()

# convert the object into a dict
files_sharing_sharees_recommended_result_dict = files_sharing_sharees_recommended_result_instance.to_dict()
# create an instance of FilesSharingShareesRecommendedResult from a dict
files_sharing_sharees_recommended_result_form_dict = files_sharing_sharees_recommended_result.from_dict(files_sharing_sharees_recommended_result_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


