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

// checks if the WeatherStatusWeatherStatusGetLocation200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WeatherStatusWeatherStatusGetLocation200Response{}

// WeatherStatusWeatherStatusGetLocation200Response struct for WeatherStatusWeatherStatusGetLocation200Response
type WeatherStatusWeatherStatusGetLocation200Response struct {
	Ocs WeatherStatusWeatherStatusGetLocation200ResponseOcs `json:"ocs"`
}

// NewWeatherStatusWeatherStatusGetLocation200Response instantiates a new WeatherStatusWeatherStatusGetLocation200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeatherStatusWeatherStatusGetLocation200Response(ocs WeatherStatusWeatherStatusGetLocation200ResponseOcs) *WeatherStatusWeatherStatusGetLocation200Response {
	this := WeatherStatusWeatherStatusGetLocation200Response{}
	this.Ocs = ocs
	return &this
}

// NewWeatherStatusWeatherStatusGetLocation200ResponseWithDefaults instantiates a new WeatherStatusWeatherStatusGetLocation200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeatherStatusWeatherStatusGetLocation200ResponseWithDefaults() *WeatherStatusWeatherStatusGetLocation200Response {
	this := WeatherStatusWeatherStatusGetLocation200Response{}
	return &this
}

// GetOcs returns the Ocs field value
func (o *WeatherStatusWeatherStatusGetLocation200Response) GetOcs() WeatherStatusWeatherStatusGetLocation200ResponseOcs {
	if o == nil {
		var ret WeatherStatusWeatherStatusGetLocation200ResponseOcs
		return ret
	}

	return o.Ocs
}

// GetOcsOk returns a tuple with the Ocs field value
// and a boolean to check if the value has been set.
func (o *WeatherStatusWeatherStatusGetLocation200Response) GetOcsOk() (*WeatherStatusWeatherStatusGetLocation200ResponseOcs, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ocs, true
}

// SetOcs sets field value
func (o *WeatherStatusWeatherStatusGetLocation200Response) SetOcs(v WeatherStatusWeatherStatusGetLocation200ResponseOcs) {
	o.Ocs = v
}

func (o WeatherStatusWeatherStatusGetLocation200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WeatherStatusWeatherStatusGetLocation200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ocs"] = o.Ocs
	return toSerialize, nil
}

type NullableWeatherStatusWeatherStatusGetLocation200Response struct {
	value *WeatherStatusWeatherStatusGetLocation200Response
	isSet bool
}

func (v NullableWeatherStatusWeatherStatusGetLocation200Response) Get() *WeatherStatusWeatherStatusGetLocation200Response {
	return v.value
}

func (v *NullableWeatherStatusWeatherStatusGetLocation200Response) Set(val *WeatherStatusWeatherStatusGetLocation200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableWeatherStatusWeatherStatusGetLocation200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableWeatherStatusWeatherStatusGetLocation200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeatherStatusWeatherStatusGetLocation200Response(val *WeatherStatusWeatherStatusGetLocation200Response) *NullableWeatherStatusWeatherStatusGetLocation200Response {
	return &NullableWeatherStatusWeatherStatusGetLocation200Response{value: val, isSet: true}
}

func (v NullableWeatherStatusWeatherStatusGetLocation200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeatherStatusWeatherStatusGetLocation200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


