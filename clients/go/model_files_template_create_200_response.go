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

// checks if the FilesTemplateCreate200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FilesTemplateCreate200Response{}

// FilesTemplateCreate200Response struct for FilesTemplateCreate200Response
type FilesTemplateCreate200Response struct {
	Ocs FilesTemplateCreate200ResponseOcs `json:"ocs"`
	AdditionalProperties map[string]interface{}
}

type _FilesTemplateCreate200Response FilesTemplateCreate200Response

// NewFilesTemplateCreate200Response instantiates a new FilesTemplateCreate200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilesTemplateCreate200Response(ocs FilesTemplateCreate200ResponseOcs) *FilesTemplateCreate200Response {
	this := FilesTemplateCreate200Response{}
	this.Ocs = ocs
	return &this
}

// NewFilesTemplateCreate200ResponseWithDefaults instantiates a new FilesTemplateCreate200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilesTemplateCreate200ResponseWithDefaults() *FilesTemplateCreate200Response {
	this := FilesTemplateCreate200Response{}
	return &this
}

// GetOcs returns the Ocs field value
func (o *FilesTemplateCreate200Response) GetOcs() FilesTemplateCreate200ResponseOcs {
	if o == nil {
		var ret FilesTemplateCreate200ResponseOcs
		return ret
	}

	return o.Ocs
}

// GetOcsOk returns a tuple with the Ocs field value
// and a boolean to check if the value has been set.
func (o *FilesTemplateCreate200Response) GetOcsOk() (*FilesTemplateCreate200ResponseOcs, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ocs, true
}

// SetOcs sets field value
func (o *FilesTemplateCreate200Response) SetOcs(v FilesTemplateCreate200ResponseOcs) {
	o.Ocs = v
}

func (o FilesTemplateCreate200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FilesTemplateCreate200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ocs"] = o.Ocs

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FilesTemplateCreate200Response) UnmarshalJSON(bytes []byte) (err error) {
	varFilesTemplateCreate200Response := _FilesTemplateCreate200Response{}

	if err = json.Unmarshal(bytes, &varFilesTemplateCreate200Response); err == nil {
		*o = FilesTemplateCreate200Response(varFilesTemplateCreate200Response)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ocs")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFilesTemplateCreate200Response struct {
	value *FilesTemplateCreate200Response
	isSet bool
}

func (v NullableFilesTemplateCreate200Response) Get() *FilesTemplateCreate200Response {
	return v.value
}

func (v *NullableFilesTemplateCreate200Response) Set(val *FilesTemplateCreate200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableFilesTemplateCreate200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableFilesTemplateCreate200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilesTemplateCreate200Response(val *FilesTemplateCreate200Response) *NullableFilesTemplateCreate200Response {
	return &NullableFilesTemplateCreate200Response{value: val, isSet: true}
}

func (v NullableFilesTemplateCreate200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilesTemplateCreate200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


