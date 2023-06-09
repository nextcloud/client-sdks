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

// checks if the FilesSharingShareapiCreateShare200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FilesSharingShareapiCreateShare200Response{}

// FilesSharingShareapiCreateShare200Response struct for FilesSharingShareapiCreateShare200Response
type FilesSharingShareapiCreateShare200Response struct {
	Ocs FilesSharingShareapiCreateShare200ResponseOcs `json:"ocs"`
}

// NewFilesSharingShareapiCreateShare200Response instantiates a new FilesSharingShareapiCreateShare200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilesSharingShareapiCreateShare200Response(ocs FilesSharingShareapiCreateShare200ResponseOcs) *FilesSharingShareapiCreateShare200Response {
	this := FilesSharingShareapiCreateShare200Response{}
	this.Ocs = ocs
	return &this
}

// NewFilesSharingShareapiCreateShare200ResponseWithDefaults instantiates a new FilesSharingShareapiCreateShare200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilesSharingShareapiCreateShare200ResponseWithDefaults() *FilesSharingShareapiCreateShare200Response {
	this := FilesSharingShareapiCreateShare200Response{}
	return &this
}

// GetOcs returns the Ocs field value
func (o *FilesSharingShareapiCreateShare200Response) GetOcs() FilesSharingShareapiCreateShare200ResponseOcs {
	if o == nil {
		var ret FilesSharingShareapiCreateShare200ResponseOcs
		return ret
	}

	return o.Ocs
}

// GetOcsOk returns a tuple with the Ocs field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareapiCreateShare200Response) GetOcsOk() (*FilesSharingShareapiCreateShare200ResponseOcs, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ocs, true
}

// SetOcs sets field value
func (o *FilesSharingShareapiCreateShare200Response) SetOcs(v FilesSharingShareapiCreateShare200ResponseOcs) {
	o.Ocs = v
}

func (o FilesSharingShareapiCreateShare200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FilesSharingShareapiCreateShare200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ocs"] = o.Ocs
	return toSerialize, nil
}

type NullableFilesSharingShareapiCreateShare200Response struct {
	value *FilesSharingShareapiCreateShare200Response
	isSet bool
}

func (v NullableFilesSharingShareapiCreateShare200Response) Get() *FilesSharingShareapiCreateShare200Response {
	return v.value
}

func (v *NullableFilesSharingShareapiCreateShare200Response) Set(val *FilesSharingShareapiCreateShare200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableFilesSharingShareapiCreateShare200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableFilesSharingShareapiCreateShare200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilesSharingShareapiCreateShare200Response(val *FilesSharingShareapiCreateShare200Response) *NullableFilesSharingShareapiCreateShare200Response {
	return &NullableFilesSharingShareapiCreateShare200Response{value: val, isSet: true}
}

func (v NullableFilesSharingShareapiCreateShare200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilesSharingShareapiCreateShare200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


