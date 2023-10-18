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



from pydantic import BaseModel, Field, StrictInt, StrictStr

class FilesSharingShareeRemoteAllOfValue(BaseModel):
    """
    FilesSharingShareeRemoteAllOfValue
    """
    share_type: StrictInt = Field(..., alias="shareType")
    share_with: StrictStr = Field(..., alias="shareWith")
    server: StrictStr = Field(...)
    __properties = ["shareType", "shareWith", "server"]

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
    def from_json(cls, json_str: str) -> FilesSharingShareeRemoteAllOfValue:
        """Create an instance of FilesSharingShareeRemoteAllOfValue from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self):
        """Returns the dictionary representation of the model using alias"""
        _dict = self.dict(by_alias=True,
                          exclude={
                          },
                          exclude_none=True)
        return _dict

    @classmethod
    def from_dict(cls, obj: dict) -> FilesSharingShareeRemoteAllOfValue:
        """Create an instance of FilesSharingShareeRemoteAllOfValue from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return FilesSharingShareeRemoteAllOfValue.parse_obj(obj)

        _obj = FilesSharingShareeRemoteAllOfValue.parse_obj({
            "share_type": obj.get("shareType"),
            "share_with": obj.get("shareWith"),
            "server": obj.get("server")
        })
        return _obj
