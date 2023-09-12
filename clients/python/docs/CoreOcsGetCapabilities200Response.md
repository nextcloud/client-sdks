# CoreOcsGetCapabilities200Response


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ocs** | [**CoreOcsGetCapabilities200ResponseOcs**](CoreOcsGetCapabilities200ResponseOcs.md) |  | 

## Example

```python
from nextcloud_client.models.core_ocs_get_capabilities200_response import CoreOcsGetCapabilities200Response

# TODO update the JSON string below
json = "{}"
# create an instance of CoreOcsGetCapabilities200Response from a JSON string
core_ocs_get_capabilities200_response_instance = CoreOcsGetCapabilities200Response.from_json(json)
# print the JSON string representation of the object
print CoreOcsGetCapabilities200Response.to_json()

# convert the object into a dict
core_ocs_get_capabilities200_response_dict = core_ocs_get_capabilities200_response_instance.to_dict()
# create an instance of CoreOcsGetCapabilities200Response from a dict
core_ocs_get_capabilities200_response_form_dict = core_ocs_get_capabilities200_response.from_dict(core_ocs_get_capabilities200_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


