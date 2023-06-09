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

// checks if the CoreAppPasswordGetAppPassword200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreAppPasswordGetAppPassword200Response{}

// CoreAppPasswordGetAppPassword200Response struct for CoreAppPasswordGetAppPassword200Response
type CoreAppPasswordGetAppPassword200Response struct {
	Ocs CoreAppPasswordGetAppPassword200ResponseOcs `json:"ocs"`
}

// NewCoreAppPasswordGetAppPassword200Response instantiates a new CoreAppPasswordGetAppPassword200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreAppPasswordGetAppPassword200Response(ocs CoreAppPasswordGetAppPassword200ResponseOcs) *CoreAppPasswordGetAppPassword200Response {
	this := CoreAppPasswordGetAppPassword200Response{}
	this.Ocs = ocs
	return &this
}

// NewCoreAppPasswordGetAppPassword200ResponseWithDefaults instantiates a new CoreAppPasswordGetAppPassword200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreAppPasswordGetAppPassword200ResponseWithDefaults() *CoreAppPasswordGetAppPassword200Response {
	this := CoreAppPasswordGetAppPassword200Response{}
	return &this
}

// GetOcs returns the Ocs field value
func (o *CoreAppPasswordGetAppPassword200Response) GetOcs() CoreAppPasswordGetAppPassword200ResponseOcs {
	if o == nil {
		var ret CoreAppPasswordGetAppPassword200ResponseOcs
		return ret
	}

	return o.Ocs
}

// GetOcsOk returns a tuple with the Ocs field value
// and a boolean to check if the value has been set.
func (o *CoreAppPasswordGetAppPassword200Response) GetOcsOk() (*CoreAppPasswordGetAppPassword200ResponseOcs, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ocs, true
}

// SetOcs sets field value
func (o *CoreAppPasswordGetAppPassword200Response) SetOcs(v CoreAppPasswordGetAppPassword200ResponseOcs) {
	o.Ocs = v
}

func (o CoreAppPasswordGetAppPassword200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreAppPasswordGetAppPassword200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ocs"] = o.Ocs
	return toSerialize, nil
}

type NullableCoreAppPasswordGetAppPassword200Response struct {
	value *CoreAppPasswordGetAppPassword200Response
	isSet bool
}

func (v NullableCoreAppPasswordGetAppPassword200Response) Get() *CoreAppPasswordGetAppPassword200Response {
	return v.value
}

func (v *NullableCoreAppPasswordGetAppPassword200Response) Set(val *CoreAppPasswordGetAppPassword200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreAppPasswordGetAppPassword200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreAppPasswordGetAppPassword200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreAppPasswordGetAppPassword200Response(val *CoreAppPasswordGetAppPassword200Response) *NullableCoreAppPasswordGetAppPassword200Response {
	return &NullableCoreAppPasswordGetAppPassword200Response{value: val, isSet: true}
}

func (v NullableCoreAppPasswordGetAppPassword200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreAppPasswordGetAppPassword200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


