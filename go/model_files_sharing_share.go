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

// checks if the FilesSharingShare type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FilesSharingShare{}

// FilesSharingShare struct for FilesSharingShare
type FilesSharingShare struct {
	Attributes NullableString `json:"attributes"`
	CanDelete bool `json:"can_delete"`
	CanEdit bool `json:"can_edit"`
	DisplaynameFileOwner string `json:"displayname_file_owner"`
	DisplaynameOwner string `json:"displayname_owner"`
	Expiration NullableString `json:"expiration"`
	FileParent int64 `json:"file_parent"`
	FileSource int64 `json:"file_source"`
	FileTarget string `json:"file_target"`
	HasPreview bool `json:"has_preview"`
	Id string `json:"id"`
	ItemSource int64 `json:"item_source"`
	ItemType string `json:"item_type"`
	Label string `json:"label"`
	MailSend int64 `json:"mail_send"`
	Mimetype string `json:"mimetype"`
	Note string `json:"note"`
	Password NullableString `json:"password"`
	PasswordExpirationTime NullableString `json:"password_expiration_time"`
	Path string `json:"path"`
	Permissions int64 `json:"permissions"`
	SendPasswordByTalk NullableBool `json:"send_password_by_talk"`
	ShareType int64 `json:"share_type"`
	ShareWith NullableString `json:"share_with"`
	ShareWithAvatar NullableString `json:"share_with_avatar"`
	ShareWithDisplayname NullableString `json:"share_with_displayname"`
	ShareWithLink NullableString `json:"share_with_link"`
	Status NullableFilesSharingShareStatus `json:"status"`
	Stime int64 `json:"stime"`
	Storage int64 `json:"storage"`
	StorageId string `json:"storage_id"`
	Token NullableString `json:"token"`
	UidFileOwner string `json:"uid_file_owner"`
	UidOwner string `json:"uid_owner"`
	Url NullableString `json:"url"`
}

// NewFilesSharingShare instantiates a new FilesSharingShare object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilesSharingShare(attributes NullableString, canDelete bool, canEdit bool, displaynameFileOwner string, displaynameOwner string, expiration NullableString, fileParent int64, fileSource int64, fileTarget string, hasPreview bool, id string, itemSource int64, itemType string, label string, mailSend int64, mimetype string, note string, password NullableString, passwordExpirationTime NullableString, path string, permissions int64, sendPasswordByTalk NullableBool, shareType int64, shareWith NullableString, shareWithAvatar NullableString, shareWithDisplayname NullableString, shareWithLink NullableString, status NullableFilesSharingShareStatus, stime int64, storage int64, storageId string, token NullableString, uidFileOwner string, uidOwner string, url NullableString) *FilesSharingShare {
	this := FilesSharingShare{}
	this.Attributes = attributes
	this.CanDelete = canDelete
	this.CanEdit = canEdit
	this.DisplaynameFileOwner = displaynameFileOwner
	this.DisplaynameOwner = displaynameOwner
	this.Expiration = expiration
	this.FileParent = fileParent
	this.FileSource = fileSource
	this.FileTarget = fileTarget
	this.HasPreview = hasPreview
	this.Id = id
	this.ItemSource = itemSource
	this.ItemType = itemType
	this.Label = label
	this.MailSend = mailSend
	this.Mimetype = mimetype
	this.Note = note
	this.Password = password
	this.PasswordExpirationTime = passwordExpirationTime
	this.Path = path
	this.Permissions = permissions
	this.SendPasswordByTalk = sendPasswordByTalk
	this.ShareType = shareType
	this.ShareWith = shareWith
	this.ShareWithAvatar = shareWithAvatar
	this.ShareWithDisplayname = shareWithDisplayname
	this.ShareWithLink = shareWithLink
	this.Status = status
	this.Stime = stime
	this.Storage = storage
	this.StorageId = storageId
	this.Token = token
	this.UidFileOwner = uidFileOwner
	this.UidOwner = uidOwner
	this.Url = url
	return &this
}

// NewFilesSharingShareWithDefaults instantiates a new FilesSharingShare object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilesSharingShareWithDefaults() *FilesSharingShare {
	this := FilesSharingShare{}
	return &this
}

