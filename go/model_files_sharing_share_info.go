/*
nextcloud

Nextcloud APIs

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the FilesSharingShareInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FilesSharingShareInfo{}

// FilesSharingShareInfo struct for FilesSharingShareInfo
type FilesSharingShareInfo struct {
	Id int64 `json:"id"`
	ParentId int64 `json:"parentId"`
	Mtime int64 `json:"mtime"`
	Name string `json:"name"`
	Permissions int64 `json:"permissions"`
	Mimetype string `json:"mimetype"`
	Size int64 `json:"size"`
	Type string `json:"type"`
	Etag string `json:"etag"`
	Children []map[string]map[string]interface{} `json:"children"`
}

// NewFilesSharingShareInfo instantiates a new FilesSharingShareInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilesSharingShareInfo(id int64, parentId int64, mtime int64, name string, permissions int64, mimetype string, size int64, type_ string, etag string, children []map[string]map[string]interface{}) *FilesSharingShareInfo {
	this := FilesSharingShareInfo{}
	this.Id = id
	this.ParentId = parentId
	this.Mtime = mtime
	this.Name = name
	this.Permissions = permissions
	this.Mimetype = mimetype
	this.Size = size
	this.Type = type_
	this.Etag = etag
	this.Children = children
	return &this
}

// NewFilesSharingShareInfoWithDefaults instantiates a new FilesSharingShareInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilesSharingShareInfoWithDefaults() *FilesSharingShareInfo {
	this := FilesSharingShareInfo{}
	return &this
}

// GetId returns the Id field value
func (o *FilesSharingShareInfo) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareInfo) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FilesSharingShareInfo) SetId(v int64) {
	o.Id = v
}

// GetParentId returns the ParentId field value
func (o *FilesSharingShareInfo) GetParentId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareInfo) GetParentIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentId, true
}

// SetParentId sets field value
func (o *FilesSharingShareInfo) SetParentId(v int64) {
	o.ParentId = v
}

// GetMtime returns the Mtime field value
func (o *FilesSharingShareInfo) GetMtime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Mtime
}

// GetMtimeOk returns a tuple with the Mtime field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareInfo) GetMtimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mtime, true
}

// SetMtime sets field value
func (o *FilesSharingShareInfo) SetMtime(v int64) {
	o.Mtime = v
}

// GetName returns the Name field value
func (o *FilesSharingShareInfo) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareInfo) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FilesSharingShareInfo) SetName(v string) {
	o.Name = v
}

// GetPermissions returns the Permissions field value
func (o *FilesSharingShareInfo) GetPermissions() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareInfo) GetPermissionsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Permissions, true
}

// SetPermissions sets field value
func (o *FilesSharingShareInfo) SetPermissions(v int64) {
	o.Permissions = v
}

// GetMimetype returns the Mimetype field value
func (o *FilesSharingShareInfo) GetMimetype() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mimetype
}

// GetMimetypeOk returns a tuple with the Mimetype field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareInfo) GetMimetypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mimetype, true
}

// SetMimetype sets field value
func (o *FilesSharingShareInfo) SetMimetype(v string) {
	o.Mimetype = v
}

// GetSize returns the Size field value
func (o *FilesSharingShareInfo) GetSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareInfo) GetSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *FilesSharingShareInfo) SetSize(v int64) {
	o.Size = v
}

// GetType returns the Type field value
func (o *FilesSharingShareInfo) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareInfo) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FilesSharingShareInfo) SetType(v string) {
	o.Type = v
}

// GetEtag returns the Etag field value
func (o *FilesSharingShareInfo) GetEtag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Etag
}

// GetEtagOk returns a tuple with the Etag field value
// and a boolean to check if the value has been set.
func (o *FilesSharingShareInfo) GetEtagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Etag, true
}

// SetEtag sets field value
func (o *FilesSharingShareInfo) SetEtag(v string) {
	o.Etag = v
}

// GetChildren returns the Children field value
// If the value is explicit nil, the zero value for []map[string]map[string]interface{} will be returned
func (o *FilesSharingShareInfo) GetChildren() []map[string]map[string]interface{} {
	if o == nil {
		var ret []map[string]map[string]interface{}
		return ret
	}

	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FilesSharingShareInfo) GetChildrenOk() ([]map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Children) {
		return nil, false
	}
	return o.Children, true
}

// SetChildren sets field value
func (o *FilesSharingShareInfo) SetChildren(v []map[string]map[string]interface{}) {
	o.Children = v
}

func (o FilesSharingShareInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FilesSharingShareInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["parentId"] = o.ParentId
	toSerialize["mtime"] = o.Mtime
	toSerialize["name"] = o.Name
	toSerialize["permissions"] = o.Permissions
	toSerialize["mimetype"] = o.Mimetype
	toSerialize["size"] = o.Size
	toSerialize["type"] = o.Type
	toSerialize["etag"] = o.Etag
	if o.Children != nil {
		toSerialize["children"] = o.Children
	}
	return toSerialize, nil
}

type NullableFilesSharingShareInfo struct {
	value *FilesSharingShareInfo
	isSet bool
}

func (v NullableFilesSharingShareInfo) Get() *FilesSharingShareInfo {
	return v.value
}

func (v *NullableFilesSharingShareInfo) Set(val *FilesSharingShareInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableFilesSharingShareInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableFilesSharingShareInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilesSharingShareInfo(val *FilesSharingShareInfo) *NullableFilesSharingShareInfo {
	return &NullableFilesSharingShareInfo{value: val, isSet: true}
}

func (v NullableFilesSharingShareInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilesSharingShareInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


