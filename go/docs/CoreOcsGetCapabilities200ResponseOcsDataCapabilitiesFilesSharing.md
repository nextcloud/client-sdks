# CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiEnabled** | **bool** |  | 
**Public** | [**CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingPublic**](CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingPublic.md) |  | 
**User** | [**CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUser**](CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUser.md) |  | 
**Resharing** | **bool** |  | 
**GroupSharing** | Pointer to **bool** |  | [optional] 
**Group** | Pointer to [**CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingGroup**](CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingGroup.md) |  | [optional] 
**DefaultPermissions** | Pointer to **int64** |  | [optional] 
**Federation** | [**CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingFederation**](CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingFederation.md) |  | 
**Sharee** | [**CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingSharee**](CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingSharee.md) |  | 

## Methods

### NewCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing

`func NewCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing(apiEnabled bool, public CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingPublic, user CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUser, resharing bool, federation CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingFederation, sharee CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingSharee, ) *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing`

NewCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing instantiates a new CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingWithDefaults

`func NewCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingWithDefaults() *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing`

NewCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingWithDefaults instantiates a new CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiEnabled

`func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing) GetApiEnabled() bool`

GetApiEnabled returns the ApiEnabled field if non-nil, zero value otherwise.

### GetApiEnabledOk

`func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing) GetApiEnabledOk() (*bool, bool)`

GetApiEnabledOk returns a tuple with the ApiEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiEnabled

`func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing) SetApiEnabled(v bool)`

SetApiEnabled sets ApiEnabled field to given value.


### GetPublic

`func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing) GetPublic() CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingPublic`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing) GetPublicOk() (*CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingPublic, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing) SetPublic(v CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingPublic)`

SetPublic sets Public field to given value.


### GetUser

`func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing) GetUser() CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing) GetUserOk() (*CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing) SetUser(v CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUser)`

SetUser sets User field to given value.


### GetResharing

`func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing) GetResharing() bool`

GetResharing returns the Resharing field if non-nil, zero value otherwise.

### GetResharingOk

`func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing) GetResharingOk() (*bool, bool)`

GetResharingOk returns a tuple with the Resharing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResharing

`func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing) SetResharing(v bool)`

SetResharing sets Resharing field to given value.


### GetGroupSharing

`func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing) GetGroupSharing() bool`

GetGroupSharing returns the GroupSharing field if non-nil, zero value otherwise.

### GetGroupSharingOk

`func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing) GetGroupSharingOk() (*bool, bool)`

GetGroupSharingOk returns a tuple with the GroupSharing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupSharing

`func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing) SetGroupSharing(v bool)`

SetGroupSharing sets GroupSharing field to given value.

### HasGroupSharing

`func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing) HasGroupSharing() bool`

HasGroupSharing returns a boolean if a field has been set.

### GetGroup

`func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing) GetGroup() CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing) GetGroupOk() (*CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing) SetGroup(v CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetDefaultPermissions

`func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing) GetDefaultPermissions() int64`

GetDefaultPermissions returns the DefaultPermissions field if non-nil, zero value otherwise.

### GetDefaultPermissionsOk

`func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing) GetDefaultPermissionsOk() (*int64, bool)`

GetDefaultPermissionsOk returns a tuple with the DefaultPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPermissions

`func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing) SetDefaultPermissions(v int64)`

SetDefaultPermissions sets DefaultPermissions field to given value.

### HasDefaultPermissions

`func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing) HasDefaultPermissions() bool`

HasDefaultPermissions returns a boolean if a field has been set.

### GetFederation

`func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing) GetFederation() CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingFederation`

GetFederation returns the Federation field if non-nil, zero value otherwise.

### GetFederationOk

`func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing) GetFederationOk() (*CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingFederation, bool)`

GetFederationOk returns a tuple with the Federation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederation

`func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing) SetFederation(v CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingFederation)`

SetFederation sets Federation field to given value.


### GetSharee

`func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing) GetSharee() CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingSharee`

GetSharee returns the Sharee field if non-nil, zero value otherwise.

### GetShareeOk

`func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing) GetShareeOk() (*CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingSharee, bool)`

GetShareeOk returns a tuple with the Sharee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharee

`func (o *CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing) SetSharee(v CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingSharee)`

SetSharee sets Sharee field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


