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

// checks if the FilesSharingShareeEmail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FilesSharingShareeEmail{}

// FilesSharingShareeEmail struct for FilesSharingShareeEmail
type FilesSharingShareeEmail struct {
	Count NullableInt64 `json:"count"`
	Label string `json:"label"`
	Uuid string `json:"uuid"`
	Name string `json:"name"`
	Type string `json:"type"`
	ShareWithDisplayNameUnique string `json:"shareWithDisplayNameUnique"`
	Value FilesSharingShareeValue `json:"value"`
	AdditionalProperties map[string]interface{}
}

type _FilesSharingShareeEmail FilesSharingShareeEmail

// NewFilesSharingShareeEmail instantiates a new FilesSharingShareeEmail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilesSharingShareeEmail(count NullableInt64, label string, uuid string, name string, type_ string, shareWithDisplayNameUnique string, value FilesSharingShareeValue) *FilesSharingShareeEmail {
	this := FilesSharingShareeEmail{}
	this.Count = count
	this.Label = label
	this.Uuid = uuid
	this.Name = name
	this.Type = type_
	this.ShareWithDisplayNameUnique = shareWithDisplayNameUnique
	this.Value = value
	return &this
}

// NewFilesSharingShareeEmailWithDefaults instantiates a new FilesSharingShareeEmail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilesSharingShareeEmailWithDefaults() *FilesSharingShareeEmail {
	this := FilesSharingShareeEmail{}
	return &this
}

// GetCount returns the Count field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *FilesSharingShareeEmail) GetCount() int64 {
	if o == nil || o.Count.Get() == nil {
		var ret int64
		return ret
	}

	return *o.Count.Get()
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesSharingShareeEmail) GetCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Count.Get(), o.Count.IsSet()
}

// SetCount sets field value
func (o *FilesSharingShareeEmail) SetCount(v int64) {
	o.Count.Set(&v)
}

// GetLabel returns the Label field value
func (o *FilesSharingShareeEmail) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareeEmail) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *FilesSharingShareeEmail) SetLabel(v string) {
	o.Label = v
}

// GetUuid returns the Uuid field value
func (o *FilesSharingShareeEmail) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareeEmail) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *FilesSharingShareeEmail) SetUuid(v string) {
	o.Uuid = v
}

// GetName returns the Name field value
func (o *FilesSharingShareeEmail) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareeEmail) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FilesSharingShareeEmail) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *FilesSharingShareeEmail) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareeEmail) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FilesSharingShareeEmail) SetType(v string) {
	o.Type = v
}

// GetShareWithDisplayNameUnique returns the ShareWithDisplayNameUnique field value
func (o *FilesSharingShareeEmail) GetShareWithDisplayNameUnique() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShareWithDisplayNameUnique
}

// GetShareWithDisplayNameUniqueOk returns a tuple with the ShareWithDisplayNameUnique field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareeEmail) GetShareWithDisplayNameUniqueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShareWithDisplayNameUnique, true
}

// SetShareWithDisplayNameUnique sets field value
func (o *FilesSharingShareeEmail) SetShareWithDisplayNameUnique(v string) {
	o.ShareWithDisplayNameUnique = v
}

// GetValue returns the Value field value
func (o *FilesSharingShareeEmail) GetValue() FilesSharingShareeValue {
	if o == nil {
		var ret FilesSharingShareeValue
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareeEmail) GetValueOk() (*FilesSharingShareeValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *FilesSharingShareeEmail) SetValue(v FilesSharingShareeValue) {
	o.Value = v
}

func (o FilesSharingShareeEmail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FilesSharingShareeEmail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["count"] = o.Count.Get()
	toSerialize["label"] = o.Label
	toSerialize["uuid"] = o.Uuid
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	toSerialize["shareWithDisplayNameUnique"] = o.ShareWithDisplayNameUnique
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FilesSharingShareeEmail) UnmarshalJSON(bytes []byte) (err error) {
	varFilesSharingShareeEmail := _FilesSharingShareeEmail{}

	if err = json.Unmarshal(bytes, &varFilesSharingShareeEmail); err == nil {
		*o = FilesSharingShareeEmail(varFilesSharingShareeEmail)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "count")
		delete(additionalProperties, "label")
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		delete(additionalProperties, "shareWithDisplayNameUnique")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFilesSharingShareeEmail struct {
	value *FilesSharingShareeEmail
	isSet bool
}

func (v NullableFilesSharingShareeEmail) Get() *FilesSharingShareeEmail {
	return v.value
}

func (v *NullableFilesSharingShareeEmail) Set(val *FilesSharingShareeEmail) {
	v.value = val
	v.isSet = true
}

func (v NullableFilesSharingShareeEmail) IsSet() bool {
	return v.isSet
}

func (v *NullableFilesSharingShareeEmail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilesSharingShareeEmail(val *FilesSharingShareeEmail) *NullableFilesSharingShareeEmail {
	return &NullableFilesSharingShareeEmail{value: val, isSet: true}
}

func (v NullableFilesSharingShareeEmail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilesSharingShareeEmail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


