# WeatherStatusForecastDataNext6HoursDetails


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**air_temperature_max** | **float** |  | 
**air_temperature_min** | **float** |  | 
**precipitation_amount** | **float** |  | 
**precipitation_amount_max** | **float** |  | 
**precipitation_amount_min** | **float** |  | 
**probability_of_precipitation** | **float** |  | 

## Example

```python
from openapi_client.models.weather_status_forecast_data_next6_hours_details import WeatherStatusForecastDataNext6HoursDetails

# TODO update the JSON string below
json = "{}"
# create an instance of WeatherStatusForecastDataNext6HoursDetails from a JSON string
weather_status_forecast_data_next6_hours_details_instance = WeatherStatusForecastDataNext6HoursDetails.from_json(json)
# print the JSON string representation of the object
print WeatherStatusForecastDataNext6HoursDetails.to_json()

# convert the object into a dict
weather_status_forecast_data_next6_hours_details_dict = weather_status_forecast_data_next6_hours_details_instance.to_dict()
# create an instance of WeatherStatusForecastDataNext6HoursDetails from a dict
weather_status_forecast_data_next6_hours_details_form_dict = weather_status_forecast_data_next6_hours_details.from_dict(weather_status_forecast_data_next6_hours_details_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


