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
from openapi_client.api.dav_direct_api import DavDirectApi  # noqa: E501
from openapi_client.rest import ApiException


class TestDavDirectApi(unittest.TestCase):
    """DavDirectApi unit test stubs"""

    def setUp(self):
        self.api = openapi_client.api.dav_direct_api.DavDirectApi()  # noqa: E501

    def tearDown(self):
        pass

    def test_dav_direct_get_url(self):
        """Test case for dav_direct_get_url

        Get a direct link to a file  # noqa: E501
        """
        pass


if __name__ == '__main__':
    unittest.main()
