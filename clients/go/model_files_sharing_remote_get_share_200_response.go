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

// checks if the FilesSharingRemoteGetShare200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FilesSharingRemoteGetShare200Response{}

// FilesSharingRemoteGetShare200Response struct for FilesSharingRemoteGetShare200Response
type FilesSharingRemoteGetShare200Response struct {
	Ocs FilesSharingRemoteGetShare200ResponseOcs `json:"ocs"`
	AdditionalProperties map[string]interface{}
}

type _FilesSharingRemoteGetShare200Response FilesSharingRemoteGetShare200Response

// NewFilesSharingRemoteGetShare200Response instantiates a new FilesSharingRemoteGetShare200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilesSharingRemoteGetShare200Response(ocs FilesSharingRemoteGetShare200ResponseOcs) *FilesSharingRemoteGetShare200Response {
	this := FilesSharingRemoteGetShare200Response{}
	this.Ocs = ocs
	return &this
}

// NewFilesSharingRemoteGetShare200ResponseWithDefaults instantiates a new FilesSharingRemoteGetShare200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilesSharingRemoteGetShare200ResponseWithDefaults() *FilesSharingRemoteGetShare200Response {
	this := FilesSharingRemoteGetShare200Response{}
	return &this
}

// GetOcs returns the Ocs field value
func (o *FilesSharingRemoteGetShare200Response) GetOcs() FilesSharingRemoteGetShare200ResponseOcs {
	if o == nil {
		var ret FilesSharingRemoteGetShare200ResponseOcs
		return ret
	}

	return o.Ocs
}

// GetOcsOk returns a tuple with the Ocs field value
// and a boolean to check if the value has been set.
func (o *FilesSharingRemoteGetShare200Response) GetOcsOk() (*FilesSharingRemoteGetShare200ResponseOcs, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ocs, true
}

// SetOcs sets field value
func (o *FilesSharingRemoteGetShare200Response) SetOcs(v FilesSharingRemoteGetShare200ResponseOcs) {
	o.Ocs = v
}

func (o FilesSharingRemoteGetShare200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FilesSharingRemoteGetShare200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ocs"] = o.Ocs

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FilesSharingRemoteGetShare200Response) UnmarshalJSON(bytes []byte) (err error) {
	varFilesSharingRemoteGetShare200Response := _FilesSharingRemoteGetShare200Response{}

	if err = json.Unmarshal(bytes, &varFilesSharingRemoteGetShare200Response); err == nil {
		*o = FilesSharingRemoteGetShare200Response(varFilesSharingRemoteGetShare200Response)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ocs")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFilesSharingRemoteGetShare200Response struct {
	value *FilesSharingRemoteGetShare200Response
	isSet bool
}

func (v NullableFilesSharingRemoteGetShare200Response) Get() *FilesSharingRemoteGetShare200Response {
	return v.value
}

func (v *NullableFilesSharingRemoteGetShare200Response) Set(val *FilesSharingRemoteGetShare200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableFilesSharingRemoteGetShare200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableFilesSharingRemoteGetShare200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilesSharingRemoteGetShare200Response(val *FilesSharingRemoteGetShare200Response) *NullableFilesSharingRemoteGetShare200Response {
	return &NullableFilesSharingRemoteGetShare200Response{value: val, isSet: true}
}

func (v NullableFilesSharingRemoteGetShare200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilesSharingRemoteGetShare200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

