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

// checks if the FilesTemplateCreate200ResponseOcs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FilesTemplateCreate200ResponseOcs{}

// FilesTemplateCreate200ResponseOcs struct for FilesTemplateCreate200ResponseOcs
type FilesTemplateCreate200ResponseOcs struct {
	Meta OCSMeta `json:"meta"`
	Data FilesTemplateFile `json:"data"`
	AdditionalProperties map[string]interface{}
}

type _FilesTemplateCreate200ResponseOcs FilesTemplateCreate200ResponseOcs

// NewFilesTemplateCreate200ResponseOcs instantiates a new FilesTemplateCreate200ResponseOcs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilesTemplateCreate200ResponseOcs(meta OCSMeta, data FilesTemplateFile) *FilesTemplateCreate200ResponseOcs {
	this := FilesTemplateCreate200ResponseOcs{}
	this.Meta = meta
	this.Data = data
	return &this
}

// NewFilesTemplateCreate200ResponseOcsWithDefaults instantiates a new FilesTemplateCreate200ResponseOcs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilesTemplateCreate200ResponseOcsWithDefaults() *FilesTemplateCreate200ResponseOcs {
	this := FilesTemplateCreate200ResponseOcs{}
	return &this
}

// GetMeta returns the Meta field value
func (o *FilesTemplateCreate200ResponseOcs) GetMeta() OCSMeta {
	if o == nil {
		var ret OCSMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *FilesTemplateCreate200ResponseOcs) GetMetaOk() (*OCSMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *FilesTemplateCreate200ResponseOcs) SetMeta(v OCSMeta) {
	o.Meta = v
}

// GetData returns the Data field value
func (o *FilesTemplateCreate200ResponseOcs) GetData() FilesTemplateFile {
	if o == nil {
		var ret FilesTemplateFile
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *FilesTemplateCreate200ResponseOcs) GetDataOk() (*FilesTemplateFile, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *FilesTemplateCreate200ResponseOcs) SetData(v FilesTemplateFile) {
	o.Data = v
}

func (o FilesTemplateCreate200ResponseOcs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FilesTemplateCreate200ResponseOcs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["data"] = o.Data

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FilesTemplateCreate200ResponseOcs) UnmarshalJSON(bytes []byte) (err error) {
	varFilesTemplateCreate200ResponseOcs := _FilesTemplateCreate200ResponseOcs{}

	if err = json.Unmarshal(bytes, &varFilesTemplateCreate200ResponseOcs); err == nil {
		*o = FilesTemplateCreate200ResponseOcs(varFilesTemplateCreate200ResponseOcs)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "meta")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFilesTemplateCreate200ResponseOcs struct {
	value *FilesTemplateCreate200ResponseOcs
	isSet bool
}

func (v NullableFilesTemplateCreate200ResponseOcs) Get() *FilesTemplateCreate200ResponseOcs {
	return v.value
}

func (v *NullableFilesTemplateCreate200ResponseOcs) Set(val *FilesTemplateCreate200ResponseOcs) {
	v.value = val
	v.isSet = true
}

func (v NullableFilesTemplateCreate200ResponseOcs) IsSet() bool {
	return v.isSet
}

func (v *NullableFilesTemplateCreate200ResponseOcs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilesTemplateCreate200ResponseOcs(val *FilesTemplateCreate200ResponseOcs) *NullableFilesTemplateCreate200ResponseOcs {
	return &NullableFilesTemplateCreate200ResponseOcs{value: val, isSet: true}
}

func (v NullableFilesTemplateCreate200ResponseOcs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilesTemplateCreate200ResponseOcs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

