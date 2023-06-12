# DashboardWidget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Title** | **string** |  | 
**Order** | **int64** |  | 
**IconClass** | **string** |  | 
**IconUrl** | **string** |  | 
**WidgetUrl** | **NullableString** |  | 
**ItemIconsRound** | **bool** |  | 
**Buttons** | Pointer to [**[]DashboardWidgetButtonsInner**](DashboardWidgetButtonsInner.md) |  | [optional] 

## Methods

### NewDashboardWidget

`func NewDashboardWidget(id string, title string, order int64, iconClass string, iconUrl string, widgetUrl NullableString, itemIconsRound bool, ) *DashboardWidget`

NewDashboardWidget instantiates a new DashboardWidget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardWidgetWithDefaults

`func NewDashboardWidgetWithDefaults() *DashboardWidget`

NewDashboardWidgetWithDefaults instantiates a new DashboardWidget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DashboardWidget) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DashboardWidget) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DashboardWidget) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *DashboardWidget) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DashboardWidget) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DashboardWidget) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetOrder

`func (o *DashboardWidget) GetOrder() int64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *DashboardWidget) GetOrderOk() (*int64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *DashboardWidget) SetOrder(v int64)`

SetOrder sets Order field to given value.


### GetIconClass

`func (o *DashboardWidget) GetIconClass() string`

GetIconClass returns the IconClass field if non-nil, zero value otherwise.

### GetIconClassOk

`func (o *DashboardWidget) GetIconClassOk() (*string, bool)`

GetIconClassOk returns a tuple with the IconClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconClass

`func (o *DashboardWidget) SetIconClass(v string)`

SetIconClass sets IconClass field to given value.


### GetIconUrl

`func (o *DashboardWidget) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *DashboardWidget) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *DashboardWidget) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.


### GetWidgetUrl

`func (o *DashboardWidget) GetWidgetUrl() string`

GetWidgetUrl returns the WidgetUrl field if non-nil, zero value otherwise.

### GetWidgetUrlOk

`func (o *DashboardWidget) GetWidgetUrlOk() (*string, bool)`

GetWidgetUrlOk returns a tuple with the WidgetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidgetUrl

`func (o *DashboardWidget) SetWidgetUrl(v string)`

SetWidgetUrl sets WidgetUrl field to given value.


### SetWidgetUrlNil

`func (o *DashboardWidget) SetWidgetUrlNil(b bool)`

 SetWidgetUrlNil sets the value for WidgetUrl to be an explicit nil

### UnsetWidgetUrl
`func (o *DashboardWidget) UnsetWidgetUrl()`

UnsetWidgetUrl ensures that no value is present for WidgetUrl, not even an explicit nil
### GetItemIconsRound

`func (o *DashboardWidget) GetItemIconsRound() bool`

GetItemIconsRound returns the ItemIconsRound field if non-nil, zero value otherwise.

### GetItemIconsRoundOk

`func (o *DashboardWidget) GetItemIconsRoundOk() (*bool, bool)`

GetItemIconsRoundOk returns a tuple with the ItemIconsRound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemIconsRound

`func (o *DashboardWidget) SetItemIconsRound(v bool)`

SetItemIconsRound sets ItemIconsRound field to given value.


### GetButtons

`func (o *DashboardWidget) GetButtons() []DashboardWidgetButtonsInner`

GetButtons returns the Buttons field if non-nil, zero value otherwise.

### GetButtonsOk

`func (o *DashboardWidget) GetButtonsOk() (*[]DashboardWidgetButtonsInner, bool)`

GetButtonsOk returns a tuple with the Buttons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtons

`func (o *DashboardWidget) SetButtons(v []DashboardWidgetButtonsInner)`

SetButtons sets Buttons field to given value.

### HasButtons

`func (o *DashboardWidget) HasButtons() bool`

HasButtons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


