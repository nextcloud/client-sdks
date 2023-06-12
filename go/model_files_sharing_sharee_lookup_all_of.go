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

// checks if the FilesSharingShareeLookupAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FilesSharingShareeLookupAllOf{}

// FilesSharingShareeLookupAllOf struct for FilesSharingShareeLookupAllOf
type FilesSharingShareeLookupAllOf struct {
	Extra FilesSharingShareeLookupAllOfExtra `json:"extra"`
	Value FilesSharingShareeLookupAllOfValue `json:"value"`
}

// NewFilesSharingShareeLookupAllOf instantiates a new FilesSharingShareeLookupAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilesSharingShareeLookupAllOf(extra FilesSharingShareeLookupAllOfExtra, value FilesSharingShareeLookupAllOfValue) *FilesSharingShareeLookupAllOf {
	this := FilesSharingShareeLookupAllOf{}
	this.Extra = extra
	this.Value = value
	return &this
}

// NewFilesSharingShareeLookupAllOfWithDefaults instantiates a new FilesSharingShareeLookupAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilesSharingShareeLookupAllOfWithDefaults() *FilesSharingShareeLookupAllOf {
	this := FilesSharingShareeLookupAllOf{}
	return &this
}

// GetExtra returns the Extra field value
func (o *FilesSharingShareeLookupAllOf) GetExtra() FilesSharingShareeLookupAllOfExtra {
	if o == nil {
		var ret FilesSharingShareeLookupAllOfExtra
		return ret
	}

	return o.Extra
}

// GetExtraOk returns a tuple with the Extra field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareeLookupAllOf) GetExtraOk() (*FilesSharingShareeLookupAllOfExtra, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Extra, true
}

// SetExtra sets field value
func (o *FilesSharingShareeLookupAllOf) SetExtra(v FilesSharingShareeLookupAllOfExtra) {
	o.Extra = v
}

// GetValue returns the Value field value
func (o *FilesSharingShareeLookupAllOf) GetValue() FilesSharingShareeLookupAllOfValue {
	if o == nil {
		var ret FilesSharingShareeLookupAllOfValue
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareeLookupAllOf) GetValueOk() (*FilesSharingShareeLookupAllOfValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *FilesSharingShareeLookupAllOf) SetValue(v FilesSharingShareeLookupAllOfValue) {
	o.Value = v
}

func (o FilesSharingShareeLookupAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FilesSharingShareeLookupAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["extra"] = o.Extra
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableFilesSharingShareeLookupAllOf struct {
	value *FilesSharingShareeLookupAllOf
	isSet bool
}

func (v NullableFilesSharingShareeLookupAllOf) Get() *FilesSharingShareeLookupAllOf {
	return v.value
}

func (v *NullableFilesSharingShareeLookupAllOf) Set(val *FilesSharingShareeLookupAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFilesSharingShareeLookupAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFilesSharingShareeLookupAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilesSharingShareeLookupAllOf(val *FilesSharingShareeLookupAllOf) *NullableFilesSharingShareeLookupAllOf {
	return &NullableFilesSharingShareeLookupAllOf{value: val, isSet: true}
}

func (v NullableFilesSharingShareeLookupAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilesSharingShareeLookupAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


