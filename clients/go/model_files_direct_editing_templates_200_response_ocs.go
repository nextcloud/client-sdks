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

// checks if the FilesDirectEditingTemplates200ResponseOcs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FilesDirectEditingTemplates200ResponseOcs{}

// FilesDirectEditingTemplates200ResponseOcs struct for FilesDirectEditingTemplates200ResponseOcs
type FilesDirectEditingTemplates200ResponseOcs struct {
	Meta OCSMeta `json:"meta"`
	Data FilesDirectEditingTemplates200ResponseOcsData `json:"data"`
	AdditionalProperties map[string]interface{}
}

type _FilesDirectEditingTemplates200ResponseOcs FilesDirectEditingTemplates200ResponseOcs

// NewFilesDirectEditingTemplates200ResponseOcs instantiates a new FilesDirectEditingTemplates200ResponseOcs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilesDirectEditingTemplates200ResponseOcs(meta OCSMeta, data FilesDirectEditingTemplates200ResponseOcsData) *FilesDirectEditingTemplates200ResponseOcs {
	this := FilesDirectEditingTemplates200ResponseOcs{}
	this.Meta = meta
	this.Data = data
	return &this
}

// NewFilesDirectEditingTemplates200ResponseOcsWithDefaults instantiates a new FilesDirectEditingTemplates200ResponseOcs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilesDirectEditingTemplates200ResponseOcsWithDefaults() *FilesDirectEditingTemplates200ResponseOcs {
	this := FilesDirectEditingTemplates200ResponseOcs{}
	return &this
}

// GetMeta returns the Meta field value
func (o *FilesDirectEditingTemplates200ResponseOcs) GetMeta() OCSMeta {
	if o == nil {
		var ret OCSMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *FilesDirectEditingTemplates200ResponseOcs) GetMetaOk() (*OCSMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *FilesDirectEditingTemplates200ResponseOcs) SetMeta(v OCSMeta) {
	o.Meta = v
}

// GetData returns the Data field value
func (o *FilesDirectEditingTemplates200ResponseOcs) GetData() FilesDirectEditingTemplates200ResponseOcsData {
	if o == nil {
		var ret FilesDirectEditingTemplates200ResponseOcsData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *FilesDirectEditingTemplates200ResponseOcs) GetDataOk() (*FilesDirectEditingTemplates200ResponseOcsData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *FilesDirectEditingTemplates200ResponseOcs) SetData(v FilesDirectEditingTemplates200ResponseOcsData) {
	o.Data = v
}

func (o FilesDirectEditingTemplates200ResponseOcs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FilesDirectEditingTemplates200ResponseOcs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["data"] = o.Data

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FilesDirectEditingTemplates200ResponseOcs) UnmarshalJSON(bytes []byte) (err error) {
	varFilesDirectEditingTemplates200ResponseOcs := _FilesDirectEditingTemplates200ResponseOcs{}

	if err = json.Unmarshal(bytes, &varFilesDirectEditingTemplates200ResponseOcs); err == nil {
		*o = FilesDirectEditingTemplates200ResponseOcs(varFilesDirectEditingTemplates200ResponseOcs)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "meta")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFilesDirectEditingTemplates200ResponseOcs struct {
	value *FilesDirectEditingTemplates200ResponseOcs
	isSet bool
}

func (v NullableFilesDirectEditingTemplates200ResponseOcs) Get() *FilesDirectEditingTemplates200ResponseOcs {
	return v.value
}

func (v *NullableFilesDirectEditingTemplates200ResponseOcs) Set(val *FilesDirectEditingTemplates200ResponseOcs) {
	v.value = val
	v.isSet = true
}

func (v NullableFilesDirectEditingTemplates200ResponseOcs) IsSet() bool {
	return v.isSet
}

func (v *NullableFilesDirectEditingTemplates200ResponseOcs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilesDirectEditingTemplates200ResponseOcs(val *FilesDirectEditingTemplates200ResponseOcs) *NullableFilesDirectEditingTemplates200ResponseOcs {
	return &NullableFilesDirectEditingTemplates200ResponseOcs{value: val, isSet: true}
}

func (v NullableFilesDirectEditingTemplates200ResponseOcs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilesDirectEditingTemplates200ResponseOcs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


