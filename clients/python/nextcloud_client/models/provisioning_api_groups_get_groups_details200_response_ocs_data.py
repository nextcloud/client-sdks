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
from nextcloud_client.models.provisioning_api_group_details import ProvisioningApiGroupDetails

class ProvisioningApiGroupsGetGroupsDetails200ResponseOcsData(BaseModel):
    """
    ProvisioningApiGroupsGetGroupsDetails200ResponseOcsData
    """
    groups: conlist(ProvisioningApiGroupDetails) = Field(...)
    additional_properties: Dict[str, Any] = {}
    __properties = ["groups"]

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
    def from_json(cls, json_str: str) -> ProvisioningApiGroupsGetGroupsDetails200ResponseOcsData:
        """Create an instance of ProvisioningApiGroupsGetGroupsDetails200ResponseOcsData from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self):
        """Returns the dictionary representation of the model using alias"""
        _dict = self.dict(by_alias=True,
                          exclude={
                            "additional_properties"
                          },
                          exclude_none=True)
        # override the default output from pydantic by calling `to_dict()` of each item in groups (list)
        _items = []
        if self.groups:
            for _item in self.groups:
                if _item:
                    _items.append(_item.to_dict())
            _dict['groups'] = _items
        # puts key-value pairs in additional_properties in the top level
        if self.additional_properties is not None:
            for _key, _value in self.additional_properties.items():
                _dict[_key] = _value

        return _dict

    @classmethod
    def from_dict(cls, obj: dict) -> ProvisioningApiGroupsGetGroupsDetails200ResponseOcsData:
        """Create an instance of ProvisioningApiGroupsGetGroupsDetails200ResponseOcsData from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return ProvisioningApiGroupsGetGroupsDetails200ResponseOcsData.parse_obj(obj)

        _obj = ProvisioningApiGroupsGetGroupsDetails200ResponseOcsData.parse_obj({
            "groups": [ProvisioningApiGroupDetails.from_dict(_item) for _item in obj.get("groups")] if obj.get("groups") is not None else None
        })
        # store additional fields in additional_properties
        for _key in obj.keys():
            if _key not in cls.__properties:
                _obj.additional_properties[_key] = obj.get(_key)

        return _obj


