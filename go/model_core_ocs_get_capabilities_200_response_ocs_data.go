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

// checks if the CoreOcsGetCapabilities200ResponseOcsData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreOcsGetCapabilities200ResponseOcsData{}

// CoreOcsGetCapabilities200ResponseOcsData struct for CoreOcsGetCapabilities200ResponseOcsData
type CoreOcsGetCapabilities200ResponseOcsData struct {
	Version CoreOcsGetCapabilities200ResponseOcsDataVersion `json:"version"`
	Capabilities CoreOcsGetCapabilities200ResponseOcsDataCapabilities `json:"capabilities"`
}

// NewCoreOcsGetCapabilities200ResponseOcsData instantiates a new CoreOcsGetCapabilities200ResponseOcsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreOcsGetCapabilities200ResponseOcsData(version CoreOcsGetCapabilities200ResponseOcsDataVersion, capabilities CoreOcsGetCapabilities200ResponseOcsDataCapabilities) *CoreOcsGetCapabilities200ResponseOcsData {
	this := CoreOcsGetCapabilities200ResponseOcsData{}
	this.Version = version
	this.Capabilities = capabilities
	return &this
}

// NewCoreOcsGetCapabilities200ResponseOcsDataWithDefaults instantiates a new CoreOcsGetCapabilities200ResponseOcsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreOcsGetCapabilities200ResponseOcsDataWithDefaults() *CoreOcsGetCapabilities200ResponseOcsData {
	this := CoreOcsGetCapabilities200ResponseOcsData{}
	return &this
}

// GetVersion returns the Version field value
func (o *CoreOcsGetCapabilities200ResponseOcsData) GetVersion() CoreOcsGetCapabilities200ResponseOcsDataVersion {
	if o == nil {
		var ret CoreOcsGetCapabilities200ResponseOcsDataVersion
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *CoreOcsGetCapabilities200ResponseOcsData) GetVersionOk() (*CoreOcsGetCapabilities200ResponseOcsDataVersion, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *CoreOcsGetCapabilities200ResponseOcsData) SetVersion(v CoreOcsGetCapabilities200ResponseOcsDataVersion) {
	o.Version = v
}

// GetCapabilities returns the Capabilities field value
func (o *CoreOcsGetCapabilities200ResponseOcsData) GetCapabilities() CoreOcsGetCapabilities200ResponseOcsDataCapabilities {
	if o == nil {
		var ret CoreOcsGetCapabilities200ResponseOcsDataCapabilities
		return ret
	}

	return o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value
// and a boolean to check if the value has been set.
func (o *CoreOcsGetCapabilities200ResponseOcsData) GetCapabilitiesOk() (*CoreOcsGetCapabilities200ResponseOcsDataCapabilities, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Capabilities, true
}

// SetCapabilities sets field value
func (o *CoreOcsGetCapabilities200ResponseOcsData) SetCapabilities(v CoreOcsGetCapabilities200ResponseOcsDataCapabilities) {
	o.Capabilities = v
}

func (o CoreOcsGetCapabilities200ResponseOcsData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreOcsGetCapabilities200ResponseOcsData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["version"] = o.Version
	toSerialize["capabilities"] = o.Capabilities
	return toSerialize, nil
}

type NullableCoreOcsGetCapabilities200ResponseOcsData struct {
	value *CoreOcsGetCapabilities200ResponseOcsData
	isSet bool
}

func (v NullableCoreOcsGetCapabilities200ResponseOcsData) Get() *CoreOcsGetCapabilities200ResponseOcsData {
	return v.value
}

func (v *NullableCoreOcsGetCapabilities200ResponseOcsData) Set(val *CoreOcsGetCapabilities200ResponseOcsData) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreOcsGetCapabilities200ResponseOcsData) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreOcsGetCapabilities200ResponseOcsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreOcsGetCapabilities200ResponseOcsData(val *CoreOcsGetCapabilities200ResponseOcsData) *NullableCoreOcsGetCapabilities200ResponseOcsData {
	return &NullableCoreOcsGetCapabilities200ResponseOcsData{value: val, isSet: true}
}

func (v NullableCoreOcsGetCapabilities200ResponseOcsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreOcsGetCapabilities200ResponseOcsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


