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

// checks if the ProvisioningApiGroupDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProvisioningApiGroupDetails{}

// ProvisioningApiGroupDetails struct for ProvisioningApiGroupDetails
type ProvisioningApiGroupDetails struct {
	Id string `json:"id"`
	Displayname string `json:"displayname"`
	Usercount ProvisioningApiGroupDetailsUsercount `json:"usercount"`
	Disabled ProvisioningApiGroupDetailsUsercount `json:"disabled"`
	CanAdd bool `json:"canAdd"`
	CanRemove bool `json:"canRemove"`
	AdditionalProperties map[string]interface{}
}

type _ProvisioningApiGroupDetails ProvisioningApiGroupDetails

// NewProvisioningApiGroupDetails instantiates a new ProvisioningApiGroupDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisioningApiGroupDetails(id string, displayname string, usercount ProvisioningApiGroupDetailsUsercount, disabled ProvisioningApiGroupDetailsUsercount, canAdd bool, canRemove bool) *ProvisioningApiGroupDetails {
	this := ProvisioningApiGroupDetails{}
	this.Id = id
	this.Displayname = displayname
	this.Usercount = usercount
	this.Disabled = disabled
	this.CanAdd = canAdd
	this.CanRemove = canRemove
	return &this
}

// NewProvisioningApiGroupDetailsWithDefaults instantiates a new ProvisioningApiGroupDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisioningApiGroupDetailsWithDefaults() *ProvisioningApiGroupDetails {
	this := ProvisioningApiGroupDetails{}
	return &this
}

// GetId returns the Id field value
func (o *ProvisioningApiGroupDetails) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProvisioningApiGroupDetails) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ProvisioningApiGroupDetails) SetId(v string) {
	o.Id = v
}

// GetDisplayname returns the Displayname field value
func (o *ProvisioningApiGroupDetails) GetDisplayname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Displayname
}

// GetDisplaynameOk returns a tuple with the Displayname field value
// and a boolean to check if the value has been set.
func (o *ProvisioningApiGroupDetails) GetDisplaynameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Displayname, true
}

// SetDisplayname sets field value
func (o *ProvisioningApiGroupDetails) SetDisplayname(v string) {
	o.Displayname = v
}

// GetUsercount returns the Usercount field value
func (o *ProvisioningApiGroupDetails) GetUsercount() ProvisioningApiGroupDetailsUsercount {
	if o == nil {
		var ret ProvisioningApiGroupDetailsUsercount
		return ret
	}

	return o.Usercount
}

// GetUsercountOk returns a tuple with the Usercount field value
// and a boolean to check if the value has been set.
func (o *ProvisioningApiGroupDetails) GetUsercountOk() (*ProvisioningApiGroupDetailsUsercount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Usercount, true
}

// SetUsercount sets field value
func (o *ProvisioningApiGroupDetails) SetUsercount(v ProvisioningApiGroupDetailsUsercount) {
	o.Usercount = v
}

// GetDisabled returns the Disabled field value
func (o *ProvisioningApiGroupDetails) GetDisabled() ProvisioningApiGroupDetailsUsercount {
	if o == nil {
		var ret ProvisioningApiGroupDetailsUsercount
		return ret
	}

	return o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value
// and a boolean to check if the value has been set.
func (o *ProvisioningApiGroupDetails) GetDisabledOk() (*ProvisioningApiGroupDetailsUsercount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Disabled, true
}

// SetDisabled sets field value
func (o *ProvisioningApiGroupDetails) SetDisabled(v ProvisioningApiGroupDetailsUsercount) {
	o.Disabled = v
}

// GetCanAdd returns the CanAdd field value
func (o *ProvisioningApiGroupDetails) GetCanAdd() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanAdd
}

// GetCanAddOk returns a tuple with the CanAdd field value
// and a boolean to check if the value has been set.
func (o *ProvisioningApiGroupDetails) GetCanAddOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanAdd, true
}

// SetCanAdd sets field value
func (o *ProvisioningApiGroupDetails) SetCanAdd(v bool) {
	o.CanAdd = v
}

// GetCanRemove returns the CanRemove field value
func (o *ProvisioningApiGroupDetails) GetCanRemove() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanRemove
}

// GetCanRemoveOk returns a tuple with the CanRemove field value
// and a boolean to check if the value has been set.
func (o *ProvisioningApiGroupDetails) GetCanRemoveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanRemove, true
}

// SetCanRemove sets field value
func (o *ProvisioningApiGroupDetails) SetCanRemove(v bool) {
	o.CanRemove = v
}

func (o ProvisioningApiGroupDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProvisioningApiGroupDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["displayname"] = o.Displayname
	toSerialize["usercount"] = o.Usercount
	toSerialize["disabled"] = o.Disabled
	toSerialize["canAdd"] = o.CanAdd
	toSerialize["canRemove"] = o.CanRemove

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProvisioningApiGroupDetails) UnmarshalJSON(bytes []byte) (err error) {
	varProvisioningApiGroupDetails := _ProvisioningApiGroupDetails{}

	if err = json.Unmarshal(bytes, &varProvisioningApiGroupDetails); err == nil {
		*o = ProvisioningApiGroupDetails(varProvisioningApiGroupDetails)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "displayname")
		delete(additionalProperties, "usercount")
		delete(additionalProperties, "disabled")
		delete(additionalProperties, "canAdd")
		delete(additionalProperties, "canRemove")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProvisioningApiGroupDetails struct {
	value *ProvisioningApiGroupDetails
	isSet bool
}

func (v NullableProvisioningApiGroupDetails) Get() *ProvisioningApiGroupDetails {
	return v.value
}

func (v *NullableProvisioningApiGroupDetails) Set(val *ProvisioningApiGroupDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisioningApiGroupDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisioningApiGroupDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisioningApiGroupDetails(val *ProvisioningApiGroupDetails) *NullableProvisioningApiGroupDetails {
	return &NullableProvisioningApiGroupDetails{value: val, isSet: true}
}

func (v NullableProvisioningApiGroupDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisioningApiGroupDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


