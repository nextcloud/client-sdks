# FilesSharingShare

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | **NullableString** |  | 
**CanDelete** | **bool** |  | 
**CanEdit** | **bool** |  | 
**DisplaynameFileOwner** | **string** |  | 
**DisplaynameOwner** | **string** |  | 
**Expiration** | **NullableString** |  | 
**FileParent** | **int64** |  | 
**FileSource** | **int64** |  | 
**FileTarget** | **string** |  | 
**HasPreview** | **bool** |  | 
**Id** | **string** |  | 
**ItemSource** | **int64** |  | 
**ItemType** | **string** |  | 
**Label** | **string** |  | 
**MailSend** | **int64** |  | 
**Mimetype** | **string** |  | 
**Note** | **string** |  | 
**Password** | **NullableString** |  | 
**PasswordExpirationTime** | **NullableString** |  | 
**Path** | **string** |  | 
**Permissions** | **int64** |  | 
**SendPasswordByTalk** | **NullableBool** |  | 
**ShareType** | **int64** |  | 
**ShareWith** | **NullableString** |  | 
**ShareWithAvatar** | **NullableString** |  | 
**ShareWithDisplayname** | **NullableString** |  | 
**ShareWithLink** | **NullableString** |  | 
**Status** | [**NullableFilesSharingShareStatus**](FilesSharingShareStatus.md) |  | 
**Stime** | **int64** |  | 
**Storage** | **int64** |  | 
**StorageId** | **string** |  | 
**Token** | **NullableString** |  | 
**UidFileOwner** | **string** |  | 
**UidOwner** | **string** |  | 
**Url** | **NullableString** |  | 

## Methods

### NewFilesSharingShare

`func NewFilesSharingShare(attributes NullableString, canDelete bool, canEdit bool, displaynameFileOwner string, displaynameOwner string, expiration NullableString, fileParent int64, fileSource int64, fileTarget string, hasPreview bool, id string, itemSource int64, itemType string, label string, mailSend int64, mimetype string, note string, password NullableString, passwordExpirationTime NullableString, path string, permissions int64, sendPasswordByTalk NullableBool, shareType int64, shareWith NullableString, shareWithAvatar NullableString, shareWithDisplayname NullableString, shareWithLink NullableString, status NullableFilesSharingShareStatus, stime int64, storage int64, storageId string, token NullableString, uidFileOwner string, uidOwner string, url NullableString, ) *FilesSharingShare`

NewFilesSharingShare instantiates a new FilesSharingShare object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilesSharingShareWithDefaults

`func NewFilesSharingShareWithDefaults() *FilesSharingShare`

NewFilesSharingShareWithDefaults instantiates a new FilesSharingShare object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *FilesSharingShare) GetAttributes() string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *FilesSharingShare) GetAttributesOk() (*string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *FilesSharingShare) SetAttributes(v string)`

SetAttributes sets Attributes field to given value.


### SetAttributesNil

`func (o *FilesSharingShare) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *FilesSharingShare) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetCanDelete

`func (o *FilesSharingShare) GetCanDelete() bool`

GetCanDelete returns the CanDelete field if non-nil, zero value otherwise.

### GetCanDeleteOk

`func (o *FilesSharingShare) GetCanDeleteOk() (*bool, bool)`

GetCanDeleteOk returns a tuple with the CanDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDelete

`func (o *FilesSharingShare) SetCanDelete(v bool)`

SetCanDelete sets CanDelete field to given value.


### GetCanEdit

`func (o *FilesSharingShare) GetCanEdit() bool`

GetCanEdit returns the CanEdit field if non-nil, zero value otherwise.

### GetCanEditOk

`func (o *FilesSharingShare) GetCanEditOk() (*bool, bool)`

GetCanEditOk returns a tuple with the CanEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEdit

`func (o *FilesSharingShare) SetCanEdit(v bool)`

SetCanEdit sets CanEdit field to given value.


