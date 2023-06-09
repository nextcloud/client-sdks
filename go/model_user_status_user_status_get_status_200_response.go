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

// checks if the UserStatusUserStatusGetStatus200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserStatusUserStatusGetStatus200Response{}

// UserStatusUserStatusGetStatus200Response struct for UserStatusUserStatusGetStatus200Response
type UserStatusUserStatusGetStatus200Response struct {
	Ocs UserStatusUserStatusGetStatus200ResponseOcs `json:"ocs"`
}

// NewUserStatusUserStatusGetStatus200Response instantiates a new UserStatusUserStatusGetStatus200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserStatusUserStatusGetStatus200Response(ocs UserStatusUserStatusGetStatus200ResponseOcs) *UserStatusUserStatusGetStatus200Response {
	this := UserStatusUserStatusGetStatus200Response{}
	this.Ocs = ocs
	return &this
}

// NewUserStatusUserStatusGetStatus200ResponseWithDefaults instantiates a new UserStatusUserStatusGetStatus200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserStatusUserStatusGetStatus200ResponseWithDefaults() *UserStatusUserStatusGetStatus200Response {
	this := UserStatusUserStatusGetStatus200Response{}
	return &this
}

// GetOcs returns the Ocs field value
func (o *UserStatusUserStatusGetStatus200Response) GetOcs() UserStatusUserStatusGetStatus200ResponseOcs {
	if o == nil {
		var ret UserStatusUserStatusGetStatus200ResponseOcs
		return ret
	}

	return o.Ocs
}

// GetOcsOk returns a tuple with the Ocs field value
// and a boolean to check if the value has been set.
func (o *UserStatusUserStatusGetStatus200Response) GetOcsOk() (*UserStatusUserStatusGetStatus200ResponseOcs, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ocs, true
}

// SetOcs sets field value
func (o *UserStatusUserStatusGetStatus200Response) SetOcs(v UserStatusUserStatusGetStatus200ResponseOcs) {
	o.Ocs = v
}

func (o UserStatusUserStatusGetStatus200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserStatusUserStatusGetStatus200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ocs"] = o.Ocs
	return toSerialize, nil
}

type NullableUserStatusUserStatusGetStatus200Response struct {
	value *UserStatusUserStatusGetStatus200Response
	isSet bool
}

func (v NullableUserStatusUserStatusGetStatus200Response) Get() *UserStatusUserStatusGetStatus200Response {
	return v.value
}

func (v *NullableUserStatusUserStatusGetStatus200Response) Set(val *UserStatusUserStatusGetStatus200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableUserStatusUserStatusGetStatus200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableUserStatusUserStatusGetStatus200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserStatusUserStatusGetStatus200Response(val *UserStatusUserStatusGetStatus200Response) *NullableUserStatusUserStatusGetStatus200Response {
	return &NullableUserStatusUserStatusGetStatus200Response{value: val, isSet: true}
}

func (v NullableUserStatusUserStatusGetStatus200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserStatusUserStatusGetStatus200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


