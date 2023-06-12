# WeatherStatusForecastDataNext1Hours


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**summary** | [**WeatherStatusForecastDataNext12HoursSummary**](WeatherStatusForecastDataNext12HoursSummary.md) |  | 
**details** | [**WeatherStatusForecastDataNext1HoursDetails**](WeatherStatusForecastDataNext1HoursDetails.md) |  | 

## Example

```python
from openapi_client.models.weather_status_forecast_data_next1_hours import WeatherStatusForecastDataNext1Hours

# TODO update the JSON string below
json = "{}"
# create an instance of WeatherStatusForecastDataNext1Hours from a JSON string
weather_status_forecast_data_next1_hours_instance = WeatherStatusForecastDataNext1Hours.from_json(json)
# print the JSON string representation of the object
print WeatherStatusForecastDataNext1Hours.to_json()

# convert the object into a dict
weather_status_forecast_data_next1_hours_dict = weather_status_forecast_data_next1_hours_instance.to_dict()
# create an instance of WeatherStatusForecastDataNext1Hours from a dict
weather_status_forecast_data_next1_hours_form_dict = weather_status_forecast_data_next1_hours.from_dict(weather_status_forecast_data_next1_hours_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


