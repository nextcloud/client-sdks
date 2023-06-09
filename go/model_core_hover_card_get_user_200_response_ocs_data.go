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

// checks if the CoreHoverCardGetUser200ResponseOcsData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreHoverCardGetUser200ResponseOcsData{}

// CoreHoverCardGetUser200ResponseOcsData struct for CoreHoverCardGetUser200ResponseOcsData
type CoreHoverCardGetUser200ResponseOcsData struct {
	UserId string `json:"userId"`
	DisplayName string `json:"displayName"`
	Actions []CoreContactsAction `json:"actions"`
}

// NewCoreHoverCardGetUser200ResponseOcsData instantiates a new CoreHoverCardGetUser200ResponseOcsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreHoverCardGetUser200ResponseOcsData(userId string, displayName string, actions []CoreContactsAction) *CoreHoverCardGetUser200ResponseOcsData {
	this := CoreHoverCardGetUser200ResponseOcsData{}
	this.UserId = userId
	this.DisplayName = displayName
	this.Actions = actions
	return &this
}

// NewCoreHoverCardGetUser200ResponseOcsDataWithDefaults instantiates a new CoreHoverCardGetUser200ResponseOcsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreHoverCardGetUser200ResponseOcsDataWithDefaults() *CoreHoverCardGetUser200ResponseOcsData {
	this := CoreHoverCardGetUser200ResponseOcsData{}
	return &this
}

// GetUserId returns the UserId field value
func (o *CoreHoverCardGetUser200ResponseOcsData) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *CoreHoverCardGetUser200ResponseOcsData) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *CoreHoverCardGetUser200ResponseOcsData) SetUserId(v string) {
	o.UserId = v
}

// GetDisplayName returns the DisplayName field value
func (o *CoreHoverCardGetUser200ResponseOcsData) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *CoreHoverCardGetUser200ResponseOcsData) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *CoreHoverCardGetUser200ResponseOcsData) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetActions returns the Actions field value
func (o *CoreHoverCardGetUser200ResponseOcsData) GetActions() []CoreContactsAction {
	if o == nil {
		var ret []CoreContactsAction
		return ret
	}

	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value
// and a boolean to check if the value has been set.
func (o *CoreHoverCardGetUser200ResponseOcsData) GetActionsOk() ([]CoreContactsAction, bool) {
	if o == nil {
		return nil, false
	}
	return o.Actions, true
}

// SetActions sets field value
func (o *CoreHoverCardGetUser200ResponseOcsData) SetActions(v []CoreContactsAction) {
	o.Actions = v
}

func (o CoreHoverCardGetUser200ResponseOcsData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreHoverCardGetUser200ResponseOcsData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["userId"] = o.UserId
	toSerialize["displayName"] = o.DisplayName
	toSerialize["actions"] = o.Actions
	return toSerialize, nil
}

type NullableCoreHoverCardGetUser200ResponseOcsData struct {
	value *CoreHoverCardGetUser200ResponseOcsData
	isSet bool
}

func (v NullableCoreHoverCardGetUser200ResponseOcsData) Get() *CoreHoverCardGetUser200ResponseOcsData {
	return v.value
}

func (v *NullableCoreHoverCardGetUser200ResponseOcsData) Set(val *CoreHoverCardGetUser200ResponseOcsData) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreHoverCardGetUser200ResponseOcsData) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreHoverCardGetUser200ResponseOcsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreHoverCardGetUser200ResponseOcsData(val *CoreHoverCardGetUser200ResponseOcsData) *NullableCoreHoverCardGetUser200ResponseOcsData {
	return &NullableCoreHoverCardGetUser200ResponseOcsData{value: val, isSet: true}
}

func (v NullableCoreHoverCardGetUser200ResponseOcsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreHoverCardGetUser200ResponseOcsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


