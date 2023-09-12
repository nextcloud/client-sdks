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

// checks if the FilesRemindersApiGet200ResponseOcs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FilesRemindersApiGet200ResponseOcs{}

// FilesRemindersApiGet200ResponseOcs struct for FilesRemindersApiGet200ResponseOcs
type FilesRemindersApiGet200ResponseOcs struct {
	Meta OCSMeta `json:"meta"`
	Data FilesRemindersApiGet200ResponseOcsData `json:"data"`
	AdditionalProperties map[string]interface{}
}

type _FilesRemindersApiGet200ResponseOcs FilesRemindersApiGet200ResponseOcs

// NewFilesRemindersApiGet200ResponseOcs instantiates a new FilesRemindersApiGet200ResponseOcs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilesRemindersApiGet200ResponseOcs(meta OCSMeta, data FilesRemindersApiGet200ResponseOcsData) *FilesRemindersApiGet200ResponseOcs {
	this := FilesRemindersApiGet200ResponseOcs{}
	this.Meta = meta
	this.Data = data
	return &this
}

// NewFilesRemindersApiGet200ResponseOcsWithDefaults instantiates a new FilesRemindersApiGet200ResponseOcs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilesRemindersApiGet200ResponseOcsWithDefaults() *FilesRemindersApiGet200ResponseOcs {
	this := FilesRemindersApiGet200ResponseOcs{}
	return &this
}

// GetMeta returns the Meta field value
func (o *FilesRemindersApiGet200ResponseOcs) GetMeta() OCSMeta {
	if o == nil {
		var ret OCSMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *FilesRemindersApiGet200ResponseOcs) GetMetaOk() (*OCSMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *FilesRemindersApiGet200ResponseOcs) SetMeta(v OCSMeta) {
	o.Meta = v
}

// GetData returns the Data field value
func (o *FilesRemindersApiGet200ResponseOcs) GetData() FilesRemindersApiGet200ResponseOcsData {
	if o == nil {
		var ret FilesRemindersApiGet200ResponseOcsData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *FilesRemindersApiGet200ResponseOcs) GetDataOk() (*FilesRemindersApiGet200ResponseOcsData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *FilesRemindersApiGet200ResponseOcs) SetData(v FilesRemindersApiGet200ResponseOcsData) {
	o.Data = v
}

func (o FilesRemindersApiGet200ResponseOcs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FilesRemindersApiGet200ResponseOcs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["data"] = o.Data

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FilesRemindersApiGet200ResponseOcs) UnmarshalJSON(bytes []byte) (err error) {
	varFilesRemindersApiGet200ResponseOcs := _FilesRemindersApiGet200ResponseOcs{}

	if err = json.Unmarshal(bytes, &varFilesRemindersApiGet200ResponseOcs); err == nil {
		*o = FilesRemindersApiGet200ResponseOcs(varFilesRemindersApiGet200ResponseOcs)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "meta")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFilesRemindersApiGet200ResponseOcs struct {
	value *FilesRemindersApiGet200ResponseOcs
	isSet bool
}

func (v NullableFilesRemindersApiGet200ResponseOcs) Get() *FilesRemindersApiGet200ResponseOcs {
	return v.value
}

func (v *NullableFilesRemindersApiGet200ResponseOcs) Set(val *FilesRemindersApiGet200ResponseOcs) {
	v.value = val
	v.isSet = true
}

func (v NullableFilesRemindersApiGet200ResponseOcs) IsSet() bool {
	return v.isSet
}

func (v *NullableFilesRemindersApiGet200ResponseOcs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilesRemindersApiGet200ResponseOcs(val *FilesRemindersApiGet200ResponseOcs) *NullableFilesRemindersApiGet200ResponseOcs {
	return &NullableFilesRemindersApiGet200ResponseOcs{value: val, isSet: true}
}

func (v NullableFilesRemindersApiGet200ResponseOcs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilesRemindersApiGet200ResponseOcs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