### GetDisplaynameFileOwner

`func (o *FilesSharingShare) GetDisplaynameFileOwner() string`

GetDisplaynameFileOwner returns the DisplaynameFileOwner field if non-nil, zero value otherwise.

### GetDisplaynameFileOwnerOk

`func (o *FilesSharingShare) GetDisplaynameFileOwnerOk() (*string, bool)`

GetDisplaynameFileOwnerOk returns a tuple with the DisplaynameFileOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplaynameFileOwner

`func (o *FilesSharingShare) SetDisplaynameFileOwner(v string)`

SetDisplaynameFileOwner sets DisplaynameFileOwner field to given value.


### GetDisplaynameOwner

`func (o *FilesSharingShare) GetDisplaynameOwner() string`

GetDisplaynameOwner returns the DisplaynameOwner field if non-nil, zero value otherwise.

### GetDisplaynameOwnerOk

`func (o *FilesSharingShare) GetDisplaynameOwnerOk() (*string, bool)`

GetDisplaynameOwnerOk returns a tuple with the DisplaynameOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplaynameOwner

`func (o *FilesSharingShare) SetDisplaynameOwner(v string)`

SetDisplaynameOwner sets DisplaynameOwner field to given value.


### GetExpiration

`func (o *FilesSharingShare) GetExpiration() string`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *FilesSharingShare) GetExpirationOk() (*string, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *FilesSharingShare) SetExpiration(v string)`

SetExpiration sets Expiration field to given value.


### SetExpirationNil

`func (o *FilesSharingShare) SetExpirationNil(b bool)`

 SetExpirationNil sets the value for Expiration to be an explicit nil

### UnsetExpiration
`func (o *FilesSharingShare) UnsetExpiration()`

UnsetExpiration ensures that no value is present for Expiration, not even an explicit nil
### GetFileParent

`func (o *FilesSharingShare) GetFileParent() int64`

GetFileParent returns the FileParent field if non-nil, zero value otherwise.

### GetFileParentOk

`func (o *FilesSharingShare) GetFileParentOk() (*int64, bool)`

GetFileParentOk returns a tuple with the FileParent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileParent

`func (o *FilesSharingShare) SetFileParent(v int64)`

SetFileParent sets FileParent field to given value.


### GetFileSource

`func (o *FilesSharingShare) GetFileSource() int64`

GetFileSource returns the FileSource field if non-nil, zero value otherwise.

### GetFileSourceOk

`func (o *FilesSharingShare) GetFileSourceOk() (*int64, bool)`

GetFileSourceOk returns a tuple with the FileSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSource

`func (o *FilesSharingShare) SetFileSource(v int64)`

SetFileSource sets FileSource field to given value.


### GetFileTarget

`func (o *FilesSharingShare) GetFileTarget() string`

GetFileTarget returns the FileTarget field if non-nil, zero value otherwise.

### GetFileTargetOk

`func (o *FilesSharingShare) GetFileTargetOk() (*string, bool)`

GetFileTargetOk returns a tuple with the FileTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileTarget

`func (o *FilesSharingShare) SetFileTarget(v string)`

SetFileTarget sets FileTarget field to given value.


### GetHasPreview

`func (o *FilesSharingShare) GetHasPreview() bool`

GetHasPreview returns the HasPreview field if non-nil, zero value otherwise.

### GetHasPreviewOk

`func (o *FilesSharingShare) GetHasPreviewOk() (*bool, bool)`

GetHasPreviewOk returns a tuple with the HasPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPreview

`func (o *FilesSharingShare) SetHasPreview(v bool)`

SetHasPreview sets HasPreview field to given value.


### GetId

`func (o *FilesSharingShare) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FilesSharingShare) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FilesSharingShare) SetId(v string)`

SetId sets Id field to given value.


### GetItemSource

`func (o *FilesSharingShare) GetItemSource() int64`

GetItemSource returns the ItemSource field if non-nil, zero value otherwise.

### GetItemSourceOk

`func (o *FilesSharingShare) GetItemSourceOk() (*int64, bool)`

GetItemSourceOk returns a tuple with the ItemSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemSource

`func (o *FilesSharingShare) SetItemSource(v int64)`

SetItemSource sets ItemSource field to given value.


### GetItemType

`func (o *FilesSharingShare) GetItemType() string`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *FilesSharingShare) GetItemTypeOk() (*string, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *FilesSharingShare) SetItemType(v string)`

