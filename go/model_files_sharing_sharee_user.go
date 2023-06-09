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

// checks if the FilesSharingShareeUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FilesSharingShareeUser{}

// FilesSharingShareeUser struct for FilesSharingShareeUser
type FilesSharingShareeUser struct {
	Count NullableInt64 `json:"count"`
	Label string `json:"label"`
	Subline string `json:"subline"`
	Icon string `json:"icon"`
	ShareWithDisplayNameUnique string `json:"shareWithDisplayNameUnique"`
	Status FilesSharingShareeUserAllOfStatus `json:"status"`
	Value FilesSharingShareeValue `json:"value"`
}

// NewFilesSharingShareeUser instantiates a new FilesSharingShareeUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilesSharingShareeUser(count NullableInt64, label string, subline string, icon string, shareWithDisplayNameUnique string, status FilesSharingShareeUserAllOfStatus, value FilesSharingShareeValue) *FilesSharingShareeUser {
	this := FilesSharingShareeUser{}
	this.Count = count
	this.Label = label
	this.Subline = subline
	this.Icon = icon
	this.ShareWithDisplayNameUnique = shareWithDisplayNameUnique
	this.Status = status
	this.Value = value
	return &this
}

// NewFilesSharingShareeUserWithDefaults instantiates a new FilesSharingShareeUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilesSharingShareeUserWithDefaults() *FilesSharingShareeUser {
	this := FilesSharingShareeUser{}
	return &this
}

// GetCount returns the Count field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *FilesSharingShareeUser) GetCount() int64 {
	if o == nil || o.Count.Get() == nil {
		var ret int64
		return ret
	}

	return *o.Count.Get()
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesSharingShareeUser) GetCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Count.Get(), o.Count.IsSet()
}

// SetCount sets field value
func (o *FilesSharingShareeUser) SetCount(v int64) {
	o.Count.Set(&v)
}

// GetLabel returns the Label field value
func (o *FilesSharingShareeUser) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareeUser) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *FilesSharingShareeUser) SetLabel(v string) {
	o.Label = v
}

// GetSubline returns the Subline field value
func (o *FilesSharingShareeUser) GetSubline() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subline
}

// GetSublineOk returns a tuple with the Subline field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareeUser) GetSublineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subline, true
}

// SetSubline sets field value
func (o *FilesSharingShareeUser) SetSubline(v string) {
	o.Subline = v
}

// GetIcon returns the Icon field value
func (o *FilesSharingShareeUser) GetIcon() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Icon
}

// GetIconOk returns a tuple with the Icon field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareeUser) GetIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Icon, true
}

// SetIcon sets field value
func (o *FilesSharingShareeUser) SetIcon(v string) {
	o.Icon = v
}

// GetShareWithDisplayNameUnique returns the ShareWithDisplayNameUnique field value
func (o *FilesSharingShareeUser) GetShareWithDisplayNameUnique() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShareWithDisplayNameUnique
}

// GetShareWithDisplayNameUniqueOk returns a tuple with the ShareWithDisplayNameUnique field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareeUser) GetShareWithDisplayNameUniqueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShareWithDisplayNameUnique, true
}

// SetShareWithDisplayNameUnique sets field value
func (o *FilesSharingShareeUser) SetShareWithDisplayNameUnique(v string) {
	o.ShareWithDisplayNameUnique = v
}

// GetStatus returns the Status field value
func (o *FilesSharingShareeUser) GetStatus() FilesSharingShareeUserAllOfStatus {
	if o == nil {
		var ret FilesSharingShareeUserAllOfStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareeUser) GetStatusOk() (*FilesSharingShareeUserAllOfStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *FilesSharingShareeUser) SetStatus(v FilesSharingShareeUserAllOfStatus) {
	o.Status = v
}

// GetValue returns the Value field value
func (o *FilesSharingShareeUser) GetValue() FilesSharingShareeValue {
	if o == nil {
		var ret FilesSharingShareeValue
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareeUser) GetValueOk() (*FilesSharingShareeValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *FilesSharingShareeUser) SetValue(v FilesSharingShareeValue) {
	o.Value = v
}

func (o FilesSharingShareeUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FilesSharingShareeUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["count"] = o.Count.Get()
	toSerialize["label"] = o.Label
	toSerialize["subline"] = o.Subline
	toSerialize["icon"] = o.Icon
	toSerialize["shareWithDisplayNameUnique"] = o.ShareWithDisplayNameUnique
	toSerialize["status"] = o.Status
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableFilesSharingShareeUser struct {
	value *FilesSharingShareeUser
	isSet bool
}

func (v NullableFilesSharingShareeUser) Get() *FilesSharingShareeUser {
	return v.value
}

func (v *NullableFilesSharingShareeUser) Set(val *FilesSharingShareeUser) {
	v.value = val
	v.isSet = true
}

func (v NullableFilesSharingShareeUser) IsSet() bool {
	return v.isSet
}

func (v *NullableFilesSharingShareeUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilesSharingShareeUser(val *FilesSharingShareeUser) *NullableFilesSharingShareeUser {
	return &NullableFilesSharingShareeUser{value: val, isSet: true}
}

func (v NullableFilesSharingShareeUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilesSharingShareeUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


