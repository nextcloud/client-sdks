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



from pydantic import BaseModel, Field
from openapi_client.models.user_status_statuses_find_all200_response_ocs import UserStatusStatusesFindAll200ResponseOcs

class UserStatusStatusesFindAll200Response(BaseModel):
    """
    UserStatusStatusesFindAll200Response
    """
    ocs: UserStatusStatusesFindAll200ResponseOcs = Field(...)
    __properties = ["ocs"]

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
    def from_json(cls, json_str: str) -> UserStatusStatusesFindAll200Response:
        """Create an instance of UserStatusStatusesFindAll200Response from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self):
        """Returns the dictionary representation of the model using alias"""
        _dict = self.dict(by_alias=True,
                          exclude={
                          },
                          exclude_none=True)
        # override the default output from pydantic by calling `to_dict()` of ocs
        if self.ocs:
            _dict['ocs'] = self.ocs.to_dict()
        return _dict

    @classmethod
    def from_dict(cls, obj: dict) -> UserStatusStatusesFindAll200Response:
        """Create an instance of UserStatusStatusesFindAll200Response from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return UserStatusStatusesFindAll200Response.parse_obj(obj)

        _obj = UserStatusStatusesFindAll200Response.parse_obj({
            "ocs": UserStatusStatusesFindAll200ResponseOcs.from_dict(obj.get("ocs")) if obj.get("ocs") is not None else None
        })
        return _obj

