# CloudFederationApiCapabilities


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ocm** | [**CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcm**](CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcm.md) |  | 

## Example

```python
from openapi_client.models.cloud_federation_api_capabilities import CloudFederationApiCapabilities

# TODO update the JSON string below
json = "{}"
# create an instance of CloudFederationApiCapabilities from a JSON string
cloud_federation_api_capabilities_instance = CloudFederationApiCapabilities.from_json(json)
# print the JSON string representation of the object
print CloudFederationApiCapabilities.to_json()

# convert the object into a dict
cloud_federation_api_capabilities_dict = cloud_federation_api_capabilities_instance.to_dict()
# create an instance of CloudFederationApiCapabilities from a dict
cloud_federation_api_capabilities_form_dict = cloud_federation_api_capabilities.from_dict(cloud_federation_api_capabilities_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


