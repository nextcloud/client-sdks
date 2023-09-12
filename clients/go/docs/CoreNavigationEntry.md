# CoreNavigationEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Order** | [**CoreNavigationEntryOrder**](CoreNavigationEntryOrder.md) |  | 
**Href** | **string** |  | 
**Icon** | **string** |  | 
**Type** | **string** |  | 
**Name** | **string** |  | 
**Active** | **bool** |  | 
**Classes** | **string** |  | 
**Unread** | **int64** |  | 

## Methods

### NewCoreNavigationEntry

`func NewCoreNavigationEntry(id string, order CoreNavigationEntryOrder, href string, icon string, type_ string, name string, active bool, classes string, unread int64, ) *CoreNavigationEntry`

NewCoreNavigationEntry instantiates a new CoreNavigationEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoreNavigationEntryWithDefaults

`func NewCoreNavigationEntryWithDefaults() *CoreNavigationEntry`

NewCoreNavigationEntryWithDefaults instantiates a new CoreNavigationEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CoreNavigationEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CoreNavigationEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CoreNavigationEntry) SetId(v string)`

SetId sets Id field to given value.


### GetOrder

`func (o *CoreNavigationEntry) GetOrder() CoreNavigationEntryOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *CoreNavigationEntry) GetOrderOk() (*CoreNavigationEntryOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *CoreNavigationEntry) SetOrder(v CoreNavigationEntryOrder)`

SetOrder sets Order field to given value.


### GetHref

`func (o *CoreNavigationEntry) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *CoreNavigationEntry) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *CoreNavigationEntry) SetHref(v string)`

SetHref sets Href field to given value.


### GetIcon

`func (o *CoreNavigationEntry) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *CoreNavigationEntry) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *CoreNavigationEntry) SetIcon(v string)`

SetIcon sets Icon field to given value.


### GetType

`func (o *CoreNavigationEntry) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CoreNavigationEntry) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CoreNavigationEntry) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *CoreNavigationEntry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CoreNavigationEntry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CoreNavigationEntry) SetName(v string)`

SetName sets Name field to given value.


### GetActive

`func (o *CoreNavigationEntry) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CoreNavigationEntry) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CoreNavigationEntry) SetActive(v bool)`

SetActive sets Active field to given value.


### GetClasses

`func (o *CoreNavigationEntry) GetClasses() string`

GetClasses returns the Classes field if non-nil, zero value otherwise.

### GetClassesOk

`func (o *CoreNavigationEntry) GetClassesOk() (*string, bool)`

GetClassesOk returns a tuple with the Classes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClasses

`func (o *CoreNavigationEntry) SetClasses(v string)`

SetClasses sets Classes field to given value.


### GetUnread

`func (o *CoreNavigationEntry) GetUnread() int64`

GetUnread returns the Unread field if non-nil, zero value otherwise.

### GetUnreadOk

`func (o *CoreNavigationEntry) GetUnreadOk() (*int64, bool)`

GetUnreadOk returns a tuple with the Unread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnread

`func (o *CoreNavigationEntry) SetUnread(v int64)`

SetUnread sets Unread field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


