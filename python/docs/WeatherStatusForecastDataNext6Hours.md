# WeatherStatusForecastDataNext6Hours


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**summary** | [**WeatherStatusForecastDataNext12HoursSummary**](WeatherStatusForecastDataNext12HoursSummary.md) |  | 
**details** | [**WeatherStatusForecastDataNext6HoursDetails**](WeatherStatusForecastDataNext6HoursDetails.md) |  | 

## Example

```python
from openapi_client.models.weather_status_forecast_data_next6_hours import WeatherStatusForecastDataNext6Hours

# TODO update the JSON string below
json = "{}"
# create an instance of WeatherStatusForecastDataNext6Hours from a JSON string
weather_status_forecast_data_next6_hours_instance = WeatherStatusForecastDataNext6Hours.from_json(json)
# print the JSON string representation of the object
print WeatherStatusForecastDataNext6Hours.to_json()

# convert the object into a dict
weather_status_forecast_data_next6_hours_dict = weather_status_forecast_data_next6_hours_instance.to_dict()
# create an instance of WeatherStatusForecastDataNext6Hours from a dict
weather_status_forecast_data_next6_hours_form_dict = weather_status_forecast_data_next6_hours.from_dict(weather_status_forecast_data_next6_hours_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


