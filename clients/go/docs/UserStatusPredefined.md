# UserStatusPredefined

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Icon** | **string** |  | 
**Message** | **string** |  | 
**ClearAt** | [**UserStatusClearAt**](UserStatusClearAt.md) |  | 
**Visible** | **NullableBool** |  | 

## Methods

### NewUserStatusPredefined

`func NewUserStatusPredefined(id string, icon string, message string, clearAt UserStatusClearAt, visible NullableBool, ) *UserStatusPredefined`

NewUserStatusPredefined instantiates a new UserStatusPredefined object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserStatusPredefinedWithDefaults

`func NewUserStatusPredefinedWithDefaults() *UserStatusPredefined`

NewUserStatusPredefinedWithDefaults instantiates a new UserStatusPredefined object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserStatusPredefined) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserStatusPredefined) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserStatusPredefined) SetId(v string)`

SetId sets Id field to given value.


### GetIcon

`func (o *UserStatusPredefined) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *UserStatusPredefined) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *UserStatusPredefined) SetIcon(v string)`

SetIcon sets Icon field to given value.


### GetMessage

`func (o *UserStatusPredefined) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UserStatusPredefined) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UserStatusPredefined) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetClearAt

`func (o *UserStatusPredefined) GetClearAt() UserStatusClearAt`

GetClearAt returns the ClearAt field if non-nil, zero value otherwise.

### GetClearAtOk

`func (o *UserStatusPredefined) GetClearAtOk() (*UserStatusClearAt, bool)`

GetClearAtOk returns a tuple with the ClearAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearAt

`func (o *UserStatusPredefined) SetClearAt(v UserStatusClearAt)`

SetClearAt sets ClearAt field to given value.


### GetVisible

`func (o *UserStatusPredefined) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *UserStatusPredefined) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *UserStatusPredefined) SetVisible(v bool)`

SetVisible sets Visible field to given value.


### SetVisibleNil

`func (o *UserStatusPredefined) SetVisibleNil(b bool)`

 SetVisibleNil sets the value for Visible to be an explicit nil

### UnsetVisible
`func (o *UserStatusPredefined) UnsetVisible()`

UnsetVisible ensures that no value is present for Visible, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


