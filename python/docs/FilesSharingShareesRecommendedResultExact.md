# FilesSharingShareesRecommendedResultExact


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**emails** | [**List[FilesSharingShareeEmail]**](FilesSharingShareeEmail.md) |  | 
**groups** | [**List[FilesSharingSharee]**](FilesSharingSharee.md) |  | 
**remote_groups** | [**List[FilesSharingShareeRemoteGroup]**](FilesSharingShareeRemoteGroup.md) |  | 
**remotes** | [**List[FilesSharingShareeRemote]**](FilesSharingShareeRemote.md) |  | 
**users** | [**List[FilesSharingShareeUser]**](FilesSharingShareeUser.md) |  | 

## Example

```python
from openapi_client.models.files_sharing_sharees_recommended_result_exact import FilesSharingShareesRecommendedResultExact

# TODO update the JSON string below
json = "{}"
# create an instance of FilesSharingShareesRecommendedResultExact from a JSON string
files_sharing_sharees_recommended_result_exact_instance = FilesSharingShareesRecommendedResultExact.from_json(json)
# print the JSON string representation of the object
print FilesSharingShareesRecommendedResultExact.to_json()

# convert the object into a dict
files_sharing_sharees_recommended_result_exact_dict = files_sharing_sharees_recommended_result_exact_instance.to_dict()
# create an instance of FilesSharingShareesRecommendedResultExact from a dict
files_sharing_sharees_recommended_result_exact_form_dict = files_sharing_sharees_recommended_result_exact.from_dict(files_sharing_sharees_recommended_result_exact_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


