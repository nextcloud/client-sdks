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

// CoreNavigationEntryOrder - struct for CoreNavigationEntryOrder
type CoreNavigationEntryOrder struct {
	Int64 *int64
	String *string
}

// int64AsCoreNavigationEntryOrder is a convenience function that returns int64 wrapped in CoreNavigationEntryOrder
func Int64AsCoreNavigationEntryOrder(v *int64) CoreNavigationEntryOrder {
	return CoreNavigationEntryOrder{
		Int64: v,
	}
}

// stringAsCoreNavigationEntryOrder is a convenience function that returns string wrapped in CoreNavigationEntryOrder
func StringAsCoreNavigationEntryOrder(v *string) CoreNavigationEntryOrder {
	return CoreNavigationEntryOrder{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CoreNavigationEntryOrder) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
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
		dst.Int64 = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CoreNavigationEntryOrder)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CoreNavigationEntryOrder)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CoreNavigationEntryOrder) MarshalJSON() ([]byte, error) {
	if src.Int64 != nil {
		return json.Marshal(&src.Int64)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CoreNavigationEntryOrder) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
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

type NullableCoreNavigationEntryOrder struct {
	value *CoreNavigationEntryOrder
	isSet bool
}

func (v NullableCoreNavigationEntryOrder) Get() *CoreNavigationEntryOrder {
	return v.value
}

func (v *NullableCoreNavigationEntryOrder) Set(val *CoreNavigationEntryOrder) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreNavigationEntryOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreNavigationEntryOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreNavigationEntryOrder(val *CoreNavigationEntryOrder) *NullableCoreNavigationEntryOrder {
	return &NullableCoreNavigationEntryOrder{value: val, isSet: true}
}

func (v NullableCoreNavigationEntryOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreNavigationEntryOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


