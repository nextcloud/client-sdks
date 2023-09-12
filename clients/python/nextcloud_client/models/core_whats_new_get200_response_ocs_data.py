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
from nextcloud_client.models.core_whats_new_get200_response_ocs_data_whats_new import CoreWhatsNewGet200ResponseOcsDataWhatsNew

class CoreWhatsNewGet200ResponseOcsData(BaseModel):
    """
    CoreWhatsNewGet200ResponseOcsData
    """
    changelog_url: StrictStr = Field(..., alias="changelogURL")
    product: StrictStr = Field(...)
    version: StrictStr = Field(...)
    whats_new: Optional[CoreWhatsNewGet200ResponseOcsDataWhatsNew] = Field(None, alias="whatsNew")
    additional_properties: Dict[str, Any] = {}
    __properties = ["changelogURL", "product", "version", "whatsNew"]

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
    def from_json(cls, json_str: str) -> CoreWhatsNewGet200ResponseOcsData:
        """Create an instance of CoreWhatsNewGet200ResponseOcsData from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self):
        """Returns the dictionary representation of the model using alias"""
        _dict = self.dict(by_alias=True,
                          exclude={
                            "additional_properties"
                          },
                          exclude_none=True)
        # override the default output from pydantic by calling `to_dict()` of whats_new
        if self.whats_new:
            _dict['whatsNew'] = self.whats_new.to_dict()
        # puts key-value pairs in additional_properties in the top level
        if self.additional_properties is not None:
            for _key, _value in self.additional_properties.items():
                _dict[_key] = _value

        return _dict

    @classmethod
    def from_dict(cls, obj: dict) -> CoreWhatsNewGet200ResponseOcsData:
        """Create an instance of CoreWhatsNewGet200ResponseOcsData from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return CoreWhatsNewGet200ResponseOcsData.parse_obj(obj)

        _obj = CoreWhatsNewGet200ResponseOcsData.parse_obj({
            "changelog_url": obj.get("changelogURL"),
            "product": obj.get("product"),
            "version": obj.get("version"),
            "whats_new": CoreWhatsNewGet200ResponseOcsDataWhatsNew.from_dict(obj.get("whatsNew")) if obj.get("whatsNew") is not None else None
        })
        # store additional fields in additional_properties
        for _key in obj.keys():
            if _key not in cls.__properties:
                _obj.additional_properties[_key] = obj.get(_key)

        return _obj

