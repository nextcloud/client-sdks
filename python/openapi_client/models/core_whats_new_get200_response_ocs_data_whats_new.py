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


from typing import List
from pydantic import BaseModel, Field, StrictStr, conlist

class CoreWhatsNewGet200ResponseOcsDataWhatsNew(BaseModel):
    """
    CoreWhatsNewGet200ResponseOcsDataWhatsNew
    """
    regular: conlist(StrictStr) = Field(...)
    admin: conlist(StrictStr) = Field(...)
    __properties = ["regular", "admin"]

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
    def from_json(cls, json_str: str) -> CoreWhatsNewGet200ResponseOcsDataWhatsNew:
        """Create an instance of CoreWhatsNewGet200ResponseOcsDataWhatsNew from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self):
        """Returns the dictionary representation of the model using alias"""
        _dict = self.dict(by_alias=True,
                          exclude={
                          },
                          exclude_none=True)
        return _dict

    @classmethod
    def from_dict(cls, obj: dict) -> CoreWhatsNewGet200ResponseOcsDataWhatsNew:
        """Create an instance of CoreWhatsNewGet200ResponseOcsDataWhatsNew from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return CoreWhatsNewGet200ResponseOcsDataWhatsNew.parse_obj(obj)

        _obj = CoreWhatsNewGet200ResponseOcsDataWhatsNew.parse_obj({
            "regular": obj.get("regular"),
            "admin": obj.get("admin")
        })
        return _obj

