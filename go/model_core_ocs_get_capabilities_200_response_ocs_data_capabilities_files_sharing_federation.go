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

// checks if the CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingFederation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingFederation{}

// CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingFederation struct for CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingFederation
type CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingFederation struct {
	Outgoing bool `json:"outgoing"`
	Incoming bool `json:"incoming"`
	ExpireDate CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUserExpireDate `json:"expire_date"`
	ExpireDateSupported CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUserExpireDate `json:"expire_date_supported"`
}

// NewCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingFederation instantiates a new CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingFederation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingFederation(outgoing bool, incoming bool, expireDate CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUserExpireDate, expireDateSupported CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUserExpireDate) *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingFederation {
	this := CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingFederation{}
	this.Outgoing = outgoing
	this.Incoming = incoming
	this.ExpireDate = expireDate
	this.ExpireDateSupported = expireDateSupported
	return &this
}

// NewCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingFederationWithDefaults instantiates a new CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingFederation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingFederationWithDefaults() *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingFederation {
	this := CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingFederation{}
	return &this
}

// GetOutgoing returns the Outgoing field value
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingFederation) GetOutgoing() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Outgoing
}

// GetOutgoingOk returns a tuple with the Outgoing field value
// and a boolean to check if the value has been set.
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingFederation) GetOutgoingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Outgoing, true
}

// SetOutgoing sets field value
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingFederation) SetOutgoing(v bool) {
	o.Outgoing = v
}

// GetIncoming returns the Incoming field value
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingFederation) GetIncoming() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Incoming
}

// GetIncomingOk returns a tuple with the Incoming field value
// and a boolean to check if the value has been set.
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingFederation) GetIncomingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Incoming, true
}

// SetIncoming sets field value
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingFederation) SetIncoming(v bool) {
	o.Incoming = v
}

// GetExpireDate returns the ExpireDate field value
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingFederation) GetExpireDate() CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUserExpireDate {
	if o == nil {
		var ret CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUserExpireDate
		return ret
	}

	return o.ExpireDate
}

// GetExpireDateOk returns a tuple with the ExpireDate field value
// and a boolean to check if the value has been set.
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingFederation) GetExpireDateOk() (*CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUserExpireDate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpireDate, true
}

// SetExpireDate sets field value
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingFederation) SetExpireDate(v CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUserExpireDate) {
	o.ExpireDate = v
}

// GetExpireDateSupported returns the ExpireDateSupported field value
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingFederation) GetExpireDateSupported() CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUserExpireDate {
	if o == nil {
		var ret CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUserExpireDate
		return ret
	}

	return o.ExpireDateSupported
}

// GetExpireDateSupportedOk returns a tuple with the ExpireDateSupported field value
// and a boolean to check if the value has been set.
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingFederation) GetExpireDateSupportedOk() (*CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUserExpireDate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpireDateSupported, true
}

// SetExpireDateSupported sets field value
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingFederation) SetExpireDateSupported(v CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUserExpireDate) {
	o.ExpireDateSupported = v
}

func (o CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingFederation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingFederation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["outgoing"] = o.Outgoing
	toSerialize["incoming"] = o.Incoming
	toSerialize["expire_date"] = o.ExpireDate
	toSerialize["expire_date_supported"] = o.ExpireDateSupported
	return toSerialize, nil
}

type NullableCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingFederation struct {
	value *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingFederation
	isSet bool
}

func (v NullableCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingFederation) Get() *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingFederation {
	return v.value
}

func (v *NullableCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingFederation) Set(val *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingFederation) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingFederation) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingFederation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingFederation(val *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingFederation) *NullableCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingFederation {
	return &NullableCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingFederation{value: val, isSet: true}
}

func (v NullableCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingFederation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingFederation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


