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

// checks if the CoreOcsGetCapabilities200ResponseOcsDataCapabilities type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreOcsGetCapabilities200ResponseOcsDataCapabilities{}

// CoreOcsGetCapabilities200ResponseOcsDataCapabilities struct for CoreOcsGetCapabilities200ResponseOcsDataCapabilities
type CoreOcsGetCapabilities200ResponseOcsDataCapabilities struct {
	Ocm *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcm `json:"ocm,omitempty"`
	Dav *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesDav `json:"dav,omitempty"`
	FilesSharing *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing `json:"files_sharing,omitempty"`
	ProvisioningApi *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesProvisioningApi `json:"provisioning_api,omitempty"`
	Theming *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesTheming `json:"theming,omitempty"`
	UserStatus *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesUserStatus `json:"user_status,omitempty"`
	WeatherStatus *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUserExpireDate `json:"weather_status,omitempty"`
}

// NewCoreOcsGetCapabilities200ResponseOcsDataCapabilities instantiates a new CoreOcsGetCapabilities200ResponseOcsDataCapabilities object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreOcsGetCapabilities200ResponseOcsDataCapabilities() *CoreOcsGetCapabilities200ResponseOcsDataCapabilities {
	this := CoreOcsGetCapabilities200ResponseOcsDataCapabilities{}
	return &this
}

// NewCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesWithDefaults instantiates a new CoreOcsGetCapabilities200ResponseOcsDataCapabilities object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesWithDefaults() *CoreOcsGetCapabilities200ResponseOcsDataCapabilities {
	this := CoreOcsGetCapabilities200ResponseOcsDataCapabilities{}
	return &this
}

// GetOcm returns the Ocm field value if set, zero value otherwise.
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilities) GetOcm() CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcm {
	if o == nil || IsNil(o.Ocm) {
		var ret CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcm
		return ret
	}
	return *o.Ocm
}

// GetOcmOk returns a tuple with the Ocm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilities) GetOcmOk() (*CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcm, bool) {
	if o == nil || IsNil(o.Ocm) {
		return nil, false
	}
	return o.Ocm, true
}

// HasOcm returns a boolean if a field has been set.
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilities) HasOcm() bool {
	if o != nil && !IsNil(o.Ocm) {
		return true
	}

	return false
}

// SetOcm gets a reference to the given CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcm and assigns it to the Ocm field.
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilities) SetOcm(v CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcm) {
	o.Ocm = &v
}

// GetDav returns the Dav field value if set, zero value otherwise.
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilities) GetDav() CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesDav {
	if o == nil || IsNil(o.Dav) {
		var ret CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesDav
		return ret
	}
	return *o.Dav
}

// GetDavOk returns a tuple with the Dav field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilities) GetDavOk() (*CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesDav, bool) {
	if o == nil || IsNil(o.Dav) {
		return nil, false
	}
	return o.Dav, true
}

// HasDav returns a boolean if a field has been set.
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilities) HasDav() bool {
	if o != nil && !IsNil(o.Dav) {
		return true
	}

	return false
}

// SetDav gets a reference to the given CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesDav and assigns it to the Dav field.
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilities) SetDav(v CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesDav) {
	o.Dav = &v
}

// GetFilesSharing returns the FilesSharing field value if set, zero value otherwise.
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilities) GetFilesSharing() CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing {
	if o == nil || IsNil(o.FilesSharing) {
		var ret CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing
		return ret
	}
	return *o.FilesSharing
}

// GetFilesSharingOk returns a tuple with the FilesSharing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilities) GetFilesSharingOk() (*CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing, bool) {
	if o == nil || IsNil(o.FilesSharing) {
		return nil, false
	}
	return o.FilesSharing, true
}

// HasFilesSharing returns a boolean if a field has been set.
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilities) HasFilesSharing() bool {
	if o != nil && !IsNil(o.FilesSharing) {
		return true
	}

	return false
}

// SetFilesSharing gets a reference to the given CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing and assigns it to the FilesSharing field.
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilities) SetFilesSharing(v CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing) {
	o.FilesSharing = &v
}

// GetProvisioningApi returns the ProvisioningApi field value if set, zero value otherwise.
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilities) GetProvisioningApi() CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesProvisioningApi {
	if o == nil || IsNil(o.ProvisioningApi) {
		var ret CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesProvisioningApi
		return ret
	}
	return *o.ProvisioningApi
}

// GetProvisioningApiOk returns a tuple with the ProvisioningApi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilities) GetProvisioningApiOk() (*CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesProvisioningApi, bool) {
	if o == nil || IsNil(o.ProvisioningApi) {
		return nil, false
	}
	return o.ProvisioningApi, true
}

// HasProvisioningApi returns a boolean if a field has been set.
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilities) HasProvisioningApi() bool {
	if o != nil && !IsNil(o.ProvisioningApi) {
		return true
	}

	return false
}

