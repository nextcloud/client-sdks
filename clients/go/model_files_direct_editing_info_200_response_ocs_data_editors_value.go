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

// checks if the FilesDirectEditingInfo200ResponseOcsDataEditorsValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FilesDirectEditingInfo200ResponseOcsDataEditorsValue{}

// FilesDirectEditingInfo200ResponseOcsDataEditorsValue struct for FilesDirectEditingInfo200ResponseOcsDataEditorsValue
type FilesDirectEditingInfo200ResponseOcsDataEditorsValue struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Mimetypes []string `json:"mimetypes"`
	OptionalMimetypes []string `json:"optionalMimetypes"`
	Secure bool `json:"secure"`
	AdditionalProperties map[string]interface{}
}

type _FilesDirectEditingInfo200ResponseOcsDataEditorsValue FilesDirectEditingInfo200ResponseOcsDataEditorsValue

// NewFilesDirectEditingInfo200ResponseOcsDataEditorsValue instantiates a new FilesDirectEditingInfo200ResponseOcsDataEditorsValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilesDirectEditingInfo200ResponseOcsDataEditorsValue(id string, name string, mimetypes []string, optionalMimetypes []string, secure bool) *FilesDirectEditingInfo200ResponseOcsDataEditorsValue {
	this := FilesDirectEditingInfo200ResponseOcsDataEditorsValue{}
	this.Id = id
	this.Name = name
	this.Mimetypes = mimetypes
	this.OptionalMimetypes = optionalMimetypes
	this.Secure = secure
	return &this
}

// NewFilesDirectEditingInfo200ResponseOcsDataEditorsValueWithDefaults instantiates a new FilesDirectEditingInfo200ResponseOcsDataEditorsValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilesDirectEditingInfo200ResponseOcsDataEditorsValueWithDefaults() *FilesDirectEditingInfo200ResponseOcsDataEditorsValue {
	this := FilesDirectEditingInfo200ResponseOcsDataEditorsValue{}
	return &this
}

// GetId returns the Id field value
func (o *FilesDirectEditingInfo200ResponseOcsDataEditorsValue) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FilesDirectEditingInfo200ResponseOcsDataEditorsValue) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FilesDirectEditingInfo200ResponseOcsDataEditorsValue) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *FilesDirectEditingInfo200ResponseOcsDataEditorsValue) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FilesDirectEditingInfo200ResponseOcsDataEditorsValue) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FilesDirectEditingInfo200ResponseOcsDataEditorsValue) SetName(v string) {
	o.Name = v
}

// GetMimetypes returns the Mimetypes field value
func (o *FilesDirectEditingInfo200ResponseOcsDataEditorsValue) GetMimetypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Mimetypes
}

// GetMimetypesOk returns a tuple with the Mimetypes field value
// and a boolean to check if the value has been set.
func (o *FilesDirectEditingInfo200ResponseOcsDataEditorsValue) GetMimetypesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mimetypes, true
}

// SetMimetypes sets field value
func (o *FilesDirectEditingInfo200ResponseOcsDataEditorsValue) SetMimetypes(v []string) {
	o.Mimetypes = v
}

// GetOptionalMimetypes returns the OptionalMimetypes field value
func (o *FilesDirectEditingInfo200ResponseOcsDataEditorsValue) GetOptionalMimetypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.OptionalMimetypes
}

// GetOptionalMimetypesOk returns a tuple with the OptionalMimetypes field value
// and a boolean to check if the value has been set.
func (o *FilesDirectEditingInfo200ResponseOcsDataEditorsValue) GetOptionalMimetypesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OptionalMimetypes, true
}

// SetOptionalMimetypes sets field value
func (o *FilesDirectEditingInfo200ResponseOcsDataEditorsValue) SetOptionalMimetypes(v []string) {
	o.OptionalMimetypes = v
}

// GetSecure returns the Secure field value
func (o *FilesDirectEditingInfo200ResponseOcsDataEditorsValue) GetSecure() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Secure
}

// GetSecureOk returns a tuple with the Secure field value
// and a boolean to check if the value has been set.
func (o *FilesDirectEditingInfo200ResponseOcsDataEditorsValue) GetSecureOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Secure, true
}

// SetSecure sets field value
func (o *FilesDirectEditingInfo200ResponseOcsDataEditorsValue) SetSecure(v bool) {
	o.Secure = v
}

func (o FilesDirectEditingInfo200ResponseOcsDataEditorsValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FilesDirectEditingInfo200ResponseOcsDataEditorsValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["mimetypes"] = o.Mimetypes
	toSerialize["optionalMimetypes"] = o.OptionalMimetypes
	toSerialize["secure"] = o.Secure

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FilesDirectEditingInfo200ResponseOcsDataEditorsValue) UnmarshalJSON(bytes []byte) (err error) {
	varFilesDirectEditingInfo200ResponseOcsDataEditorsValue := _FilesDirectEditingInfo200ResponseOcsDataEditorsValue{}

	if err = json.Unmarshal(bytes, &varFilesDirectEditingInfo200ResponseOcsDataEditorsValue); err == nil {
		*o = FilesDirectEditingInfo200ResponseOcsDataEditorsValue(varFilesDirectEditingInfo200ResponseOcsDataEditorsValue)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "mimetypes")
		delete(additionalProperties, "optionalMimetypes")
		delete(additionalProperties, "secure")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFilesDirectEditingInfo200ResponseOcsDataEditorsValue struct {
	value *FilesDirectEditingInfo200ResponseOcsDataEditorsValue
	isSet bool
}

func (v NullableFilesDirectEditingInfo200ResponseOcsDataEditorsValue) Get() *FilesDirectEditingInfo200ResponseOcsDataEditorsValue {
	return v.value
}

func (v *NullableFilesDirectEditingInfo200ResponseOcsDataEditorsValue) Set(val *FilesDirectEditingInfo200ResponseOcsDataEditorsValue) {
	v.value = val
	v.isSet = true
}

func (v NullableFilesDirectEditingInfo200ResponseOcsDataEditorsValue) IsSet() bool {
	return v.isSet
}

func (v *NullableFilesDirectEditingInfo200ResponseOcsDataEditorsValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilesDirectEditingInfo200ResponseOcsDataEditorsValue(val *FilesDirectEditingInfo200ResponseOcsDataEditorsValue) *NullableFilesDirectEditingInfo200ResponseOcsDataEditorsValue {
	return &NullableFilesDirectEditingInfo200ResponseOcsDataEditorsValue{value: val, isSet: true}
}

func (v NullableFilesDirectEditingInfo200ResponseOcsDataEditorsValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilesDirectEditingInfo200ResponseOcsDataEditorsValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

