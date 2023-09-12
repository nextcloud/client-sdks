# DavCapabilities


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**dav** | [**DavCapabilitiesDav**](DavCapabilitiesDav.md) |  | 

## Example

```python
from nextcloud_client.models.dav_capabilities import DavCapabilities

# TODO update the JSON string below
json = "{}"
# create an instance of DavCapabilities from a JSON string
dav_capabilities_instance = DavCapabilities.from_json(json)
# print the JSON string representation of the object
print DavCapabilities.to_json()

# convert the object into a dict
dav_capabilities_dict = dav_capabilities_instance.to_dict()
# create an instance of DavCapabilities from a dict
dav_capabilities_form_dict = dav_capabilities.from_dict(dav_capabilities_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


