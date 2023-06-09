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

// checks if the FilesSharingDeletedShare type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FilesSharingDeletedShare{}

// FilesSharingDeletedShare struct for FilesSharingDeletedShare
type FilesSharingDeletedShare struct {
	Id string `json:"id"`
	ShareType int64 `json:"share_type"`
	UidOwner string `json:"uid_owner"`
	DisplaynameOwner string `json:"displayname_owner"`
	Permissions int64 `json:"permissions"`
	Stime int64 `json:"stime"`
	UidFileOwner string `json:"uid_file_owner"`
	DisplaynameFileOwner string `json:"displayname_file_owner"`
	Path string `json:"path"`
	ItemType string `json:"item_type"`
	Mimetype string `json:"mimetype"`
	Storage int64 `json:"storage"`
	ItemSource int64 `json:"item_source"`
	FileSource int64 `json:"file_source"`
	FileParent int64 `json:"file_parent"`
	FileTarget int64 `json:"file_target"`
	Expiration NullableString `json:"expiration"`
	ShareWith NullableString `json:"share_with"`
	ShareWithDisplayname NullableString `json:"share_with_displayname"`
	ShareWithLink NullableString `json:"share_with_link"`
}

// NewFilesSharingDeletedShare instantiates a new FilesSharingDeletedShare object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilesSharingDeletedShare(id string, shareType int64, uidOwner string, displaynameOwner string, permissions int64, stime int64, uidFileOwner string, displaynameFileOwner string, path string, itemType string, mimetype string, storage int64, itemSource int64, fileSource int64, fileParent int64, fileTarget int64, expiration NullableString, shareWith NullableString, shareWithDisplayname NullableString, shareWithLink NullableString) *FilesSharingDeletedShare {
	this := FilesSharingDeletedShare{}
	this.Id = id
	this.ShareType = shareType
	this.UidOwner = uidOwner
	this.DisplaynameOwner = displaynameOwner
	this.Permissions = permissions
	this.Stime = stime
	this.UidFileOwner = uidFileOwner
	this.DisplaynameFileOwner = displaynameFileOwner
	this.Path = path
	this.ItemType = itemType
	this.Mimetype = mimetype
	this.Storage = storage
	this.ItemSource = itemSource
	this.FileSource = fileSource
	this.FileParent = fileParent
	this.FileTarget = fileTarget
	this.Expiration = expiration
	this.ShareWith = shareWith
	this.ShareWithDisplayname = shareWithDisplayname
	this.ShareWithLink = shareWithLink
	return &this
}

// NewFilesSharingDeletedShareWithDefaults instantiates a new FilesSharingDeletedShare object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilesSharingDeletedShareWithDefaults() *FilesSharingDeletedShare {
	this := FilesSharingDeletedShare{}
	return &this
}

// GetId returns the Id field value
func (o *FilesSharingDeletedShare) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FilesSharingDeletedShare) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FilesSharingDeletedShare) SetId(v string) {
	o.Id = v
}

// GetShareType returns the ShareType field value
func (o *FilesSharingDeletedShare) GetShareType() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ShareType
}

// GetShareTypeOk returns a tuple with the ShareType field value
// and a boolean to check if the value has been set.
func (o *FilesSharingDeletedShare) GetShareTypeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShareType, true
}

// SetShareType sets field value
func (o *FilesSharingDeletedShare) SetShareType(v int64) {
	o.ShareType = v
}

// GetUidOwner returns the UidOwner field value
func (o *FilesSharingDeletedShare) GetUidOwner() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UidOwner
}

// GetUidOwnerOk returns a tuple with the UidOwner field value
// and a boolean to check if the value has been set.
func (o *FilesSharingDeletedShare) GetUidOwnerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UidOwner, true
}

// SetUidOwner sets field value
func (o *FilesSharingDeletedShare) SetUidOwner(v string) {
	o.UidOwner = v
}

// GetDisplaynameOwner returns the DisplaynameOwner field value
func (o *FilesSharingDeletedShare) GetDisplaynameOwner() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplaynameOwner
}

// GetDisplaynameOwnerOk returns a tuple with the DisplaynameOwner field value
// and a boolean to check if the value has been set.
func (o *FilesSharingDeletedShare) GetDisplaynameOwnerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplaynameOwner, true
}

// SetDisplaynameOwner sets field value
func (o *FilesSharingDeletedShare) SetDisplaynameOwner(v string) {
	o.DisplaynameOwner = v
}

// GetPermissions returns the Permissions field value
func (o *FilesSharingDeletedShare) GetPermissions() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
func (o *FilesSharingDeletedShare) GetPermissionsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Permissions, true
}

