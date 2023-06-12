# coding: utf-8

"""
    nextcloud

    Nextcloud APIs  # noqa: E501

    The version of the OpenAPI document: 0.0.1
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""


import unittest

import openapi_client
from openapi_client.api.theming_icon_api import ThemingIconApi  # noqa: E501
from openapi_client.rest import ApiException


class TestThemingIconApi(unittest.TestCase):
    """ThemingIconApi unit test stubs"""

    def setUp(self):
        self.api = openapi_client.api.theming_icon_api.ThemingIconApi()  # noqa: E501

    def tearDown(self):
        pass

    def test_theming_icon_get_favicon(self):
        """Test case for theming_icon_get_favicon

        Return a 32x32 favicon as png  # noqa: E501
        """
        pass

    def test_theming_icon_get_themed_icon(self):
        """Test case for theming_icon_get_themed_icon

        Get a themed icon  # noqa: E501
        """
        pass

    def test_theming_icon_get_touch_icon(self):
        """Test case for theming_icon_get_touch_icon

        Return a 512x512 icon for touch devices  # noqa: E501
        """
        pass


if __name__ == '__main__':
    unittest.main()