SetItemType sets ItemType field to given value.


### GetLabel

`func (o *FilesSharingShare) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FilesSharingShare) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FilesSharingShare) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetMailSend

`func (o *FilesSharingShare) GetMailSend() int64`

GetMailSend returns the MailSend field if non-nil, zero value otherwise.

### GetMailSendOk

`func (o *FilesSharingShare) GetMailSendOk() (*int64, bool)`

GetMailSendOk returns a tuple with the MailSend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailSend

`func (o *FilesSharingShare) SetMailSend(v int64)`

SetMailSend sets MailSend field to given value.


### GetMimetype

`func (o *FilesSharingShare) GetMimetype() string`

GetMimetype returns the Mimetype field if non-nil, zero value otherwise.

### GetMimetypeOk

`func (o *FilesSharingShare) GetMimetypeOk() (*string, bool)`

GetMimetypeOk returns a tuple with the Mimetype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimetype

`func (o *FilesSharingShare) SetMimetype(v string)`

SetMimetype sets Mimetype field to given value.


### GetNote

`func (o *FilesSharingShare) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *FilesSharingShare) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *FilesSharingShare) SetNote(v string)`

SetNote sets Note field to given value.


### GetPassword

`func (o *FilesSharingShare) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *FilesSharingShare) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *FilesSharingShare) SetPassword(v string)`

SetPassword sets Password field to given value.


### SetPasswordNil

`func (o *FilesSharingShare) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *FilesSharingShare) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetPasswordExpirationTime

`func (o *FilesSharingShare) GetPasswordExpirationTime() string`

GetPasswordExpirationTime returns the PasswordExpirationTime field if non-nil, zero value otherwise.

### GetPasswordExpirationTimeOk

`func (o *FilesSharingShare) GetPasswordExpirationTimeOk() (*string, bool)`

GetPasswordExpirationTimeOk returns a tuple with the PasswordExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordExpirationTime

`func (o *FilesSharingShare) SetPasswordExpirationTime(v string)`

SetPasswordExpirationTime sets PasswordExpirationTime field to given value.


### SetPasswordExpirationTimeNil

`func (o *FilesSharingShare) SetPasswordExpirationTimeNil(b bool)`

 SetPasswordExpirationTimeNil sets the value for PasswordExpirationTime to be an explicit nil

### UnsetPasswordExpirationTime
`func (o *FilesSharingShare) UnsetPasswordExpirationTime()`

UnsetPasswordExpirationTime ensures that no value is present for PasswordExpirationTime, not even an explicit nil
### GetPath

`func (o *FilesSharingShare) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *FilesSharingShare) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *FilesSharingShare) SetPath(v string)`

SetPath sets Path field to given value.


### GetPermissions

