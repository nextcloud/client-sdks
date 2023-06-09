# coding: utf-8

"""
    nextcloud

    Nextcloud APIs  # noqa: E501

    The version of the OpenAPI document: 0.0.1
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""


import unittest
import datetime

import openapi_client
from openapi_client.models.weather_status_weather_status_use_personal_address200_response_ocs import WeatherStatusWeatherStatusUsePersonalAddress200ResponseOcs  # noqa: E501
from openapi_client.rest import ApiException

class TestWeatherStatusWeatherStatusUsePersonalAddress200ResponseOcs(unittest.TestCase):
    """WeatherStatusWeatherStatusUsePersonalAddress200ResponseOcs unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test WeatherStatusWeatherStatusUsePersonalAddress200ResponseOcs
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `WeatherStatusWeatherStatusUsePersonalAddress200ResponseOcs`
        """
        model = openapi_client.models.weather_status_weather_status_use_personal_address200_response_ocs.WeatherStatusWeatherStatusUsePersonalAddress200ResponseOcs()  # noqa: E501
        if include_optional :
            return WeatherStatusWeatherStatusUsePersonalAddress200ResponseOcs(
                meta = openapi_client.models.ocs_meta.OCSMeta(
                    status = '', 
                    statuscode = 56, 
                    message = '', 
                    totalitems = '', 
                    itemsperpage = '', ), 
                data = openapi_client.models.weather_status_weather_status_use_personal_address_200_response_ocs_data.weather_status_weather_status_use_personal_address_200_response_ocs_data(
                    success = True, 
                    lat = 1.337, 
                    lon = 1.337, 
                    address = '', )
            )
        else :
            return WeatherStatusWeatherStatusUsePersonalAddress200ResponseOcs(
                meta = openapi_client.models.ocs_meta.OCSMeta(
                    status = '', 
                    statuscode = 56, 
                    message = '', 
                    totalitems = '', 
                    itemsperpage = '', ),
                data = openapi_client.models.weather_status_weather_status_use_personal_address_200_response_ocs_data.weather_status_weather_status_use_personal_address_200_response_ocs_data(
                    success = True, 
                    lat = 1.337, 
                    lon = 1.337, 
                    address = '', ),
        )
        """

    def testWeatherStatusWeatherStatusUsePersonalAddress200ResponseOcs(self):
        """Test WeatherStatusWeatherStatusUsePersonalAddress200ResponseOcs"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
