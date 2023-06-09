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

// checks if the FilesSharingShareesSearchResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FilesSharingShareesSearchResult{}

// FilesSharingShareesSearchResult struct for FilesSharingShareesSearchResult
type FilesSharingShareesSearchResult struct {
	Exact FilesSharingShareesSearchResultExact `json:"exact"`
	Circles []FilesSharingShareeCircle `json:"circles"`
	Emails []FilesSharingShareeEmail `json:"emails"`
	Groups []FilesSharingSharee `json:"groups"`
	Lookup []FilesSharingShareeLookup `json:"lookup"`
	RemoteGroups []FilesSharingShareeRemoteGroup `json:"remote_groups"`
	Remotes []FilesSharingShareeRemote `json:"remotes"`
	Rooms []FilesSharingSharee `json:"rooms"`
	Users []FilesSharingShareeUser `json:"users"`
	LookupEnabled bool `json:"lookupEnabled"`
}

// NewFilesSharingShareesSearchResult instantiates a new FilesSharingShareesSearchResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilesSharingShareesSearchResult(exact FilesSharingShareesSearchResultExact, circles []FilesSharingShareeCircle, emails []FilesSharingShareeEmail, groups []FilesSharingSharee, lookup []FilesSharingShareeLookup, remoteGroups []FilesSharingShareeRemoteGroup, remotes []FilesSharingShareeRemote, rooms []FilesSharingSharee, users []FilesSharingShareeUser, lookupEnabled bool) *FilesSharingShareesSearchResult {
	this := FilesSharingShareesSearchResult{}
	this.Exact = exact
	this.Circles = circles
	this.Emails = emails
	this.Groups = groups
	this.Lookup = lookup
	this.RemoteGroups = remoteGroups
	this.Remotes = remotes
	this.Rooms = rooms
	this.Users = users
	this.LookupEnabled = lookupEnabled
	return &this
}

// NewFilesSharingShareesSearchResultWithDefaults instantiates a new FilesSharingShareesSearchResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilesSharingShareesSearchResultWithDefaults() *FilesSharingShareesSearchResult {
	this := FilesSharingShareesSearchResult{}
	return &this
}

// GetExact returns the Exact field value
func (o *FilesSharingShareesSearchResult) GetExact() FilesSharingShareesSearchResultExact {
	if o == nil {
		var ret FilesSharingShareesSearchResultExact
		return ret
	}

	return o.Exact
}

// GetExactOk returns a tuple with the Exact field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareesSearchResult) GetExactOk() (*FilesSharingShareesSearchResultExact, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Exact, true
}

// SetExact sets field value
func (o *FilesSharingShareesSearchResult) SetExact(v FilesSharingShareesSearchResultExact) {
	o.Exact = v
}

// GetCircles returns the Circles field value
func (o *FilesSharingShareesSearchResult) GetCircles() []FilesSharingShareeCircle {
	if o == nil {
		var ret []FilesSharingShareeCircle
		return ret
	}

	return o.Circles
}

// GetCirclesOk returns a tuple with the Circles field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareesSearchResult) GetCirclesOk() ([]FilesSharingShareeCircle, bool) {
	if o == nil {
		return nil, false
	}
	return o.Circles, true
}

// SetCircles sets field value
func (o *FilesSharingShareesSearchResult) SetCircles(v []FilesSharingShareeCircle) {
	o.Circles = v
}

// GetEmails returns the Emails field value
func (o *FilesSharingShareesSearchResult) GetEmails() []FilesSharingShareeEmail {
	if o == nil {
		var ret []FilesSharingShareeEmail
		return ret
	}

	return o.Emails
}

// GetEmailsOk returns a tuple with the Emails field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareesSearchResult) GetEmailsOk() ([]FilesSharingShareeEmail, bool) {
	if o == nil {
		return nil, false
	}
	return o.Emails, true
}

// SetEmails sets field value
func (o *FilesSharingShareesSearchResult) SetEmails(v []FilesSharingShareeEmail) {
	o.Emails = v
}

// GetGroups returns the Groups field value
func (o *FilesSharingShareesSearchResult) GetGroups() []FilesSharingSharee {
	if o == nil {
		var ret []FilesSharingSharee
		return ret
	}

	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareesSearchResult) GetGroupsOk() ([]FilesSharingSharee, bool) {
	if o == nil {
		return nil, false
	}
	return o.Groups, true
}

// SetGroups sets field value
func (o *FilesSharingShareesSearchResult) SetGroups(v []FilesSharingSharee) {
	o.Groups = v
}

// GetLookup returns the Lookup field value
func (o *FilesSharingShareesSearchResult) GetLookup() []FilesSharingShareeLookup {
	if o == nil {
		var ret []FilesSharingShareeLookup
		return ret
	}

	return o.Lookup
}

