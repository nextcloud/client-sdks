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


from typing import Any, Dict, List
from pydantic import BaseModel, Field, conlist
from nextcloud_client.models.updatenotification_app import UpdatenotificationApp

class UpdatenotificationApiGetAppList200ResponseOcsData(BaseModel):
    """
    UpdatenotificationApiGetAppList200ResponseOcsData
    """
    missing: conlist(UpdatenotificationApp) = Field(...)
    available: conlist(UpdatenotificationApp) = Field(...)
    additional_properties: Dict[str, Any] = {}
    __properties = ["missing", "available"]

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
    def from_json(cls, json_str: str) -> UpdatenotificationApiGetAppList200ResponseOcsData:
        """Create an instance of UpdatenotificationApiGetAppList200ResponseOcsData from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self):
        """Returns the dictionary representation of the model using alias"""
        _dict = self.dict(by_alias=True,
                          exclude={
                            "additional_properties"
                          },
                          exclude_none=True)
        # override the default output from pydantic by calling `to_dict()` of each item in missing (list)
        _items = []
        if self.missing:
            for _item in self.missing:
                if _item:
                    _items.append(_item.to_dict())
            _dict['missing'] = _items
        # override the default output from pydantic by calling `to_dict()` of each item in available (list)
        _items = []
        if self.available:
            for _item in self.available:
                if _item:
                    _items.append(_item.to_dict())
            _dict['available'] = _items
        # puts key-value pairs in additional_properties in the top level
        if self.additional_properties is not None:
            for _key, _value in self.additional_properties.items():
                _dict[_key] = _value

        return _dict

    @classmethod
    def from_dict(cls, obj: dict) -> UpdatenotificationApiGetAppList200ResponseOcsData:
        """Create an instance of UpdatenotificationApiGetAppList200ResponseOcsData from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return UpdatenotificationApiGetAppList200ResponseOcsData.parse_obj(obj)

        _obj = UpdatenotificationApiGetAppList200ResponseOcsData.parse_obj({
            "missing": [UpdatenotificationApp.from_dict(_item) for _item in obj.get("missing")] if obj.get("missing") is not None else None,
            "available": [UpdatenotificationApp.from_dict(_item) for _item in obj.get("available")] if obj.get("available") is not None else None
        })
        # store additional fields in additional_properties
        for _key in obj.keys():
            if _key not in cls.__properties:
                _obj.additional_properties[_key] = obj.get(_key)

        return _obj