`func (o *FilesSharingShare) GetPermissions() int64`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *FilesSharingShare) GetPermissionsOk() (*int64, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *FilesSharingShare) SetPermissions(v int64)`

SetPermissions sets Permissions field to given value.


### GetSendPasswordByTalk

`func (o *FilesSharingShare) GetSendPasswordByTalk() bool`

GetSendPasswordByTalk returns the SendPasswordByTalk field if non-nil, zero value otherwise.

### GetSendPasswordByTalkOk

`func (o *FilesSharingShare) GetSendPasswordByTalkOk() (*bool, bool)`

GetSendPasswordByTalkOk returns a tuple with the SendPasswordByTalk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendPasswordByTalk

`func (o *FilesSharingShare) SetSendPasswordByTalk(v bool)`

SetSendPasswordByTalk sets SendPasswordByTalk field to given value.


### SetSendPasswordByTalkNil

`func (o *FilesSharingShare) SetSendPasswordByTalkNil(b bool)`

 SetSendPasswordByTalkNil sets the value for SendPasswordByTalk to be an explicit nil

### UnsetSendPasswordByTalk
`func (o *FilesSharingShare) UnsetSendPasswordByTalk()`

UnsetSendPasswordByTalk ensures that no value is present for SendPasswordByTalk, not even an explicit nil
### GetShareType

`func (o *FilesSharingShare) GetShareType() int64`

GetShareType returns the ShareType field if non-nil, zero value otherwise.

### GetShareTypeOk

`func (o *FilesSharingShare) GetShareTypeOk() (*int64, bool)`

GetShareTypeOk returns a tuple with the ShareType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareType

`func (o *FilesSharingShare) SetShareType(v int64)`

SetShareType sets ShareType field to given value.


### GetShareWith

`func (o *FilesSharingShare) GetShareWith() string`

GetShareWith returns the ShareWith field if non-nil, zero value otherwise.

### GetShareWithOk

`func (o *FilesSharingShare) GetShareWithOk() (*string, bool)`

GetShareWithOk returns a tuple with the ShareWith field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareWith

`func (o *FilesSharingShare) SetShareWith(v string)`

SetShareWith sets ShareWith field to given value.


### SetShareWithNil

`func (o *FilesSharingShare) SetShareWithNil(b bool)`

 SetShareWithNil sets the value for ShareWith to be an explicit nil

### UnsetShareWith
`func (o *FilesSharingShare) UnsetShareWith()`

UnsetShareWith ensures that no value is present for ShareWith, not even an explicit nil
### GetShareWithAvatar

`func (o *FilesSharingShare) GetShareWithAvatar() string`

GetShareWithAvatar returns the ShareWithAvatar field if non-nil, zero value otherwise.

### GetShareWithAvatarOk

`func (o *FilesSharingShare) GetShareWithAvatarOk() (*string, bool)`

GetShareWithAvatarOk returns a tuple with the ShareWithAvatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareWithAvatar

`func (o *FilesSharingShare) SetShareWithAvatar(v string)`

SetShareWithAvatar sets ShareWithAvatar field to given value.


### SetShareWithAvatarNil

`func (o *FilesSharingShare) SetShareWithAvatarNil(b bool)`

 SetShareWithAvatarNil sets the value for ShareWithAvatar to be an explicit nil

### UnsetShareWithAvatar
`func (o *FilesSharingShare) UnsetShareWithAvatar()`

UnsetShareWithAvatar ensures that no value is present for ShareWithAvatar, not even an explicit nil
### GetShareWithDisplayname

`func (o *FilesSharingShare) GetShareWithDisplayname() string`

GetShareWithDisplayname returns the ShareWithDisplayname field if non-nil, zero value otherwise.

### GetShareWithDisplaynameOk

`func (o *FilesSharingShare) GetShareWithDisplaynameOk() (*string, bool)`

GetShareWithDisplaynameOk returns a tuple with the ShareWithDisplayname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareWithDisplayname

`func (o *FilesSharingShare) SetShareWithDisplayname(v string)`

SetShareWithDisplayname sets ShareWithDisplayname field to given value.


### SetShareWithDisplaynameNil

`func (o *FilesSharingShare) SetShareWithDisplaynameNil(b bool)`

 SetShareWithDisplaynameNil sets the value for ShareWithDisplayname to be an explicit nil

### UnsetShareWithDisplayname
`func (o *FilesSharingShare) UnsetShareWithDisplayname()`

UnsetShareWithDisplayname ensures that no value is present for ShareWithDisplayname, not even an explicit nil
### GetShareWithLink

`func (o *FilesSharingShare) GetShareWithLink() string`

GetShareWithLink returns the ShareWithLink field if non-nil, zero value otherwise.

### GetShareWithLinkOk

`func (o *FilesSharingShare) GetShareWithLinkOk() (*string, bool)`

GetShareWithLinkOk returns a tuple with the ShareWithLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareWithLink

`func (o *FilesSharingShare) SetShareWithLink(v string)`

SetShareWithLink sets ShareWithLink field to given value.


### SetShareWithLinkNil

`func (o *FilesSharingShare) SetShareWithLinkNil(b bool)`

 SetShareWithLinkNil sets the value for ShareWithLink to be an explicit nil

### UnsetShareWithLink
`func (o *FilesSharingShare) UnsetShareWithLink()`

UnsetShareWithLink ensures that no value is present for ShareWithLink, not even an explicit nil
### GetStatus

`func (o *FilesSharingShare) GetStatus() FilesSharingShareStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FilesSharingShare) GetStatusOk() (*FilesSharingShareStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FilesSharingShare) SetStatus(v FilesSharingShareStatus)`

SetStatus sets Status field to given value.


### SetStatusNil

`func (o *FilesSharingShare) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *FilesSharingShare) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetStime

