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

// checks if the FilesSharingShareeRemote type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FilesSharingShareeRemote{}

// FilesSharingShareeRemote struct for FilesSharingShareeRemote
type FilesSharingShareeRemote struct {
	Count NullableInt64 `json:"count"`
	Label string `json:"label"`
	Uuid string `json:"uuid"`
	Name string `json:"name"`
	Type string `json:"type"`
	Value FilesSharingShareeRemoteAllOfValue `json:"value"`
	AdditionalProperties map[string]interface{}
}

type _FilesSharingShareeRemote FilesSharingShareeRemote

// NewFilesSharingShareeRemote instantiates a new FilesSharingShareeRemote object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilesSharingShareeRemote(count NullableInt64, label string, uuid string, name string, type_ string, value FilesSharingShareeRemoteAllOfValue) *FilesSharingShareeRemote {
	this := FilesSharingShareeRemote{}
	this.Count = count
	this.Label = label
	this.Uuid = uuid
	this.Name = name
	this.Type = type_
	this.Value = value
	return &this
}

// NewFilesSharingShareeRemoteWithDefaults instantiates a new FilesSharingShareeRemote object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilesSharingShareeRemoteWithDefaults() *FilesSharingShareeRemote {
	this := FilesSharingShareeRemote{}
	return &this
}

// GetCount returns the Count field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *FilesSharingShareeRemote) GetCount() int64 {
	if o == nil || o.Count.Get() == nil {
		var ret int64
		return ret
	}

	return *o.Count.Get()
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesSharingShareeRemote) GetCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Count.Get(), o.Count.IsSet()
}

// SetCount sets field value
func (o *FilesSharingShareeRemote) SetCount(v int64) {
	o.Count.Set(&v)
}

// GetLabel returns the Label field value
func (o *FilesSharingShareeRemote) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareeRemote) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *FilesSharingShareeRemote) SetLabel(v string) {
	o.Label = v
}

// GetUuid returns the Uuid field value
func (o *FilesSharingShareeRemote) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareeRemote) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *FilesSharingShareeRemote) SetUuid(v string) {
	o.Uuid = v
}

// GetName returns the Name field value
func (o *FilesSharingShareeRemote) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareeRemote) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FilesSharingShareeRemote) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *FilesSharingShareeRemote) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareeRemote) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FilesSharingShareeRemote) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *FilesSharingShareeRemote) GetValue() FilesSharingShareeRemoteAllOfValue {
	if o == nil {
		var ret FilesSharingShareeRemoteAllOfValue
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareeRemote) GetValueOk() (*FilesSharingShareeRemoteAllOfValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *FilesSharingShareeRemote) SetValue(v FilesSharingShareeRemoteAllOfValue) {
	o.Value = v
}

func (o FilesSharingShareeRemote) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FilesSharingShareeRemote) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["count"] = o.Count.Get()
	toSerialize["label"] = o.Label
	toSerialize["uuid"] = o.Uuid
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FilesSharingShareeRemote) UnmarshalJSON(bytes []byte) (err error) {
	varFilesSharingShareeRemote := _FilesSharingShareeRemote{}

	if err = json.Unmarshal(bytes, &varFilesSharingShareeRemote); err == nil {
		*o = FilesSharingShareeRemote(varFilesSharingShareeRemote)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "count")
		delete(additionalProperties, "label")
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFilesSharingShareeRemote struct {
	value *FilesSharingShareeRemote
	isSet bool
}

func (v NullableFilesSharingShareeRemote) Get() *FilesSharingShareeRemote {
	return v.value
}

func (v *NullableFilesSharingShareeRemote) Set(val *FilesSharingShareeRemote) {
	v.value = val
	v.isSet = true
}

func (v NullableFilesSharingShareeRemote) IsSet() bool {
	return v.isSet
}

func (v *NullableFilesSharingShareeRemote) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilesSharingShareeRemote(val *FilesSharingShareeRemote) *NullableFilesSharingShareeRemote {
	return &NullableFilesSharingShareeRemote{value: val, isSet: true}
}

func (v NullableFilesSharingShareeRemote) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilesSharingShareeRemote) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


