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


from typing import Any, Dict
from pydantic import BaseModel, Field, StrictBool, StrictStr
from nextcloud_client.models.core_open_graph_object import CoreOpenGraphObject

class CoreResource(BaseModel):
    """
    CoreResource
    """
    rich_object_type: StrictStr = Field(..., alias="richObjectType")
    rich_object: Dict[str, Dict[str, Any]] = Field(..., alias="richObject")
    open_graph_object: CoreOpenGraphObject = Field(..., alias="openGraphObject")
    accessible: StrictBool = Field(...)
    additional_properties: Dict[str, Any] = {}
    __properties = ["richObjectType", "richObject", "openGraphObject", "accessible"]

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
    def from_json(cls, json_str: str) -> CoreResource:
        """Create an instance of CoreResource from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self):
        """Returns the dictionary representation of the model using alias"""
        _dict = self.dict(by_alias=True,
                          exclude={
                            "additional_properties"
                          },
                          exclude_none=True)
        # override the default output from pydantic by calling `to_dict()` of open_graph_object
        if self.open_graph_object:
            _dict['openGraphObject'] = self.open_graph_object.to_dict()
        # puts key-value pairs in additional_properties in the top level
        if self.additional_properties is not None:
            for _key, _value in self.additional_properties.items():
                _dict[_key] = _value

        return _dict

    @classmethod
    def from_dict(cls, obj: dict) -> CoreResource:
        """Create an instance of CoreResource from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return CoreResource.parse_obj(obj)

        _obj = CoreResource.parse_obj({
            "rich_object_type": obj.get("richObjectType"),
            "rich_object": obj.get("richObject"),
            "open_graph_object": CoreOpenGraphObject.from_dict(obj.get("openGraphObject")) if obj.get("openGraphObject") is not None else None,
            "accessible": obj.get("accessible")
        })
        # store additional fields in additional_properties
        for _key in obj.keys():
            if _key not in cls.__properties:
                _obj.additional_properties[_key] = obj.get(_key)

        return _obj