// SetPermissions sets field value
func (o *FilesSharingDeletedShare) SetPermissions(v int64) {
	o.Permissions = v
}

// GetStime returns the Stime field value
func (o *FilesSharingDeletedShare) GetStime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Stime
}

// GetStimeOk returns a tuple with the Stime field value
// and a boolean to check if the value has been set.
func (o *FilesSharingDeletedShare) GetStimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stime, true
}

// SetStime sets field value
func (o *FilesSharingDeletedShare) SetStime(v int64) {
	o.Stime = v
}

// GetUidFileOwner returns the UidFileOwner field value
func (o *FilesSharingDeletedShare) GetUidFileOwner() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UidFileOwner
}

// GetUidFileOwnerOk returns a tuple with the UidFileOwner field value
// and a boolean to check if the value has been set.
func (o *FilesSharingDeletedShare) GetUidFileOwnerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UidFileOwner, true
}

// SetUidFileOwner sets field value
func (o *FilesSharingDeletedShare) SetUidFileOwner(v string) {
	o.UidFileOwner = v
}

// GetDisplaynameFileOwner returns the DisplaynameFileOwner field value
func (o *FilesSharingDeletedShare) GetDisplaynameFileOwner() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplaynameFileOwner
}

// GetDisplaynameFileOwnerOk returns a tuple with the DisplaynameFileOwner field value
// and a boolean to check if the value has been set.
func (o *FilesSharingDeletedShare) GetDisplaynameFileOwnerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplaynameFileOwner, true
}

// SetDisplaynameFileOwner sets field value
func (o *FilesSharingDeletedShare) SetDisplaynameFileOwner(v string) {
	o.DisplaynameFileOwner = v
}

// GetPath returns the Path field value
func (o *FilesSharingDeletedShare) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *FilesSharingDeletedShare) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *FilesSharingDeletedShare) SetPath(v string) {
	o.Path = v
}

// GetItemType returns the ItemType field value
func (o *FilesSharingDeletedShare) GetItemType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ItemType
}

// GetItemTypeOk returns a tuple with the ItemType field value
// and a boolean to check if the value has been set.
func (o *FilesSharingDeletedShare) GetItemTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ItemType, true
}

// SetItemType sets field value
func (o *FilesSharingDeletedShare) SetItemType(v string) {
	o.ItemType = v
}

// GetMimetype returns the Mimetype field value
func (o *FilesSharingDeletedShare) GetMimetype() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mimetype
}

// GetMimetypeOk returns a tuple with the Mimetype field value
// and a boolean to check if the value has been set.
func (o *FilesSharingDeletedShare) GetMimetypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mimetype, true
}

// SetMimetype sets field value
func (o *FilesSharingDeletedShare) SetMimetype(v string) {
	o.Mimetype = v
}

// GetStorage returns the Storage field value
func (o *FilesSharingDeletedShare) GetStorage() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value
// and a boolean to check if the value has been set.
func (o *FilesSharingDeletedShare) GetStorageOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Storage, true
}

// SetStorage sets field value
func (o *FilesSharingDeletedShare) SetStorage(v int64) {
	o.Storage = v
}

// GetItemSource returns the ItemSource field value
func (o *FilesSharingDeletedShare) GetItemSource() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ItemSource
}

// GetItemSourceOk returns a tuple with the ItemSource field value
// and a boolean to check if the value has been set.
func (o *FilesSharingDeletedShare) GetItemSourceOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ItemSource, true
}

// SetItemSource sets field value
func (o *FilesSharingDeletedShare) SetItemSource(v int64) {
	o.ItemSource = v
}

// GetFileSource returns the FileSource field value
func (o *FilesSharingDeletedShare) GetFileSource() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.FileSource
}

// GetFileSourceOk returns a tuple with the FileSource field value
// and a boolean to check if the value has been set.
func (o *FilesSharingDeletedShare) GetFileSourceOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileSource, true
}

// SetFileSource sets field value
func (o *FilesSharingDeletedShare) SetFileSource(v int64) {
	o.FileSource = v
}

// GetFileParent returns the FileParent field value
func (o *FilesSharingDeletedShare) GetFileParent() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.FileParent
}

// GetFileParentOk returns a tuple with the FileParent field value
// and a boolean to check if the value has been set.
func (o *FilesSharingDeletedShare) GetFileParentOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileParent, true
}

// SetFileParent sets field value
func (o *FilesSharingDeletedShare) SetFileParent(v int64) {
	o.FileParent = v
}

// GetFileTarget returns the FileTarget field value
func (o *FilesSharingDeletedShare) GetFileTarget() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.FileTarget
}

// GetFileTargetOk returns a tuple with the FileTarget field value
// and a boolean to check if the value has been set.
func (o *FilesSharingDeletedShare) GetFileTargetOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileTarget, true
}

