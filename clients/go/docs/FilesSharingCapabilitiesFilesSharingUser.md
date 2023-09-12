# FilesSharingCapabilitiesFilesSharingUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SendMail** | **bool** |  | 
**ExpireDate** | Pointer to [**FilesSharingCapabilitiesFilesSharingUserExpireDate**](FilesSharingCapabilitiesFilesSharingUserExpireDate.md) |  | [optional] 

## Methods

### NewFilesSharingCapabilitiesFilesSharingUser

`func NewFilesSharingCapabilitiesFilesSharingUser(sendMail bool, ) *FilesSharingCapabilitiesFilesSharingUser`

NewFilesSharingCapabilitiesFilesSharingUser instantiates a new FilesSharingCapabilitiesFilesSharingUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilesSharingCapabilitiesFilesSharingUserWithDefaults

`func NewFilesSharingCapabilitiesFilesSharingUserWithDefaults() *FilesSharingCapabilitiesFilesSharingUser`

NewFilesSharingCapabilitiesFilesSharingUserWithDefaults instantiates a new FilesSharingCapabilitiesFilesSharingUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSendMail

`func (o *FilesSharingCapabilitiesFilesSharingUser) GetSendMail() bool`

GetSendMail returns the SendMail field if non-nil, zero value otherwise.

### GetSendMailOk

`func (o *FilesSharingCapabilitiesFilesSharingUser) GetSendMailOk() (*bool, bool)`

GetSendMailOk returns a tuple with the SendMail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendMail

`func (o *FilesSharingCapabilitiesFilesSharingUser) SetSendMail(v bool)`

SetSendMail sets SendMail field to given value.


### GetExpireDate

`func (o *FilesSharingCapabilitiesFilesSharingUser) GetExpireDate() FilesSharingCapabilitiesFilesSharingUserExpireDate`

GetExpireDate returns the ExpireDate field if non-nil, zero value otherwise.

### GetExpireDateOk

`func (o *FilesSharingCapabilitiesFilesSharingUser) GetExpireDateOk() (*FilesSharingCapabilitiesFilesSharingUserExpireDate, bool)`

GetExpireDateOk returns a tuple with the ExpireDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireDate

`func (o *FilesSharingCapabilitiesFilesSharingUser) SetExpireDate(v FilesSharingCapabilitiesFilesSharingUserExpireDate)`

SetExpireDate sets ExpireDate field to given value.

### HasExpireDate

`func (o *FilesSharingCapabilitiesFilesSharingUser) HasExpireDate() bool`

HasExpireDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


