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

// checks if the ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcs{}

// ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcs struct for ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcs
type ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcs struct {
	Meta OCSMeta `json:"meta"`
	Data ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsData `json:"data"`
}

// NewProvisioningApiGroupsGetGroupUsersDetails200ResponseOcs instantiates a new ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisioningApiGroupsGetGroupUsersDetails200ResponseOcs(meta OCSMeta, data ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsData) *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcs {
	this := ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcs{}
	this.Meta = meta
	this.Data = data
	return &this
}

// NewProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsWithDefaults instantiates a new ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsWithDefaults() *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcs {
	this := ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcs{}
	return &this
}

// GetMeta returns the Meta field value
func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcs) GetMeta() OCSMeta {
	if o == nil {
		var ret OCSMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcs) GetMetaOk() (*OCSMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcs) SetMeta(v OCSMeta) {
	o.Meta = v
}

// GetData returns the Data field value
func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcs) GetData() ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsData {
	if o == nil {
		var ret ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcs) GetDataOk() (*ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcs) SetData(v ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsData) {
	o.Data = v
}

func (o ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableProvisioningApiGroupsGetGroupUsersDetails200ResponseOcs struct {
	value *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcs
	isSet bool
}

func (v NullableProvisioningApiGroupsGetGroupUsersDetails200ResponseOcs) Get() *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcs {
	return v.value
}

func (v *NullableProvisioningApiGroupsGetGroupUsersDetails200ResponseOcs) Set(val *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcs) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisioningApiGroupsGetGroupUsersDetails200ResponseOcs) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisioningApiGroupsGetGroupUsersDetails200ResponseOcs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisioningApiGroupsGetGroupUsersDetails200ResponseOcs(val *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcs) *NullableProvisioningApiGroupsGetGroupUsersDetails200ResponseOcs {
	return &NullableProvisioningApiGroupsGetGroupUsersDetails200ResponseOcs{value: val, isSet: true}
}

func (v NullableProvisioningApiGroupsGetGroupUsersDetails200ResponseOcs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisioningApiGroupsGetGroupUsersDetails200ResponseOcs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

