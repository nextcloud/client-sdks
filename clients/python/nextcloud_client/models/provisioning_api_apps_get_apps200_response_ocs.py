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
from pydantic import BaseModel, Field
from nextcloud_client.models.ocs_meta import OCSMeta
from nextcloud_client.models.provisioning_api_apps_get_apps200_response_ocs_data import ProvisioningApiAppsGetApps200ResponseOcsData

class ProvisioningApiAppsGetApps200ResponseOcs(BaseModel):
    """
    ProvisioningApiAppsGetApps200ResponseOcs
    """
    meta: OCSMeta = Field(...)
    data: ProvisioningApiAppsGetApps200ResponseOcsData = Field(...)
    additional_properties: Dict[str, Any] = {}
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
    def from_json(cls, json_str: str) -> ProvisioningApiAppsGetApps200ResponseOcs:
        """Create an instance of ProvisioningApiAppsGetApps200ResponseOcs from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self):
        """Returns the dictionary representation of the model using alias"""
        _dict = self.dict(by_alias=True,
                          exclude={
                            "additional_properties"
                          },
                          exclude_none=True)
        # override the default output from pydantic by calling `to_dict()` of meta
        if self.meta:
            _dict['meta'] = self.meta.to_dict()
        # override the default output from pydantic by calling `to_dict()` of data
        if self.data:
            _dict['data'] = self.data.to_dict()
        # puts key-value pairs in additional_properties in the top level
        if self.additional_properties is not None:
            for _key, _value in self.additional_properties.items():
                _dict[_key] = _value

        return _dict

    @classmethod
    def from_dict(cls, obj: dict) -> ProvisioningApiAppsGetApps200ResponseOcs:
        """Create an instance of ProvisioningApiAppsGetApps200ResponseOcs from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return ProvisioningApiAppsGetApps200ResponseOcs.parse_obj(obj)

        _obj = ProvisioningApiAppsGetApps200ResponseOcs.parse_obj({
            "meta": OCSMeta.from_dict(obj.get("meta")) if obj.get("meta") is not None else None,
            "data": ProvisioningApiAppsGetApps200ResponseOcsData.from_dict(obj.get("data")) if obj.get("data") is not None else None
        })
        # store additional fields in additional_properties
        for _key in obj.keys():
            if _key not in cls.__properties:
                _obj.additional_properties[_key] = obj.get(_key)

        return _obj


