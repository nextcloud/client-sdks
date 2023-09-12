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

// checks if the FilesSharingDeletedShareapiList200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FilesSharingDeletedShareapiList200Response{}

// FilesSharingDeletedShareapiList200Response struct for FilesSharingDeletedShareapiList200Response
type FilesSharingDeletedShareapiList200Response struct {
	Ocs FilesSharingDeletedShareapiList200ResponseOcs `json:"ocs"`
	AdditionalProperties map[string]interface{}
}

type _FilesSharingDeletedShareapiList200Response FilesSharingDeletedShareapiList200Response

// NewFilesSharingDeletedShareapiList200Response instantiates a new FilesSharingDeletedShareapiList200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilesSharingDeletedShareapiList200Response(ocs FilesSharingDeletedShareapiList200ResponseOcs) *FilesSharingDeletedShareapiList200Response {
	this := FilesSharingDeletedShareapiList200Response{}
	this.Ocs = ocs
	return &this
}

// NewFilesSharingDeletedShareapiList200ResponseWithDefaults instantiates a new FilesSharingDeletedShareapiList200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilesSharingDeletedShareapiList200ResponseWithDefaults() *FilesSharingDeletedShareapiList200Response {
	this := FilesSharingDeletedShareapiList200Response{}
	return &this
}

// GetOcs returns the Ocs field value
func (o *FilesSharingDeletedShareapiList200Response) GetOcs() FilesSharingDeletedShareapiList200ResponseOcs {
	if o == nil {
		var ret FilesSharingDeletedShareapiList200ResponseOcs
		return ret
	}

	return o.Ocs
}

// GetOcsOk returns a tuple with the Ocs field value
// and a boolean to check if the value has been set.
func (o *FilesSharingDeletedShareapiList200Response) GetOcsOk() (*FilesSharingDeletedShareapiList200ResponseOcs, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ocs, true
}

// SetOcs sets field value
func (o *FilesSharingDeletedShareapiList200Response) SetOcs(v FilesSharingDeletedShareapiList200ResponseOcs) {
	o.Ocs = v
}

func (o FilesSharingDeletedShareapiList200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FilesSharingDeletedShareapiList200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ocs"] = o.Ocs

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FilesSharingDeletedShareapiList200Response) UnmarshalJSON(bytes []byte) (err error) {
	varFilesSharingDeletedShareapiList200Response := _FilesSharingDeletedShareapiList200Response{}

	if err = json.Unmarshal(bytes, &varFilesSharingDeletedShareapiList200Response); err == nil {
		*o = FilesSharingDeletedShareapiList200Response(varFilesSharingDeletedShareapiList200Response)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ocs")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFilesSharingDeletedShareapiList200Response struct {
	value *FilesSharingDeletedShareapiList200Response
	isSet bool
}

func (v NullableFilesSharingDeletedShareapiList200Response) Get() *FilesSharingDeletedShareapiList200Response {
	return v.value
}

func (v *NullableFilesSharingDeletedShareapiList200Response) Set(val *FilesSharingDeletedShareapiList200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableFilesSharingDeletedShareapiList200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableFilesSharingDeletedShareapiList200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilesSharingDeletedShareapiList200Response(val *FilesSharingDeletedShareapiList200Response) *NullableFilesSharingDeletedShareapiList200Response {
	return &NullableFilesSharingDeletedShareapiList200Response{value: val, isSet: true}
}

func (v NullableFilesSharingDeletedShareapiList200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilesSharingDeletedShareapiList200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

