# CoreUnifiedSearchResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**IsPaginated** | **bool** |  | 
**Entries** | [**[]CoreUnifiedSearchResultEntry**](CoreUnifiedSearchResultEntry.md) |  | 
**Cursor** | [**NullableCoreUnifiedSearchResultCursor**](CoreUnifiedSearchResultCursor.md) |  | 

## Methods

### NewCoreUnifiedSearchResult

`func NewCoreUnifiedSearchResult(name string, isPaginated bool, entries []CoreUnifiedSearchResultEntry, cursor NullableCoreUnifiedSearchResultCursor, ) *CoreUnifiedSearchResult`

NewCoreUnifiedSearchResult instantiates a new CoreUnifiedSearchResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoreUnifiedSearchResultWithDefaults

`func NewCoreUnifiedSearchResultWithDefaults() *CoreUnifiedSearchResult`

NewCoreUnifiedSearchResultWithDefaults instantiates a new CoreUnifiedSearchResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CoreUnifiedSearchResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CoreUnifiedSearchResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CoreUnifiedSearchResult) SetName(v string)`

SetName sets Name field to given value.


### GetIsPaginated

`func (o *CoreUnifiedSearchResult) GetIsPaginated() bool`

GetIsPaginated returns the IsPaginated field if non-nil, zero value otherwise.

### GetIsPaginatedOk

`func (o *CoreUnifiedSearchResult) GetIsPaginatedOk() (*bool, bool)`

GetIsPaginatedOk returns a tuple with the IsPaginated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPaginated

`func (o *CoreUnifiedSearchResult) SetIsPaginated(v bool)`

SetIsPaginated sets IsPaginated field to given value.


### GetEntries

`func (o *CoreUnifiedSearchResult) GetEntries() []CoreUnifiedSearchResultEntry`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *CoreUnifiedSearchResult) GetEntriesOk() (*[]CoreUnifiedSearchResultEntry, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *CoreUnifiedSearchResult) SetEntries(v []CoreUnifiedSearchResultEntry)`

SetEntries sets Entries field to given value.


### GetCursor

`func (o *CoreUnifiedSearchResult) GetCursor() CoreUnifiedSearchResultCursor`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *CoreUnifiedSearchResult) GetCursorOk() (*CoreUnifiedSearchResultCursor, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *CoreUnifiedSearchResult) SetCursor(v CoreUnifiedSearchResultCursor)`

SetCursor sets Cursor field to given value.


### SetCursorNil

`func (o *CoreUnifiedSearchResult) SetCursorNil(b bool)`

 SetCursorNil sets the value for Cursor to be an explicit nil

### UnsetCursor
`func (o *CoreUnifiedSearchResult) UnsetCursor()`

UnsetCursor ensures that no value is present for Cursor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


