# CoreReferenceApiGetProvidersInfo200ResponseOcs


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**meta** | [**OCSMeta**](OCSMeta.md) |  | 
**data** | [**List[CoreReferenceProvider]**](CoreReferenceProvider.md) |  | 

## Example

```python
from nextcloud_client.models.core_reference_api_get_providers_info200_response_ocs import CoreReferenceApiGetProvidersInfo200ResponseOcs

# TODO update the JSON string below
json = "{}"
# create an instance of CoreReferenceApiGetProvidersInfo200ResponseOcs from a JSON string
core_reference_api_get_providers_info200_response_ocs_instance = CoreReferenceApiGetProvidersInfo200ResponseOcs.from_json(json)
# print the JSON string representation of the object
print CoreReferenceApiGetProvidersInfo200ResponseOcs.to_json()

# convert the object into a dict
core_reference_api_get_providers_info200_response_ocs_dict = core_reference_api_get_providers_info200_response_ocs_instance.to_dict()
# create an instance of CoreReferenceApiGetProvidersInfo200ResponseOcs from a dict
core_reference_api_get_providers_info200_response_ocs_form_dict = core_reference_api_get_providers_info200_response_ocs.from_dict(core_reference_api_get_providers_info200_response_ocs_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


