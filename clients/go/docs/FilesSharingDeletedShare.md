# FilesSharingDeletedShare

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ShareType** | **int64** |  | 
**UidOwner** | **string** |  | 
**DisplaynameOwner** | **string** |  | 
**Permissions** | **int64** |  | 
**Stime** | **int64** |  | 
**UidFileOwner** | **string** |  | 
**DisplaynameFileOwner** | **string** |  | 
**Path** | **string** |  | 
**ItemType** | **string** |  | 
**Mimetype** | **string** |  | 
**Storage** | **int64** |  | 
**ItemSource** | **int64** |  | 
**FileSource** | **int64** |  | 
**FileParent** | **int64** |  | 
**FileTarget** | **int64** |  | 
**Expiration** | **NullableString** |  | 
**ShareWith** | **NullableString** |  | 
**ShareWithDisplayname** | **NullableString** |  | 
**ShareWithLink** | **NullableString** |  | 

## Methods

### NewFilesSharingDeletedShare

`func NewFilesSharingDeletedShare(id string, shareType int64, uidOwner string, displaynameOwner string, permissions int64, stime int64, uidFileOwner string, displaynameFileOwner string, path string, itemType string, mimetype string, storage int64, itemSource int64, fileSource int64, fileParent int64, fileTarget int64, expiration NullableString, shareWith NullableString, shareWithDisplayname NullableString, shareWithLink NullableString, ) *FilesSharingDeletedShare`

NewFilesSharingDeletedShare instantiates a new FilesSharingDeletedShare object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilesSharingDeletedShareWithDefaults

`func NewFilesSharingDeletedShareWithDefaults() *FilesSharingDeletedShare`

NewFilesSharingDeletedShareWithDefaults instantiates a new FilesSharingDeletedShare object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FilesSharingDeletedShare) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FilesSharingDeletedShare) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FilesSharingDeletedShare) SetId(v string)`

SetId sets Id field to given value.


### GetShareType

`func (o *FilesSharingDeletedShare) GetShareType() int64`

GetShareType returns the ShareType field if non-nil, zero value otherwise.

### GetShareTypeOk

`func (o *FilesSharingDeletedShare) GetShareTypeOk() (*int64, bool)`

GetShareTypeOk returns a tuple with the ShareType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareType

`func (o *FilesSharingDeletedShare) SetShareType(v int64)`

SetShareType sets ShareType field to given value.


### GetUidOwner

`func (o *FilesSharingDeletedShare) GetUidOwner() string`

GetUidOwner returns the UidOwner field if non-nil, zero value otherwise.

### GetUidOwnerOk

`func (o *FilesSharingDeletedShare) GetUidOwnerOk() (*string, bool)`

GetUidOwnerOk returns a tuple with the UidOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidOwner

`func (o *FilesSharingDeletedShare) SetUidOwner(v string)`

SetUidOwner sets UidOwner field to given value.


### GetDisplaynameOwner

`func (o *FilesSharingDeletedShare) GetDisplaynameOwner() string`

GetDisplaynameOwner returns the DisplaynameOwner field if non-nil, zero value otherwise.

### GetDisplaynameOwnerOk

`func (o *FilesSharingDeletedShare) GetDisplaynameOwnerOk() (*string, bool)`

GetDisplaynameOwnerOk returns a tuple with the DisplaynameOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplaynameOwner

`func (o *FilesSharingDeletedShare) SetDisplaynameOwner(v string)`

SetDisplaynameOwner sets DisplaynameOwner field to given value.


### GetPermissions

`func (o *FilesSharingDeletedShare) GetPermissions() int64`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *FilesSharingDeletedShare) GetPermissionsOk() (*int64, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *FilesSharingDeletedShare) SetPermissions(v int64)`

SetPermissions sets Permissions field to given value.


### GetStime

`func (o *FilesSharingDeletedShare) GetStime() int64`

GetStime returns the Stime field if non-nil, zero value otherwise.

### GetStimeOk

`func (o *FilesSharingDeletedShare) GetStimeOk() (*int64, bool)`

