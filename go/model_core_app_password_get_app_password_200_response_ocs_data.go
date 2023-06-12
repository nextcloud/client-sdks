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

// checks if the CoreAppPasswordGetAppPassword200ResponseOcsData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreAppPasswordGetAppPassword200ResponseOcsData{}

// CoreAppPasswordGetAppPassword200ResponseOcsData struct for CoreAppPasswordGetAppPassword200ResponseOcsData
type CoreAppPasswordGetAppPassword200ResponseOcsData struct {
	Apppassword string `json:"apppassword"`
}

// NewCoreAppPasswordGetAppPassword200ResponseOcsData instantiates a new CoreAppPasswordGetAppPassword200ResponseOcsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreAppPasswordGetAppPassword200ResponseOcsData(apppassword string) *CoreAppPasswordGetAppPassword200ResponseOcsData {
	this := CoreAppPasswordGetAppPassword200ResponseOcsData{}
	this.Apppassword = apppassword
	return &this
}

// NewCoreAppPasswordGetAppPassword200ResponseOcsDataWithDefaults instantiates a new CoreAppPasswordGetAppPassword200ResponseOcsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreAppPasswordGetAppPassword200ResponseOcsDataWithDefaults() *CoreAppPasswordGetAppPassword200ResponseOcsData {
	this := CoreAppPasswordGetAppPassword200ResponseOcsData{}
	return &this
}

// GetApppassword returns the Apppassword field value
func (o *CoreAppPasswordGetAppPassword200ResponseOcsData) GetApppassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Apppassword
}

// GetApppasswordOk returns a tuple with the Apppassword field value
// and a boolean to check if the value has been set.
func (o *CoreAppPasswordGetAppPassword200ResponseOcsData) GetApppasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Apppassword, true
}

// SetApppassword sets field value
func (o *CoreAppPasswordGetAppPassword200ResponseOcsData) SetApppassword(v string) {
	o.Apppassword = v
}

func (o CoreAppPasswordGetAppPassword200ResponseOcsData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreAppPasswordGetAppPassword200ResponseOcsData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apppassword"] = o.Apppassword
	return toSerialize, nil
}

type NullableCoreAppPasswordGetAppPassword200ResponseOcsData struct {
	value *CoreAppPasswordGetAppPassword200ResponseOcsData
	isSet bool
}

func (v NullableCoreAppPasswordGetAppPassword200ResponseOcsData) Get() *CoreAppPasswordGetAppPassword200ResponseOcsData {
	return v.value
}

func (v *NullableCoreAppPasswordGetAppPassword200ResponseOcsData) Set(val *CoreAppPasswordGetAppPassword200ResponseOcsData) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreAppPasswordGetAppPassword200ResponseOcsData) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreAppPasswordGetAppPassword200ResponseOcsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreAppPasswordGetAppPassword200ResponseOcsData(val *CoreAppPasswordGetAppPassword200ResponseOcsData) *NullableCoreAppPasswordGetAppPassword200ResponseOcsData {
	return &NullableCoreAppPasswordGetAppPassword200ResponseOcsData{value: val, isSet: true}
}

func (v NullableCoreAppPasswordGetAppPassword200ResponseOcsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreAppPasswordGetAppPassword200ResponseOcsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


