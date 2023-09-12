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

// checks if the CoreOpenGraphObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreOpenGraphObject{}

// CoreOpenGraphObject struct for CoreOpenGraphObject
type CoreOpenGraphObject struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Description NullableString `json:"description"`
	Thumb NullableString `json:"thumb"`
	Link string `json:"link"`
	AdditionalProperties map[string]interface{}
}

type _CoreOpenGraphObject CoreOpenGraphObject

// NewCoreOpenGraphObject instantiates a new CoreOpenGraphObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreOpenGraphObject(id string, name string, description NullableString, thumb NullableString, link string) *CoreOpenGraphObject {
	this := CoreOpenGraphObject{}
	this.Id = id
	this.Name = name
	this.Description = description
	this.Thumb = thumb
	this.Link = link
	return &this
}

// NewCoreOpenGraphObjectWithDefaults instantiates a new CoreOpenGraphObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreOpenGraphObjectWithDefaults() *CoreOpenGraphObject {
	this := CoreOpenGraphObject{}
	return &this
}

// GetId returns the Id field value
func (o *CoreOpenGraphObject) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CoreOpenGraphObject) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CoreOpenGraphObject) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *CoreOpenGraphObject) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CoreOpenGraphObject) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CoreOpenGraphObject) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CoreOpenGraphObject) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CoreOpenGraphObject) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *CoreOpenGraphObject) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetThumb returns the Thumb field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CoreOpenGraphObject) GetThumb() string {
	if o == nil || o.Thumb.Get() == nil {
		var ret string
		return ret
	}

	return *o.Thumb.Get()
}

// GetThumbOk returns a tuple with the Thumb field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CoreOpenGraphObject) GetThumbOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Thumb.Get(), o.Thumb.IsSet()
}

// SetThumb sets field value
func (o *CoreOpenGraphObject) SetThumb(v string) {
	o.Thumb.Set(&v)
}

// GetLink returns the Link field value
func (o *CoreOpenGraphObject) GetLink() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Link
}

// GetLinkOk returns a tuple with the Link field value
// and a boolean to check if the value has been set.
func (o *CoreOpenGraphObject) GetLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Link, true
}

// SetLink sets field value
func (o *CoreOpenGraphObject) SetLink(v string) {
	o.Link = v
}

func (o CoreOpenGraphObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreOpenGraphObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description.Get()
	toSerialize["thumb"] = o.Thumb.Get()
	toSerialize["link"] = o.Link

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CoreOpenGraphObject) UnmarshalJSON(bytes []byte) (err error) {
	varCoreOpenGraphObject := _CoreOpenGraphObject{}

	if err = json.Unmarshal(bytes, &varCoreOpenGraphObject); err == nil {
		*o = CoreOpenGraphObject(varCoreOpenGraphObject)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "thumb")
		delete(additionalProperties, "link")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCoreOpenGraphObject struct {
	value *CoreOpenGraphObject
	isSet bool
}

func (v NullableCoreOpenGraphObject) Get() *CoreOpenGraphObject {
	return v.value
}

func (v *NullableCoreOpenGraphObject) Set(val *CoreOpenGraphObject) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreOpenGraphObject) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreOpenGraphObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreOpenGraphObject(val *CoreOpenGraphObject) *NullableCoreOpenGraphObject {
	return &NullableCoreOpenGraphObject{value: val, isSet: true}
}

func (v NullableCoreOpenGraphObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreOpenGraphObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