GetStimeOk returns a tuple with the Stime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStime

`func (o *FilesSharingDeletedShare) SetStime(v int64)`

SetStime sets Stime field to given value.


### GetUidFileOwner

`func (o *FilesSharingDeletedShare) GetUidFileOwner() string`

GetUidFileOwner returns the UidFileOwner field if non-nil, zero value otherwise.

### GetUidFileOwnerOk

`func (o *FilesSharingDeletedShare) GetUidFileOwnerOk() (*string, bool)`

GetUidFileOwnerOk returns a tuple with the UidFileOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidFileOwner

`func (o *FilesSharingDeletedShare) SetUidFileOwner(v string)`

SetUidFileOwner sets UidFileOwner field to given value.


### GetDisplaynameFileOwner

`func (o *FilesSharingDeletedShare) GetDisplaynameFileOwner() string`

GetDisplaynameFileOwner returns the DisplaynameFileOwner field if non-nil, zero value otherwise.

### GetDisplaynameFileOwnerOk

`func (o *FilesSharingDeletedShare) GetDisplaynameFileOwnerOk() (*string, bool)`

GetDisplaynameFileOwnerOk returns a tuple with the DisplaynameFileOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplaynameFileOwner

`func (o *FilesSharingDeletedShare) SetDisplaynameFileOwner(v string)`

SetDisplaynameFileOwner sets DisplaynameFileOwner field to given value.


### GetPath

`func (o *FilesSharingDeletedShare) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *FilesSharingDeletedShare) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *FilesSharingDeletedShare) SetPath(v string)`

SetPath sets Path field to given value.


### GetItemType

`func (o *FilesSharingDeletedShare) GetItemType() string`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *FilesSharingDeletedShare) GetItemTypeOk() (*string, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *FilesSharingDeletedShare) SetItemType(v string)`

SetItemType sets ItemType field to given value.


### GetMimetype

`func (o *FilesSharingDeletedShare) GetMimetype() string`

GetMimetype returns the Mimetype field if non-nil, zero value otherwise.

### GetMimetypeOk

`func (o *FilesSharingDeletedShare) GetMimetypeOk() (*string, bool)`

GetMimetypeOk returns a tuple with the Mimetype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimetype

`func (o *FilesSharingDeletedShare) SetMimetype(v string)`

SetMimetype sets Mimetype field to given value.


### GetStorage

`func (o *FilesSharingDeletedShare) GetStorage() int64`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *FilesSharingDeletedShare) GetStorageOk() (*int64, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *FilesSharingDeletedShare) SetStorage(v int64)`

SetStorage sets Storage field to given value.


### GetItemSource

`func (o *FilesSharingDeletedShare) GetItemSource() int64`

GetItemSource returns the ItemSource field if non-nil, zero value otherwise.

### GetItemSourceOk

`func (o *FilesSharingDeletedShare) GetItemSourceOk() (*int64, bool)`

GetItemSourceOk returns a tuple with the ItemSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemSource

`func (o *FilesSharingDeletedShare) SetItemSource(v int64)`

SetItemSource sets ItemSource field to given value.


### GetFileSource

`func (o *FilesSharingDeletedShare) GetFileSource() int64`

GetFileSource returns the FileSource field if non-nil, zero value otherwise.

### GetFileSourceOk

`func (o *FilesSharingDeletedShare) GetFileSourceOk() (*int64, bool)`

GetFileSourceOk returns a tuple with the FileSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSource

`func (o *FilesSharingDeletedShare) SetFileSource(v int64)`

SetFileSource sets FileSource field to given value.


### GetFileParent

`func (o *FilesSharingDeletedShare) GetFileParent() int64`

GetFileParent returns the FileParent field if non-nil, zero value otherwise.

### GetFileParentOk

`func (o *FilesSharingDeletedShare) GetFileParentOk() (*int64, bool)`

GetFileParentOk returns a tuple with the FileParent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileParent

`func (o *FilesSharingDeletedShare) SetFileParent(v int64)`

