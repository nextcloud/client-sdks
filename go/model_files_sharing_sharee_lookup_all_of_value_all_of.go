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

// checks if the FilesSharingShareeLookupAllOfValueAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FilesSharingShareeLookupAllOfValueAllOf{}

// FilesSharingShareeLookupAllOfValueAllOf struct for FilesSharingShareeLookupAllOfValueAllOf
type FilesSharingShareeLookupAllOfValueAllOf struct {
	GlobalScale bool `json:"globalScale"`
}

// NewFilesSharingShareeLookupAllOfValueAllOf instantiates a new FilesSharingShareeLookupAllOfValueAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilesSharingShareeLookupAllOfValueAllOf(globalScale bool) *FilesSharingShareeLookupAllOfValueAllOf {
	this := FilesSharingShareeLookupAllOfValueAllOf{}
	this.GlobalScale = globalScale
	return &this
}

// NewFilesSharingShareeLookupAllOfValueAllOfWithDefaults instantiates a new FilesSharingShareeLookupAllOfValueAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilesSharingShareeLookupAllOfValueAllOfWithDefaults() *FilesSharingShareeLookupAllOfValueAllOf {
	this := FilesSharingShareeLookupAllOfValueAllOf{}
	return &this
}

// GetGlobalScale returns the GlobalScale field value
func (o *FilesSharingShareeLookupAllOfValueAllOf) GetGlobalScale() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.GlobalScale
}

// GetGlobalScaleOk returns a tuple with the GlobalScale field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareeLookupAllOfValueAllOf) GetGlobalScaleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GlobalScale, true
}

// SetGlobalScale sets field value
func (o *FilesSharingShareeLookupAllOfValueAllOf) SetGlobalScale(v bool) {
	o.GlobalScale = v
}

func (o FilesSharingShareeLookupAllOfValueAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FilesSharingShareeLookupAllOfValueAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["globalScale"] = o.GlobalScale
	return toSerialize, nil
}

type NullableFilesSharingShareeLookupAllOfValueAllOf struct {
	value *FilesSharingShareeLookupAllOfValueAllOf
	isSet bool
}

func (v NullableFilesSharingShareeLookupAllOfValueAllOf) Get() *FilesSharingShareeLookupAllOfValueAllOf {
	return v.value
}

func (v *NullableFilesSharingShareeLookupAllOfValueAllOf) Set(val *FilesSharingShareeLookupAllOfValueAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFilesSharingShareeLookupAllOfValueAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFilesSharingShareeLookupAllOfValueAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilesSharingShareeLookupAllOfValueAllOf(val *FilesSharingShareeLookupAllOfValueAllOf) *NullableFilesSharingShareeLookupAllOfValueAllOf {
	return &NullableFilesSharingShareeLookupAllOfValueAllOf{value: val, isSet: true}
}

func (v NullableFilesSharingShareeLookupAllOfValueAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilesSharingShareeLookupAllOfValueAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

