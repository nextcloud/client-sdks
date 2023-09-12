# SharebymailCapabilities


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**files_sharing** | [**SharebymailCapabilitiesFilesSharing**](SharebymailCapabilitiesFilesSharing.md) |  | 

## Example

```python
from nextcloud_client.models.sharebymail_capabilities import SharebymailCapabilities

# TODO update the JSON string below
json = "{}"
# create an instance of SharebymailCapabilities from a JSON string
sharebymail_capabilities_instance = SharebymailCapabilities.from_json(json)
# print the JSON string representation of the object
print SharebymailCapabilities.to_json()

# convert the object into a dict
sharebymail_capabilities_dict = sharebymail_capabilities_instance.to_dict()
# create an instance of SharebymailCapabilities from a dict
sharebymail_capabilities_form_dict = sharebymail_capabilities.from_dict(sharebymail_capabilities_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