// GetAttributes returns the Attributes field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FilesSharingShare) GetAttributes() string {
	if o == nil || o.Attributes.Get() == nil {
		var ret string
		return ret
	}

	return *o.Attributes.Get()
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesSharingShare) GetAttributesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attributes.Get(), o.Attributes.IsSet()
}

// SetAttributes sets field value
func (o *FilesSharingShare) SetAttributes(v string) {
	o.Attributes.Set(&v)
}

// GetCanDelete returns the CanDelete field value
func (o *FilesSharingShare) GetCanDelete() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanDelete
}

// GetCanDeleteOk returns a tuple with the CanDelete field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShare) GetCanDeleteOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanDelete, true
}

// SetCanDelete sets field value
func (o *FilesSharingShare) SetCanDelete(v bool) {
	o.CanDelete = v
}

// GetCanEdit returns the CanEdit field value
func (o *FilesSharingShare) GetCanEdit() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanEdit
}

// GetCanEditOk returns a tuple with the CanEdit field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShare) GetCanEditOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanEdit, true
}

// SetCanEdit sets field value
func (o *FilesSharingShare) SetCanEdit(v bool) {
	o.CanEdit = v
}

// GetDisplaynameFileOwner returns the DisplaynameFileOwner field value
func (o *FilesSharingShare) GetDisplaynameFileOwner() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplaynameFileOwner
}

// GetDisplaynameFileOwnerOk returns a tuple with the DisplaynameFileOwner field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShare) GetDisplaynameFileOwnerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplaynameFileOwner, true
}

// SetDisplaynameFileOwner sets field value
func (o *FilesSharingShare) SetDisplaynameFileOwner(v string) {
	o.DisplaynameFileOwner = v
}

// GetDisplaynameOwner returns the DisplaynameOwner field value
func (o *FilesSharingShare) GetDisplaynameOwner() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplaynameOwner
}

// GetDisplaynameOwnerOk returns a tuple with the DisplaynameOwner field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShare) GetDisplaynameOwnerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplaynameOwner, true
}

// SetDisplaynameOwner sets field value
func (o *FilesSharingShare) SetDisplaynameOwner(v string) {
	o.DisplaynameOwner = v
}

// GetExpiration returns the Expiration field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FilesSharingShare) GetExpiration() string {
	if o == nil || o.Expiration.Get() == nil {
		var ret string
		return ret
	}

	return *o.Expiration.Get()
}

// GetExpirationOk returns a tuple with the Expiration field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesSharingShare) GetExpirationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expiration.Get(), o.Expiration.IsSet()
}

// SetExpiration sets field value
func (o *FilesSharingShare) SetExpiration(v string) {
	o.Expiration.Set(&v)
}

// GetFileParent returns the FileParent field value
func (o *FilesSharingShare) GetFileParent() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.FileParent
}

// GetFileParentOk returns a tuple with the FileParent field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShare) GetFileParentOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileParent, true
}

// SetFileParent sets field value
func (o *FilesSharingShare) SetFileParent(v int64) {
	o.FileParent = v
}

// GetFileSource returns the FileSource field value
func (o *FilesSharingShare) GetFileSource() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.FileSource
}

// GetFileSourceOk returns a tuple with the FileSource field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShare) GetFileSourceOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileSource, true
}

// SetFileSource sets field value
func (o *FilesSharingShare) SetFileSource(v int64) {
	o.FileSource = v
}

// GetFileTarget returns the FileTarget field value
func (o *FilesSharingShare) GetFileTarget() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileTarget
}

// GetFileTargetOk returns a tuple with the FileTarget field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShare) GetFileTargetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileTarget, true
}

// SetFileTarget sets field value
func (o *FilesSharingShare) SetFileTarget(v string) {
	o.FileTarget = v
}

// GetHasPreview returns the HasPreview field value
func (o *FilesSharingShare) GetHasPreview() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasPreview
}

// GetHasPreviewOk returns a tuple with the HasPreview field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShare) GetHasPreviewOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasPreview, true
}

// SetHasPreview sets field value
func (o *FilesSharingShare) SetHasPreview(v bool) {
	o.HasPreview = v
}

