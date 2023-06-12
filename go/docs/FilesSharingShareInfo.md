# FilesSharingShareInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**ParentId** | **int64** |  | 
**Mtime** | **int64** |  | 
**Name** | **string** |  | 
**Permissions** | **int64** |  | 
**Mimetype** | **string** |  | 
**Size** | **int64** |  | 
**Type** | **string** |  | 
**Etag** | **string** |  | 
**Children** | **[]map[string]map[string]interface{}** |  | 

## Methods

### NewFilesSharingShareInfo

`func NewFilesSharingShareInfo(id int64, parentId int64, mtime int64, name string, permissions int64, mimetype string, size int64, type_ string, etag string, children []map[string]map[string]interface{}, ) *FilesSharingShareInfo`

NewFilesSharingShareInfo instantiates a new FilesSharingShareInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilesSharingShareInfoWithDefaults

`func NewFilesSharingShareInfoWithDefaults() *FilesSharingShareInfo`

NewFilesSharingShareInfoWithDefaults instantiates a new FilesSharingShareInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FilesSharingShareInfo) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FilesSharingShareInfo) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FilesSharingShareInfo) SetId(v int64)`

SetId sets Id field to given value.


### GetParentId

`func (o *FilesSharingShareInfo) GetParentId() int64`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *FilesSharingShareInfo) GetParentIdOk() (*int64, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *FilesSharingShareInfo) SetParentId(v int64)`

SetParentId sets ParentId field to given value.


### GetMtime

`func (o *FilesSharingShareInfo) GetMtime() int64`

GetMtime returns the Mtime field if non-nil, zero value otherwise.

### GetMtimeOk

`func (o *FilesSharingShareInfo) GetMtimeOk() (*int64, bool)`

GetMtimeOk returns a tuple with the Mtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtime

`func (o *FilesSharingShareInfo) SetMtime(v int64)`

SetMtime sets Mtime field to given value.


### GetName

`func (o *FilesSharingShareInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FilesSharingShareInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FilesSharingShareInfo) SetName(v string)`

SetName sets Name field to given value.


### GetPermissions

`func (o *FilesSharingShareInfo) GetPermissions() int64`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *FilesSharingShareInfo) GetPermissionsOk() (*int64, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *FilesSharingShareInfo) SetPermissions(v int64)`

SetPermissions sets Permissions field to given value.


### GetMimetype

`func (o *FilesSharingShareInfo) GetMimetype() string`

GetMimetype returns the Mimetype field if non-nil, zero value otherwise.

### GetMimetypeOk

`func (o *FilesSharingShareInfo) GetMimetypeOk() (*string, bool)`

GetMimetypeOk returns a tuple with the Mimetype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimetype

`func (o *FilesSharingShareInfo) SetMimetype(v string)`

SetMimetype sets Mimetype field to given value.


### GetSize

`func (o *FilesSharingShareInfo) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *FilesSharingShareInfo) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *FilesSharingShareInfo) SetSize(v int64)`

SetSize sets Size field to given value.


### GetType

`func (o *FilesSharingShareInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FilesSharingShareInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FilesSharingShareInfo) SetType(v string)`

SetType sets Type field to given value.


### GetEtag

`func (o *FilesSharingShareInfo) GetEtag() string`

GetEtag returns the Etag field if non-nil, zero value otherwise.

### GetEtagOk

`func (o *FilesSharingShareInfo) GetEtagOk() (*string, bool)`

GetEtagOk returns a tuple with the Etag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtag

`func (o *FilesSharingShareInfo) SetEtag(v string)`

SetEtag sets Etag field to given value.


### GetChildren

`func (o *FilesSharingShareInfo) GetChildren() []map[string]map[string]interface{}`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *FilesSharingShareInfo) GetChildrenOk() (*[]map[string]map[string]interface{}, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *FilesSharingShareInfo) SetChildren(v []map[string]map[string]interface{})`

SetChildren sets Children field to given value.


### SetChildrenNil

`func (o *FilesSharingShareInfo) SetChildrenNil(b bool)`

 SetChildrenNil sets the value for Children to be an explicit nil

### UnsetChildren
`func (o *FilesSharingShareInfo) UnsetChildren()`

UnsetChildren ensures that no value is present for Children, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


