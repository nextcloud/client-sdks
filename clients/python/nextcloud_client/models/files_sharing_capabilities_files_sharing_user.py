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
from pydantic import BaseModel, Field, StrictBool
from nextcloud_client.models.files_sharing_capabilities_files_sharing_user_expire_date import FilesSharingCapabilitiesFilesSharingUserExpireDate

class FilesSharingCapabilitiesFilesSharingUser(BaseModel):
    """
    FilesSharingCapabilitiesFilesSharingUser
    """
    send_mail: StrictBool = Field(...)
    expire_date: Optional[FilesSharingCapabilitiesFilesSharingUserExpireDate] = None
    additional_properties: Dict[str, Any] = {}
    __properties = ["send_mail", "expire_date"]

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
    def from_json(cls, json_str: str) -> FilesSharingCapabilitiesFilesSharingUser:
        """Create an instance of FilesSharingCapabilitiesFilesSharingUser from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self):
        """Returns the dictionary representation of the model using alias"""
        _dict = self.dict(by_alias=True,
                          exclude={
                            "additional_properties"
                          },
                          exclude_none=True)
        # override the default output from pydantic by calling `to_dict()` of expire_date
        if self.expire_date:
            _dict['expire_date'] = self.expire_date.to_dict()
        # puts key-value pairs in additional_properties in the top level
        if self.additional_properties is not None:
            for _key, _value in self.additional_properties.items():
                _dict[_key] = _value

        return _dict

    @classmethod
    def from_dict(cls, obj: dict) -> FilesSharingCapabilitiesFilesSharingUser:
        """Create an instance of FilesSharingCapabilitiesFilesSharingUser from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return FilesSharingCapabilitiesFilesSharingUser.parse_obj(obj)

        _obj = FilesSharingCapabilitiesFilesSharingUser.parse_obj({
            "send_mail": obj.get("send_mail"),
            "expire_date": FilesSharingCapabilitiesFilesSharingUserExpireDate.from_dict(obj.get("expire_date")) if obj.get("expire_date") is not None else None
        })
        # store additional fields in additional_properties
        for _key in obj.keys():
            if _key not in cls.__properties:
                _obj.additional_properties[_key] = obj.get(_key)

        return _obj


