# WeatherStatusWeatherStatusGetForecast200ResponseOcs


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**meta** | [**OCSMeta**](OCSMeta.md) |  | 
**data** | [**List[WeatherStatusForecast]**](WeatherStatusForecast.md) |  | 

## Example

```python
from nextcloud_client.models.weather_status_weather_status_get_forecast200_response_ocs import WeatherStatusWeatherStatusGetForecast200ResponseOcs

# TODO update the JSON string below
json = "{}"
# create an instance of WeatherStatusWeatherStatusGetForecast200ResponseOcs from a JSON string
weather_status_weather_status_get_forecast200_response_ocs_instance = WeatherStatusWeatherStatusGetForecast200ResponseOcs.from_json(json)
# print the JSON string representation of the object
print WeatherStatusWeatherStatusGetForecast200ResponseOcs.to_json()

# convert the object into a dict
weather_status_weather_status_get_forecast200_response_ocs_dict = weather_status_weather_status_get_forecast200_response_ocs_instance.to_dict()
# create an instance of WeatherStatusWeatherStatusGetForecast200ResponseOcs from a dict
weather_status_weather_status_get_forecast200_response_ocs_form_dict = weather_status_weather_status_get_forecast200_response_ocs.from_dict(weather_status_weather_status_get_forecast200_response_ocs_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


