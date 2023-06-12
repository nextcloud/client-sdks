# coding: utf-8

"""
    nextcloud

    Nextcloud APIs  # noqa: E501

    The version of the OpenAPI document: 0.0.1
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""


from __future__ import annotations
import pprint
import re  # noqa: F401
import json


from typing import Any, Dict, List, Optional
from pydantic import BaseModel, Field, StrictInt, StrictStr, conlist

class FilesSharingShareInfo(BaseModel):
    """
    FilesSharingShareInfo
    """
    id: StrictInt = Field(...)
    parent_id: StrictInt = Field(..., alias="parentId")
    mtime: StrictInt = Field(...)
    name: StrictStr = Field(...)
    permissions: StrictInt = Field(...)
    mimetype: StrictStr = Field(...)
    size: StrictInt = Field(...)
    type: StrictStr = Field(...)
    etag: StrictStr = Field(...)
    children: Optional[conlist(Dict[str, Dict[str, Any]])] = Field(...)
    __properties = ["id", "parentId", "mtime", "name", "permissions", "mimetype", "size", "type", "etag", "children"]

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
    def from_json(cls, json_str: str) -> FilesSharingShareInfo:
        """Create an instance of FilesSharingShareInfo from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self):
        """Returns the dictionary representation of the model using alias"""
        _dict = self.dict(by_alias=True,
                          exclude={
                          },
                          exclude_none=True)
        # set to None if children (nullable) is None
        # and __fields_set__ contains the field
        if self.children is None and "children" in self.__fields_set__:
            _dict['children'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: dict) -> FilesSharingShareInfo:
        """Create an instance of FilesSharingShareInfo from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return FilesSharingShareInfo.parse_obj(obj)

        _obj = FilesSharingShareInfo.parse_obj({
            "id": obj.get("id"),
            "parent_id": obj.get("parentId"),
            "mtime": obj.get("mtime"),
            "name": obj.get("name"),
            "permissions": obj.get("permissions"),
            "mimetype": obj.get("mimetype"),
            "size": obj.get("size"),
            "type": obj.get("type"),
            "etag": obj.get("etag"),
            "children": obj.get("children")
        })
        return _obj

