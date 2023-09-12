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

// checks if the FilesSharingCapabilitiesFilesSharingSharee type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FilesSharingCapabilitiesFilesSharingSharee{}

// FilesSharingCapabilitiesFilesSharingSharee struct for FilesSharingCapabilitiesFilesSharingSharee
type FilesSharingCapabilitiesFilesSharingSharee struct {
	QueryLookupDefault bool `json:"query_lookup_default"`
	AlwaysShowUnique bool `json:"always_show_unique"`
	AdditionalProperties map[string]interface{}
}

type _FilesSharingCapabilitiesFilesSharingSharee FilesSharingCapabilitiesFilesSharingSharee

// NewFilesSharingCapabilitiesFilesSharingSharee instantiates a new FilesSharingCapabilitiesFilesSharingSharee object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilesSharingCapabilitiesFilesSharingSharee(queryLookupDefault bool, alwaysShowUnique bool) *FilesSharingCapabilitiesFilesSharingSharee {
	this := FilesSharingCapabilitiesFilesSharingSharee{}
	this.QueryLookupDefault = queryLookupDefault
	this.AlwaysShowUnique = alwaysShowUnique
	return &this
}

// NewFilesSharingCapabilitiesFilesSharingShareeWithDefaults instantiates a new FilesSharingCapabilitiesFilesSharingSharee object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilesSharingCapabilitiesFilesSharingShareeWithDefaults() *FilesSharingCapabilitiesFilesSharingSharee {
	this := FilesSharingCapabilitiesFilesSharingSharee{}
	return &this
}

// GetQueryLookupDefault returns the QueryLookupDefault field value
func (o *FilesSharingCapabilitiesFilesSharingSharee) GetQueryLookupDefault() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.QueryLookupDefault
}

// GetQueryLookupDefaultOk returns a tuple with the QueryLookupDefault field value
// and a boolean to check if the value has been set.
func (o *FilesSharingCapabilitiesFilesSharingSharee) GetQueryLookupDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryLookupDefault, true
}

// SetQueryLookupDefault sets field value
func (o *FilesSharingCapabilitiesFilesSharingSharee) SetQueryLookupDefault(v bool) {
	o.QueryLookupDefault = v
}

// GetAlwaysShowUnique returns the AlwaysShowUnique field value
func (o *FilesSharingCapabilitiesFilesSharingSharee) GetAlwaysShowUnique() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AlwaysShowUnique
}

// GetAlwaysShowUniqueOk returns a tuple with the AlwaysShowUnique field value
// and a boolean to check if the value has been set.
func (o *FilesSharingCapabilitiesFilesSharingSharee) GetAlwaysShowUniqueOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlwaysShowUnique, true
}

// SetAlwaysShowUnique sets field value
func (o *FilesSharingCapabilitiesFilesSharingSharee) SetAlwaysShowUnique(v bool) {
	o.AlwaysShowUnique = v
}

func (o FilesSharingCapabilitiesFilesSharingSharee) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FilesSharingCapabilitiesFilesSharingSharee) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["query_lookup_default"] = o.QueryLookupDefault
	toSerialize["always_show_unique"] = o.AlwaysShowUnique

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FilesSharingCapabilitiesFilesSharingSharee) UnmarshalJSON(bytes []byte) (err error) {
	varFilesSharingCapabilitiesFilesSharingSharee := _FilesSharingCapabilitiesFilesSharingSharee{}

	if err = json.Unmarshal(bytes, &varFilesSharingCapabilitiesFilesSharingSharee); err == nil {
		*o = FilesSharingCapabilitiesFilesSharingSharee(varFilesSharingCapabilitiesFilesSharingSharee)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "query_lookup_default")
		delete(additionalProperties, "always_show_unique")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFilesSharingCapabilitiesFilesSharingSharee struct {
	value *FilesSharingCapabilitiesFilesSharingSharee
	isSet bool
}

func (v NullableFilesSharingCapabilitiesFilesSharingSharee) Get() *FilesSharingCapabilitiesFilesSharingSharee {
	return v.value
}

func (v *NullableFilesSharingCapabilitiesFilesSharingSharee) Set(val *FilesSharingCapabilitiesFilesSharingSharee) {
	v.value = val
	v.isSet = true
}

func (v NullableFilesSharingCapabilitiesFilesSharingSharee) IsSet() bool {
	return v.isSet
}

func (v *NullableFilesSharingCapabilitiesFilesSharingSharee) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilesSharingCapabilitiesFilesSharingSharee(val *FilesSharingCapabilitiesFilesSharingSharee) *NullableFilesSharingCapabilitiesFilesSharingSharee {
	return &NullableFilesSharingCapabilitiesFilesSharingSharee{value: val, isSet: true}
}

func (v NullableFilesSharingCapabilitiesFilesSharingSharee) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilesSharingCapabilitiesFilesSharingSharee) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

