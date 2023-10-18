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
from openapi_client.models.weather_status_forecast_data_next1_hours import WeatherStatusForecastDataNext1Hours  # noqa: E501
from openapi_client.rest import ApiException

class TestWeatherStatusForecastDataNext1Hours(unittest.TestCase):
    """WeatherStatusForecastDataNext1Hours unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test WeatherStatusForecastDataNext1Hours
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `WeatherStatusForecastDataNext1Hours`
        """
        model = openapi_client.models.weather_status_forecast_data_next1_hours.WeatherStatusForecastDataNext1Hours()  # noqa: E501
        if include_optional :
            return WeatherStatusForecastDataNext1Hours(
                summary = openapi_client.models.weather_status_forecast_data_next_12_hours_summary.WeatherStatusForecast_data_next_12_hours_summary(
                    symbol_code = '', ), 
                details = openapi_client.models.weather_status_forecast_data_next_1_hours_details.WeatherStatusForecast_data_next_1_hours_details(
                    precipitation_amount = 1.337, 
                    precipitation_amount_max = 1.337, 
                    precipitation_amount_min = 1.337, 
                    probability_of_precipitation = 1.337, 
                    probability_of_thunder = 1.337, )
            )
        else :
            return WeatherStatusForecastDataNext1Hours(
                summary = openapi_client.models.weather_status_forecast_data_next_12_hours_summary.WeatherStatusForecast_data_next_12_hours_summary(
                    symbol_code = '', ),
                details = openapi_client.models.weather_status_forecast_data_next_1_hours_details.WeatherStatusForecast_data_next_1_hours_details(
                    precipitation_amount = 1.337, 
                    precipitation_amount_max = 1.337, 
                    precipitation_amount_min = 1.337, 
                    probability_of_precipitation = 1.337, 
                    probability_of_thunder = 1.337, ),
        )
        """

    def testWeatherStatusForecastDataNext1Hours(self):
        """Test WeatherStatusForecastDataNext1Hours"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()