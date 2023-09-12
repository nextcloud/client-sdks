# UpdatenotificationApiGetAppList200ResponseOcsData


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**missing** | [**List[UpdatenotificationApp]**](UpdatenotificationApp.md) |  | 
**available** | [**List[UpdatenotificationApp]**](UpdatenotificationApp.md) |  | 

## Example

```python
from nextcloud_client.models.updatenotification_api_get_app_list200_response_ocs_data import UpdatenotificationApiGetAppList200ResponseOcsData

# TODO update the JSON string below
json = "{}"
# create an instance of UpdatenotificationApiGetAppList200ResponseOcsData from a JSON string
updatenotification_api_get_app_list200_response_ocs_data_instance = UpdatenotificationApiGetAppList200ResponseOcsData.from_json(json)
# print the JSON string representation of the object
print UpdatenotificationApiGetAppList200ResponseOcsData.to_json()

# convert the object into a dict
updatenotification_api_get_app_list200_response_ocs_data_dict = updatenotification_api_get_app_list200_response_ocs_data_instance.to_dict()
# create an instance of UpdatenotificationApiGetAppList200ResponseOcsData from a dict
updatenotification_api_get_app_list200_response_ocs_data_form_dict = updatenotification_api_get_app_list200_response_ocs_data.from_dict(updatenotification_api_get_app_list200_response_ocs_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


