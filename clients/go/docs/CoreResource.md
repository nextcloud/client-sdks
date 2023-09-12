# CoreResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RichObjectType** | **string** |  | 
**RichObject** | **map[string]map[string]interface{}** |  | 
**OpenGraphObject** | [**CoreOpenGraphObject**](CoreOpenGraphObject.md) |  | 
**Accessible** | **bool** |  | 

## Methods

### NewCoreResource

`func NewCoreResource(richObjectType string, richObject map[string]map[string]interface{}, openGraphObject CoreOpenGraphObject, accessible bool, ) *CoreResource`

NewCoreResource instantiates a new CoreResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoreResourceWithDefaults

`func NewCoreResourceWithDefaults() *CoreResource`

NewCoreResourceWithDefaults instantiates a new CoreResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRichObjectType

`func (o *CoreResource) GetRichObjectType() string`

GetRichObjectType returns the RichObjectType field if non-nil, zero value otherwise.

### GetRichObjectTypeOk

`func (o *CoreResource) GetRichObjectTypeOk() (*string, bool)`

GetRichObjectTypeOk returns a tuple with the RichObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRichObjectType

`func (o *CoreResource) SetRichObjectType(v string)`

SetRichObjectType sets RichObjectType field to given value.


### GetRichObject

`func (o *CoreResource) GetRichObject() map[string]map[string]interface{}`

GetRichObject returns the RichObject field if non-nil, zero value otherwise.

### GetRichObjectOk

`func (o *CoreResource) GetRichObjectOk() (*map[string]map[string]interface{}, bool)`

GetRichObjectOk returns a tuple with the RichObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRichObject

`func (o *CoreResource) SetRichObject(v map[string]map[string]interface{})`

SetRichObject sets RichObject field to given value.


### GetOpenGraphObject

`func (o *CoreResource) GetOpenGraphObject() CoreOpenGraphObject`

GetOpenGraphObject returns the OpenGraphObject field if non-nil, zero value otherwise.

### GetOpenGraphObjectOk

`func (o *CoreResource) GetOpenGraphObjectOk() (*CoreOpenGraphObject, bool)`

GetOpenGraphObjectOk returns a tuple with the OpenGraphObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenGraphObject

`func (o *CoreResource) SetOpenGraphObject(v CoreOpenGraphObject)`

SetOpenGraphObject sets OpenGraphObject field to given value.


### GetAccessible

`func (o *CoreResource) GetAccessible() bool`

GetAccessible returns the Accessible field if non-nil, zero value otherwise.

### GetAccessibleOk

`func (o *CoreResource) GetAccessibleOk() (*bool, bool)`

GetAccessibleOk returns a tuple with the Accessible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessible

`func (o *CoreResource) SetAccessible(v bool)`

SetAccessible sets Accessible field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


