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

// checks if the CoreTextProcessingApiSchedule200ResponseOcsData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreTextProcessingApiSchedule200ResponseOcsData{}

// CoreTextProcessingApiSchedule200ResponseOcsData struct for CoreTextProcessingApiSchedule200ResponseOcsData
type CoreTextProcessingApiSchedule200ResponseOcsData struct {
	Task CoreTextProcessingTask `json:"task"`
	AdditionalProperties map[string]interface{}
}

type _CoreTextProcessingApiSchedule200ResponseOcsData CoreTextProcessingApiSchedule200ResponseOcsData

// NewCoreTextProcessingApiSchedule200ResponseOcsData instantiates a new CoreTextProcessingApiSchedule200ResponseOcsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreTextProcessingApiSchedule200ResponseOcsData(task CoreTextProcessingTask) *CoreTextProcessingApiSchedule200ResponseOcsData {
	this := CoreTextProcessingApiSchedule200ResponseOcsData{}
	this.Task = task
	return &this
}

// NewCoreTextProcessingApiSchedule200ResponseOcsDataWithDefaults instantiates a new CoreTextProcessingApiSchedule200ResponseOcsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreTextProcessingApiSchedule200ResponseOcsDataWithDefaults() *CoreTextProcessingApiSchedule200ResponseOcsData {
	this := CoreTextProcessingApiSchedule200ResponseOcsData{}
	return &this
}

// GetTask returns the Task field value
func (o *CoreTextProcessingApiSchedule200ResponseOcsData) GetTask() CoreTextProcessingTask {
	if o == nil {
		var ret CoreTextProcessingTask
		return ret
	}

	return o.Task
}

// GetTaskOk returns a tuple with the Task field value
// and a boolean to check if the value has been set.
func (o *CoreTextProcessingApiSchedule200ResponseOcsData) GetTaskOk() (*CoreTextProcessingTask, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Task, true
}

// SetTask sets field value
func (o *CoreTextProcessingApiSchedule200ResponseOcsData) SetTask(v CoreTextProcessingTask) {
	o.Task = v
}

func (o CoreTextProcessingApiSchedule200ResponseOcsData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreTextProcessingApiSchedule200ResponseOcsData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["task"] = o.Task

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CoreTextProcessingApiSchedule200ResponseOcsData) UnmarshalJSON(bytes []byte) (err error) {
	varCoreTextProcessingApiSchedule200ResponseOcsData := _CoreTextProcessingApiSchedule200ResponseOcsData{}

	if err = json.Unmarshal(bytes, &varCoreTextProcessingApiSchedule200ResponseOcsData); err == nil {
		*o = CoreTextProcessingApiSchedule200ResponseOcsData(varCoreTextProcessingApiSchedule200ResponseOcsData)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "task")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCoreTextProcessingApiSchedule200ResponseOcsData struct {
	value *CoreTextProcessingApiSchedule200ResponseOcsData
	isSet bool
}

func (v NullableCoreTextProcessingApiSchedule200ResponseOcsData) Get() *CoreTextProcessingApiSchedule200ResponseOcsData {
	return v.value
}

func (v *NullableCoreTextProcessingApiSchedule200ResponseOcsData) Set(val *CoreTextProcessingApiSchedule200ResponseOcsData) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreTextProcessingApiSchedule200ResponseOcsData) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreTextProcessingApiSchedule200ResponseOcsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreTextProcessingApiSchedule200ResponseOcsData(val *CoreTextProcessingApiSchedule200ResponseOcsData) *NullableCoreTextProcessingApiSchedule200ResponseOcsData {
	return &NullableCoreTextProcessingApiSchedule200ResponseOcsData{value: val, isSet: true}
}

func (v NullableCoreTextProcessingApiSchedule200ResponseOcsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreTextProcessingApiSchedule200ResponseOcsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

