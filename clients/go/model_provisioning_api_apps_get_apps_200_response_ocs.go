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

// checks if the ProvisioningApiAppsGetApps200ResponseOcs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProvisioningApiAppsGetApps200ResponseOcs{}

// ProvisioningApiAppsGetApps200ResponseOcs struct for ProvisioningApiAppsGetApps200ResponseOcs
type ProvisioningApiAppsGetApps200ResponseOcs struct {
	Meta OCSMeta `json:"meta"`
	Data ProvisioningApiAppsGetApps200ResponseOcsData `json:"data"`
	AdditionalProperties map[string]interface{}
}

type _ProvisioningApiAppsGetApps200ResponseOcs ProvisioningApiAppsGetApps200ResponseOcs

// NewProvisioningApiAppsGetApps200ResponseOcs instantiates a new ProvisioningApiAppsGetApps200ResponseOcs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisioningApiAppsGetApps200ResponseOcs(meta OCSMeta, data ProvisioningApiAppsGetApps200ResponseOcsData) *ProvisioningApiAppsGetApps200ResponseOcs {
	this := ProvisioningApiAppsGetApps200ResponseOcs{}
	this.Meta = meta
	this.Data = data
	return &this
}

// NewProvisioningApiAppsGetApps200ResponseOcsWithDefaults instantiates a new ProvisioningApiAppsGetApps200ResponseOcs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisioningApiAppsGetApps200ResponseOcsWithDefaults() *ProvisioningApiAppsGetApps200ResponseOcs {
	this := ProvisioningApiAppsGetApps200ResponseOcs{}
	return &this
}

// GetMeta returns the Meta field value
func (o *ProvisioningApiAppsGetApps200ResponseOcs) GetMeta() OCSMeta {
	if o == nil {
		var ret OCSMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *ProvisioningApiAppsGetApps200ResponseOcs) GetMetaOk() (*OCSMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *ProvisioningApiAppsGetApps200ResponseOcs) SetMeta(v OCSMeta) {
	o.Meta = v
}

// GetData returns the Data field value
func (o *ProvisioningApiAppsGetApps200ResponseOcs) GetData() ProvisioningApiAppsGetApps200ResponseOcsData {
	if o == nil {
		var ret ProvisioningApiAppsGetApps200ResponseOcsData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ProvisioningApiAppsGetApps200ResponseOcs) GetDataOk() (*ProvisioningApiAppsGetApps200ResponseOcsData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ProvisioningApiAppsGetApps200ResponseOcs) SetData(v ProvisioningApiAppsGetApps200ResponseOcsData) {
	o.Data = v
}

func (o ProvisioningApiAppsGetApps200ResponseOcs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProvisioningApiAppsGetApps200ResponseOcs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["data"] = o.Data

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProvisioningApiAppsGetApps200ResponseOcs) UnmarshalJSON(bytes []byte) (err error) {
	varProvisioningApiAppsGetApps200ResponseOcs := _ProvisioningApiAppsGetApps200ResponseOcs{}

	if err = json.Unmarshal(bytes, &varProvisioningApiAppsGetApps200ResponseOcs); err == nil {
		*o = ProvisioningApiAppsGetApps200ResponseOcs(varProvisioningApiAppsGetApps200ResponseOcs)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "meta")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProvisioningApiAppsGetApps200ResponseOcs struct {
	value *ProvisioningApiAppsGetApps200ResponseOcs
	isSet bool
}

func (v NullableProvisioningApiAppsGetApps200ResponseOcs) Get() *ProvisioningApiAppsGetApps200ResponseOcs {
	return v.value
}

func (v *NullableProvisioningApiAppsGetApps200ResponseOcs) Set(val *ProvisioningApiAppsGetApps200ResponseOcs) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisioningApiAppsGetApps200ResponseOcs) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisioningApiAppsGetApps200ResponseOcs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisioningApiAppsGetApps200ResponseOcs(val *ProvisioningApiAppsGetApps200ResponseOcs) *NullableProvisioningApiAppsGetApps200ResponseOcs {
	return &NullableProvisioningApiAppsGetApps200ResponseOcs{value: val, isSet: true}
}

func (v NullableProvisioningApiAppsGetApps200ResponseOcs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisioningApiAppsGetApps200ResponseOcs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

