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

// checks if the FilesSharingShareStatusOneOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FilesSharingShareStatusOneOf{}

// FilesSharingShareStatusOneOf struct for FilesSharingShareStatusOneOf
type FilesSharingShareStatusOneOf struct {
	Status string `json:"status"`
	Message NullableString `json:"message"`
	Icon NullableString `json:"icon"`
	ClearAt NullableInt64 `json:"clearAt"`
}

// NewFilesSharingShareStatusOneOf instantiates a new FilesSharingShareStatusOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilesSharingShareStatusOneOf(status string, message NullableString, icon NullableString, clearAt NullableInt64) *FilesSharingShareStatusOneOf {
	this := FilesSharingShareStatusOneOf{}
	this.Status = status
	this.Message = message
	this.Icon = icon
	this.ClearAt = clearAt
	return &this
}

// NewFilesSharingShareStatusOneOfWithDefaults instantiates a new FilesSharingShareStatusOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilesSharingShareStatusOneOfWithDefaults() *FilesSharingShareStatusOneOf {
	this := FilesSharingShareStatusOneOf{}
	return &this
}

// GetStatus returns the Status field value
func (o *FilesSharingShareStatusOneOf) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareStatusOneOf) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *FilesSharingShareStatusOneOf) SetStatus(v string) {
	o.Status = v
}

// GetMessage returns the Message field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FilesSharingShareStatusOneOf) GetMessage() string {
	if o == nil || o.Message.Get() == nil {
		var ret string
		return ret
	}

	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesSharingShareStatusOneOf) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// SetMessage sets field value
func (o *FilesSharingShareStatusOneOf) SetMessage(v string) {
	o.Message.Set(&v)
}

// GetIcon returns the Icon field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FilesSharingShareStatusOneOf) GetIcon() string {
	if o == nil || o.Icon.Get() == nil {
		var ret string
		return ret
	}

	return *o.Icon.Get()
}

// GetIconOk returns a tuple with the Icon field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesSharingShareStatusOneOf) GetIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Icon.Get(), o.Icon.IsSet()
}

// SetIcon sets field value
func (o *FilesSharingShareStatusOneOf) SetIcon(v string) {
	o.Icon.Set(&v)
}

// GetClearAt returns the ClearAt field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *FilesSharingShareStatusOneOf) GetClearAt() int64 {
	if o == nil || o.ClearAt.Get() == nil {
		var ret int64
		return ret
	}

	return *o.ClearAt.Get()
}

// GetClearAtOk returns a tuple with the ClearAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesSharingShareStatusOneOf) GetClearAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClearAt.Get(), o.ClearAt.IsSet()
}

// SetClearAt sets field value
func (o *FilesSharingShareStatusOneOf) SetClearAt(v int64) {
	o.ClearAt.Set(&v)
}

func (o FilesSharingShareStatusOneOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FilesSharingShareStatusOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	toSerialize["message"] = o.Message.Get()
	toSerialize["icon"] = o.Icon.Get()
	toSerialize["clearAt"] = o.ClearAt.Get()
	return toSerialize, nil
}

type NullableFilesSharingShareStatusOneOf struct {
	value *FilesSharingShareStatusOneOf
	isSet bool
}

func (v NullableFilesSharingShareStatusOneOf) Get() *FilesSharingShareStatusOneOf {
	return v.value
}

func (v *NullableFilesSharingShareStatusOneOf) Set(val *FilesSharingShareStatusOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFilesSharingShareStatusOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFilesSharingShareStatusOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilesSharingShareStatusOneOf(val *FilesSharingShareStatusOneOf) *NullableFilesSharingShareStatusOneOf {
	return &NullableFilesSharingShareStatusOneOf{value: val, isSet: true}
}

func (v NullableFilesSharingShareStatusOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilesSharingShareStatusOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


