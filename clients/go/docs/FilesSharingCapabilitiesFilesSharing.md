# FilesSharingCapabilitiesFilesSharing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiEnabled** | **bool** |  | 
**Public** | [**FilesSharingCapabilitiesFilesSharingPublic**](FilesSharingCapabilitiesFilesSharingPublic.md) |  | 
**User** | [**FilesSharingCapabilitiesFilesSharingUser**](FilesSharingCapabilitiesFilesSharingUser.md) |  | 
**Resharing** | **bool** |  | 
**GroupSharing** | Pointer to **bool** |  | [optional] 
**Group** | Pointer to [**FilesSharingCapabilitiesFilesSharingGroup**](FilesSharingCapabilitiesFilesSharingGroup.md) |  | [optional] 
**DefaultPermissions** | Pointer to **int64** |  | [optional] 
**Federation** | [**FilesSharingCapabilitiesFilesSharingFederation**](FilesSharingCapabilitiesFilesSharingFederation.md) |  | 
**Sharee** | [**FilesSharingCapabilitiesFilesSharingSharee**](FilesSharingCapabilitiesFilesSharingSharee.md) |  | 

## Methods

### NewFilesSharingCapabilitiesFilesSharing

`func NewFilesSharingCapabilitiesFilesSharing(apiEnabled bool, public FilesSharingCapabilitiesFilesSharingPublic, user FilesSharingCapabilitiesFilesSharingUser, resharing bool, federation FilesSharingCapabilitiesFilesSharingFederation, sharee FilesSharingCapabilitiesFilesSharingSharee, ) *FilesSharingCapabilitiesFilesSharing`

NewFilesSharingCapabilitiesFilesSharing instantiates a new FilesSharingCapabilitiesFilesSharing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilesSharingCapabilitiesFilesSharingWithDefaults

`func NewFilesSharingCapabilitiesFilesSharingWithDefaults() *FilesSharingCapabilitiesFilesSharing`

NewFilesSharingCapabilitiesFilesSharingWithDefaults instantiates a new FilesSharingCapabilitiesFilesSharing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiEnabled

`func (o *FilesSharingCapabilitiesFilesSharing) GetApiEnabled() bool`

GetApiEnabled returns the ApiEnabled field if non-nil, zero value otherwise.

### GetApiEnabledOk

`func (o *FilesSharingCapabilitiesFilesSharing) GetApiEnabledOk() (*bool, bool)`

GetApiEnabledOk returns a tuple with the ApiEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiEnabled

`func (o *FilesSharingCapabilitiesFilesSharing) SetApiEnabled(v bool)`

SetApiEnabled sets ApiEnabled field to given value.


### GetPublic

`func (o *FilesSharingCapabilitiesFilesSharing) GetPublic() FilesSharingCapabilitiesFilesSharingPublic`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *FilesSharingCapabilitiesFilesSharing) GetPublicOk() (*FilesSharingCapabilitiesFilesSharingPublic, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *FilesSharingCapabilitiesFilesSharing) SetPublic(v FilesSharingCapabilitiesFilesSharingPublic)`

SetPublic sets Public field to given value.


### GetUser

`func (o *FilesSharingCapabilitiesFilesSharing) GetUser() FilesSharingCapabilitiesFilesSharingUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *FilesSharingCapabilitiesFilesSharing) GetUserOk() (*FilesSharingCapabilitiesFilesSharingUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *FilesSharingCapabilitiesFilesSharing) SetUser(v FilesSharingCapabilitiesFilesSharingUser)`

SetUser sets User field to given value.


### GetResharing

`func (o *FilesSharingCapabilitiesFilesSharing) GetResharing() bool`

GetResharing returns the Resharing field if non-nil, zero value otherwise.

### GetResharingOk

`func (o *FilesSharingCapabilitiesFilesSharing) GetResharingOk() (*bool, bool)`

GetResharingOk returns a tuple with the Resharing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResharing

`func (o *FilesSharingCapabilitiesFilesSharing) SetResharing(v bool)`

SetResharing sets Resharing field to given value.


### GetGroupSharing

`func (o *FilesSharingCapabilitiesFilesSharing) GetGroupSharing() bool`

GetGroupSharing returns the GroupSharing field if non-nil, zero value otherwise.

### GetGroupSharingOk

`func (o *FilesSharingCapabilitiesFilesSharing) GetGroupSharingOk() (*bool, bool)`

GetGroupSharingOk returns a tuple with the GroupSharing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupSharing

`func (o *FilesSharingCapabilitiesFilesSharing) SetGroupSharing(v bool)`

SetGroupSharing sets GroupSharing field to given value.

### HasGroupSharing

`func (o *FilesSharingCapabilitiesFilesSharing) HasGroupSharing() bool`

HasGroupSharing returns a boolean if a field has been set.

### GetGroup

`func (o *FilesSharingCapabilitiesFilesSharing) GetGroup() FilesSharingCapabilitiesFilesSharingGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *FilesSharingCapabilitiesFilesSharing) GetGroupOk() (*FilesSharingCapabilitiesFilesSharingGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *FilesSharingCapabilitiesFilesSharing) SetGroup(v FilesSharingCapabilitiesFilesSharingGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *FilesSharingCapabilitiesFilesSharing) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetDefaultPermissions

`func (o *FilesSharingCapabilitiesFilesSharing) GetDefaultPermissions() int64`

GetDefaultPermissions returns the DefaultPermissions field if non-nil, zero value otherwise.

### GetDefaultPermissionsOk

`func (o *FilesSharingCapabilitiesFilesSharing) GetDefaultPermissionsOk() (*int64, bool)`

GetDefaultPermissionsOk returns a tuple with the DefaultPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPermissions

`func (o *FilesSharingCapabilitiesFilesSharing) SetDefaultPermissions(v int64)`

SetDefaultPermissions sets DefaultPermissions field to given value.

### HasDefaultPermissions

`func (o *FilesSharingCapabilitiesFilesSharing) HasDefaultPermissions() bool`

HasDefaultPermissions returns a boolean if a field has been set.

### GetFederation

`func (o *FilesSharingCapabilitiesFilesSharing) GetFederation() FilesSharingCapabilitiesFilesSharingFederation`

GetFederation returns the Federation field if non-nil, zero value otherwise.

### GetFederationOk

`func (o *FilesSharingCapabilitiesFilesSharing) GetFederationOk() (*FilesSharingCapabilitiesFilesSharingFederation, bool)`

GetFederationOk returns a tuple with the Federation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederation

`func (o *FilesSharingCapabilitiesFilesSharing) SetFederation(v FilesSharingCapabilitiesFilesSharingFederation)`

SetFederation sets Federation field to given value.


### GetSharee

`func (o *FilesSharingCapabilitiesFilesSharing) GetSharee() FilesSharingCapabilitiesFilesSharingSharee`

GetSharee returns the Sharee field if non-nil, zero value otherwise.

### GetShareeOk

`func (o *FilesSharingCapabilitiesFilesSharing) GetShareeOk() (*FilesSharingCapabilitiesFilesSharingSharee, bool)`

GetShareeOk returns a tuple with the Sharee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharee

`func (o *FilesSharingCapabilitiesFilesSharing) SetSharee(v FilesSharingCapabilitiesFilesSharingSharee)`

SetSharee sets Sharee field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


