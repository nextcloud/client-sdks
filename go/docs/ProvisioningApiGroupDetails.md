# ProvisioningApiGroupDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Displayname** | **string** |  | 
**Usercount** | [**ProvisioningApiGroupDetailsUsercount**](ProvisioningApiGroupDetailsUsercount.md) |  | 
**Disabled** | [**ProvisioningApiGroupDetailsUsercount**](ProvisioningApiGroupDetailsUsercount.md) |  | 
**CanAdd** | **bool** |  | 
**CanRemove** | **bool** |  | 

## Methods

### NewProvisioningApiGroupDetails

`func NewProvisioningApiGroupDetails(id string, displayname string, usercount ProvisioningApiGroupDetailsUsercount, disabled ProvisioningApiGroupDetailsUsercount, canAdd bool, canRemove bool, ) *ProvisioningApiGroupDetails`

NewProvisioningApiGroupDetails instantiates a new ProvisioningApiGroupDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisioningApiGroupDetailsWithDefaults

`func NewProvisioningApiGroupDetailsWithDefaults() *ProvisioningApiGroupDetails`

NewProvisioningApiGroupDetailsWithDefaults instantiates a new ProvisioningApiGroupDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProvisioningApiGroupDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProvisioningApiGroupDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProvisioningApiGroupDetails) SetId(v string)`

SetId sets Id field to given value.


### GetDisplayname

`func (o *ProvisioningApiGroupDetails) GetDisplayname() string`

GetDisplayname returns the Displayname field if non-nil, zero value otherwise.

### GetDisplaynameOk

`func (o *ProvisioningApiGroupDetails) GetDisplaynameOk() (*string, bool)`

GetDisplaynameOk returns a tuple with the Displayname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayname

`func (o *ProvisioningApiGroupDetails) SetDisplayname(v string)`

SetDisplayname sets Displayname field to given value.


### GetUsercount

`func (o *ProvisioningApiGroupDetails) GetUsercount() ProvisioningApiGroupDetailsUsercount`

GetUsercount returns the Usercount field if non-nil, zero value otherwise.

### GetUsercountOk

`func (o *ProvisioningApiGroupDetails) GetUsercountOk() (*ProvisioningApiGroupDetailsUsercount, bool)`

GetUsercountOk returns a tuple with the Usercount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsercount

`func (o *ProvisioningApiGroupDetails) SetUsercount(v ProvisioningApiGroupDetailsUsercount)`

SetUsercount sets Usercount field to given value.


### GetDisabled

`func (o *ProvisioningApiGroupDetails) GetDisabled() ProvisioningApiGroupDetailsUsercount`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *ProvisioningApiGroupDetails) GetDisabledOk() (*ProvisioningApiGroupDetailsUsercount, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *ProvisioningApiGroupDetails) SetDisabled(v ProvisioningApiGroupDetailsUsercount)`

SetDisabled sets Disabled field to given value.


### GetCanAdd

`func (o *ProvisioningApiGroupDetails) GetCanAdd() bool`

GetCanAdd returns the CanAdd field if non-nil, zero value otherwise.

### GetCanAddOk

`func (o *ProvisioningApiGroupDetails) GetCanAddOk() (*bool, bool)`

GetCanAddOk returns a tuple with the CanAdd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanAdd

`func (o *ProvisioningApiGroupDetails) SetCanAdd(v bool)`

SetCanAdd sets CanAdd field to given value.


### GetCanRemove

`func (o *ProvisioningApiGroupDetails) GetCanRemove() bool`

GetCanRemove returns the CanRemove field if non-nil, zero value otherwise.

### GetCanRemoveOk

`func (o *ProvisioningApiGroupDetails) GetCanRemoveOk() (*bool, bool)`

GetCanRemoveOk returns a tuple with the CanRemove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanRemove

`func (o *ProvisioningApiGroupDetails) SetCanRemove(v bool)`

SetCanRemove sets CanRemove field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


