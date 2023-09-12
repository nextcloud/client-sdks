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

// checks if the FilesOpenLocalEditorCreate200ResponseOcs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FilesOpenLocalEditorCreate200ResponseOcs{}

// FilesOpenLocalEditorCreate200ResponseOcs struct for FilesOpenLocalEditorCreate200ResponseOcs
type FilesOpenLocalEditorCreate200ResponseOcs struct {
	Meta OCSMeta `json:"meta"`
	Data FilesOpenLocalEditorCreate200ResponseOcsData `json:"data"`
	AdditionalProperties map[string]interface{}
}

type _FilesOpenLocalEditorCreate200ResponseOcs FilesOpenLocalEditorCreate200ResponseOcs

// NewFilesOpenLocalEditorCreate200ResponseOcs instantiates a new FilesOpenLocalEditorCreate200ResponseOcs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilesOpenLocalEditorCreate200ResponseOcs(meta OCSMeta, data FilesOpenLocalEditorCreate200ResponseOcsData) *FilesOpenLocalEditorCreate200ResponseOcs {
	this := FilesOpenLocalEditorCreate200ResponseOcs{}
	this.Meta = meta
	this.Data = data
	return &this
}

// NewFilesOpenLocalEditorCreate200ResponseOcsWithDefaults instantiates a new FilesOpenLocalEditorCreate200ResponseOcs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilesOpenLocalEditorCreate200ResponseOcsWithDefaults() *FilesOpenLocalEditorCreate200ResponseOcs {
	this := FilesOpenLocalEditorCreate200ResponseOcs{}
	return &this
}

// GetMeta returns the Meta field value
func (o *FilesOpenLocalEditorCreate200ResponseOcs) GetMeta() OCSMeta {
	if o == nil {
		var ret OCSMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *FilesOpenLocalEditorCreate200ResponseOcs) GetMetaOk() (*OCSMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *FilesOpenLocalEditorCreate200ResponseOcs) SetMeta(v OCSMeta) {
	o.Meta = v
}

// GetData returns the Data field value
func (o *FilesOpenLocalEditorCreate200ResponseOcs) GetData() FilesOpenLocalEditorCreate200ResponseOcsData {
	if o == nil {
		var ret FilesOpenLocalEditorCreate200ResponseOcsData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *FilesOpenLocalEditorCreate200ResponseOcs) GetDataOk() (*FilesOpenLocalEditorCreate200ResponseOcsData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *FilesOpenLocalEditorCreate200ResponseOcs) SetData(v FilesOpenLocalEditorCreate200ResponseOcsData) {
	o.Data = v
}

func (o FilesOpenLocalEditorCreate200ResponseOcs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FilesOpenLocalEditorCreate200ResponseOcs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["data"] = o.Data

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FilesOpenLocalEditorCreate200ResponseOcs) UnmarshalJSON(bytes []byte) (err error) {
	varFilesOpenLocalEditorCreate200ResponseOcs := _FilesOpenLocalEditorCreate200ResponseOcs{}

	if err = json.Unmarshal(bytes, &varFilesOpenLocalEditorCreate200ResponseOcs); err == nil {
		*o = FilesOpenLocalEditorCreate200ResponseOcs(varFilesOpenLocalEditorCreate200ResponseOcs)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "meta")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFilesOpenLocalEditorCreate200ResponseOcs struct {
	value *FilesOpenLocalEditorCreate200ResponseOcs
	isSet bool
}

func (v NullableFilesOpenLocalEditorCreate200ResponseOcs) Get() *FilesOpenLocalEditorCreate200ResponseOcs {
	return v.value
}

func (v *NullableFilesOpenLocalEditorCreate200ResponseOcs) Set(val *FilesOpenLocalEditorCreate200ResponseOcs) {
	v.value = val
	v.isSet = true
}

func (v NullableFilesOpenLocalEditorCreate200ResponseOcs) IsSet() bool {
	return v.isSet
}

func (v *NullableFilesOpenLocalEditorCreate200ResponseOcs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilesOpenLocalEditorCreate200ResponseOcs(val *FilesOpenLocalEditorCreate200ResponseOcs) *NullableFilesOpenLocalEditorCreate200ResponseOcs {
	return &NullableFilesOpenLocalEditorCreate200ResponseOcs{value: val, isSet: true}
}

func (v NullableFilesOpenLocalEditorCreate200ResponseOcs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilesOpenLocalEditorCreate200ResponseOcs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

