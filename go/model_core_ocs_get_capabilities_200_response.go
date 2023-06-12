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

// checks if the CoreOcsGetCapabilities200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreOcsGetCapabilities200Response{}

// CoreOcsGetCapabilities200Response struct for CoreOcsGetCapabilities200Response
type CoreOcsGetCapabilities200Response struct {
	Ocs CoreOcsGetCapabilities200ResponseOcs `json:"ocs"`
}

// NewCoreOcsGetCapabilities200Response instantiates a new CoreOcsGetCapabilities200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreOcsGetCapabilities200Response(ocs CoreOcsGetCapabilities200ResponseOcs) *CoreOcsGetCapabilities200Response {
	this := CoreOcsGetCapabilities200Response{}
	this.Ocs = ocs
	return &this
}

// NewCoreOcsGetCapabilities200ResponseWithDefaults instantiates a new CoreOcsGetCapabilities200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreOcsGetCapabilities200ResponseWithDefaults() *CoreOcsGetCapabilities200Response {
	this := CoreOcsGetCapabilities200Response{}
	return &this
}

// GetOcs returns the Ocs field value
func (o *CoreOcsGetCapabilities200Response) GetOcs() CoreOcsGetCapabilities200ResponseOcs {
	if o == nil {
		var ret CoreOcsGetCapabilities200ResponseOcs
		return ret
	}

	return o.Ocs
}

// GetOcsOk returns a tuple with the Ocs field value
// and a boolean to check if the value has been set.
func (o *CoreOcsGetCapabilities200Response) GetOcsOk() (*CoreOcsGetCapabilities200ResponseOcs, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ocs, true
}

// SetOcs sets field value
func (o *CoreOcsGetCapabilities200Response) SetOcs(v CoreOcsGetCapabilities200ResponseOcs) {
	o.Ocs = v
}

func (o CoreOcsGetCapabilities200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreOcsGetCapabilities200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ocs"] = o.Ocs
	return toSerialize, nil
}

type NullableCoreOcsGetCapabilities200Response struct {
	value *CoreOcsGetCapabilities200Response
	isSet bool
}

func (v NullableCoreOcsGetCapabilities200Response) Get() *CoreOcsGetCapabilities200Response {
	return v.value
}

func (v *NullableCoreOcsGetCapabilities200Response) Set(val *CoreOcsGetCapabilities200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreOcsGetCapabilities200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreOcsGetCapabilities200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreOcsGetCapabilities200Response(val *CoreOcsGetCapabilities200Response) *NullableCoreOcsGetCapabilities200Response {
	return &NullableCoreOcsGetCapabilities200Response{value: val, isSet: true}
}

func (v NullableCoreOcsGetCapabilities200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreOcsGetCapabilities200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


