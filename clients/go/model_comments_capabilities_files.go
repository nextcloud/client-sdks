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

// checks if the CommentsCapabilitiesFiles type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommentsCapabilitiesFiles{}

// CommentsCapabilitiesFiles struct for CommentsCapabilitiesFiles
type CommentsCapabilitiesFiles struct {
	Comments bool `json:"comments"`
	AdditionalProperties map[string]interface{}
}

type _CommentsCapabilitiesFiles CommentsCapabilitiesFiles

// NewCommentsCapabilitiesFiles instantiates a new CommentsCapabilitiesFiles object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommentsCapabilitiesFiles(comments bool) *CommentsCapabilitiesFiles {
	this := CommentsCapabilitiesFiles{}
	this.Comments = comments
	return &this
}

// NewCommentsCapabilitiesFilesWithDefaults instantiates a new CommentsCapabilitiesFiles object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommentsCapabilitiesFilesWithDefaults() *CommentsCapabilitiesFiles {
	this := CommentsCapabilitiesFiles{}
	return &this
}

// GetComments returns the Comments field value
func (o *CommentsCapabilitiesFiles) GetComments() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value
// and a boolean to check if the value has been set.
func (o *CommentsCapabilitiesFiles) GetCommentsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Comments, true
}

// SetComments sets field value
func (o *CommentsCapabilitiesFiles) SetComments(v bool) {
	o.Comments = v
}

func (o CommentsCapabilitiesFiles) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommentsCapabilitiesFiles) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["comments"] = o.Comments

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CommentsCapabilitiesFiles) UnmarshalJSON(bytes []byte) (err error) {
	varCommentsCapabilitiesFiles := _CommentsCapabilitiesFiles{}

	if err = json.Unmarshal(bytes, &varCommentsCapabilitiesFiles); err == nil {
		*o = CommentsCapabilitiesFiles(varCommentsCapabilitiesFiles)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "comments")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCommentsCapabilitiesFiles struct {
	value *CommentsCapabilitiesFiles
	isSet bool
}

func (v NullableCommentsCapabilitiesFiles) Get() *CommentsCapabilitiesFiles {
	return v.value
}

func (v *NullableCommentsCapabilitiesFiles) Set(val *CommentsCapabilitiesFiles) {
	v.value = val
	v.isSet = true
}

func (v NullableCommentsCapabilitiesFiles) IsSet() bool {
	return v.isSet
}

func (v *NullableCommentsCapabilitiesFiles) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommentsCapabilitiesFiles(val *CommentsCapabilitiesFiles) *NullableCommentsCapabilitiesFiles {
	return &NullableCommentsCapabilitiesFiles{value: val, isSet: true}
}

func (v NullableCommentsCapabilitiesFiles) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommentsCapabilitiesFiles) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

