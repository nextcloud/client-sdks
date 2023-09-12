# ThemingPublicCapabilitiesTheming


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** |  | 
**url** | **str** |  | 
**slogan** | **str** |  | 
**color** | **str** |  | 
**color_text** | **str** |  | 
**color_element** | **str** |  | 
**color_element_bright** | **str** |  | 
**color_element_dark** | **str** |  | 
**logo** | **str** |  | 
**background** | **str** |  | 
**background_plain** | **bool** |  | 
**background_default** | **bool** |  | 
**logoheader** | **str** |  | 
**favicon** | **str** |  | 

## Example

```python
from nextcloud_client.models.theming_public_capabilities_theming import ThemingPublicCapabilitiesTheming

# TODO update the JSON string below
json = "{}"
# create an instance of ThemingPublicCapabilitiesTheming from a JSON string
theming_public_capabilities_theming_instance = ThemingPublicCapabilitiesTheming.from_json(json)
# print the JSON string representation of the object
print ThemingPublicCapabilitiesTheming.to_json()

# convert the object into a dict
theming_public_capabilities_theming_dict = theming_public_capabilities_theming_instance.to_dict()
# create an instance of ThemingPublicCapabilitiesTheming from a dict
theming_public_capabilities_theming_form_dict = theming_public_capabilities_theming.from_dict(theming_public_capabilities_theming_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


