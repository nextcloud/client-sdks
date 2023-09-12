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

// checks if the FilesSharingShareeLookupAllOfValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FilesSharingShareeLookupAllOfValue{}

// FilesSharingShareeLookupAllOfValue struct for FilesSharingShareeLookupAllOfValue
type FilesSharingShareeLookupAllOfValue struct {
	ShareType int64 `json:"shareType"`
	ShareWith string `json:"shareWith"`
	GlobalScale bool `json:"globalScale"`
	AdditionalProperties map[string]interface{}
}

type _FilesSharingShareeLookupAllOfValue FilesSharingShareeLookupAllOfValue

// NewFilesSharingShareeLookupAllOfValue instantiates a new FilesSharingShareeLookupAllOfValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilesSharingShareeLookupAllOfValue(shareType int64, shareWith string, globalScale bool) *FilesSharingShareeLookupAllOfValue {
	this := FilesSharingShareeLookupAllOfValue{}
	this.ShareType = shareType
	this.ShareWith = shareWith
	this.GlobalScale = globalScale
	return &this
}

// NewFilesSharingShareeLookupAllOfValueWithDefaults instantiates a new FilesSharingShareeLookupAllOfValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilesSharingShareeLookupAllOfValueWithDefaults() *FilesSharingShareeLookupAllOfValue {
	this := FilesSharingShareeLookupAllOfValue{}
	return &this
}

// GetShareType returns the ShareType field value
func (o *FilesSharingShareeLookupAllOfValue) GetShareType() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ShareType
}

// GetShareTypeOk returns a tuple with the ShareType field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareeLookupAllOfValue) GetShareTypeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShareType, true
}

// SetShareType sets field value
func (o *FilesSharingShareeLookupAllOfValue) SetShareType(v int64) {
	o.ShareType = v
}

// GetShareWith returns the ShareWith field value
func (o *FilesSharingShareeLookupAllOfValue) GetShareWith() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShareWith
}

// GetShareWithOk returns a tuple with the ShareWith field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareeLookupAllOfValue) GetShareWithOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShareWith, true
}

// SetShareWith sets field value
func (o *FilesSharingShareeLookupAllOfValue) SetShareWith(v string) {
	o.ShareWith = v
}

// GetGlobalScale returns the GlobalScale field value
func (o *FilesSharingShareeLookupAllOfValue) GetGlobalScale() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.GlobalScale
}

// GetGlobalScaleOk returns a tuple with the GlobalScale field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareeLookupAllOfValue) GetGlobalScaleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GlobalScale, true
}

// SetGlobalScale sets field value
func (o *FilesSharingShareeLookupAllOfValue) SetGlobalScale(v bool) {
	o.GlobalScale = v
}

func (o FilesSharingShareeLookupAllOfValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FilesSharingShareeLookupAllOfValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["shareType"] = o.ShareType
	toSerialize["shareWith"] = o.ShareWith
	toSerialize["globalScale"] = o.GlobalScale

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FilesSharingShareeLookupAllOfValue) UnmarshalJSON(bytes []byte) (err error) {
	varFilesSharingShareeLookupAllOfValue := _FilesSharingShareeLookupAllOfValue{}

	if err = json.Unmarshal(bytes, &varFilesSharingShareeLookupAllOfValue); err == nil {
		*o = FilesSharingShareeLookupAllOfValue(varFilesSharingShareeLookupAllOfValue)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "shareType")
		delete(additionalProperties, "shareWith")
		delete(additionalProperties, "globalScale")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFilesSharingShareeLookupAllOfValue struct {
	value *FilesSharingShareeLookupAllOfValue
	isSet bool
}

func (v NullableFilesSharingShareeLookupAllOfValue) Get() *FilesSharingShareeLookupAllOfValue {
	return v.value
}

func (v *NullableFilesSharingShareeLookupAllOfValue) Set(val *FilesSharingShareeLookupAllOfValue) {
	v.value = val
	v.isSet = true
}

func (v NullableFilesSharingShareeLookupAllOfValue) IsSet() bool {
	return v.isSet
}

func (v *NullableFilesSharingShareeLookupAllOfValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilesSharingShareeLookupAllOfValue(val *FilesSharingShareeLookupAllOfValue) *NullableFilesSharingShareeLookupAllOfValue {
	return &NullableFilesSharingShareeLookupAllOfValue{value: val, isSet: true}
}

func (v NullableFilesSharingShareeLookupAllOfValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilesSharingShareeLookupAllOfValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

