# FilesSharingShareStatus


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**clear_at** | **int** |  | [optional] 
**icon** | **str** |  | [optional] 
**message** | **str** |  | [optional] 
**status** | **str** |  | [optional] 

## Example

```python
from nextcloud_client.models.files_sharing_share_status import FilesSharingShareStatus

# TODO update the JSON string below
json = "{}"
# create an instance of FilesSharingShareStatus from a JSON string
files_sharing_share_status_instance = FilesSharingShareStatus.from_json(json)
# print the JSON string representation of the object
print FilesSharingShareStatus.to_json()

# convert the object into a dict
files_sharing_share_status_dict = files_sharing_share_status_instance.to_dict()
# create an instance of FilesSharingShareStatus from a dict
files_sharing_share_status_form_dict = files_sharing_share_status.from_dict(files_sharing_share_status_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


