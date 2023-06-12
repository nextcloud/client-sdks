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

// checks if the CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUser{}

// CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUser struct for CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUser
type CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUser struct {
	SendMail bool `json:"send_mail"`
	ExpireDate *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUserExpireDate `json:"expire_date,omitempty"`
}

// NewCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUser instantiates a new CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUser(sendMail bool) *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUser {
	this := CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUser{}
	this.SendMail = sendMail
	return &this
}

// NewCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUserWithDefaults instantiates a new CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUserWithDefaults() *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUser {
	this := CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUser{}
	return &this
}

// GetSendMail returns the SendMail field value
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUser) GetSendMail() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SendMail
}

// GetSendMailOk returns a tuple with the SendMail field value
// and a boolean to check if the value has been set.
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUser) GetSendMailOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SendMail, true
}

// SetSendMail sets field value
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUser) SetSendMail(v bool) {
	o.SendMail = v
}

// GetExpireDate returns the ExpireDate field value if set, zero value otherwise.
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUser) GetExpireDate() CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUserExpireDate {
	if o == nil || IsNil(o.ExpireDate) {
		var ret CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUserExpireDate
		return ret
	}
	return *o.ExpireDate
}

// GetExpireDateOk returns a tuple with the ExpireDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUser) GetExpireDateOk() (*CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUserExpireDate, bool) {
	if o == nil || IsNil(o.ExpireDate) {
		return nil, false
	}
	return o.ExpireDate, true
}

// HasExpireDate returns a boolean if a field has been set.
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUser) HasExpireDate() bool {
	if o != nil && !IsNil(o.ExpireDate) {
		return true
	}

	return false
}

// SetExpireDate gets a reference to the given CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUserExpireDate and assigns it to the ExpireDate field.
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUser) SetExpireDate(v CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUserExpireDate) {
	o.ExpireDate = &v
}

func (o CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["send_mail"] = o.SendMail
	if !IsNil(o.ExpireDate) {
		toSerialize["expire_date"] = o.ExpireDate
	}
	return toSerialize, nil
}

type NullableCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUser struct {
	value *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUser
	isSet bool
}

func (v NullableCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUser) Get() *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUser {
	return v.value
}

func (v *NullableCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUser) Set(val *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUser) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUser) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUser(val *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUser) *NullableCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUser {
	return &NullableCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUser{value: val, isSet: true}
}

func (v NullableCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


