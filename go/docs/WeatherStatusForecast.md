# WeatherStatusForecast

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | **string** |  | 
**Data** | [**WeatherStatusForecastData**](WeatherStatusForecastData.md) |  | 

## Methods

### NewWeatherStatusForecast

`func NewWeatherStatusForecast(time string, data WeatherStatusForecastData, ) *WeatherStatusForecast`

NewWeatherStatusForecast instantiates a new WeatherStatusForecast object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWeatherStatusForecastWithDefaults

`func NewWeatherStatusForecastWithDefaults() *WeatherStatusForecast`

NewWeatherStatusForecastWithDefaults instantiates a new WeatherStatusForecast object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *WeatherStatusForecast) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *WeatherStatusForecast) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *WeatherStatusForecast) SetTime(v string)`

SetTime sets Time field to given value.


### GetData

`func (o *WeatherStatusForecast) GetData() WeatherStatusForecastData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WeatherStatusForecast) GetDataOk() (*WeatherStatusForecastData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WeatherStatusForecast) SetData(v WeatherStatusForecastData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


