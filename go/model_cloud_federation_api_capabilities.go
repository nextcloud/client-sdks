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

// checks if the CloudFederationApiCapabilities type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudFederationApiCapabilities{}

// CloudFederationApiCapabilities struct for CloudFederationApiCapabilities
type CloudFederationApiCapabilities struct {
	Ocm CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcm `json:"ocm"`
}

// NewCloudFederationApiCapabilities instantiates a new CloudFederationApiCapabilities object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudFederationApiCapabilities(ocm CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcm) *CloudFederationApiCapabilities {
	this := CloudFederationApiCapabilities{}
	this.Ocm = ocm
	return &this
}

// NewCloudFederationApiCapabilitiesWithDefaults instantiates a new CloudFederationApiCapabilities object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudFederationApiCapabilitiesWithDefaults() *CloudFederationApiCapabilities {
	this := CloudFederationApiCapabilities{}
	return &this
}

// GetOcm returns the Ocm field value
func (o *CloudFederationApiCapabilities) GetOcm() CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcm {
	if o == nil {
		var ret CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcm
		return ret
	}

	return o.Ocm
}

// GetOcmOk returns a tuple with the Ocm field value
// and a boolean to check if the value has been set.
func (o *CloudFederationApiCapabilities) GetOcmOk() (*CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcm, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ocm, true
}

// SetOcm sets field value
func (o *CloudFederationApiCapabilities) SetOcm(v CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcm) {
	o.Ocm = v
}

func (o CloudFederationApiCapabilities) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudFederationApiCapabilities) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ocm"] = o.Ocm
	return toSerialize, nil
}

type NullableCloudFederationApiCapabilities struct {
	value *CloudFederationApiCapabilities
	isSet bool
}

func (v NullableCloudFederationApiCapabilities) Get() *CloudFederationApiCapabilities {
	return v.value
}

func (v *NullableCloudFederationApiCapabilities) Set(val *CloudFederationApiCapabilities) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudFederationApiCapabilities) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudFederationApiCapabilities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudFederationApiCapabilities(val *CloudFederationApiCapabilities) *NullableCloudFederationApiCapabilities {
	return &NullableCloudFederationApiCapabilities{value: val, isSet: true}
}

func (v NullableCloudFederationApiCapabilities) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudFederationApiCapabilities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


