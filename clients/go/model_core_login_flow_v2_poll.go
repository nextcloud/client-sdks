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

// checks if the CoreLoginFlowV2Poll type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreLoginFlowV2Poll{}

// CoreLoginFlowV2Poll struct for CoreLoginFlowV2Poll
type CoreLoginFlowV2Poll struct {
	Token string `json:"token"`
	Endpoint string `json:"endpoint"`
	AdditionalProperties map[string]interface{}
}

type _CoreLoginFlowV2Poll CoreLoginFlowV2Poll

// NewCoreLoginFlowV2Poll instantiates a new CoreLoginFlowV2Poll object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreLoginFlowV2Poll(token string, endpoint string) *CoreLoginFlowV2Poll {
	this := CoreLoginFlowV2Poll{}
	this.Token = token
	this.Endpoint = endpoint
	return &this
}

// NewCoreLoginFlowV2PollWithDefaults instantiates a new CoreLoginFlowV2Poll object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreLoginFlowV2PollWithDefaults() *CoreLoginFlowV2Poll {
	this := CoreLoginFlowV2Poll{}
	return &this
}

// GetToken returns the Token field value
func (o *CoreLoginFlowV2Poll) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *CoreLoginFlowV2Poll) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *CoreLoginFlowV2Poll) SetToken(v string) {
	o.Token = v
}

// GetEndpoint returns the Endpoint field value
func (o *CoreLoginFlowV2Poll) GetEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value
// and a boolean to check if the value has been set.
func (o *CoreLoginFlowV2Poll) GetEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Endpoint, true
}

// SetEndpoint sets field value
func (o *CoreLoginFlowV2Poll) SetEndpoint(v string) {
	o.Endpoint = v
}

func (o CoreLoginFlowV2Poll) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreLoginFlowV2Poll) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["token"] = o.Token
	toSerialize["endpoint"] = o.Endpoint

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CoreLoginFlowV2Poll) UnmarshalJSON(bytes []byte) (err error) {
	varCoreLoginFlowV2Poll := _CoreLoginFlowV2Poll{}

	if err = json.Unmarshal(bytes, &varCoreLoginFlowV2Poll); err == nil {
		*o = CoreLoginFlowV2Poll(varCoreLoginFlowV2Poll)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "token")
		delete(additionalProperties, "endpoint")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCoreLoginFlowV2Poll struct {
	value *CoreLoginFlowV2Poll
	isSet bool
}

func (v NullableCoreLoginFlowV2Poll) Get() *CoreLoginFlowV2Poll {
	return v.value
}

func (v *NullableCoreLoginFlowV2Poll) Set(val *CoreLoginFlowV2Poll) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreLoginFlowV2Poll) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreLoginFlowV2Poll) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreLoginFlowV2Poll(val *CoreLoginFlowV2Poll) *NullableCoreLoginFlowV2Poll {
	return &NullableCoreLoginFlowV2Poll{value: val, isSet: true}
}

func (v NullableCoreLoginFlowV2Poll) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreLoginFlowV2Poll) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


