# CoreStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Installed** | **bool** |  | 
**Maintenance** | **bool** |  | 
**NeedsDbUpgrade** | **bool** |  | 
**Version** | **string** |  | 
**Versionstring** | **string** |  | 
**Edition** | **string** |  | 
**Productname** | **string** |  | 
**ExtendedSupport** | **bool** |  | 

## Methods

### NewCoreStatus

`func NewCoreStatus(installed bool, maintenance bool, needsDbUpgrade bool, version string, versionstring string, edition string, productname string, extendedSupport bool, ) *CoreStatus`

NewCoreStatus instantiates a new CoreStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoreStatusWithDefaults

`func NewCoreStatusWithDefaults() *CoreStatus`

NewCoreStatusWithDefaults instantiates a new CoreStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstalled

`func (o *CoreStatus) GetInstalled() bool`

GetInstalled returns the Installed field if non-nil, zero value otherwise.

### GetInstalledOk

`func (o *CoreStatus) GetInstalledOk() (*bool, bool)`

GetInstalledOk returns a tuple with the Installed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalled

`func (o *CoreStatus) SetInstalled(v bool)`

SetInstalled sets Installed field to given value.


### GetMaintenance

`func (o *CoreStatus) GetMaintenance() bool`

GetMaintenance returns the Maintenance field if non-nil, zero value otherwise.

### GetMaintenanceOk

`func (o *CoreStatus) GetMaintenanceOk() (*bool, bool)`

GetMaintenanceOk returns a tuple with the Maintenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenance

`func (o *CoreStatus) SetMaintenance(v bool)`

SetMaintenance sets Maintenance field to given value.


### GetNeedsDbUpgrade

`func (o *CoreStatus) GetNeedsDbUpgrade() bool`

GetNeedsDbUpgrade returns the NeedsDbUpgrade field if non-nil, zero value otherwise.

### GetNeedsDbUpgradeOk

`func (o *CoreStatus) GetNeedsDbUpgradeOk() (*bool, bool)`

GetNeedsDbUpgradeOk returns a tuple with the NeedsDbUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedsDbUpgrade

`func (o *CoreStatus) SetNeedsDbUpgrade(v bool)`

SetNeedsDbUpgrade sets NeedsDbUpgrade field to given value.


### GetVersion

`func (o *CoreStatus) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CoreStatus) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CoreStatus) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetVersionstring

`func (o *CoreStatus) GetVersionstring() string`

GetVersionstring returns the Versionstring field if non-nil, zero value otherwise.

### GetVersionstringOk

`func (o *CoreStatus) GetVersionstringOk() (*string, bool)`

GetVersionstringOk returns a tuple with the Versionstring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionstring

`func (o *CoreStatus) SetVersionstring(v string)`

SetVersionstring sets Versionstring field to given value.


### GetEdition

`func (o *CoreStatus) GetEdition() string`

GetEdition returns the Edition field if non-nil, zero value otherwise.

### GetEditionOk

`func (o *CoreStatus) GetEditionOk() (*string, bool)`

GetEditionOk returns a tuple with the Edition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdition

`func (o *CoreStatus) SetEdition(v string)`

SetEdition sets Edition field to given value.


### GetProductname

`func (o *CoreStatus) GetProductname() string`

GetProductname returns the Productname field if non-nil, zero value otherwise.

### GetProductnameOk

`func (o *CoreStatus) GetProductnameOk() (*string, bool)`

GetProductnameOk returns a tuple with the Productname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductname

`func (o *CoreStatus) SetProductname(v string)`

SetProductname sets Productname field to given value.


### GetExtendedSupport

`func (o *CoreStatus) GetExtendedSupport() bool`

GetExtendedSupport returns the ExtendedSupport field if non-nil, zero value otherwise.

### GetExtendedSupportOk

`func (o *CoreStatus) GetExtendedSupportOk() (*bool, bool)`

GetExtendedSupportOk returns a tuple with the ExtendedSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedSupport

`func (o *CoreStatus) SetExtendedSupport(v bool)`

SetExtendedSupport sets ExtendedSupport field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


