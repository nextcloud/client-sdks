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

// checks if the UserStatusPrivateAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserStatusPrivateAllOf{}

// UserStatusPrivateAllOf struct for UserStatusPrivateAllOf
type UserStatusPrivateAllOf struct {
	MessageId NullableString `json:"messageId"`
	MessageIsPredefined bool `json:"messageIsPredefined"`
	StatusIsUserDefined bool `json:"statusIsUserDefined"`
}

// NewUserStatusPrivateAllOf instantiates a new UserStatusPrivateAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserStatusPrivateAllOf(messageId NullableString, messageIsPredefined bool, statusIsUserDefined bool) *UserStatusPrivateAllOf {
	this := UserStatusPrivateAllOf{}
	this.MessageId = messageId
	this.MessageIsPredefined = messageIsPredefined
	this.StatusIsUserDefined = statusIsUserDefined
	return &this
}

// NewUserStatusPrivateAllOfWithDefaults instantiates a new UserStatusPrivateAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserStatusPrivateAllOfWithDefaults() *UserStatusPrivateAllOf {
	this := UserStatusPrivateAllOf{}
	return &this
}

// GetMessageId returns the MessageId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UserStatusPrivateAllOf) GetMessageId() string {
	if o == nil || o.MessageId.Get() == nil {
		var ret string
		return ret
	}

	return *o.MessageId.Get()
}

// GetMessageIdOk returns a tuple with the MessageId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserStatusPrivateAllOf) GetMessageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MessageId.Get(), o.MessageId.IsSet()
}

// SetMessageId sets field value
func (o *UserStatusPrivateAllOf) SetMessageId(v string) {
	o.MessageId.Set(&v)
}

// GetMessageIsPredefined returns the MessageIsPredefined field value
func (o *UserStatusPrivateAllOf) GetMessageIsPredefined() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.MessageIsPredefined
}

// GetMessageIsPredefinedOk returns a tuple with the MessageIsPredefined field value
// and a boolean to check if the value has been set.
func (o *UserStatusPrivateAllOf) GetMessageIsPredefinedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageIsPredefined, true
}

// SetMessageIsPredefined sets field value
func (o *UserStatusPrivateAllOf) SetMessageIsPredefined(v bool) {
	o.MessageIsPredefined = v
}

// GetStatusIsUserDefined returns the StatusIsUserDefined field value
func (o *UserStatusPrivateAllOf) GetStatusIsUserDefined() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.StatusIsUserDefined
}

// GetStatusIsUserDefinedOk returns a tuple with the StatusIsUserDefined field value
// and a boolean to check if the value has been set.
func (o *UserStatusPrivateAllOf) GetStatusIsUserDefinedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusIsUserDefined, true
}

// SetStatusIsUserDefined sets field value
func (o *UserStatusPrivateAllOf) SetStatusIsUserDefined(v bool) {
	o.StatusIsUserDefined = v
}

func (o UserStatusPrivateAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserStatusPrivateAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["messageId"] = o.MessageId.Get()
	toSerialize["messageIsPredefined"] = o.MessageIsPredefined
	toSerialize["statusIsUserDefined"] = o.StatusIsUserDefined
	return toSerialize, nil
}

type NullableUserStatusPrivateAllOf struct {
	value *UserStatusPrivateAllOf
	isSet bool
}

func (v NullableUserStatusPrivateAllOf) Get() *UserStatusPrivateAllOf {
	return v.value
}

func (v *NullableUserStatusPrivateAllOf) Set(val *UserStatusPrivateAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUserStatusPrivateAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUserStatusPrivateAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserStatusPrivateAllOf(val *UserStatusPrivateAllOf) *NullableUserStatusPrivateAllOf {
	return &NullableUserStatusPrivateAllOf{value: val, isSet: true}
}

func (v NullableUserStatusPrivateAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserStatusPrivateAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

