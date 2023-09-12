# FilesTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TemplateType** | **string** |  | 
**TemplateId** | **string** |  | 
**Basename** | **string** |  | 
**Etag** | **string** |  | 
**Fileid** | **int64** |  | 
**Filename** | **string** |  | 
**Lastmod** | **int64** |  | 
**Mime** | **string** |  | 
**Size** | **int64** |  | 
**Type** | **string** |  | 
**HasPreview** | **bool** |  | 
**PreviewUrl** | **NullableString** |  | 

## Methods

### NewFilesTemplate

`func NewFilesTemplate(templateType string, templateId string, basename string, etag string, fileid int64, filename string, lastmod int64, mime string, size int64, type_ string, hasPreview bool, previewUrl NullableString, ) *FilesTemplate`

NewFilesTemplate instantiates a new FilesTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilesTemplateWithDefaults

`func NewFilesTemplateWithDefaults() *FilesTemplate`

NewFilesTemplateWithDefaults instantiates a new FilesTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplateType

`func (o *FilesTemplate) GetTemplateType() string`

GetTemplateType returns the TemplateType field if non-nil, zero value otherwise.

### GetTemplateTypeOk

`func (o *FilesTemplate) GetTemplateTypeOk() (*string, bool)`

GetTemplateTypeOk returns a tuple with the TemplateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateType

`func (o *FilesTemplate) SetTemplateType(v string)`

SetTemplateType sets TemplateType field to given value.


### GetTemplateId

`func (o *FilesTemplate) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *FilesTemplate) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *FilesTemplate) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.


### GetBasename

`func (o *FilesTemplate) GetBasename() string`

GetBasename returns the Basename field if non-nil, zero value otherwise.

### GetBasenameOk

`func (o *FilesTemplate) GetBasenameOk() (*string, bool)`

GetBasenameOk returns a tuple with the Basename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasename

`func (o *FilesTemplate) SetBasename(v string)`

SetBasename sets Basename field to given value.


### GetEtag

`func (o *FilesTemplate) GetEtag() string`

GetEtag returns the Etag field if non-nil, zero value otherwise.

### GetEtagOk

`func (o *FilesTemplate) GetEtagOk() (*string, bool)`

GetEtagOk returns a tuple with the Etag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtag

`func (o *FilesTemplate) SetEtag(v string)`

SetEtag sets Etag field to given value.


### GetFileid

`func (o *FilesTemplate) GetFileid() int64`

GetFileid returns the Fileid field if non-nil, zero value otherwise.

### GetFileidOk

`func (o *FilesTemplate) GetFileidOk() (*int64, bool)`

GetFileidOk returns a tuple with the Fileid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileid

`func (o *FilesTemplate) SetFileid(v int64)`

SetFileid sets Fileid field to given value.


### GetFilename

`func (o *FilesTemplate) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *FilesTemplate) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *FilesTemplate) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetLastmod

`func (o *FilesTemplate) GetLastmod() int64`

GetLastmod returns the Lastmod field if non-nil, zero value otherwise.

### GetLastmodOk

`func (o *FilesTemplate) GetLastmodOk() (*int64, bool)`

GetLastmodOk returns a tuple with the Lastmod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastmod

`func (o *FilesTemplate) SetLastmod(v int64)`

SetLastmod sets Lastmod field to given value.


### GetMime

`func (o *FilesTemplate) GetMime() string`

GetMime returns the Mime field if non-nil, zero value otherwise.

### GetMimeOk

`func (o *FilesTemplate) GetMimeOk() (*string, bool)`

GetMimeOk returns a tuple with the Mime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMime

`func (o *FilesTemplate) SetMime(v string)`

SetMime sets Mime field to given value.


### GetSize

`func (o *FilesTemplate) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *FilesTemplate) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *FilesTemplate) SetSize(v int64)`

SetSize sets Size field to given value.


### GetType

`func (o *FilesTemplate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FilesTemplate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FilesTemplate) SetType(v string)`

SetType sets Type field to given value.


### GetHasPreview

`func (o *FilesTemplate) GetHasPreview() bool`

GetHasPreview returns the HasPreview field if non-nil, zero value otherwise.

### GetHasPreviewOk

`func (o *FilesTemplate) GetHasPreviewOk() (*bool, bool)`

GetHasPreviewOk returns a tuple with the HasPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPreview

`func (o *FilesTemplate) SetHasPreview(v bool)`

SetHasPreview sets HasPreview field to given value.


### GetPreviewUrl

`func (o *FilesTemplate) GetPreviewUrl() string`

GetPreviewUrl returns the PreviewUrl field if non-nil, zero value otherwise.

### GetPreviewUrlOk

`func (o *FilesTemplate) GetPreviewUrlOk() (*string, bool)`

GetPreviewUrlOk returns a tuple with the PreviewUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewUrl

`func (o *FilesTemplate) SetPreviewUrl(v string)`

SetPreviewUrl sets PreviewUrl field to given value.


### SetPreviewUrlNil

`func (o *FilesTemplate) SetPreviewUrlNil(b bool)`

 SetPreviewUrlNil sets the value for PreviewUrl to be an explicit nil

### UnsetPreviewUrl
`func (o *FilesTemplate) UnsetPreviewUrl()`

UnsetPreviewUrl ensures that no value is present for PreviewUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


