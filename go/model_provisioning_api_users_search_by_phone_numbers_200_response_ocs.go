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

// checks if the ProvisioningApiUsersSearchByPhoneNumbers200ResponseOcs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProvisioningApiUsersSearchByPhoneNumbers200ResponseOcs{}

// ProvisioningApiUsersSearchByPhoneNumbers200ResponseOcs struct for ProvisioningApiUsersSearchByPhoneNumbers200ResponseOcs
type ProvisioningApiUsersSearchByPhoneNumbers200ResponseOcs struct {
	Meta OCSMeta `json:"meta"`
	Data map[string]string `json:"data"`
}

// NewProvisioningApiUsersSearchByPhoneNumbers200ResponseOcs instantiates a new ProvisioningApiUsersSearchByPhoneNumbers200ResponseOcs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisioningApiUsersSearchByPhoneNumbers200ResponseOcs(meta OCSMeta, data map[string]string) *ProvisioningApiUsersSearchByPhoneNumbers200ResponseOcs {
	this := ProvisioningApiUsersSearchByPhoneNumbers200ResponseOcs{}
	this.Meta = meta
	this.Data = data
	return &this
}

// NewProvisioningApiUsersSearchByPhoneNumbers200ResponseOcsWithDefaults instantiates a new ProvisioningApiUsersSearchByPhoneNumbers200ResponseOcs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisioningApiUsersSearchByPhoneNumbers200ResponseOcsWithDefaults() *ProvisioningApiUsersSearchByPhoneNumbers200ResponseOcs {
	this := ProvisioningApiUsersSearchByPhoneNumbers200ResponseOcs{}
	return &this
}

// GetMeta returns the Meta field value
func (o *ProvisioningApiUsersSearchByPhoneNumbers200ResponseOcs) GetMeta() OCSMeta {
	if o == nil {
		var ret OCSMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *ProvisioningApiUsersSearchByPhoneNumbers200ResponseOcs) GetMetaOk() (*OCSMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *ProvisioningApiUsersSearchByPhoneNumbers200ResponseOcs) SetMeta(v OCSMeta) {
	o.Meta = v
}

// GetData returns the Data field value
func (o *ProvisioningApiUsersSearchByPhoneNumbers200ResponseOcs) GetData() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ProvisioningApiUsersSearchByPhoneNumbers200ResponseOcs) GetDataOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ProvisioningApiUsersSearchByPhoneNumbers200ResponseOcs) SetData(v map[string]string) {
	o.Data = v
}

func (o ProvisioningApiUsersSearchByPhoneNumbers200ResponseOcs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProvisioningApiUsersSearchByPhoneNumbers200ResponseOcs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableProvisioningApiUsersSearchByPhoneNumbers200ResponseOcs struct {
	value *ProvisioningApiUsersSearchByPhoneNumbers200ResponseOcs
	isSet bool
}

func (v NullableProvisioningApiUsersSearchByPhoneNumbers200ResponseOcs) Get() *ProvisioningApiUsersSearchByPhoneNumbers200ResponseOcs {
	return v.value
}

func (v *NullableProvisioningApiUsersSearchByPhoneNumbers200ResponseOcs) Set(val *ProvisioningApiUsersSearchByPhoneNumbers200ResponseOcs) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisioningApiUsersSearchByPhoneNumbers200ResponseOcs) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisioningApiUsersSearchByPhoneNumbers200ResponseOcs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisioningApiUsersSearchByPhoneNumbers200ResponseOcs(val *ProvisioningApiUsersSearchByPhoneNumbers200ResponseOcs) *NullableProvisioningApiUsersSearchByPhoneNumbers200ResponseOcs {
	return &NullableProvisioningApiUsersSearchByPhoneNumbers200ResponseOcs{value: val, isSet: true}
}

func (v NullableProvisioningApiUsersSearchByPhoneNumbers200ResponseOcs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisioningApiUsersSearchByPhoneNumbers200ResponseOcs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


