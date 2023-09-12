# FilesExternalStorageConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicableGroups** | Pointer to **[]string** |  | [optional] 
**ApplicableUsers** | Pointer to **[]string** |  | [optional] 
**AuthMechanism** | **string** |  | 
**Backend** | **string** |  | 
**BackendOptions** | **map[string]map[string]interface{}** |  | 
**Id** | Pointer to **int64** |  | [optional] 
**MountOptions** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**MountPoint** | **string** |  | 
**Priority** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **int64** |  | [optional] 
**StatusMessage** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 
**UserProvided** | **bool** |  | 

## Methods

### NewFilesExternalStorageConfig

`func NewFilesExternalStorageConfig(authMechanism string, backend string, backendOptions map[string]map[string]interface{}, mountPoint string, type_ string, userProvided bool, ) *FilesExternalStorageConfig`

NewFilesExternalStorageConfig instantiates a new FilesExternalStorageConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilesExternalStorageConfigWithDefaults

`func NewFilesExternalStorageConfigWithDefaults() *FilesExternalStorageConfig`

NewFilesExternalStorageConfigWithDefaults instantiates a new FilesExternalStorageConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicableGroups

`func (o *FilesExternalStorageConfig) GetApplicableGroups() []string`

GetApplicableGroups returns the ApplicableGroups field if non-nil, zero value otherwise.

### GetApplicableGroupsOk

`func (o *FilesExternalStorageConfig) GetApplicableGroupsOk() (*[]string, bool)`

GetApplicableGroupsOk returns a tuple with the ApplicableGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicableGroups

`func (o *FilesExternalStorageConfig) SetApplicableGroups(v []string)`

SetApplicableGroups sets ApplicableGroups field to given value.

### HasApplicableGroups

`func (o *FilesExternalStorageConfig) HasApplicableGroups() bool`

HasApplicableGroups returns a boolean if a field has been set.

### GetApplicableUsers

`func (o *FilesExternalStorageConfig) GetApplicableUsers() []string`

GetApplicableUsers returns the ApplicableUsers field if non-nil, zero value otherwise.

### GetApplicableUsersOk

`func (o *FilesExternalStorageConfig) GetApplicableUsersOk() (*[]string, bool)`

GetApplicableUsersOk returns a tuple with the ApplicableUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicableUsers

`func (o *FilesExternalStorageConfig) SetApplicableUsers(v []string)`

SetApplicableUsers sets ApplicableUsers field to given value.

### HasApplicableUsers

`func (o *FilesExternalStorageConfig) HasApplicableUsers() bool`

HasApplicableUsers returns a boolean if a field has been set.

### GetAuthMechanism

`func (o *FilesExternalStorageConfig) GetAuthMechanism() string`

GetAuthMechanism returns the AuthMechanism field if non-nil, zero value otherwise.

### GetAuthMechanismOk

`func (o *FilesExternalStorageConfig) GetAuthMechanismOk() (*string, bool)`

GetAuthMechanismOk returns a tuple with the AuthMechanism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMechanism

`func (o *FilesExternalStorageConfig) SetAuthMechanism(v string)`

SetAuthMechanism sets AuthMechanism field to given value.


### GetBackend

`func (o *FilesExternalStorageConfig) GetBackend() string`

GetBackend returns the Backend field if non-nil, zero value otherwise.

### GetBackendOk

`func (o *FilesExternalStorageConfig) GetBackendOk() (*string, bool)`

GetBackendOk returns a tuple with the Backend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackend

`func (o *FilesExternalStorageConfig) SetBackend(v string)`

SetBackend sets Backend field to given value.


### GetBackendOptions

`func (o *FilesExternalStorageConfig) GetBackendOptions() map[string]map[string]interface{}`

GetBackendOptions returns the BackendOptions field if non-nil, zero value otherwise.

### GetBackendOptionsOk

`func (o *FilesExternalStorageConfig) GetBackendOptionsOk() (*map[string]map[string]interface{}, bool)`

GetBackendOptionsOk returns a tuple with the BackendOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendOptions

`func (o *FilesExternalStorageConfig) SetBackendOptions(v map[string]map[string]interface{})`

SetBackendOptions sets BackendOptions field to given value.


### GetId

`func (o *FilesExternalStorageConfig) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FilesExternalStorageConfig) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FilesExternalStorageConfig) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *FilesExternalStorageConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMountOptions

`func (o *FilesExternalStorageConfig) GetMountOptions() map[string]map[string]interface{}`

GetMountOptions returns the MountOptions field if non-nil, zero value otherwise.

### GetMountOptionsOk

`func (o *FilesExternalStorageConfig) GetMountOptionsOk() (*map[string]map[string]interface{}, bool)`

GetMountOptionsOk returns a tuple with the MountOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountOptions

`func (o *FilesExternalStorageConfig) SetMountOptions(v map[string]map[string]interface{})`

SetMountOptions sets MountOptions field to given value.

### HasMountOptions

`func (o *FilesExternalStorageConfig) HasMountOptions() bool`

HasMountOptions returns a boolean if a field has been set.

### GetMountPoint

`func (o *FilesExternalStorageConfig) GetMountPoint() string`

GetMountPoint returns the MountPoint field if non-nil, zero value otherwise.

### GetMountPointOk

`func (o *FilesExternalStorageConfig) GetMountPointOk() (*string, bool)`

GetMountPointOk returns a tuple with the MountPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPoint

`func (o *FilesExternalStorageConfig) SetMountPoint(v string)`

SetMountPoint sets MountPoint field to given value.


### GetPriority

`func (o *FilesExternalStorageConfig) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *FilesExternalStorageConfig) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *FilesExternalStorageConfig) SetPriority(v int64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *FilesExternalStorageConfig) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetStatus

`func (o *FilesExternalStorageConfig) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FilesExternalStorageConfig) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FilesExternalStorageConfig) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FilesExternalStorageConfig) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *FilesExternalStorageConfig) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *FilesExternalStorageConfig) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *FilesExternalStorageConfig) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *FilesExternalStorageConfig) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetType

`func (o *FilesExternalStorageConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FilesExternalStorageConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FilesExternalStorageConfig) SetType(v string)`

SetType sets Type field to given value.


### GetUserProvided

`func (o *FilesExternalStorageConfig) GetUserProvided() bool`

GetUserProvided returns the UserProvided field if non-nil, zero value otherwise.

### GetUserProvidedOk

`func (o *FilesExternalStorageConfig) GetUserProvidedOk() (*bool, bool)`

GetUserProvidedOk returns a tuple with the UserProvided field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserProvided

`func (o *FilesExternalStorageConfig) SetUserProvided(v bool)`

SetUserProvided sets UserProvided field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


