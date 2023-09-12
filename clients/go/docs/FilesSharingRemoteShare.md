# FilesSharingRemoteShare

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accepted** | **bool** |  | 
**FileId** | **NullableInt64** |  | 
**Id** | **int64** |  | 
**Mimetype** | **NullableString** |  | 
**Mountpoint** | **string** |  | 
**Mtime** | **NullableInt64** |  | 
**Name** | **string** |  | 
**Owner** | **string** |  | 
**Parent** | **NullableInt64** |  | 
**Permissions** | **NullableInt64** |  | 
**Remote** | **string** |  | 
**RemoteId** | **string** |  | 
**ShareToken** | **string** |  | 
**ShareType** | **int64** |  | 
**Type** | **NullableString** |  | 
**User** | **string** |  | 

## Methods

### NewFilesSharingRemoteShare

`func NewFilesSharingRemoteShare(accepted bool, fileId NullableInt64, id int64, mimetype NullableString, mountpoint string, mtime NullableInt64, name string, owner string, parent NullableInt64, permissions NullableInt64, remote string, remoteId string, shareToken string, shareType int64, type_ NullableString, user string, ) *FilesSharingRemoteShare`

NewFilesSharingRemoteShare instantiates a new FilesSharingRemoteShare object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilesSharingRemoteShareWithDefaults

`func NewFilesSharingRemoteShareWithDefaults() *FilesSharingRemoteShare`

NewFilesSharingRemoteShareWithDefaults instantiates a new FilesSharingRemoteShare object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccepted

`func (o *FilesSharingRemoteShare) GetAccepted() bool`

GetAccepted returns the Accepted field if non-nil, zero value otherwise.

### GetAcceptedOk

`func (o *FilesSharingRemoteShare) GetAcceptedOk() (*bool, bool)`

GetAcceptedOk returns a tuple with the Accepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccepted

`func (o *FilesSharingRemoteShare) SetAccepted(v bool)`

SetAccepted sets Accepted field to given value.


### GetFileId

`func (o *FilesSharingRemoteShare) GetFileId() int64`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *FilesSharingRemoteShare) GetFileIdOk() (*int64, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *FilesSharingRemoteShare) SetFileId(v int64)`

SetFileId sets FileId field to given value.


### SetFileIdNil

`func (o *FilesSharingRemoteShare) SetFileIdNil(b bool)`

 SetFileIdNil sets the value for FileId to be an explicit nil

### UnsetFileId
`func (o *FilesSharingRemoteShare) UnsetFileId()`

UnsetFileId ensures that no value is present for FileId, not even an explicit nil
### GetId

`func (o *FilesSharingRemoteShare) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FilesSharingRemoteShare) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FilesSharingRemoteShare) SetId(v int64)`

SetId sets Id field to given value.


### GetMimetype

`func (o *FilesSharingRemoteShare) GetMimetype() string`

GetMimetype returns the Mimetype field if non-nil, zero value otherwise.

### GetMimetypeOk

`func (o *FilesSharingRemoteShare) GetMimetypeOk() (*string, bool)`

GetMimetypeOk returns a tuple with the Mimetype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimetype

`func (o *FilesSharingRemoteShare) SetMimetype(v string)`

SetMimetype sets Mimetype field to given value.


### SetMimetypeNil

`func (o *FilesSharingRemoteShare) SetMimetypeNil(b bool)`

 SetMimetypeNil sets the value for Mimetype to be an explicit nil

### UnsetMimetype
`func (o *FilesSharingRemoteShare) UnsetMimetype()`

UnsetMimetype ensures that no value is present for Mimetype, not even an explicit nil
### GetMountpoint

`func (o *FilesSharingRemoteShare) GetMountpoint() string`

GetMountpoint returns the Mountpoint field if non-nil, zero value otherwise.

### GetMountpointOk

`func (o *FilesSharingRemoteShare) GetMountpointOk() (*string, bool)`

GetMountpointOk returns a tuple with the Mountpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountpoint

`func (o *FilesSharingRemoteShare) SetMountpoint(v string)`

SetMountpoint sets Mountpoint field to given value.


### GetMtime

`func (o *FilesSharingRemoteShare) GetMtime() int64`

GetMtime returns the Mtime field if non-nil, zero value otherwise.

### GetMtimeOk

`func (o *FilesSharingRemoteShare) GetMtimeOk() (*int64, bool)`

GetMtimeOk returns a tuple with the Mtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtime

`func (o *FilesSharingRemoteShare) SetMtime(v int64)`

SetMtime sets Mtime field to given value.


### SetMtimeNil

`func (o *FilesSharingRemoteShare) SetMtimeNil(b bool)`

 SetMtimeNil sets the value for Mtime to be an explicit nil

### UnsetMtime
`func (o *FilesSharingRemoteShare) UnsetMtime()`

UnsetMtime ensures that no value is present for Mtime, not even an explicit nil
### GetName

`func (o *FilesSharingRemoteShare) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FilesSharingRemoteShare) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FilesSharingRemoteShare) SetName(v string)`

SetName sets Name field to given value.


### GetOwner

`func (o *FilesSharingRemoteShare) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *FilesSharingRemoteShare) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *FilesSharingRemoteShare) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetParent

`func (o *FilesSharingRemoteShare) GetParent() int64`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *FilesSharingRemoteShare) GetParentOk() (*int64, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *FilesSharingRemoteShare) SetParent(v int64)`

SetParent sets Parent field to given value.


### SetParentNil

`func (o *FilesSharingRemoteShare) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *FilesSharingRemoteShare) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetPermissions

`func (o *FilesSharingRemoteShare) GetPermissions() int64`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *FilesSharingRemoteShare) GetPermissionsOk() (*int64, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *FilesSharingRemoteShare) SetPermissions(v int64)`

SetPermissions sets Permissions field to given value.


### SetPermissionsNil

`func (o *FilesSharingRemoteShare) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *FilesSharingRemoteShare) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
### GetRemote

`func (o *FilesSharingRemoteShare) GetRemote() string`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *FilesSharingRemoteShare) GetRemoteOk() (*string, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *FilesSharingRemoteShare) SetRemote(v string)`

SetRemote sets Remote field to given value.


### GetRemoteId

`func (o *FilesSharingRemoteShare) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *FilesSharingRemoteShare) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *FilesSharingRemoteShare) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.


### GetShareToken

`func (o *FilesSharingRemoteShare) GetShareToken() string`

GetShareToken returns the ShareToken field if non-nil, zero value otherwise.

### GetShareTokenOk

`func (o *FilesSharingRemoteShare) GetShareTokenOk() (*string, bool)`

GetShareTokenOk returns a tuple with the ShareToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareToken

`func (o *FilesSharingRemoteShare) SetShareToken(v string)`

SetShareToken sets ShareToken field to given value.


### GetShareType

`func (o *FilesSharingRemoteShare) GetShareType() int64`

GetShareType returns the ShareType field if non-nil, zero value otherwise.

### GetShareTypeOk

`func (o *FilesSharingRemoteShare) GetShareTypeOk() (*int64, bool)`

GetShareTypeOk returns a tuple with the ShareType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareType

`func (o *FilesSharingRemoteShare) SetShareType(v int64)`

SetShareType sets ShareType field to given value.


### GetType

`func (o *FilesSharingRemoteShare) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FilesSharingRemoteShare) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FilesSharingRemoteShare) SetType(v string)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *FilesSharingRemoteShare) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *FilesSharingRemoteShare) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUser

`func (o *FilesSharingRemoteShare) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *FilesSharingRemoteShare) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *FilesSharingRemoteShare) SetUser(v string)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


