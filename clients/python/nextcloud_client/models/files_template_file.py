# coding: utf-8

"""
    nextcloud

    Nextcloud APIs

    The version of the OpenAPI document: 0.0.1
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


from __future__ import annotations
import pprint
import re  # noqa: F401
import json


from typing import Any, Dict, Optional
from pydantic import BaseModel, Field, StrictBool, StrictInt, StrictStr

class FilesTemplateFile(BaseModel):
    """
    FilesTemplateFile
    """
    basename: StrictStr = Field(...)
    etag: StrictStr = Field(...)
    fileid: StrictInt = Field(...)
    filename: Optional[StrictStr] = Field(...)
    lastmod: StrictInt = Field(...)
    mime: StrictStr = Field(...)
    size: StrictInt = Field(...)
    type: StrictStr = Field(...)
    has_preview: StrictBool = Field(..., alias="hasPreview")
    additional_properties: Dict[str, Any] = {}
    __properties = ["basename", "etag", "fileid", "filename", "lastmod", "mime", "size", "type", "hasPreview"]

    class Config:
        """Pydantic configuration"""
        allow_population_by_field_name = True
        validate_assignment = True

    def to_str(self) -> str:
        """Returns the string representation of the model using alias"""
        return pprint.pformat(self.dict(by_alias=True))

    def to_json(self) -> str:
        """Returns the JSON representation of the model using alias"""
        return json.dumps(self.to_dict())

    @classmethod
    def from_json(cls, json_str: str) -> FilesTemplateFile:
        """Create an instance of FilesTemplateFile from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self):
        """Returns the dictionary representation of the model using alias"""
        _dict = self.dict(by_alias=True,
                          exclude={
                            "additional_properties"
                          },
                          exclude_none=True)
        # puts key-value pairs in additional_properties in the top level
        if self.additional_properties is not None:
            for _key, _value in self.additional_properties.items():
                _dict[_key] = _value

        # set to None if filename (nullable) is None
        # and __fields_set__ contains the field
        if self.filename is None and "filename" in self.__fields_set__:
            _dict['filename'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: dict) -> FilesTemplateFile:
        """Create an instance of FilesTemplateFile from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return FilesTemplateFile.parse_obj(obj)

        _obj = FilesTemplateFile.parse_obj({
            "basename": obj.get("basename"),
            "etag": obj.get("etag"),
            "fileid": obj.get("fileid"),
            "filename": obj.get("filename"),
            "lastmod": obj.get("lastmod"),
            "mime": obj.get("mime"),
            "size": obj.get("size"),
            "type": obj.get("type"),
            "has_preview": obj.get("hasPreview")
        })
        # store additional fields in additional_properties
        for _key in obj.keys():
            if _key not in cls.__properties:
                _obj.additional_properties[_key] = obj.get(_key)

        return _obj


