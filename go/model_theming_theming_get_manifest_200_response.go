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

// checks if the ThemingThemingGetManifest200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThemingThemingGetManifest200Response{}

// ThemingThemingGetManifest200Response struct for ThemingThemingGetManifest200Response
type ThemingThemingGetManifest200Response struct {
	Name string `json:"name"`
	ShortName string `json:"short_name"`
	StartUrl string `json:"start_url"`
	ThemeColor string `json:"theme_color"`
	BackgroundColor string `json:"background_color"`
	Description string `json:"description"`
	Icons []ThemingThemingGetManifest200ResponseIconsInner `json:"icons"`
	Display string `json:"display"`
}

// NewThemingThemingGetManifest200Response instantiates a new ThemingThemingGetManifest200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThemingThemingGetManifest200Response(name string, shortName string, startUrl string, themeColor string, backgroundColor string, description string, icons []ThemingThemingGetManifest200ResponseIconsInner, display string) *ThemingThemingGetManifest200Response {
	this := ThemingThemingGetManifest200Response{}
	this.Name = name
	this.ShortName = shortName
	this.StartUrl = startUrl
	this.ThemeColor = themeColor
	this.BackgroundColor = backgroundColor
	this.Description = description
	this.Icons = icons
	this.Display = display
	return &this
}

// NewThemingThemingGetManifest200ResponseWithDefaults instantiates a new ThemingThemingGetManifest200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThemingThemingGetManifest200ResponseWithDefaults() *ThemingThemingGetManifest200Response {
	this := ThemingThemingGetManifest200Response{}
	return &this
}

// GetName returns the Name field value
func (o *ThemingThemingGetManifest200Response) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ThemingThemingGetManifest200Response) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ThemingThemingGetManifest200Response) SetName(v string) {
	o.Name = v
}

// GetShortName returns the ShortName field value
func (o *ThemingThemingGetManifest200Response) GetShortName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value
// and a boolean to check if the value has been set.
func (o *ThemingThemingGetManifest200Response) GetShortNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShortName, true
}

// SetShortName sets field value
func (o *ThemingThemingGetManifest200Response) SetShortName(v string) {
	o.ShortName = v
}

// GetStartUrl returns the StartUrl field value
func (o *ThemingThemingGetManifest200Response) GetStartUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartUrl
}

// GetStartUrlOk returns a tuple with the StartUrl field value
// and a boolean to check if the value has been set.
func (o *ThemingThemingGetManifest200Response) GetStartUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartUrl, true
}

// SetStartUrl sets field value
func (o *ThemingThemingGetManifest200Response) SetStartUrl(v string) {
	o.StartUrl = v
}

// GetThemeColor returns the ThemeColor field value
func (o *ThemingThemingGetManifest200Response) GetThemeColor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThemeColor
}

// GetThemeColorOk returns a tuple with the ThemeColor field value
// and a boolean to check if the value has been set.
func (o *ThemingThemingGetManifest200Response) GetThemeColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThemeColor, true
}

// SetThemeColor sets field value
func (o *ThemingThemingGetManifest200Response) SetThemeColor(v string) {
	o.ThemeColor = v
}

// GetBackgroundColor returns the BackgroundColor field value
func (o *ThemingThemingGetManifest200Response) GetBackgroundColor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BackgroundColor
}

// GetBackgroundColorOk returns a tuple with the BackgroundColor field value
// and a boolean to check if the value has been set.
func (o *ThemingThemingGetManifest200Response) GetBackgroundColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BackgroundColor, true
}

// SetBackgroundColor sets field value
func (o *ThemingThemingGetManifest200Response) SetBackgroundColor(v string) {
	o.BackgroundColor = v
}

// GetDescription returns the Description field value
func (o *ThemingThemingGetManifest200Response) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ThemingThemingGetManifest200Response) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ThemingThemingGetManifest200Response) SetDescription(v string) {
	o.Description = v
}

// GetIcons returns the Icons field value
func (o *ThemingThemingGetManifest200Response) GetIcons() []ThemingThemingGetManifest200ResponseIconsInner {
	if o == nil {
		var ret []ThemingThemingGetManifest200ResponseIconsInner
		return ret
	}

	return o.Icons
}

// GetIconsOk returns a tuple with the Icons field value
// and a boolean to check if the value has been set.
func (o *ThemingThemingGetManifest200Response) GetIconsOk() ([]ThemingThemingGetManifest200ResponseIconsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Icons, true
}

// SetIcons sets field value
func (o *ThemingThemingGetManifest200Response) SetIcons(v []ThemingThemingGetManifest200ResponseIconsInner) {
	o.Icons = v
}

// GetDisplay returns the Display field value
func (o *ThemingThemingGetManifest200Response) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *ThemingThemingGetManifest200Response) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *ThemingThemingGetManifest200Response) SetDisplay(v string) {
	o.Display = v
}

func (o ThemingThemingGetManifest200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThemingThemingGetManifest200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["short_name"] = o.ShortName
	toSerialize["start_url"] = o.StartUrl
	toSerialize["theme_color"] = o.ThemeColor
	toSerialize["background_color"] = o.BackgroundColor
	toSerialize["description"] = o.Description
	toSerialize["icons"] = o.Icons
	toSerialize["display"] = o.Display
	return toSerialize, nil
}

type NullableThemingThemingGetManifest200Response struct {
	value *ThemingThemingGetManifest200Response
	isSet bool
}

func (v NullableThemingThemingGetManifest200Response) Get() *ThemingThemingGetManifest200Response {
	return v.value
}

func (v *NullableThemingThemingGetManifest200Response) Set(val *ThemingThemingGetManifest200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableThemingThemingGetManifest200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableThemingThemingGetManifest200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThemingThemingGetManifest200Response(val *ThemingThemingGetManifest200Response) *NullableThemingThemingGetManifest200Response {
	return &NullableThemingThemingGetManifest200Response{value: val, isSet: true}
}

func (v NullableThemingThemingGetManifest200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThemingThemingGetManifest200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


