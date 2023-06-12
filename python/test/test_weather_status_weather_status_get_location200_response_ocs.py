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
from openapi_client.models.weather_status_weather_status_get_location200_response_ocs import WeatherStatusWeatherStatusGetLocation200ResponseOcs  # noqa: E501
from openapi_client.rest import ApiException

class TestWeatherStatusWeatherStatusGetLocation200ResponseOcs(unittest.TestCase):
    """WeatherStatusWeatherStatusGetLocation200ResponseOcs unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test WeatherStatusWeatherStatusGetLocation200ResponseOcs
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `WeatherStatusWeatherStatusGetLocation200ResponseOcs`
        """
        model = openapi_client.models.weather_status_weather_status_get_location200_response_ocs.WeatherStatusWeatherStatusGetLocation200ResponseOcs()  # noqa: E501
        if include_optional :
            return WeatherStatusWeatherStatusGetLocation200ResponseOcs(
                meta = openapi_client.models.ocs_meta.OCSMeta(
                    status = '', 
                    statuscode = 56, 
                    message = '', 
                    totalitems = '', 
                    itemsperpage = '', ), 
                data = openapi_client.models.weather_status_weather_status_get_location_200_response_ocs_data.weather_status_weather_status_get_location_200_response_ocs_data(
                    lat = 1.337, 
                    lon = 1.337, 
                    address = '', 
                    mode = 56, )
            )
        else :
            return WeatherStatusWeatherStatusGetLocation200ResponseOcs(
                meta = openapi_client.models.ocs_meta.OCSMeta(
                    status = '', 
                    statuscode = 56, 
                    message = '', 
                    totalitems = '', 
                    itemsperpage = '', ),
                data = openapi_client.models.weather_status_weather_status_get_location_200_response_ocs_data.weather_status_weather_status_get_location_200_response_ocs_data(
                    lat = 1.337, 
                    lon = 1.337, 
                    address = '', 
                    mode = 56, ),
        )
        """

    def testWeatherStatusWeatherStatusGetLocation200ResponseOcs(self):
        """Test WeatherStatusWeatherStatusGetLocation200ResponseOcs"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