// GetId returns the Id field value
func (o *FilesSharingShare) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShare) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FilesSharingShare) SetId(v string) {
	o.Id = v
}

// GetItemSource returns the ItemSource field value
func (o *FilesSharingShare) GetItemSource() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ItemSource
}

// GetItemSourceOk returns a tuple with the ItemSource field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShare) GetItemSourceOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ItemSource, true
}

// SetItemSource sets field value
func (o *FilesSharingShare) SetItemSource(v int64) {
	o.ItemSource = v
}

// GetItemType returns the ItemType field value
func (o *FilesSharingShare) GetItemType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ItemType
}

// GetItemTypeOk returns a tuple with the ItemType field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShare) GetItemTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ItemType, true
}

// SetItemType sets field value
func (o *FilesSharingShare) SetItemType(v string) {
	o.ItemType = v
}

// GetLabel returns the Label field value
func (o *FilesSharingShare) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShare) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *FilesSharingShare) SetLabel(v string) {
	o.Label = v
}

// GetMailSend returns the MailSend field value
func (o *FilesSharingShare) GetMailSend() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MailSend
}

// GetMailSendOk returns a tuple with the MailSend field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShare) GetMailSendOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MailSend, true
}

// SetMailSend sets field value
func (o *FilesSharingShare) SetMailSend(v int64) {
	o.MailSend = v
}

// GetMimetype returns the Mimetype field value
func (o *FilesSharingShare) GetMimetype() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mimetype
}

// GetMimetypeOk returns a tuple with the Mimetype field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShare) GetMimetypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mimetype, true
}

// SetMimetype sets field value
func (o *FilesSharingShare) SetMimetype(v string) {
	o.Mimetype = v
}

// GetNote returns the Note field value
func (o *FilesSharingShare) GetNote() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Note
}

// GetNoteOk returns a tuple with the Note field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShare) GetNoteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Note, true
}

// SetNote sets field value
func (o *FilesSharingShare) SetNote(v string) {
	o.Note = v
}

// GetPassword returns the Password field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FilesSharingShare) GetPassword() string {
	if o == nil || o.Password.Get() == nil {
		var ret string
		return ret
	}

	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesSharingShare) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// SetPassword sets field value
func (o *FilesSharingShare) SetPassword(v string) {
	o.Password.Set(&v)
}

// GetPasswordExpirationTime returns the PasswordExpirationTime field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FilesSharingShare) GetPasswordExpirationTime() string {
	if o == nil || o.PasswordExpirationTime.Get() == nil {
		var ret string
		return ret
	}

	return *o.PasswordExpirationTime.Get()
}

// GetPasswordExpirationTimeOk returns a tuple with the PasswordExpirationTime field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesSharingShare) GetPasswordExpirationTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PasswordExpirationTime.Get(), o.PasswordExpirationTime.IsSet()
}

// SetPasswordExpirationTime sets field value
func (o *FilesSharingShare) SetPasswordExpirationTime(v string) {
	o.PasswordExpirationTime.Set(&v)
}

// GetPath returns the Path field value
func (o *FilesSharingShare) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShare) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *FilesSharingShare) SetPath(v string) {
	o.Path = v
}

// GetPermissions returns the Permissions field value
func (o *FilesSharingShare) GetPermissions() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShare) GetPermissionsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Permissions, true
}

// SetPermissions sets field value
func (o *FilesSharingShare) SetPermissions(v int64) {
	o.Permissions = v
}

// GetSendPasswordByTalk returns the SendPasswordByTalk field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *FilesSharingShare) GetSendPasswordByTalk() bool {
	if o == nil || o.SendPasswordByTalk.Get() == nil {
		var ret bool
		return ret
	}

	return *o.SendPasswordByTalk.Get()
}

// GetSendPasswordByTalkOk returns a tuple with the SendPasswordByTalk field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesSharingShare) GetSendPasswordByTalkOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SendPasswordByTalk.Get(), o.SendPasswordByTalk.IsSet()
}

// SetSendPasswordByTalk sets field value
func (o *FilesSharingShare) SetSendPasswordByTalk(v bool) {
	o.SendPasswordByTalk.Set(&v)
}

// GetShareType returns the ShareType field value
func (o *FilesSharingShare) GetShareType() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ShareType
}

