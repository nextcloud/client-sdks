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

// checks if the CoreTranslationApiLanguages200ResponseOcs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreTranslationApiLanguages200ResponseOcs{}

// CoreTranslationApiLanguages200ResponseOcs struct for CoreTranslationApiLanguages200ResponseOcs
type CoreTranslationApiLanguages200ResponseOcs struct {
	Meta OCSMeta `json:"meta"`
	Data CoreTranslationApiLanguages200ResponseOcsData `json:"data"`
}

// NewCoreTranslationApiLanguages200ResponseOcs instantiates a new CoreTranslationApiLanguages200ResponseOcs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreTranslationApiLanguages200ResponseOcs(meta OCSMeta, data CoreTranslationApiLanguages200ResponseOcsData) *CoreTranslationApiLanguages200ResponseOcs {
	this := CoreTranslationApiLanguages200ResponseOcs{}
	this.Meta = meta
	this.Data = data
	return &this
}

// NewCoreTranslationApiLanguages200ResponseOcsWithDefaults instantiates a new CoreTranslationApiLanguages200ResponseOcs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreTranslationApiLanguages200ResponseOcsWithDefaults() *CoreTranslationApiLanguages200ResponseOcs {
	this := CoreTranslationApiLanguages200ResponseOcs{}
	return &this
}

// GetMeta returns the Meta field value
func (o *CoreTranslationApiLanguages200ResponseOcs) GetMeta() OCSMeta {
	if o == nil {
		var ret OCSMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *CoreTranslationApiLanguages200ResponseOcs) GetMetaOk() (*OCSMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *CoreTranslationApiLanguages200ResponseOcs) SetMeta(v OCSMeta) {
	o.Meta = v
}

// GetData returns the Data field value
func (o *CoreTranslationApiLanguages200ResponseOcs) GetData() CoreTranslationApiLanguages200ResponseOcsData {
	if o == nil {
		var ret CoreTranslationApiLanguages200ResponseOcsData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CoreTranslationApiLanguages200ResponseOcs) GetDataOk() (*CoreTranslationApiLanguages200ResponseOcsData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CoreTranslationApiLanguages200ResponseOcs) SetData(v CoreTranslationApiLanguages200ResponseOcsData) {
	o.Data = v
}

func (o CoreTranslationApiLanguages200ResponseOcs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreTranslationApiLanguages200ResponseOcs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableCoreTranslationApiLanguages200ResponseOcs struct {
	value *CoreTranslationApiLanguages200ResponseOcs
	isSet bool
}

func (v NullableCoreTranslationApiLanguages200ResponseOcs) Get() *CoreTranslationApiLanguages200ResponseOcs {
	return v.value
}

func (v *NullableCoreTranslationApiLanguages200ResponseOcs) Set(val *CoreTranslationApiLanguages200ResponseOcs) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreTranslationApiLanguages200ResponseOcs) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreTranslationApiLanguages200ResponseOcs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreTranslationApiLanguages200ResponseOcs(val *CoreTranslationApiLanguages200ResponseOcs) *NullableCoreTranslationApiLanguages200ResponseOcs {
	return &NullableCoreTranslationApiLanguages200ResponseOcs{value: val, isSet: true}
}

func (v NullableCoreTranslationApiLanguages200ResponseOcs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreTranslationApiLanguages200ResponseOcs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


