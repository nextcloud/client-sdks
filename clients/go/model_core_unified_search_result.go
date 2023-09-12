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

// checks if the CoreUnifiedSearchResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreUnifiedSearchResult{}

// CoreUnifiedSearchResult struct for CoreUnifiedSearchResult
type CoreUnifiedSearchResult struct {
	Name string `json:"name"`
	IsPaginated bool `json:"isPaginated"`
	Entries []CoreUnifiedSearchResultEntry `json:"entries"`
	Cursor NullableCoreUnifiedSearchResultCursor `json:"cursor"`
	AdditionalProperties map[string]interface{}
}

type _CoreUnifiedSearchResult CoreUnifiedSearchResult

// NewCoreUnifiedSearchResult instantiates a new CoreUnifiedSearchResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreUnifiedSearchResult(name string, isPaginated bool, entries []CoreUnifiedSearchResultEntry, cursor NullableCoreUnifiedSearchResultCursor) *CoreUnifiedSearchResult {
	this := CoreUnifiedSearchResult{}
	this.Name = name
	this.IsPaginated = isPaginated
	this.Entries = entries
	this.Cursor = cursor
	return &this
}

// NewCoreUnifiedSearchResultWithDefaults instantiates a new CoreUnifiedSearchResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreUnifiedSearchResultWithDefaults() *CoreUnifiedSearchResult {
	this := CoreUnifiedSearchResult{}
	return &this
}

// GetName returns the Name field value
func (o *CoreUnifiedSearchResult) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CoreUnifiedSearchResult) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CoreUnifiedSearchResult) SetName(v string) {
	o.Name = v
}

// GetIsPaginated returns the IsPaginated field value
func (o *CoreUnifiedSearchResult) GetIsPaginated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsPaginated
}

// GetIsPaginatedOk returns a tuple with the IsPaginated field value
// and a boolean to check if the value has been set.
func (o *CoreUnifiedSearchResult) GetIsPaginatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPaginated, true
}

// SetIsPaginated sets field value
func (o *CoreUnifiedSearchResult) SetIsPaginated(v bool) {
	o.IsPaginated = v
}

// GetEntries returns the Entries field value
func (o *CoreUnifiedSearchResult) GetEntries() []CoreUnifiedSearchResultEntry {
	if o == nil {
		var ret []CoreUnifiedSearchResultEntry
		return ret
	}

	return o.Entries
}

// GetEntriesOk returns a tuple with the Entries field value
// and a boolean to check if the value has been set.
func (o *CoreUnifiedSearchResult) GetEntriesOk() ([]CoreUnifiedSearchResultEntry, bool) {
	if o == nil {
		return nil, false
	}
	return o.Entries, true
}

// SetEntries sets field value
func (o *CoreUnifiedSearchResult) SetEntries(v []CoreUnifiedSearchResultEntry) {
	o.Entries = v
}

// GetCursor returns the Cursor field value
// If the value is explicit nil, the zero value for CoreUnifiedSearchResultCursor will be returned
func (o *CoreUnifiedSearchResult) GetCursor() CoreUnifiedSearchResultCursor {
	if o == nil || o.Cursor.Get() == nil {
		var ret CoreUnifiedSearchResultCursor
		return ret
	}

	return *o.Cursor.Get()
}

// GetCursorOk returns a tuple with the Cursor field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CoreUnifiedSearchResult) GetCursorOk() (*CoreUnifiedSearchResultCursor, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cursor.Get(), o.Cursor.IsSet()
}

// SetCursor sets field value
func (o *CoreUnifiedSearchResult) SetCursor(v CoreUnifiedSearchResultCursor) {
	o.Cursor.Set(&v)
}

func (o CoreUnifiedSearchResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreUnifiedSearchResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["isPaginated"] = o.IsPaginated
	toSerialize["entries"] = o.Entries
	toSerialize["cursor"] = o.Cursor.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CoreUnifiedSearchResult) UnmarshalJSON(bytes []byte) (err error) {
	varCoreUnifiedSearchResult := _CoreUnifiedSearchResult{}

	if err = json.Unmarshal(bytes, &varCoreUnifiedSearchResult); err == nil {
		*o = CoreUnifiedSearchResult(varCoreUnifiedSearchResult)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "isPaginated")
		delete(additionalProperties, "entries")
		delete(additionalProperties, "cursor")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCoreUnifiedSearchResult struct {
	value *CoreUnifiedSearchResult
	isSet bool
}

func (v NullableCoreUnifiedSearchResult) Get() *CoreUnifiedSearchResult {
	return v.value
}

func (v *NullableCoreUnifiedSearchResult) Set(val *CoreUnifiedSearchResult) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreUnifiedSearchResult) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreUnifiedSearchResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreUnifiedSearchResult(val *CoreUnifiedSearchResult) *NullableCoreUnifiedSearchResult {
	return &NullableCoreUnifiedSearchResult{value: val, isSet: true}
}

func (v NullableCoreUnifiedSearchResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreUnifiedSearchResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


