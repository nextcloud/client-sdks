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

// checks if the FilesSharingShareesRecommendedResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FilesSharingShareesRecommendedResult{}

// FilesSharingShareesRecommendedResult struct for FilesSharingShareesRecommendedResult
type FilesSharingShareesRecommendedResult struct {
	Exact FilesSharingShareesRecommendedResultExact `json:"exact"`
	Emails []FilesSharingShareeEmail `json:"emails"`
	Groups []FilesSharingSharee `json:"groups"`
	RemoteGroups []FilesSharingShareeRemoteGroup `json:"remote_groups"`
	Remotes []FilesSharingShareeRemote `json:"remotes"`
	Users []FilesSharingShareeUser `json:"users"`
}

// NewFilesSharingShareesRecommendedResult instantiates a new FilesSharingShareesRecommendedResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilesSharingShareesRecommendedResult(exact FilesSharingShareesRecommendedResultExact, emails []FilesSharingShareeEmail, groups []FilesSharingSharee, remoteGroups []FilesSharingShareeRemoteGroup, remotes []FilesSharingShareeRemote, users []FilesSharingShareeUser) *FilesSharingShareesRecommendedResult {
	this := FilesSharingShareesRecommendedResult{}
	this.Exact = exact
	this.Emails = emails
	this.Groups = groups
	this.RemoteGroups = remoteGroups
	this.Remotes = remotes
	this.Users = users
	return &this
}

// NewFilesSharingShareesRecommendedResultWithDefaults instantiates a new FilesSharingShareesRecommendedResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilesSharingShareesRecommendedResultWithDefaults() *FilesSharingShareesRecommendedResult {
	this := FilesSharingShareesRecommendedResult{}
	return &this
}

// GetExact returns the Exact field value
func (o *FilesSharingShareesRecommendedResult) GetExact() FilesSharingShareesRecommendedResultExact {
	if o == nil {
		var ret FilesSharingShareesRecommendedResultExact
		return ret
	}

	return o.Exact
}

// GetExactOk returns a tuple with the Exact field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareesRecommendedResult) GetExactOk() (*FilesSharingShareesRecommendedResultExact, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Exact, true
}

// SetExact sets field value
func (o *FilesSharingShareesRecommendedResult) SetExact(v FilesSharingShareesRecommendedResultExact) {
	o.Exact = v
}

// GetEmails returns the Emails field value
func (o *FilesSharingShareesRecommendedResult) GetEmails() []FilesSharingShareeEmail {
	if o == nil {
		var ret []FilesSharingShareeEmail
		return ret
	}

	return o.Emails
}

// GetEmailsOk returns a tuple with the Emails field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareesRecommendedResult) GetEmailsOk() ([]FilesSharingShareeEmail, bool) {
	if o == nil {
		return nil, false
	}
	return o.Emails, true
}

// SetEmails sets field value
func (o *FilesSharingShareesRecommendedResult) SetEmails(v []FilesSharingShareeEmail) {
	o.Emails = v
}

// GetGroups returns the Groups field value
func (o *FilesSharingShareesRecommendedResult) GetGroups() []FilesSharingSharee {
	if o == nil {
		var ret []FilesSharingSharee
		return ret
	}

	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareesRecommendedResult) GetGroupsOk() ([]FilesSharingSharee, bool) {
	if o == nil {
		return nil, false
	}
	return o.Groups, true
}

// SetGroups sets field value
func (o *FilesSharingShareesRecommendedResult) SetGroups(v []FilesSharingSharee) {
	o.Groups = v
}

// GetRemoteGroups returns the RemoteGroups field value
func (o *FilesSharingShareesRecommendedResult) GetRemoteGroups() []FilesSharingShareeRemoteGroup {
	if o == nil {
		var ret []FilesSharingShareeRemoteGroup
		return ret
	}

	return o.RemoteGroups
}

// GetRemoteGroupsOk returns a tuple with the RemoteGroups field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareesRecommendedResult) GetRemoteGroupsOk() ([]FilesSharingShareeRemoteGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.RemoteGroups, true
}

// SetRemoteGroups sets field value
func (o *FilesSharingShareesRecommendedResult) SetRemoteGroups(v []FilesSharingShareeRemoteGroup) {
	o.RemoteGroups = v
}

// GetRemotes returns the Remotes field value
func (o *FilesSharingShareesRecommendedResult) GetRemotes() []FilesSharingShareeRemote {
	if o == nil {
		var ret []FilesSharingShareeRemote
		return ret
	}

	return o.Remotes
}

// GetRemotesOk returns a tuple with the Remotes field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareesRecommendedResult) GetRemotesOk() ([]FilesSharingShareeRemote, bool) {
	if o == nil {
		return nil, false
	}
	return o.Remotes, true
}

// SetRemotes sets field value
func (o *FilesSharingShareesRecommendedResult) SetRemotes(v []FilesSharingShareeRemote) {
	o.Remotes = v
}

// GetUsers returns the Users field value
func (o *FilesSharingShareesRecommendedResult) GetUsers() []FilesSharingShareeUser {
	if o == nil {
		var ret []FilesSharingShareeUser
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareesRecommendedResult) GetUsersOk() ([]FilesSharingShareeUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.Users, true
}

// SetUsers sets field value
func (o *FilesSharingShareesRecommendedResult) SetUsers(v []FilesSharingShareeUser) {
	o.Users = v
}

func (o FilesSharingShareesRecommendedResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FilesSharingShareesRecommendedResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["exact"] = o.Exact
	toSerialize["emails"] = o.Emails
	toSerialize["groups"] = o.Groups
	toSerialize["remote_groups"] = o.RemoteGroups
	toSerialize["remotes"] = o.Remotes
	toSerialize["users"] = o.Users
	return toSerialize, nil
}

type NullableFilesSharingShareesRecommendedResult struct {
	value *FilesSharingShareesRecommendedResult
	isSet bool
}

func (v NullableFilesSharingShareesRecommendedResult) Get() *FilesSharingShareesRecommendedResult {
	return v.value
}

func (v *NullableFilesSharingShareesRecommendedResult) Set(val *FilesSharingShareesRecommendedResult) {
	v.value = val
	v.isSet = true
}

func (v NullableFilesSharingShareesRecommendedResult) IsSet() bool {
	return v.isSet
}

func (v *NullableFilesSharingShareesRecommendedResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilesSharingShareesRecommendedResult(val *FilesSharingShareesRecommendedResult) *NullableFilesSharingShareesRecommendedResult {
	return &NullableFilesSharingShareesRecommendedResult{value: val, isSet: true}
}

func (v NullableFilesSharingShareesRecommendedResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilesSharingShareesRecommendedResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


