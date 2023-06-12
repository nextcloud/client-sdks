# OCSMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Statuscode** | **int32** |  | 
**Message** | Pointer to **string** |  | [optional] 
**Totalitems** | Pointer to **string** |  | [optional] 
**Itemsperpage** | Pointer to **string** |  | [optional] 

## Methods

### NewOCSMeta

`func NewOCSMeta(status string, statuscode int32, ) *OCSMeta`

NewOCSMeta instantiates a new OCSMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOCSMetaWithDefaults

`func NewOCSMetaWithDefaults() *OCSMeta`

NewOCSMetaWithDefaults instantiates a new OCSMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *OCSMeta) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OCSMeta) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OCSMeta) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStatuscode

`func (o *OCSMeta) GetStatuscode() int32`

GetStatuscode returns the Statuscode field if non-nil, zero value otherwise.

### GetStatuscodeOk

`func (o *OCSMeta) GetStatuscodeOk() (*int32, bool)`

GetStatuscodeOk returns a tuple with the Statuscode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuscode

`func (o *OCSMeta) SetStatuscode(v int32)`

SetStatuscode sets Statuscode field to given value.


### GetMessage

`func (o *OCSMeta) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *OCSMeta) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *OCSMeta) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *OCSMeta) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetTotalitems

`func (o *OCSMeta) GetTotalitems() string`

GetTotalitems returns the Totalitems field if non-nil, zero value otherwise.

### GetTotalitemsOk

`func (o *OCSMeta) GetTotalitemsOk() (*string, bool)`

GetTotalitemsOk returns a tuple with the Totalitems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalitems

`func (o *OCSMeta) SetTotalitems(v string)`

SetTotalitems sets Totalitems field to given value.

### HasTotalitems

`func (o *OCSMeta) HasTotalitems() bool`

HasTotalitems returns a boolean if a field has been set.

### GetItemsperpage

`func (o *OCSMeta) GetItemsperpage() string`

GetItemsperpage returns the Itemsperpage field if non-nil, zero value otherwise.

### GetItemsperpageOk

`func (o *OCSMeta) GetItemsperpageOk() (*string, bool)`

GetItemsperpageOk returns a tuple with the Itemsperpage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsperpage

`func (o *OCSMeta) SetItemsperpage(v string)`

SetItemsperpage sets Itemsperpage field to given value.

### HasItemsperpage

`func (o *OCSMeta) HasItemsperpage() bool`

HasItemsperpage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


