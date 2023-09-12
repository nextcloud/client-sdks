# WeatherStatusForecastDataNext1HoursDetails


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**precipitation_amount** | **float** |  | 
**precipitation_amount_max** | **float** |  | 
**precipitation_amount_min** | **float** |  | 
**probability_of_precipitation** | **float** |  | 
**probability_of_thunder** | **float** |  | 

## Example

```python
from nextcloud_client.models.weather_status_forecast_data_next1_hours_details import WeatherStatusForecastDataNext1HoursDetails

# TODO update the JSON string below
json = "{}"
# create an instance of WeatherStatusForecastDataNext1HoursDetails from a JSON string
weather_status_forecast_data_next1_hours_details_instance = WeatherStatusForecastDataNext1HoursDetails.from_json(json)
# print the JSON string representation of the object
print WeatherStatusForecastDataNext1HoursDetails.to_json()

# convert the object into a dict
weather_status_forecast_data_next1_hours_details_dict = weather_status_forecast_data_next1_hours_details_instance.to_dict()
# create an instance of WeatherStatusForecastDataNext1HoursDetails from a dict
weather_status_forecast_data_next1_hours_details_form_dict = weather_status_forecast_data_next1_hours_details.from_dict(weather_status_forecast_data_next1_hours_details_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