// GetLookupOk returns a tuple with the Lookup field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareesSearchResult) GetLookupOk() ([]FilesSharingShareeLookup, bool) {
	if o == nil {
		return nil, false
	}
	return o.Lookup, true
}

// SetLookup sets field value
func (o *FilesSharingShareesSearchResult) SetLookup(v []FilesSharingShareeLookup) {
	o.Lookup = v
}

// GetRemoteGroups returns the RemoteGroups field value
func (o *FilesSharingShareesSearchResult) GetRemoteGroups() []FilesSharingShareeRemoteGroup {
	if o == nil {
		var ret []FilesSharingShareeRemoteGroup
		return ret
	}

	return o.RemoteGroups
}

// GetRemoteGroupsOk returns a tuple with the RemoteGroups field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareesSearchResult) GetRemoteGroupsOk() ([]FilesSharingShareeRemoteGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.RemoteGroups, true
}

// SetRemoteGroups sets field value
func (o *FilesSharingShareesSearchResult) SetRemoteGroups(v []FilesSharingShareeRemoteGroup) {
	o.RemoteGroups = v
}

// GetRemotes returns the Remotes field value
func (o *FilesSharingShareesSearchResult) GetRemotes() []FilesSharingShareeRemote {
	if o == nil {
		var ret []FilesSharingShareeRemote
		return ret
	}

	return o.Remotes
}

// GetRemotesOk returns a tuple with the Remotes field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareesSearchResult) GetRemotesOk() ([]FilesSharingShareeRemote, bool) {
	if o == nil {
		return nil, false
	}
	return o.Remotes, true
}

// SetRemotes sets field value
func (o *FilesSharingShareesSearchResult) SetRemotes(v []FilesSharingShareeRemote) {
	o.Remotes = v
}

// GetRooms returns the Rooms field value
func (o *FilesSharingShareesSearchResult) GetRooms() []FilesSharingSharee {
	if o == nil {
		var ret []FilesSharingSharee
		return ret
	}

	return o.Rooms
}

// GetRoomsOk returns a tuple with the Rooms field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareesSearchResult) GetRoomsOk() ([]FilesSharingSharee, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rooms, true
}

// SetRooms sets field value
func (o *FilesSharingShareesSearchResult) SetRooms(v []FilesSharingSharee) {
	o.Rooms = v
}

// GetUsers returns the Users field value
func (o *FilesSharingShareesSearchResult) GetUsers() []FilesSharingShareeUser {
	if o == nil {
		var ret []FilesSharingShareeUser
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareesSearchResult) GetUsersOk() ([]FilesSharingShareeUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.Users, true
}

// SetUsers sets field value
func (o *FilesSharingShareesSearchResult) SetUsers(v []FilesSharingShareeUser) {
	o.Users = v
}

// GetLookupEnabled returns the LookupEnabled field value
func (o *FilesSharingShareesSearchResult) GetLookupEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.LookupEnabled
}

// GetLookupEnabledOk returns a tuple with the LookupEnabled field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareesSearchResult) GetLookupEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LookupEnabled, true
}

// SetLookupEnabled sets field value
func (o *FilesSharingShareesSearchResult) SetLookupEnabled(v bool) {
	o.LookupEnabled = v
}

func (o FilesSharingShareesSearchResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FilesSharingShareesSearchResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["exact"] = o.Exact
	toSerialize["circles"] = o.Circles
	toSerialize["emails"] = o.Emails
	toSerialize["groups"] = o.Groups
	toSerialize["lookup"] = o.Lookup
	toSerialize["remote_groups"] = o.RemoteGroups
	toSerialize["remotes"] = o.Remotes
	toSerialize["rooms"] = o.Rooms
	toSerialize["users"] = o.Users
	toSerialize["lookupEnabled"] = o.LookupEnabled
	return toSerialize, nil
}

type NullableFilesSharingShareesSearchResult struct {
	value *FilesSharingShareesSearchResult
	isSet bool
}

func (v NullableFilesSharingShareesSearchResult) Get() *FilesSharingShareesSearchResult {
	return v.value
}

func (v *NullableFilesSharingShareesSearchResult) Set(val *FilesSharingShareesSearchResult) {
	v.value = val
	v.isSet = true
}

func (v NullableFilesSharingShareesSearchResult) IsSet() bool {
	return v.isSet
}

func (v *NullableFilesSharingShareesSearchResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilesSharingShareesSearchResult(val *FilesSharingShareesSearchResult) *NullableFilesSharingShareesSearchResult {
	return &NullableFilesSharingShareesSearchResult{value: val, isSet: true}
}

func (v NullableFilesSharingShareesSearchResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilesSharingShareesSearchResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


