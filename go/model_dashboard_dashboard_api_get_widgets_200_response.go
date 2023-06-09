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

// checks if the DashboardDashboardApiGetWidgets200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DashboardDashboardApiGetWidgets200Response{}

// DashboardDashboardApiGetWidgets200Response struct for DashboardDashboardApiGetWidgets200Response
type DashboardDashboardApiGetWidgets200Response struct {
	Ocs DashboardDashboardApiGetWidgets200ResponseOcs `json:"ocs"`
}

// NewDashboardDashboardApiGetWidgets200Response instantiates a new DashboardDashboardApiGetWidgets200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDashboardDashboardApiGetWidgets200Response(ocs DashboardDashboardApiGetWidgets200ResponseOcs) *DashboardDashboardApiGetWidgets200Response {
	this := DashboardDashboardApiGetWidgets200Response{}
	this.Ocs = ocs
	return &this
}

// NewDashboardDashboardApiGetWidgets200ResponseWithDefaults instantiates a new DashboardDashboardApiGetWidgets200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDashboardDashboardApiGetWidgets200ResponseWithDefaults() *DashboardDashboardApiGetWidgets200Response {
	this := DashboardDashboardApiGetWidgets200Response{}
	return &this
}

// GetOcs returns the Ocs field value
func (o *DashboardDashboardApiGetWidgets200Response) GetOcs() DashboardDashboardApiGetWidgets200ResponseOcs {
	if o == nil {
		var ret DashboardDashboardApiGetWidgets200ResponseOcs
		return ret
	}

	return o.Ocs
}

// GetOcsOk returns a tuple with the Ocs field value
// and a boolean to check if the value has been set.
func (o *DashboardDashboardApiGetWidgets200Response) GetOcsOk() (*DashboardDashboardApiGetWidgets200ResponseOcs, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ocs, true
}

// SetOcs sets field value
func (o *DashboardDashboardApiGetWidgets200Response) SetOcs(v DashboardDashboardApiGetWidgets200ResponseOcs) {
	o.Ocs = v
}

func (o DashboardDashboardApiGetWidgets200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DashboardDashboardApiGetWidgets200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ocs"] = o.Ocs
	return toSerialize, nil
}

type NullableDashboardDashboardApiGetWidgets200Response struct {
	value *DashboardDashboardApiGetWidgets200Response
	isSet bool
}

func (v NullableDashboardDashboardApiGetWidgets200Response) Get() *DashboardDashboardApiGetWidgets200Response {
	return v.value
}

func (v *NullableDashboardDashboardApiGetWidgets200Response) Set(val *DashboardDashboardApiGetWidgets200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboardDashboardApiGetWidgets200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboardDashboardApiGetWidgets200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboardDashboardApiGetWidgets200Response(val *DashboardDashboardApiGetWidgets200Response) *NullableDashboardDashboardApiGetWidgets200Response {
	return &NullableDashboardDashboardApiGetWidgets200Response{value: val, isSet: true}
}

func (v NullableDashboardDashboardApiGetWidgets200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboardDashboardApiGetWidgets200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


