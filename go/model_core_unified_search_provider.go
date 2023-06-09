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

// checks if the CoreUnifiedSearchProvider type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreUnifiedSearchProvider{}

// CoreUnifiedSearchProvider struct for CoreUnifiedSearchProvider
type CoreUnifiedSearchProvider struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Order int64 `json:"order"`
}

// NewCoreUnifiedSearchProvider instantiates a new CoreUnifiedSearchProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreUnifiedSearchProvider(id string, name string, order int64) *CoreUnifiedSearchProvider {
	this := CoreUnifiedSearchProvider{}
	this.Id = id
	this.Name = name
	this.Order = order
	return &this
}

// NewCoreUnifiedSearchProviderWithDefaults instantiates a new CoreUnifiedSearchProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreUnifiedSearchProviderWithDefaults() *CoreUnifiedSearchProvider {
	this := CoreUnifiedSearchProvider{}
	return &this
}

// GetId returns the Id field value
func (o *CoreUnifiedSearchProvider) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CoreUnifiedSearchProvider) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CoreUnifiedSearchProvider) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *CoreUnifiedSearchProvider) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CoreUnifiedSearchProvider) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CoreUnifiedSearchProvider) SetName(v string) {
	o.Name = v
}

// GetOrder returns the Order field value
func (o *CoreUnifiedSearchProvider) GetOrder() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Order
}

// GetOrderOk returns a tuple with the Order field value
// and a boolean to check if the value has been set.
func (o *CoreUnifiedSearchProvider) GetOrderOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Order, true
}

// SetOrder sets field value
func (o *CoreUnifiedSearchProvider) SetOrder(v int64) {
	o.Order = v
}

func (o CoreUnifiedSearchProvider) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreUnifiedSearchProvider) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["order"] = o.Order
	return toSerialize, nil
}

type NullableCoreUnifiedSearchProvider struct {
	value *CoreUnifiedSearchProvider
	isSet bool
}

func (v NullableCoreUnifiedSearchProvider) Get() *CoreUnifiedSearchProvider {
	return v.value
}

func (v *NullableCoreUnifiedSearchProvider) Set(val *CoreUnifiedSearchProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreUnifiedSearchProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreUnifiedSearchProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreUnifiedSearchProvider(val *CoreUnifiedSearchProvider) *NullableCoreUnifiedSearchProvider {
	return &NullableCoreUnifiedSearchProvider{value: val, isSet: true}
}

func (v NullableCoreUnifiedSearchProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreUnifiedSearchProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


