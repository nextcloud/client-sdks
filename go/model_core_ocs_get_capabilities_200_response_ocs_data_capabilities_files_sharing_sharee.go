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

// checks if the CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingSharee type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingSharee{}

// CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingSharee struct for CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingSharee
type CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingSharee struct {
	QueryLookupDefault bool `json:"query_lookup_default"`
	AlwaysShowUnique bool `json:"always_show_unique"`
}

// NewCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingSharee instantiates a new CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingSharee object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingSharee(queryLookupDefault bool, alwaysShowUnique bool) *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingSharee {
	this := CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingSharee{}
	this.QueryLookupDefault = queryLookupDefault
	this.AlwaysShowUnique = alwaysShowUnique
	return &this
}

// NewCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingShareeWithDefaults instantiates a new CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingSharee object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingShareeWithDefaults() *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingSharee {
	this := CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingSharee{}
	return &this
}

// GetQueryLookupDefault returns the QueryLookupDefault field value
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingSharee) GetQueryLookupDefault() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.QueryLookupDefault
}

// GetQueryLookupDefaultOk returns a tuple with the QueryLookupDefault field value
// and a boolean to check if the value has been set.
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingSharee) GetQueryLookupDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryLookupDefault, true
}

// SetQueryLookupDefault sets field value
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingSharee) SetQueryLookupDefault(v bool) {
	o.QueryLookupDefault = v
}

// GetAlwaysShowUnique returns the AlwaysShowUnique field value
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingSharee) GetAlwaysShowUnique() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AlwaysShowUnique
}

// GetAlwaysShowUniqueOk returns a tuple with the AlwaysShowUnique field value
// and a boolean to check if the value has been set.
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingSharee) GetAlwaysShowUniqueOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlwaysShowUnique, true
}

// SetAlwaysShowUnique sets field value
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingSharee) SetAlwaysShowUnique(v bool) {
	o.AlwaysShowUnique = v
}

func (o CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingSharee) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingSharee) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["query_lookup_default"] = o.QueryLookupDefault
	toSerialize["always_show_unique"] = o.AlwaysShowUnique
	return toSerialize, nil
}

type NullableCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingSharee struct {
	value *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingSharee
	isSet bool
}

func (v NullableCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingSharee) Get() *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingSharee {
	return v.value
}

func (v *NullableCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingSharee) Set(val *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingSharee) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingSharee) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingSharee) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingSharee(val *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingSharee) *NullableCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingSharee {
	return &NullableCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingSharee{value: val, isSet: true}
}

func (v NullableCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingSharee) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingSharee) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


