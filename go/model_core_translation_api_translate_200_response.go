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

// checks if the CoreTranslationApiTranslate200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreTranslationApiTranslate200Response{}

// CoreTranslationApiTranslate200Response struct for CoreTranslationApiTranslate200Response
type CoreTranslationApiTranslate200Response struct {
	Ocs CoreTranslationApiTranslate200ResponseOcs `json:"ocs"`
}

// NewCoreTranslationApiTranslate200Response instantiates a new CoreTranslationApiTranslate200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreTranslationApiTranslate200Response(ocs CoreTranslationApiTranslate200ResponseOcs) *CoreTranslationApiTranslate200Response {
	this := CoreTranslationApiTranslate200Response{}
	this.Ocs = ocs
	return &this
}

// NewCoreTranslationApiTranslate200ResponseWithDefaults instantiates a new CoreTranslationApiTranslate200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreTranslationApiTranslate200ResponseWithDefaults() *CoreTranslationApiTranslate200Response {
	this := CoreTranslationApiTranslate200Response{}
	return &this
}

// GetOcs returns the Ocs field value
func (o *CoreTranslationApiTranslate200Response) GetOcs() CoreTranslationApiTranslate200ResponseOcs {
	if o == nil {
		var ret CoreTranslationApiTranslate200ResponseOcs
		return ret
	}

	return o.Ocs
}

// GetOcsOk returns a tuple with the Ocs field value
// and a boolean to check if the value has been set.
func (o *CoreTranslationApiTranslate200Response) GetOcsOk() (*CoreTranslationApiTranslate200ResponseOcs, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ocs, true
}

// SetOcs sets field value
func (o *CoreTranslationApiTranslate200Response) SetOcs(v CoreTranslationApiTranslate200ResponseOcs) {
	o.Ocs = v
}

func (o CoreTranslationApiTranslate200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreTranslationApiTranslate200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ocs"] = o.Ocs
	return toSerialize, nil
}

type NullableCoreTranslationApiTranslate200Response struct {
	value *CoreTranslationApiTranslate200Response
	isSet bool
}

func (v NullableCoreTranslationApiTranslate200Response) Get() *CoreTranslationApiTranslate200Response {
	return v.value
}

func (v *NullableCoreTranslationApiTranslate200Response) Set(val *CoreTranslationApiTranslate200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreTranslationApiTranslate200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreTranslationApiTranslate200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreTranslationApiTranslate200Response(val *CoreTranslationApiTranslate200Response) *NullableCoreTranslationApiTranslate200Response {
	return &NullableCoreTranslationApiTranslate200Response{value: val, isSet: true}
}

func (v NullableCoreTranslationApiTranslate200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreTranslationApiTranslate200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