// GetShareTypeOk returns a tuple with the ShareType field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShare) GetShareTypeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShareType, true
}

// SetShareType sets field value
func (o *FilesSharingShare) SetShareType(v int64) {
	o.ShareType = v
}

// GetShareWith returns the ShareWith field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FilesSharingShare) GetShareWith() string {
	if o == nil || o.ShareWith.Get() == nil {
		var ret string
		return ret
	}

	return *o.ShareWith.Get()
}

// GetShareWithOk returns a tuple with the ShareWith field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesSharingShare) GetShareWithOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShareWith.Get(), o.ShareWith.IsSet()
}

// SetShareWith sets field value
func (o *FilesSharingShare) SetShareWith(v string) {
	o.ShareWith.Set(&v)
}

// GetShareWithAvatar returns the ShareWithAvatar field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FilesSharingShare) GetShareWithAvatar() string {
	if o == nil || o.ShareWithAvatar.Get() == nil {
		var ret string
		return ret
	}

	return *o.ShareWithAvatar.Get()
}

// GetShareWithAvatarOk returns a tuple with the ShareWithAvatar field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesSharingShare) GetShareWithAvatarOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShareWithAvatar.Get(), o.ShareWithAvatar.IsSet()
}

// SetShareWithAvatar sets field value
func (o *FilesSharingShare) SetShareWithAvatar(v string) {
	o.ShareWithAvatar.Set(&v)
}

// GetShareWithDisplayname returns the ShareWithDisplayname field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FilesSharingShare) GetShareWithDisplayname() string {
	if o == nil || o.ShareWithDisplayname.Get() == nil {
		var ret string
		return ret
	}

	return *o.ShareWithDisplayname.Get()
}

// GetShareWithDisplaynameOk returns a tuple with the ShareWithDisplayname field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesSharingShare) GetShareWithDisplaynameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShareWithDisplayname.Get(), o.ShareWithDisplayname.IsSet()
}

// SetShareWithDisplayname sets field value
func (o *FilesSharingShare) SetShareWithDisplayname(v string) {
	o.ShareWithDisplayname.Set(&v)
}

// GetShareWithLink returns the ShareWithLink field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FilesSharingShare) GetShareWithLink() string {
	if o == nil || o.ShareWithLink.Get() == nil {
		var ret string
		return ret
	}

	return *o.ShareWithLink.Get()
}

// GetShareWithLinkOk returns a tuple with the ShareWithLink field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesSharingShare) GetShareWithLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShareWithLink.Get(), o.ShareWithLink.IsSet()
}

// SetShareWithLink sets field value
func (o *FilesSharingShare) SetShareWithLink(v string) {
	o.ShareWithLink.Set(&v)
}

// GetStatus returns the Status field value
// If the value is explicit nil, the zero value for FilesSharingShareStatus will be returned
func (o *FilesSharingShare) GetStatus() FilesSharingShareStatus {
	if o == nil || o.Status.Get() == nil {
		var ret FilesSharingShareStatus
		return ret
	}

	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesSharingShare) GetStatusOk() (*FilesSharingShareStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// SetStatus sets field value
func (o *FilesSharingShare) SetStatus(v FilesSharingShareStatus) {
	o.Status.Set(&v)
}

// GetStime returns the Stime field value
func (o *FilesSharingShare) GetStime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Stime
}

// GetStimeOk returns a tuple with the Stime field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShare) GetStimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stime, true
}

// SetStime sets field value
func (o *FilesSharingShare) SetStime(v int64) {
	o.Stime = v
}

// GetStorage returns the Storage field value
func (o *FilesSharingShare) GetStorage() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShare) GetStorageOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Storage, true
}

// SetStorage sets field value
func (o *FilesSharingShare) SetStorage(v int64) {
	o.Storage = v
}

// GetStorageId returns the StorageId field value
func (o *FilesSharingShare) GetStorageId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StorageId
}

// GetStorageIdOk returns a tuple with the StorageId field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShare) GetStorageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageId, true
}

// SetStorageId sets field value
func (o *FilesSharingShare) SetStorageId(v string) {
	o.StorageId = v
}

