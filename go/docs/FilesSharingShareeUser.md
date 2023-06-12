# FilesSharingShareeUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **NullableInt64** |  | 
**Label** | **string** |  | 
**Subline** | **string** |  | 
**Icon** | **string** |  | 
**ShareWithDisplayNameUnique** | **string** |  | 
**Status** | [**FilesSharingShareeUserAllOfStatus**](FilesSharingShareeUserAllOfStatus.md) |  | 
**Value** | [**FilesSharingShareeValue**](FilesSharingShareeValue.md) |  | 

## Methods

### NewFilesSharingShareeUser

`func NewFilesSharingShareeUser(count NullableInt64, label string, subline string, icon string, shareWithDisplayNameUnique string, status FilesSharingShareeUserAllOfStatus, value FilesSharingShareeValue, ) *FilesSharingShareeUser`

NewFilesSharingShareeUser instantiates a new FilesSharingShareeUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilesSharingShareeUserWithDefaults

`func NewFilesSharingShareeUserWithDefaults() *FilesSharingShareeUser`

NewFilesSharingShareeUserWithDefaults instantiates a new FilesSharingShareeUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *FilesSharingShareeUser) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *FilesSharingShareeUser) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *FilesSharingShareeUser) SetCount(v int64)`

SetCount sets Count field to given value.


### SetCountNil

`func (o *FilesSharingShareeUser) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *FilesSharingShareeUser) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetLabel

`func (o *FilesSharingShareeUser) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FilesSharingShareeUser) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FilesSharingShareeUser) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetSubline

`func (o *FilesSharingShareeUser) GetSubline() string`

GetSubline returns the Subline field if non-nil, zero value otherwise.

### GetSublineOk

`func (o *FilesSharingShareeUser) GetSublineOk() (*string, bool)`

GetSublineOk returns a tuple with the Subline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubline

`func (o *FilesSharingShareeUser) SetSubline(v string)`

SetSubline sets Subline field to given value.


### GetIcon

`func (o *FilesSharingShareeUser) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *FilesSharingShareeUser) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *FilesSharingShareeUser) SetIcon(v string)`

SetIcon sets Icon field to given value.


### GetShareWithDisplayNameUnique

`func (o *FilesSharingShareeUser) GetShareWithDisplayNameUnique() string`

GetShareWithDisplayNameUnique returns the ShareWithDisplayNameUnique field if non-nil, zero value otherwise.

### GetShareWithDisplayNameUniqueOk

`func (o *FilesSharingShareeUser) GetShareWithDisplayNameUniqueOk() (*string, bool)`

GetShareWithDisplayNameUniqueOk returns a tuple with the ShareWithDisplayNameUnique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareWithDisplayNameUnique

`func (o *FilesSharingShareeUser) SetShareWithDisplayNameUnique(v string)`

SetShareWithDisplayNameUnique sets ShareWithDisplayNameUnique field to given value.


### GetStatus

`func (o *FilesSharingShareeUser) GetStatus() FilesSharingShareeUserAllOfStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FilesSharingShareeUser) GetStatusOk() (*FilesSharingShareeUserAllOfStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FilesSharingShareeUser) SetStatus(v FilesSharingShareeUserAllOfStatus)`

SetStatus sets Status field to given value.


### GetValue

`func (o *FilesSharingShareeUser) GetValue() FilesSharingShareeValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FilesSharingShareeUser) GetValueOk() (*FilesSharingShareeValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FilesSharingShareeUser) SetValue(v FilesSharingShareeValue)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