SetFileParent sets FileParent field to given value.


### GetFileTarget

`func (o *FilesSharingDeletedShare) GetFileTarget() int64`

GetFileTarget returns the FileTarget field if non-nil, zero value otherwise.

### GetFileTargetOk

`func (o *FilesSharingDeletedShare) GetFileTargetOk() (*int64, bool)`

GetFileTargetOk returns a tuple with the FileTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileTarget

`func (o *FilesSharingDeletedShare) SetFileTarget(v int64)`

SetFileTarget sets FileTarget field to given value.


### GetExpiration

`func (o *FilesSharingDeletedShare) GetExpiration() string`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *FilesSharingDeletedShare) GetExpirationOk() (*string, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *FilesSharingDeletedShare) SetExpiration(v string)`

SetExpiration sets Expiration field to given value.


### SetExpirationNil

`func (o *FilesSharingDeletedShare) SetExpirationNil(b bool)`

 SetExpirationNil sets the value for Expiration to be an explicit nil

### UnsetExpiration
`func (o *FilesSharingDeletedShare) UnsetExpiration()`

UnsetExpiration ensures that no value is present for Expiration, not even an explicit nil
### GetShareWith

`func (o *FilesSharingDeletedShare) GetShareWith() string`

GetShareWith returns the ShareWith field if non-nil, zero value otherwise.

### GetShareWithOk

`func (o *FilesSharingDeletedShare) GetShareWithOk() (*string, bool)`

GetShareWithOk returns a tuple with the ShareWith field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareWith

`func (o *FilesSharingDeletedShare) SetShareWith(v string)`

SetShareWith sets ShareWith field to given value.


### SetShareWithNil

`func (o *FilesSharingDeletedShare) SetShareWithNil(b bool)`

 SetShareWithNil sets the value for ShareWith to be an explicit nil

### UnsetShareWith
`func (o *FilesSharingDeletedShare) UnsetShareWith()`

UnsetShareWith ensures that no value is present for ShareWith, not even an explicit nil
### GetShareWithDisplayname

`func (o *FilesSharingDeletedShare) GetShareWithDisplayname() string`

GetShareWithDisplayname returns the ShareWithDisplayname field if non-nil, zero value otherwise.

### GetShareWithDisplaynameOk

`func (o *FilesSharingDeletedShare) GetShareWithDisplaynameOk() (*string, bool)`

GetShareWithDisplaynameOk returns a tuple with the ShareWithDisplayname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareWithDisplayname

`func (o *FilesSharingDeletedShare) SetShareWithDisplayname(v string)`

SetShareWithDisplayname sets ShareWithDisplayname field to given value.


### SetShareWithDisplaynameNil

`func (o *FilesSharingDeletedShare) SetShareWithDisplaynameNil(b bool)`

 SetShareWithDisplaynameNil sets the value for ShareWithDisplayname to be an explicit nil

### UnsetShareWithDisplayname
`func (o *FilesSharingDeletedShare) UnsetShareWithDisplayname()`

UnsetShareWithDisplayname ensures that no value is present for ShareWithDisplayname, not even an explicit nil
### GetShareWithLink

`func (o *FilesSharingDeletedShare) GetShareWithLink() string`

GetShareWithLink returns the ShareWithLink field if non-nil, zero value otherwise.

### GetShareWithLinkOk

`func (o *FilesSharingDeletedShare) GetShareWithLinkOk() (*string, bool)`

GetShareWithLinkOk returns a tuple with the ShareWithLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareWithLink

`func (o *FilesSharingDeletedShare) SetShareWithLink(v string)`

SetShareWithLink sets ShareWithLink field to given value.


### SetShareWithLinkNil

`func (o *FilesSharingDeletedShare) SetShareWithLinkNil(b bool)`

 SetShareWithLinkNil sets the value for ShareWithLink to be an explicit nil

### UnsetShareWithLink
`func (o *FilesSharingDeletedShare) UnsetShareWithLink()`

UnsetShareWithLink ensures that no value is present for ShareWithLink, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


