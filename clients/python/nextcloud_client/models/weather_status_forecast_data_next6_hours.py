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
from nextcloud_client.models.weather_status_forecast_data_next12_hours_summary import WeatherStatusForecastDataNext12HoursSummary
from nextcloud_client.models.weather_status_forecast_data_next6_hours_details import WeatherStatusForecastDataNext6HoursDetails

class WeatherStatusForecastDataNext6Hours(BaseModel):
    """
    WeatherStatusForecastDataNext6Hours
    """
    summary: WeatherStatusForecastDataNext12HoursSummary = Field(...)
    details: WeatherStatusForecastDataNext6HoursDetails = Field(...)
    additional_properties: Dict[str, Any] = {}
    __properties = ["summary", "details"]

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
    def from_json(cls, json_str: str) -> WeatherStatusForecastDataNext6Hours:
        """Create an instance of WeatherStatusForecastDataNext6Hours from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self):
        """Returns the dictionary representation of the model using alias"""
        _dict = self.dict(by_alias=True,
                          exclude={
                            "additional_properties"
                          },
                          exclude_none=True)
        # override the default output from pydantic by calling `to_dict()` of summary
        if self.summary:
            _dict['summary'] = self.summary.to_dict()
        # override the default output from pydantic by calling `to_dict()` of details
        if self.details:
            _dict['details'] = self.details.to_dict()
        # puts key-value pairs in additional_properties in the top level
        if self.additional_properties is not None:
            for _key, _value in self.additional_properties.items():
                _dict[_key] = _value

        return _dict

    @classmethod
    def from_dict(cls, obj: dict) -> WeatherStatusForecastDataNext6Hours:
        """Create an instance of WeatherStatusForecastDataNext6Hours from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return WeatherStatusForecastDataNext6Hours.parse_obj(obj)

        _obj = WeatherStatusForecastDataNext6Hours.parse_obj({
            "summary": WeatherStatusForecastDataNext12HoursSummary.from_dict(obj.get("summary")) if obj.get("summary") is not None else None,
            "details": WeatherStatusForecastDataNext6HoursDetails.from_dict(obj.get("details")) if obj.get("details") is not None else None
        })
        # store additional fields in additional_properties
        for _key in obj.keys():
            if _key not in cls.__properties:
                _obj.additional_properties[_key] = obj.get(_key)

        return _obj


