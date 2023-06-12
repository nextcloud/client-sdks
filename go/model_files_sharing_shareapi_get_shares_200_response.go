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

// checks if the FilesSharingShareapiGetShares200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FilesSharingShareapiGetShares200Response{}

// FilesSharingShareapiGetShares200Response struct for FilesSharingShareapiGetShares200Response
type FilesSharingShareapiGetShares200Response struct {
	Ocs FilesSharingShareapiGetShares200ResponseOcs `json:"ocs"`
}

// NewFilesSharingShareapiGetShares200Response instantiates a new FilesSharingShareapiGetShares200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilesSharingShareapiGetShares200Response(ocs FilesSharingShareapiGetShares200ResponseOcs) *FilesSharingShareapiGetShares200Response {
	this := FilesSharingShareapiGetShares200Response{}
	this.Ocs = ocs
	return &this
}

// NewFilesSharingShareapiGetShares200ResponseWithDefaults instantiates a new FilesSharingShareapiGetShares200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilesSharingShareapiGetShares200ResponseWithDefaults() *FilesSharingShareapiGetShares200Response {
	this := FilesSharingShareapiGetShares200Response{}
	return &this
}

// GetOcs returns the Ocs field value
func (o *FilesSharingShareapiGetShares200Response) GetOcs() FilesSharingShareapiGetShares200ResponseOcs {
	if o == nil {
		var ret FilesSharingShareapiGetShares200ResponseOcs
		return ret
	}

	return o.Ocs
}

// GetOcsOk returns a tuple with the Ocs field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareapiGetShares200Response) GetOcsOk() (*FilesSharingShareapiGetShares200ResponseOcs, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ocs, true
}

// SetOcs sets field value
func (o *FilesSharingShareapiGetShares200Response) SetOcs(v FilesSharingShareapiGetShares200ResponseOcs) {
	o.Ocs = v
}

func (o FilesSharingShareapiGetShares200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FilesSharingShareapiGetShares200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ocs"] = o.Ocs
	return toSerialize, nil
}

type NullableFilesSharingShareapiGetShares200Response struct {
	value *FilesSharingShareapiGetShares200Response
	isSet bool
}

func (v NullableFilesSharingShareapiGetShares200Response) Get() *FilesSharingShareapiGetShares200Response {
	return v.value
}

func (v *NullableFilesSharingShareapiGetShares200Response) Set(val *FilesSharingShareapiGetShares200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableFilesSharingShareapiGetShares200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableFilesSharingShareapiGetShares200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilesSharingShareapiGetShares200Response(val *FilesSharingShareapiGetShares200Response) *NullableFilesSharingShareapiGetShares200Response {
	return &NullableFilesSharingShareapiGetShares200Response{value: val, isSet: true}
}

func (v NullableFilesSharingShareapiGetShares200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilesSharingShareapiGetShares200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


