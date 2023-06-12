# FilesSharingShareesSearchResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exact** | [**FilesSharingShareesSearchResultExact**](FilesSharingShareesSearchResultExact.md) |  | 
**Circles** | [**[]FilesSharingShareeCircle**](FilesSharingShareeCircle.md) |  | 
**Emails** | [**[]FilesSharingShareeEmail**](FilesSharingShareeEmail.md) |  | 
**Groups** | [**[]FilesSharingSharee**](FilesSharingSharee.md) |  | 
**Lookup** | [**[]FilesSharingShareeLookup**](FilesSharingShareeLookup.md) |  | 
**RemoteGroups** | [**[]FilesSharingShareeRemoteGroup**](FilesSharingShareeRemoteGroup.md) |  | 
**Remotes** | [**[]FilesSharingShareeRemote**](FilesSharingShareeRemote.md) |  | 
**Rooms** | [**[]FilesSharingSharee**](FilesSharingSharee.md) |  | 
**Users** | [**[]FilesSharingShareeUser**](FilesSharingShareeUser.md) |  | 
**LookupEnabled** | **bool** |  | 

## Methods

### NewFilesSharingShareesSearchResult

`func NewFilesSharingShareesSearchResult(exact FilesSharingShareesSearchResultExact, circles []FilesSharingShareeCircle, emails []FilesSharingShareeEmail, groups []FilesSharingSharee, lookup []FilesSharingShareeLookup, remoteGroups []FilesSharingShareeRemoteGroup, remotes []FilesSharingShareeRemote, rooms []FilesSharingSharee, users []FilesSharingShareeUser, lookupEnabled bool, ) *FilesSharingShareesSearchResult`

NewFilesSharingShareesSearchResult instantiates a new FilesSharingShareesSearchResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilesSharingShareesSearchResultWithDefaults

`func NewFilesSharingShareesSearchResultWithDefaults() *FilesSharingShareesSearchResult`

NewFilesSharingShareesSearchResultWithDefaults instantiates a new FilesSharingShareesSearchResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExact

`func (o *FilesSharingShareesSearchResult) GetExact() FilesSharingShareesSearchResultExact`

GetExact returns the Exact field if non-nil, zero value otherwise.

### GetExactOk

`func (o *FilesSharingShareesSearchResult) GetExactOk() (*FilesSharingShareesSearchResultExact, bool)`

GetExactOk returns a tuple with the Exact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExact

`func (o *FilesSharingShareesSearchResult) SetExact(v FilesSharingShareesSearchResultExact)`

SetExact sets Exact field to given value.


### GetCircles

`func (o *FilesSharingShareesSearchResult) GetCircles() []FilesSharingShareeCircle`

GetCircles returns the Circles field if non-nil, zero value otherwise.

### GetCirclesOk

`func (o *FilesSharingShareesSearchResult) GetCirclesOk() (*[]FilesSharingShareeCircle, bool)`

GetCirclesOk returns a tuple with the Circles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircles

`func (o *FilesSharingShareesSearchResult) SetCircles(v []FilesSharingShareeCircle)`

SetCircles sets Circles field to given value.


### GetEmails

`func (o *FilesSharingShareesSearchResult) GetEmails() []FilesSharingShareeEmail`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *FilesSharingShareesSearchResult) GetEmailsOk() (*[]FilesSharingShareeEmail, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *FilesSharingShareesSearchResult) SetEmails(v []FilesSharingShareeEmail)`

SetEmails sets Emails field to given value.


### GetGroups

`func (o *FilesSharingShareesSearchResult) GetGroups() []FilesSharingSharee`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *FilesSharingShareesSearchResult) GetGroupsOk() (*[]FilesSharingSharee, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *FilesSharingShareesSearchResult) SetGroups(v []FilesSharingSharee)`

SetGroups sets Groups field to given value.


### GetLookup

`func (o *FilesSharingShareesSearchResult) GetLookup() []FilesSharingShareeLookup`

GetLookup returns the Lookup field if non-nil, zero value otherwise.

### GetLookupOk

`func (o *FilesSharingShareesSearchResult) GetLookupOk() (*[]FilesSharingShareeLookup, bool)`

GetLookupOk returns a tuple with the Lookup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookup

`func (o *FilesSharingShareesSearchResult) SetLookup(v []FilesSharingShareeLookup)`

SetLookup sets Lookup field to given value.


### GetRemoteGroups

`func (o *FilesSharingShareesSearchResult) GetRemoteGroups() []FilesSharingShareeRemoteGroup`

GetRemoteGroups returns the RemoteGroups field if non-nil, zero value otherwise.

### GetRemoteGroupsOk

`func (o *FilesSharingShareesSearchResult) GetRemoteGroupsOk() (*[]FilesSharingShareeRemoteGroup, bool)`

GetRemoteGroupsOk returns a tuple with the RemoteGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteGroups

`func (o *FilesSharingShareesSearchResult) SetRemoteGroups(v []FilesSharingShareeRemoteGroup)`

SetRemoteGroups sets RemoteGroups field to given value.


### GetRemotes

`func (o *FilesSharingShareesSearchResult) GetRemotes() []FilesSharingShareeRemote`

GetRemotes returns the Remotes field if non-nil, zero value otherwise.

### GetRemotesOk

`func (o *FilesSharingShareesSearchResult) GetRemotesOk() (*[]FilesSharingShareeRemote, bool)`

GetRemotesOk returns a tuple with the Remotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotes

`func (o *FilesSharingShareesSearchResult) SetRemotes(v []FilesSharingShareeRemote)`

SetRemotes sets Remotes field to given value.


### GetRooms

`func (o *FilesSharingShareesSearchResult) GetRooms() []FilesSharingSharee`

GetRooms returns the Rooms field if non-nil, zero value otherwise.

### GetRoomsOk

`func (o *FilesSharingShareesSearchResult) GetRoomsOk() (*[]FilesSharingSharee, bool)`

GetRoomsOk returns a tuple with the Rooms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRooms

`func (o *FilesSharingShareesSearchResult) SetRooms(v []FilesSharingSharee)`

SetRooms sets Rooms field to given value.


### GetUsers

`func (o *FilesSharingShareesSearchResult) GetUsers() []FilesSharingShareeUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *FilesSharingShareesSearchResult) GetUsersOk() (*[]FilesSharingShareeUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *FilesSharingShareesSearchResult) SetUsers(v []FilesSharingShareeUser)`

SetUsers sets Users field to given value.


### GetLookupEnabled

`func (o *FilesSharingShareesSearchResult) GetLookupEnabled() bool`

GetLookupEnabled returns the LookupEnabled field if non-nil, zero value otherwise.

### GetLookupEnabledOk

`func (o *FilesSharingShareesSearchResult) GetLookupEnabledOk() (*bool, bool)`

GetLookupEnabledOk returns a tuple with the LookupEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookupEnabled

`func (o *FilesSharingShareesSearchResult) SetLookupEnabled(v bool)`

SetLookupEnabled sets LookupEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


