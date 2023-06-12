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

// checks if the ProvisioningApiAppConfigGetValue200ResponseOcsData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProvisioningApiAppConfigGetValue200ResponseOcsData{}

// ProvisioningApiAppConfigGetValue200ResponseOcsData struct for ProvisioningApiAppConfigGetValue200ResponseOcsData
type ProvisioningApiAppConfigGetValue200ResponseOcsData struct {
	Data string `json:"data"`
}

// NewProvisioningApiAppConfigGetValue200ResponseOcsData instantiates a new ProvisioningApiAppConfigGetValue200ResponseOcsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisioningApiAppConfigGetValue200ResponseOcsData(data string) *ProvisioningApiAppConfigGetValue200ResponseOcsData {
	this := ProvisioningApiAppConfigGetValue200ResponseOcsData{}
	this.Data = data
	return &this
}

// NewProvisioningApiAppConfigGetValue200ResponseOcsDataWithDefaults instantiates a new ProvisioningApiAppConfigGetValue200ResponseOcsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisioningApiAppConfigGetValue200ResponseOcsDataWithDefaults() *ProvisioningApiAppConfigGetValue200ResponseOcsData {
	this := ProvisioningApiAppConfigGetValue200ResponseOcsData{}
	return &this
}

// GetData returns the Data field value
func (o *ProvisioningApiAppConfigGetValue200ResponseOcsData) GetData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ProvisioningApiAppConfigGetValue200ResponseOcsData) GetDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ProvisioningApiAppConfigGetValue200ResponseOcsData) SetData(v string) {
	o.Data = v
}

func (o ProvisioningApiAppConfigGetValue200ResponseOcsData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProvisioningApiAppConfigGetValue200ResponseOcsData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableProvisioningApiAppConfigGetValue200ResponseOcsData struct {
	value *ProvisioningApiAppConfigGetValue200ResponseOcsData
	isSet bool
}

func (v NullableProvisioningApiAppConfigGetValue200ResponseOcsData) Get() *ProvisioningApiAppConfigGetValue200ResponseOcsData {
	return v.value
}

func (v *NullableProvisioningApiAppConfigGetValue200ResponseOcsData) Set(val *ProvisioningApiAppConfigGetValue200ResponseOcsData) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisioningApiAppConfigGetValue200ResponseOcsData) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisioningApiAppConfigGetValue200ResponseOcsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisioningApiAppConfigGetValue200ResponseOcsData(val *ProvisioningApiAppConfigGetValue200ResponseOcsData) *NullableProvisioningApiAppConfigGetValue200ResponseOcsData {
	return &NullableProvisioningApiAppConfigGetValue200ResponseOcsData{value: val, isSet: true}
}

func (v NullableProvisioningApiAppConfigGetValue200ResponseOcsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisioningApiAppConfigGetValue200ResponseOcsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


