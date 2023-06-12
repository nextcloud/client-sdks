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


from typing import Optional
from pydantic import BaseModel, Field, StrictInt, StrictStr
from openapi_client.models.files_sharing_sharee_lookup_all_of_extra import FilesSharingShareeLookupAllOfExtra
from openapi_client.models.files_sharing_sharee_lookup_all_of_value import FilesSharingShareeLookupAllOfValue

class FilesSharingShareeLookup(BaseModel):
    """
    FilesSharingShareeLookup
    """
    count: Optional[StrictInt] = Field(...)
    label: StrictStr = Field(...)
    extra: FilesSharingShareeLookupAllOfExtra = Field(...)
    value: FilesSharingShareeLookupAllOfValue = Field(...)
    __properties = ["count", "label", "extra", "value"]

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
    def from_json(cls, json_str: str) -> FilesSharingShareeLookup:
        """Create an instance of FilesSharingShareeLookup from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self):
        """Returns the dictionary representation of the model using alias"""
        _dict = self.dict(by_alias=True,
                          exclude={
                          },
                          exclude_none=True)
        # override the default output from pydantic by calling `to_dict()` of extra
        if self.extra:
            _dict['extra'] = self.extra.to_dict()
        # override the default output from pydantic by calling `to_dict()` of value
        if self.value:
            _dict['value'] = self.value.to_dict()
        # set to None if count (nullable) is None
        # and __fields_set__ contains the field
        if self.count is None and "count" in self.__fields_set__:
            _dict['count'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: dict) -> FilesSharingShareeLookup:
        """Create an instance of FilesSharingShareeLookup from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return FilesSharingShareeLookup.parse_obj(obj)

        _obj = FilesSharingShareeLookup.parse_obj({
            "count": obj.get("count"),
            "label": obj.get("label"),
            "extra": FilesSharingShareeLookupAllOfExtra.from_dict(obj.get("extra")) if obj.get("extra") is not None else None,
            "value": FilesSharingShareeLookupAllOfValue.from_dict(obj.get("value")) if obj.get("value") is not None else None
        })
        return _obj

