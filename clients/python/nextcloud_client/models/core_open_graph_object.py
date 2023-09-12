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
from pydantic import BaseModel, Field, StrictStr

class CoreOpenGraphObject(BaseModel):
    """
    CoreOpenGraphObject
    """
    id: StrictStr = Field(...)
    name: StrictStr = Field(...)
    description: Optional[StrictStr] = Field(...)
    thumb: Optional[StrictStr] = Field(...)
    link: StrictStr = Field(...)
    additional_properties: Dict[str, Any] = {}
    __properties = ["id", "name", "description", "thumb", "link"]

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
    def from_json(cls, json_str: str) -> CoreOpenGraphObject:
        """Create an instance of CoreOpenGraphObject from a JSON string"""
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

        # set to None if description (nullable) is None
        # and __fields_set__ contains the field
        if self.description is None and "description" in self.__fields_set__:
            _dict['description'] = None

        # set to None if thumb (nullable) is None
        # and __fields_set__ contains the field
        if self.thumb is None and "thumb" in self.__fields_set__:
            _dict['thumb'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: dict) -> CoreOpenGraphObject:
        """Create an instance of CoreOpenGraphObject from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return CoreOpenGraphObject.parse_obj(obj)

        _obj = CoreOpenGraphObject.parse_obj({
            "id": obj.get("id"),
            "name": obj.get("name"),
            "description": obj.get("description"),
            "thumb": obj.get("thumb"),
            "link": obj.get("link")
        })
        # store additional fields in additional_properties
        for _key in obj.keys():
            if _key not in cls.__properties:
                _obj.additional_properties[_key] = obj.get(_key)

        return _obj


