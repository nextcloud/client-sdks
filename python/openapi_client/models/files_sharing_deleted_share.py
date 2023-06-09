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

class FilesSharingDeletedShare(BaseModel):
    """
    FilesSharingDeletedShare
    """
    id: StrictStr = Field(...)
    share_type: StrictInt = Field(...)
    uid_owner: StrictStr = Field(...)
    displayname_owner: StrictStr = Field(...)
    permissions: StrictInt = Field(...)
    stime: StrictInt = Field(...)
    uid_file_owner: StrictStr = Field(...)
    displayname_file_owner: StrictStr = Field(...)
    path: StrictStr = Field(...)
    item_type: StrictStr = Field(...)
    mimetype: StrictStr = Field(...)
    storage: StrictInt = Field(...)
    item_source: StrictInt = Field(...)
    file_source: StrictInt = Field(...)
    file_parent: StrictInt = Field(...)
    file_target: StrictInt = Field(...)
    expiration: Optional[StrictStr] = Field(...)
    share_with: Optional[StrictStr] = Field(...)
    share_with_displayname: Optional[StrictStr] = Field(...)
    share_with_link: Optional[StrictStr] = Field(...)
    __properties = ["id", "share_type", "uid_owner", "displayname_owner", "permissions", "stime", "uid_file_owner", "displayname_file_owner", "path", "item_type", "mimetype", "storage", "item_source", "file_source", "file_parent", "file_target", "expiration", "share_with", "share_with_displayname", "share_with_link"]

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
    def from_json(cls, json_str: str) -> FilesSharingDeletedShare:
        """Create an instance of FilesSharingDeletedShare from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self):
        """Returns the dictionary representation of the model using alias"""
        _dict = self.dict(by_alias=True,
                          exclude={
                          },
                          exclude_none=True)
        # set to None if expiration (nullable) is None
        # and __fields_set__ contains the field
        if self.expiration is None and "expiration" in self.__fields_set__:
            _dict['expiration'] = None

        # set to None if share_with (nullable) is None
        # and __fields_set__ contains the field
        if self.share_with is None and "share_with" in self.__fields_set__:
            _dict['share_with'] = None

        # set to None if share_with_displayname (nullable) is None
        # and __fields_set__ contains the field
        if self.share_with_displayname is None and "share_with_displayname" in self.__fields_set__:
            _dict['share_with_displayname'] = None

        # set to None if share_with_link (nullable) is None
        # and __fields_set__ contains the field
        if self.share_with_link is None and "share_with_link" in self.__fields_set__:
            _dict['share_with_link'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: dict) -> FilesSharingDeletedShare:
        """Create an instance of FilesSharingDeletedShare from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return FilesSharingDeletedShare.parse_obj(obj)

        _obj = FilesSharingDeletedShare.parse_obj({
            "id": obj.get("id"),
            "share_type": obj.get("share_type"),
            "uid_owner": obj.get("uid_owner"),
            "displayname_owner": obj.get("displayname_owner"),
            "permissions": obj.get("permissions"),
            "stime": obj.get("stime"),
            "uid_file_owner": obj.get("uid_file_owner"),
            "displayname_file_owner": obj.get("displayname_file_owner"),
            "path": obj.get("path"),
            "item_type": obj.get("item_type"),
            "mimetype": obj.get("mimetype"),
            "storage": obj.get("storage"),
            "item_source": obj.get("item_source"),
            "file_source": obj.get("file_source"),
            "file_parent": obj.get("file_parent"),
            "file_target": obj.get("file_target"),
            "expiration": obj.get("expiration"),
            "share_with": obj.get("share_with"),
            "share_with_displayname": obj.get("share_with_displayname"),
            "share_with_link": obj.get("share_with_link")
        })
        return _obj

