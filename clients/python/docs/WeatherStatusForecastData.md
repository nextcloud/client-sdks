# WeatherStatusForecastData


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**instant** | [**WeatherStatusForecastDataInstant**](WeatherStatusForecastDataInstant.md) |  | 
**next_12_hours** | [**WeatherStatusForecastDataNext12Hours**](WeatherStatusForecastDataNext12Hours.md) |  | 
**next_1_hours** | [**WeatherStatusForecastDataNext1Hours**](WeatherStatusForecastDataNext1Hours.md) |  | 
**next_6_hours** | [**WeatherStatusForecastDataNext6Hours**](WeatherStatusForecastDataNext6Hours.md) |  | 

## Example

```python
from nextcloud_client.models.weather_status_forecast_data import WeatherStatusForecastData

# TODO update the JSON string below
json = "{}"
# create an instance of WeatherStatusForecastData from a JSON string
weather_status_forecast_data_instance = WeatherStatusForecastData.from_json(json)
# print the JSON string representation of the object
print WeatherStatusForecastData.to_json()

# convert the object into a dict
weather_status_forecast_data_dict = weather_status_forecast_data_instance.to_dict()
# create an instance of WeatherStatusForecastData from a dict
weather_status_forecast_data_form_dict = weather_status_forecast_data.from_dict(weather_status_forecast_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


