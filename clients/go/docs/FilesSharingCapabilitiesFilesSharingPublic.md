# FilesSharingCapabilitiesFilesSharingPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** |  | 
**Password** | Pointer to [**FilesSharingCapabilitiesFilesSharingPublicPassword**](FilesSharingCapabilitiesFilesSharingPublicPassword.md) |  | [optional] 
**MultipleLinks** | Pointer to **bool** |  | [optional] 
**ExpireDate** | Pointer to [**FilesSharingCapabilitiesFilesSharingPublicExpireDate**](FilesSharingCapabilitiesFilesSharingPublicExpireDate.md) |  | [optional] 
**ExpireDateInternal** | Pointer to [**FilesSharingCapabilitiesFilesSharingPublicExpireDate**](FilesSharingCapabilitiesFilesSharingPublicExpireDate.md) |  | [optional] 
**ExpireDateRemote** | Pointer to [**FilesSharingCapabilitiesFilesSharingPublicExpireDate**](FilesSharingCapabilitiesFilesSharingPublicExpireDate.md) |  | [optional] 
**SendMail** | Pointer to **bool** |  | [optional] 
**Upload** | Pointer to **bool** |  | [optional] 
**UploadFilesDrop** | Pointer to **bool** |  | [optional] 

## Methods

### NewFilesSharingCapabilitiesFilesSharingPublic

`func NewFilesSharingCapabilitiesFilesSharingPublic(enabled bool, ) *FilesSharingCapabilitiesFilesSharingPublic`

NewFilesSharingCapabilitiesFilesSharingPublic instantiates a new FilesSharingCapabilitiesFilesSharingPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilesSharingCapabilitiesFilesSharingPublicWithDefaults

`func NewFilesSharingCapabilitiesFilesSharingPublicWithDefaults() *FilesSharingCapabilitiesFilesSharingPublic`

NewFilesSharingCapabilitiesFilesSharingPublicWithDefaults instantiates a new FilesSharingCapabilitiesFilesSharingPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *FilesSharingCapabilitiesFilesSharingPublic) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FilesSharingCapabilitiesFilesSharingPublic) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FilesSharingCapabilitiesFilesSharingPublic) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetPassword

`func (o *FilesSharingCapabilitiesFilesSharingPublic) GetPassword() FilesSharingCapabilitiesFilesSharingPublicPassword`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *FilesSharingCapabilitiesFilesSharingPublic) GetPasswordOk() (*FilesSharingCapabilitiesFilesSharingPublicPassword, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *FilesSharingCapabilitiesFilesSharingPublic) SetPassword(v FilesSharingCapabilitiesFilesSharingPublicPassword)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *FilesSharingCapabilitiesFilesSharingPublic) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetMultipleLinks

`func (o *FilesSharingCapabilitiesFilesSharingPublic) GetMultipleLinks() bool`

GetMultipleLinks returns the MultipleLinks field if non-nil, zero value otherwise.

### GetMultipleLinksOk

`func (o *FilesSharingCapabilitiesFilesSharingPublic) GetMultipleLinksOk() (*bool, bool)`

GetMultipleLinksOk returns a tuple with the MultipleLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleLinks

`func (o *FilesSharingCapabilitiesFilesSharingPublic) SetMultipleLinks(v bool)`

SetMultipleLinks sets MultipleLinks field to given value.

### HasMultipleLinks

`func (o *FilesSharingCapabilitiesFilesSharingPublic) HasMultipleLinks() bool`

HasMultipleLinks returns a boolean if a field has been set.

### GetExpireDate

`func (o *FilesSharingCapabilitiesFilesSharingPublic) GetExpireDate() FilesSharingCapabilitiesFilesSharingPublicExpireDate`

GetExpireDate returns the ExpireDate field if non-nil, zero value otherwise.

### GetExpireDateOk

`func (o *FilesSharingCapabilitiesFilesSharingPublic) GetExpireDateOk() (*FilesSharingCapabilitiesFilesSharingPublicExpireDate, bool)`

GetExpireDateOk returns a tuple with the ExpireDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireDate

`func (o *FilesSharingCapabilitiesFilesSharingPublic) SetExpireDate(v FilesSharingCapabilitiesFilesSharingPublicExpireDate)`

SetExpireDate sets ExpireDate field to given value.

### HasExpireDate

`func (o *FilesSharingCapabilitiesFilesSharingPublic) HasExpireDate() bool`

