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

// checks if the CloudFederationApiValidationErrorAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudFederationApiValidationErrorAllOf{}

// CloudFederationApiValidationErrorAllOf struct for CloudFederationApiValidationErrorAllOf
type CloudFederationApiValidationErrorAllOf struct {
	ValidationErrors []CloudFederationApiValidationErrorAllOfValidationErrors `json:"validationErrors"`
}

// NewCloudFederationApiValidationErrorAllOf instantiates a new CloudFederationApiValidationErrorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudFederationApiValidationErrorAllOf(validationErrors []CloudFederationApiValidationErrorAllOfValidationErrors) *CloudFederationApiValidationErrorAllOf {
	this := CloudFederationApiValidationErrorAllOf{}
	this.ValidationErrors = validationErrors
	return &this
}

// NewCloudFederationApiValidationErrorAllOfWithDefaults instantiates a new CloudFederationApiValidationErrorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudFederationApiValidationErrorAllOfWithDefaults() *CloudFederationApiValidationErrorAllOf {
	this := CloudFederationApiValidationErrorAllOf{}
	return &this
}

// GetValidationErrors returns the ValidationErrors field value
func (o *CloudFederationApiValidationErrorAllOf) GetValidationErrors() []CloudFederationApiValidationErrorAllOfValidationErrors {
	if o == nil {
		var ret []CloudFederationApiValidationErrorAllOfValidationErrors
		return ret
	}

	return o.ValidationErrors
}

// GetValidationErrorsOk returns a tuple with the ValidationErrors field value
// and a boolean to check if the value has been set.
func (o *CloudFederationApiValidationErrorAllOf) GetValidationErrorsOk() ([]CloudFederationApiValidationErrorAllOfValidationErrors, bool) {
	if o == nil {
		return nil, false
	}
	return o.ValidationErrors, true
}

// SetValidationErrors sets field value
func (o *CloudFederationApiValidationErrorAllOf) SetValidationErrors(v []CloudFederationApiValidationErrorAllOfValidationErrors) {
	o.ValidationErrors = v
}

func (o CloudFederationApiValidationErrorAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudFederationApiValidationErrorAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["validationErrors"] = o.ValidationErrors
	return toSerialize, nil
}

type NullableCloudFederationApiValidationErrorAllOf struct {
	value *CloudFederationApiValidationErrorAllOf
	isSet bool
}

func (v NullableCloudFederationApiValidationErrorAllOf) Get() *CloudFederationApiValidationErrorAllOf {
	return v.value
}

func (v *NullableCloudFederationApiValidationErrorAllOf) Set(val *CloudFederationApiValidationErrorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudFederationApiValidationErrorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudFederationApiValidationErrorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudFederationApiValidationErrorAllOf(val *CloudFederationApiValidationErrorAllOf) *NullableCloudFederationApiValidationErrorAllOf {
	return &NullableCloudFederationApiValidationErrorAllOf{value: val, isSet: true}
}

func (v NullableCloudFederationApiValidationErrorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudFederationApiValidationErrorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


