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

// checks if the UserStatusPublic type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserStatusPublic{}

// UserStatusPublic struct for UserStatusPublic
type UserStatusPublic struct {
	UserId string `json:"userId"`
	Message NullableString `json:"message"`
	Icon NullableString `json:"icon"`
	ClearAt NullableInt64 `json:"clearAt"`
	Status string `json:"status"`
}

// NewUserStatusPublic instantiates a new UserStatusPublic object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserStatusPublic(userId string, message NullableString, icon NullableString, clearAt NullableInt64, status string) *UserStatusPublic {
	this := UserStatusPublic{}
	this.UserId = userId
	this.Message = message
	this.Icon = icon
	this.ClearAt = clearAt
	this.Status = status
	return &this
}

// NewUserStatusPublicWithDefaults instantiates a new UserStatusPublic object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserStatusPublicWithDefaults() *UserStatusPublic {
	this := UserStatusPublic{}
	return &this
}

// GetUserId returns the UserId field value
func (o *UserStatusPublic) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *UserStatusPublic) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *UserStatusPublic) SetUserId(v string) {
	o.UserId = v
}

// GetMessage returns the Message field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UserStatusPublic) GetMessage() string {
	if o == nil || o.Message.Get() == nil {
		var ret string
		return ret
	}

	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserStatusPublic) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// SetMessage sets field value
func (o *UserStatusPublic) SetMessage(v string) {
	o.Message.Set(&v)
}

// GetIcon returns the Icon field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UserStatusPublic) GetIcon() string {
	if o == nil || o.Icon.Get() == nil {
		var ret string
		return ret
	}

	return *o.Icon.Get()
}

// GetIconOk returns a tuple with the Icon field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserStatusPublic) GetIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Icon.Get(), o.Icon.IsSet()
}

// SetIcon sets field value
func (o *UserStatusPublic) SetIcon(v string) {
	o.Icon.Set(&v)
}

// GetClearAt returns the ClearAt field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *UserStatusPublic) GetClearAt() int64 {
	if o == nil || o.ClearAt.Get() == nil {
		var ret int64
		return ret
	}

	return *o.ClearAt.Get()
}

// GetClearAtOk returns a tuple with the ClearAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserStatusPublic) GetClearAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClearAt.Get(), o.ClearAt.IsSet()
}

// SetClearAt sets field value
func (o *UserStatusPublic) SetClearAt(v int64) {
	o.ClearAt.Set(&v)
}

// GetStatus returns the Status field value
func (o *UserStatusPublic) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *UserStatusPublic) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *UserStatusPublic) SetStatus(v string) {
	o.Status = v
}

func (o UserStatusPublic) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserStatusPublic) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["userId"] = o.UserId
	toSerialize["message"] = o.Message.Get()
	toSerialize["icon"] = o.Icon.Get()
	toSerialize["clearAt"] = o.ClearAt.Get()
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

type NullableUserStatusPublic struct {
	value *UserStatusPublic
	isSet bool
}

func (v NullableUserStatusPublic) Get() *UserStatusPublic {
	return v.value
}

func (v *NullableUserStatusPublic) Set(val *UserStatusPublic) {
	v.value = val
	v.isSet = true
}

func (v NullableUserStatusPublic) IsSet() bool {
	return v.isSet
}

func (v *NullableUserStatusPublic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserStatusPublic(val *UserStatusPublic) *NullableUserStatusPublic {
	return &NullableUserStatusPublic{value: val, isSet: true}
}

func (v NullableUserStatusPublic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserStatusPublic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


