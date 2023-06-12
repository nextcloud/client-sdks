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

// checks if the CoreLoginFlowV2Credentials type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreLoginFlowV2Credentials{}

// CoreLoginFlowV2Credentials struct for CoreLoginFlowV2Credentials
type CoreLoginFlowV2Credentials struct {
	Server string `json:"server"`
	LoginName string `json:"loginName"`
	AppPassword string `json:"appPassword"`
}

// NewCoreLoginFlowV2Credentials instantiates a new CoreLoginFlowV2Credentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreLoginFlowV2Credentials(server string, loginName string, appPassword string) *CoreLoginFlowV2Credentials {
	this := CoreLoginFlowV2Credentials{}
	this.Server = server
	this.LoginName = loginName
	this.AppPassword = appPassword
	return &this
}

// NewCoreLoginFlowV2CredentialsWithDefaults instantiates a new CoreLoginFlowV2Credentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreLoginFlowV2CredentialsWithDefaults() *CoreLoginFlowV2Credentials {
	this := CoreLoginFlowV2Credentials{}
	return &this
}

// GetServer returns the Server field value
func (o *CoreLoginFlowV2Credentials) GetServer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Server
}

// GetServerOk returns a tuple with the Server field value
// and a boolean to check if the value has been set.
func (o *CoreLoginFlowV2Credentials) GetServerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Server, true
}

// SetServer sets field value
func (o *CoreLoginFlowV2Credentials) SetServer(v string) {
	o.Server = v
}

// GetLoginName returns the LoginName field value
func (o *CoreLoginFlowV2Credentials) GetLoginName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LoginName
}

// GetLoginNameOk returns a tuple with the LoginName field value
// and a boolean to check if the value has been set.
func (o *CoreLoginFlowV2Credentials) GetLoginNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LoginName, true
}

// SetLoginName sets field value
func (o *CoreLoginFlowV2Credentials) SetLoginName(v string) {
	o.LoginName = v
}

// GetAppPassword returns the AppPassword field value
func (o *CoreLoginFlowV2Credentials) GetAppPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppPassword
}

// GetAppPasswordOk returns a tuple with the AppPassword field value
// and a boolean to check if the value has been set.
func (o *CoreLoginFlowV2Credentials) GetAppPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppPassword, true
}

// SetAppPassword sets field value
func (o *CoreLoginFlowV2Credentials) SetAppPassword(v string) {
	o.AppPassword = v
}

func (o CoreLoginFlowV2Credentials) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreLoginFlowV2Credentials) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["server"] = o.Server
	toSerialize["loginName"] = o.LoginName
	toSerialize["appPassword"] = o.AppPassword
	return toSerialize, nil
}

type NullableCoreLoginFlowV2Credentials struct {
	value *CoreLoginFlowV2Credentials
	isSet bool
}

func (v NullableCoreLoginFlowV2Credentials) Get() *CoreLoginFlowV2Credentials {
	return v.value
}

func (v *NullableCoreLoginFlowV2Credentials) Set(val *CoreLoginFlowV2Credentials) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreLoginFlowV2Credentials) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreLoginFlowV2Credentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreLoginFlowV2Credentials(val *CoreLoginFlowV2Credentials) *NullableCoreLoginFlowV2Credentials {
	return &NullableCoreLoginFlowV2Credentials{value: val, isSet: true}
}

func (v NullableCoreLoginFlowV2Credentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreLoginFlowV2Credentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


