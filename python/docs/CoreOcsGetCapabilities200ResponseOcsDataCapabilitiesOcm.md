# CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcm


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**enabled** | **bool** |  | 
**api_version** | **str** |  | 
**end_point** | **str** |  | 
**resource_types** | [**List[CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcmResourceTypesInner]**](CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcmResourceTypesInner.md) |  | 

## Example

```python
from openapi_client.models.core_ocs_get_capabilities200_response_ocs_data_capabilities_ocm import CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcm

# TODO update the JSON string below
json = "{}"
# create an instance of CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcm from a JSON string
core_ocs_get_capabilities200_response_ocs_data_capabilities_ocm_instance = CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcm.from_json(json)
# print the JSON string representation of the object
print CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcm.to_json()

# convert the object into a dict
core_ocs_get_capabilities200_response_ocs_data_capabilities_ocm_dict = core_ocs_get_capabilities200_response_ocs_data_capabilities_ocm_instance.to_dict()
# create an instance of CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcm from a dict
core_ocs_get_capabilities200_response_ocs_data_capabilities_ocm_form_dict = core_ocs_get_capabilities200_response_ocs_data_capabilities_ocm.from_dict(core_ocs_get_capabilities200_response_ocs_data_capabilities_ocm_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


