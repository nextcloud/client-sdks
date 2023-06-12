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

// checks if the ProvisioningApiGroupsGetGroupUsers200ResponseOcsData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProvisioningApiGroupsGetGroupUsers200ResponseOcsData{}

// ProvisioningApiGroupsGetGroupUsers200ResponseOcsData struct for ProvisioningApiGroupsGetGroupUsers200ResponseOcsData
type ProvisioningApiGroupsGetGroupUsers200ResponseOcsData struct {
	Users []string `json:"users"`
}

// NewProvisioningApiGroupsGetGroupUsers200ResponseOcsData instantiates a new ProvisioningApiGroupsGetGroupUsers200ResponseOcsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisioningApiGroupsGetGroupUsers200ResponseOcsData(users []string) *ProvisioningApiGroupsGetGroupUsers200ResponseOcsData {
	this := ProvisioningApiGroupsGetGroupUsers200ResponseOcsData{}
	this.Users = users
	return &this
}

// NewProvisioningApiGroupsGetGroupUsers200ResponseOcsDataWithDefaults instantiates a new ProvisioningApiGroupsGetGroupUsers200ResponseOcsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisioningApiGroupsGetGroupUsers200ResponseOcsDataWithDefaults() *ProvisioningApiGroupsGetGroupUsers200ResponseOcsData {
	this := ProvisioningApiGroupsGetGroupUsers200ResponseOcsData{}
	return &this
}

// GetUsers returns the Users field value
func (o *ProvisioningApiGroupsGetGroupUsers200ResponseOcsData) GetUsers() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *ProvisioningApiGroupsGetGroupUsers200ResponseOcsData) GetUsersOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Users, true
}

// SetUsers sets field value
func (o *ProvisioningApiGroupsGetGroupUsers200ResponseOcsData) SetUsers(v []string) {
	o.Users = v
}

func (o ProvisioningApiGroupsGetGroupUsers200ResponseOcsData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProvisioningApiGroupsGetGroupUsers200ResponseOcsData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["users"] = o.Users
	return toSerialize, nil
}

type NullableProvisioningApiGroupsGetGroupUsers200ResponseOcsData struct {
	value *ProvisioningApiGroupsGetGroupUsers200ResponseOcsData
	isSet bool
}

func (v NullableProvisioningApiGroupsGetGroupUsers200ResponseOcsData) Get() *ProvisioningApiGroupsGetGroupUsers200ResponseOcsData {
	return v.value
}

func (v *NullableProvisioningApiGroupsGetGroupUsers200ResponseOcsData) Set(val *ProvisioningApiGroupsGetGroupUsers200ResponseOcsData) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisioningApiGroupsGetGroupUsers200ResponseOcsData) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisioningApiGroupsGetGroupUsers200ResponseOcsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisioningApiGroupsGetGroupUsers200ResponseOcsData(val *ProvisioningApiGroupsGetGroupUsers200ResponseOcsData) *NullableProvisioningApiGroupsGetGroupUsers200ResponseOcsData {
	return &NullableProvisioningApiGroupsGetGroupUsers200ResponseOcsData{value: val, isSet: true}
}

func (v NullableProvisioningApiGroupsGetGroupUsers200ResponseOcsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisioningApiGroupsGetGroupUsers200ResponseOcsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


