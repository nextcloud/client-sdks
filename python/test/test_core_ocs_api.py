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
from openapi_client.api.core_ocs_api import CoreOcsApi  # noqa: E501
from openapi_client.rest import ApiException


class TestCoreOcsApi(unittest.TestCase):
    """CoreOcsApi unit test stubs"""

    def setUp(self):
        self.api = openapi_client.api.core_ocs_api.CoreOcsApi()  # noqa: E501

    def tearDown(self):
        pass

    def test_core_ocs_get_capabilities(self):
        """Test case for core_ocs_get_capabilities

        Get the capabilities  # noqa: E501
        """
        pass


if __name__ == '__main__':
    unittest.main()