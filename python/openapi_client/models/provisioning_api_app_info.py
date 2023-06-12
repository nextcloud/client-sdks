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


from typing import Any, Dict, Optional
from pydantic import BaseModel, Field, StrictBool, StrictInt, StrictStr

class ProvisioningApiAppInfo(BaseModel):
    """
    ProvisioningApiAppInfo
    """
    active: Optional[StrictBool] = Field(...)
    activity: Optional[Dict[str, Any]] = Field(...)
    author: Optional[Dict[str, Any]] = Field(...)
    background_jobs: Optional[Dict[str, Any]] = Field(..., alias="background-jobs")
    bugs: Optional[Dict[str, Any]] = Field(...)
    category: Optional[Dict[str, Any]] = Field(...)
    collaboration: Optional[Dict[str, Any]] = Field(...)
    commands: Optional[Dict[str, Any]] = Field(...)
    default_enable: Optional[Dict[str, Any]] = Field(...)
    dependencies: Optional[Dict[str, Any]] = Field(...)
    description: StrictStr = Field(...)
    discussion: Optional[Dict[str, Any]] = Field(...)
    documentation: Optional[Dict[str, Any]] = Field(...)
    groups: Optional[Dict[str, Any]] = Field(...)
    id: StrictStr = Field(...)
    info: Optional[Dict[str, Any]] = Field(...)
    internal: Optional[StrictBool] = Field(...)
    level: Optional[StrictInt] = Field(...)
    licence: Optional[Dict[str, Any]] = Field(...)
    name: StrictStr = Field(...)
    namespace: Optional[Dict[str, Any]] = Field(...)
    navigations: Optional[Dict[str, Any]] = Field(...)
    preview: Optional[Dict[str, Any]] = Field(...)
    preview_as_icon: Optional[StrictBool] = Field(..., alias="previewAsIcon")
    public: Optional[Dict[str, Any]] = Field(...)
    remote: Optional[Dict[str, Any]] = Field(...)
    removable: Optional[StrictBool] = Field(...)
    repair_steps: Optional[Dict[str, Any]] = Field(..., alias="repair-steps")
    repository: Optional[Dict[str, Any]] = Field(...)
    sabre: Optional[Dict[str, Any]] = Field(...)
    screenshot: Optional[Dict[str, Any]] = Field(...)
    settings: Optional[Dict[str, Any]] = Field(...)
    summary: StrictStr = Field(...)
    trash: Optional[Dict[str, Any]] = Field(...)
    two_factor_providers: Optional[Dict[str, Any]] = Field(..., alias="two-factor-providers")
    types: Optional[Dict[str, Any]] = Field(...)
    version: StrictStr = Field(...)
    versions: Optional[Dict[str, Any]] = Field(...)
    website: Optional[Dict[str, Any]] = Field(...)
    __properties = ["active", "activity", "author", "background-jobs", "bugs", "category", "collaboration", "commands", "default_enable", "dependencies", "description", "discussion", "documentation", "groups", "id", "info", "internal", "level", "licence", "name", "namespace", "navigations", "preview", "previewAsIcon", "public", "remote", "removable", "repair-steps", "repository", "sabre", "screenshot", "settings", "summary", "trash", "two-factor-providers", "types", "version", "versions", "website"]

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
    def from_json(cls, json_str: str) -> ProvisioningApiAppInfo:
        """Create an instance of ProvisioningApiAppInfo from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self):
        """Returns the dictionary representation of the model using alias"""
        _dict = self.dict(by_alias=True,
                          exclude={
                          },
                          exclude_none=True)
        # set to None if active (nullable) is None
        # and __fields_set__ contains the field
        if self.active is None and "active" in self.__fields_set__:
            _dict['active'] = None

        # set to None if activity (nullable) is None
        # and __fields_set__ contains the field
        if self.activity is None and "activity" in self.__fields_set__:
            _dict['activity'] = None

        # set to None if author (nullable) is None
        # and __fields_set__ contains the field
        if self.author is None and "author" in self.__fields_set__:
            _dict['author'] = None

        # set to None if background_jobs (nullable) is None
        # and __fields_set__ contains the field
        if self.background_jobs is None and "background_jobs" in self.__fields_set__:
            _dict['background-jobs'] = None

        # set to None if bugs (nullable) is None
        # and __fields_set__ contains the field
        if self.bugs is None and "bugs" in self.__fields_set__:
            _dict['bugs'] = None

        # set to None if category (nullable) is None
        # and __fields_set__ contains the field
        if self.category is None and "category" in self.__fields_set__:
            _dict['category'] = None

        # set to None if collaboration (nullable) is None
        # and __fields_set__ contains the field
        if self.collaboration is None and "collaboration" in self.__fields_set__:
            _dict['collaboration'] = None

        # set to None if commands (nullable) is None
        # and __fields_set__ contains the field
        if self.commands is None and "commands" in self.__fields_set__:
            _dict['commands'] = None

        # set to None if default_enable (nullable) is None
        # and __fields_set__ contains the field
        if self.default_enable is None and "default_enable" in self.__fields_set__:
            _dict['default_enable'] = None

        # set to None if dependencies (nullable) is None
        # and __fields_set__ contains the field
        if self.dependencies is None and "dependencies" in self.__fields_set__:
            _dict['dependencies'] = None

        # set to None if discussion (nullable) is None
        # and __fields_set__ contains the field
        if self.discussion is None and "discussion" in self.__fields_set__:
            _dict['discussion'] = None

        # set to None if documentation (nullable) is None
        # and __fields_set__ contains the field
        if self.documentation is None and "documentation" in self.__fields_set__:
            _dict['documentation'] = None

        # set to None if groups (nullable) is None
        # and __fields_set__ contains the field
        if self.groups is None and "groups" in self.__fields_set__:
            _dict['groups'] = None

        # set to None if info (nullable) is None
        # and __fields_set__ contains the field
        if self.info is None and "info" in self.__fields_set__:
            _dict['info'] = None

        # set to None if internal (nullable) is None
        # and __fields_set__ contains the field
        if self.internal is None and "internal" in self.__fields_set__:
            _dict['internal'] = None

        # set to None if level (nullable) is None
        # and __fields_set__ contains the field
        if self.level is None and "level" in self.__fields_set__:
            _dict['level'] = None

        # set to None if licence (nullable) is None
        # and __fields_set__ contains the field
        if self.licence is None and "licence" in self.__fields_set__:
            _dict['licence'] = None

        # set to None if namespace (nullable) is None
        # and __fields_set__ contains the field
        if self.namespace is None and "namespace" in self.__fields_set__:
            _dict['namespace'] = None

        # set to None if navigations (nullable) is None
        # and __fields_set__ contains the field
        if self.navigations is None and "navigations" in self.__fields_set__:
            _dict['navigations'] = None

        # set to None if preview (nullable) is None
        # and __fields_set__ contains the field
        if self.preview is None and "preview" in self.__fields_set__:
            _dict['preview'] = None

        # set to None if preview_as_icon (nullable) is None
        # and __fields_set__ contains the field
        if self.preview_as_icon is None and "preview_as_icon" in self.__fields_set__:
            _dict['previewAsIcon'] = None

        # set to None if public (nullable) is None
        # and __fields_set__ contains the field
        if self.public is None and "public" in self.__fields_set__:
            _dict['public'] = None

        # set to None if remote (nullable) is None
        # and __fields_set__ contains the field
        if self.remote is None and "remote" in self.__fields_set__:
            _dict['remote'] = None

        # set to None if removable (nullable) is None
        # and __fields_set__ contains the field
        if self.removable is None and "removable" in self.__fields_set__:
            _dict['removable'] = None

        # set to None if repair_steps (nullable) is None
        # and __fields_set__ contains the field
        if self.repair_steps is None and "repair_steps" in self.__fields_set__:
            _dict['repair-steps'] = None

        # set to None if repository (nullable) is None
        # and __fields_set__ contains the field
        if self.repository is None and "repository" in self.__fields_set__:
            _dict['repository'] = None

        # set to None if sabre (nullable) is None
        # and __fields_set__ contains the field
        if self.sabre is None and "sabre" in self.__fields_set__:
            _dict['sabre'] = None

        # set to None if screenshot (nullable) is None
        # and __fields_set__ contains the field
        if self.screenshot is None and "screenshot" in self.__fields_set__:
            _dict['screenshot'] = None

        # set to None if settings (nullable) is None
        # and __fields_set__ contains the field
        if self.settings is None and "settings" in self.__fields_set__:
            _dict['settings'] = None

        # set to None if trash (nullable) is None
        # and __fields_set__ contains the field
        if self.trash is None and "trash" in self.__fields_set__:
            _dict['trash'] = None

        # set to None if two_factor_providers (nullable) is None
        # and __fields_set__ contains the field
        if self.two_factor_providers is None and "two_factor_providers" in self.__fields_set__:
            _dict['two-factor-providers'] = None

        # set to None if types (nullable) is None
        # and __fields_set__ contains the field
        if self.types is None and "types" in self.__fields_set__:
            _dict['types'] = None

        # set to None if versions (nullable) is None
        # and __fields_set__ contains the field
        if self.versions is None and "versions" in self.__fields_set__:
            _dict['versions'] = None

        # set to None if website (nullable) is None
        # and __fields_set__ contains the field
        if self.website is None and "website" in self.__fields_set__:
            _dict['website'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: dict) -> ProvisioningApiAppInfo:
        """Create an instance of ProvisioningApiAppInfo from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return ProvisioningApiAppInfo.parse_obj(obj)

        _obj = ProvisioningApiAppInfo.parse_obj({
            "active": obj.get("active"),
            "activity": obj.get("activity"),
            "author": obj.get("author"),
            "background_jobs": obj.get("background-jobs"),
            "bugs": obj.get("bugs"),
            "category": obj.get("category"),
            "collaboration": obj.get("collaboration"),
            "commands": obj.get("commands"),
            "default_enable": obj.get("default_enable"),
            "dependencies": obj.get("dependencies"),
            "description": obj.get("description"),
            "discussion": obj.get("discussion"),
            "documentation": obj.get("documentation"),
            "groups": obj.get("groups"),
            "id": obj.get("id"),
            "info": obj.get("info"),
            "internal": obj.get("internal"),
            "level": obj.get("level"),
            "licence": obj.get("licence"),
            "name": obj.get("name"),
            "namespace": obj.get("namespace"),
            "navigations": obj.get("navigations"),
            "preview": obj.get("preview"),
            "preview_as_icon": obj.get("previewAsIcon"),
            "public": obj.get("public"),
            "remote": obj.get("remote"),
            "removable": obj.get("removable"),
            "repair_steps": obj.get("repair-steps"),
            "repository": obj.get("repository"),
            "sabre": obj.get("sabre"),
            "screenshot": obj.get("screenshot"),
            "settings": obj.get("settings"),
            "summary": obj.get("summary"),
            "trash": obj.get("trash"),
            "two_factor_providers": obj.get("two-factor-providers"),
            "types": obj.get("types"),
            "version": obj.get("version"),
            "versions": obj.get("versions"),
            "website": obj.get("website")
        })
        return _obj

