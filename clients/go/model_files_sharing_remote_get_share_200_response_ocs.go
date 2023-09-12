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

// checks if the FilesSharingRemoteGetShare200ResponseOcs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FilesSharingRemoteGetShare200ResponseOcs{}

// FilesSharingRemoteGetShare200ResponseOcs struct for FilesSharingRemoteGetShare200ResponseOcs
type FilesSharingRemoteGetShare200ResponseOcs struct {
	Meta OCSMeta `json:"meta"`
	Data FilesSharingRemoteShare `json:"data"`
	AdditionalProperties map[string]interface{}
}

type _FilesSharingRemoteGetShare200ResponseOcs FilesSharingRemoteGetShare200ResponseOcs

// NewFilesSharingRemoteGetShare200ResponseOcs instantiates a new FilesSharingRemoteGetShare200ResponseOcs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilesSharingRemoteGetShare200ResponseOcs(meta OCSMeta, data FilesSharingRemoteShare) *FilesSharingRemoteGetShare200ResponseOcs {
	this := FilesSharingRemoteGetShare200ResponseOcs{}
	this.Meta = meta
	this.Data = data
	return &this
}

// NewFilesSharingRemoteGetShare200ResponseOcsWithDefaults instantiates a new FilesSharingRemoteGetShare200ResponseOcs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilesSharingRemoteGetShare200ResponseOcsWithDefaults() *FilesSharingRemoteGetShare200ResponseOcs {
	this := FilesSharingRemoteGetShare200ResponseOcs{}
	return &this
}

// GetMeta returns the Meta field value
func (o *FilesSharingRemoteGetShare200ResponseOcs) GetMeta() OCSMeta {
	if o == nil {
		var ret OCSMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *FilesSharingRemoteGetShare200ResponseOcs) GetMetaOk() (*OCSMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *FilesSharingRemoteGetShare200ResponseOcs) SetMeta(v OCSMeta) {
	o.Meta = v
}

// GetData returns the Data field value
func (o *FilesSharingRemoteGetShare200ResponseOcs) GetData() FilesSharingRemoteShare {
	if o == nil {
		var ret FilesSharingRemoteShare
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *FilesSharingRemoteGetShare200ResponseOcs) GetDataOk() (*FilesSharingRemoteShare, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *FilesSharingRemoteGetShare200ResponseOcs) SetData(v FilesSharingRemoteShare) {
	o.Data = v
}

func (o FilesSharingRemoteGetShare200ResponseOcs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FilesSharingRemoteGetShare200ResponseOcs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["data"] = o.Data

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FilesSharingRemoteGetShare200ResponseOcs) UnmarshalJSON(bytes []byte) (err error) {
	varFilesSharingRemoteGetShare200ResponseOcs := _FilesSharingRemoteGetShare200ResponseOcs{}

	if err = json.Unmarshal(bytes, &varFilesSharingRemoteGetShare200ResponseOcs); err == nil {
		*o = FilesSharingRemoteGetShare200ResponseOcs(varFilesSharingRemoteGetShare200ResponseOcs)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "meta")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFilesSharingRemoteGetShare200ResponseOcs struct {
	value *FilesSharingRemoteGetShare200ResponseOcs
	isSet bool
}

func (v NullableFilesSharingRemoteGetShare200ResponseOcs) Get() *FilesSharingRemoteGetShare200ResponseOcs {
	return v.value
}

func (v *NullableFilesSharingRemoteGetShare200ResponseOcs) Set(val *FilesSharingRemoteGetShare200ResponseOcs) {
	v.value = val
	v.isSet = true
}

func (v NullableFilesSharingRemoteGetShare200ResponseOcs) IsSet() bool {
	return v.isSet
}

func (v *NullableFilesSharingRemoteGetShare200ResponseOcs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilesSharingRemoteGetShare200ResponseOcs(val *FilesSharingRemoteGetShare200ResponseOcs) *NullableFilesSharingRemoteGetShare200ResponseOcs {
	return &NullableFilesSharingRemoteGetShare200ResponseOcs{value: val, isSet: true}
}

func (v NullableFilesSharingRemoteGetShare200ResponseOcs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilesSharingRemoteGetShare200ResponseOcs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

