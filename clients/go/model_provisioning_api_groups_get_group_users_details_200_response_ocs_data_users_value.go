/*
nextcloud

Nextcloud APIs

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client_sdk

import (
	"encoding/json"
	"fmt"
)

// ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue - struct for ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue
type ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue struct {
	ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValueOneOf *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValueOneOf
	ProvisioningApiUserDetails *ProvisioningApiUserDetails
}

// ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValueOneOfAsProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue is a convenience function that returns ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValueOneOf wrapped in ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue
func ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValueOneOfAsProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue(v *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValueOneOf) ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue {
	return ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue{
		ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValueOneOf: v,
	}
}

// ProvisioningApiUserDetailsAsProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue is a convenience function that returns ProvisioningApiUserDetails wrapped in ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue
func ProvisioningApiUserDetailsAsProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue(v *ProvisioningApiUserDetails) ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue {
	return ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue{
		ProvisioningApiUserDetails: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValueOneOf
	err = newStrictDecoder(data).Decode(&dst.ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValueOneOf)
	if err == nil {
		jsonProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValueOneOf, _ := json.Marshal(dst.ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValueOneOf)
		if string(jsonProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValueOneOf) == "{}" { // empty struct
			dst.ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValueOneOf = nil
		} else {
			match++
		}
	} else {
		dst.ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValueOneOf = nil
	}

	// try to unmarshal data into ProvisioningApiUserDetails
	err = newStrictDecoder(data).Decode(&dst.ProvisioningApiUserDetails)
	if err == nil {
		jsonProvisioningApiUserDetails, _ := json.Marshal(dst.ProvisioningApiUserDetails)
		if string(jsonProvisioningApiUserDetails) == "{}" { // empty struct
			dst.ProvisioningApiUserDetails = nil
		} else {
			match++
		}
	} else {
		dst.ProvisioningApiUserDetails = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValueOneOf = nil
		dst.ProvisioningApiUserDetails = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) MarshalJSON() ([]byte, error) {
	if src.ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValueOneOf != nil {
		return json.Marshal(&src.ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValueOneOf)
	}

	if src.ProvisioningApiUserDetails != nil {
		return json.Marshal(&src.ProvisioningApiUserDetails)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValueOneOf != nil {
		return obj.ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValueOneOf
	}

	if obj.ProvisioningApiUserDetails != nil {
		return obj.ProvisioningApiUserDetails
	}

	// all schemas are nil
	return nil
}

type NullableProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue struct {
	value *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue
	isSet bool
}

func (v NullableProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) Get() *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue {
	return v.value
}

func (v *NullableProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) Set(val *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue(val *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) *NullableProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue {
	return &NullableProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue{value: val, isSet: true}
}

func (v NullableProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


