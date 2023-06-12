# CoreReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RichObjectType** | **string** |  | 
**RichObject** | **map[string]map[string]interface{}** |  | 
**OpenGraphObject** | [**CoreReferenceOpenGraphObject**](CoreReferenceOpenGraphObject.md) |  | 
**Accessible** | **bool** |  | 

## Methods

### NewCoreReference

`func NewCoreReference(richObjectType string, richObject map[string]map[string]interface{}, openGraphObject CoreReferenceOpenGraphObject, accessible bool, ) *CoreReference`

NewCoreReference instantiates a new CoreReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoreReferenceWithDefaults

`func NewCoreReferenceWithDefaults() *CoreReference`

NewCoreReferenceWithDefaults instantiates a new CoreReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRichObjectType

`func (o *CoreReference) GetRichObjectType() string`

GetRichObjectType returns the RichObjectType field if non-nil, zero value otherwise.

### GetRichObjectTypeOk

`func (o *CoreReference) GetRichObjectTypeOk() (*string, bool)`

GetRichObjectTypeOk returns a tuple with the RichObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRichObjectType

`func (o *CoreReference) SetRichObjectType(v string)`

SetRichObjectType sets RichObjectType field to given value.


### GetRichObject

`func (o *CoreReference) GetRichObject() map[string]map[string]interface{}`

GetRichObject returns the RichObject field if non-nil, zero value otherwise.

### GetRichObjectOk

`func (o *CoreReference) GetRichObjectOk() (*map[string]map[string]interface{}, bool)`

GetRichObjectOk returns a tuple with the RichObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRichObject

`func (o *CoreReference) SetRichObject(v map[string]map[string]interface{})`

SetRichObject sets RichObject field to given value.


### GetOpenGraphObject

`func (o *CoreReference) GetOpenGraphObject() CoreReferenceOpenGraphObject`

GetOpenGraphObject returns the OpenGraphObject field if non-nil, zero value otherwise.

### GetOpenGraphObjectOk

`func (o *CoreReference) GetOpenGraphObjectOk() (*CoreReferenceOpenGraphObject, bool)`

GetOpenGraphObjectOk returns a tuple with the OpenGraphObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenGraphObject

`func (o *CoreReference) SetOpenGraphObject(v CoreReferenceOpenGraphObject)`

SetOpenGraphObject sets OpenGraphObject field to given value.


### GetAccessible

`func (o *CoreReference) GetAccessible() bool`

GetAccessible returns the Accessible field if non-nil, zero value otherwise.

### GetAccessibleOk

`func (o *CoreReference) GetAccessibleOk() (*bool, bool)`

GetAccessibleOk returns a tuple with the Accessible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessible

`func (o *CoreReference) SetAccessible(v bool)`

SetAccessible sets Accessible field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


