# FilesRemindersApiGet200ResponseOcs


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**meta** | [**OCSMeta**](OCSMeta.md) |  | 
**data** | [**FilesRemindersApiGet200ResponseOcsData**](FilesRemindersApiGet200ResponseOcsData.md) |  | 

## Example

```python
from nextcloud_client.models.files_reminders_api_get200_response_ocs import FilesRemindersApiGet200ResponseOcs

# TODO update the JSON string below
json = "{}"
# create an instance of FilesRemindersApiGet200ResponseOcs from a JSON string
files_reminders_api_get200_response_ocs_instance = FilesRemindersApiGet200ResponseOcs.from_json(json)
# print the JSON string representation of the object
print FilesRemindersApiGet200ResponseOcs.to_json()

# convert the object into a dict
files_reminders_api_get200_response_ocs_dict = files_reminders_api_get200_response_ocs_instance.to_dict()
# create an instance of FilesRemindersApiGet200ResponseOcs from a dict
files_reminders_api_get200_response_ocs_form_dict = files_reminders_api_get200_response_ocs.from_dict(files_reminders_api_get200_response_ocs_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

