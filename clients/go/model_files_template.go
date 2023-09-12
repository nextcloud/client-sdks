/*
nextcloud

Nextcloud APIs

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client_sdk

import (
	"encoding/json"
)

// checks if the FilesTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FilesTemplate{}

// FilesTemplate struct for FilesTemplate
type FilesTemplate struct {
	TemplateType string `json:"templateType"`
	TemplateId string `json:"templateId"`
	Basename string `json:"basename"`
	Etag string `json:"etag"`
	Fileid int64 `json:"fileid"`
	Filename string `json:"filename"`
	Lastmod int64 `json:"lastmod"`
	Mime string `json:"mime"`
	Size int64 `json:"size"`
	Type string `json:"type"`
	HasPreview bool `json:"hasPreview"`
	PreviewUrl NullableString `json:"previewUrl"`
	AdditionalProperties map[string]interface{}
}

type _FilesTemplate FilesTemplate

// NewFilesTemplate instantiates a new FilesTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilesTemplate(templateType string, templateId string, basename string, etag string, fileid int64, filename string, lastmod int64, mime string, size int64, type_ string, hasPreview bool, previewUrl NullableString) *FilesTemplate {
	this := FilesTemplate{}
	this.TemplateType = templateType
	this.TemplateId = templateId
	this.Basename = basename
	this.Etag = etag
	this.Fileid = fileid
	this.Filename = filename
	this.Lastmod = lastmod
	this.Mime = mime
	this.Size = size
	this.Type = type_
	this.HasPreview = hasPreview
	this.PreviewUrl = previewUrl
	return &this
}

// NewFilesTemplateWithDefaults instantiates a new FilesTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilesTemplateWithDefaults() *FilesTemplate {
	this := FilesTemplate{}
	return &this
}

// GetTemplateType returns the TemplateType field value
func (o *FilesTemplate) GetTemplateType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TemplateType
}

// GetTemplateTypeOk returns a tuple with the TemplateType field value
// and a boolean to check if the value has been set.
func (o *FilesTemplate) GetTemplateTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateType, true
}

// SetTemplateType sets field value
func (o *FilesTemplate) SetTemplateType(v string) {
	o.TemplateType = v
}

// GetTemplateId returns the TemplateId field value
func (o *FilesTemplate) GetTemplateId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value
// and a boolean to check if the value has been set.
func (o *FilesTemplate) GetTemplateIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateId, true
}

// SetTemplateId sets field value
func (o *FilesTemplate) SetTemplateId(v string) {
	o.TemplateId = v
}

// GetBasename returns the Basename field value
func (o *FilesTemplate) GetBasename() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Basename
}

// GetBasenameOk returns a tuple with the Basename field value
// and a boolean to check if the value has been set.
func (o *FilesTemplate) GetBasenameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Basename, true
}

// SetBasename sets field value
func (o *FilesTemplate) SetBasename(v string) {
	o.Basename = v
}

// GetEtag returns the Etag field value
func (o *FilesTemplate) GetEtag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Etag
}

// GetEtagOk returns a tuple with the Etag field value
// and a boolean to check if the value has been set.
func (o *FilesTemplate) GetEtagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Etag, true
}

// SetEtag sets field value
func (o *FilesTemplate) SetEtag(v string) {
	o.Etag = v
}

// GetFileid returns the Fileid field value
func (o *FilesTemplate) GetFileid() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Fileid
}

// GetFileidOk returns a tuple with the Fileid field value
// and a boolean to check if the value has been set.
func (o *FilesTemplate) GetFileidOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fileid, true
}

// SetFileid sets field value
func (o *FilesTemplate) SetFileid(v int64) {
	o.Fileid = v
}

// GetFilename returns the Filename field value
func (o *FilesTemplate) GetFilename() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value
// and a boolean to check if the value has been set.
func (o *FilesTemplate) GetFilenameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filename, true
}

// SetFilename sets field value
func (o *FilesTemplate) SetFilename(v string) {
	o.Filename = v
}

// GetLastmod returns the Lastmod field value
func (o *FilesTemplate) GetLastmod() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Lastmod
}

// GetLastmodOk returns a tuple with the Lastmod field value
// and a boolean to check if the value has been set.
func (o *FilesTemplate) GetLastmodOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Lastmod, true
}

// SetLastmod sets field value
func (o *FilesTemplate) SetLastmod(v int64) {
	o.Lastmod = v
}

// GetMime returns the Mime field value
func (o *FilesTemplate) GetMime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mime
}

// GetMimeOk returns a tuple with the Mime field value
// and a boolean to check if the value has been set.
func (o *FilesTemplate) GetMimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mime, true
}

// SetMime sets field value
func (o *FilesTemplate) SetMime(v string) {
	o.Mime = v
}

// GetSize returns the Size field value
func (o *FilesTemplate) GetSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *FilesTemplate) GetSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *FilesTemplate) SetSize(v int64) {
	o.Size = v
}

// GetType returns the Type field value
func (o *FilesTemplate) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FilesTemplate) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FilesTemplate) SetType(v string) {
	o.Type = v
}

// GetHasPreview returns the HasPreview field value
func (o *FilesTemplate) GetHasPreview() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasPreview
}

// GetHasPreviewOk returns a tuple with the HasPreview field value
// and a boolean to check if the value has been set.
func (o *FilesTemplate) GetHasPreviewOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasPreview, true
}

// SetHasPreview sets field value
func (o *FilesTemplate) SetHasPreview(v bool) {
	o.HasPreview = v
}

// GetPreviewUrl returns the PreviewUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FilesTemplate) GetPreviewUrl() string {
	if o == nil || o.PreviewUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.PreviewUrl.Get()
}

// GetPreviewUrlOk returns a tuple with the PreviewUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesTemplate) GetPreviewUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PreviewUrl.Get(), o.PreviewUrl.IsSet()
}

// SetPreviewUrl sets field value
func (o *FilesTemplate) SetPreviewUrl(v string) {
	o.PreviewUrl.Set(&v)
}

func (o FilesTemplate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FilesTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["templateType"] = o.TemplateType
	toSerialize["templateId"] = o.TemplateId
	toSerialize["basename"] = o.Basename
	toSerialize["etag"] = o.Etag
	toSerialize["fileid"] = o.Fileid
	toSerialize["filename"] = o.Filename
	toSerialize["lastmod"] = o.Lastmod
	toSerialize["mime"] = o.Mime
	toSerialize["size"] = o.Size
	toSerialize["type"] = o.Type
	toSerialize["hasPreview"] = o.HasPreview
	toSerialize["previewUrl"] = o.PreviewUrl.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FilesTemplate) UnmarshalJSON(bytes []byte) (err error) {
	varFilesTemplate := _FilesTemplate{}

	if err = json.Unmarshal(bytes, &varFilesTemplate); err == nil {
		*o = FilesTemplate(varFilesTemplate)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "templateType")
		delete(additionalProperties, "templateId")
		delete(additionalProperties, "basename")
		delete(additionalProperties, "etag")
		delete(additionalProperties, "fileid")
		delete(additionalProperties, "filename")
		delete(additionalProperties, "lastmod")
		delete(additionalProperties, "mime")
		delete(additionalProperties, "size")
		delete(additionalProperties, "type")
		delete(additionalProperties, "hasPreview")
		delete(additionalProperties, "previewUrl")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFilesTemplate struct {
	value *FilesTemplate
	isSet bool
}

func (v NullableFilesTemplate) Get() *FilesTemplate {
	return v.value
}

func (v *NullableFilesTemplate) Set(val *FilesTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableFilesTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableFilesTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilesTemplate(val *FilesTemplate) *NullableFilesTemplate {
	return &NullableFilesTemplate{value: val, isSet: true}
}

func (v NullableFilesTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilesTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