`func (o *FilesSharingShare) GetStime() int64`

GetStime returns the Stime field if non-nil, zero value otherwise.

### GetStimeOk

`func (o *FilesSharingShare) GetStimeOk() (*int64, bool)`

GetStimeOk returns a tuple with the Stime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStime

`func (o *FilesSharingShare) SetStime(v int64)`

SetStime sets Stime field to given value.


### GetStorage

`func (o *FilesSharingShare) GetStorage() int64`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *FilesSharingShare) GetStorageOk() (*int64, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *FilesSharingShare) SetStorage(v int64)`

SetStorage sets Storage field to given value.


### GetStorageId

`func (o *FilesSharingShare) GetStorageId() string`

GetStorageId returns the StorageId field if non-nil, zero value otherwise.

### GetStorageIdOk

`func (o *FilesSharingShare) GetStorageIdOk() (*string, bool)`

GetStorageIdOk returns a tuple with the StorageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageId

`func (o *FilesSharingShare) SetStorageId(v string)`

SetStorageId sets StorageId field to given value.


### GetToken

`func (o *FilesSharingShare) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FilesSharingShare) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FilesSharingShare) SetToken(v string)`

SetToken sets Token field to given value.


### SetTokenNil

`func (o *FilesSharingShare) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *FilesSharingShare) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil
### GetUidFileOwner

`func (o *FilesSharingShare) GetUidFileOwner() string`

GetUidFileOwner returns the UidFileOwner field if non-nil, zero value otherwise.

### GetUidFileOwnerOk

`func (o *FilesSharingShare) GetUidFileOwnerOk() (*string, bool)`

GetUidFileOwnerOk returns a tuple with the UidFileOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidFileOwner

`func (o *FilesSharingShare) SetUidFileOwner(v string)`

SetUidFileOwner sets UidFileOwner field to given value.


### GetUidOwner

`func (o *FilesSharingShare) GetUidOwner() string`

GetUidOwner returns the UidOwner field if non-nil, zero value otherwise.

### GetUidOwnerOk

`func (o *FilesSharingShare) GetUidOwnerOk() (*string, bool)`

GetUidOwnerOk returns a tuple with the UidOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidOwner

`func (o *FilesSharingShare) SetUidOwner(v string)`

SetUidOwner sets UidOwner field to given value.


### GetUrl

`func (o *FilesSharingShare) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FilesSharingShare) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FilesSharingShare) SetUrl(v string)`

SetUrl sets Url field to given value.


### SetUrlNil

`func (o *FilesSharingShare) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *FilesSharingShare) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


