# UserStatusPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** |  | 
**Message** | **NullableString** |  | 
**Icon** | **NullableString** |  | 
**ClearAt** | **NullableInt64** |  | 
**Status** | **string** |  | 

## Methods

### NewUserStatusPublic

`func NewUserStatusPublic(userId string, message NullableString, icon NullableString, clearAt NullableInt64, status string, ) *UserStatusPublic`

NewUserStatusPublic instantiates a new UserStatusPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserStatusPublicWithDefaults

`func NewUserStatusPublicWithDefaults() *UserStatusPublic`

NewUserStatusPublicWithDefaults instantiates a new UserStatusPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *UserStatusPublic) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserStatusPublic) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserStatusPublic) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetMessage

`func (o *UserStatusPublic) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UserStatusPublic) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UserStatusPublic) SetMessage(v string)`

SetMessage sets Message field to given value.


### SetMessageNil

`func (o *UserStatusPublic) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *UserStatusPublic) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetIcon

`func (o *UserStatusPublic) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *UserStatusPublic) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *UserStatusPublic) SetIcon(v string)`

SetIcon sets Icon field to given value.


### SetIconNil

`func (o *UserStatusPublic) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *UserStatusPublic) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil
### GetClearAt

`func (o *UserStatusPublic) GetClearAt() int64`

GetClearAt returns the ClearAt field if non-nil, zero value otherwise.

### GetClearAtOk

`func (o *UserStatusPublic) GetClearAtOk() (*int64, bool)`

GetClearAtOk returns a tuple with the ClearAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearAt

`func (o *UserStatusPublic) SetClearAt(v int64)`

SetClearAt sets ClearAt field to given value.


### SetClearAtNil

`func (o *UserStatusPublic) SetClearAtNil(b bool)`

 SetClearAtNil sets the value for ClearAt to be an explicit nil

### UnsetClearAt
`func (o *UserStatusPublic) UnsetClearAt()`

UnsetClearAt ensures that no value is present for ClearAt, not even an explicit nil
### GetStatus

`func (o *UserStatusPublic) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserStatusPublic) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserStatusPublic) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


