/*
nextcloud

Nextcloud APIs

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// ProvisioningApiUserDetailsQuotaQuota - struct for ProvisioningApiUserDetailsQuotaQuota
type ProvisioningApiUserDetailsQuotaQuota struct {
	Bool *bool
	Int64 *int64
	String *string
}

// boolAsProvisioningApiUserDetailsQuotaQuota is a convenience function that returns bool wrapped in ProvisioningApiUserDetailsQuotaQuota
func BoolAsProvisioningApiUserDetailsQuotaQuota(v *bool) ProvisioningApiUserDetailsQuotaQuota {
	return ProvisioningApiUserDetailsQuotaQuota{
		Bool: v,
	}
}

// int64AsProvisioningApiUserDetailsQuotaQuota is a convenience function that returns int64 wrapped in ProvisioningApiUserDetailsQuotaQuota
func Int64AsProvisioningApiUserDetailsQuotaQuota(v *int64) ProvisioningApiUserDetailsQuotaQuota {
	return ProvisioningApiUserDetailsQuotaQuota{
		Int64: v,
	}
}

// stringAsProvisioningApiUserDetailsQuotaQuota is a convenience function that returns string wrapped in ProvisioningApiUserDetailsQuotaQuota
func StringAsProvisioningApiUserDetailsQuotaQuota(v *string) ProvisioningApiUserDetailsQuotaQuota {
	return ProvisioningApiUserDetailsQuotaQuota{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ProvisioningApiUserDetailsQuotaQuota) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Bool
	err = newStrictDecoder(data).Decode(&dst.Bool)
	if err == nil {
		jsonBool, _ := json.Marshal(dst.Bool)
		if string(jsonBool) == "{}" { // empty struct
			dst.Bool = nil
		} else {
			match++
		}
	} else {
		dst.Bool = nil
	}

	// try to unmarshal data into Int64
	err = newStrictDecoder(data).Decode(&dst.Int64)
	if err == nil {
		jsonInt64, _ := json.Marshal(dst.Int64)
		if string(jsonInt64) == "{}" { // empty struct
			dst.Int64 = nil
		} else {
			match++
		}
	} else {
		dst.Int64 = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Bool = nil
		dst.Int64 = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ProvisioningApiUserDetailsQuotaQuota)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ProvisioningApiUserDetailsQuotaQuota)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ProvisioningApiUserDetailsQuotaQuota) MarshalJSON() ([]byte, error) {
	if src.Bool != nil {
		return json.Marshal(&src.Bool)
	}

	if src.Int64 != nil {
		return json.Marshal(&src.Int64)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ProvisioningApiUserDetailsQuotaQuota) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Bool != nil {
		return obj.Bool
	}

	if obj.Int64 != nil {
		return obj.Int64
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableProvisioningApiUserDetailsQuotaQuota struct {
	value *ProvisioningApiUserDetailsQuotaQuota
	isSet bool
}

func (v NullableProvisioningApiUserDetailsQuotaQuota) Get() *ProvisioningApiUserDetailsQuotaQuota {
	return v.value
}

func (v *NullableProvisioningApiUserDetailsQuotaQuota) Set(val *ProvisioningApiUserDetailsQuotaQuota) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisioningApiUserDetailsQuotaQuota) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisioningApiUserDetailsQuotaQuota) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisioningApiUserDetailsQuotaQuota(val *ProvisioningApiUserDetailsQuotaQuota) *NullableProvisioningApiUserDetailsQuotaQuota {
	return &NullableProvisioningApiUserDetailsQuotaQuota{value: val, isSet: true}
}

func (v NullableProvisioningApiUserDetailsQuotaQuota) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisioningApiUserDetailsQuotaQuota) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


