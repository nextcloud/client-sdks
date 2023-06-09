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
from pydantic import BaseModel, Field, StrictBool, StrictInt
from openapi_client.models.core_ocs_get_capabilities200_response_ocs_data_capabilities_files_sharing_federation import CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingFederation
from openapi_client.models.core_ocs_get_capabilities200_response_ocs_data_capabilities_files_sharing_group import CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingGroup
from openapi_client.models.core_ocs_get_capabilities200_response_ocs_data_capabilities_files_sharing_public import CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingPublic
from openapi_client.models.core_ocs_get_capabilities200_response_ocs_data_capabilities_files_sharing_sharee import CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingSharee
from openapi_client.models.core_ocs_get_capabilities200_response_ocs_data_capabilities_files_sharing_user import CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUser

class CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing(BaseModel):
    """
    CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing
    """
    api_enabled: StrictBool = Field(...)
    public: CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingPublic = Field(...)
    user: CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUser = Field(...)
    resharing: StrictBool = Field(...)
    group_sharing: Optional[StrictBool] = None
    group: Optional[CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingGroup] = None
    default_permissions: Optional[StrictInt] = None
    federation: CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingFederation = Field(...)
    sharee: CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingSharee = Field(...)
    __properties = ["api_enabled", "public", "user", "resharing", "group_sharing", "group", "default_permissions", "federation", "sharee"]

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
    def from_json(cls, json_str: str) -> CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing:
        """Create an instance of CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self):
        """Returns the dictionary representation of the model using alias"""
        _dict = self.dict(by_alias=True,
                          exclude={
                          },
                          exclude_none=True)
        # override the default output from pydantic by calling `to_dict()` of public
        if self.public:
            _dict['public'] = self.public.to_dict()
        # override the default output from pydantic by calling `to_dict()` of user
        if self.user:
            _dict['user'] = self.user.to_dict()
        # override the default output from pydantic by calling `to_dict()` of group
        if self.group:
            _dict['group'] = self.group.to_dict()
        # override the default output from pydantic by calling `to_dict()` of federation
        if self.federation:
            _dict['federation'] = self.federation.to_dict()
        # override the default output from pydantic by calling `to_dict()` of sharee
        if self.sharee:
            _dict['sharee'] = self.sharee.to_dict()
        return _dict

    @classmethod
    def from_dict(cls, obj: dict) -> CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing:
        """Create an instance of CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing.parse_obj(obj)

        _obj = CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing.parse_obj({
            "api_enabled": obj.get("api_enabled"),
            "public": CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingPublic.from_dict(obj.get("public")) if obj.get("public") is not None else None,
            "user": CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUser.from_dict(obj.get("user")) if obj.get("user") is not None else None,
            "resharing": obj.get("resharing"),
            "group_sharing": obj.get("group_sharing"),
            "group": CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingGroup.from_dict(obj.get("group")) if obj.get("group") is not None else None,
            "default_permissions": obj.get("default_permissions"),
            "federation": CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingFederation.from_dict(obj.get("federation")) if obj.get("federation") is not None else None,
            "sharee": CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingSharee.from_dict(obj.get("sharee")) if obj.get("sharee") is not None else None
        })
        return _obj

