# WeatherStatusForecastDataInstant


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**details** | [**WeatherStatusForecastDataInstantDetails**](WeatherStatusForecastDataInstantDetails.md) |  | 

## Example

```python
from nextcloud_client.models.weather_status_forecast_data_instant import WeatherStatusForecastDataInstant

# TODO update the JSON string below
json = "{}"
# create an instance of WeatherStatusForecastDataInstant from a JSON string
weather_status_forecast_data_instant_instance = WeatherStatusForecastDataInstant.from_json(json)
# print the JSON string representation of the object
print WeatherStatusForecastDataInstant.to_json()

# convert the object into a dict
weather_status_forecast_data_instant_dict = weather_status_forecast_data_instant_instance.to_dict()
# create an instance of WeatherStatusForecastDataInstant from a dict
weather_status_forecast_data_instant_form_dict = weather_status_forecast_data_instant.from_dict(weather_status_forecast_data_instant_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


