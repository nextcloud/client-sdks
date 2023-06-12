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

// checks if the FilesSharingShareesapiFindRecommended200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FilesSharingShareesapiFindRecommended200Response{}

// FilesSharingShareesapiFindRecommended200Response struct for FilesSharingShareesapiFindRecommended200Response
type FilesSharingShareesapiFindRecommended200Response struct {
	Ocs FilesSharingShareesapiFindRecommended200ResponseOcs `json:"ocs"`
}

// NewFilesSharingShareesapiFindRecommended200Response instantiates a new FilesSharingShareesapiFindRecommended200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilesSharingShareesapiFindRecommended200Response(ocs FilesSharingShareesapiFindRecommended200ResponseOcs) *FilesSharingShareesapiFindRecommended200Response {
	this := FilesSharingShareesapiFindRecommended200Response{}
	this.Ocs = ocs
	return &this
}

// NewFilesSharingShareesapiFindRecommended200ResponseWithDefaults instantiates a new FilesSharingShareesapiFindRecommended200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilesSharingShareesapiFindRecommended200ResponseWithDefaults() *FilesSharingShareesapiFindRecommended200Response {
	this := FilesSharingShareesapiFindRecommended200Response{}
	return &this
}

// GetOcs returns the Ocs field value
func (o *FilesSharingShareesapiFindRecommended200Response) GetOcs() FilesSharingShareesapiFindRecommended200ResponseOcs {
	if o == nil {
		var ret FilesSharingShareesapiFindRecommended200ResponseOcs
		return ret
	}

	return o.Ocs
}

// GetOcsOk returns a tuple with the Ocs field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareesapiFindRecommended200Response) GetOcsOk() (*FilesSharingShareesapiFindRecommended200ResponseOcs, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ocs, true
}

// SetOcs sets field value
func (o *FilesSharingShareesapiFindRecommended200Response) SetOcs(v FilesSharingShareesapiFindRecommended200ResponseOcs) {
	o.Ocs = v
}

func (o FilesSharingShareesapiFindRecommended200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FilesSharingShareesapiFindRecommended200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ocs"] = o.Ocs
	return toSerialize, nil
}

type NullableFilesSharingShareesapiFindRecommended200Response struct {
	value *FilesSharingShareesapiFindRecommended200Response
	isSet bool
}

func (v NullableFilesSharingShareesapiFindRecommended200Response) Get() *FilesSharingShareesapiFindRecommended200Response {
	return v.value
}

func (v *NullableFilesSharingShareesapiFindRecommended200Response) Set(val *FilesSharingShareesapiFindRecommended200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableFilesSharingShareesapiFindRecommended200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableFilesSharingShareesapiFindRecommended200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilesSharingShareesapiFindRecommended200Response(val *FilesSharingShareesapiFindRecommended200Response) *NullableFilesSharingShareesapiFindRecommended200Response {
	return &NullableFilesSharingShareesapiFindRecommended200Response{value: val, isSet: true}
}

func (v NullableFilesSharingShareesapiFindRecommended200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilesSharingShareesapiFindRecommended200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


