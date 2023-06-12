# ProvisioningApiUserDetailsQuota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Free** | **NullableInt64** |  | 
**Quota** | [**ProvisioningApiUserDetailsQuotaQuota**](ProvisioningApiUserDetailsQuotaQuota.md) |  | 
**Relative** | **NullableFloat32** |  | 
**Total** | **NullableInt64** |  | 
**Used** | **int64** |  | 

## Methods

### NewProvisioningApiUserDetailsQuota

`func NewProvisioningApiUserDetailsQuota(free NullableInt64, quota ProvisioningApiUserDetailsQuotaQuota, relative NullableFloat32, total NullableInt64, used int64, ) *ProvisioningApiUserDetailsQuota`

NewProvisioningApiUserDetailsQuota instantiates a new ProvisioningApiUserDetailsQuota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisioningApiUserDetailsQuotaWithDefaults

`func NewProvisioningApiUserDetailsQuotaWithDefaults() *ProvisioningApiUserDetailsQuota`

NewProvisioningApiUserDetailsQuotaWithDefaults instantiates a new ProvisioningApiUserDetailsQuota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFree

`func (o *ProvisioningApiUserDetailsQuota) GetFree() int64`

GetFree returns the Free field if non-nil, zero value otherwise.

### GetFreeOk

`func (o *ProvisioningApiUserDetailsQuota) GetFreeOk() (*int64, bool)`

GetFreeOk returns a tuple with the Free field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree

`func (o *ProvisioningApiUserDetailsQuota) SetFree(v int64)`

SetFree sets Free field to given value.


### SetFreeNil

`func (o *ProvisioningApiUserDetailsQuota) SetFreeNil(b bool)`

 SetFreeNil sets the value for Free to be an explicit nil

### UnsetFree
`func (o *ProvisioningApiUserDetailsQuota) UnsetFree()`

UnsetFree ensures that no value is present for Free, not even an explicit nil
### GetQuota

`func (o *ProvisioningApiUserDetailsQuota) GetQuota() ProvisioningApiUserDetailsQuotaQuota`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *ProvisioningApiUserDetailsQuota) GetQuotaOk() (*ProvisioningApiUserDetailsQuotaQuota, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *ProvisioningApiUserDetailsQuota) SetQuota(v ProvisioningApiUserDetailsQuotaQuota)`

SetQuota sets Quota field to given value.


### GetRelative

`func (o *ProvisioningApiUserDetailsQuota) GetRelative() float32`

GetRelative returns the Relative field if non-nil, zero value otherwise.

### GetRelativeOk

`func (o *ProvisioningApiUserDetailsQuota) GetRelativeOk() (*float32, bool)`

GetRelativeOk returns a tuple with the Relative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelative

`func (o *ProvisioningApiUserDetailsQuota) SetRelative(v float32)`

SetRelative sets Relative field to given value.


### SetRelativeNil

`func (o *ProvisioningApiUserDetailsQuota) SetRelativeNil(b bool)`

 SetRelativeNil sets the value for Relative to be an explicit nil

### UnsetRelative
`func (o *ProvisioningApiUserDetailsQuota) UnsetRelative()`

UnsetRelative ensures that no value is present for Relative, not even an explicit nil
### GetTotal

`func (o *ProvisioningApiUserDetailsQuota) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ProvisioningApiUserDetailsQuota) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ProvisioningApiUserDetailsQuota) SetTotal(v int64)`

SetTotal sets Total field to given value.


### SetTotalNil

`func (o *ProvisioningApiUserDetailsQuota) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *ProvisioningApiUserDetailsQuota) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil
### GetUsed

`func (o *ProvisioningApiUserDetailsQuota) GetUsed() int64`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *ProvisioningApiUserDetailsQuota) GetUsedOk() (*int64, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *ProvisioningApiUserDetailsQuota) SetUsed(v int64)`

SetUsed sets Used field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


