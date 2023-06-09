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

// checks if the WeatherStatusWeatherStatusGetForecast200ResponseOcs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WeatherStatusWeatherStatusGetForecast200ResponseOcs{}

// WeatherStatusWeatherStatusGetForecast200ResponseOcs struct for WeatherStatusWeatherStatusGetForecast200ResponseOcs
type WeatherStatusWeatherStatusGetForecast200ResponseOcs struct {
	Meta OCSMeta `json:"meta"`
	Data []WeatherStatusForecast `json:"data"`
}

// NewWeatherStatusWeatherStatusGetForecast200ResponseOcs instantiates a new WeatherStatusWeatherStatusGetForecast200ResponseOcs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeatherStatusWeatherStatusGetForecast200ResponseOcs(meta OCSMeta, data []WeatherStatusForecast) *WeatherStatusWeatherStatusGetForecast200ResponseOcs {
	this := WeatherStatusWeatherStatusGetForecast200ResponseOcs{}
	this.Meta = meta
	this.Data = data
	return &this
}

// NewWeatherStatusWeatherStatusGetForecast200ResponseOcsWithDefaults instantiates a new WeatherStatusWeatherStatusGetForecast200ResponseOcs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeatherStatusWeatherStatusGetForecast200ResponseOcsWithDefaults() *WeatherStatusWeatherStatusGetForecast200ResponseOcs {
	this := WeatherStatusWeatherStatusGetForecast200ResponseOcs{}
	return &this
}

// GetMeta returns the Meta field value
func (o *WeatherStatusWeatherStatusGetForecast200ResponseOcs) GetMeta() OCSMeta {
	if o == nil {
		var ret OCSMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *WeatherStatusWeatherStatusGetForecast200ResponseOcs) GetMetaOk() (*OCSMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *WeatherStatusWeatherStatusGetForecast200ResponseOcs) SetMeta(v OCSMeta) {
	o.Meta = v
}

// GetData returns the Data field value
func (o *WeatherStatusWeatherStatusGetForecast200ResponseOcs) GetData() []WeatherStatusForecast {
	if o == nil {
		var ret []WeatherStatusForecast
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *WeatherStatusWeatherStatusGetForecast200ResponseOcs) GetDataOk() ([]WeatherStatusForecast, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *WeatherStatusWeatherStatusGetForecast200ResponseOcs) SetData(v []WeatherStatusForecast) {
	o.Data = v
}

func (o WeatherStatusWeatherStatusGetForecast200ResponseOcs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WeatherStatusWeatherStatusGetForecast200ResponseOcs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableWeatherStatusWeatherStatusGetForecast200ResponseOcs struct {
	value *WeatherStatusWeatherStatusGetForecast200ResponseOcs
	isSet bool
}

func (v NullableWeatherStatusWeatherStatusGetForecast200ResponseOcs) Get() *WeatherStatusWeatherStatusGetForecast200ResponseOcs {
	return v.value
}

func (v *NullableWeatherStatusWeatherStatusGetForecast200ResponseOcs) Set(val *WeatherStatusWeatherStatusGetForecast200ResponseOcs) {
	v.value = val
	v.isSet = true
}

func (v NullableWeatherStatusWeatherStatusGetForecast200ResponseOcs) IsSet() bool {
	return v.isSet
}

func (v *NullableWeatherStatusWeatherStatusGetForecast200ResponseOcs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeatherStatusWeatherStatusGetForecast200ResponseOcs(val *WeatherStatusWeatherStatusGetForecast200ResponseOcs) *NullableWeatherStatusWeatherStatusGetForecast200ResponseOcs {
	return &NullableWeatherStatusWeatherStatusGetForecast200ResponseOcs{value: val, isSet: true}
}

func (v NullableWeatherStatusWeatherStatusGetForecast200ResponseOcs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeatherStatusWeatherStatusGetForecast200ResponseOcs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


