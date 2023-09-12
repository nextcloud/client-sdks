# CoreCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Name** | **string** |  | 
**Resources** | [**[]CoreResource**](CoreResource.md) |  | 

## Methods

### NewCoreCollection

`func NewCoreCollection(id int64, name string, resources []CoreResource, ) *CoreCollection`

NewCoreCollection instantiates a new CoreCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoreCollectionWithDefaults

`func NewCoreCollectionWithDefaults() *CoreCollection`

NewCoreCollectionWithDefaults instantiates a new CoreCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CoreCollection) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CoreCollection) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CoreCollection) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *CoreCollection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CoreCollection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CoreCollection) SetName(v string)`

SetName sets Name field to given value.


### GetResources

`func (o *CoreCollection) GetResources() []CoreResource`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *CoreCollection) GetResourcesOk() (*[]CoreResource, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *CoreCollection) SetResources(v []CoreResource)`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


