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
from pydantic import BaseModel, Field, StrictBool, StrictInt, StrictStr

class UserStatusPrivate(BaseModel):
    """
    UserStatusPrivate
    """
    user_id: StrictStr = Field(..., alias="userId")
    message: Optional[StrictStr] = Field(...)
    icon: Optional[StrictStr] = Field(...)
    clear_at: Optional[StrictInt] = Field(..., alias="clearAt")
    status: StrictStr = Field(...)
    message_id: Optional[StrictStr] = Field(..., alias="messageId")
    message_is_predefined: StrictBool = Field(..., alias="messageIsPredefined")
    status_is_user_defined: StrictBool = Field(..., alias="statusIsUserDefined")
    __properties = ["userId", "message", "icon", "clearAt", "status", "messageId", "messageIsPredefined", "statusIsUserDefined"]

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
    def from_json(cls, json_str: str) -> UserStatusPrivate:
        """Create an instance of UserStatusPrivate from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self):
        """Returns the dictionary representation of the model using alias"""
        _dict = self.dict(by_alias=True,
                          exclude={
                          },
                          exclude_none=True)
        # set to None if message (nullable) is None
        # and __fields_set__ contains the field
        if self.message is None and "message" in self.__fields_set__:
            _dict['message'] = None

        # set to None if icon (nullable) is None
        # and __fields_set__ contains the field
        if self.icon is None and "icon" in self.__fields_set__:
            _dict['icon'] = None

        # set to None if clear_at (nullable) is None
        # and __fields_set__ contains the field
        if self.clear_at is None and "clear_at" in self.__fields_set__:
            _dict['clearAt'] = None

        # set to None if message_id (nullable) is None
        # and __fields_set__ contains the field
        if self.message_id is None and "message_id" in self.__fields_set__:
            _dict['messageId'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: dict) -> UserStatusPrivate:
        """Create an instance of UserStatusPrivate from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return UserStatusPrivate.parse_obj(obj)

        _obj = UserStatusPrivate.parse_obj({
            "user_id": obj.get("userId"),
            "message": obj.get("message"),
            "icon": obj.get("icon"),
            "clear_at": obj.get("clearAt"),
            "status": obj.get("status"),
            "message_id": obj.get("messageId"),
            "message_is_predefined": obj.get("messageIsPredefined"),
            "status_is_user_defined": obj.get("statusIsUserDefined")
        })
        return _obj

