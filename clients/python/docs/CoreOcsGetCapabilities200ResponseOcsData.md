# CoreOcsGetCapabilities200ResponseOcsData


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**version** | [**CoreOcsGetCapabilities200ResponseOcsDataVersion**](CoreOcsGetCapabilities200ResponseOcsDataVersion.md) |  | 
**capabilities** | [**CoreOcsGetCapabilities200ResponseOcsDataCapabilities**](CoreOcsGetCapabilities200ResponseOcsDataCapabilities.md) |  | 

## Example

```python
from nextcloud_client.models.core_ocs_get_capabilities200_response_ocs_data import CoreOcsGetCapabilities200ResponseOcsData

# TODO update the JSON string below
json = "{}"
# create an instance of CoreOcsGetCapabilities200ResponseOcsData from a JSON string
core_ocs_get_capabilities200_response_ocs_data_instance = CoreOcsGetCapabilities200ResponseOcsData.from_json(json)
# print the JSON string representation of the object
print CoreOcsGetCapabilities200ResponseOcsData.to_json()

# convert the object into a dict
core_ocs_get_capabilities200_response_ocs_data_dict = core_ocs_get_capabilities200_response_ocs_data_instance.to_dict()
# create an instance of CoreOcsGetCapabilities200ResponseOcsData from a dict
core_ocs_get_capabilities200_response_ocs_data_form_dict = core_ocs_get_capabilities200_response_ocs_data.from_dict(core_ocs_get_capabilities200_response_ocs_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


