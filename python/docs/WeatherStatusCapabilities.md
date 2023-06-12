# WeatherStatusCapabilities


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**weather_status** | [**CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUserExpireDate**](CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUserExpireDate.md) |  | 

## Example

```python
from openapi_client.models.weather_status_capabilities import WeatherStatusCapabilities

# TODO update the JSON string below
json = "{}"
# create an instance of WeatherStatusCapabilities from a JSON string
weather_status_capabilities_instance = WeatherStatusCapabilities.from_json(json)
# print the JSON string representation of the object
print WeatherStatusCapabilities.to_json()

# convert the object into a dict
weather_status_capabilities_dict = weather_status_capabilities_instance.to_dict()
# create an instance of WeatherStatusCapabilities from a dict
weather_status_capabilities_form_dict = weather_status_capabilities.from_dict(weather_status_capabilities_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


