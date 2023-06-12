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

// checks if the FilesSharingShareesRecommendedResultExact type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FilesSharingShareesRecommendedResultExact{}

// FilesSharingShareesRecommendedResultExact struct for FilesSharingShareesRecommendedResultExact
type FilesSharingShareesRecommendedResultExact struct {
	Emails []FilesSharingShareeEmail `json:"emails"`
	Groups []FilesSharingSharee `json:"groups"`
	RemoteGroups []FilesSharingShareeRemoteGroup `json:"remote_groups"`
	Remotes []FilesSharingShareeRemote `json:"remotes"`
	Users []FilesSharingShareeUser `json:"users"`
}

// NewFilesSharingShareesRecommendedResultExact instantiates a new FilesSharingShareesRecommendedResultExact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilesSharingShareesRecommendedResultExact(emails []FilesSharingShareeEmail, groups []FilesSharingSharee, remoteGroups []FilesSharingShareeRemoteGroup, remotes []FilesSharingShareeRemote, users []FilesSharingShareeUser) *FilesSharingShareesRecommendedResultExact {
	this := FilesSharingShareesRecommendedResultExact{}
	this.Emails = emails
	this.Groups = groups
	this.RemoteGroups = remoteGroups
	this.Remotes = remotes
	this.Users = users
	return &this
}

// NewFilesSharingShareesRecommendedResultExactWithDefaults instantiates a new FilesSharingShareesRecommendedResultExact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilesSharingShareesRecommendedResultExactWithDefaults() *FilesSharingShareesRecommendedResultExact {
	this := FilesSharingShareesRecommendedResultExact{}
	return &this
}

// GetEmails returns the Emails field value
func (o *FilesSharingShareesRecommendedResultExact) GetEmails() []FilesSharingShareeEmail {
	if o == nil {
		var ret []FilesSharingShareeEmail
		return ret
	}

	return o.Emails
}

// GetEmailsOk returns a tuple with the Emails field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareesRecommendedResultExact) GetEmailsOk() ([]FilesSharingShareeEmail, bool) {
	if o == nil {
		return nil, false
	}
	return o.Emails, true
}

// SetEmails sets field value
func (o *FilesSharingShareesRecommendedResultExact) SetEmails(v []FilesSharingShareeEmail) {
	o.Emails = v
}

// GetGroups returns the Groups field value
func (o *FilesSharingShareesRecommendedResultExact) GetGroups() []FilesSharingSharee {
	if o == nil {
		var ret []FilesSharingSharee
		return ret
	}

	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareesRecommendedResultExact) GetGroupsOk() ([]FilesSharingSharee, bool) {
	if o == nil {
		return nil, false
	}
	return o.Groups, true
}

// SetGroups sets field value
func (o *FilesSharingShareesRecommendedResultExact) SetGroups(v []FilesSharingSharee) {
	o.Groups = v
}

// GetRemoteGroups returns the RemoteGroups field value
func (o *FilesSharingShareesRecommendedResultExact) GetRemoteGroups() []FilesSharingShareeRemoteGroup {
	if o == nil {
		var ret []FilesSharingShareeRemoteGroup
		return ret
	}

	return o.RemoteGroups
}

// GetRemoteGroupsOk returns a tuple with the RemoteGroups field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareesRecommendedResultExact) GetRemoteGroupsOk() ([]FilesSharingShareeRemoteGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.RemoteGroups, true
}

// SetRemoteGroups sets field value
func (o *FilesSharingShareesRecommendedResultExact) SetRemoteGroups(v []FilesSharingShareeRemoteGroup) {
	o.RemoteGroups = v
}

// GetRemotes returns the Remotes field value
func (o *FilesSharingShareesRecommendedResultExact) GetRemotes() []FilesSharingShareeRemote {
	if o == nil {
		var ret []FilesSharingShareeRemote
		return ret
	}

	return o.Remotes
}

// GetRemotesOk returns a tuple with the Remotes field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareesRecommendedResultExact) GetRemotesOk() ([]FilesSharingShareeRemote, bool) {
	if o == nil {
		return nil, false
	}
	return o.Remotes, true
}

// SetRemotes sets field value
func (o *FilesSharingShareesRecommendedResultExact) SetRemotes(v []FilesSharingShareeRemote) {
	o.Remotes = v
}

// GetUsers returns the Users field value
func (o *FilesSharingShareesRecommendedResultExact) GetUsers() []FilesSharingShareeUser {
	if o == nil {
		var ret []FilesSharingShareeUser
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareesRecommendedResultExact) GetUsersOk() ([]FilesSharingShareeUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.Users, true
}

// SetUsers sets field value
func (o *FilesSharingShareesRecommendedResultExact) SetUsers(v []FilesSharingShareeUser) {
	o.Users = v
}

func (o FilesSharingShareesRecommendedResultExact) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FilesSharingShareesRecommendedResultExact) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["emails"] = o.Emails
	toSerialize["groups"] = o.Groups
	toSerialize["remote_groups"] = o.RemoteGroups
	toSerialize["remotes"] = o.Remotes
	toSerialize["users"] = o.Users
	return toSerialize, nil
}

type NullableFilesSharingShareesRecommendedResultExact struct {
	value *FilesSharingShareesRecommendedResultExact
	isSet bool
}

func (v NullableFilesSharingShareesRecommendedResultExact) Get() *FilesSharingShareesRecommendedResultExact {
	return v.value
}

func (v *NullableFilesSharingShareesRecommendedResultExact) Set(val *FilesSharingShareesRecommendedResultExact) {
	v.value = val
	v.isSet = true
}

func (v NullableFilesSharingShareesRecommendedResultExact) IsSet() bool {
	return v.isSet
}

func (v *NullableFilesSharingShareesRecommendedResultExact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilesSharingShareesRecommendedResultExact(val *FilesSharingShareesRecommendedResultExact) *NullableFilesSharingShareesRecommendedResultExact {
	return &NullableFilesSharingShareesRecommendedResultExact{value: val, isSet: true}
}

func (v NullableFilesSharingShareesRecommendedResultExact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilesSharingShareesRecommendedResultExact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


