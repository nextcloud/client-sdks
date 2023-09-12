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

// checks if the CoreReference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreReference{}

// CoreReference struct for CoreReference
type CoreReference struct {
	RichObjectType string `json:"richObjectType"`
	RichObject map[string]map[string]interface{} `json:"richObject"`
	OpenGraphObject CoreOpenGraphObject `json:"openGraphObject"`
	Accessible bool `json:"accessible"`
	AdditionalProperties map[string]interface{}
}

type _CoreReference CoreReference

// NewCoreReference instantiates a new CoreReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreReference(richObjectType string, richObject map[string]map[string]interface{}, openGraphObject CoreOpenGraphObject, accessible bool) *CoreReference {
	this := CoreReference{}
	this.RichObjectType = richObjectType
	this.RichObject = richObject
	this.OpenGraphObject = openGraphObject
	this.Accessible = accessible
	return &this
}

// NewCoreReferenceWithDefaults instantiates a new CoreReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreReferenceWithDefaults() *CoreReference {
	this := CoreReference{}
	return &this
}

// GetRichObjectType returns the RichObjectType field value
func (o *CoreReference) GetRichObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RichObjectType
}

// GetRichObjectTypeOk returns a tuple with the RichObjectType field value
// and a boolean to check if the value has been set.
func (o *CoreReference) GetRichObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RichObjectType, true
}

// SetRichObjectType sets field value
func (o *CoreReference) SetRichObjectType(v string) {
	o.RichObjectType = v
}

// GetRichObject returns the RichObject field value
func (o *CoreReference) GetRichObject() map[string]map[string]interface{} {
	if o == nil {
		var ret map[string]map[string]interface{}
		return ret
	}

	return o.RichObject
}

// GetRichObjectOk returns a tuple with the RichObject field value
// and a boolean to check if the value has been set.
func (o *CoreReference) GetRichObjectOk() (map[string]map[string]interface{}, bool) {
	if o == nil {
		return map[string]map[string]interface{}{}, false
	}
	return o.RichObject, true
}

// SetRichObject sets field value
func (o *CoreReference) SetRichObject(v map[string]map[string]interface{}) {
	o.RichObject = v
}

// GetOpenGraphObject returns the OpenGraphObject field value
func (o *CoreReference) GetOpenGraphObject() CoreOpenGraphObject {
	if o == nil {
		var ret CoreOpenGraphObject
		return ret
	}

	return o.OpenGraphObject
}

// GetOpenGraphObjectOk returns a tuple with the OpenGraphObject field value
// and a boolean to check if the value has been set.
func (o *CoreReference) GetOpenGraphObjectOk() (*CoreOpenGraphObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OpenGraphObject, true
}

// SetOpenGraphObject sets field value
func (o *CoreReference) SetOpenGraphObject(v CoreOpenGraphObject) {
	o.OpenGraphObject = v
}

// GetAccessible returns the Accessible field value
func (o *CoreReference) GetAccessible() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Accessible
}

// GetAccessibleOk returns a tuple with the Accessible field value
// and a boolean to check if the value has been set.
func (o *CoreReference) GetAccessibleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Accessible, true
}

// SetAccessible sets field value
func (o *CoreReference) SetAccessible(v bool) {
	o.Accessible = v
}

func (o CoreReference) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreReference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["richObjectType"] = o.RichObjectType
	toSerialize["richObject"] = o.RichObject
	toSerialize["openGraphObject"] = o.OpenGraphObject
	toSerialize["accessible"] = o.Accessible

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CoreReference) UnmarshalJSON(bytes []byte) (err error) {
	varCoreReference := _CoreReference{}

	if err = json.Unmarshal(bytes, &varCoreReference); err == nil {
		*o = CoreReference(varCoreReference)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "richObjectType")
		delete(additionalProperties, "richObject")
		delete(additionalProperties, "openGraphObject")
		delete(additionalProperties, "accessible")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCoreReference struct {
	value *CoreReference
	isSet bool
}

func (v NullableCoreReference) Get() *CoreReference {
	return v.value
}

func (v *NullableCoreReference) Set(val *CoreReference) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreReference) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreReference(val *CoreReference) *NullableCoreReference {
	return &NullableCoreReference{value: val, isSet: true}
}

func (v NullableCoreReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

