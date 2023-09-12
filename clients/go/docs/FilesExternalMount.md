# FilesExternalMount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Path** | **string** |  | 
**Type** | **string** |  | 
**Backend** | **string** |  | 
**Scope** | **string** |  | 
**Permissions** | **int64** |  | 
**Id** | **int64** |  | 
**Class** | **string** |  | 
**Config** | [**FilesExternalStorageConfig**](FilesExternalStorageConfig.md) |  | 

## Methods

### NewFilesExternalMount

`func NewFilesExternalMount(name string, path string, type_ string, backend string, scope string, permissions int64, id int64, class string, config FilesExternalStorageConfig, ) *FilesExternalMount`

NewFilesExternalMount instantiates a new FilesExternalMount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilesExternalMountWithDefaults

`func NewFilesExternalMountWithDefaults() *FilesExternalMount`

NewFilesExternalMountWithDefaults instantiates a new FilesExternalMount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FilesExternalMount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FilesExternalMount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FilesExternalMount) SetName(v string)`

SetName sets Name field to given value.


### GetPath

`func (o *FilesExternalMount) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *FilesExternalMount) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *FilesExternalMount) SetPath(v string)`

SetPath sets Path field to given value.


### GetType

`func (o *FilesExternalMount) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FilesExternalMount) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FilesExternalMount) SetType(v string)`

SetType sets Type field to given value.


### GetBackend

`func (o *FilesExternalMount) GetBackend() string`

GetBackend returns the Backend field if non-nil, zero value otherwise.

### GetBackendOk

`func (o *FilesExternalMount) GetBackendOk() (*string, bool)`

GetBackendOk returns a tuple with the Backend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackend

`func (o *FilesExternalMount) SetBackend(v string)`

SetBackend sets Backend field to given value.


### GetScope

`func (o *FilesExternalMount) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *FilesExternalMount) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *FilesExternalMount) SetScope(v string)`

SetScope sets Scope field to given value.


### GetPermissions

`func (o *FilesExternalMount) GetPermissions() int64`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *FilesExternalMount) GetPermissionsOk() (*int64, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *FilesExternalMount) SetPermissions(v int64)`

SetPermissions sets Permissions field to given value.


### GetId

`func (o *FilesExternalMount) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FilesExternalMount) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FilesExternalMount) SetId(v int64)`

SetId sets Id field to given value.


### GetClass

`func (o *FilesExternalMount) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *FilesExternalMount) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *FilesExternalMount) SetClass(v string)`

SetClass sets Class field to given value.


### GetConfig

`func (o *FilesExternalMount) GetConfig() FilesExternalStorageConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *FilesExternalMount) GetConfigOk() (*FilesExternalStorageConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *FilesExternalMount) SetConfig(v FilesExternalStorageConfig)`

SetConfig sets Config field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


