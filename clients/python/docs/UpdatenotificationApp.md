# UpdatenotificationApp


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**app_id** | **str** |  | 
**app_name** | **str** |  | 

## Example

```python
from nextcloud_client.models.updatenotification_app import UpdatenotificationApp

# TODO update the JSON string below
json = "{}"
# create an instance of UpdatenotificationApp from a JSON string
updatenotification_app_instance = UpdatenotificationApp.from_json(json)
# print the JSON string representation of the object
print UpdatenotificationApp.to_json()

# convert the object into a dict
updatenotification_app_dict = updatenotification_app_instance.to_dict()
# create an instance of UpdatenotificationApp from a dict
updatenotification_app_form_dict = updatenotification_app.from_dict(updatenotification_app_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


