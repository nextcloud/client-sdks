# WeatherStatusForecastDataInstantDetails


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**air_pressure_at_sea_level** | **float** |  | 
**air_temperature** | **float** |  | 
**cloud_area_fraction** | **float** |  | 
**cloud_area_fraction_high** | **float** |  | 
**cloud_area_fraction_low** | **float** |  | 
**cloud_area_fraction_medium** | **float** |  | 
**dew_point_temperature** | **float** |  | 
**fog_area_fraction** | **float** |  | 
**relative_humidity** | **float** |  | 
**ultraviolet_index_clear_sky** | **float** |  | 
**wind_from_direction** | **float** |  | 
**wind_speed** | **float** |  | 
**wind_speed_of_gust** | **float** |  | 

## Example

```python
from nextcloud_client.models.weather_status_forecast_data_instant_details import WeatherStatusForecastDataInstantDetails

# TODO update the JSON string below
json = "{}"
# create an instance of WeatherStatusForecastDataInstantDetails from a JSON string
weather_status_forecast_data_instant_details_instance = WeatherStatusForecastDataInstantDetails.from_json(json)
# print the JSON string representation of the object
print WeatherStatusForecastDataInstantDetails.to_json()

# convert the object into a dict
weather_status_forecast_data_instant_details_dict = weather_status_forecast_data_instant_details_instance.to_dict()
# create an instance of WeatherStatusForecastDataInstantDetails from a dict
weather_status_forecast_data_instant_details_form_dict = weather_status_forecast_data_instant_details.from_dict(weather_status_forecast_data_instant_details_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


