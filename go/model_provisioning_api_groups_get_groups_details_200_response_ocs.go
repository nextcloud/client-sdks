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

// checks if the ProvisioningApiGroupsGetGroupsDetails200ResponseOcs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProvisioningApiGroupsGetGroupsDetails200ResponseOcs{}

// ProvisioningApiGroupsGetGroupsDetails200ResponseOcs struct for ProvisioningApiGroupsGetGroupsDetails200ResponseOcs
type ProvisioningApiGroupsGetGroupsDetails200ResponseOcs struct {
	Meta OCSMeta `json:"meta"`
	Data ProvisioningApiGroupsGetGroupsDetails200ResponseOcsData `json:"data"`
}

// NewProvisioningApiGroupsGetGroupsDetails200ResponseOcs instantiates a new ProvisioningApiGroupsGetGroupsDetails200ResponseOcs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisioningApiGroupsGetGroupsDetails200ResponseOcs(meta OCSMeta, data ProvisioningApiGroupsGetGroupsDetails200ResponseOcsData) *ProvisioningApiGroupsGetGroupsDetails200ResponseOcs {
	this := ProvisioningApiGroupsGetGroupsDetails200ResponseOcs{}
	this.Meta = meta
	this.Data = data
	return &this
}

// NewProvisioningApiGroupsGetGroupsDetails200ResponseOcsWithDefaults instantiates a new ProvisioningApiGroupsGetGroupsDetails200ResponseOcs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisioningApiGroupsGetGroupsDetails200ResponseOcsWithDefaults() *ProvisioningApiGroupsGetGroupsDetails200ResponseOcs {
	this := ProvisioningApiGroupsGetGroupsDetails200ResponseOcs{}
	return &this
}

// GetMeta returns the Meta field value
func (o *ProvisioningApiGroupsGetGroupsDetails200ResponseOcs) GetMeta() OCSMeta {
	if o == nil {
		var ret OCSMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *ProvisioningApiGroupsGetGroupsDetails200ResponseOcs) GetMetaOk() (*OCSMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *ProvisioningApiGroupsGetGroupsDetails200ResponseOcs) SetMeta(v OCSMeta) {
	o.Meta = v
}

// GetData returns the Data field value
func (o *ProvisioningApiGroupsGetGroupsDetails200ResponseOcs) GetData() ProvisioningApiGroupsGetGroupsDetails200ResponseOcsData {
	if o == nil {
		var ret ProvisioningApiGroupsGetGroupsDetails200ResponseOcsData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ProvisioningApiGroupsGetGroupsDetails200ResponseOcs) GetDataOk() (*ProvisioningApiGroupsGetGroupsDetails200ResponseOcsData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ProvisioningApiGroupsGetGroupsDetails200ResponseOcs) SetData(v ProvisioningApiGroupsGetGroupsDetails200ResponseOcsData) {
	o.Data = v
}

func (o ProvisioningApiGroupsGetGroupsDetails200ResponseOcs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProvisioningApiGroupsGetGroupsDetails200ResponseOcs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableProvisioningApiGroupsGetGroupsDetails200ResponseOcs struct {
	value *ProvisioningApiGroupsGetGroupsDetails200ResponseOcs
	isSet bool
}

func (v NullableProvisioningApiGroupsGetGroupsDetails200ResponseOcs) Get() *ProvisioningApiGroupsGetGroupsDetails200ResponseOcs {
	return v.value
}

func (v *NullableProvisioningApiGroupsGetGroupsDetails200ResponseOcs) Set(val *ProvisioningApiGroupsGetGroupsDetails200ResponseOcs) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisioningApiGroupsGetGroupsDetails200ResponseOcs) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisioningApiGroupsGetGroupsDetails200ResponseOcs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisioningApiGroupsGetGroupsDetails200ResponseOcs(val *ProvisioningApiGroupsGetGroupsDetails200ResponseOcs) *NullableProvisioningApiGroupsGetGroupsDetails200ResponseOcs {
	return &NullableProvisioningApiGroupsGetGroupsDetails200ResponseOcs{value: val, isSet: true}
}

func (v NullableProvisioningApiGroupsGetGroupsDetails200ResponseOcs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisioningApiGroupsGetGroupsDetails200ResponseOcs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

