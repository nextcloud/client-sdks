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

// checks if the UserStatusPredefinedStatusFindAll200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserStatusPredefinedStatusFindAll200Response{}

// UserStatusPredefinedStatusFindAll200Response struct for UserStatusPredefinedStatusFindAll200Response
type UserStatusPredefinedStatusFindAll200Response struct {
	Ocs UserStatusPredefinedStatusFindAll200ResponseOcs `json:"ocs"`
}

// NewUserStatusPredefinedStatusFindAll200Response instantiates a new UserStatusPredefinedStatusFindAll200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserStatusPredefinedStatusFindAll200Response(ocs UserStatusPredefinedStatusFindAll200ResponseOcs) *UserStatusPredefinedStatusFindAll200Response {
	this := UserStatusPredefinedStatusFindAll200Response{}
	this.Ocs = ocs
	return &this
}

// NewUserStatusPredefinedStatusFindAll200ResponseWithDefaults instantiates a new UserStatusPredefinedStatusFindAll200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserStatusPredefinedStatusFindAll200ResponseWithDefaults() *UserStatusPredefinedStatusFindAll200Response {
	this := UserStatusPredefinedStatusFindAll200Response{}
	return &this
}

// GetOcs returns the Ocs field value
func (o *UserStatusPredefinedStatusFindAll200Response) GetOcs() UserStatusPredefinedStatusFindAll200ResponseOcs {
	if o == nil {
		var ret UserStatusPredefinedStatusFindAll200ResponseOcs
		return ret
	}

	return o.Ocs
}

// GetOcsOk returns a tuple with the Ocs field value
// and a boolean to check if the value has been set.
func (o *UserStatusPredefinedStatusFindAll200Response) GetOcsOk() (*UserStatusPredefinedStatusFindAll200ResponseOcs, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ocs, true
}

// SetOcs sets field value
func (o *UserStatusPredefinedStatusFindAll200Response) SetOcs(v UserStatusPredefinedStatusFindAll200ResponseOcs) {
	o.Ocs = v
}

func (o UserStatusPredefinedStatusFindAll200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserStatusPredefinedStatusFindAll200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ocs"] = o.Ocs
	return toSerialize, nil
}

type NullableUserStatusPredefinedStatusFindAll200Response struct {
	value *UserStatusPredefinedStatusFindAll200Response
	isSet bool
}

func (v NullableUserStatusPredefinedStatusFindAll200Response) Get() *UserStatusPredefinedStatusFindAll200Response {
	return v.value
}

func (v *NullableUserStatusPredefinedStatusFindAll200Response) Set(val *UserStatusPredefinedStatusFindAll200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableUserStatusPredefinedStatusFindAll200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableUserStatusPredefinedStatusFindAll200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserStatusPredefinedStatusFindAll200Response(val *UserStatusPredefinedStatusFindAll200Response) *NullableUserStatusPredefinedStatusFindAll200Response {
	return &NullableUserStatusPredefinedStatusFindAll200Response{value: val, isSet: true}
}

func (v NullableUserStatusPredefinedStatusFindAll200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserStatusPredefinedStatusFindAll200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


