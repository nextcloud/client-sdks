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

// checks if the CoreUnifiedSearchResultEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CoreUnifiedSearchResultEntry{}

// CoreUnifiedSearchResultEntry struct for CoreUnifiedSearchResultEntry
type CoreUnifiedSearchResultEntry struct {
	ThumbnailUrl string `json:"thumbnailUrl"`
	Title string `json:"title"`
	Subline string `json:"subline"`
	ResourceUrl string `json:"resourceUrl"`
	Icon string `json:"icon"`
	Rounded bool `json:"rounded"`
	Attributes []string `json:"attributes"`
}

// NewCoreUnifiedSearchResultEntry instantiates a new CoreUnifiedSearchResultEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreUnifiedSearchResultEntry(thumbnailUrl string, title string, subline string, resourceUrl string, icon string, rounded bool, attributes []string) *CoreUnifiedSearchResultEntry {
	this := CoreUnifiedSearchResultEntry{}
	this.ThumbnailUrl = thumbnailUrl
	this.Title = title
	this.Subline = subline
	this.ResourceUrl = resourceUrl
	this.Icon = icon
	this.Rounded = rounded
	this.Attributes = attributes
	return &this
}

// NewCoreUnifiedSearchResultEntryWithDefaults instantiates a new CoreUnifiedSearchResultEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreUnifiedSearchResultEntryWithDefaults() *CoreUnifiedSearchResultEntry {
	this := CoreUnifiedSearchResultEntry{}
	return &this
}

// GetThumbnailUrl returns the ThumbnailUrl field value
func (o *CoreUnifiedSearchResultEntry) GetThumbnailUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThumbnailUrl
}

// GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field value
// and a boolean to check if the value has been set.
func (o *CoreUnifiedSearchResultEntry) GetThumbnailUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThumbnailUrl, true
}

// SetThumbnailUrl sets field value
func (o *CoreUnifiedSearchResultEntry) SetThumbnailUrl(v string) {
	o.ThumbnailUrl = v
}

// GetTitle returns the Title field value
func (o *CoreUnifiedSearchResultEntry) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *CoreUnifiedSearchResultEntry) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *CoreUnifiedSearchResultEntry) SetTitle(v string) {
	o.Title = v
}

// GetSubline returns the Subline field value
func (o *CoreUnifiedSearchResultEntry) GetSubline() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subline
}

// GetSublineOk returns a tuple with the Subline field value
// and a boolean to check if the value has been set.
func (o *CoreUnifiedSearchResultEntry) GetSublineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subline, true
}

// SetSubline sets field value
func (o *CoreUnifiedSearchResultEntry) SetSubline(v string) {
	o.Subline = v
}

// GetResourceUrl returns the ResourceUrl field value
func (o *CoreUnifiedSearchResultEntry) GetResourceUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceUrl
}

// GetResourceUrlOk returns a tuple with the ResourceUrl field value
// and a boolean to check if the value has been set.
func (o *CoreUnifiedSearchResultEntry) GetResourceUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceUrl, true
}

// SetResourceUrl sets field value
func (o *CoreUnifiedSearchResultEntry) SetResourceUrl(v string) {
	o.ResourceUrl = v
}

// GetIcon returns the Icon field value
func (o *CoreUnifiedSearchResultEntry) GetIcon() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Icon
}

// GetIconOk returns a tuple with the Icon field value
// and a boolean to check if the value has been set.
func (o *CoreUnifiedSearchResultEntry) GetIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Icon, true
}

// SetIcon sets field value
func (o *CoreUnifiedSearchResultEntry) SetIcon(v string) {
	o.Icon = v
}

// GetRounded returns the Rounded field value
func (o *CoreUnifiedSearchResultEntry) GetRounded() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Rounded
}

// GetRoundedOk returns a tuple with the Rounded field value
// and a boolean to check if the value has been set.
func (o *CoreUnifiedSearchResultEntry) GetRoundedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rounded, true
}

// SetRounded sets field value
func (o *CoreUnifiedSearchResultEntry) SetRounded(v bool) {
	o.Rounded = v
}

// GetAttributes returns the Attributes field value
func (o *CoreUnifiedSearchResultEntry) GetAttributes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CoreUnifiedSearchResultEntry) GetAttributesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attributes, true
}

// SetAttributes sets field value
func (o *CoreUnifiedSearchResultEntry) SetAttributes(v []string) {
	o.Attributes = v
}

func (o CoreUnifiedSearchResultEntry) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreUnifiedSearchResultEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["thumbnailUrl"] = o.ThumbnailUrl
	toSerialize["title"] = o.Title
	toSerialize["subline"] = o.Subline
	toSerialize["resourceUrl"] = o.ResourceUrl
	toSerialize["icon"] = o.Icon
	toSerialize["rounded"] = o.Rounded
	toSerialize["attributes"] = o.Attributes
	return toSerialize, nil
}

type NullableCoreUnifiedSearchResultEntry struct {
	value *CoreUnifiedSearchResultEntry
	isSet bool
}

func (v NullableCoreUnifiedSearchResultEntry) Get() *CoreUnifiedSearchResultEntry {
	return v.value
}

func (v *NullableCoreUnifiedSearchResultEntry) Set(val *CoreUnifiedSearchResultEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreUnifiedSearchResultEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreUnifiedSearchResultEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreUnifiedSearchResultEntry(val *CoreUnifiedSearchResultEntry) *NullableCoreUnifiedSearchResultEntry {
	return &NullableCoreUnifiedSearchResultEntry{value: val, isSet: true}
}

func (v NullableCoreUnifiedSearchResultEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreUnifiedSearchResultEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