// GetToken returns the Token field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FilesSharingShare) GetToken() string {
	if o == nil || o.Token.Get() == nil {
		var ret string
		return ret
	}

	return *o.Token.Get()
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesSharingShare) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Token.Get(), o.Token.IsSet()
}

// SetToken sets field value
func (o *FilesSharingShare) SetToken(v string) {
	o.Token.Set(&v)
}

// GetUidFileOwner returns the UidFileOwner field value
func (o *FilesSharingShare) GetUidFileOwner() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UidFileOwner
}

// GetUidFileOwnerOk returns a tuple with the UidFileOwner field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShare) GetUidFileOwnerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UidFileOwner, true
}

// SetUidFileOwner sets field value
func (o *FilesSharingShare) SetUidFileOwner(v string) {
	o.UidFileOwner = v
}

// GetUidOwner returns the UidOwner field value
func (o *FilesSharingShare) GetUidOwner() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UidOwner
}

// GetUidOwnerOk returns a tuple with the UidOwner field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShare) GetUidOwnerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UidOwner, true
}

// SetUidOwner sets field value
func (o *FilesSharingShare) SetUidOwner(v string) {
	o.UidOwner = v
}

// GetUrl returns the Url field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FilesSharingShare) GetUrl() string {
	if o == nil || o.Url.Get() == nil {
		var ret string
		return ret
	}

	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesSharingShare) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// SetUrl sets field value
func (o *FilesSharingShare) SetUrl(v string) {
	o.Url.Set(&v)
}

func (o FilesSharingShare) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FilesSharingShare) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["attributes"] = o.Attributes.Get()
	toSerialize["can_delete"] = o.CanDelete
	toSerialize["can_edit"] = o.CanEdit
	toSerialize["displayname_file_owner"] = o.DisplaynameFileOwner
	toSerialize["displayname_owner"] = o.DisplaynameOwner
	toSerialize["expiration"] = o.Expiration.Get()
	toSerialize["file_parent"] = o.FileParent
	toSerialize["file_source"] = o.FileSource
	toSerialize["file_target"] = o.FileTarget
	toSerialize["has_preview"] = o.HasPreview
	toSerialize["id"] = o.Id
	toSerialize["item_source"] = o.ItemSource
	toSerialize["item_type"] = o.ItemType
	toSerialize["label"] = o.Label
	toSerialize["mail_send"] = o.MailSend
	toSerialize["mimetype"] = o.Mimetype
	toSerialize["note"] = o.Note
	toSerialize["password"] = o.Password.Get()
	toSerialize["password_expiration_time"] = o.PasswordExpirationTime.Get()
	toSerialize["path"] = o.Path
	toSerialize["permissions"] = o.Permissions
	toSerialize["send_password_by_talk"] = o.SendPasswordByTalk.Get()
	toSerialize["share_type"] = o.ShareType
	toSerialize["share_with"] = o.ShareWith.Get()
	toSerialize["share_with_avatar"] = o.ShareWithAvatar.Get()
	toSerialize["share_with_displayname"] = o.ShareWithDisplayname.Get()
	toSerialize["share_with_link"] = o.ShareWithLink.Get()
	toSerialize["status"] = o.Status.Get()
	toSerialize["stime"] = o.Stime
	toSerialize["storage"] = o.Storage
	toSerialize["storage_id"] = o.StorageId
	toSerialize["token"] = o.Token.Get()
	toSerialize["uid_file_owner"] = o.UidFileOwner
	toSerialize["uid_owner"] = o.UidOwner
	toSerialize["url"] = o.Url.Get()
	return toSerialize, nil
}

type NullableFilesSharingShare struct {
	value *FilesSharingShare
	isSet bool
}

func (v NullableFilesSharingShare) Get() *FilesSharingShare {
	return v.value
}

func (v *NullableFilesSharingShare) Set(val *FilesSharingShare) {
	v.value = val
	v.isSet = true
}

func (v NullableFilesSharingShare) IsSet() bool {
	return v.isSet
}

func (v *NullableFilesSharingShare) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilesSharingShare(val *FilesSharingShare) *NullableFilesSharingShare {
	return &NullableFilesSharingShare{value: val, isSet: true}
}

func (v NullableFilesSharingShare) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilesSharingShare) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

