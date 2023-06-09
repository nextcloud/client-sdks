/*
nextcloud

Nextcloud APIs

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the WeatherStatusForecastDataNext1HoursDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WeatherStatusForecastDataNext1HoursDetails{}

// WeatherStatusForecastDataNext1HoursDetails struct for WeatherStatusForecastDataNext1HoursDetails
type WeatherStatusForecastDataNext1HoursDetails struct {
	PrecipitationAmount float32 `json:"precipitation_amount"`
	PrecipitationAmountMax float32 `json:"precipitation_amount_max"`
	PrecipitationAmountMin float32 `json:"precipitation_amount_min"`
	ProbabilityOfPrecipitation float32 `json:"probability_of_precipitation"`
	ProbabilityOfThunder float32 `json:"probability_of_thunder"`
}

// NewWeatherStatusForecastDataNext1HoursDetails instantiates a new WeatherStatusForecastDataNext1HoursDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeatherStatusForecastDataNext1HoursDetails(precipitationAmount float32, precipitationAmountMax float32, precipitationAmountMin float32, probabilityOfPrecipitation float32, probabilityOfThunder float32) *WeatherStatusForecastDataNext1HoursDetails {
	this := WeatherStatusForecastDataNext1HoursDetails{}
	this.PrecipitationAmount = precipitationAmount
	this.PrecipitationAmountMax = precipitationAmountMax
	this.PrecipitationAmountMin = precipitationAmountMin
	this.ProbabilityOfPrecipitation = probabilityOfPrecipitation
	this.ProbabilityOfThunder = probabilityOfThunder
	return &this
}

// NewWeatherStatusForecastDataNext1HoursDetailsWithDefaults instantiates a new WeatherStatusForecastDataNext1HoursDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeatherStatusForecastDataNext1HoursDetailsWithDefaults() *WeatherStatusForecastDataNext1HoursDetails {
	this := WeatherStatusForecastDataNext1HoursDetails{}
	return &this
}

// GetPrecipitationAmount returns the PrecipitationAmount field value
func (o *WeatherStatusForecastDataNext1HoursDetails) GetPrecipitationAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PrecipitationAmount
}

// GetPrecipitationAmountOk returns a tuple with the PrecipitationAmount field value
// and a boolean to check if the value has been set.
func (o *WeatherStatusForecastDataNext1HoursDetails) GetPrecipitationAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrecipitationAmount, true
}

// SetPrecipitationAmount sets field value
func (o *WeatherStatusForecastDataNext1HoursDetails) SetPrecipitationAmount(v float32) {
	o.PrecipitationAmount = v
}

// GetPrecipitationAmountMax returns the PrecipitationAmountMax field value
func (o *WeatherStatusForecastDataNext1HoursDetails) GetPrecipitationAmountMax() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PrecipitationAmountMax
}

// GetPrecipitationAmountMaxOk returns a tuple with the PrecipitationAmountMax field value
// and a boolean to check if the value has been set.
func (o *WeatherStatusForecastDataNext1HoursDetails) GetPrecipitationAmountMaxOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrecipitationAmountMax, true
}

// SetPrecipitationAmountMax sets field value
func (o *WeatherStatusForecastDataNext1HoursDetails) SetPrecipitationAmountMax(v float32) {
	o.PrecipitationAmountMax = v
}

// GetPrecipitationAmountMin returns the PrecipitationAmountMin field value
func (o *WeatherStatusForecastDataNext1HoursDetails) GetPrecipitationAmountMin() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PrecipitationAmountMin
}

// GetPrecipitationAmountMinOk returns a tuple with the PrecipitationAmountMin field value
// and a boolean to check if the value has been set.
func (o *WeatherStatusForecastDataNext1HoursDetails) GetPrecipitationAmountMinOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrecipitationAmountMin, true
}

// SetPrecipitationAmountMin sets field value
func (o *WeatherStatusForecastDataNext1HoursDetails) SetPrecipitationAmountMin(v float32) {
	o.PrecipitationAmountMin = v
}

// GetProbabilityOfPrecipitation returns the ProbabilityOfPrecipitation field value
func (o *WeatherStatusForecastDataNext1HoursDetails) GetProbabilityOfPrecipitation() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ProbabilityOfPrecipitation
}

// GetProbabilityOfPrecipitationOk returns a tuple with the ProbabilityOfPrecipitation field value
// and a boolean to check if the value has been set.
func (o *WeatherStatusForecastDataNext1HoursDetails) GetProbabilityOfPrecipitationOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProbabilityOfPrecipitation, true
}

// SetProbabilityOfPrecipitation sets field value
func (o *WeatherStatusForecastDataNext1HoursDetails) SetProbabilityOfPrecipitation(v float32) {
	o.ProbabilityOfPrecipitation = v
}

// GetProbabilityOfThunder returns the ProbabilityOfThunder field value
func (o *WeatherStatusForecastDataNext1HoursDetails) GetProbabilityOfThunder() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ProbabilityOfThunder
}

// GetProbabilityOfThunderOk returns a tuple with the ProbabilityOfThunder field value
// and a boolean to check if the value has been set.
func (o *WeatherStatusForecastDataNext1HoursDetails) GetProbabilityOfThunderOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProbabilityOfThunder, true
}

// SetProbabilityOfThunder sets field value
func (o *WeatherStatusForecastDataNext1HoursDetails) SetProbabilityOfThunder(v float32) {
	o.ProbabilityOfThunder = v
}

func (o WeatherStatusForecastDataNext1HoursDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WeatherStatusForecastDataNext1HoursDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["precipitation_amount"] = o.PrecipitationAmount
	toSerialize["precipitation_amount_max"] = o.PrecipitationAmountMax
	toSerialize["precipitation_amount_min"] = o.PrecipitationAmountMin
	toSerialize["probability_of_precipitation"] = o.ProbabilityOfPrecipitation
	toSerialize["probability_of_thunder"] = o.ProbabilityOfThunder
	return toSerialize, nil
}

type NullableWeatherStatusForecastDataNext1HoursDetails struct {
	value *WeatherStatusForecastDataNext1HoursDetails
	isSet bool
}

func (v NullableWeatherStatusForecastDataNext1HoursDetails) Get() *WeatherStatusForecastDataNext1HoursDetails {
	return v.value
}

func (v *NullableWeatherStatusForecastDataNext1HoursDetails) Set(val *WeatherStatusForecastDataNext1HoursDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableWeatherStatusForecastDataNext1HoursDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableWeatherStatusForecastDataNext1HoursDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeatherStatusForecastDataNext1HoursDetails(val *WeatherStatusForecastDataNext1HoursDetails) *NullableWeatherStatusForecastDataNext1HoursDetails {
	return &NullableWeatherStatusForecastDataNext1HoursDetails{value: val, isSet: true}
}

func (v NullableWeatherStatusForecastDataNext1HoursDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeatherStatusForecastDataNext1HoursDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


