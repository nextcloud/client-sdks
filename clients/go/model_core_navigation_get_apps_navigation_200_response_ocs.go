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

// checks if the CoreNavigationGetAppsNavigation200ResponseOcs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreNavigationGetAppsNavigation200ResponseOcs{}

// CoreNavigationGetAppsNavigation200ResponseOcs struct for CoreNavigationGetAppsNavigation200ResponseOcs
type CoreNavigationGetAppsNavigation200ResponseOcs struct {
	Meta OCSMeta `json:"meta"`
	Data []CoreNavigationEntry `json:"data"`
	AdditionalProperties map[string]interface{}
}

type _CoreNavigationGetAppsNavigation200ResponseOcs CoreNavigationGetAppsNavigation200ResponseOcs

// NewCoreNavigationGetAppsNavigation200ResponseOcs instantiates a new CoreNavigationGetAppsNavigation200ResponseOcs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreNavigationGetAppsNavigation200ResponseOcs(meta OCSMeta, data []CoreNavigationEntry) *CoreNavigationGetAppsNavigation200ResponseOcs {
	this := CoreNavigationGetAppsNavigation200ResponseOcs{}
	this.Meta = meta
	this.Data = data
	return &this
}

// NewCoreNavigationGetAppsNavigation200ResponseOcsWithDefaults instantiates a new CoreNavigationGetAppsNavigation200ResponseOcs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreNavigationGetAppsNavigation200ResponseOcsWithDefaults() *CoreNavigationGetAppsNavigation200ResponseOcs {
	this := CoreNavigationGetAppsNavigation200ResponseOcs{}
	return &this
}

// GetMeta returns the Meta field value
func (o *CoreNavigationGetAppsNavigation200ResponseOcs) GetMeta() OCSMeta {
	if o == nil {
		var ret OCSMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *CoreNavigationGetAppsNavigation200ResponseOcs) GetMetaOk() (*OCSMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *CoreNavigationGetAppsNavigation200ResponseOcs) SetMeta(v OCSMeta) {
	o.Meta = v
}

// GetData returns the Data field value
func (o *CoreNavigationGetAppsNavigation200ResponseOcs) GetData() []CoreNavigationEntry {
	if o == nil {
		var ret []CoreNavigationEntry
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CoreNavigationGetAppsNavigation200ResponseOcs) GetDataOk() ([]CoreNavigationEntry, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *CoreNavigationGetAppsNavigation200ResponseOcs) SetData(v []CoreNavigationEntry) {
	o.Data = v
}

func (o CoreNavigationGetAppsNavigation200ResponseOcs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreNavigationGetAppsNavigation200ResponseOcs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meta"] = o.Meta
	toSerialize["data"] = o.Data

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CoreNavigationGetAppsNavigation200ResponseOcs) UnmarshalJSON(bytes []byte) (err error) {
	varCoreNavigationGetAppsNavigation200ResponseOcs := _CoreNavigationGetAppsNavigation200ResponseOcs{}

	if err = json.Unmarshal(bytes, &varCoreNavigationGetAppsNavigation200ResponseOcs); err == nil {
		*o = CoreNavigationGetAppsNavigation200ResponseOcs(varCoreNavigationGetAppsNavigation200ResponseOcs)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "meta")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCoreNavigationGetAppsNavigation200ResponseOcs struct {
	value *CoreNavigationGetAppsNavigation200ResponseOcs
	isSet bool
}

func (v NullableCoreNavigationGetAppsNavigation200ResponseOcs) Get() *CoreNavigationGetAppsNavigation200ResponseOcs {
	return v.value
}

func (v *NullableCoreNavigationGetAppsNavigation200ResponseOcs) Set(val *CoreNavigationGetAppsNavigation200ResponseOcs) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreNavigationGetAppsNavigation200ResponseOcs) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreNavigationGetAppsNavigation200ResponseOcs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreNavigationGetAppsNavigation200ResponseOcs(val *CoreNavigationGetAppsNavigation200ResponseOcs) *NullableCoreNavigationGetAppsNavigation200ResponseOcs {
	return &NullableCoreNavigationGetAppsNavigation200ResponseOcs{value: val, isSet: true}
}

func (v NullableCoreNavigationGetAppsNavigation200ResponseOcs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreNavigationGetAppsNavigation200ResponseOcs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

