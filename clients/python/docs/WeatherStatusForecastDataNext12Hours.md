# WeatherStatusForecastDataNext12Hours


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**summary** | [**WeatherStatusForecastDataNext12HoursSummary**](WeatherStatusForecastDataNext12HoursSummary.md) |  | 
**details** | [**WeatherStatusForecastDataNext12HoursDetails**](WeatherStatusForecastDataNext12HoursDetails.md) |  | 

## Example

```python
from nextcloud_client.models.weather_status_forecast_data_next12_hours import WeatherStatusForecastDataNext12Hours

# TODO update the JSON string below
json = "{}"
# create an instance of WeatherStatusForecastDataNext12Hours from a JSON string
weather_status_forecast_data_next12_hours_instance = WeatherStatusForecastDataNext12Hours.from_json(json)
# print the JSON string representation of the object
print WeatherStatusForecastDataNext12Hours.to_json()

# convert the object into a dict
weather_status_forecast_data_next12_hours_dict = weather_status_forecast_data_next12_hours_instance.to_dict()
# create an instance of WeatherStatusForecastDataNext12Hours from a dict
weather_status_forecast_data_next12_hours_form_dict = weather_status_forecast_data_next12_hours.from_dict(weather_status_forecast_data_next12_hours_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


