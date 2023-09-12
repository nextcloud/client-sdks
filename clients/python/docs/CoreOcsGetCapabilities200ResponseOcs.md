# CoreOcsGetCapabilities200ResponseOcs


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**meta** | [**OCSMeta**](OCSMeta.md) |  | 
**data** | [**CoreOcsGetCapabilities200ResponseOcsData**](CoreOcsGetCapabilities200ResponseOcsData.md) |  | 

## Example

```python
from nextcloud_client.models.core_ocs_get_capabilities200_response_ocs import CoreOcsGetCapabilities200ResponseOcs

# TODO update the JSON string below
json = "{}"
# create an instance of CoreOcsGetCapabilities200ResponseOcs from a JSON string
core_ocs_get_capabilities200_response_ocs_instance = CoreOcsGetCapabilities200ResponseOcs.from_json(json)
# print the JSON string representation of the object
print CoreOcsGetCapabilities200ResponseOcs.to_json()

# convert the object into a dict
core_ocs_get_capabilities200_response_ocs_dict = core_ocs_get_capabilities200_response_ocs_instance.to_dict()
# create an instance of CoreOcsGetCapabilities200ResponseOcs from a dict
core_ocs_get_capabilities200_response_ocs_form_dict = core_ocs_get_capabilities200_response_ocs.from_dict(core_ocs_get_capabilities200_response_ocs_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


