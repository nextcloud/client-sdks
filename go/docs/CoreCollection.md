# CoreCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Resources** | **[]map[string]map[string]interface{}** |  | 

## Methods

### NewCoreCollection

`func NewCoreCollection(id string, name string, resources []map[string]map[string]interface{}, ) *CoreCollection`

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

`func (o *CoreCollection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CoreCollection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CoreCollection) SetId(v string)`

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

`func (o *CoreCollection) GetResources() []map[string]map[string]interface{}`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *CoreCollection) GetResourcesOk() (*[]map[string]map[string]interface{}, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *CoreCollection) SetResources(v []map[string]map[string]interface{})`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


