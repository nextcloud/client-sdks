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

// checks if the CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcmResourceTypesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcmResourceTypesInner{}

// CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcmResourceTypesInner struct for CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcmResourceTypesInner
type CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcmResourceTypesInner struct {
	Name string `json:"name"`
	ShareTypes []string `json:"shareTypes"`
	Protocols CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcmResourceTypesInnerProtocols `json:"protocols"`
}

// NewCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcmResourceTypesInner instantiates a new CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcmResourceTypesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcmResourceTypesInner(name string, shareTypes []string, protocols CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcmResourceTypesInnerProtocols) *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcmResourceTypesInner {
	this := CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcmResourceTypesInner{}
	this.Name = name
	this.ShareTypes = shareTypes
	this.Protocols = protocols
	return &this
}

// NewCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcmResourceTypesInnerWithDefaults instantiates a new CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcmResourceTypesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcmResourceTypesInnerWithDefaults() *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcmResourceTypesInner {
	this := CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcmResourceTypesInner{}
	return &this
}

// GetName returns the Name field value
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcmResourceTypesInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcmResourceTypesInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcmResourceTypesInner) SetName(v string) {
	o.Name = v
}

// GetShareTypes returns the ShareTypes field value
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcmResourceTypesInner) GetShareTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ShareTypes
}

// GetShareTypesOk returns a tuple with the ShareTypes field value
// and a boolean to check if the value has been set.
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcmResourceTypesInner) GetShareTypesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShareTypes, true
}

// SetShareTypes sets field value
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcmResourceTypesInner) SetShareTypes(v []string) {
	o.ShareTypes = v
}

// GetProtocols returns the Protocols field value
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcmResourceTypesInner) GetProtocols() CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcmResourceTypesInnerProtocols {
	if o == nil {
		var ret CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcmResourceTypesInnerProtocols
		return ret
	}

	return o.Protocols
}

// GetProtocolsOk returns a tuple with the Protocols field value
// and a boolean to check if the value has been set.
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcmResourceTypesInner) GetProtocolsOk() (*CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcmResourceTypesInnerProtocols, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Protocols, true
}

// SetProtocols sets field value
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcmResourceTypesInner) SetProtocols(v CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcmResourceTypesInnerProtocols) {
	o.Protocols = v
}

func (o CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcmResourceTypesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcmResourceTypesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["shareTypes"] = o.ShareTypes
	toSerialize["protocols"] = o.Protocols
	return toSerialize, nil
}

type NullableCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcmResourceTypesInner struct {
	value *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcmResourceTypesInner
	isSet bool
}

func (v NullableCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcmResourceTypesInner) Get() *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcmResourceTypesInner {
	return v.value
}

func (v *NullableCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcmResourceTypesInner) Set(val *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcmResourceTypesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcmResourceTypesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcmResourceTypesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcmResourceTypesInner(val *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcmResourceTypesInner) *NullableCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcmResourceTypesInner {
	return &NullableCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcmResourceTypesInner{value: val, isSet: true}
}

func (v NullableCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcmResourceTypesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcmResourceTypesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


