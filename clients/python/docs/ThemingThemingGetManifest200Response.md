# ThemingThemingGetManifest200Response


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** |  | 
**short_name** | **str** |  | 
**start_url** | **str** |  | 
**theme_color** | **str** |  | 
**background_color** | **str** |  | 
**description** | **str** |  | 
**icons** | [**List[ThemingThemingGetManifest200ResponseIconsInner]**](ThemingThemingGetManifest200ResponseIconsInner.md) |  | 
**display** | **str** |  | 

## Example

```python
from nextcloud_client.models.theming_theming_get_manifest200_response import ThemingThemingGetManifest200Response

# TODO update the JSON string below
json = "{}"
# create an instance of ThemingThemingGetManifest200Response from a JSON string
theming_theming_get_manifest200_response_instance = ThemingThemingGetManifest200Response.from_json(json)
# print the JSON string representation of the object
print ThemingThemingGetManifest200Response.to_json()

# convert the object into a dict
theming_theming_get_manifest200_response_dict = theming_theming_get_manifest200_response_instance.to_dict()
# create an instance of ThemingThemingGetManifest200Response from a dict
theming_theming_get_manifest200_response_form_dict = theming_theming_get_manifest200_response.from_dict(theming_theming_get_manifest200_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


