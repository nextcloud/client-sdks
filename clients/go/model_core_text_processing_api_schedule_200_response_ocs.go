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

// checks if the CoreTextProcessingApiSchedule200ResponseOcs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreTextProcessingApiSchedule200ResponseOcs{}

// CoreTextProcessingApiSchedule200ResponseOcs struct for CoreTextProcessingApiSchedule200ResponseOcs
type CoreTextProcessingApiSchedule200ResponseOcs struct {
	Meta OCSMeta `json:"meta"`
	Data CoreTextProcessingApiSchedule200ResponseOcsData `json:"data"`
	AdditionalProperties map[string]interface{}
}

type _CoreTextProcessingApiSchedule200ResponseOcs CoreTextProcessingApiSchedule200ResponseOcs

// NewCoreTextProcessingApiSchedule200ResponseOcs instantiates a new CoreTextProcessingApiSchedule200ResponseOcs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreTextProcessingApiSchedule200ResponseOcs(meta OCSMeta, data CoreTextProcessingApiSchedule200ResponseOcsData) *CoreTextProcessingApiSchedule200ResponseOcs {
	this := CoreTextProcessingApiSchedule200ResponseOcs{}
	this.Meta = meta
	this.Data = data
	return &this
}

// NewCoreTextProcessingApiSchedule200ResponseOcsWithDefaults instantiates a new CoreTextProcessingApiSchedule200ResponseOcs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreTextProcessingApiSchedule200ResponseOcsWithDefaults() *CoreTextProcessingApiSchedule200ResponseOcs {
	this := CoreTextProcessingApiSchedule200ResponseOcs{}
	return &this
}

// GetMeta returns the Meta field value
func (o *CoreTextProcessingApiSchedule200ResponseOcs) GetMeta() OCSMeta {
	if o == nil {
		var ret OCSMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *CoreTextProcessingApiSchedule200ResponseOcs) GetMetaOk() (*OCSMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *CoreTextProcessingApiSchedule200ResponseOcs) SetMeta(v OCSMeta) {
	o.Meta = v
}

// GetData returns the Data field value
func (o *CoreTextProcessingApiSchedule200ResponseOcs) GetData() CoreTextProcessingApiSchedule200ResponseOcsData {
	if o == nil {
		var ret CoreTextProcessingApiSchedule200ResponseOcsData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CoreTextProcessingApiSchedule200ResponseOcs) GetDataOk() (*CoreTextProcessingApiSchedule200ResponseOcsData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CoreTextProcessingApiSchedule200ResponseOcs) SetData(v CoreTextProcessingApiSchedule200ResponseOcsData) {
	o.Data = v
}

func (o CoreTextProcessingApiSchedule200ResponseOcs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreTextProcessingApiSchedule200ResponseOcs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["data"] = o.Data

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CoreTextProcessingApiSchedule200ResponseOcs) UnmarshalJSON(bytes []byte) (err error) {
	varCoreTextProcessingApiSchedule200ResponseOcs := _CoreTextProcessingApiSchedule200ResponseOcs{}

	if err = json.Unmarshal(bytes, &varCoreTextProcessingApiSchedule200ResponseOcs); err == nil {
		*o = CoreTextProcessingApiSchedule200ResponseOcs(varCoreTextProcessingApiSchedule200ResponseOcs)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "meta")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCoreTextProcessingApiSchedule200ResponseOcs struct {
	value *CoreTextProcessingApiSchedule200ResponseOcs
	isSet bool
}

func (v NullableCoreTextProcessingApiSchedule200ResponseOcs) Get() *CoreTextProcessingApiSchedule200ResponseOcs {
	return v.value
}

func (v *NullableCoreTextProcessingApiSchedule200ResponseOcs) Set(val *CoreTextProcessingApiSchedule200ResponseOcs) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreTextProcessingApiSchedule200ResponseOcs) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreTextProcessingApiSchedule200ResponseOcs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreTextProcessingApiSchedule200ResponseOcs(val *CoreTextProcessingApiSchedule200ResponseOcs) *NullableCoreTextProcessingApiSchedule200ResponseOcs {
	return &NullableCoreTextProcessingApiSchedule200ResponseOcs{value: val, isSet: true}
}

func (v NullableCoreTextProcessingApiSchedule200ResponseOcs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreTextProcessingApiSchedule200ResponseOcs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


