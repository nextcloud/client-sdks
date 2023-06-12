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

// checks if the CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesDav type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesDav{}

// CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesDav struct for CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesDav
type CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesDav struct {
	Chunking string `json:"chunking"`
	Bulkupload *string `json:"bulkupload,omitempty"`
}

// NewCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesDav instantiates a new CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesDav object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesDav(chunking string) *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesDav {
	this := CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesDav{}
	this.Chunking = chunking
	return &this
}

// NewCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesDavWithDefaults instantiates a new CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesDav object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesDavWithDefaults() *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesDav {
	this := CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesDav{}
	return &this
}

// GetChunking returns the Chunking field value
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesDav) GetChunking() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Chunking
}

// GetChunkingOk returns a tuple with the Chunking field value
// and a boolean to check if the value has been set.
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesDav) GetChunkingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Chunking, true
}

// SetChunking sets field value
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesDav) SetChunking(v string) {
	o.Chunking = v
}

// GetBulkupload returns the Bulkupload field value if set, zero value otherwise.
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesDav) GetBulkupload() string {
	if o == nil || IsNil(o.Bulkupload) {
		var ret string
		return ret
	}
	return *o.Bulkupload
}

// GetBulkuploadOk returns a tuple with the Bulkupload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesDav) GetBulkuploadOk() (*string, bool) {
	if o == nil || IsNil(o.Bulkupload) {
		return nil, false
	}
	return o.Bulkupload, true
}

// HasBulkupload returns a boolean if a field has been set.
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesDav) HasBulkupload() bool {
	if o != nil && !IsNil(o.Bulkupload) {
		return true
	}

	return false
}

// SetBulkupload gets a reference to the given string and assigns it to the Bulkupload field.
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesDav) SetBulkupload(v string) {
	o.Bulkupload = &v
}

func (o CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesDav) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesDav) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["chunking"] = o.Chunking
	if !IsNil(o.Bulkupload) {
		toSerialize["bulkupload"] = o.Bulkupload
	}
	return toSerialize, nil
}

type NullableCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesDav struct {
	value *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesDav
	isSet bool
}

func (v NullableCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesDav) Get() *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesDav {
	return v.value
}

func (v *NullableCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesDav) Set(val *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesDav) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesDav) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesDav) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesDav(val *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesDav) *NullableCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesDav {
	return &NullableCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesDav{value: val, isSet: true}
}

func (v NullableCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesDav) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesDav) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


