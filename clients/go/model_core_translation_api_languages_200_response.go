/*
nextcloud

Nextcloud APIs

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client_sdk

import (
	"encoding/json"
)

// checks if the CoreTranslationApiLanguages200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreTranslationApiLanguages200Response{}

// CoreTranslationApiLanguages200Response struct for CoreTranslationApiLanguages200Response
type CoreTranslationApiLanguages200Response struct {
	Ocs CoreTranslationApiLanguages200ResponseOcs `json:"ocs"`
	AdditionalProperties map[string]interface{}
}

type _CoreTranslationApiLanguages200Response CoreTranslationApiLanguages200Response

// NewCoreTranslationApiLanguages200Response instantiates a new CoreTranslationApiLanguages200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreTranslationApiLanguages200Response(ocs CoreTranslationApiLanguages200ResponseOcs) *CoreTranslationApiLanguages200Response {
	this := CoreTranslationApiLanguages200Response{}
	this.Ocs = ocs
	return &this
}

// NewCoreTranslationApiLanguages200ResponseWithDefaults instantiates a new CoreTranslationApiLanguages200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreTranslationApiLanguages200ResponseWithDefaults() *CoreTranslationApiLanguages200Response {
	this := CoreTranslationApiLanguages200Response{}
	return &this
}

// GetOcs returns the Ocs field value
func (o *CoreTranslationApiLanguages200Response) GetOcs() CoreTranslationApiLanguages200ResponseOcs {
	if o == nil {
		var ret CoreTranslationApiLanguages200ResponseOcs
		return ret
	}

	return o.Ocs
}

// GetOcsOk returns a tuple with the Ocs field value
// and a boolean to check if the value has been set.
func (o *CoreTranslationApiLanguages200Response) GetOcsOk() (*CoreTranslationApiLanguages200ResponseOcs, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ocs, true
}

// SetOcs sets field value
func (o *CoreTranslationApiLanguages200Response) SetOcs(v CoreTranslationApiLanguages200ResponseOcs) {
	o.Ocs = v
}

func (o CoreTranslationApiLanguages200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreTranslationApiLanguages200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ocs"] = o.Ocs

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CoreTranslationApiLanguages200Response) UnmarshalJSON(bytes []byte) (err error) {
	varCoreTranslationApiLanguages200Response := _CoreTranslationApiLanguages200Response{}

	if err = json.Unmarshal(bytes, &varCoreTranslationApiLanguages200Response); err == nil {
		*o = CoreTranslationApiLanguages200Response(varCoreTranslationApiLanguages200Response)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ocs")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCoreTranslationApiLanguages200Response struct {
	value *CoreTranslationApiLanguages200Response
	isSet bool
}

func (v NullableCoreTranslationApiLanguages200Response) Get() *CoreTranslationApiLanguages200Response {
	return v.value
}

func (v *NullableCoreTranslationApiLanguages200Response) Set(val *CoreTranslationApiLanguages200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreTranslationApiLanguages200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreTranslationApiLanguages200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreTranslationApiLanguages200Response(val *CoreTranslationApiLanguages200Response) *NullableCoreTranslationApiLanguages200Response {
	return &NullableCoreTranslationApiLanguages200Response{value: val, isSet: true}
}

func (v NullableCoreTranslationApiLanguages200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreTranslationApiLanguages200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


