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

// checks if the UserStatusUserStatusRevertStatus200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserStatusUserStatusRevertStatus200Response{}

// UserStatusUserStatusRevertStatus200Response struct for UserStatusUserStatusRevertStatus200Response
type UserStatusUserStatusRevertStatus200Response struct {
	Ocs UserStatusUserStatusRevertStatus200ResponseOcs `json:"ocs"`
	AdditionalProperties map[string]interface{}
}

type _UserStatusUserStatusRevertStatus200Response UserStatusUserStatusRevertStatus200Response

// NewUserStatusUserStatusRevertStatus200Response instantiates a new UserStatusUserStatusRevertStatus200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserStatusUserStatusRevertStatus200Response(ocs UserStatusUserStatusRevertStatus200ResponseOcs) *UserStatusUserStatusRevertStatus200Response {
	this := UserStatusUserStatusRevertStatus200Response{}
	this.Ocs = ocs
	return &this
}

// NewUserStatusUserStatusRevertStatus200ResponseWithDefaults instantiates a new UserStatusUserStatusRevertStatus200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserStatusUserStatusRevertStatus200ResponseWithDefaults() *UserStatusUserStatusRevertStatus200Response {
	this := UserStatusUserStatusRevertStatus200Response{}
	return &this
}

// GetOcs returns the Ocs field value
func (o *UserStatusUserStatusRevertStatus200Response) GetOcs() UserStatusUserStatusRevertStatus200ResponseOcs {
	if o == nil {
		var ret UserStatusUserStatusRevertStatus200ResponseOcs
		return ret
	}

	return o.Ocs
}

// GetOcsOk returns a tuple with the Ocs field value
// and a boolean to check if the value has been set.
func (o *UserStatusUserStatusRevertStatus200Response) GetOcsOk() (*UserStatusUserStatusRevertStatus200ResponseOcs, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ocs, true
}

// SetOcs sets field value
func (o *UserStatusUserStatusRevertStatus200Response) SetOcs(v UserStatusUserStatusRevertStatus200ResponseOcs) {
	o.Ocs = v
}

func (o UserStatusUserStatusRevertStatus200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserStatusUserStatusRevertStatus200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ocs"] = o.Ocs

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserStatusUserStatusRevertStatus200Response) UnmarshalJSON(bytes []byte) (err error) {
	varUserStatusUserStatusRevertStatus200Response := _UserStatusUserStatusRevertStatus200Response{}

	if err = json.Unmarshal(bytes, &varUserStatusUserStatusRevertStatus200Response); err == nil {
		*o = UserStatusUserStatusRevertStatus200Response(varUserStatusUserStatusRevertStatus200Response)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ocs")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserStatusUserStatusRevertStatus200Response struct {
	value *UserStatusUserStatusRevertStatus200Response
	isSet bool
}

func (v NullableUserStatusUserStatusRevertStatus200Response) Get() *UserStatusUserStatusRevertStatus200Response {
	return v.value
}

func (v *NullableUserStatusUserStatusRevertStatus200Response) Set(val *UserStatusUserStatusRevertStatus200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableUserStatusUserStatusRevertStatus200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableUserStatusUserStatusRevertStatus200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserStatusUserStatusRevertStatus200Response(val *UserStatusUserStatusRevertStatus200Response) *NullableUserStatusUserStatusRevertStatus200Response {
	return &NullableUserStatusUserStatusRevertStatus200Response{value: val, isSet: true}
}

func (v NullableUserStatusUserStatusRevertStatus200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserStatusUserStatusRevertStatus200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

