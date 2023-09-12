# DavCapabilitiesDav


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**chunking** | **str** |  | 
**bulkupload** | **str** |  | [optional] 

## Example

```python
from nextcloud_client.models.dav_capabilities_dav import DavCapabilitiesDav

# TODO update the JSON string below
json = "{}"
# create an instance of DavCapabilitiesDav from a JSON string
dav_capabilities_dav_instance = DavCapabilitiesDav.from_json(json)
# print the JSON string representation of the object
print DavCapabilitiesDav.to_json()

# convert the object into a dict
dav_capabilities_dav_dict = dav_capabilities_dav_instance.to_dict()
# create an instance of DavCapabilitiesDav from a dict
dav_capabilities_dav_form_dict = dav_capabilities_dav.from_dict(dav_capabilities_dav_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


