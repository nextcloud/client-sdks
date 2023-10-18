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

// checks if the FilesSharingShareeEmailAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FilesSharingShareeEmailAllOf{}

// FilesSharingShareeEmailAllOf struct for FilesSharingShareeEmailAllOf
type FilesSharingShareeEmailAllOf struct {
	Uuid string `json:"uuid"`
	Name string `json:"name"`
	Type string `json:"type"`
	ShareWithDisplayNameUnique string `json:"shareWithDisplayNameUnique"`
	Value FilesSharingShareeValue `json:"value"`
}

// NewFilesSharingShareeEmailAllOf instantiates a new FilesSharingShareeEmailAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilesSharingShareeEmailAllOf(uuid string, name string, type_ string, shareWithDisplayNameUnique string, value FilesSharingShareeValue) *FilesSharingShareeEmailAllOf {
	this := FilesSharingShareeEmailAllOf{}
	this.Uuid = uuid
	this.Name = name
	this.Type = type_
	this.ShareWithDisplayNameUnique = shareWithDisplayNameUnique
	this.Value = value
	return &this
}

// NewFilesSharingShareeEmailAllOfWithDefaults instantiates a new FilesSharingShareeEmailAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilesSharingShareeEmailAllOfWithDefaults() *FilesSharingShareeEmailAllOf {
	this := FilesSharingShareeEmailAllOf{}
	return &this
}

// GetUuid returns the Uuid field value
func (o *FilesSharingShareeEmailAllOf) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareeEmailAllOf) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *FilesSharingShareeEmailAllOf) SetUuid(v string) {
	o.Uuid = v
}

// GetName returns the Name field value
func (o *FilesSharingShareeEmailAllOf) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareeEmailAllOf) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FilesSharingShareeEmailAllOf) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *FilesSharingShareeEmailAllOf) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareeEmailAllOf) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FilesSharingShareeEmailAllOf) SetType(v string) {
	o.Type = v
}

// GetShareWithDisplayNameUnique returns the ShareWithDisplayNameUnique field value
func (o *FilesSharingShareeEmailAllOf) GetShareWithDisplayNameUnique() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShareWithDisplayNameUnique
}

// GetShareWithDisplayNameUniqueOk returns a tuple with the ShareWithDisplayNameUnique field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareeEmailAllOf) GetShareWithDisplayNameUniqueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShareWithDisplayNameUnique, true
}

// SetShareWithDisplayNameUnique sets field value
func (o *FilesSharingShareeEmailAllOf) SetShareWithDisplayNameUnique(v string) {
	o.ShareWithDisplayNameUnique = v
}

// GetValue returns the Value field value
func (o *FilesSharingShareeEmailAllOf) GetValue() FilesSharingShareeValue {
	if o == nil {
		var ret FilesSharingShareeValue
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareeEmailAllOf) GetValueOk() (*FilesSharingShareeValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *FilesSharingShareeEmailAllOf) SetValue(v FilesSharingShareeValue) {
	o.Value = v
}

func (o FilesSharingShareeEmailAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FilesSharingShareeEmailAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["uuid"] = o.Uuid
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	toSerialize["shareWithDisplayNameUnique"] = o.ShareWithDisplayNameUnique
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableFilesSharingShareeEmailAllOf struct {
	value *FilesSharingShareeEmailAllOf
	isSet bool
}

func (v NullableFilesSharingShareeEmailAllOf) Get() *FilesSharingShareeEmailAllOf {
	return v.value
}

func (v *NullableFilesSharingShareeEmailAllOf) Set(val *FilesSharingShareeEmailAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFilesSharingShareeEmailAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFilesSharingShareeEmailAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilesSharingShareeEmailAllOf(val *FilesSharingShareeEmailAllOf) *NullableFilesSharingShareeEmailAllOf {
	return &NullableFilesSharingShareeEmailAllOf{value: val, isSet: true}
}

func (v NullableFilesSharingShareeEmailAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilesSharingShareeEmailAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

