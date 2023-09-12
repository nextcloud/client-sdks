# FilesCapabilitiesFiles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bigfilechunking** | **bool** |  | 
**BlacklistedFiles** | **[]map[string]interface{}** |  | 
**DirectEditing** | [**FilesCapabilitiesFilesDirectEditing**](FilesCapabilitiesFilesDirectEditing.md) |  | 

## Methods

### NewFilesCapabilitiesFiles

`func NewFilesCapabilitiesFiles(bigfilechunking bool, blacklistedFiles []map[string]interface{}, directEditing FilesCapabilitiesFilesDirectEditing, ) *FilesCapabilitiesFiles`

NewFilesCapabilitiesFiles instantiates a new FilesCapabilitiesFiles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilesCapabilitiesFilesWithDefaults

`func NewFilesCapabilitiesFilesWithDefaults() *FilesCapabilitiesFiles`

NewFilesCapabilitiesFilesWithDefaults instantiates a new FilesCapabilitiesFiles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBigfilechunking

`func (o *FilesCapabilitiesFiles) GetBigfilechunking() bool`

GetBigfilechunking returns the Bigfilechunking field if non-nil, zero value otherwise.

### GetBigfilechunkingOk

`func (o *FilesCapabilitiesFiles) GetBigfilechunkingOk() (*bool, bool)`

GetBigfilechunkingOk returns a tuple with the Bigfilechunking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBigfilechunking

`func (o *FilesCapabilitiesFiles) SetBigfilechunking(v bool)`

SetBigfilechunking sets Bigfilechunking field to given value.


### GetBlacklistedFiles

`func (o *FilesCapabilitiesFiles) GetBlacklistedFiles() []map[string]interface{}`

GetBlacklistedFiles returns the BlacklistedFiles field if non-nil, zero value otherwise.

### GetBlacklistedFilesOk

`func (o *FilesCapabilitiesFiles) GetBlacklistedFilesOk() (*[]map[string]interface{}, bool)`

GetBlacklistedFilesOk returns a tuple with the BlacklistedFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlacklistedFiles

`func (o *FilesCapabilitiesFiles) SetBlacklistedFiles(v []map[string]interface{})`

SetBlacklistedFiles sets BlacklistedFiles field to given value.


### GetDirectEditing

`func (o *FilesCapabilitiesFiles) GetDirectEditing() FilesCapabilitiesFilesDirectEditing`

GetDirectEditing returns the DirectEditing field if non-nil, zero value otherwise.

### GetDirectEditingOk

`func (o *FilesCapabilitiesFiles) GetDirectEditingOk() (*FilesCapabilitiesFilesDirectEditing, bool)`

GetDirectEditingOk returns a tuple with the DirectEditing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectEditing

`func (o *FilesCapabilitiesFiles) SetDirectEditing(v FilesCapabilitiesFilesDirectEditing)`

SetDirectEditing sets DirectEditing field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


