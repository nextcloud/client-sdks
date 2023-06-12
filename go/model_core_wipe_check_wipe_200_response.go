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

// checks if the CoreWipeCheckWipe200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreWipeCheckWipe200Response{}

// CoreWipeCheckWipe200Response struct for CoreWipeCheckWipe200Response
type CoreWipeCheckWipe200Response struct {
	Wipe bool `json:"wipe"`
}

// NewCoreWipeCheckWipe200Response instantiates a new CoreWipeCheckWipe200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreWipeCheckWipe200Response(wipe bool) *CoreWipeCheckWipe200Response {
	this := CoreWipeCheckWipe200Response{}
	this.Wipe = wipe
	return &this
}

// NewCoreWipeCheckWipe200ResponseWithDefaults instantiates a new CoreWipeCheckWipe200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreWipeCheckWipe200ResponseWithDefaults() *CoreWipeCheckWipe200Response {
	this := CoreWipeCheckWipe200Response{}
	return &this
}

// GetWipe returns the Wipe field value
func (o *CoreWipeCheckWipe200Response) GetWipe() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Wipe
}

// GetWipeOk returns a tuple with the Wipe field value
// and a boolean to check if the value has been set.
func (o *CoreWipeCheckWipe200Response) GetWipeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Wipe, true
}

// SetWipe sets field value
func (o *CoreWipeCheckWipe200Response) SetWipe(v bool) {
	o.Wipe = v
}

func (o CoreWipeCheckWipe200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreWipeCheckWipe200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["wipe"] = o.Wipe
	return toSerialize, nil
}

type NullableCoreWipeCheckWipe200Response struct {
	value *CoreWipeCheckWipe200Response
	isSet bool
}

func (v NullableCoreWipeCheckWipe200Response) Get() *CoreWipeCheckWipe200Response {
	return v.value
}

func (v *NullableCoreWipeCheckWipe200Response) Set(val *CoreWipeCheckWipe200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreWipeCheckWipe200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreWipeCheckWipe200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreWipeCheckWipe200Response(val *CoreWipeCheckWipe200Response) *NullableCoreWipeCheckWipe200Response {
	return &NullableCoreWipeCheckWipe200Response{value: val, isSet: true}
}

func (v NullableCoreWipeCheckWipe200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreWipeCheckWipe200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


