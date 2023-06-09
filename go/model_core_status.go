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

// checks if the CoreStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreStatus{}

// CoreStatus struct for CoreStatus
type CoreStatus struct {
	Installed bool `json:"installed"`
	Maintenance bool `json:"maintenance"`
	NeedsDbUpgrade bool `json:"needsDbUpgrade"`
	Version string `json:"version"`
	Versionstring string `json:"versionstring"`
	Edition string `json:"edition"`
	Productname string `json:"productname"`
	ExtendedSupport bool `json:"extendedSupport"`
}

// NewCoreStatus instantiates a new CoreStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreStatus(installed bool, maintenance bool, needsDbUpgrade bool, version string, versionstring string, edition string, productname string, extendedSupport bool) *CoreStatus {
	this := CoreStatus{}
	this.Installed = installed
	this.Maintenance = maintenance
	this.NeedsDbUpgrade = needsDbUpgrade
	this.Version = version
	this.Versionstring = versionstring
	this.Edition = edition
	this.Productname = productname
	this.ExtendedSupport = extendedSupport
	return &this
}

// NewCoreStatusWithDefaults instantiates a new CoreStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreStatusWithDefaults() *CoreStatus {
	this := CoreStatus{}
	return &this
}

// GetInstalled returns the Installed field value
func (o *CoreStatus) GetInstalled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Installed
}

// GetInstalledOk returns a tuple with the Installed field value
// and a boolean to check if the value has been set.
func (o *CoreStatus) GetInstalledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Installed, true
}

// SetInstalled sets field value
func (o *CoreStatus) SetInstalled(v bool) {
	o.Installed = v
}

// GetMaintenance returns the Maintenance field value
func (o *CoreStatus) GetMaintenance() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Maintenance
}

// GetMaintenanceOk returns a tuple with the Maintenance field value
// and a boolean to check if the value has been set.
func (o *CoreStatus) GetMaintenanceOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Maintenance, true
}

// SetMaintenance sets field value
func (o *CoreStatus) SetMaintenance(v bool) {
	o.Maintenance = v
}

// GetNeedsDbUpgrade returns the NeedsDbUpgrade field value
func (o *CoreStatus) GetNeedsDbUpgrade() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.NeedsDbUpgrade
}

// GetNeedsDbUpgradeOk returns a tuple with the NeedsDbUpgrade field value
// and a boolean to check if the value has been set.
func (o *CoreStatus) GetNeedsDbUpgradeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NeedsDbUpgrade, true
}

// SetNeedsDbUpgrade sets field value
func (o *CoreStatus) SetNeedsDbUpgrade(v bool) {
	o.NeedsDbUpgrade = v
}

// GetVersion returns the Version field value
func (o *CoreStatus) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *CoreStatus) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *CoreStatus) SetVersion(v string) {
	o.Version = v
}

// GetVersionstring returns the Versionstring field value
func (o *CoreStatus) GetVersionstring() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Versionstring
}

// GetVersionstringOk returns a tuple with the Versionstring field value
// and a boolean to check if the value has been set.
func (o *CoreStatus) GetVersionstringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Versionstring, true
}

// SetVersionstring sets field value
func (o *CoreStatus) SetVersionstring(v string) {
	o.Versionstring = v
}

// GetEdition returns the Edition field value
func (o *CoreStatus) GetEdition() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Edition
}

// GetEditionOk returns a tuple with the Edition field value
// and a boolean to check if the value has been set.
func (o *CoreStatus) GetEditionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Edition, true
}

// SetEdition sets field value
func (o *CoreStatus) SetEdition(v string) {
	o.Edition = v
}

// GetProductname returns the Productname field value
func (o *CoreStatus) GetProductname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Productname
}

// GetProductnameOk returns a tuple with the Productname field value
// and a boolean to check if the value has been set.
func (o *CoreStatus) GetProductnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Productname, true
}

// SetProductname sets field value
func (o *CoreStatus) SetProductname(v string) {
	o.Productname = v
}

// GetExtendedSupport returns the ExtendedSupport field value
func (o *CoreStatus) GetExtendedSupport() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ExtendedSupport
}

// GetExtendedSupportOk returns a tuple with the ExtendedSupport field value
// and a boolean to check if the value has been set.
func (o *CoreStatus) GetExtendedSupportOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExtendedSupport, true
}

// SetExtendedSupport sets field value
func (o *CoreStatus) SetExtendedSupport(v bool) {
	o.ExtendedSupport = v
}

func (o CoreStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["installed"] = o.Installed
	toSerialize["maintenance"] = o.Maintenance
	toSerialize["needsDbUpgrade"] = o.NeedsDbUpgrade
	toSerialize["version"] = o.Version
	toSerialize["versionstring"] = o.Versionstring
	toSerialize["edition"] = o.Edition
	toSerialize["productname"] = o.Productname
	toSerialize["extendedSupport"] = o.ExtendedSupport
	return toSerialize, nil
}

type NullableCoreStatus struct {
	value *CoreStatus
	isSet bool
}

func (v NullableCoreStatus) Get() *CoreStatus {
	return v.value
}

func (v *NullableCoreStatus) Set(val *CoreStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreStatus(val *CoreStatus) *NullableCoreStatus {
	return &NullableCoreStatus{value: val, isSet: true}
}

func (v NullableCoreStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


