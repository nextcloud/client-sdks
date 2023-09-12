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

// checks if the FilesSharingShareeCircleAllOfValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FilesSharingShareeCircleAllOfValue{}

// FilesSharingShareeCircleAllOfValue struct for FilesSharingShareeCircleAllOfValue
type FilesSharingShareeCircleAllOfValue struct {
	ShareType int64 `json:"shareType"`
	ShareWith string `json:"shareWith"`
	Circle string `json:"circle"`
	AdditionalProperties map[string]interface{}
}

type _FilesSharingShareeCircleAllOfValue FilesSharingShareeCircleAllOfValue

// NewFilesSharingShareeCircleAllOfValue instantiates a new FilesSharingShareeCircleAllOfValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilesSharingShareeCircleAllOfValue(shareType int64, shareWith string, circle string) *FilesSharingShareeCircleAllOfValue {
	this := FilesSharingShareeCircleAllOfValue{}
	this.ShareType = shareType
	this.ShareWith = shareWith
	this.Circle = circle
	return &this
}

// NewFilesSharingShareeCircleAllOfValueWithDefaults instantiates a new FilesSharingShareeCircleAllOfValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilesSharingShareeCircleAllOfValueWithDefaults() *FilesSharingShareeCircleAllOfValue {
	this := FilesSharingShareeCircleAllOfValue{}
	return &this
}

// GetShareType returns the ShareType field value
func (o *FilesSharingShareeCircleAllOfValue) GetShareType() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ShareType
}

// GetShareTypeOk returns a tuple with the ShareType field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareeCircleAllOfValue) GetShareTypeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShareType, true
}

// SetShareType sets field value
func (o *FilesSharingShareeCircleAllOfValue) SetShareType(v int64) {
	o.ShareType = v
}

// GetShareWith returns the ShareWith field value
func (o *FilesSharingShareeCircleAllOfValue) GetShareWith() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShareWith
}

// GetShareWithOk returns a tuple with the ShareWith field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareeCircleAllOfValue) GetShareWithOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShareWith, true
}

// SetShareWith sets field value
func (o *FilesSharingShareeCircleAllOfValue) SetShareWith(v string) {
	o.ShareWith = v
}

// GetCircle returns the Circle field value
func (o *FilesSharingShareeCircleAllOfValue) GetCircle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Circle
}

// GetCircleOk returns a tuple with the Circle field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareeCircleAllOfValue) GetCircleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Circle, true
}

// SetCircle sets field value
func (o *FilesSharingShareeCircleAllOfValue) SetCircle(v string) {
	o.Circle = v
}

func (o FilesSharingShareeCircleAllOfValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FilesSharingShareeCircleAllOfValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["shareType"] = o.ShareType
	toSerialize["shareWith"] = o.ShareWith
	toSerialize["circle"] = o.Circle

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FilesSharingShareeCircleAllOfValue) UnmarshalJSON(bytes []byte) (err error) {
	varFilesSharingShareeCircleAllOfValue := _FilesSharingShareeCircleAllOfValue{}

	if err = json.Unmarshal(bytes, &varFilesSharingShareeCircleAllOfValue); err == nil {
		*o = FilesSharingShareeCircleAllOfValue(varFilesSharingShareeCircleAllOfValue)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "shareType")
		delete(additionalProperties, "shareWith")
		delete(additionalProperties, "circle")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFilesSharingShareeCircleAllOfValue struct {
	value *FilesSharingShareeCircleAllOfValue
	isSet bool
}

func (v NullableFilesSharingShareeCircleAllOfValue) Get() *FilesSharingShareeCircleAllOfValue {
	return v.value
}

func (v *NullableFilesSharingShareeCircleAllOfValue) Set(val *FilesSharingShareeCircleAllOfValue) {
	v.value = val
	v.isSet = true
}

func (v NullableFilesSharingShareeCircleAllOfValue) IsSet() bool {
	return v.isSet
}

func (v *NullableFilesSharingShareeCircleAllOfValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilesSharingShareeCircleAllOfValue(val *FilesSharingShareeCircleAllOfValue) *NullableFilesSharingShareeCircleAllOfValue {
	return &NullableFilesSharingShareeCircleAllOfValue{value: val, isSet: true}
}

func (v NullableFilesSharingShareeCircleAllOfValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilesSharingShareeCircleAllOfValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