// SetFileTarget sets field value
func (o *FilesSharingDeletedShare) SetFileTarget(v int64) {
	o.FileTarget = v
}

// GetExpiration returns the Expiration field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FilesSharingDeletedShare) GetExpiration() string {
	if o == nil || o.Expiration.Get() == nil {
		var ret string
		return ret
	}

	return *o.Expiration.Get()
}

// GetExpirationOk returns a tuple with the Expiration field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesSharingDeletedShare) GetExpirationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expiration.Get(), o.Expiration.IsSet()
}

// SetExpiration sets field value
func (o *FilesSharingDeletedShare) SetExpiration(v string) {
	o.Expiration.Set(&v)
}

// GetShareWith returns the ShareWith field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FilesSharingDeletedShare) GetShareWith() string {
	if o == nil || o.ShareWith.Get() == nil {
		var ret string
		return ret
	}

	return *o.ShareWith.Get()
}

// GetShareWithOk returns a tuple with the ShareWith field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesSharingDeletedShare) GetShareWithOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShareWith.Get(), o.ShareWith.IsSet()
}

// SetShareWith sets field value
func (o *FilesSharingDeletedShare) SetShareWith(v string) {
	o.ShareWith.Set(&v)
}

// GetShareWithDisplayname returns the ShareWithDisplayname field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FilesSharingDeletedShare) GetShareWithDisplayname() string {
	if o == nil || o.ShareWithDisplayname.Get() == nil {
		var ret string
		return ret
	}

	return *o.ShareWithDisplayname.Get()
}

// GetShareWithDisplaynameOk returns a tuple with the ShareWithDisplayname field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesSharingDeletedShare) GetShareWithDisplaynameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShareWithDisplayname.Get(), o.ShareWithDisplayname.IsSet()
}

// SetShareWithDisplayname sets field value
func (o *FilesSharingDeletedShare) SetShareWithDisplayname(v string) {
	o.ShareWithDisplayname.Set(&v)
}

// GetShareWithLink returns the ShareWithLink field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FilesSharingDeletedShare) GetShareWithLink() string {
	if o == nil || o.ShareWithLink.Get() == nil {
		var ret string
		return ret
	}

	return *o.ShareWithLink.Get()
}

// GetShareWithLinkOk returns a tuple with the ShareWithLink field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesSharingDeletedShare) GetShareWithLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShareWithLink.Get(), o.ShareWithLink.IsSet()
}

// SetShareWithLink sets field value
func (o *FilesSharingDeletedShare) SetShareWithLink(v string) {
	o.ShareWithLink.Set(&v)
}

func (o FilesSharingDeletedShare) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FilesSharingDeletedShare) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["share_type"] = o.ShareType
	toSerialize["uid_owner"] = o.UidOwner
	toSerialize["displayname_owner"] = o.DisplaynameOwner
	toSerialize["permissions"] = o.Permissions
	toSerialize["stime"] = o.Stime
	toSerialize["uid_file_owner"] = o.UidFileOwner
	toSerialize["displayname_file_owner"] = o.DisplaynameFileOwner
	toSerialize["path"] = o.Path
	toSerialize["item_type"] = o.ItemType
	toSerialize["mimetype"] = o.Mimetype
	toSerialize["storage"] = o.Storage
	toSerialize["item_source"] = o.ItemSource
	toSerialize["file_source"] = o.FileSource
	toSerialize["file_parent"] = o.FileParent
	toSerialize["file_target"] = o.FileTarget
	toSerialize["expiration"] = o.Expiration.Get()
	toSerialize["share_with"] = o.ShareWith.Get()
	toSerialize["share_with_displayname"] = o.ShareWithDisplayname.Get()
	toSerialize["share_with_link"] = o.ShareWithLink.Get()
	return toSerialize, nil
}

type NullableFilesSharingDeletedShare struct {
	value *FilesSharingDeletedShare
	isSet bool
}

func (v NullableFilesSharingDeletedShare) Get() *FilesSharingDeletedShare {
	return v.value
}

func (v *NullableFilesSharingDeletedShare) Set(val *FilesSharingDeletedShare) {
	v.value = val
	v.isSet = true
}

func (v NullableFilesSharingDeletedShare) IsSet() bool {
	return v.isSet
}

func (v *NullableFilesSharingDeletedShare) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilesSharingDeletedShare(val *FilesSharingDeletedShare) *NullableFilesSharingDeletedShare {
	return &NullableFilesSharingDeletedShare{value: val, isSet: true}
}

func (v NullableFilesSharingDeletedShare) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilesSharingDeletedShare) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


