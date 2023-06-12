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

// checks if the ProvisioningApiGroupsGetGroups200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProvisioningApiGroupsGetGroups200Response{}

// ProvisioningApiGroupsGetGroups200Response struct for ProvisioningApiGroupsGetGroups200Response
type ProvisioningApiGroupsGetGroups200Response struct {
	Ocs ProvisioningApiGroupsGetGroups200ResponseOcs `json:"ocs"`
}

// NewProvisioningApiGroupsGetGroups200Response instantiates a new ProvisioningApiGroupsGetGroups200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisioningApiGroupsGetGroups200Response(ocs ProvisioningApiGroupsGetGroups200ResponseOcs) *ProvisioningApiGroupsGetGroups200Response {
	this := ProvisioningApiGroupsGetGroups200Response{}
	this.Ocs = ocs
	return &this
}

// NewProvisioningApiGroupsGetGroups200ResponseWithDefaults instantiates a new ProvisioningApiGroupsGetGroups200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisioningApiGroupsGetGroups200ResponseWithDefaults() *ProvisioningApiGroupsGetGroups200Response {
	this := ProvisioningApiGroupsGetGroups200Response{}
	return &this
}

// GetOcs returns the Ocs field value
func (o *ProvisioningApiGroupsGetGroups200Response) GetOcs() ProvisioningApiGroupsGetGroups200ResponseOcs {
	if o == nil {
		var ret ProvisioningApiGroupsGetGroups200ResponseOcs
		return ret
	}

	return o.Ocs
}

// GetOcsOk returns a tuple with the Ocs field value
// and a boolean to check if the value has been set.
func (o *ProvisioningApiGroupsGetGroups200Response) GetOcsOk() (*ProvisioningApiGroupsGetGroups200ResponseOcs, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ocs, true
}

// SetOcs sets field value
func (o *ProvisioningApiGroupsGetGroups200Response) SetOcs(v ProvisioningApiGroupsGetGroups200ResponseOcs) {
	o.Ocs = v
}

func (o ProvisioningApiGroupsGetGroups200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProvisioningApiGroupsGetGroups200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ocs"] = o.Ocs
	return toSerialize, nil
}

type NullableProvisioningApiGroupsGetGroups200Response struct {
	value *ProvisioningApiGroupsGetGroups200Response
	isSet bool
}

func (v NullableProvisioningApiGroupsGetGroups200Response) Get() *ProvisioningApiGroupsGetGroups200Response {
	return v.value
}

func (v *NullableProvisioningApiGroupsGetGroups200Response) Set(val *ProvisioningApiGroupsGetGroups200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisioningApiGroupsGetGroups200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisioningApiGroupsGetGroups200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisioningApiGroupsGetGroups200Response(val *ProvisioningApiGroupsGetGroups200Response) *NullableProvisioningApiGroupsGetGroups200Response {
	return &NullableProvisioningApiGroupsGetGroups200Response{value: val, isSet: true}
}

func (v NullableProvisioningApiGroupsGetGroups200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisioningApiGroupsGetGroups200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


