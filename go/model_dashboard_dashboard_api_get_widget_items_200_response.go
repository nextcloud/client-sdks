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

// checks if the DashboardDashboardApiGetWidgetItems200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DashboardDashboardApiGetWidgetItems200Response{}

// DashboardDashboardApiGetWidgetItems200Response struct for DashboardDashboardApiGetWidgetItems200Response
type DashboardDashboardApiGetWidgetItems200Response struct {
	Ocs DashboardDashboardApiGetWidgetItems200ResponseOcs `json:"ocs"`
}

// NewDashboardDashboardApiGetWidgetItems200Response instantiates a new DashboardDashboardApiGetWidgetItems200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDashboardDashboardApiGetWidgetItems200Response(ocs DashboardDashboardApiGetWidgetItems200ResponseOcs) *DashboardDashboardApiGetWidgetItems200Response {
	this := DashboardDashboardApiGetWidgetItems200Response{}
	this.Ocs = ocs
	return &this
}

// NewDashboardDashboardApiGetWidgetItems200ResponseWithDefaults instantiates a new DashboardDashboardApiGetWidgetItems200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDashboardDashboardApiGetWidgetItems200ResponseWithDefaults() *DashboardDashboardApiGetWidgetItems200Response {
	this := DashboardDashboardApiGetWidgetItems200Response{}
	return &this
}

// GetOcs returns the Ocs field value
func (o *DashboardDashboardApiGetWidgetItems200Response) GetOcs() DashboardDashboardApiGetWidgetItems200ResponseOcs {
	if o == nil {
		var ret DashboardDashboardApiGetWidgetItems200ResponseOcs
		return ret
	}

	return o.Ocs
}

// GetOcsOk returns a tuple with the Ocs field value
// and a boolean to check if the value has been set.
func (o *DashboardDashboardApiGetWidgetItems200Response) GetOcsOk() (*DashboardDashboardApiGetWidgetItems200ResponseOcs, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ocs, true
}

// SetOcs sets field value
func (o *DashboardDashboardApiGetWidgetItems200Response) SetOcs(v DashboardDashboardApiGetWidgetItems200ResponseOcs) {
	o.Ocs = v
}

func (o DashboardDashboardApiGetWidgetItems200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DashboardDashboardApiGetWidgetItems200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ocs"] = o.Ocs
	return toSerialize, nil
}

type NullableDashboardDashboardApiGetWidgetItems200Response struct {
	value *DashboardDashboardApiGetWidgetItems200Response
	isSet bool
}

func (v NullableDashboardDashboardApiGetWidgetItems200Response) Get() *DashboardDashboardApiGetWidgetItems200Response {
	return v.value
}

func (v *NullableDashboardDashboardApiGetWidgetItems200Response) Set(val *DashboardDashboardApiGetWidgetItems200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboardDashboardApiGetWidgetItems200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboardDashboardApiGetWidgetItems200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboardDashboardApiGetWidgetItems200Response(val *DashboardDashboardApiGetWidgetItems200Response) *NullableDashboardDashboardApiGetWidgetItems200Response {
	return &NullableDashboardDashboardApiGetWidgetItems200Response{value: val, isSet: true}
}

func (v NullableDashboardDashboardApiGetWidgetItems200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboardDashboardApiGetWidgetItems200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


