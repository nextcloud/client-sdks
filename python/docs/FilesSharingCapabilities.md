# FilesSharingCapabilities


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**files_sharing** | [**CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing**](CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing.md) |  | 

## Example

```python
from openapi_client.models.files_sharing_capabilities import FilesSharingCapabilities

# TODO update the JSON string below
json = "{}"
# create an instance of FilesSharingCapabilities from a JSON string
files_sharing_capabilities_instance = FilesSharingCapabilities.from_json(json)
# print the JSON string representation of the object
print FilesSharingCapabilities.to_json()

# convert the object into a dict
files_sharing_capabilities_dict = files_sharing_capabilities_instance.to_dict()
# create an instance of FilesSharingCapabilities from a dict
files_sharing_capabilities_form_dict = files_sharing_capabilities.from_dict(files_sharing_capabilities_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


