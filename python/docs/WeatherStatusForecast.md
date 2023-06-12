# WeatherStatusForecast


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**time** | **str** |  | 
**data** | [**WeatherStatusForecastData**](WeatherStatusForecastData.md) |  | 

## Example

```python
from openapi_client.models.weather_status_forecast import WeatherStatusForecast

# TODO update the JSON string below
json = "{}"
# create an instance of WeatherStatusForecast from a JSON string
weather_status_forecast_instance = WeatherStatusForecast.from_json(json)
# print the JSON string representation of the object
print WeatherStatusForecast.to_json()

# convert the object into a dict
weather_status_forecast_dict = weather_status_forecast_instance.to_dict()
# create an instance of WeatherStatusForecast from a dict
weather_status_forecast_form_dict = weather_status_forecast.from_dict(weather_status_forecast_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


