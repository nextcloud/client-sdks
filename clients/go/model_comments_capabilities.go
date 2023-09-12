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

// checks if the CommentsCapabilities type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommentsCapabilities{}

// CommentsCapabilities struct for CommentsCapabilities
type CommentsCapabilities struct {
	Files CommentsCapabilitiesFiles `json:"files"`
	AdditionalProperties map[string]interface{}
}

type _CommentsCapabilities CommentsCapabilities

// NewCommentsCapabilities instantiates a new CommentsCapabilities object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommentsCapabilities(files CommentsCapabilitiesFiles) *CommentsCapabilities {
	this := CommentsCapabilities{}
	this.Files = files
	return &this
}

// NewCommentsCapabilitiesWithDefaults instantiates a new CommentsCapabilities object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommentsCapabilitiesWithDefaults() *CommentsCapabilities {
	this := CommentsCapabilities{}
	return &this
}

// GetFiles returns the Files field value
func (o *CommentsCapabilities) GetFiles() CommentsCapabilitiesFiles {
	if o == nil {
		var ret CommentsCapabilitiesFiles
		return ret
	}

	return o.Files
}

// GetFilesOk returns a tuple with the Files field value
// and a boolean to check if the value has been set.
func (o *CommentsCapabilities) GetFilesOk() (*CommentsCapabilitiesFiles, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Files, true
}

// SetFiles sets field value
func (o *CommentsCapabilities) SetFiles(v CommentsCapabilitiesFiles) {
	o.Files = v
}

func (o CommentsCapabilities) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommentsCapabilities) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["files"] = o.Files

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CommentsCapabilities) UnmarshalJSON(bytes []byte) (err error) {
	varCommentsCapabilities := _CommentsCapabilities{}

	if err = json.Unmarshal(bytes, &varCommentsCapabilities); err == nil {
		*o = CommentsCapabilities(varCommentsCapabilities)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "files")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCommentsCapabilities struct {
	value *CommentsCapabilities
	isSet bool
}

func (v NullableCommentsCapabilities) Get() *CommentsCapabilities {
	return v.value
}

func (v *NullableCommentsCapabilities) Set(val *CommentsCapabilities) {
	v.value = val
	v.isSet = true
}

func (v NullableCommentsCapabilities) IsSet() bool {
	return v.isSet
}

func (v *NullableCommentsCapabilities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommentsCapabilities(val *CommentsCapabilities) *NullableCommentsCapabilities {
	return &NullableCommentsCapabilities{value: val, isSet: true}
}

func (v NullableCommentsCapabilities) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommentsCapabilities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

