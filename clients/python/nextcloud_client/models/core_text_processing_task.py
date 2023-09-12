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
from pydantic import BaseModel, Field, StrictInt, StrictStr

class CoreTextProcessingTask(BaseModel):
    """
    CoreTextProcessingTask
    """
    id: Optional[StrictInt] = Field(...)
    type: StrictStr = Field(...)
    status: StrictInt = Field(...)
    user_id: Optional[StrictStr] = Field(..., alias="userId")
    app_id: StrictStr = Field(..., alias="appId")
    input: StrictStr = Field(...)
    output: Optional[StrictStr] = Field(...)
    identifier: StrictStr = Field(...)
    additional_properties: Dict[str, Any] = {}
    __properties = ["id", "type", "status", "userId", "appId", "input", "output", "identifier"]

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
    def from_json(cls, json_str: str) -> CoreTextProcessingTask:
        """Create an instance of CoreTextProcessingTask from a JSON string"""
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

        # set to None if id (nullable) is None
        # and __fields_set__ contains the field
        if self.id is None and "id" in self.__fields_set__:
            _dict['id'] = None

        # set to None if user_id (nullable) is None
        # and __fields_set__ contains the field
        if self.user_id is None and "user_id" in self.__fields_set__:
            _dict['userId'] = None

        # set to None if output (nullable) is None
        # and __fields_set__ contains the field
        if self.output is None and "output" in self.__fields_set__:
            _dict['output'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: dict) -> CoreTextProcessingTask:
        """Create an instance of CoreTextProcessingTask from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return CoreTextProcessingTask.parse_obj(obj)

        _obj = CoreTextProcessingTask.parse_obj({
            "id": obj.get("id"),
            "type": obj.get("type"),
            "status": obj.get("status"),
            "user_id": obj.get("userId"),
            "app_id": obj.get("appId"),
            "input": obj.get("input"),
            "output": obj.get("output"),
            "identifier": obj.get("identifier")
        })
        # store additional fields in additional_properties
        for _key in obj.keys():
            if _key not in cls.__properties:
                _obj.additional_properties[_key] = obj.get(_key)

        return _obj


