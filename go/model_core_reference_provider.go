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

// checks if the CoreReferenceProvider type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreReferenceProvider{}

// CoreReferenceProvider struct for CoreReferenceProvider
type CoreReferenceProvider struct {
	Id string `json:"id"`
	Title string `json:"title"`
	IconUrl string `json:"icon_url"`
	Order int64 `json:"order"`
	SearchProvidersIds []string `json:"search_providers_ids"`
}

// NewCoreReferenceProvider instantiates a new CoreReferenceProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreReferenceProvider(id string, title string, iconUrl string, order int64, searchProvidersIds []string) *CoreReferenceProvider {
	this := CoreReferenceProvider{}
	this.Id = id
	this.Title = title
	this.IconUrl = iconUrl
	this.Order = order
	this.SearchProvidersIds = searchProvidersIds
	return &this
}

// NewCoreReferenceProviderWithDefaults instantiates a new CoreReferenceProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreReferenceProviderWithDefaults() *CoreReferenceProvider {
	this := CoreReferenceProvider{}
	return &this
}

// GetId returns the Id field value
func (o *CoreReferenceProvider) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CoreReferenceProvider) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CoreReferenceProvider) SetId(v string) {
	o.Id = v
}

// GetTitle returns the Title field value
func (o *CoreReferenceProvider) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *CoreReferenceProvider) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *CoreReferenceProvider) SetTitle(v string) {
	o.Title = v
}

// GetIconUrl returns the IconUrl field value
func (o *CoreReferenceProvider) GetIconUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IconUrl
}

// GetIconUrlOk returns a tuple with the IconUrl field value
// and a boolean to check if the value has been set.
func (o *CoreReferenceProvider) GetIconUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IconUrl, true
}

// SetIconUrl sets field value
func (o *CoreReferenceProvider) SetIconUrl(v string) {
	o.IconUrl = v
}

// GetOrder returns the Order field value
func (o *CoreReferenceProvider) GetOrder() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Order
}

// GetOrderOk returns a tuple with the Order field value
// and a boolean to check if the value has been set.
func (o *CoreReferenceProvider) GetOrderOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Order, true
}

// SetOrder sets field value
func (o *CoreReferenceProvider) SetOrder(v int64) {
	o.Order = v
}

// GetSearchProvidersIds returns the SearchProvidersIds field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *CoreReferenceProvider) GetSearchProvidersIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SearchProvidersIds
}

// GetSearchProvidersIdsOk returns a tuple with the SearchProvidersIds field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CoreReferenceProvider) GetSearchProvidersIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.SearchProvidersIds) {
		return nil, false
	}
	return o.SearchProvidersIds, true
}

// SetSearchProvidersIds sets field value
func (o *CoreReferenceProvider) SetSearchProvidersIds(v []string) {
	o.SearchProvidersIds = v
}

func (o CoreReferenceProvider) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreReferenceProvider) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["title"] = o.Title
	toSerialize["icon_url"] = o.IconUrl
	toSerialize["order"] = o.Order
	if o.SearchProvidersIds != nil {
		toSerialize["search_providers_ids"] = o.SearchProvidersIds
	}
	return toSerialize, nil
}

type NullableCoreReferenceProvider struct {
	value *CoreReferenceProvider
	isSet bool
}

func (v NullableCoreReferenceProvider) Get() *CoreReferenceProvider {
	return v.value
}

func (v *NullableCoreReferenceProvider) Set(val *CoreReferenceProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreReferenceProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreReferenceProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreReferenceProvider(val *CoreReferenceProvider) *NullableCoreReferenceProvider {
	return &NullableCoreReferenceProvider{value: val, isSet: true}
}

func (v NullableCoreReferenceProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreReferenceProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


