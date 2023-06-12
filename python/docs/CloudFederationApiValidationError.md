# CloudFederationApiValidationError


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**message** | **str** |  | 
**validation_errors** | [**List[CloudFederationApiValidationErrorAllOfValidationErrors]**](CloudFederationApiValidationErrorAllOfValidationErrors.md) |  | 

## Example

```python
from openapi_client.models.cloud_federation_api_validation_error import CloudFederationApiValidationError

# TODO update the JSON string below
json = "{}"
# create an instance of CloudFederationApiValidationError from a JSON string
cloud_federation_api_validation_error_instance = CloudFederationApiValidationError.from_json(json)
# print the JSON string representation of the object
print CloudFederationApiValidationError.to_json()

# convert the object into a dict
cloud_federation_api_validation_error_dict = cloud_federation_api_validation_error_instance.to_dict()
# create an instance of CloudFederationApiValidationError from a dict
cloud_federation_api_validation_error_form_dict = cloud_federation_api_validation_error.from_dict(cloud_federation_api_validation_error_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