HasExpireDate returns a boolean if a field has been set.

### GetExpireDateInternal

`func (o *FilesSharingCapabilitiesFilesSharingPublic) GetExpireDateInternal() FilesSharingCapabilitiesFilesSharingPublicExpireDate`

GetExpireDateInternal returns the ExpireDateInternal field if non-nil, zero value otherwise.

### GetExpireDateInternalOk

`func (o *FilesSharingCapabilitiesFilesSharingPublic) GetExpireDateInternalOk() (*FilesSharingCapabilitiesFilesSharingPublicExpireDate, bool)`

GetExpireDateInternalOk returns a tuple with the ExpireDateInternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireDateInternal

`func (o *FilesSharingCapabilitiesFilesSharingPublic) SetExpireDateInternal(v FilesSharingCapabilitiesFilesSharingPublicExpireDate)`

SetExpireDateInternal sets ExpireDateInternal field to given value.

### HasExpireDateInternal

`func (o *FilesSharingCapabilitiesFilesSharingPublic) HasExpireDateInternal() bool`

HasExpireDateInternal returns a boolean if a field has been set.

### GetExpireDateRemote

`func (o *FilesSharingCapabilitiesFilesSharingPublic) GetExpireDateRemote() FilesSharingCapabilitiesFilesSharingPublicExpireDate`

GetExpireDateRemote returns the ExpireDateRemote field if non-nil, zero value otherwise.

### GetExpireDateRemoteOk

`func (o *FilesSharingCapabilitiesFilesSharingPublic) GetExpireDateRemoteOk() (*FilesSharingCapabilitiesFilesSharingPublicExpireDate, bool)`

GetExpireDateRemoteOk returns a tuple with the ExpireDateRemote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireDateRemote

`func (o *FilesSharingCapabilitiesFilesSharingPublic) SetExpireDateRemote(v FilesSharingCapabilitiesFilesSharingPublicExpireDate)`

SetExpireDateRemote sets ExpireDateRemote field to given value.

### HasExpireDateRemote

`func (o *FilesSharingCapabilitiesFilesSharingPublic) HasExpireDateRemote() bool`

HasExpireDateRemote returns a boolean if a field has been set.

### GetSendMail

`func (o *FilesSharingCapabilitiesFilesSharingPublic) GetSendMail() bool`

GetSendMail returns the SendMail field if non-nil, zero value otherwise.

### GetSendMailOk

`func (o *FilesSharingCapabilitiesFilesSharingPublic) GetSendMailOk() (*bool, bool)`

GetSendMailOk returns a tuple with the SendMail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendMail

`func (o *FilesSharingCapabilitiesFilesSharingPublic) SetSendMail(v bool)`

SetSendMail sets SendMail field to given value.

### HasSendMail

`func (o *FilesSharingCapabilitiesFilesSharingPublic) HasSendMail() bool`

HasSendMail returns a boolean if a field has been set.

### GetUpload

`func (o *FilesSharingCapabilitiesFilesSharingPublic) GetUpload() bool`

GetUpload returns the Upload field if non-nil, zero value otherwise.

### GetUploadOk

`func (o *FilesSharingCapabilitiesFilesSharingPublic) GetUploadOk() (*bool, bool)`

GetUploadOk returns a tuple with the Upload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpload

`func (o *FilesSharingCapabilitiesFilesSharingPublic) SetUpload(v bool)`

SetUpload sets Upload field to given value.

### HasUpload

`func (o *FilesSharingCapabilitiesFilesSharingPublic) HasUpload() bool`

HasUpload returns a boolean if a field has been set.

### GetUploadFilesDrop

`func (o *FilesSharingCapabilitiesFilesSharingPublic) GetUploadFilesDrop() bool`

GetUploadFilesDrop returns the UploadFilesDrop field if non-nil, zero value otherwise.

### GetUploadFilesDropOk

`func (o *FilesSharingCapabilitiesFilesSharingPublic) GetUploadFilesDropOk() (*bool, bool)`

GetUploadFilesDropOk returns a tuple with the UploadFilesDrop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadFilesDrop

`func (o *FilesSharingCapabilitiesFilesSharingPublic) SetUploadFilesDrop(v bool)`

SetUploadFilesDrop sets UploadFilesDrop field to given value.

### HasUploadFilesDrop

`func (o *FilesSharingCapabilitiesFilesSharingPublic) HasUploadFilesDrop() bool`

HasUploadFilesDrop returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


