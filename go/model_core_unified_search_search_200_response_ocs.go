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

// checks if the CoreUnifiedSearchSearch200ResponseOcs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreUnifiedSearchSearch200ResponseOcs{}

// CoreUnifiedSearchSearch200ResponseOcs struct for CoreUnifiedSearchSearch200ResponseOcs
type CoreUnifiedSearchSearch200ResponseOcs struct {
	Meta OCSMeta `json:"meta"`
	Data CoreUnifiedSearchResult `json:"data"`
}

// NewCoreUnifiedSearchSearch200ResponseOcs instantiates a new CoreUnifiedSearchSearch200ResponseOcs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreUnifiedSearchSearch200ResponseOcs(meta OCSMeta, data CoreUnifiedSearchResult) *CoreUnifiedSearchSearch200ResponseOcs {
	this := CoreUnifiedSearchSearch200ResponseOcs{}
	this.Meta = meta
	this.Data = data
	return &this
}

// NewCoreUnifiedSearchSearch200ResponseOcsWithDefaults instantiates a new CoreUnifiedSearchSearch200ResponseOcs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreUnifiedSearchSearch200ResponseOcsWithDefaults() *CoreUnifiedSearchSearch200ResponseOcs {
	this := CoreUnifiedSearchSearch200ResponseOcs{}
	return &this
}

// GetMeta returns the Meta field value
func (o *CoreUnifiedSearchSearch200ResponseOcs) GetMeta() OCSMeta {
	if o == nil {
		var ret OCSMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *CoreUnifiedSearchSearch200ResponseOcs) GetMetaOk() (*OCSMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *CoreUnifiedSearchSearch200ResponseOcs) SetMeta(v OCSMeta) {
	o.Meta = v
}

// GetData returns the Data field value
func (o *CoreUnifiedSearchSearch200ResponseOcs) GetData() CoreUnifiedSearchResult {
	if o == nil {
		var ret CoreUnifiedSearchResult
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CoreUnifiedSearchSearch200ResponseOcs) GetDataOk() (*CoreUnifiedSearchResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CoreUnifiedSearchSearch200ResponseOcs) SetData(v CoreUnifiedSearchResult) {
	o.Data = v
}

func (o CoreUnifiedSearchSearch200ResponseOcs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreUnifiedSearchSearch200ResponseOcs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableCoreUnifiedSearchSearch200ResponseOcs struct {
	value *CoreUnifiedSearchSearch200ResponseOcs
	isSet bool
}

func (v NullableCoreUnifiedSearchSearch200ResponseOcs) Get() *CoreUnifiedSearchSearch200ResponseOcs {
	return v.value
}

func (v *NullableCoreUnifiedSearchSearch200ResponseOcs) Set(val *CoreUnifiedSearchSearch200ResponseOcs) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreUnifiedSearchSearch200ResponseOcs) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreUnifiedSearchSearch200ResponseOcs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreUnifiedSearchSearch200ResponseOcs(val *CoreUnifiedSearchSearch200ResponseOcs) *NullableCoreUnifiedSearchSearch200ResponseOcs {
	return &NullableCoreUnifiedSearchSearch200ResponseOcs{value: val, isSet: true}
}

func (v NullableCoreUnifiedSearchSearch200ResponseOcs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreUnifiedSearchSearch200ResponseOcs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


