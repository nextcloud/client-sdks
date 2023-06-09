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


from typing import Dict, List
from pydantic import BaseModel, Field, conlist
from openapi_client.models.dashboard_widget_item import DashboardWidgetItem
from openapi_client.models.ocs_meta import OCSMeta

class DashboardDashboardApiGetWidgetItems200ResponseOcs(BaseModel):
    """
    DashboardDashboardApiGetWidgetItems200ResponseOcs
    """
    meta: OCSMeta = Field(...)
    data: Dict[str, conlist(DashboardWidgetItem)] = Field(...)
    __properties = ["meta", "data"]

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
    def from_json(cls, json_str: str) -> DashboardDashboardApiGetWidgetItems200ResponseOcs:
        """Create an instance of DashboardDashboardApiGetWidgetItems200ResponseOcs from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self):
        """Returns the dictionary representation of the model using alias"""
        _dict = self.dict(by_alias=True,
                          exclude={
                          },
                          exclude_none=True)
        # override the default output from pydantic by calling `to_dict()` of meta
        if self.meta:
            _dict['meta'] = self.meta.to_dict()
        # override the default output from pydantic by calling `to_dict()` of each value in data (dict)
        _field_dict = {}
        if self.data:
            for _key in self.data:
                if self.data[_key]:
                    _field_dict[_key] = self.data[_key].to_dict()
            _dict['data'] = _field_dict
        return _dict

    @classmethod
    def from_dict(cls, obj: dict) -> DashboardDashboardApiGetWidgetItems200ResponseOcs:
        """Create an instance of DashboardDashboardApiGetWidgetItems200ResponseOcs from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return DashboardDashboardApiGetWidgetItems200ResponseOcs.parse_obj(obj)

        _obj = DashboardDashboardApiGetWidgetItems200ResponseOcs.parse_obj({
            "meta": OCSMeta.from_dict(obj.get("meta")) if obj.get("meta") is not None else None,
            "data": dict(
                (_k, [(_ik, DashboardWidgetItem.from_dict(_iv))]
                    for _ik, _iv in _v.items()
                    if _v is not None
                    else None
                )
                for _k, _v in obj.get("data").items()
            )
            if obj.get("data") is not None
            else None
        })
        return _obj

