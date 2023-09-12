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

// checks if the FilesCapabilities type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FilesCapabilities{}

// FilesCapabilities struct for FilesCapabilities
type FilesCapabilities struct {
	Files FilesCapabilitiesFiles `json:"files"`
	AdditionalProperties map[string]interface{}
}

type _FilesCapabilities FilesCapabilities

// NewFilesCapabilities instantiates a new FilesCapabilities object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilesCapabilities(files FilesCapabilitiesFiles) *FilesCapabilities {
	this := FilesCapabilities{}
	this.Files = files
	return &this
}

// NewFilesCapabilitiesWithDefaults instantiates a new FilesCapabilities object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilesCapabilitiesWithDefaults() *FilesCapabilities {
	this := FilesCapabilities{}
	return &this
}

// GetFiles returns the Files field value
func (o *FilesCapabilities) GetFiles() FilesCapabilitiesFiles {
	if o == nil {
		var ret FilesCapabilitiesFiles
		return ret
	}

	return o.Files
}

// GetFilesOk returns a tuple with the Files field value
// and a boolean to check if the value has been set.
func (o *FilesCapabilities) GetFilesOk() (*FilesCapabilitiesFiles, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Files, true
}

// SetFiles sets field value
func (o *FilesCapabilities) SetFiles(v FilesCapabilitiesFiles) {
	o.Files = v
}

func (o FilesCapabilities) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FilesCapabilities) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["files"] = o.Files

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FilesCapabilities) UnmarshalJSON(bytes []byte) (err error) {
	varFilesCapabilities := _FilesCapabilities{}

	if err = json.Unmarshal(bytes, &varFilesCapabilities); err == nil {
		*o = FilesCapabilities(varFilesCapabilities)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "files")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFilesCapabilities struct {
	value *FilesCapabilities
	isSet bool
}

func (v NullableFilesCapabilities) Get() *FilesCapabilities {
	return v.value
}

func (v *NullableFilesCapabilities) Set(val *FilesCapabilities) {
	v.value = val
	v.isSet = true
}

func (v NullableFilesCapabilities) IsSet() bool {
	return v.isSet
}

func (v *NullableFilesCapabilities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilesCapabilities(val *FilesCapabilities) *NullableFilesCapabilities {
	return &NullableFilesCapabilities{value: val, isSet: true}
}

func (v NullableFilesCapabilities) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilesCapabilities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

