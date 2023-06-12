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
from openapi_client.api.core_wipe_api import CoreWipeApi  # noqa: E501
from openapi_client.rest import ApiException


class TestCoreWipeApi(unittest.TestCase):
    """CoreWipeApi unit test stubs"""

    def setUp(self):
        self.api = openapi_client.api.core_wipe_api.CoreWipeApi()  # noqa: E501

    def tearDown(self):
        pass

    def test_core_wipe_check_wipe(self):
        """Test case for core_wipe_check_wipe

        Check if the device should be wiped  # noqa: E501
        """
        pass

    def test_core_wipe_wipe_done(self):
        """Test case for core_wipe_wipe_done

        Finish the wipe  # noqa: E501
        """
        pass


if __name__ == '__main__':
    unittest.main()
