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

// checks if the CoreTranslationApiTranslate200ResponseOcs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreTranslationApiTranslate200ResponseOcs{}

// CoreTranslationApiTranslate200ResponseOcs struct for CoreTranslationApiTranslate200ResponseOcs
type CoreTranslationApiTranslate200ResponseOcs struct {
	Meta OCSMeta `json:"meta"`
	Data CoreTranslationApiTranslate200ResponseOcsData `json:"data"`
	AdditionalProperties map[string]interface{}
}

type _CoreTranslationApiTranslate200ResponseOcs CoreTranslationApiTranslate200ResponseOcs

// NewCoreTranslationApiTranslate200ResponseOcs instantiates a new CoreTranslationApiTranslate200ResponseOcs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreTranslationApiTranslate200ResponseOcs(meta OCSMeta, data CoreTranslationApiTranslate200ResponseOcsData) *CoreTranslationApiTranslate200ResponseOcs {
	this := CoreTranslationApiTranslate200ResponseOcs{}
	this.Meta = meta
	this.Data = data
	return &this
}

// NewCoreTranslationApiTranslate200ResponseOcsWithDefaults instantiates a new CoreTranslationApiTranslate200ResponseOcs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreTranslationApiTranslate200ResponseOcsWithDefaults() *CoreTranslationApiTranslate200ResponseOcs {
	this := CoreTranslationApiTranslate200ResponseOcs{}
	return &this
}

// GetMeta returns the Meta field value
func (o *CoreTranslationApiTranslate200ResponseOcs) GetMeta() OCSMeta {
	if o == nil {
		var ret OCSMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *CoreTranslationApiTranslate200ResponseOcs) GetMetaOk() (*OCSMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *CoreTranslationApiTranslate200ResponseOcs) SetMeta(v OCSMeta) {
	o.Meta = v
}

// GetData returns the Data field value
func (o *CoreTranslationApiTranslate200ResponseOcs) GetData() CoreTranslationApiTranslate200ResponseOcsData {
	if o == nil {
		var ret CoreTranslationApiTranslate200ResponseOcsData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CoreTranslationApiTranslate200ResponseOcs) GetDataOk() (*CoreTranslationApiTranslate200ResponseOcsData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CoreTranslationApiTranslate200ResponseOcs) SetData(v CoreTranslationApiTranslate200ResponseOcsData) {
	o.Data = v
}

func (o CoreTranslationApiTranslate200ResponseOcs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreTranslationApiTranslate200ResponseOcs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["data"] = o.Data

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CoreTranslationApiTranslate200ResponseOcs) UnmarshalJSON(bytes []byte) (err error) {
	varCoreTranslationApiTranslate200ResponseOcs := _CoreTranslationApiTranslate200ResponseOcs{}

	if err = json.Unmarshal(bytes, &varCoreTranslationApiTranslate200ResponseOcs); err == nil {
		*o = CoreTranslationApiTranslate200ResponseOcs(varCoreTranslationApiTranslate200ResponseOcs)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "meta")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCoreTranslationApiTranslate200ResponseOcs struct {
	value *CoreTranslationApiTranslate200ResponseOcs
	isSet bool
}

func (v NullableCoreTranslationApiTranslate200ResponseOcs) Get() *CoreTranslationApiTranslate200ResponseOcs {
	return v.value
}

func (v *NullableCoreTranslationApiTranslate200ResponseOcs) Set(val *CoreTranslationApiTranslate200ResponseOcs) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreTranslationApiTranslate200ResponseOcs) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreTranslationApiTranslate200ResponseOcs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreTranslationApiTranslate200ResponseOcs(val *CoreTranslationApiTranslate200ResponseOcs) *NullableCoreTranslationApiTranslate200ResponseOcs {
	return &NullableCoreTranslationApiTranslate200ResponseOcs{value: val, isSet: true}
}

func (v NullableCoreTranslationApiTranslate200ResponseOcs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreTranslationApiTranslate200ResponseOcs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

