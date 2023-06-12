# \WeatherStatusWeatherStatusApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**weather_status_weather_status_get_favorites**](WeatherStatusWeatherStatusApi.md#weather_status_weather_status_get_favorites) | **GET** /ocs/v2.php/apps/weather_status/api/v1/favorites | Get favorites list
[**weather_status_weather_status_get_forecast**](WeatherStatusWeatherStatusApi.md#weather_status_weather_status_get_forecast) | **GET** /ocs/v2.php/apps/weather_status/api/v1/forecast | Get forecast for current location
[**weather_status_weather_status_get_location**](WeatherStatusWeatherStatusApi.md#weather_status_weather_status_get_location) | **GET** /ocs/v2.php/apps/weather_status/api/v1/location | Get stored user location
[**weather_status_weather_status_set_favorites**](WeatherStatusWeatherStatusApi.md#weather_status_weather_status_set_favorites) | **PUT** /ocs/v2.php/apps/weather_status/api/v1/favorites | Set favorites list
[**weather_status_weather_status_set_location**](WeatherStatusWeatherStatusApi.md#weather_status_weather_status_set_location) | **PUT** /ocs/v2.php/apps/weather_status/api/v1/location | Set address and resolve it to get coordinates or directly set coordinates and get address with reverse geocoding
[**weather_status_weather_status_set_mode**](WeatherStatusWeatherStatusApi.md#weather_status_weather_status_set_mode) | **PUT** /ocs/v2.php/apps/weather_status/api/v1/mode | Change the weather status mode. There are currently 2 modes: - ask the browser - use the user defined address
[**weather_status_weather_status_use_personal_address**](WeatherStatusWeatherStatusApi.md#weather_status_weather_status_use_personal_address) | **PUT** /ocs/v2.php/apps/weather_status/api/v1/use-personal | Try to use the address set in user personal settings as weather location



## weather_status_weather_status_get_favorites

> crate::models::ProvisioningApiGroupsGetSubAdminsOfGroup200Response weather_status_weather_status_get_favorites(ocs_api_request)
Get favorites list

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::ProvisioningApiGroupsGetSubAdminsOfGroup200Response**](provisioning_api_groups_get_sub_admins_of_group_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## weather_status_weather_status_get_forecast

> crate::models::WeatherStatusWeatherStatusGetForecast200Response weather_status_weather_status_get_forecast(ocs_api_request)
Get forecast for current location

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::WeatherStatusWeatherStatusGetForecast200Response**](weather_status_weather_status_get_forecast_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## weather_status_weather_status_get_location

> crate::models::WeatherStatusWeatherStatusGetLocation200Response weather_status_weather_status_get_location(ocs_api_request)
Get stored user location

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::WeatherStatusWeatherStatusGetLocation200Response**](weather_status_weather_status_get_location_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## weather_status_weather_status_set_favorites

> crate::models::CoreReferenceApiTouchProvider200Response weather_status_weather_status_set_favorites(favorites, ocs_api_request)
Set favorites list

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**favorites** | **String** | Favorite addresses | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::CoreReferenceApiTouchProvider200Response**](core_reference_api_touch_provider_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## weather_status_weather_status_set_location

> crate::models::WeatherStatusWeatherStatusUsePersonalAddress200Response weather_status_weather_status_set_location(ocs_api_request, address, lat, lon)
Set address and resolve it to get coordinates or directly set coordinates and get address with reverse geocoding

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**ocs_api_request** | **String** |  | [required] |[default to true]
**address** | Option<**String**> | Any approximative or exact address |  |
**lat** | Option<**f32**> | Latitude in decimal degree format |  |
**lon** | Option<**f32**> | Longitude in decimal degree format |  |

### Return type

[**crate::models::WeatherStatusWeatherStatusUsePersonalAddress200Response**](weather_status_weather_status_use_personal_address_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## weather_status_weather_status_set_mode

> crate::models::CoreReferenceApiTouchProvider200Response weather_status_weather_status_set_mode(mode, ocs_api_request)
Change the weather status mode. There are currently 2 modes: - ask the browser - use the user defined address

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**mode** | **i64** | New mode | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::CoreReferenceApiTouchProvider200Response**](core_reference_api_touch_provider_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## weather_status_weather_status_use_personal_address

> crate::models::WeatherStatusWeatherStatusUsePersonalAddress200Response weather_status_weather_status_use_personal_address(ocs_api_request)
Try to use the address set in user personal settings as weather location

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::WeatherStatusWeatherStatusUsePersonalAddress200Response**](weather_status_weather_status_use_personal_address_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

