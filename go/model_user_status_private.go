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

// checks if the UserStatusPrivate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserStatusPrivate{}

// UserStatusPrivate struct for UserStatusPrivate
type UserStatusPrivate struct {
	UserId string `json:"userId"`
	Message NullableString `json:"message"`
	Icon NullableString `json:"icon"`
	ClearAt NullableInt64 `json:"clearAt"`
	Status string `json:"status"`
	MessageId NullableString `json:"messageId"`
	MessageIsPredefined bool `json:"messageIsPredefined"`
	StatusIsUserDefined bool `json:"statusIsUserDefined"`
}

// NewUserStatusPrivate instantiates a new UserStatusPrivate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserStatusPrivate(userId string, message NullableString, icon NullableString, clearAt NullableInt64, status string, messageId NullableString, messageIsPredefined bool, statusIsUserDefined bool) *UserStatusPrivate {
	this := UserStatusPrivate{}
	this.UserId = userId
	this.Message = message
	this.Icon = icon
	this.ClearAt = clearAt
	this.Status = status
	this.MessageId = messageId
	this.MessageIsPredefined = messageIsPredefined
	this.StatusIsUserDefined = statusIsUserDefined
	return &this
}

// NewUserStatusPrivateWithDefaults instantiates a new UserStatusPrivate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserStatusPrivateWithDefaults() *UserStatusPrivate {
	this := UserStatusPrivate{}
	return &this
}

// GetUserId returns the UserId field value
func (o *UserStatusPrivate) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *UserStatusPrivate) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *UserStatusPrivate) SetUserId(v string) {
	o.UserId = v
}

// GetMessage returns the Message field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UserStatusPrivate) GetMessage() string {
	if o == nil || o.Message.Get() == nil {
		var ret string
		return ret
	}

	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserStatusPrivate) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// SetMessage sets field value
func (o *UserStatusPrivate) SetMessage(v string) {
	o.Message.Set(&v)
}

// GetIcon returns the Icon field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UserStatusPrivate) GetIcon() string {
	if o == nil || o.Icon.Get() == nil {
		var ret string
		return ret
	}

	return *o.Icon.Get()
}

// GetIconOk returns a tuple with the Icon field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserStatusPrivate) GetIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Icon.Get(), o.Icon.IsSet()
}

// SetIcon sets field value
func (o *UserStatusPrivate) SetIcon(v string) {
	o.Icon.Set(&v)
}

// GetClearAt returns the ClearAt field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *UserStatusPrivate) GetClearAt() int64 {
	if o == nil || o.ClearAt.Get() == nil {
		var ret int64
		return ret
	}

	return *o.ClearAt.Get()
}

// GetClearAtOk returns a tuple with the ClearAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserStatusPrivate) GetClearAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClearAt.Get(), o.ClearAt.IsSet()
}

// SetClearAt sets field value
func (o *UserStatusPrivate) SetClearAt(v int64) {
	o.ClearAt.Set(&v)
}

// GetStatus returns the Status field value
func (o *UserStatusPrivate) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *UserStatusPrivate) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *UserStatusPrivate) SetStatus(v string) {
	o.Status = v
}

// GetMessageId returns the MessageId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UserStatusPrivate) GetMessageId() string {
	if o == nil || o.MessageId.Get() == nil {
		var ret string
		return ret
	}

	return *o.MessageId.Get()
}

// GetMessageIdOk returns a tuple with the MessageId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserStatusPrivate) GetMessageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MessageId.Get(), o.MessageId.IsSet()
}

// SetMessageId sets field value
func (o *UserStatusPrivate) SetMessageId(v string) {
	o.MessageId.Set(&v)
}

// GetMessageIsPredefined returns the MessageIsPredefined field value
func (o *UserStatusPrivate) GetMessageIsPredefined() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.MessageIsPredefined
}

// GetMessageIsPredefinedOk returns a tuple with the MessageIsPredefined field value
// and a boolean to check if the value has been set.
func (o *UserStatusPrivate) GetMessageIsPredefinedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageIsPredefined, true
}

// SetMessageIsPredefined sets field value
func (o *UserStatusPrivate) SetMessageIsPredefined(v bool) {
	o.MessageIsPredefined = v
}

// GetStatusIsUserDefined returns the StatusIsUserDefined field value
func (o *UserStatusPrivate) GetStatusIsUserDefined() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.StatusIsUserDefined
}

// GetStatusIsUserDefinedOk returns a tuple with the StatusIsUserDefined field value
// and a boolean to check if the value has been set.
func (o *UserStatusPrivate) GetStatusIsUserDefinedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusIsUserDefined, true
}

// SetStatusIsUserDefined sets field value
func (o *UserStatusPrivate) SetStatusIsUserDefined(v bool) {
	o.StatusIsUserDefined = v
}

func (o UserStatusPrivate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserStatusPrivate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["userId"] = o.UserId
	toSerialize["message"] = o.Message.Get()
	toSerialize["icon"] = o.Icon.Get()
	toSerialize["clearAt"] = o.ClearAt.Get()
	toSerialize["status"] = o.Status
	toSerialize["messageId"] = o.MessageId.Get()
	toSerialize["messageIsPredefined"] = o.MessageIsPredefined
	toSerialize["statusIsUserDefined"] = o.StatusIsUserDefined
	return toSerialize, nil
}

type NullableUserStatusPrivate struct {
	value *UserStatusPrivate
	isSet bool
}

func (v NullableUserStatusPrivate) Get() *UserStatusPrivate {
	return v.value
}

func (v *NullableUserStatusPrivate) Set(val *UserStatusPrivate) {
	v.value = val
	v.isSet = true
}

func (v NullableUserStatusPrivate) IsSet() bool {
	return v.isSet
}

func (v *NullableUserStatusPrivate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserStatusPrivate(val *UserStatusPrivate) *NullableUserStatusPrivate {
	return &NullableUserStatusPrivate{value: val, isSet: true}
}

func (v NullableUserStatusPrivate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserStatusPrivate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


