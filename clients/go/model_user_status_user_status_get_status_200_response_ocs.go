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

// checks if the UserStatusUserStatusGetStatus200ResponseOcs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserStatusUserStatusGetStatus200ResponseOcs{}

// UserStatusUserStatusGetStatus200ResponseOcs struct for UserStatusUserStatusGetStatus200ResponseOcs
type UserStatusUserStatusGetStatus200ResponseOcs struct {
	Meta OCSMeta `json:"meta"`
	Data UserStatusPrivate `json:"data"`
	AdditionalProperties map[string]interface{}
}

type _UserStatusUserStatusGetStatus200ResponseOcs UserStatusUserStatusGetStatus200ResponseOcs

// NewUserStatusUserStatusGetStatus200ResponseOcs instantiates a new UserStatusUserStatusGetStatus200ResponseOcs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserStatusUserStatusGetStatus200ResponseOcs(meta OCSMeta, data UserStatusPrivate) *UserStatusUserStatusGetStatus200ResponseOcs {
	this := UserStatusUserStatusGetStatus200ResponseOcs{}
	this.Meta = meta
	this.Data = data
	return &this
}

// NewUserStatusUserStatusGetStatus200ResponseOcsWithDefaults instantiates a new UserStatusUserStatusGetStatus200ResponseOcs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserStatusUserStatusGetStatus200ResponseOcsWithDefaults() *UserStatusUserStatusGetStatus200ResponseOcs {
	this := UserStatusUserStatusGetStatus200ResponseOcs{}
	return &this
}

// GetMeta returns the Meta field value
func (o *UserStatusUserStatusGetStatus200ResponseOcs) GetMeta() OCSMeta {
	if o == nil {
		var ret OCSMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *UserStatusUserStatusGetStatus200ResponseOcs) GetMetaOk() (*OCSMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *UserStatusUserStatusGetStatus200ResponseOcs) SetMeta(v OCSMeta) {
	o.Meta = v
}

// GetData returns the Data field value
func (o *UserStatusUserStatusGetStatus200ResponseOcs) GetData() UserStatusPrivate {
	if o == nil {
		var ret UserStatusPrivate
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *UserStatusUserStatusGetStatus200ResponseOcs) GetDataOk() (*UserStatusPrivate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *UserStatusUserStatusGetStatus200ResponseOcs) SetData(v UserStatusPrivate) {
	o.Data = v
}

func (o UserStatusUserStatusGetStatus200ResponseOcs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserStatusUserStatusGetStatus200ResponseOcs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["data"] = o.Data

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserStatusUserStatusGetStatus200ResponseOcs) UnmarshalJSON(bytes []byte) (err error) {
	varUserStatusUserStatusGetStatus200ResponseOcs := _UserStatusUserStatusGetStatus200ResponseOcs{}

	if err = json.Unmarshal(bytes, &varUserStatusUserStatusGetStatus200ResponseOcs); err == nil {
		*o = UserStatusUserStatusGetStatus200ResponseOcs(varUserStatusUserStatusGetStatus200ResponseOcs)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "meta")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserStatusUserStatusGetStatus200ResponseOcs struct {
	value *UserStatusUserStatusGetStatus200ResponseOcs
	isSet bool
}

func (v NullableUserStatusUserStatusGetStatus200ResponseOcs) Get() *UserStatusUserStatusGetStatus200ResponseOcs {
	return v.value
}

func (v *NullableUserStatusUserStatusGetStatus200ResponseOcs) Set(val *UserStatusUserStatusGetStatus200ResponseOcs) {
	v.value = val
	v.isSet = true
}

func (v NullableUserStatusUserStatusGetStatus200ResponseOcs) IsSet() bool {
	return v.isSet
}

func (v *NullableUserStatusUserStatusGetStatus200ResponseOcs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserStatusUserStatusGetStatus200ResponseOcs(val *UserStatusUserStatusGetStatus200ResponseOcs) *NullableUserStatusUserStatusGetStatus200ResponseOcs {
	return &NullableUserStatusUserStatusGetStatus200ResponseOcs{value: val, isSet: true}
}

func (v NullableUserStatusUserStatusGetStatus200ResponseOcs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserStatusUserStatusGetStatus200ResponseOcs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


