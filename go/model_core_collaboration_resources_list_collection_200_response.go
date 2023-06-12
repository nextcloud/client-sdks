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

// checks if the CoreCollaborationResourcesListCollection200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreCollaborationResourcesListCollection200Response{}

// CoreCollaborationResourcesListCollection200Response struct for CoreCollaborationResourcesListCollection200Response
type CoreCollaborationResourcesListCollection200Response struct {
	Ocs CoreCollaborationResourcesListCollection200ResponseOcs `json:"ocs"`
}

// NewCoreCollaborationResourcesListCollection200Response instantiates a new CoreCollaborationResourcesListCollection200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreCollaborationResourcesListCollection200Response(ocs CoreCollaborationResourcesListCollection200ResponseOcs) *CoreCollaborationResourcesListCollection200Response {
	this := CoreCollaborationResourcesListCollection200Response{}
	this.Ocs = ocs
	return &this
}

// NewCoreCollaborationResourcesListCollection200ResponseWithDefaults instantiates a new CoreCollaborationResourcesListCollection200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreCollaborationResourcesListCollection200ResponseWithDefaults() *CoreCollaborationResourcesListCollection200Response {
	this := CoreCollaborationResourcesListCollection200Response{}
	return &this
}

// GetOcs returns the Ocs field value
func (o *CoreCollaborationResourcesListCollection200Response) GetOcs() CoreCollaborationResourcesListCollection200ResponseOcs {
	if o == nil {
		var ret CoreCollaborationResourcesListCollection200ResponseOcs
		return ret
	}

	return o.Ocs
}

// GetOcsOk returns a tuple with the Ocs field value
// and a boolean to check if the value has been set.
func (o *CoreCollaborationResourcesListCollection200Response) GetOcsOk() (*CoreCollaborationResourcesListCollection200ResponseOcs, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ocs, true
}

// SetOcs sets field value
func (o *CoreCollaborationResourcesListCollection200Response) SetOcs(v CoreCollaborationResourcesListCollection200ResponseOcs) {
	o.Ocs = v
}

func (o CoreCollaborationResourcesListCollection200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreCollaborationResourcesListCollection200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ocs"] = o.Ocs
	return toSerialize, nil
}

type NullableCoreCollaborationResourcesListCollection200Response struct {
	value *CoreCollaborationResourcesListCollection200Response
	isSet bool
}

func (v NullableCoreCollaborationResourcesListCollection200Response) Get() *CoreCollaborationResourcesListCollection200Response {
	return v.value
}

func (v *NullableCoreCollaborationResourcesListCollection200Response) Set(val *CoreCollaborationResourcesListCollection200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreCollaborationResourcesListCollection200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreCollaborationResourcesListCollection200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreCollaborationResourcesListCollection200Response(val *CoreCollaborationResourcesListCollection200Response) *NullableCoreCollaborationResourcesListCollection200Response {
	return &NullableCoreCollaborationResourcesListCollection200Response{value: val, isSet: true}
}

func (v NullableCoreCollaborationResourcesListCollection200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreCollaborationResourcesListCollection200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


