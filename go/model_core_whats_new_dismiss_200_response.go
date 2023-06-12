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

// checks if the CoreWhatsNewDismiss200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreWhatsNewDismiss200Response{}

// CoreWhatsNewDismiss200Response struct for CoreWhatsNewDismiss200Response
type CoreWhatsNewDismiss200Response struct {
	Ocs CoreWhatsNewDismiss200ResponseOcs `json:"ocs"`
}

// NewCoreWhatsNewDismiss200Response instantiates a new CoreWhatsNewDismiss200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreWhatsNewDismiss200Response(ocs CoreWhatsNewDismiss200ResponseOcs) *CoreWhatsNewDismiss200Response {
	this := CoreWhatsNewDismiss200Response{}
	this.Ocs = ocs
	return &this
}

// NewCoreWhatsNewDismiss200ResponseWithDefaults instantiates a new CoreWhatsNewDismiss200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreWhatsNewDismiss200ResponseWithDefaults() *CoreWhatsNewDismiss200Response {
	this := CoreWhatsNewDismiss200Response{}
	return &this
}

// GetOcs returns the Ocs field value
func (o *CoreWhatsNewDismiss200Response) GetOcs() CoreWhatsNewDismiss200ResponseOcs {
	if o == nil {
		var ret CoreWhatsNewDismiss200ResponseOcs
		return ret
	}

	return o.Ocs
}

// GetOcsOk returns a tuple with the Ocs field value
// and a boolean to check if the value has been set.
func (o *CoreWhatsNewDismiss200Response) GetOcsOk() (*CoreWhatsNewDismiss200ResponseOcs, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ocs, true
}

// SetOcs sets field value
func (o *CoreWhatsNewDismiss200Response) SetOcs(v CoreWhatsNewDismiss200ResponseOcs) {
	o.Ocs = v
}

func (o CoreWhatsNewDismiss200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreWhatsNewDismiss200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ocs"] = o.Ocs
	return toSerialize, nil
}

type NullableCoreWhatsNewDismiss200Response struct {
	value *CoreWhatsNewDismiss200Response
	isSet bool
}

func (v NullableCoreWhatsNewDismiss200Response) Get() *CoreWhatsNewDismiss200Response {
	return v.value
}

func (v *NullableCoreWhatsNewDismiss200Response) Set(val *CoreWhatsNewDismiss200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreWhatsNewDismiss200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreWhatsNewDismiss200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreWhatsNewDismiss200Response(val *CoreWhatsNewDismiss200Response) *NullableCoreWhatsNewDismiss200Response {
	return &NullableCoreWhatsNewDismiss200Response{value: val, isSet: true}
}

func (v NullableCoreWhatsNewDismiss200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreWhatsNewDismiss200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


