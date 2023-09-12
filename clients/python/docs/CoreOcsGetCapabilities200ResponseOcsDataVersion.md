# CoreOcsGetCapabilities200ResponseOcsDataVersion


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**major** | **int** |  | 
**minor** | **int** |  | 
**micro** | **int** |  | 
**string** | **str** |  | 
**edition** | **str** |  | 
**extended_support** | **bool** |  | 

## Example

```python
from nextcloud_client.models.core_ocs_get_capabilities200_response_ocs_data_version import CoreOcsGetCapabilities200ResponseOcsDataVersion

# TODO update the JSON string below
json = "{}"
# create an instance of CoreOcsGetCapabilities200ResponseOcsDataVersion from a JSON string
core_ocs_get_capabilities200_response_ocs_data_version_instance = CoreOcsGetCapabilities200ResponseOcsDataVersion.from_json(json)
# print the JSON string representation of the object
print CoreOcsGetCapabilities200ResponseOcsDataVersion.to_json()

# convert the object into a dict
core_ocs_get_capabilities200_response_ocs_data_version_dict = core_ocs_get_capabilities200_response_ocs_data_version_instance.to_dict()
# create an instance of CoreOcsGetCapabilities200ResponseOcsDataVersion from a dict
core_ocs_get_capabilities200_response_ocs_data_version_form_dict = core_ocs_get_capabilities200_response_ocs_data_version.from_dict(core_ocs_get_capabilities200_response_ocs_data_version_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


