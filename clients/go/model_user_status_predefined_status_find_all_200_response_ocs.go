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

// checks if the UserStatusPredefinedStatusFindAll200ResponseOcs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserStatusPredefinedStatusFindAll200ResponseOcs{}

// UserStatusPredefinedStatusFindAll200ResponseOcs struct for UserStatusPredefinedStatusFindAll200ResponseOcs
type UserStatusPredefinedStatusFindAll200ResponseOcs struct {
	Meta OCSMeta `json:"meta"`
	Data []UserStatusPredefined `json:"data"`
	AdditionalProperties map[string]interface{}
}

type _UserStatusPredefinedStatusFindAll200ResponseOcs UserStatusPredefinedStatusFindAll200ResponseOcs

// NewUserStatusPredefinedStatusFindAll200ResponseOcs instantiates a new UserStatusPredefinedStatusFindAll200ResponseOcs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserStatusPredefinedStatusFindAll200ResponseOcs(meta OCSMeta, data []UserStatusPredefined) *UserStatusPredefinedStatusFindAll200ResponseOcs {
	this := UserStatusPredefinedStatusFindAll200ResponseOcs{}
	this.Meta = meta
	this.Data = data
	return &this
}

// NewUserStatusPredefinedStatusFindAll200ResponseOcsWithDefaults instantiates a new UserStatusPredefinedStatusFindAll200ResponseOcs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserStatusPredefinedStatusFindAll200ResponseOcsWithDefaults() *UserStatusPredefinedStatusFindAll200ResponseOcs {
	this := UserStatusPredefinedStatusFindAll200ResponseOcs{}
	return &this
}

// GetMeta returns the Meta field value
func (o *UserStatusPredefinedStatusFindAll200ResponseOcs) GetMeta() OCSMeta {
	if o == nil {
		var ret OCSMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *UserStatusPredefinedStatusFindAll200ResponseOcs) GetMetaOk() (*OCSMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *UserStatusPredefinedStatusFindAll200ResponseOcs) SetMeta(v OCSMeta) {
	o.Meta = v
}

// GetData returns the Data field value
func (o *UserStatusPredefinedStatusFindAll200ResponseOcs) GetData() []UserStatusPredefined {
	if o == nil {
		var ret []UserStatusPredefined
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *UserStatusPredefinedStatusFindAll200ResponseOcs) GetDataOk() ([]UserStatusPredefined, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *UserStatusPredefinedStatusFindAll200ResponseOcs) SetData(v []UserStatusPredefined) {
	o.Data = v
}

func (o UserStatusPredefinedStatusFindAll200ResponseOcs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserStatusPredefinedStatusFindAll200ResponseOcs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["data"] = o.Data

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserStatusPredefinedStatusFindAll200ResponseOcs) UnmarshalJSON(bytes []byte) (err error) {
	varUserStatusPredefinedStatusFindAll200ResponseOcs := _UserStatusPredefinedStatusFindAll200ResponseOcs{}

	if err = json.Unmarshal(bytes, &varUserStatusPredefinedStatusFindAll200ResponseOcs); err == nil {
		*o = UserStatusPredefinedStatusFindAll200ResponseOcs(varUserStatusPredefinedStatusFindAll200ResponseOcs)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "meta")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserStatusPredefinedStatusFindAll200ResponseOcs struct {
	value *UserStatusPredefinedStatusFindAll200ResponseOcs
	isSet bool
}

func (v NullableUserStatusPredefinedStatusFindAll200ResponseOcs) Get() *UserStatusPredefinedStatusFindAll200ResponseOcs {
	return v.value
}

func (v *NullableUserStatusPredefinedStatusFindAll200ResponseOcs) Set(val *UserStatusPredefinedStatusFindAll200ResponseOcs) {
	v.value = val
	v.isSet = true
}

func (v NullableUserStatusPredefinedStatusFindAll200ResponseOcs) IsSet() bool {
	return v.isSet
}

func (v *NullableUserStatusPredefinedStatusFindAll200ResponseOcs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserStatusPredefinedStatusFindAll200ResponseOcs(val *UserStatusPredefinedStatusFindAll200ResponseOcs) *NullableUserStatusPredefinedStatusFindAll200ResponseOcs {
	return &NullableUserStatusPredefinedStatusFindAll200ResponseOcs{value: val, isSet: true}
}

func (v NullableUserStatusPredefinedStatusFindAll200ResponseOcs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserStatusPredefinedStatusFindAll200ResponseOcs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