// SetProvisioningApi gets a reference to the given CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesProvisioningApi and assigns it to the ProvisioningApi field.
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilities) SetProvisioningApi(v CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesProvisioningApi) {
	o.ProvisioningApi = &v
}

// GetTheming returns the Theming field value if set, zero value otherwise.
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilities) GetTheming() CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesTheming {
	if o == nil || IsNil(o.Theming) {
		var ret CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesTheming
		return ret
	}
	return *o.Theming
}

// GetThemingOk returns a tuple with the Theming field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilities) GetThemingOk() (*CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesTheming, bool) {
	if o == nil || IsNil(o.Theming) {
		return nil, false
	}
	return o.Theming, true
}

// HasTheming returns a boolean if a field has been set.
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilities) HasTheming() bool {
	if o != nil && !IsNil(o.Theming) {
		return true
	}

	return false
}

// SetTheming gets a reference to the given CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesTheming and assigns it to the Theming field.
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilities) SetTheming(v CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesTheming) {
	o.Theming = &v
}

// GetUserStatus returns the UserStatus field value if set, zero value otherwise.
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilities) GetUserStatus() CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesUserStatus {
	if o == nil || IsNil(o.UserStatus) {
		var ret CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesUserStatus
		return ret
	}
	return *o.UserStatus
}

// GetUserStatusOk returns a tuple with the UserStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilities) GetUserStatusOk() (*CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesUserStatus, bool) {
	if o == nil || IsNil(o.UserStatus) {
		return nil, false
	}
	return o.UserStatus, true
}

// HasUserStatus returns a boolean if a field has been set.
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilities) HasUserStatus() bool {
	if o != nil && !IsNil(o.UserStatus) {
		return true
	}

	return false
}

// SetUserStatus gets a reference to the given CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesUserStatus and assigns it to the UserStatus field.
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilities) SetUserStatus(v CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesUserStatus) {
	o.UserStatus = &v
}

// GetWeatherStatus returns the WeatherStatus field value if set, zero value otherwise.
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilities) GetWeatherStatus() CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUserExpireDate {
	if o == nil || IsNil(o.WeatherStatus) {
		var ret CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUserExpireDate
		return ret
	}
	return *o.WeatherStatus
}

// GetWeatherStatusOk returns a tuple with the WeatherStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilities) GetWeatherStatusOk() (*CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUserExpireDate, bool) {
	if o == nil || IsNil(o.WeatherStatus) {
		return nil, false
	}
	return o.WeatherStatus, true
}

// HasWeatherStatus returns a boolean if a field has been set.
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilities) HasWeatherStatus() bool {
	if o != nil && !IsNil(o.WeatherStatus) {
		return true
	}

	return false
}

// SetWeatherStatus gets a reference to the given CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUserExpireDate and assigns it to the WeatherStatus field.
func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilities) SetWeatherStatus(v CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUserExpireDate) {
	o.WeatherStatus = &v
}

func (o CoreOcsGetCapabilities200ResponseOcsDataCapabilities) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreOcsGetCapabilities200ResponseOcsDataCapabilities) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ocm) {
		toSerialize["ocm"] = o.Ocm
	}
	if !IsNil(o.Dav) {
		toSerialize["dav"] = o.Dav
	}
	if !IsNil(o.FilesSharing) {
		toSerialize["files_sharing"] = o.FilesSharing
	}
	if !IsNil(o.ProvisioningApi) {
		toSerialize["provisioning_api"] = o.ProvisioningApi
	}
	if !IsNil(o.Theming) {
		toSerialize["theming"] = o.Theming
	}
	if !IsNil(o.UserStatus) {
		toSerialize["user_status"] = o.UserStatus
	}
	if !IsNil(o.WeatherStatus) {
		toSerialize["weather_status"] = o.WeatherStatus
	}
	return toSerialize, nil
}

type NullableCoreOcsGetCapabilities200ResponseOcsDataCapabilities struct {
	value *CoreOcsGetCapabilities200ResponseOcsDataCapabilities
	isSet bool
}

func (v NullableCoreOcsGetCapabilities200ResponseOcsDataCapabilities) Get() *CoreOcsGetCapabilities200ResponseOcsDataCapabilities {
	return v.value
}

func (v *NullableCoreOcsGetCapabilities200ResponseOcsDataCapabilities) Set(val *CoreOcsGetCapabilities200ResponseOcsDataCapabilities) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreOcsGetCapabilities200ResponseOcsDataCapabilities) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreOcsGetCapabilities200ResponseOcsDataCapabilities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreOcsGetCapabilities200ResponseOcsDataCapabilities(val *CoreOcsGetCapabilities200ResponseOcsDataCapabilities) *NullableCoreOcsGetCapabilities200ResponseOcsDataCapabilities {
	return &NullableCoreOcsGetCapabilities200ResponseOcsDataCapabilities{value: val, isSet: true}
}

func (v NullableCoreOcsGetCapabilities200ResponseOcsDataCapabilities) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreOcsGetCapabilities200ResponseOcsDataCapabilities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


