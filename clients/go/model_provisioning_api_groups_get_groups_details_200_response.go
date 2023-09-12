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

// checks if the ProvisioningApiGroupsGetGroupsDetails200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProvisioningApiGroupsGetGroupsDetails200Response{}

// ProvisioningApiGroupsGetGroupsDetails200Response struct for ProvisioningApiGroupsGetGroupsDetails200Response
type ProvisioningApiGroupsGetGroupsDetails200Response struct {
	Ocs ProvisioningApiGroupsGetGroupsDetails200ResponseOcs `json:"ocs"`
	AdditionalProperties map[string]interface{}
}

type _ProvisioningApiGroupsGetGroupsDetails200Response ProvisioningApiGroupsGetGroupsDetails200Response

// NewProvisioningApiGroupsGetGroupsDetails200Response instantiates a new ProvisioningApiGroupsGetGroupsDetails200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisioningApiGroupsGetGroupsDetails200Response(ocs ProvisioningApiGroupsGetGroupsDetails200ResponseOcs) *ProvisioningApiGroupsGetGroupsDetails200Response {
	this := ProvisioningApiGroupsGetGroupsDetails200Response{}
	this.Ocs = ocs
	return &this
}

// NewProvisioningApiGroupsGetGroupsDetails200ResponseWithDefaults instantiates a new ProvisioningApiGroupsGetGroupsDetails200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisioningApiGroupsGetGroupsDetails200ResponseWithDefaults() *ProvisioningApiGroupsGetGroupsDetails200Response {
	this := ProvisioningApiGroupsGetGroupsDetails200Response{}
	return &this
}

// GetOcs returns the Ocs field value
func (o *ProvisioningApiGroupsGetGroupsDetails200Response) GetOcs() ProvisioningApiGroupsGetGroupsDetails200ResponseOcs {
	if o == nil {
		var ret ProvisioningApiGroupsGetGroupsDetails200ResponseOcs
		return ret
	}

	return o.Ocs
}

// GetOcsOk returns a tuple with the Ocs field value
// and a boolean to check if the value has been set.
func (o *ProvisioningApiGroupsGetGroupsDetails200Response) GetOcsOk() (*ProvisioningApiGroupsGetGroupsDetails200ResponseOcs, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ocs, true
}

// SetOcs sets field value
func (o *ProvisioningApiGroupsGetGroupsDetails200Response) SetOcs(v ProvisioningApiGroupsGetGroupsDetails200ResponseOcs) {
	o.Ocs = v
}

func (o ProvisioningApiGroupsGetGroupsDetails200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProvisioningApiGroupsGetGroupsDetails200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ocs"] = o.Ocs

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProvisioningApiGroupsGetGroupsDetails200Response) UnmarshalJSON(bytes []byte) (err error) {
	varProvisioningApiGroupsGetGroupsDetails200Response := _ProvisioningApiGroupsGetGroupsDetails200Response{}

	if err = json.Unmarshal(bytes, &varProvisioningApiGroupsGetGroupsDetails200Response); err == nil {
		*o = ProvisioningApiGroupsGetGroupsDetails200Response(varProvisioningApiGroupsGetGroupsDetails200Response)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ocs")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProvisioningApiGroupsGetGroupsDetails200Response struct {
	value *ProvisioningApiGroupsGetGroupsDetails200Response
	isSet bool
}

func (v NullableProvisioningApiGroupsGetGroupsDetails200Response) Get() *ProvisioningApiGroupsGetGroupsDetails200Response {
	return v.value
}

func (v *NullableProvisioningApiGroupsGetGroupsDetails200Response) Set(val *ProvisioningApiGroupsGetGroupsDetails200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisioningApiGroupsGetGroupsDetails200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisioningApiGroupsGetGroupsDetails200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisioningApiGroupsGetGroupsDetails200Response(val *ProvisioningApiGroupsGetGroupsDetails200Response) *NullableProvisioningApiGroupsGetGroupsDetails200Response {
	return &NullableProvisioningApiGroupsGetGroupsDetails200Response{value: val, isSet: true}
}

func (v NullableProvisioningApiGroupsGetGroupsDetails200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisioningApiGroupsGetGroupsDetails200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


