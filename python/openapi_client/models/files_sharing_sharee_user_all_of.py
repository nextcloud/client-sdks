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



from pydantic import BaseModel, Field, StrictStr
from openapi_client.models.files_sharing_sharee_user_all_of_status import FilesSharingShareeUserAllOfStatus
from openapi_client.models.files_sharing_sharee_value import FilesSharingShareeValue

class FilesSharingShareeUserAllOf(BaseModel):
    """
    FilesSharingShareeUserAllOf
    """
    subline: StrictStr = Field(...)
    icon: StrictStr = Field(...)
    share_with_display_name_unique: StrictStr = Field(..., alias="shareWithDisplayNameUnique")
    status: FilesSharingShareeUserAllOfStatus = Field(...)
    value: FilesSharingShareeValue = Field(...)
    __properties = ["subline", "icon", "shareWithDisplayNameUnique", "status", "value"]

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
    def from_json(cls, json_str: str) -> FilesSharingShareeUserAllOf:
        """Create an instance of FilesSharingShareeUserAllOf from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self):
        """Returns the dictionary representation of the model using alias"""
        _dict = self.dict(by_alias=True,
                          exclude={
                          },
                          exclude_none=True)
        # override the default output from pydantic by calling `to_dict()` of status
        if self.status:
            _dict['status'] = self.status.to_dict()
        # override the default output from pydantic by calling `to_dict()` of value
        if self.value:
            _dict['value'] = self.value.to_dict()
        return _dict

    @classmethod
    def from_dict(cls, obj: dict) -> FilesSharingShareeUserAllOf:
        """Create an instance of FilesSharingShareeUserAllOf from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return FilesSharingShareeUserAllOf.parse_obj(obj)

        _obj = FilesSharingShareeUserAllOf.parse_obj({
            "subline": obj.get("subline"),
            "icon": obj.get("icon"),
            "share_with_display_name_unique": obj.get("shareWithDisplayNameUnique"),
            "status": FilesSharingShareeUserAllOfStatus.from_dict(obj.get("status")) if obj.get("status") is not None else None,
            "value": FilesSharingShareeValue.from_dict(obj.get("value")) if obj.get("value") is not None else None
        })
        return _obj

